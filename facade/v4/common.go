package v4

import (
	"github.com/nutanix-cloud-native/prism-go-client/facade"
)

// Common generic function for CRUD operations

func CommonGetEntity[R ApiResponse, T any](apiCall func() (R, error), entityName string) (*T, error) {
	result, err := CallAPI[R, T](apiCall())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func CommonListEntities[R ApiResponse, T any](apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) ([]T, error) {
	reqParams, err := OptsToV4ODataParams(options...)

	if err != nil {
		return nil, err
	}

	result, err := CallAPI[R, []T](apiCall(reqParams))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CommonListAllEntities[R ApiResponse, T any](apiCall func(reqParams *V4ODataParams) (R, error), reqParams *V4ODataParams, entitiesName string) ([]T, error) {
	result := []T{}
	page := 0

	if reqParams == nil {
		reqParams = &V4ODataParams{}
	}

	reqParams.Page = &page
	reqParams.Limit = nil // Let API use the default limit

	items, totalCount, err := CallListAPI[R, T](apiCall(reqParams))
	if err != nil {
		return nil, err
	}
	result = append(result, items...)

	for len(result) < totalCount {
		page++
		reqParams.Page = &page
		moreItems, _, err := CallListAPI[R, T](apiCall(reqParams))
		if err != nil {
			return nil, err
		}
		result = append(result, moreItems...)
	}

	return result, nil
}

func CommonGetListIterator[R ApiResponse, T any](f *FacadeV4Client, apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) facade.ODataListIterator[T] {
	return NewFacadeV4ODataIterator[R, T](
		f.client,
		func(reqParams *V4ODataParams) (R, error) {
			return apiCall(reqParams)
		},
		options...,
	)
}
