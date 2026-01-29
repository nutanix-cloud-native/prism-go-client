package v4

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/converged"

	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"k8s.io/utils/ptr"
)

// V4ODataParams holds OData query parameters for Prism Central V4 list APIs (page, limit, filter, orderBy, expand, select, apply).
type V4ODataParams struct {
	Page    *int
	Limit   *int
	Filter  *string
	OrderBy *string
	Expand  *string
	Select  *string
	Apply   *string
}

// ToV4ODataParams converts converged ODataOptions to *V4ODataParams if the underlying type is V4ODataParams.
func ToV4ODataParams(params converged.ODataOptions) (*V4ODataParams, error) {
	if params == nil {
		return nil, nil
	}

	if v4Params, ok := params.(*V4ODataParams); ok {
		return v4Params, nil
	}

	return nil, fmt.Errorf("expected *V4ODataParams, got %T", params)
}

// SetPageOption sets the page number for list requests.
func (o *V4ODataParams) SetPageOption(page int) error {
	o.Page = &page
	return nil
}

// SetLimitOption sets the page size for list requests.
func (o *V4ODataParams) SetLimitOption(limit int) error {
	o.Limit = &limit
	return nil
}

// SetFilterOption sets the OData filter expression.
func (o *V4ODataParams) SetFilterOption(filter string) error {
	o.Filter = &filter
	return nil
}

// SetOrderByOption sets the OData order-by expression.
func (o *V4ODataParams) SetOrderByOption(orderBy string) error {
	o.OrderBy = &orderBy
	return nil
}

// SetExpandOption sets the OData expand expression.
func (o *V4ODataParams) SetExpandOption(expand string) error {
	o.Expand = &expand
	return nil
}

// SetSelectOption sets the OData select expression.
func (o *V4ODataParams) SetSelectOption(selectFields string) error {
	o.Select = &selectFields
	return nil
}

// SetApplyOption sets the OData apply expression.
func (o *V4ODataParams) SetApplyOption(apply string) error {
	o.Apply = &apply
	return nil
}

// GetEtag extracts the ETag from an API entity (via Reserved_ map) for conditional updates.
func GetEtag(object any) string {
	var reserved reflect.Value
	if reflect.TypeOf(object).Kind() == reflect.Struct {
		reserved = reflect.ValueOf(object).FieldByName("Reserved_")
	} else if reflect.TypeOf(object).Kind() == reflect.Interface || reflect.TypeOf(object).Kind() == reflect.Ptr {
		reserved = reflect.ValueOf(object).Elem().FieldByName("Reserved_")
	} else {
		return ""
	}

	if reserved.IsValid() {
		etagKey := strings.ToLower("Etag")
		reservedMap := reserved.Interface().(map[string]any)
		for k, v := range reservedMap {
			if strings.ToLower(k) == etagKey {
				return v.(string)
			}
		}
	}

	return ""
}

// DropEtag removes the ETag from the object's Reserved_ map.
func DropEtag(object any) any {
	if reflect.TypeOf(object).Kind() == reflect.Struct {
		reserved := reflect.ValueOf(object).FieldByName("Reserved_")
		if reserved.IsValid() {
			reservedMap := reserved.Interface().(map[string]any)
			allEtagKeys := make([]string, 0)
			for k := range reservedMap {
				if strings.ToLower(k) == "etag" {
					allEtagKeys = append(allEtagKeys, k)
				}
			}
			for _, k := range allEtagKeys {
				delete(reservedMap, k)
			}
		}
	}
	return object
}

// CopyEtag copies the ETag from source to destination (used when building update payloads).
func CopyEtag(source, destination any) any {
	var reserved reflect.Value

	destination = DropEtag(destination)
	if reflect.TypeOf(source).Kind() == reflect.Struct {
		reserved = reflect.ValueOf(source).FieldByName("Reserved_")
	} else if reflect.TypeOf(source).Kind() == reflect.Interface || reflect.TypeOf(source).Kind() == reflect.Ptr {
		reserved = reflect.ValueOf(source).Elem().FieldByName("Reserved_")
	} else {
		return destination
	}

	if reserved.IsValid() {
		etag := GetEtag(source)
		if etag == "" {
			return destination
		}
		reservedMap := reserved.Interface().(map[string]any)
		reservedMap["Etag"] = etag
	}

	return destination
}

// V4TaskStatusesMap maps Prism Central V4 task status values to converged.TaskStatus.
var V4TaskStatusesMap = map[v4prismModels.TaskStatus]converged.TaskStatus{
		v4prismModels.TASKSTATUS_CANCELED:  converged.TaskStatusCanceled,
		v4prismModels.TASKSTATUS_CANCELING: converged.TaskStatusCanceling,
		v4prismModels.TASKSTATUS_FAILED:    converged.TaskStatusFailed,
		v4prismModels.TASKSTATUS_QUEUED:    converged.TaskStatusQueued,
		v4prismModels.TASKSTATUS_RUNNING:   converged.TaskStatusRunning,
		v4prismModels.TASKSTATUS_SUSPENDED: converged.TaskStatusSuspended,
		v4prismModels.TASKSTATUS_SUCCEEDED: converged.TaskStatusSucceeded,
		v4prismModels.TASKSTATUS_UNKNOWN:   converged.TaskStatusUnknown,
		v4prismModels.TASKSTATUS_REDACTED:  converged.TaskStatusRedacted,
	}

// ConvertTaskStatus converts a V4 API task status to converged.TaskStatus.
func ConvertTaskStatus(status v4prismModels.TaskStatus) converged.TaskStatus {
	if convertedStatus, ok := V4TaskStatusesMap[status]; ok {
		return convertedStatus
	}
	return converged.TaskStatusUnknown
}

