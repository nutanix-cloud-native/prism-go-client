package facade

import (
	"context"

	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

type ImagesFacadeV4 interface {
	// GetImage returns the image for the given UUID.
	GetImage(ctx context.Context, uuid string) (*imageModels.Image, error)

	// ListImages returns a list of images.
	ListImages(ctx context.Context, opts ...ODataOption) ([]imageModels.Image, error)

	// ListAllImages returns all images without pagination.
	ListAllImages(ctx context.Context, filterParam *string, orderbyParam *string, selectParam *string) ([]imageModels.Image, error)

	// GetListIteratorImages returns an iterator for listing images.
	GetListIteratorImages(ctx context.Context, opts ...ODataOption) ODataListIterator[imageModels.Image]
}
