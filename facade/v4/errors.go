package v4

import (
	"errors"
	"net"
	"net/url"
	"reflect"
	"strings"
	"syscall"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
)

// GetCategorisedV4ApiCallError return categorised and wrapped errors.
func GetCategorisedV4ApiCallError[R ApiResponse, Rerr ApiErrorResponseError](response R, err error) error {
	if err == nil {
		return nil
	}

	if isNil(response) {
		if isNetworkError(err) {
			return ferrors.NewErrNetworkError("API call failed", err)
		}

		return ferrors.NewErrUncategorisedError("API call failed", err)
	}

	return getCategorisedV4ApiError[Rerr](response)
}

func isNil(resp interface{}) bool {
	if resp == nil {
		return true
	}

	v := reflect.ValueOf(resp)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}

// isNetworkError returns true if the given error is network-related.
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return true
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		return true
	}

	// Check for syscall-level errors (connection refused, reset, etc.)
	switch {
	case errors.Is(err, syscall.ECONNREFUSED),
		errors.Is(err, syscall.ECONNRESET),
		errors.Is(err, syscall.ECONNABORTED),
		errors.Is(err, syscall.ENETUNREACH),
		errors.Is(err, syscall.EHOSTUNREACH),
		errors.Is(err, syscall.ETIMEDOUT):
		return true
	}

	return false
}

func getCategorisedV4ApiError[Rerr ApiErrorResponseError](apiResponse ApiResponse) error {
	errorData := apiResponse.GetData()

	unstructuredErrorFieldPtr, err := getErrorFieldValue(errorData) // *OneOfErrorResponseError
	if err != nil {
		return err
	}

	objectType, err := getObjectType(unstructuredErrorFieldPtr)
	if err != nil {
		return err
	}

	if strings.Contains(objectType, "SchemaValidationError") {
		return ferrors.NewErrV4ApiSchemaValidationError("", errorData)
	}
	if strings.Contains(objectType, "AppMessage") {
		errorFieldIntf, ok := unstructuredErrorFieldPtr.(Rerr)
		if !ok {
			return ferrors.NewErrTypeAssertionError("Invalid value of field `Data` in ApiResponse", errorData)
		}

		return getCategorisedV4ApiErrorFromAppMessages(errorFieldIntf.GetValue(), errorData)
	}

	return ferrors.NewErrV4ApiUncategorisedError("Invalid value of field `Data` in ApiResponse", errorData)
}

func getErrorFieldValue(errorData interface{}) (interface{}, error) {
	v := reflect.ValueOf(errorData)
	if v.Kind() != reflect.Struct {
		return nil, ferrors.NewErrTypeAssertionError("API Error Response is not a struct", errorData)
	}
	fbn := v.FieldByName("Error")
	if !fbn.IsValid() {
		return nil, ferrors.NewErrInternalError("Error field missing in API Error Response", errorData)
	}

	return fbn.Interface(), nil
}

func getObjectType(unstructuredErrorFieldPtr interface{}) (string, error) {
	v := reflect.ValueOf(unstructuredErrorFieldPtr)
	if v.Kind() != reflect.Ptr {
		return "", ferrors.NewErrTypeAssertionError("Invalid value of `Error` field in ApiResponse, should be a pointer", unstructuredErrorFieldPtr)
	}

	vElem := v.Elem()
	if vElem.Kind() != reflect.Struct {
		return "", ferrors.NewErrTypeAssertionError("Invalid value of `Error` field pointer, should be a struct", unstructuredErrorFieldPtr)
	}

	fbn := vElem.FieldByName("ObjectType_")
	if !fbn.IsValid() {
		return "", ferrors.NewErrInternalError("`ObjectType_` field missing in `Error` field value", unstructuredErrorFieldPtr)
	}

	strPtr, ok := fbn.Interface().(*string)
	if !ok || strPtr == nil {
		return "", ferrors.NewErrInternalError("Invalid value for field `ObjectType_`", unstructuredErrorFieldPtr)
	}

	return *strPtr, nil
}

func getCategorisedV4ApiErrorFromAppMessages(appMessages interface{}, errorData interface{}) error {
	v := reflect.ValueOf(appMessages)
	if v.Kind() != reflect.Slice {
		return ferrors.NewErrTypeAssertionError("Invalid type for `[]AppMessage`", errorData)
	}

	apiErrSubType := ferrors.ErrorSubTypeV4ApiUncategorisedError

	len := v.Len()
	// DOUBT should we expect only 1???
	for i := range len {
		errArg := map[string]interface{}{"index": i}
		elem := v.Index(i)
		if elem.Kind() != reflect.Struct {
			return ferrors.NewErrTypeAssertionError("Invalid `AppMessage` type", errorData, errArg)
		}

		fbn := elem.FieldByName("ErrorGroup")
		if !fbn.IsValid() {
			return ferrors.NewErrInternalError("`ErrorGroup` field missing in `AppMessage`", errorData, errArg)
		}
		errorGroupStrPtr, ok := fbn.Interface().(*string)
		if !ok || errorGroupStrPtr == nil {
			return ferrors.NewErrInternalError("Invalid value for field `ErrorGroup`", errorData, errArg)
		}

		subType := getV4ApiErrorSubTypeForErrorGroup(*errorGroupStrPtr)
		// DOUBT ranking???
		if subType != "" {
			apiErrSubType = subType
		}
	}

	return getErrorForV4ApiErrSubType(apiErrSubType, errorData)
}

func getV4ApiErrorSubTypeForErrorGroup(errorGroup string) ferrors.ErrorSubTypeV4Api {
	searchKeyToSubTypeMap := map[string]ferrors.ErrorSubTypeV4Api{
		"AUTHORIZATION": ferrors.ErrorSubTypeV4ApiAuthorizationError,
		"RATE_LIMIT":    ferrors.ErrorSubTypeV4ApiRateLimitError,
		"NOT_FOUND":     ferrors.ErrorSubTypeV4ApiResourceNotFoundError,
	}

	for key, subType := range searchKeyToSubTypeMap {
		if strings.Contains(errorGroup, key) {
			return subType
		}
	}

	return ""
}

func getErrorForV4ApiErrSubType(subType ferrors.ErrorSubTypeV4Api, errorData interface{}) error {
	subTypeToConstructorMap := map[ferrors.ErrorSubTypeV4Api]interface{}{
		ferrors.ErrorSubTypeV4ApiUncategorisedError:    ferrors.NewErrV4ApiUncategorisedError,
		ferrors.ErrorSubTypeV4ApiAuthorizationError:    ferrors.NewErrV4ApiAuthorizationError,
		ferrors.ErrorSubTypeV4ApiResourceNotFoundError: ferrors.NewErrV4ApiResourceNotFoundError,
		ferrors.ErrorSubTypeV4ApiRateLimitError:        ferrors.NewErrV4ApiRateLimitError,
	}

	if constructor, found := subTypeToConstructorMap[subType]; found {
		return constructor.(func(string, interface{}, ...map[string]interface{}) error)("", errorData)
	}

	return ferrors.NewErrInternalError("Invalid v4 API error sub-type", errorData)
}