// APIResponse is implemented by V4 API response types that expose GetData().
type APIResponse interface {
	GetData() any
}

// CallAPI unwraps a V4 API response into the data type T, or returns an error.
func CallAPI[R APIResponse, T any](response R, err error) (T, error) {
	var zero, result T
	if err != nil {
		return zero, fmt.Errorf("API call failed: %w", err)
	}

	data := response.GetData()
	if data == nil {
		return zero, nil
	}

	result, ok := data.(T)
	if !ok {
		return zero, fmt.Errorf("unexpected type for API response data: %T", data)
	}

	return result, nil
}

// GetMetadataTotalResults returns Metadata.TotalAvailableResults from a V4 list API response.
func GetMetadataTotalResults[R APIResponse](response R) (int, error) {
	hasMetadataField := reflect.ValueOf(response).Elem().FieldByName("Metadata")
	if !hasMetadataField.IsValid() {
		return 0, fmt.Errorf("response does not have Metadata field")
	}
	metadata := hasMetadataField.Interface()
	if reflect.ValueOf(metadata).IsNil() {
		return 0, fmt.Errorf("no metadata found in response")
	}

	totalCountField := reflect.ValueOf(metadata).Elem().FieldByName("TotalAvailableResults")
	if !totalCountField.IsValid() || totalCountField.IsNil() {
		return 0, fmt.Errorf("metadata does not have TotalAvailableResults field")
	}
	totalCount := totalCountField.Interface().(*int)
	if totalCount == nil || *totalCount < 0 {
		return 0, fmt.Errorf("invalid total count: %v", totalCount)
	}
	return int(*totalCount), nil
}

// CallListAPI unwraps a V4 list API response into a slice of T and total count, or returns an error.
func CallListAPI[R APIResponse, T any](response R, err error) ([]T, int, error) {
	var zero []T
	if err != nil {
		return zero, 0, fmt.Errorf("API call failed: %w", err)
	}

	totalCount, err := GetMetadataTotalResults(response)
	if err != nil {
		return zero, 0, fmt.Errorf("failed to get total results from response metadata: %w", err)
	}

	data := response.GetData()
	if data == nil {
		return zero, totalCount, nil
	}

	result, ok := data.([]T)
	if !ok {
		return zero, 0, fmt.Errorf("unexpected type for API response data: %T", data)
	}

	return result, totalCount, nil
}

// GetEntityAndEtag returns the entity, its ETag, and If-Match args for conditional update/delete.
func GetEntityAndEtag[T any](entity T, err error) (T, map[string]any, error) {
	var zero T

	if err != nil {
		return zero, nil, fmt.Errorf("failed to get entity: %w", err)
	}

	etag := GetEtag(entity)
	if etag == "" {
		return zero, nil, fmt.Errorf("no ETag found for entity of type %T", entity)
	}

	args := map[string]any{
		"If-Match": &etag,
	}

	return entity, args, nil
}

// OptsToV4ODataParams builds V4ODataParams from converged OData options.
func OptsToV4ODataParams(opts ...converged.ODataOption) (*V4ODataParams, error) {
	params := &V4ODataParams{}
	for _, opt := range opts {
		if opt != nil {
			if err := opt(params); err != nil {
				return nil, fmt.Errorf("failed to apply OData option: %w", err)
			}
		}
	}
	return params, nil
}

// GenericGetEntity calls the API and returns the entity data, or an error including entityName.
func GenericGetEntity[R APIResponse, T any](apiCall func() (R, error), entityName string) (*T, error) {
	result, err := CallAPI[R, T](apiCall())
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", entityName, err)
	}
	return &result, nil
}

// GenericListEntities calls the list API with OData options and returns all entities (paginating if needed).
func GenericListEntities[R APIResponse, T any](apiCall func(reqParams *V4ODataParams) (R, error), options []converged.ODataOption, entitiesName string) ([]T, error) {
	returnAll := false
	page := 0

	reqParams, err := OptsToV4ODataParams(options...)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to V4ODataParams: %w", err)
	}

	if reqParams.Limit != nil && reqParams.Page == nil {
		reqParams.Page = ptr.To(page)
	}

	if reqParams.Page == nil {
		returnAll = true
		reqParams.Page = ptr.To(page)
		reqParams.Limit = nil // Let API use the default limit
	}

	result := []T{}

	items, totalCount, err := CallListAPI[R, T](apiCall(reqParams))
	if err != nil {
		return nil, fmt.Errorf("failed to list all %s: %w", entitiesName, err)
	}
	result = append(result, items...)

	if returnAll {
		for len(result) < totalCount {
			page++
			reqParams.Page = ptr.To(page)
			moreItems, _, err := CallListAPI[R, T](apiCall(reqParams))
			if err != nil {
				return nil, fmt.Errorf("failed to list all %s on page %d: %w", entitiesName, *reqParams.Page, err)
			}
			result = append(result, moreItems...)
		}
	}

	return result, nil
}

// GenericNewIterator returns an iterator that calls the list API with the given options.
func GenericNewIterator[R APIResponse, T any](ctx context.Context, apiCall func(ctx context.Context, reqParams *V4ODataParams) (R, error), options []converged.ODataOption, entitiesName string) converged.Iterator[T] {
	return NewIterator[R, T](
		ctx,
		func(ctx context.Context, reqParams *V4ODataParams) (R, error) {
			return apiCall(ctx, reqParams)
		},
		options...,
	)
}
