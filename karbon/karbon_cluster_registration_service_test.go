package karbon

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

const (
	noCategoryMappingErrorMsg  = "No category mappings provided"
	noUUIDErrorMsg             = "uuid in body is required"
	noTaskIDCreateRespErrorMsg = "Task ID was not include the create response."
	timeout                    = 60 * time.Second
	taskSucceed                = "SUCCEEDED"
	taskFailed                 = "FAILED"
	taskAborted                = "ABORTED"
	taskSuspended              = "SUSPENDED"
)

func validateK8sClusterRegistration(t *testing.T, k8sClusterReg *K8sClusterRegistration) {
	assert.NotEmpty(t, *k8sClusterReg.Name)
	assert.NotEmpty(t, k8sClusterReg.Status)
	assert.NotEmpty(t, k8sClusterReg.UUID)
	assert.NotEmpty(t, k8sClusterReg.CategoriesMapping)
	assert.NotZero(t, k8sClusterReg.CategoriesMapping)
}

func validateK8sClusterRegistrationGetResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string,
	responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.NotZero(t, len(responseGetReg.CategoriesMapping))
}

func validateK8sClusterRegistrationDeleteResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string,
	delResp *K8sClusterRegistrationDeleteResponse,
) {
	assert.NotEmpty(t, delResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, delResp.ClusterName)
	assert.NotEmpty(t, delResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, delResp.ClusterUUID)
	assert.NotEmpty(t, *delResp.TaskUUID)
}

func validateK8sClusterRegistrationCreateResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string,
	createResp *K8sCreateClusterRegistrationResponse,
) {
	assert.NotEmpty(t, createResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, *createResp.ClusterName)
	assert.NotEmpty(t, createResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, *createResp.ClusterUUID)
	assert.NotEmpty(t, *createResp.TaskUUID)
}

func validateK8sClusterRegistrationList(t *testing.T, clusterRegList *K8sClusterRegistrationList) {
	for _, k8sRegistration := range *clusterRegList {
		validateK8sClusterRegistration(t, k8sRegistration)
	}
}

/*func validateK8sClusterRegistrationTaskStatus(t *testing.T, kctx context.Context, taskID string, creds prismgoclient.Credentials) {
	v3Client, err := v3.NewV3Client(creds)
	assert.NoError(t, err)
	assert.NotNil(t, v3Client)
	timeStart := time.Now()
	for {
		timeCur := time.Now()
		if timeCur.Sub(timeStart) > timeout {
			assert.FailNow(t, fmt.Sprintf("Task %s was failed: Task was not completed by timeout.\n", taskID))
		}
		responseGetTask, err := v3Client.V3.GetTask(kctx, taskID)
		assert.NoError(t, err)
		assert.NotNil(t, responseGetTask.Status)
		taskStatus := *responseGetTask.Status
		if taskStatus == taskSucceed {
			break
		} else if taskStatus == taskFailed || taskStatus == taskAborted || taskStatus == taskSuspended {
			assert.FailNow(t, fmt.Sprintf("Task %s was failed: the task status is %s.\n", taskID, taskStatus))
		}
	}
}*/

func TestKarbonCreateClusterRegistration(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := strings.ToLower("cluster1")
	test_cluster_uuid := strings.ToLower("634F5852-CE1B-465C-A95A-9B8DFFBDDE42")
	test_category_mapping := map[string]string{
		fmt.Sprintf("kubernetes-io-cluster-%s", test_cluster_name): "owned",
		"KubernetesClusterName": test_cluster_name,
		"KubernetesClusterUUID": test_cluster_uuid,
	}
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, responseGetReg)
		// Registration exists. delete it so that we can create it
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		// Valid k8s cluster registration delete response
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
		// Get task uuid status and to check if the task complete successfully before timeout
		assert.NotNil(t, responseDelReg.TaskUUID)
		//validateK8sClusterRegistrationTaskStatus(t, kctx, *responseDelReg.TaskUUID, creds)
	}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
	}

	// returns type K8sCreateClusterRegistrationResponse
	responseCreateReg, err := nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	assert.NoError(t, err)

	// Valid k8s cluster registration create response
	validateK8sClusterRegistrationCreateResponse(t, test_cluster_name, test_cluster_uuid, responseCreateReg)

	// Get task uuid status and to check if the task complete successfully before timeout
	assert.NotNil(t, responseCreateReg.TaskUUID)
	//validateK8sClusterRegistrationTaskStatus(t, kctx, *responseCreateReg.TaskUUID, creds)
}

func TestKarbonCreateClusterRegistrationWithNoCategory(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := strings.ToLower("cluster2")
	test_cluster_uuid := strings.ToLower("E33FD0FD-5673-45FA-825D-7EF869A91577")
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, responseGetReg)
		// Registration exists. delete it so that we can create it
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		// Valid k8s cluster registration delete response
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
		// Get task uuid status and to check if the task complete successfully before timeout
		assert.NotNil(t, responseDelReg.TaskUUID)
		//validateK8sClusterRegistrationTaskStatus(t, kctx, *responseDelReg.TaskUUID, creds)
	}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:     &test_cluster_name,
		UUID:     test_cluster_uuid,
		Metadata: test_metadata,
	}

	// check if the error is expected
	_, err = nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	if assert.Error(t, err) {
		assert.Contains(t, fmt.Sprint(err), noCategoryMappingErrorMsg)
	}
}

func TestKarbonCreateClusterRegistrationWithNoUUID(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := strings.ToLower("cluster3")
	test_category_mapping := map[string]string{
		fmt.Sprintf("kubernetes-io-cluster-%s", test_cluster_name): "owned",
		"KubernetesClusterName": test_cluster_name,
	}
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
	}

	// check if the error is expected
	_, err = nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	if assert.Error(t, err) {
		assert.Contains(t, fmt.Sprint(err), noUUIDErrorMsg)
	}
}

func TestKarbonGetK8sRegistrationList(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	// returns type K8sCreateClusterRegistrationResponse
	response, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistrationList(kctx)
	assert.NoError(t, err)
	validateK8sClusterRegistrationList(t, response)
}
