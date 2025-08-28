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

func TestClassifyV4APICallErrorCases(t *testing.T) {
	tests := []struct {
		name            string
		responseBuilder func() v4.ApiResponse
		err             error
		expectedType    ferrors.ErrorType
		expectedSubType ferrors.ErrorSubType
	}{
		{
			name:            "Network error",
			responseBuilder: func() v4.ApiResponse { return nil },
			err:             &net.OpError{},
			expectedType:    ferrors.ErrorTypeNetworkError,
		},
		{
			name:            "Random error",
			responseBuilder: func() v4.ApiResponse { return nil },
			err:             errors.New(""),
			expectedType:    ferrors.ErrorTypeUncategorisedError,
		},
		{
			name:            "V4 api 4xx sample error",
			err:             errors.New(""),
			expectedType:    ferrors.ErrorTypeV4ApiError,
			expectedSubType: ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiAuthorizationError),
			responseBuilder: func() v4.ApiResponse {
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v4.GetCategorisedV4ApiCallError[v4.ApiResponse, v4.ApiErrorResponse](tt.responseBuilder(), tt.err)
			if err == nil {
				t.Fatal("Expected error to be returned, got nil")
			}

			if ferr, ok := err.(ferrors.Err); ok {
				if ferr.Type != tt.expectedType {
					t.Errorf("Expected error type %v, got %v", tt.expectedType, ferr.Type)
				}
				if ferr.SubType != tt.expectedSubType {
					t.Errorf("Expected error sub-type %v, got %v", tt.expectedSubType, ferr.SubType)
				}
			} else {
				t.Errorf("Expected error to be of type ferrors.Err, got %T", err)
			}
		})
	}
}
