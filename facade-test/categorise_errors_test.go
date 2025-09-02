package facadetest

import (
	"net"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4 "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	clusterClient "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	vmmConfigModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmmErrorModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"k8s.io/utils/ptr"
)

func TestCategoriseV4APICallErrorForVM(t *testing.T) {
	testCases := []struct {
		name                 string
		responseBuilder      func() *vmmConfigModels.GetVmApiResponse // provide one of responseBuilder or sampleError
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
			expectedErrorType:    ferrors.ErrorTypeV4ApiError,
			expectedErrorSubType: ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiSchemaValidationError),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var in error
			if tt.sampleErr != nil {
				in = tt.sampleErr
			} else {
				jsBytes, err := tt.responseBuilder().MarshalJSON()
				if err != nil {
					t.Fatal("Unable to marshal response json")
				}

				in = clusterClient.GenericOpenAPIError{
					Body:   jsBytes,
					Status: "4xx",
				}
			}

			err := v4.GetCategorisedV4ApiCallError(in)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}

			if ferr, ok := err.(*ferrors.FacadeError); ok {
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
