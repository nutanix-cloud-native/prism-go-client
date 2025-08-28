package v4

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"syscall"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4VmmError "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
)

// GetCategorisedV4ApiCallError return classified and wrapped errors.
func GetCategorisedV4ApiCallError[R ApiResponse, Rerr ApiErrorResponse](response R, err error) error {
	if err == nil {
		return nil
	}

	if ApiResponse(response) == nil {
		if isNetworkError(err) {
			return ferrors.NewErrNetworkError("API call failed", err)
		}

		return ferrors.NewErrUncategorisedError("API call failed", err)
	}

	data := response.GetData()
	if data == nil {
		return ferrors.NewErrUncategorisedError("", err)
	}

	// 4xx, 5xx from v4
	// ERROR: Rerr is an interface how can I get the struct?
	unstructuredErrorResponse, ok := data.(Rerr) // 400 -> ErrorResponse 401 -> TProjection
	if !ok {
		return ferrors.NewErrTypeAssertionError("", fmt.Errorf("unexpected type for API response data: %T", data))
	}

	return getCategorisedV4ApiErrorFromErrorResponse(unstructuredErrorResponse)
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

func getCategorisedV4ApiErrorFromErrorResponse[Rerr ApiErrorResponse](errorResponse Rerr) error {
	errResponseError := errorResponse.GetError() // either of []AppMessage or SchemaValidationError

	if result, ok := errResponseError.(v4VmmError.SchemaValidationError); ok {
		return ferrors.NewErrV4ApiSchemaValidationError("", result)
	} else if result, ok := errResponseError.([]v4VmmError.AppMessage); ok {
		return getCategorisedV4ApiErrorFromAppMessages(result)
	}
	return ferrors.NewErrV4ApiUncategorisedError("API call failed", errResponseError)
}

func getCategorisedV4ApiErrorFromAppMessages(appMessages []v4VmmError.AppMessage) error {
	err := ferrors.NewErrUncategorisedError("", appMessages)

	fmt.Println("AHO")
	for _, appMessage := range appMessages {
		// string search each appMessage.errorGroup and add it to a categorised subtype
		// but do we need to rank them if multiple are present????
		fmt.Printf("%s ,", *appMessage.ErrorGroup)
	}

	return err
}
