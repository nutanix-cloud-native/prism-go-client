package facade

import (
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type CategoriesFacadeV4 interface {
	// GetCategory returns the category for the given UUID.
	GetCategory(uuid string) (*v4prismModels.Category, error)

	// ListCategories returns a list of categories. If page and limit are not provided,
	// it returns all the categories available (filtered if filter is provided).
	ListCategories(opts ...ODataOption) ([]v4prismModels.Category, error)

	// GetListIteratorCategories returns an iterator for listing categories.
	GetListIteratorCategories(opts ...ODataOption) ODataListIterator[*v4prismModels.Category]

	// CreateCategory creates a new category.
	CreateCategory(category *v4prismModels.Category) (*v4prismModels.Category, error)

	// UpdateCategory updates an existing category.
	UpdateCategory(uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error)

	// DeleteCategory deletes the category with the given UUID.
	DeleteCategory(uuid string) error
}
