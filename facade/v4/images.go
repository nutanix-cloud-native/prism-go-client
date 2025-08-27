package v4

import (
	"github.com/nutanix-cloud-native/prism-go-client/facade"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// GetImage returns the image for the given UUID.
func (f *FacadeV4Client) GetImage(uuid string) (*imageModels.Image, error) {
	return CommonGetEntity[*imageModels.GetImageApiResponse, imageModels.Image](
		func() (*imageModels.GetImageApiResponse, error) {
			return f.client.ImagesApiInstance.GetImageById(&uuid)
		},
		"image",
	)
}

// ListImages returns a list of images.
func (f *FacadeV4Client) ListImages(opts ...facade.ODataOption) ([]imageModels.Image, error) {
	return CommonListEntities[*imageModels.ListImagesApiResponse, imageModels.Image](
		func(reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return f.client.ImagesApiInstance.ListImages(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		"images",
	)
}

// ListAllImages returns all images without pagination.
func (f *FacadeV4Client) ListAllImages(filterParam *string, orderbyParam *string, selectParam *string) ([]imageModels.Image, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*imageModels.ListImagesApiResponse, imageModels.Image](
		func(reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return f.client.ImagesApiInstance.ListImages(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		reqParams,
		"images",
	)
}

// GetListIteratorImages returns an iterator for listing images.
func (f *FacadeV4Client) GetListIteratorImages(opts ...facade.ODataOption) facade.ODataListIterator[imageModels.Image] {
	return CommonGetListIterator[*imageModels.ListImagesApiResponse, imageModels.Image](
		f,
		func(reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return f.client.ImagesApiInstance.ListImages(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Select,
			)
		},
		opts,
		"images",
	)
}
