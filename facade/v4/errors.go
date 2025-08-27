package v4

import (
	"fmt"
	"reflect"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
)

func ClassifyV4APICallError[R APIResponse, Rerr APIResponseData, Terr APIOneOfErrorResponse](response R, err error) error {
	if err == nil {
		return nil
	}

	if reflect.ValueOf(response).IsNil() {
		return ferrors.NewErrUncategorisedError("", err)
	}

	data := response.GetData()
	if data == nil {
		return ferrors.NewErrUncategorisedError("", err)
	}

	// 4xx, 5xx from v4
	result, ok := data.(Rerr)
	if !ok {
		return ferrors.NewErrTypeAssertionError("", fmt.Errorf("unexpected type for API response data: %T", data))
	}

	unstructured := result.GetValue()
	oneOfApiError, ok := unstructured.(Terr)
	if !ok {
		return ferrors.NewErrTypeAssertionError("", fmt.Errorf("unexpected type for one of API response error: %T", oneOfApiError))
	}

	// oneOfApiError.GetError is one of []AppMessage or SchemaValidationError
	return ferrors.NewErrApiError("API call failed", oneOfApiError.GetError())
}
