package v4

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
)

type V4ODataParams struct {
	Page    *int
	Limit   *int
	Filter  *string
	OrderBy *string
	Expand  *string
	Select  *string
	Apply   *string
}

func ToV4ODataParams(params facade.ODataOptions) (*V4ODataParams, error) {
	if params == nil {
		return nil, nil
	}

	if v4Params, ok := params.(*V4ODataParams); ok {
		return v4Params, nil
	}

	return nil, fmt.Errorf("expected *V4ODataParams, got %T", params)
}

func (o *V4ODataParams) SetPageOption(page int) error {
	o.Page = &page
	return nil
}

func (o *V4ODataParams) SetLimitOption(limit int) error {
	o.Limit = &limit
	return nil
}

func (o *V4ODataParams) SetFilterOption(filter string) error {
	o.Filter = &filter
	return nil
}

func (o *V4ODataParams) SetOrderByOption(orderBy string) error {
	o.OrderBy = &orderBy
	return nil
}

func (o *V4ODataParams) SetExpandOption(expand string) error {
	o.Expand = &expand
	return nil
}

func (o *V4ODataParams) SetSelectOption(selectFields string) error {
	o.Select = &selectFields
	return nil
}

func (o *V4ODataParams) SetApplyOption(apply string) error {
	o.Apply = &apply
	return nil
}

func GetEtag(object interface{}) string {
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
		reservedMap := reserved.Interface().(map[string]interface{})
		for k, v := range reservedMap {
			if strings.ToLower(k) == etagKey {
				return v.(string)
			}
		}
	}

	return ""
}

func DropEtag(object interface{}) interface{} {
	if reflect.TypeOf(object).Kind() == reflect.Struct {
		reserved := reflect.ValueOf(object).FieldByName("Reserved_")
		if reserved.IsValid() {
			reservedMap := reserved.Interface().(map[string]interface{})
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

func CopyEtag(source, destination interface{}) interface{} {
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
		reservedMap := reserved.Interface().(map[string]interface{})
		reservedMap["Etag"] = etag
	}

	return destination
}

var (
	V4TaskStatusesMap = map[v4prismModels.TaskStatus]facade.TaskStatus{
		v4prismModels.TASKSTATUS_CANCELED:  facade.TaskStatusCanceled,
		v4prismModels.TASKSTATUS_CANCELING: facade.TaskStatusCanceling,
		v4prismModels.TASKSTATUS_FAILED:    facade.TaskStatusFailed,
		v4prismModels.TASKSTATUS_QUEUED:    facade.TaskStatusQueued,
		v4prismModels.TASKSTATUS_RUNNING:   facade.TaskStatusRunning,
		v4prismModels.TASKSTATUS_SUSPENDED: facade.TaskStatusSuspended,
		v4prismModels.TASKSTATUS_SUCCEEDED: facade.TaskStatusSucceeded,
		v4prismModels.TASKSTATUS_UNKNOWN:   facade.TaskStatusUnknown,
		v4prismModels.TASKSTATUS_REDACTED:  facade.TaskStatusRedacted,
	}
)

func ConvertTaskStatus(status v4prismModels.TaskStatus) facade.TaskStatus {
	if convertedStatus, ok := V4TaskStatusesMap[status]; ok {
		return convertedStatus
	}
	return facade.TaskStatusUnknown
}

type APIResponse interface {
	GetData() interface{}
}

type APIResponseData interface {
	GetValue() interface{}
}

type APIOneOfErrorResponseError interface {
	GetError() interface{} // either AppMessage or SchemaValidationError
}

func CallAPI[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](response R, err error) (T, error) {
	var zero, result T
	if err != nil {
		return zero, ClassifyV4APICallError[R, Rerr, Terr](response, err)
	}

	data := response.GetData()
	if data == nil {
		return zero, nil
	}

	result, ok := data.(T)
	if !ok {
		errStr := fmt.Sprintf("unexpected type for API response data: %T", data)
		return zero, ferrors.NewErrTypeAssertionError(errStr, errors.New(errStr))
	}

	return result, nil
}

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

func CallListAPI[R APIResponse, T any, Rerr APIResponseData, Terr APIOneOfErrorResponseError](response R, err error) ([]T, int, error) {
	var zero []T
	if err != nil {
		return zero, 0, ClassifyV4APICallError[R, Rerr, Terr](response, err)
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

func GetEntityAndEtag[T any](entity T, err error) (T, map[string]interface{}, error) {
	var zero T

	if err != nil {
		return zero, nil, fmt.Errorf("failed to get entity: %w", err)
	}

	etag := GetEtag(entity)
	if etag == "" {
		return zero, nil, fmt.Errorf("no ETag found for entity of type %T", entity)
	}

	args := map[string]interface{}{
		"If-Match": &etag,
	}

	return entity, args, nil
}

func OptsToV4ODataParams(opts ...facade.ODataOption) (*V4ODataParams, error) {
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
