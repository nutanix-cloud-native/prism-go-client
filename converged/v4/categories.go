package v4

import (
	"context"
	"errors"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	prismMessages "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

// CategoriesService provides default "not implemented" implementation for all Categories interface methods.
type CategoriesService struct {
	client   *v4prismGoClient.Client
	entities string
}

// NewCategoriesService creates a new CategoriesService instance.
func NewCategoriesService(client *v4prismGoClient.Client) *CategoriesService {
	return &CategoriesService{client: client, entities: "category"}
}

// Get returns the category for the given UUID.
func (s *CategoriesService) Get(ctx context.Context, uuid string) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	return GenericGetEntity[*prismModels.GetCategoryApiResponse, prismModels.Category](
		func() (*prismModels.GetCategoryApiResponse, error) {
			return s.client.CategoriesApiInstance.GetCategoryById(&uuid, nil)
		},
		s.entities,
	)
}

// List returns a list of categories.
func (s *CategoriesService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}

	// Check if Apply option is provided and return error if it is
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}
	if reqParams != nil && reqParams.Apply != nil {
		return nil, errors.New("apply is not supported")
	}

	return GenericListEntities[*prismModels.ListCategoriesApiResponse, prismModels.Category](
		func(reqParams *V4ODataParams) (*prismModels.ListCategoriesApiResponse, error) {
			return s.client.CategoriesApiInstance.ListCategories(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Expand, reqParams.Select)
		},
		opts,
		"categories",
	)
}

// NewIterator returns an iterator for listing categories.
func (s *CategoriesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[prismModels.Category] {
	if s.client == nil {
		return nil
	}
	return GenericNewIterator[*prismModels.ListCategoriesApiResponse, prismModels.Category](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (*prismModels.ListCategoriesApiResponse, error) {
			return s.client.CategoriesApiInstance.ListCategories(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Expand, reqParams.Select)
		},
		opts,
		"categories",
	)
}

// Create creates a new category.
func (s *CategoriesService) Create(ctx context.Context, entity *prismModels.Category) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	newCategory, err := CallAPI[*prismModels.CreateCategoryApiResponse, prismModels.Category](s.client.CategoriesApiInstance.CreateCategory(entity))
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return &newCategory, nil
}

// Update updates an existing category.
func (s *CategoriesService) Update(ctx context.Context, uuid string, entity *prismModels.Category) (*prismModels.Category, error) {
	if s.client == nil {
		return nil, errors.New("client is not initialized")
	}
	existingCategory, args, err := GetEntityAndEtag(
		s.Get(ctx, uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get category with UUID %s: %w", uuid, err)
	}
	if existingCategory == nil {
		return nil, fmt.Errorf("no category found with UUID %s", uuid)
	}

	_, err = CallAPI[*prismModels.UpdateCategoryApiResponse, []prismMessages.AppMessage](
		s.client.CategoriesApiInstance.UpdateCategoryById(&uuid, entity, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update category with UUID %s: %w", uuid, err)
	}

	updatedCategory, err := s.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated category with UUID %s: %w", uuid, err)
	}
	if updatedCategory == nil {
		return nil, fmt.Errorf("no updated category found with UUID %s", uuid)
	}
	if updatedCategory.ExtId == nil {
		return nil, fmt.Errorf("updated category ExtId is nil for UUID %s", uuid)
	}
	return updatedCategory, nil

}

// Delete deletes an existing category.
func (s *CategoriesService) Delete(ctx context.Context, uuid string) error {
	if s.client == nil {
		return errors.New("client is not initialized")
	}
	category, args, err := GetEntityAndEtag(
		s.Get(ctx, uuid),
	)
	if err != nil {
		return fmt.Errorf("failed to get category with UUID %s: %w", uuid, err)
	}
	if category == nil {
		return fmt.Errorf("no category found with UUID %s", uuid)
	}

	_, err = s.client.CategoriesApiInstance.DeleteCategoryById(&uuid, args)
	if err != nil {
		return fmt.Errorf("failed to delete category with UUID %s: %w", uuid, err)
	}

	return nil
}
