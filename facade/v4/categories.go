package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

// FacadeV4Client implements the CategoriesFacadeV4 interface.

// GetCategory returns the category for the given UUID.
func (f *FacadeV4Client) GetCategory(uuid string) (*v4prismModels.Category, error) {
	category, err := CallAPI[*v4prismModels.GetCategoryApiResponse, v4prismModels.Category](
		f.client.CategoriesApiInstance.GetCategoryById(&uuid, nil),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get category with UUID %s: %w", uuid, err)
	}

	return &category, nil
}

// ListCategories returns a list of categories.
func (f *FacadeV4Client) ListCategories(opts ...facade.ODataOption) ([]v4prismModels.Category, error) {
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4 OData params: %w", err)
	}

	categories, err := CallAPI[*v4prismModels.ListCategoriesApiResponse, []v4prismModels.Category](
		f.client.CategoriesApiInstance.ListCategories(reqParams.Page, reqParams.Limit, reqParams.Filter, reqParams.OrderBy, reqParams.Expand, reqParams.Select),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	return categories, nil
}

// ListAllCategories returns all categories without pagination.
func (f *FacadeV4Client) ListAllCategories(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]v4prismModels.Category, error) {
	result := []v4prismModels.Category{}

	page := 0
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Expand:  expandParam,
		Select:  selectParam,
		Page:    &page,
		Limit:   nil, // Let API use the default limit
	}
	categories, totalCount, err := CallListAPI[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category](
		f.client.CategoriesApiInstance.ListCategories(
			reqParams.Page,
			reqParams.Limit,
			reqParams.Filter,
			reqParams.OrderBy,
			reqParams.Expand,
			reqParams.Select,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list all categories: %w", err)
	}
	result = append(result, categories...)
	for len(result) < totalCount {
		page++
		categories, totalCount, err = CallListAPI[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category](
			f.client.CategoriesApiInstance.ListCategories(
				&page,
				nil,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to list all categories: %w", err)
		}
		result = append(result, categories...)
	}
	return result, nil
}

// GetListIteratorCategories returns an iterator for listing categories.
func (f *FacadeV4Client) GetListIteratorCategories(opts ...facade.ODataOption) (facade.ODataListIterator[*v4prismModels.Category], error) {
	reqParams, err := OptsToV4ODataParams(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4 OData params: %w", err)
	}
	page := 0
	reqParams.Limit = nil  // Let API use the default limit
	reqParams.Page = &page // Start from the page 0
	itemsBuffer, totalCount, err := CallListAPI[*v4prismModels.ListCategoriesApiResponse, *v4prismModels.Category](
		f.client.CategoriesApiInstance.ListCategories(&page, nil, reqParams.Filter, reqParams.OrderBy, reqParams.Expand, reqParams.Select),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories iterator: %w", err)
	}

	return NewFacadeV4ODataIterator(
		f.client,
		totalCount,
		itemsBuffer,
		func(params *V4ODataParams) (*v4prismModels.ListCategoriesApiResponse, error) {
			return f.client.CategoriesApiInstance.ListCategories(
				params.Page,
				params.Limit,
				params.Filter,
				params.OrderBy,
				params.Expand,
				params.Select,
			)
		},
		opts...,
	), nil
}

// CreateCategory creates a new category.
func (f *FacadeV4Client) CreateCategory(category *v4prismModels.Category) (*v4prismModels.Category, error) {
	newCategory, err := CallAPI[*v4prismModels.CreateCategoryApiResponse, v4prismModels.Category](
		f.client.CategoriesApiInstance.CreateCategory(category),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}
	return &newCategory, nil
}

// UpdateCategory updates an existing category.
func (f *FacadeV4Client) UpdateCategory(uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error) {
	updatedCategory, err := CallAPI[*v4prismModels.UpdateCategoryApiResponse, v4prismModels.Category](
		f.client.CategoriesApiInstance.UpdateCategoryById(&uuid, category),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update category with UUID %s: %w", uuid, err)
	}
	return &updatedCategory, nil
}

// DeleteCategory deletes the category with the given UUID.
func (f *FacadeV4Client) DeleteCategory(uuid string) error {
	category, args, err := GetEntityAndEtag(
		f.GetCategory(uuid),
	)
	if err != nil {
		return fmt.Errorf("failed to get category with UUID %s: %w", uuid, err)
	}
	if category == nil {
		return fmt.Errorf("no category found with UUID %s", uuid)
	}

	_, err = f.client.CategoriesApiInstance.DeleteCategoryById(&uuid, args)
	if err != nil {
		return fmt.Errorf("failed to delete category with UUID %s: %w", uuid, err)
	}

	return nil
}
