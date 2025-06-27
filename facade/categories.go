package facade

import (
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type CategoriesFacadeV4 interface {
	// GetCategory returns the category for the given UUID.
	GetCategory(uuid string) (*v4prismModels.Category, error)

	// ListCategories returns a list of categories.
	ListCategories(opts ...ODataOption) ([]v4prismModels.Category, error)

	// CreateCategory creates a new category.
	CreateCategory(category *v4prismModels.Category) (*v4prismModels.Category, error)

	// UpdateCategory updates an existing category.
	UpdateCategory(uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error)

	// DeleteCategory deletes the category with the given UUID.
	DeleteCategory(uuid string) error
}
