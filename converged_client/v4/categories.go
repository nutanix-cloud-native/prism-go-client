package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	v4prismGoClient "github.com/nutanix-cloud-native/prism-go-client/v4"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// CategoriesService provides default "not implemented" implementation for all Categories interface methods.
type CategoriesService struct {
	client *v4prismGoClient.Client
}

// NewCategoriesService creates a new CategoriesService instance.
func NewCategoriesService(client *v4prismGoClient.Client) *CategoriesService {
	return &CategoriesService{client: client}
}

// Get returns the category for the given UUID.
func (s *CategoriesService) Get(ctx context.Context, uuid string) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of categories.
func (s *CategoriesService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing categories.
func (s *CategoriesService) NewIterator(ctx context.Context, opts ...converged.ODataOption) converged.Iterator[prismModels.Category] {
	return nil
}

// Create creates a new category.
func (s *CategoriesService) Create(ctx context.Context, entity *prismModels.Category) (*prismModels.Category, error) {
	newCategory, err := CallAPI[*prismModels.CreateCategoryApiResponse, prismModels.Category](s.client.CategoriesApiInstance.CreateCategory(entity))
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return &newCategory, nil
}

// Update updates an existing category.
func (s *CategoriesService) Update(ctx context.Context, uuid string, entity *prismModels.Category) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete deletes an existing category.
func (s *CategoriesService) Delete(ctx context.Context, uuid string) error {
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
