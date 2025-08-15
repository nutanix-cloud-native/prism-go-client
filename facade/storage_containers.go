package facade

import (
	scmodels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// StorageContainersFacadeV4 defines the interface for managing storage containers.
type StorageContainersFacadeV4 interface {
	// GetStorageContainer returns the storage container for the given UUID.
	GetStorageContainer(uuid string) (*scmodels.StorageContainer, error)

	// ListStorageContainers returns a list of storage containers.
	ListStorageContainers(opts ...ODataOption) ([]scmodels.StorageContainer, error)

	// ListAllStorageContainers returns all storage containers without pagination.
	ListAllStorageContainers(filterParam *string, orderbyParam *string, selectParam *string) ([]scmodels.StorageContainer, error)

	// GetListIteratorStorageContainers returns an iterator for listing storage containers.
	GetListIteratorStorageContainers(opts ...ODataOption) ODataListIterator[scmodels.StorageContainer]
}
