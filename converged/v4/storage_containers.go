package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	scModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// StorageContainersService provides implementation for all StorageContainers interface methods.
type StorageContainersService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewStorageContainersService creates a new StorageContainersService instance.
func NewStorageContainersService(client *v4prismGoClient.Client) *StorageContainersService {
	return &StorageContainersService{client: client, entitiesName: "storage container"}
}

// Get returns the storage container for the given UUID.
func (s *StorageContainersService) Get(ctx context.Context, uuid string) (*scModels.StorageContainer, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*scModels.GetStorageContainerApiResponse, scModels.StorageContainer](
		func() (*scModels.GetStorageContainerApiResponse, error) {
			return s.client.StorageContainerAPI.GetStorageContainerById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of storage containers.
func (s *StorageContainersService) List(ctx context.Context, opts ...converged.ODataOption) ([]scModels.StorageContainer, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing StorageContainers")
	}

	return GenericListEntities[*scModels.ListStorageContainersApiResponse, scModels.StorageContainer](
		func(reqParams *V4ODataParams) (*scModels.ListStorageContainersApiResponse, error) {
			return s.client.StorageContainerAPI.ListStorageContainers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing storage containers.
func (s *StorageContainersService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[scModels.StorageContainer] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*scModels.ListStorageContainersApiResponse, scModels.StorageContainer](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*scModels.ListStorageContainersApiResponse, error) {
			return s.client.StorageContainerAPI.ListStorageContainers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		s.entitiesName,
	)
}
