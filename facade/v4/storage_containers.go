package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	storageContainerModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// GetStorageContainer returns the storage container for the given UUID.
func (f *FacadeV4Client) GetStorageContainer(uuid string) (*storageContainerModels.StorageContainer, error) {
	return nil, fmt.Errorf("GetStorageContainer not implemented in FacadeV4Client")
}

// ListStorageContainers returns a list of storage containers.
func (f *FacadeV4Client) ListStorageContainers(opts ...facade.ODataOption) ([]storageContainerModels.StorageContainer, error) {
	return nil, fmt.Errorf("ListStorageContainers not implemented in FacadeV4Client")
}

// ListAllStorageContainers returns all storage containers without pagination.
func (f *FacadeV4Client) ListAllStorageContainers(filterParam *string, orderbyParam *string, selectParam *string) ([]storageContainerModels.StorageContainer, error) {
	return nil, fmt.Errorf("ListAllStorageContainers not implemented in FacadeV4Client")
}

// GetListIteratorStorageContainers returns an iterator for listing storage containers.
func (f *FacadeV4Client) GetListIteratorStorageContainers(opts ...facade.ODataOption) (facade.ODataListIterator[*storageContainerModels.StorageContainer], error) {
	return nil, fmt.Errorf("GetListIteratorStorageContainers not implemented in FacadeV4Client")
}
