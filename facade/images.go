package facade

import (
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

type ImagesFacadeV4 interface {
	// GetImage returns the image for the given UUID.
	GetImage(uuid string) (*imageModels.Image, error)

	// ListImages returns a list of images.
	ListImages(opts ...ODataOption) ([]imageModels.Image, error)

	// ListAllImages returns all images without pagination.
	ListAllImages(filterParam *string, orderbyParam *string, selectParam *string) ([]imageModels.Image, error)

	// GetListIteratorImages returns an iterator for listing images.
	GetListIteratorImages(opts ...ODataOption) (ODataListIterator[imageModels.Image], error)
}
