package v4

import (
	"encoding/json"
	"errors"
	"net"
	"net/url"
	"strings"
	"syscall"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	clusterClient "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
)

// GetCategorisedV4ApiCallError returns categorised and wrapped errors.
func GetCategorisedV4ApiCallError(err error) error {
	if err == nil {
		return nil
	}

	if openApiError, ok := err.(clusterClient.GenericOpenAPIError); ok {
		return getCategorisedV4ApiResponseError(openApiError)
	}

	if isNetworkError(err) {
		return ferrors.NewErrNetworkError("API call failed", err)
	}

	return ferrors.NewErrUncategorisedError("API call failed", err)
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

type sampleErrorResponse struct {
	ObjectType_ interface{} `json:"$errorItemDiscriminator"`
	Data        interface{} `json:"data"`
}

// getCategorisedV4ApiResponseError categorises the error based on the error data in the API response.
func getCategorisedV4ApiResponseError(openApiError clusterClient.GenericOpenAPIError) error {
	errResp := sampleErrorResponse{}

	err := json.Unmarshal(openApiError.Body, &errResp)
	if err != nil {
		return err
	}

	errorData, ok := errResp.Data.(map[string]interface{})
	if !ok {
		return ferrors.NewErrTypeAssertionError("Invalid API error response", openApiError.Error())
	}
	errorType, ok := errorData["$errorItemDiscriminator"].(string)
	if !ok {
		return ferrors.NewErrTypeAssertionError("Invalid API error response", openApiError.Error())
	}

	if strings.Contains(errorType, "SchemaValidationError") {
		return ferrors.NewErrV4ApiSchemaValidationError("API error response", openApiError.Error())
	}
	if strings.Contains(errorType, "AppMessage") {
		return getCategorisedV4ApiErrorFromAppMessages(openApiError, errorData["error"])
	}

	return ferrors.NewErrV4ApiUncategorisedError("Invalid value of field `Data` in ApiResponse", openApiError.Error())
}

func getCategorisedV4ApiErrorFromAppMessages(openApiError clusterClient.GenericOpenAPIError, appMessagesIntf interface{}) error {
	apiErrSubType := ferrors.ErrorSubTypeV4ApiUncategorisedError

	appMessages, ok := appMessagesIntf.([]interface{})
	if !ok {
		return ferrors.NewErrTypeAssertionError("Invalid []AppMessage type in API error response", openApiError.Error())
	}

	for _, appMessageIntf := range appMessages {
		appMessages, ok := appMessageIntf.(map[string]interface{})
		if !ok {
			return ferrors.NewErrTypeAssertionError("Invalid AppMessage type in API error response", openApiError.Error())
		}

		errorGroup, ok := appMessages["errorGroup"].(string)
		if !ok {
			return ferrors.NewErrTypeAssertionError("Invalid ErrorGroup type in API error response", openApiError.Error())
		}

		subType := getV4ApiErrorSubTypeForErrorGroup(errorGroup)
		// DOUBT ranking???
		if subType != "" {
			apiErrSubType = subType
		}
	}

	return getErrorForV4ApiErrSubType(apiErrSubType, openApiError)
}

func getV4ApiErrorSubTypeForErrorGroup(errorGroup string) ferrors.ErrorSubTypeV4Api {
	searchKeyToSubTypeMap := map[string]ferrors.ErrorSubTypeV4Api{
		"AUTHORIZATION": ferrors.ErrorSubTypeV4ApiAuthorizationError,
		"RATE_LIMIT":    ferrors.ErrorSubTypeV4ApiRateLimitError,
		"NOT_FOUND":     ferrors.ErrorSubTypeV4ApiResourceNotFoundError,
		"SERVICE_ERROR": ferrors.ErrorSubTypeV4ApiInternalServiceError,
		"INVALID_INPUT": ferrors.ErrorSubTypeV4ApiInvalidInputError,
	}

	for key, subType := range searchKeyToSubTypeMap {
		if strings.Contains(errorGroup, key) {
			return subType
		}
	}

	return ""
}

func getErrorForV4ApiErrSubType(subType ferrors.ErrorSubTypeV4Api, openApiError clusterClient.GenericOpenAPIError) error {
	subTypeToConstructorMap := map[ferrors.ErrorSubTypeV4Api]interface{}{
		ferrors.ErrorSubTypeV4ApiUncategorisedError:    ferrors.NewErrV4ApiUncategorisedError,
		ferrors.ErrorSubTypeV4ApiAuthorizationError:    ferrors.NewErrV4ApiAuthorizationError,
		ferrors.ErrorSubTypeV4ApiResourceNotFoundError: ferrors.NewErrV4ApiResourceNotFoundError,
		ferrors.ErrorSubTypeV4ApiRateLimitError:        ferrors.NewErrV4ApiRateLimitError,
	}

	if constructor, found := subTypeToConstructorMap[subType]; found {
		return constructor.(func(string, interface{}, ...map[string]interface{}) error)("", openApiError.Error())
	}

	return ferrors.NewErrInternalError("Invalid v4 API error sub-type", openApiError.Error())
}
