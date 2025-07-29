package v4

import (
	"fmt"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
)

// Common generic function for CRUD operations

func CommonGetEntity[R APIResponse, T any](apiCall func() (R, error), entityName string) (*T, error) {
	result, err := CallAPI[R, T](apiCall())
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", entityName, err)
	}
	return &result, nil
}

func CommonListEntities[R APIResponse, T any](apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) ([]T, error) {
	reqParams, err := OptsToV4ODataParams(options...)

	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	result, err := CallAPI[R, []T](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to list %s: %w", entitiesName, err)
	}
	return result, nil
}

func CommonListAllEntities[R APIResponse, T any](apiCall func(reqParams *V4ODataParams) (R, error), reqParams *V4ODataParams, entitiesName string) ([]T, error) {
	result := []T{}
	page := 0

	if reqParams == nil {
		reqParams = &V4ODataParams{}
	}

	reqParams.Page = &page
	reqParams.Limit = nil // Let API use the default limit

	items, totalCount, err := CallListAPI[R, T](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to list all %s: %w", entitiesName, err)
	}
	result = append(result, items...)

	for len(result) < totalCount {
		page++
		reqParams.Page = &page
		moreItems, _, err := CallListAPI[R, T](apiCall(reqParams))
		if err != nil {
			return nil, fmt.Errorf("failed to list all %s on page %d: %w", entitiesName, page, err)
		}
		result = append(result, moreItems...)
	}

	return result, nil
}

func CommonGetListIterator[R APIResponse, T any](f *FacadeV4Client, apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) (facade.ODataListIterator[T], error) {
	reqParams, err := OptsToV4ODataParams(options...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4 OData params: %w", err)
	}

	page := 0
	reqParams.Limit = nil  // Let API use the default limit
	reqParams.Page = &page // Start from the page 0

	itemsBuffer, totalCount, err := CallListAPI[R, T](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to get list iterator for %s: %w", entitiesName, err)
	}

	return NewFacadeV4ODataIterator(
		f.client,
		totalCount,
		itemsBuffer,
		func(reqParams *V4ODataParams) (R, error) {
			return apiCall(reqParams)
		},
		options...,
	), nil
}
