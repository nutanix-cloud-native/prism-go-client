package facadetest

import (
	"errors"
	"net"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4 "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	vmmConfigModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmmErrorModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"k8s.io/utils/ptr"
)

func TestCategoriseV4APICallErrorForVM(t *testing.T) {
	testCases := []struct {
		name                 string
		responseBuilder      func() *vmmConfigModels.GetVmApiResponse
		sampleErr            error
		expectedErrorType    ferrors.ErrorType
		expectedErrorSubType ferrors.ErrorSubType
	}{
		{
			name:                 "Network/Transport error",
			responseBuilder:      func() *vmmConfigModels.GetVmApiResponse { return nil },
			sampleErr:            &net.OpError{},
			expectedErrorType:    ferrors.ErrorTypeNetworkError,
			expectedErrorSubType: "",
		},
		{
			name:                 "Random error",
			responseBuilder:      func() *vmmConfigModels.GetVmApiResponse { return nil },
			sampleErr:            errors.New("random error"),
			expectedErrorType:    ferrors.ErrorTypeUncategorisedError,
			expectedErrorSubType: "",
		},
		{
			name: "ApiErrorResponse : AppMessage : Uncategorised error group",
			responseBuilder: func() *vmmConfigModels.GetVmApiResponse {
				appMessage := vmmErrorModels.NewAppMessage()
				appMessage.Code = ptr.To("TEST-1000")
				appMessage.ErrorGroup = ptr.To("TEST_ERROR")
				appMessages := []vmmErrorModels.AppMessage{*appMessage}

				errorResp := vmmErrorModels.NewErrorResponse()
				errorResp.SetError(appMessages)

				resp := vmmConfigModels.NewGetVmApiResponse()
				resp.SetData(*errorResp)
				return resp
			},
			sampleErr:            errors.New(""),
			expectedErrorType:    ferrors.ErrorTypeV4ApiError,
			expectedErrorSubType: ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiUncategorisedError),
		},
		{
			name: "ApiErrorResponse : AppMessage : Rate limit error group",
			responseBuilder: func() *vmmConfigModels.GetVmApiResponse {
				appMessage := vmmErrorModels.NewAppMessage()
				appMessage.Code = ptr.To("PLAT-10003")
				appMessage.ErrorGroup = ptr.To("RATE_LIMIT_EXCEEDED")
				appMessages := []vmmErrorModels.AppMessage{*appMessage}

				errorResp := vmmErrorModels.NewErrorResponse()
				errorResp.SetError(appMessages)

				resp := vmmConfigModels.NewGetVmApiResponse()
				resp.SetData(*errorResp)
				return resp
			},
			sampleErr:            errors.New(""),
			expectedErrorType:    ferrors.ErrorTypeV4ApiError,
			expectedErrorSubType: ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiRateLimitError),
		},
		{
			name: "ApiErrorResponse : SchemaValidationError",
			responseBuilder: func() *vmmConfigModels.GetVmApiResponse {
				schemaValidationError := *vmmErrorModels.NewSchemaValidationError()

				errorResp := vmmErrorModels.NewErrorResponse()
				errorResp.SetError(schemaValidationError)

				resp := vmmConfigModels.NewGetVmApiResponse()
				resp.SetData(*errorResp)
				return resp
			},
			sampleErr:            errors.New(""),
			expectedErrorType:    ferrors.ErrorTypeV4ApiError,
			expectedErrorSubType: ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiSchemaValidationError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := v4.GetCategorisedV4ApiCallError[*vmmConfigModels.GetVmApiResponse, *vmmErrorModels.OneOfErrorResponseError](tt.responseBuilder(), tt.sampleErr)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if ferr, ok := err.(*ferrors.Err); ok {
				if ferr.Type != tt.expectedErrorType {
					t.Errorf("Expected error type %v, got %v", tt.expectedErrorType, ferr.Type)
				}
				if ferr.SubType != tt.expectedErrorSubType {
					t.Errorf("Expected error sub-type %v, got %v", tt.expectedErrorSubType, ferr.SubType)
				}
			} else {
				t.Errorf("Expected error to be of type ferrors.Err, got %T", err)
			}
		})
	}
}
