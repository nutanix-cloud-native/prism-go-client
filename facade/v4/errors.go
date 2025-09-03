package v4

import (
	"encoding/json"
	"fmt"
	"strings"

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

	return fmt.Errorf("API call failed with error: %v", err)
}

type sampleErrorResponse struct {
	Data any `json:"data"`
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
		return ferrors.NewUnexpectedTypeError(map[string]interface{}{}, errResp.Data)
	}
	errorType, ok := errorData["$errorItemDiscriminator"].(string)
	if !ok {
		return ferrors.NewUnexpectedTypeError("", errorData["$errorItemDiscriminator"])
	}

	if strings.Contains(errorType, "SchemaValidationError") {
		return ferrors.NewApiError(ferrors.ErrSchemaValidationError, openApiError)
	}
	if strings.Contains(errorType, "AppMessage") {
		return getCategorisedV4ApiErrorFromAppMessages(openApiError, errorData["error"])
	}

	return fmt.Errorf("invalid value of field `Data` in APIResponse: %s", openApiError.Error())
}

func getCategorisedV4ApiErrorFromAppMessages(openApiError clusterClient.GenericOpenAPIError, appMessagesIntf interface{}) error {
	apiErrType := ferrors.ErrUnknownError

	appMessages, ok := appMessagesIntf.([]interface{})
	if !ok {
		return ferrors.NewUnexpectedTypeError([]interface{}{}, appMessagesIntf)
	}

	for _, appMessageIntf := range appMessages {
		appMessage, ok := appMessageIntf.(map[string]interface{})
		if !ok {
			return ferrors.NewUnexpectedTypeError(map[string]interface{}{}, appMessageIntf)
		}

		errorGroupIntf, ok := appMessage["errorGroup"]
		if !ok {
			return fmt.Errorf("field `errorGroup` missing from AppMessage: %v", appMessage)
		}

		errorGroup, ok := errorGroupIntf.(string)
		if !ok {
			return ferrors.NewUnexpectedTypeError("", errorGroupIntf)
		}

		newErr := getV4ApiErrorSubTypeForErrorGroup(errorGroup)
		// DOUBT ranking???
		if newErr != nil {
			apiErrType = newErr
		}
	}

	return ferrors.NewApiError(apiErrType, openApiError)
}

func getV4ApiErrorSubTypeForErrorGroup(errorGroup string) error {
	searchKeyToSubTypeMap := map[string]error{
		"AUTHORIZATION": ferrors.ErrAuthorizationError,
		"NOT_FOUND":     ferrors.ErrResourceNotFound,
		"RATE_LIMIT":    ferrors.ErrRateLimitExceeded,
		"SERVICE_ERROR": ferrors.ErrInternalServiceError,
		"INVALID_INPUT": ferrors.ErrInvalidInput,
	}

	for key, subType := range searchKeyToSubTypeMap {
		if strings.Contains(errorGroup, key) {
			return subType
		}
	}

	return nil
}
