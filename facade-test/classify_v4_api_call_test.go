package facadetest

import (
	"errors"
	"testing"

	"github.com/nutanix-cloud-native/prism-go-client/facade/ferrors"
	v4 "github.com/nutanix-cloud-native/prism-go-client/facade/v4"
	vmmConfigModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	vmmErrorModels "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"k8s.io/utils/ptr"
)

func TestClassifyV4APICallErrorForVM(t *testing.T) {
	appMessage := vmmErrorModels.NewAppMessage()
	appMessage.Code = ptr.To("TEST-1000")
	appMessage.ErrorGroup = ptr.To("TEST_ERROR")
	appMessages := []vmmErrorModels.AppMessage{*appMessage}

	errorResp := vmmErrorModels.NewErrorResponse()
	errorResp.SetError(appMessages)

	sampleResponse := vmmConfigModels.NewGetVmApiResponse()
	sampleResponse.SetData(*errorResp)
	expectedType := ferrors.ErrorTypeV4ApiError
	expectedSubType := ferrors.ErrorSubType(ferrors.ErrorSubTypeV4ApiUncategorisedError)
	sampleErr := errors.New("")

	err := v4.GetCategorisedV4ApiCallError[*vmmConfigModels.GetVmApiResponse, *vmmErrorModels.OneOfErrorResponseError](sampleResponse, sampleErr)
	if err == nil {
		t.Fatal("Expected error to be returned, got nil")
	}

	if ferr, ok := err.(*ferrors.Err); ok {
		if ferr.Type != expectedType {
			t.Errorf("Expected error type %v, got %v", expectedType, ferr.Type)
		}
		if ferr.SubType != expectedSubType {
			t.Errorf("Expected error sub-type %v, got %v", expectedSubType, ferr.SubType)
		}
	} else {
		t.Errorf("Expected error to be of type ferrors.Err, got %T", err)
	}
}
