package v4

import (
	"errors"
	"fmt"
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

	if ApiResponse(response) == nil {
		if isNetworkError(err) {
			return ferrors.NewErrNetworkError("API call failed", err)
		}

		return ferrors.NewErrUncategorisedError("API call failed", err)
	}

	errorResponse := response.GetData()

	return getCategorisedV4ApiError[Rerr](errorResponse)
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

func getCategorisedV4ApiError[Rerr ApiErrorResponseError](errorResponse interface{}) error {
	unstructuredErrorResponseError, err := getErrorFieldValue(errorResponse)
	if err != nil {
		return err
	}

	errorResponseError, ok := unstructuredErrorResponseError.(Rerr)
	if !ok {
		return ferrors.NewErrTypeAssertionError("Invalid API response format", errorResponse)
	}
	errorVal := errorResponseError.GetValue()

	objectType, err := getObjectType(unstructuredErrorResponseError)
	if err != nil {
		return err
	}

	if strings.Contains(objectType, "SchemaValidationError") {
		return ferrors.NewErrV4ApiSchemaValidationError("", errorResponse)
	}

	if strings.Contains(objectType, "AppMessage") {
		return getCategorisedV4ApiErrorFromAppMessages(errorResponse, errorVal)
	}

	return ferrors.NewErrV4ApiUncategorisedError("Invalid OneOfErrorResponseError", errorResponse)
}

func getErrorFieldValue(unstructuredErrorResponse interface{}) (interface{}, error) {
	v := reflect.ValueOf(unstructuredErrorResponse)
	if v.Kind() != reflect.Struct {
		return nil, ferrors.NewErrTypeAssertionError("API Error Response is not a struct", unstructuredErrorResponse)
	}
	fbn := v.FieldByName("Error")
	if !fbn.IsValid() {
		return nil, ferrors.NewErrInternalError("Error field missing in API Error Response", unstructuredErrorResponse)
	}

	return fbn.Interface(), nil
}

func getObjectType(unstructuredErrorResponseError interface{}) (string, error) {
	v := reflect.ValueOf(unstructuredErrorResponseError)
	if v.Kind() != reflect.Ptr {
		return "", ferrors.NewErrTypeAssertionError("Invalid API Error Response Error, should be a pointer", unstructuredErrorResponseError)
	}

	vElem := v.Elem()
	if vElem.Kind() != reflect.Struct {
		return "", ferrors.NewErrTypeAssertionError("API Error Response is not a struct", unstructuredErrorResponseError)
	}

	fbn := vElem.FieldByName("ObjectType_")
	if !fbn.IsValid() {
		return "", ferrors.NewErrInternalError("`ObjectType_` field missing in API Error Response Error", unstructuredErrorResponseError)
	}

	strPtr, ok := fbn.Interface().(*string)
	if !ok || strPtr == nil {
		return "", ferrors.NewErrInternalError("Invalid value for field `ObjectType_`", unstructuredErrorResponseError)
	}

	return *strPtr, nil
}

func getCategorisedV4ApiErrorFromAppMessages(errorResponse interface{}, appMessages interface{}) error {
	v := reflect.ValueOf(appMessages)
	if v.Kind() != reflect.Slice {
		return ferrors.NewErrTypeAssertionError("Invalid type for []AppMessage in API Error Response", errorResponse)
	}

	err := ferrors.NewErrV4ApiUncategorisedError("", errorResponse)

	len := v.Len()
	for i := range len {
		elem := v.Index(i)

		if elem.Kind() != reflect.Struct {
			return ferrors.NewErrTypeAssertionError("Invalid AppMessage type", errorResponse, map[string]interface{}{"index": i})
		}

		fbn := elem.FieldByName("ErrorGroup")
		if !fbn.IsValid() {
			return ferrors.NewErrInternalError("`ErrorGroup` field missing in API Error Response", errorResponse)
		}

		strPtr, ok := fbn.Interface().(*string)
		if !ok || strPtr == nil {
			return ferrors.NewErrInternalError("Invalid value for field `Error`", errorResponse)
		}

		errorGroup := *strPtr
		fmt.Println(errorGroup)
	}

	return err
}
