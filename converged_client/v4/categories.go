package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// CategoriesService provides implementation for all Categories interface methods.
type CategoriesService struct {
	client *v4prismGoClient.Client
}

// NewCategoriesService creates a new CategoriesService instance.
func NewCategoriesService(client *v4prismGoClient.Client) *CategoriesService {
	return &CategoriesService{
		client: client,
	}
}

// Get returns the category for the given UUID.
func (s *CategoriesService) Get(ctx context.Context, uuid string) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericGetEntity[
		*prismModels.GetCategoryApiResponse,
		prismModels.Category,
	](
		func() (*prismModels.GetCategoryApiResponse, error) {
			return s.client.CategoriesApiInstance.GetCategoryById(&uuid, nil, nil)
		},
		"category",
	)
}

// List returns a list of categories.
func (s *CategoriesService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	return GenericListEntities[
		*prismModels.ListCategoriesApiResponse,
		prismModels.Category,
	](
		func(reqParams *V4ODataParams) (*prismModels.ListCategoriesApiResponse, error) {
			return s.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		opts,
		"categories",
	)
}

// ListAll returns all categories without pagination.
func (s *CategoriesService) ListAll(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	return GenericListAllEntities[
		*prismModels.ListCategoriesApiResponse,
		prismModels.Category,
	](
		func(reqParams *V4ODataParams) (*prismModels.ListCategoriesApiResponse, error) {
			return s.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		reqParams,
		"categories",
	)
}

// NewIterator returns an iterator for listing categories.
func (s *CategoriesService) NewIterator(opts ...converged.ODataOption) converged.Iterator[prismModels.Category] {
	if s.client == nil {
		return nil
	}

	return NewIterator[
		*prismModels.ListCategoriesApiResponse,
		prismModels.Category,
	](
		context.Background(),
		s.client,
		func(ctx context.Context, reqParams *V4ODataParams) (*prismModels.ListCategoriesApiResponse, error) {
			return s.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
				nil,
			)
		},
		opts...,
	)
}

// Create creates a new category.
func (s *CategoriesService) Create(ctx context.Context, entity *prismModels.Category) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	if entity == nil {
		return nil, fmt.Errorf("entity cannot be nil")
	}

	return GenericGetEntity[
		*prismModels.CreateCategoryApiResponse,
		prismModels.Category,
	](
		func() (*prismModels.CreateCategoryApiResponse, error) {
			return s.client.CategoriesApiInstance.CreateCategory(entity, nil)
		},
		"category",
	)
}

// Update updates an existing category.
func (s *CategoriesService) Update(ctx context.Context, uuid string, entity *prismModels.Category) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	if entity == nil {
		return nil, fmt.Errorf("entity cannot be nil")
	}

	// Get the current entity to obtain the ETag
	current, args, err := GetEntityAndEtag(s.client.CategoriesApiInstance.GetCategoryById(&uuid, nil, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to get current category for update: %w", err)
	}

	// Copy ETag to the entity being updated
	entity = CopyEtag(current, entity).(*prismModels.Category)

	return GenericGetEntity[
		*prismModels.UpdateCategoryApiResponse,
		prismModels.Category,
	](
		func() (*prismModels.UpdateCategoryApiResponse, error) {
			return s.client.CategoriesApiInstance.UpdateCategoryById(&uuid, entity, args)
		},
		"category",
	)
}

// Delete deletes an existing category.
func (s *CategoriesService) Delete(ctx context.Context, uuid string) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	// Get the current entity before deletion
	currentResponse, args, err := GetEntityAndEtag(s.client.CategoriesApiInstance.GetCategoryById(&uuid, nil, nil))
	if err != nil {
		return nil, fmt.Errorf("failed to get current category for deletion: %w", err)
	}

	// Extract the category from the response
	current, err := CallAPI[*prismModels.GetCategoryApiResponse, prismModels.Category](currentResponse, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to extract category from response: %w", err)
	}

	// Perform the deletion
	_, err = s.client.CategoriesApiInstance.DeleteCategoryById(&uuid, args)
	if err != nil {
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	// Return the deleted entity
	return &current, nil
}
