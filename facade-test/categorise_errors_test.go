package facadetest

import (
	"errors"
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
		responseBuilder      func() *vmmConfigModels.GetVmApiResponse
		expectedApiErrorType error
	}{
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
			expectedApiErrorType: ferrors.ErrUnknownError,
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
			expectedApiErrorType: ferrors.ErrRateLimitExceeded,
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
			expectedApiErrorType: ferrors.ErrSchemaValidationError,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			jsBytes, err := tt.responseBuilder().MarshalJSON()
			if err != nil {
				t.Errorf("Unable to marshal response json")
			}

			in := clusterClient.GenericOpenAPIError{
				Body:   jsBytes,
				Status: "4xx",
			}
			err = v4.GetCategorisedV4ApiCallError(in)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}

			var apiError *ferrors.ApiError
			found := errors.As(err, &apiError)
			if !found {
				t.Errorf("Invalid error type")
			}
			if tt.expectedApiErrorType == nil || tt.expectedApiErrorType.Error() != apiError.Type.Error() {
				t.Errorf("Invalid api error type")
			}
		})
	}
}
