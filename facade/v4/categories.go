package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	v4prismError "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

// FacadeV4Client implements the CategoriesFacadeV4 interface.

// GetCategory returns the category for the given UUID.
func (f *FacadeV4Client) GetCategory(uuid string) (*v4prismModels.Category, error) {
	return CommonGetEntity[*v4prismModels.GetCategoryApiResponse, v4prismModels.Category, *v4prismModels.OneOfGetCategoryApiResponseData, *v4prismError.ErrorResponse](
		func() (*v4prismModels.GetCategoryApiResponse, error) {
			return f.client.CategoriesApiInstance.GetCategoryById(&uuid, nil)
		},
		"category",
	)
}

// ListCategories returns a list of categories.
func (f *FacadeV4Client) ListCategories(opts ...facade.ODataOption) ([]v4prismModels.Category, error) {
	return CommonListEntities[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category, *v4prismModels.OneOfListCategoriesApiResponseData, *v4prismError.ErrorResponse](
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
func (f *FacadeV4Client) ListAllCategories(filterParam *string, orderbyParam *string, expandParam *string, selectParam *string) ([]v4prismModels.Category, error) {
	reqParams := &V4ODataParams{
		Filter:  filterParam,
		OrderBy: orderbyParam,
		Expand:  expandParam,
		Select:  selectParam,
	}

	return CommonListAllEntities[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category, *v4prismModels.OneOfListCategoriesApiResponseData, *v4prismError.ErrorResponse](
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
func (f *FacadeV4Client) GetListIteratorCategories(opts ...facade.ODataOption) facade.ODataListIterator[v4prismModels.Category] {
	return CommonGetListIterator[*v4prismModels.ListCategoriesApiResponse, v4prismModels.Category, *v4prismModels.OneOfListCategoriesApiResponseData, *v4prismError.ErrorResponse](
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
func (f *FacadeV4Client) CreateCategory(category *v4prismModels.Category) (*v4prismModels.Category, error) {
	newCategory, err := CallAPI[*v4prismModels.CreateCategoryApiResponse, v4prismModels.Category, *v4prismModels.OneOfCreateCategoryApiResponseData, *v4prismError.ErrorResponse](
		f.client.CategoriesApiInstance.CreateCategory(category),
	)
	if err != nil {
		return nil, err
	}
	return &newCategory, nil
}

// UpdateCategory updates an existing category.
func (f *FacadeV4Client) UpdateCategory(uuid string, category *v4prismModels.Category) (*v4prismModels.Category, error) {
	existingCategory, args, err := GetEntityAndEtag(
		f.GetCategory(uuid),
	)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, ferrors.NewErrUncategorisedError("", fmt.Errorf("no category found with UUID %s", uuid))
	}

	_, err = CallAPI[*v4prismModels.UpdateCategoryApiResponse, []v4prismError.AppMessage, *v4prismModels.OneOfUpdateCategoryApiResponseData, *v4prismError.ErrorResponse](
		f.client.CategoriesApiInstance.UpdateCategoryById(&uuid, category, args),
	)
	if err != nil {
		return nil, err
	}

	updatedCategory, err := f.GetCategory(uuid)
	if err != nil {
		return nil, err
	}
	if updatedCategory == nil {
		return nil, ferrors.NewErrUncategorisedError("", fmt.Errorf("no updated category found with UUID %s", uuid))
	}
	if updatedCategory.ExtId == nil {
		return nil, ferrors.NewErrUncategorisedError("", fmt.Errorf("updated category ExtId is nil for UUID %s", uuid))
	}
	return updatedCategory, nil
}

// DeleteCategory deletes the category with the given UUID.
func (f *FacadeV4Client) DeleteCategory(uuid string) error {
	category, args, err := GetEntityAndEtag(
		f.GetCategory(uuid),
	)
	if err != nil {
		return err
	}
	if category == nil {
		return ferrors.NewErrUncategorisedError("", fmt.Errorf("no category found with UUID %s", uuid))
	}

	_, err = f.client.CategoriesApiInstance.DeleteCategoryById(&uuid, args)
	if err != nil {
		return err
	}

	return nil
}
