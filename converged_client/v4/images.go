package v4

import (
	"context"
	"fmt"

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
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of images.
func (s *ImagesService) List(ctx context.Context, opts ...converged.ODataOption) ([]imageModels.Image, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing images.
func (s *ImagesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[imageModels.Image] {
	return nil
}
