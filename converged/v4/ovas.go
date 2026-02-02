package v4

import (
	"context"
	"errors"
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	vmmPrismModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	ovaModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// OvasService provides implementation for all Ovas interface methods.
type OvasService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewOvasService creates a new OvasService instance.
func NewOvasService(client *v4prismGoClient.Client) *OvasService {
	return &OvasService{client: client, entitiesName: "ova"}
}

// Get returns the OVA for the given UUID.
func (s *OvasService) Get(ctx context.Context, uuid string) (*ovaModels.Ova, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*ovaModels.GetOvaApiResponse, ovaModels.Ova](
		func() (*ovaModels.GetOvaApiResponse, error) {
			return s.client.OvasApiInstance.GetOvaById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of OVAs.
func (s *OvasService) List(ctx context.Context, opts ...converged.ODataOption) ([]ovaModels.Ova, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing OVAs")
	}

	return GenericListEntities[*ovaModels.ListOvasApiResponse, ovaModels.Ova](
		func(reqParams *V4ODataParams) (*ovaModels.ListOvasApiResponse, error) {
			return s.client.OvasApiInstance.ListOvas(
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

// NewIterator returns an iterator for listing OVAs.
func (s *OvasService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[ovaModels.Ova] {
	if s.client == nil {
		return nil
	}

	return GenericNewIterator[*ovaModels.ListOvasApiResponse, ovaModels.Ova](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*ovaModels.ListOvasApiResponse, error) {
			return s.client.OvasApiInstance.ListOvas(
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

// Create creates a new OVA.
func (s *OvasService) Create(ctx context.Context, ova *ovaModels.Ova) (*ovaModels.Ova, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.CreateAsync(ctx, ova)
	if err != nil {
		return nil, fmt.Errorf("failed to create OVA: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create OVA: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create OVA: operation completed but no OVA returned")
	}

	return result[0], nil
}

// CreateAsync creates a new OVA asynchronously.
func (s *OvasService) CreateAsync(ctx context.Context, ova *ovaModels.Ova) (converged.Operation[ovaModels.Ova], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*ovaModels.CreateOvaApiResponse, vmmPrismModels.TaskReference](
		s.client.OvasApiInstance.CreateOva(ova),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OVA: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created OVA")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// GetFile downloads the OVA file for the given UUID.
func (s *OvasService) GetFile(ctx context.Context, uuid string) (*ovaModels.FileDetail, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	result, err := CallAPI[*ovaModels.GetOvaFileApiResponse, ovaModels.FileDetail](
		s.client.OvasApiInstance.GetFileByOvaId(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get OVA file: %w", err)
	}

	return &result, nil
}
