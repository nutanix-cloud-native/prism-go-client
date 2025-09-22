package v4

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/nutanix-cloud-native/prism-go-client/facade"
	clusterClient "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
)

// GetCategorisedV4ApiCallError returns categorised and wrapped errors.
func GetCategorisedV4ApiCallError(err error) error {
	if err == nil {
		return nil
	}

	var openApiError clusterClient.GenericOpenAPIError
	if errors.As(err, &openApiError) {
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

	oneOfError, errorType, err := getErrorAndTypeFrom(errResp)
	if err != nil {
		return fmt.Errorf("failed to get error and type from `GenericOpenApiError` with: %w", err)
	}

	if strings.Contains(errorType, "SchemaValidationError") {
		return facade.NewApiError(facade.ErrSchemaValidationError, openApiError)
	}
	if strings.Contains(errorType, "AppMessage") {
		return getCategorisedV4ApiErrorFromAppMessages(openApiError, oneOfError)
	}

	return fmt.Errorf("invalid value of field `Data` in APIResponse: %s", openApiError.Error())
}

func getErrorAndTypeFrom(errResp sampleErrorResponse) (interface{}, string, error) {
	errorDataMap, ok := errResp.Data.(map[string]interface{})
	if !ok {
		return nil, "", facade.NewUnexpectedTypeError(map[string]interface{}{}, errResp.Data)
	}

	errorTypeIntf, ok := errorDataMap["$errorItemDiscriminator"]
	if !ok {
		return nil, "", fmt.Errorf("field `$errorItemDiscriminator` missing from error data: %v", errResp)
	}
	errorType, ok := errorTypeIntf.(string)
	if !ok {
		return nil, "", facade.NewUnexpectedTypeError("", errorDataMap["$errorItemDiscriminator"])
	}

	oneOfError, ok := errorDataMap["error"]
	if !ok {
		return nil, "", fmt.Errorf("field `error` missing from error data: %v", errResp)
	}

	return oneOfError, errorType, nil
}

func getCategorisedV4ApiErrorFromAppMessages(openApiError clusterClient.GenericOpenAPIError, appMessagesIntf interface{}) error {
	apiErrType := errors.New("unknown error")

	appMessages, ok := appMessagesIntf.([]interface{})
	if !ok {
		return facade.NewUnexpectedTypeError([]interface{}{}, appMessagesIntf)
	}

	for _, appMessageIntf := range appMessages {
		appMessage, ok := appMessageIntf.(map[string]interface{})
		if !ok {
			return facade.NewUnexpectedTypeError(map[string]interface{}{}, appMessageIntf)
		}

		errorGroupIntf, ok := appMessage["errorGroup"]
		if !ok {
			return fmt.Errorf("field `errorGroup` missing from AppMessage: %v", appMessage)
		}

		errorGroup, ok := errorGroupIntf.(string)
		if !ok {
			return facade.NewUnexpectedTypeError("", errorGroupIntf)
		}

		apiErrType = getV4ApiErrorSubTypeForErrorGroup(errorGroup)
	}

	return facade.NewApiError(apiErrType, openApiError)
}

func getV4ApiErrorSubTypeForErrorGroup(errorGroup string) error {
	searchKeyToSubTypeMap := map[string]error{
		"AUTHORIZATION": facade.ErrAuthorizationError,
		"NOT_FOUND":     facade.ErrResourceNotFound,
		"RATE_LIMIT":    facade.ErrRateLimitExceeded,
		"SERVICE_ERROR": facade.ErrInternalServiceError,
		"INVALID_INPUT": facade.ErrInvalidInput,
	}

	for key, subType := range searchKeyToSubTypeMap {
		if strings.Contains(errorGroup, key) {
			return subType
		}
	}

	return errors.New(errorGroup)
}
