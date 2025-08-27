package v4

import (
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	storageContainerModels "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	storageContainerModelsError "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
)

// GetStorageContainer returns the storage container for the given UUID.
func (f *FacadeV4Client) GetStorageContainer(uuid string) (*storageContainerModels.StorageContainer, error) {
	return CommonGetEntity[*storageContainerModels.GetStorageContainerApiResponse, storageContainerModels.StorageContainer, *storageContainerModels.OneOfGetStorageContainerApiResponseData, *storageContainerModelsError.ErrorResponse](
		func() (*storageContainerModels.GetStorageContainerApiResponse, error) {
			return f.client.StorageContainerAPI.GetStorageContainerById(&uuid)
		},
		"storage container",
	)
}

// ListStorageContainers returns a list of storage containers.
func (f *FacadeV4Client) ListStorageContainers(opts ...facade.ODataOption) ([]storageContainerModels.StorageContainer, error) {
	return CommonListEntities[*storageContainerModels.ListStorageContainersApiResponse, storageContainerModels.StorageContainer, *storageContainerModels.OneOfListStorageContainersApiResponseData, *storageContainerModelsError.ErrorResponse](
		func(reqParams *V4ODataParams) (*storageContainerModels.ListStorageContainersApiResponse, error) {
			return f.client.StorageContainerAPI.ListStorageContainers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		"storage containers",
	)
}

// ListAllStorageContainers returns all storage containers without pagination.
func (f *FacadeV4Client) ListAllStorageContainers(filterParam *string, orderbyParam *string, selectParam *string) ([]storageContainerModels.StorageContainer, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*storageContainerModels.ListStorageContainersApiResponse, storageContainerModels.StorageContainer, *storageContainerModels.OneOfListStorageContainersApiResponseData, *storageContainerModelsError.ErrorResponse](
		func(reqParams *V4ODataParams) (*storageContainerModels.ListStorageContainersApiResponse, error) {
			return f.client.StorageContainerAPI.ListStorageContainers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		reqParams,
		"storage containers",
	)
}

// GetListIteratorStorageContainers returns an iterator for listing storage containers.
func (f *FacadeV4Client) GetListIteratorStorageContainers(opts ...facade.ODataOption) facade.ODataListIterator[storageContainerModels.StorageContainer] {
	return CommonGetListIterator[*storageContainerModels.ListStorageContainersApiResponse, storageContainerModels.StorageContainer, *storageContainerModels.OneOfListStorageContainersApiResponseData, *storageContainerModelsError.ErrorResponse](
		f,
		func(reqParams *V4ODataParams) (*storageContainerModels.ListStorageContainersApiResponse, error) {
			return f.client.StorageContainerAPI.ListStorageContainers(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		"storage containers",
	)
}
