package facade

import (
	"context"

	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type CategoriesFacadeV4 interface {
	// GetCategory returns the category for the given UUID.
	GetCategory(ctx context.Context, uuid string) (*v4prismModels.Category, error)

	// ListCategories returns a list of categories.
	ListCategories(ctx context.Context, opts ...ODataOption) ([]v4prismModels.Category, error)

	// ListAllCategories returns all categories without pagination.
	ListAllCategories(ctx context.Context, filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]v4prismModels.Category, error)

	// GetListIteratorCategories returns an iterator for listing categories.
	GetListIteratorCategories(ctx context.Context, opts ...ODataOption) ODataListIterator[v4prismModels.Category]

	// CreateCategory creates a new category.
	CreateCategory(ctx context.Context, category *v4prismModels.Category) (*v4prismModels.Category, error)

	// UpdateCategory updates an existing category.
	UpdateCategory(ctx context.Context, uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error)

	// DeleteCategory deletes the category with the given UUID.
	DeleteCategory(ctx context.Context, uuid string) error
}
