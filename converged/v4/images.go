package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"

	vmmConfig "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	imageModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// ImagesService implements the converged Images interface for the Prism Central V4 API.
type ImagesService struct {
	client       *v4prismGoClient.Client
	entitiesName string
}

// NewImagesService returns a new ImagesService for the given V4 client.
func NewImagesService(client *v4prismGoClient.Client) *ImagesService {
	return &ImagesService{client: client, entitiesName: "image"}
}

// Get returns the image for the given UUID.
func (s *ImagesService) Get(ctx context.Context, uuid string) (*imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	return GenericGetEntity[*imageModels.GetImageApiResponse, imageModels.Image](
		func() (*imageModels.GetImageApiResponse, error) {
			return s.client.ImagesApiInstance.GetImageById(&uuid)
		},
		s.entitiesName,
	)
}

// List returns a list of images.
func (s *ImagesService) List(ctx context.Context, opts ...converged.ODataOption) ([]imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	myParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if myParams.Expand != nil || myParams.Apply != nil {
		return nil, fmt.Errorf("expand and apply are not supported for listing Images")
	}

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
		s.entitiesName,
	)
}

// NewIterator returns an iterator for listing images.
func (s *ImagesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[imageModels.Image] {
	if s.client == nil {
		return nil
	}

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
		s.entitiesName,
	)
}

// Create creates a new Image.
func (s *ImagesService) Create(ctx context.Context, image *imageModels.Image) (*imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.CreateAsync(ctx, image)
	if err != nil {
		return nil, fmt.Errorf("failed to create Image: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Image: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to create Image: operation completed but no Image returned")
	}

	return result[0], nil
}

// Update updates an existing Image.
func (s *ImagesService) Update(ctx context.Context, uuid string, entity *imageModels.Image) (*imageModels.Image, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	operation, err := s.UpdateAsync(ctx, uuid, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to update Image: %w", err)
	}

	result, err := operation.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update Image: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("failed to update Image: operation completed but no Image returned")
	}

	return result[0], nil
}

// Delete deletes an existing Image.
func (s *ImagesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	operation, err := s.DeleteAsync(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to delete Image: %w", err)
	}

	_, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete Image: %w", err)
	}

	return nil
}

// CreateAsync creates a new Image asynchronously.
func (s *ImagesService) CreateAsync(ctx context.Context, image *imageModels.Image) (converged.Operation[imageModels.Image], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	taskRef, err := CallAPI[*imageModels.CreateImageApiResponse, vmmConfig.TaskReference](
		s.client.ImagesApiInstance.CreateImage(image),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Image: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for created Image")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// UpdateAsync updates an existing Image asynchronously.
func (s *ImagesService) UpdateAsync(ctx context.Context, uuid string, image *imageModels.Image) (converged.Operation[imageModels.Image], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	currentImage, args, err := GetEntityAndEtag(
		s.client.ImagesApiInstance.GetImageById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get Image for update: %w", err)
	}

	image = CopyEtag(currentImage, image).(*imageModels.Image)

	taskRef, err := CallAPI[*imageModels.UpdateImageApiResponse, vmmConfig.TaskReference](
		s.client.ImagesApiInstance.UpdateImageById(&uuid, image, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update Image: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for updated Image")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		s.Get,
	), nil
}

// DeleteAsync deletes an existing Image asynchronously.
func (s *ImagesService) DeleteAsync(ctx context.Context, uuid string) (converged.Operation[converged.NoEntity], error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Get the Image first to retrieve its eTag
	_, args, err := GetEntityAndEtag(
		s.client.ImagesApiInstance.GetImageById(&uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get Image for deletion: %w", err)
	}

	taskRef, err := CallAPI[*imageModels.DeleteImageApiResponse, vmmConfig.TaskReference](
		s.client.ImagesApiInstance.DeleteImageById(&uuid, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete Image: %w", err)
	}

	if taskRef.ExtId == nil {
		return nil, fmt.Errorf("task reference ExtId is nil for deleted Image")
	}

	return NewOperation(
		*taskRef.ExtId,
		s.client,
		func(ctx context.Context, uuid string) (*converged.NoEntity, error) {
			return converged.NoEntityGetter(ctx, uuid)
		},
	), nil
}
