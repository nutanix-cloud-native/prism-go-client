package v4

import (
	"context"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// ImagesService provides default "not implemented" implementation for all Images interface methods.
type ImagesService struct {
	client *v4prismGoClient.Client
}

// NewImagesService creates a new ImagesService instance.
func NewImagesService(client *v4prismGoClient.Client) *ImagesService {
	return &ImagesService{client: client}
}

// Get returns the image for the given UUID.
func (s *ImagesService) Get(ctx context.Context, uuid string) (*imageModels.Image, error) {
	return GenericGetEntity[*imageModels.GetImageApiResponse, imageModels.Image](
		func() (*imageModels.GetImageApiResponse, error) {
			return s.client.ImagesApiInstance.GetImageById(&uuid)
		},
		"image",
	)
}

// List returns a list of images.
func (s *ImagesService) List(ctx context.Context, opts ...converged.ODataOption) ([]imageModels.Image, error) {
	return GenericListEntities[*imageModels.ListImagesApiResponse, imageModels.Image](
		func(reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return s.client.ImagesApiInstance.ListImages(
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

// NewIterator returns an iterator for listing images.
func (s *ImagesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[imageModels.Image] {
	return GenericNewIterator[*imageModels.ListImagesApiResponse, imageModels.Image](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*imageModels.ListImagesApiResponse, error) {
			return s.client.ImagesApiInstance.ListImages(
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
