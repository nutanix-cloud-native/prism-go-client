package v4

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
)

// Common generic function for CRUD operations

func CommonGetEntity[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](apiCall func() (R, error), entityName string) (*T, error) {
	result, err := CallAPI[R, T, Rerr, Terr](apiCall())
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", entityName, err)
	}
	return &result, nil
}

func CommonListEntities[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) ([]T, error) {
	reqParams, err := OptsToV4ODataParams(options...)

	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	result, err := CallAPI[R, []T, Rerr, Terr](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to list %s: %w", entitiesName, err)
	}
	return result, nil
}

func CommonListAllEntities[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](apiCall func(reqParams *V4ODataParams) (R, error), reqParams *V4ODataParams, entitiesName string) ([]T, error) {
	result := []T{}
	page := 0

	if reqParams == nil {
		reqParams = &V4ODataParams{}
	}

	reqParams.Page = &page
	reqParams.Limit = nil // Let API use the default limit

	items, totalCount, err := CallListAPI[R, T, Rerr, Terr](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to list all %s: %w", entitiesName, err)
	}
	result = append(result, items...)

	for len(result) < totalCount {
		page++
		reqParams.Page = &page
		moreItems, _, err := CallListAPI[R, T, Rerr, Terr](apiCall(reqParams))
		if err != nil {
			return nil, fmt.Errorf("failed to list all %s on page %d: %w", entitiesName, page, err)
		}
		result = append(result, moreItems...)
	}

	return result, nil
}

func CommonGetListIterator[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](f *FacadeV4Client, apiCall func(reqParams *V4ODataParams) (R, error), options []facade.ODataOption, entitiesName string) facade.ODataListIterator[T] {
	return NewFacadeV4ODataIterator[R, T, Rerr, Terr](
		f.client,
		func(reqParams *V4ODataParams) (R, error) {
			return apiCall(reqParams)
		},
		options...,
	)
}

func ClassifyV4APICallError[R APIResponse, Rerr APIResponseData, Terr APIOneOfErrorResponseError](response R, err error) error {
	if err == nil {
		return nil
	}

	if reflect.ValueOf(response).IsNil() {
		return ferrors.NewErrUnknownError("", err)
	}

	data := response.GetData()
	if data == nil {
		return ferrors.NewErrUnknownError("", err)
	}

	result, ok := data.(Rerr)
	if !ok {
		errStr := fmt.Sprintf("unexpected type for API response data: %T", data)
		return ferrors.NewErrTypeAssertionError(errStr, err)
	}

	apiError := result.GetValue()
	oneOfApiError, ok := apiError.(Terr)
	if !ok {
		errStr := fmt.Sprintf("unexpected type for API response data error: %T", oneOfApiError)
		return ferrors.NewErrTypeAssertionError("", errors.New(errStr))
	}

	return ferrors.NewErrApiError("API call failed", oneOfApiError.GetError())
}
