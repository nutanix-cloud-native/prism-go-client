package v4

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	responseModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
)

type V4ODataParams struct {
	Page    *int
	Limit   *int
	Filter  *string
	OrderBy *string
	Expand  *string
	Select  *string
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

func GetMetadata[R APIResponse](response R) (*responseModels.ApiResponseMetadata, error) {
	hasMetadataField := reflect.ValueOf(response).Elem().FieldByName("Metadata")
	if !hasMetadataField.IsValid() {
		return nil, fmt.Errorf("response does not have Metadata field")
	}
	metadata := hasMetadataField.Interface().(responseModels.ApiResponseMetadata)
	return &metadata, nil
}

func CallListAPI[R APIResponse, T any](response R, err error) ([]T, int, error) {
	var zero []T
	if err != nil {
		return zero, 0, fmt.Errorf("API call failed: %w", err)
	}

	totalCount := 0
	metadata, err := GetMetadata(response)
	if err != nil {
		return zero, 0, fmt.Errorf("failed to get metadata: %w", err)
	}
	if metadata != nil && metadata.TotalAvailableResults != nil {
		totalCount = *metadata.TotalAvailableResults
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
		if err := opt(params); err != nil {
			return nil, fmt.Errorf("failed to apply OData option: %w", err)
		}
	}
	return params, nil
}
