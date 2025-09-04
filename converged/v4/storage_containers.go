package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	scModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// StorageContainersService provides default "not implemented" implementation for all StorageContainers interface methods.
type StorageContainersService struct {
	client *v4prismGoClient.Client
}

// NewStorageContainersService creates a new StorageContainersService instance.
func NewStorageContainersService(client *v4prismGoClient.Client) *StorageContainersService {
	return &StorageContainersService{client: client}
}

// Get returns the storage container for the given UUID.
func (s *StorageContainersService) Get(ctx context.Context, uuid string) (*scModels.StorageContainer, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of storage containers.
func (s *StorageContainersService) List(ctx context.Context, opts ...converged.ODataOption) ([]scModels.StorageContainer, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing storage containers.
func (s *StorageContainersService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[scModels.StorageContainer] {
	return nil
}
