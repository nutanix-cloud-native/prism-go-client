package v4

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	v4prismModels "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
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

func (o *V4ODataParams) SetPageOption(page int) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}
		options.Page = &page
		return nil
	}
}

func (o *V4ODataParams) SetLimitOption(limit int) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}

		options.Limit = &limit
		return nil
	}
}

func (o *V4ODataParams) SetFilterOption(filter string) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}

		options.Filter = &filter
		return nil
	}
}

func (o *V4ODataParams) SetOrderByOption(orderBy string) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}

		options.OrderBy = &orderBy
		return nil
	}
}

func (o *V4ODataParams) SetExpandOption(expand string) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}

		options.Expand = &expand
		return nil
	}
}

func (o *V4ODataParams) SetSelectOption(selectFields string) facade.ODataOption {
	return func(params facade.ODataOptions) error {
		options, err := ToV4ODataParams(params)
		if err != nil {
			return err
		}

		options.Select = &selectFields
		return nil
	}
}

func GetEtag(object interface{}) *string {
	var reserved reflect.Value
	if reflect.TypeOf(object).Kind() == reflect.Struct {
		reserved = reflect.ValueOf(object).FieldByName("Reserved_")
	} else if reflect.TypeOf(object).Kind() == reflect.Interface || reflect.TypeOf(object).Kind() == reflect.Ptr {
		reserved = reflect.ValueOf(object).Elem().FieldByName("Reserved_")
	} else {
		return nil
	}

	if reserved.IsValid() {
		etagKey := strings.ToLower("Etag")
		reservedMap := reserved.Interface().(map[string]interface{})
		for k, v := range reservedMap {
			if strings.ToLower(k) == etagKey {
				return v.(*string)
			}
		}
	}

	return nil
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
		return zero, fmt.Errorf("no data found in API response")
	}

	result, ok := data.(T)
	if !ok {
		return zero, fmt.Errorf("unexpected type for API response data: %T", data)
	}

	return result, nil
}

func GetEntityAndEtag[T any](entity T, err error) (T, map[string]interface{}, error) {
	var zero T

	if err != nil {
		return zero, nil, fmt.Errorf("failed to get entity: %w", err)
	}

	etag := GetEtag(entity)
	if etag == nil {
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
