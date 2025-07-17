package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// GetImage returns the image for the given UUID.
func (f *FacadeV4Client) GetImage(uuid string) (*imageModels.Image, error) {
	return nil, fmt.Errorf("GetImage not implemented in FacadeV4Client")
}

// ListImages returns a list of images.
func (f *FacadeV4Client) ListImages(opts ...facade.ODataOption) ([]imageModels.Image, error) {
	return nil, fmt.Errorf("ListImages not implemented in FacadeV4Client")
}

// ListAllImages returns all images without pagination.
func (f *FacadeV4Client) ListAllImages(filterParam *string, orderbyParam *string, selectParam *string) ([]imageModels.Image, error) {
	return nil, fmt.Errorf("ListAllImages not implemented in FacadeV4Client")
}

// GetListIteratorImages returns an iterator for listing images.
func (f *FacadeV4Client) GetListIteratorImages(opts ...facade.ODataOption) (facade.ODataListIterator[*imageModels.Image], error) {
	return nil, fmt.Errorf("GetListIteratorImages not implemented in FacadeV4Client")
}
