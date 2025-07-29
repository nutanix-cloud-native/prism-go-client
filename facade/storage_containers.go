package facade

import (
	storageContainerModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// StorageContainersFacadeV4 defines the interface for managing storage containers.
type StorageContainersFacadeV4 interface {
	// GetStorageContainer returns the storage container for the given UUID.
	GetStorageContainer(uuid string) (*storageContainerModels.StorageContainer, error)

	// ListStorageContainers returns a list of storage containers.
	ListStorageContainers(opts ...ODataOption) ([]storageContainerModels.StorageContainer, error)

	// ListAllStorageContainers returns all storage containers without pagination.
	ListAllStorageContainers(filterParam *string, orderbyParam *string, selectParam *string) ([]storageContainerModels.StorageContainer, error)

	// GetListIteratorStorageContainers returns an iterator for listing storage containers.
	GetListIteratorStorageContainers(opts ...ODataOption) (ODataListIterator[storageContainerModels.StorageContainer], error)
}
