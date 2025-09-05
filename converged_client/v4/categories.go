package v4

import (
	"context"
	"fmt"

	converged "github.com/nutanix-cloud-native/prism-go-client/converged_client"
	prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// CategoriesService provides default "not implemented" implementation for all Categories interface methods.
type CategoriesService struct{}

// NewCategoriesService creates a new CategoriesService instance.
func NewCategoriesService() *CategoriesService {
	return &CategoriesService{}
}

// Get returns the category for the given UUID.
func (s *CategoriesService) Get(ctx context.Context, uuid string) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// List returns a list of categories.
func (s *CategoriesService) List(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListAll returns all categories without pagination.
func (s *CategoriesService) ListAll(ctx context.Context, opts ...converged.ODataOption) ([]prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// NewIterator returns an iterator for listing categories.
func (s *CategoriesService) NewIterator(opts ...converged.ODataOption) converged.Iterator[prismModels.Category] {
	return nil
}

// Create creates a new category.
func (s *CategoriesService) Create(ctx context.Context, entity *prismModels.Category) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update updates an existing category.
func (s *CategoriesService) Update(ctx context.Context, uuid string, entity *prismModels.Category) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete deletes an existing category.
func (s *CategoriesService) Delete(ctx context.Context, uuid string) (*prismModels.Category, error) {
	return nil, fmt.Errorf("not implemented")
}
