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
	unstructuredErrorResponseError, err := getErrorField(errorResponse)
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

func getErrorField(unstructuredErrorResponse interface{}) (interface{}, error) {
	// TODO av
	v := reflect.ValueOf(unstructuredErrorResponse)
	fbn := v.FieldByName("Error").Interface()

	return fbn, nil
}

func getObjectType(unstructuredErrorResponseError interface{}) (string, error) {
	// TODO av
	v := reflect.ValueOf(unstructuredErrorResponseError).Elem()
	fbn := v.FieldByName("ObjectType_")
	objectType := *fbn.Interface().(*string)
	return objectType, nil
}

func getCategorisedV4ApiErrorFromAppMessages(errorResponse interface{}, appMessages interface{}) error {
	v := reflect.ValueOf(appMessages)
	if v.Kind() != reflect.Slice {
		return ferrors.NewErrTypeAssertionError("Invalid type for []AppMessage", errorResponse)
	}

	err := ferrors.NewErrV4ApiUncategorisedError("", errorResponse)

	len := v.Len()
	for i := range len {
		elem := v.Index(i)

		// TODO av
		errorGroup := *reflect.ValueOf(elem.Interface()).FieldByName("ErrorGroup").Interface().(*string)

		fmt.Println(errorGroup)
	}

	return err
}
