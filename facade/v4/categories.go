package v4

import (
	"context"
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	prismMessages "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

// FacadeV4Client implements the CategoriesFacadeV4 interface.

// GetCategory returns the category for the given UUID.
func (f *FacadeV4Client) GetCategory(ctx context.Context, uuid string) (*v4prismModels.Category, error) {
	return CommonGetEntity[*v4prismModels.GetCategoryApiResponse, v4prismModels.Category](
		func() (*v4prismModels.GetCategoryApiResponse, error) {
			return f.client.CategoriesApiInstance.GetCategoryById(&uuid, nil)
		},
		"category",
	)
}

// ListCategories returns a list of categories.
func (f *FacadeV4Client) ListCategories(ctx context.Context, opts ...facade.ODataOption) ([]v4prismModels.Category, error) {
	return CommonListEntities[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category](
		func(reqParams *V4ODataParams) (*v4prismModels.ListCategoriesApiResponse, error) {
			return f.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"categories",
	)
}

// ListAllCategories returns all categories without pagination.
func (f *FacadeV4Client) ListAllCategories(ctx context.Context, filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]v4prismModels.Category, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Expand:  expandParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category](
		func(reqParams *V4ODataParams) (*v4prismModels.ListCategoriesApiResponse, error) {
			return f.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		reqParams,
		"categories",
	)
}

// GetListIteratorCategories returns an iterator for listing categories.
func (f *FacadeV4Client) GetListIteratorCategories(ctx context.Context, opts ...facade.ODataOption) facade.ODataListIterator[v4prismModels.Category] {
	return CommonGetListIterator[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category](
		ctx,
		f,
		func(reqParams *V4ODataParams) (*v4prismModels.ListCategoriesApiResponse, error) {
			return f.client.CategoriesApiInstance.ListCategories(
				reqParams.Page,
				reqParams.Limit,
				reqParams.Filter,
				reqParams.OrderBy,
				reqParams.Expand,
				reqParams.Select,
			)
		},
		opts,
		"categories",
	)
}

// CreateCategory creates a new category.
func (f *FacadeV4Client) CreateCategory(ctx context.Context, category *v4prismModels.Category) (*v4prismModels.Category, error) {
	newCategory, err := CallAPI[*v4prismModels.CreateCategoryApiResponse, v4prismModels.Category](
		f.client.CategoriesApiInstance.CreateCategory(category),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}
	return &newCategory, nil
}

// UpdateCategory updates an existing category.
func (f *FacadeV4Client) UpdateCategory(ctx context.Context, uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error) {
	existingCategory, args, err := GetEntityAndEtag(
		f.GetCategory(ctx, uuid),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get category with UUID %s: %w", uuid, err)
	}
	if existingCategory == nil {
		return nil, fmt.Errorf("no category found with UUID %s", uuid)
	}

	_, err = CallAPI[*v4prismModels.UpdateCategoryApiResponse, []prismMessages.AppMessage](
		f.client.CategoriesApiInstance.UpdateCategoryById(&uuid, category, args),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update category with UUID %s: %w", uuid, err)
	}

	updatedCategory, err := f.GetCategory(ctx, uuid)
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

// DeleteCategory deletes the category with the given UUID.
func (f *FacadeV4Client) DeleteCategory(ctx context.Context, uuid string) error {
	category, args, err := GetEntityAndEtag(
		f.GetCategory(ctx, uuid),
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
