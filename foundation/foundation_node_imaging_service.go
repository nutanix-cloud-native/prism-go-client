package foundation

import (
	"context"
	"net/http"

	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

type NodeImagingService interface {
	ImageNodes(context.Context, *ImageNodesInput) (*ImageNodesAPIResponse, error)
	ImageNodesProgress(context.Context, string) (*ImageNodesProgressResponse, error)
}

type NodeImagingOperations struct {
	client *internal.Client
}

func (op NodeImagingOperations) ImageNodes(ctx context.Context, imageNodeInput *ImageNodesInput) (*ImageNodesAPIResponse, error) {
	path := "/image_nodes"
	req, err := op.client.NewUnAuthRequest(http.MethodPost, path, imageNodeInput)
	if err != nil {
		return nil, err
	}

	imageNodesAPIResponse := new(ImageNodesAPIResponse)
	return imageNodesAPIResponse, op.client.Do(ctx, req, imageNodesAPIResponse)
}

// Gets progress of imaging session.
func (op NodeImagingOperations) ImageNodesProgress(ctx context.Context, sessionID string) (*ImageNodesProgressResponse, error) {
	path := "/progress?session_id=" + sessionID
	req, err := op.client.NewUnAuthRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	imageNodesProgressResponse := new(ImageNodesProgressResponse)
	return imageNodesProgressResponse, op.client.Do(ctx, req, imageNodesProgressResponse)
}
