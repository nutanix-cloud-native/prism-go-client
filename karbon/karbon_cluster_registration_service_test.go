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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

const (
	_noCategoryMappingErrorMsg = "No category mappings provided"
	_noUUIDErrorMsg            = "uuid in body is required"
	_noK8sDistributionMsg      = "k8s_distribution in body is required"
	_testWaitTimeout           = 60 * time.Second
	_testWaitInterval          = 2 * time.Second
	_taskSucceed               = "SUCCEEDED"
	_taskFailed                = "FAILED"
	_taskAborted               = "ABORTED"
	_taskSuspended             = "SUSPENDED"
)

func validateK8sClusterRegistration(t *testing.T, k8sClusterReg *K8sClusterRegistration) {
	assert.NotEmpty(t, *k8sClusterReg.Name)
	assert.NotEmpty(t, k8sClusterReg.Status)
	assert.NotEmpty(t, k8sClusterReg.UUID)
	assert.NotEmpty(t, k8sClusterReg.K8sDistribution)
	assert.NotEmpty(t, k8sClusterReg.CategoriesMapping)
	assert.NotZero(t, k8sClusterReg.CategoriesMapping)
}

func validateK8sClusterRegistrationGetResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid,
	expected_k8s_distribution string, responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.NotEmpty(t, responseGetReg.K8sDistribution)
	assert.Equal(t, expected_k8s_distribution, responseGetReg.K8sDistribution)
	assert.NotZero(t, len(responseGetReg.CategoriesMapping))
}

func validateK8sClusterRegistrationGetResponseWithClusterInfo(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid,
	expected_k8s_distribution string, responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.NotEmpty(t, responseGetReg.K8sDistribution)
	assert.Equal(t, expected_k8s_distribution, responseGetReg.K8sDistribution)
	assert.NotEmpty(t, responseGetReg.ClusterInfo)
	assert.Empty(t, responseGetReg.AddonsInfo)
}

func validateK8sClusterRegistrationGetResponseWithAddonInfo(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid,
	expected_k8s_distribution string, responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.NotEmpty(t, responseGetReg.K8sDistribution)
	assert.Equal(t, expected_k8s_distribution, responseGetReg.K8sDistribution)
	assert.Empty(t, responseGetReg.ClusterInfo)
	assert.NotEmpty(t, responseGetReg.AddonsInfo)
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

func validateK8sClusterRegistrationTaskStatus(t *testing.T, kctx context.Context, v3Client *v3.Client, taskID string, creds prismgoclient.Credentials) {
	maxRetryTimes := _testWaitTimeout / _testWaitInterval
	ticker := time.NewTicker(_testWaitInterval)
	retryTimes := 0
	for range ticker.C {
		if retryTimes > int(maxRetryTimes) {
			require.FailNow(t, fmt.Sprintf("Task %s was failed: Task was not completed by timeout.\n", taskID))
			break
		}
		retryTimes += 1
		responseGetTask, err := v3Client.V3.GetTask(kctx, taskID)
		require.NoError(t, err)
		require.NotNil(t, responseGetTask.Status)
		taskStatus := *responseGetTask.Status
		if taskStatus == _taskSucceed {
			break
		} else if taskStatus == _taskFailed || taskStatus == _taskAborted || taskStatus == _taskSuspended {
			assert.FailNow(t, fmt.Sprintf("Task %s was failed: the task status is %s.\n", taskID, taskStatus))
		}
	}
}

func validateK8sClusterRegistrationInfoResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string, clusterRegInfoResp *K8sUpdateClusterRegistrationInfoResponse) {
	assert.NotEmpty(t, clusterRegInfoResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, clusterRegInfoResp.ClusterName)
	assert.NotEmpty(t, clusterRegInfoResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, clusterRegInfoResp.ClusterUUID)
}

func validateK8sClusterRegistrationAddonInfoResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string, clusterRegAddonInfoResp *K8sUpdateClusterRegistrationAddonInfoResponse) {
	assert.NotEmpty(t, clusterRegAddonInfoResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, clusterRegAddonInfoResp.ClusterName)
	assert.NotEmpty(t, clusterRegAddonInfoResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, clusterRegAddonInfoResp.ClusterUUID)
}

func TestKarbonCreateClusterRegistration(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)

	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)
	require.NotNil(t, nkeClient)

	v3Client, err := v3.NewV3Client(creds, v3.WithRoundTripper(interceptor))
	require.NoError(t, err)
	require.NotNil(t, v3Client)

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
	test_k8s_distribution := "Openshift"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)
		// Registration exists. delete it so that we can create it
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)

		// Valid k8s cluster registration delete response
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)

		// Get task uuid status and to check if the task complete successfully before timeout
		require.NotNil(t, responseDelReg.TaskUUID)
		validateK8sClusterRegistrationTaskStatus(t, kctx, v3Client, *responseDelReg.TaskUUID, creds)
	}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              &test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
		K8sDistribution:   &test_k8s_distribution,
	}

	// returns type K8sCreateClusterRegistrationResponse
	responseCreateReg, err := nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	assert.NoError(t, err)

	// Valid k8s cluster registration create response
	validateK8sClusterRegistrationCreateResponse(t, test_cluster_name, test_cluster_uuid, responseCreateReg)

	// Get task uuid status and to check if the task complete successfully before timeout
	require.NotNil(t, responseCreateReg.TaskUUID)
	validateK8sClusterRegistrationTaskStatus(t, kctx, v3Client, *responseCreateReg.TaskUUID, creds)
}

func TestKarbonCreateClusterRegistrationWithNoCategory(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)

	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)
	require.NotNil(t, nkeClient)

	v3Client, err := v3.NewV3Client(creds, v3.WithRoundTripper(interceptor))
	require.NoError(t, err)
	require.NotNil(t, v3Client)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := strings.ToLower("cluster2")
	test_cluster_uuid := strings.ToLower("E33FD0FD-5673-45FA-825D-7EF869A91577")
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}
	test_k8s_distribution := "Openshift"

	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)
		// Registration exists. delete it so that we can create it
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		// Valid k8s cluster registration delete response
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
		// Get task uuid status and to check if the task complete successfully before timeout
		assert.NotNil(t, responseDelReg.TaskUUID)
		validateK8sClusterRegistrationTaskStatus(t, kctx, v3Client, *responseDelReg.TaskUUID, creds)
	}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:            &test_cluster_name,
		UUID:            &test_cluster_uuid,
		Metadata:        test_metadata,
		K8sDistribution: &test_k8s_distribution,
	}

	// check if the error is expected
	_, err = nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	if assert.Error(t, err) {
		assert.Contains(t, fmt.Sprint(err), _noCategoryMappingErrorMsg)
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
	test_k8s_distribution := "Openshift"

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
		K8sDistribution:   &test_k8s_distribution,
	}

	// check if the error is expected
	_, err = nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	if assert.Error(t, err) {
		assert.Contains(t, fmt.Sprint(err), _noUUIDErrorMsg)
	}
}

func TestKarbonCreateClusterRegistrationWithNoK8sDistribution(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testhelpers.CredentialsFromEnvironment(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := strings.ToLower("cluster4")
	test_cluster_uuid := strings.ToLower("FE139669-50F3-4307-8867-F7630241F0DF")
	test_category_mapping := map[string]string{
		fmt.Sprintf("kubernetes-io-cluster-%s", test_cluster_name): "owned",
		"KubernetesClusterName": test_cluster_name,
	}
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              &test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
	}

	// check if the error is expected
	_, err = nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	if assert.Error(t, err) {
		assert.Contains(t, fmt.Sprint(err), _noK8sDistributionMsg)
	}
}

func TestKarbonCreateClusterRegistrationAndSetInfo(t *testing.T) {
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
	test_k8s_distribution := "Openshift"

	t.Log("Get Cluster registration")
	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		t.Log("Get Cluster registration exists")
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)
		// Registration exists. delete it so that we can create it
		t.Log("Delete Cluster registration")
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
	}

	t.Log("Create Cluster registration")
	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              &test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		K8sDistribution:   &test_k8s_distribution,
		Metadata:          test_metadata,
	}

	// returns type K8sCreateClusterRegistrationResponse
	responseCreateReg, err := nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	assert.NoError(t, err)
	validateK8sClusterRegistrationCreateResponse(t, test_cluster_name, test_cluster_uuid, responseCreateReg)

	// TODO get task uuid status

	t.Log("Update K8S Info")
	test_k8s_distribution = "Openshift"
	test_k8s_version := "v1.25.0"

	test_cluster_info := map[string]string{
		"k8s_version": test_k8s_version,
	}

	updateInfoRequest := &K8sUpdateClusterRegistrationInfoRequest{
		ClusterInfo: test_cluster_info,
	}
	responseUpdateInfo, err := nkeClient.ClusterRegistrationOperations.UpdateK8sRegistrationInfo(kctx, test_cluster_uuid, updateInfoRequest)
	assert.NoError(t, err)
	validateK8sClusterRegistrationInfoResponse(t, test_cluster_name, test_cluster_uuid, responseUpdateInfo)

	t.Log("Get Cluster registration")
	responseGetReg, err = nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		t.Log("Get Cluster registration exists")
		validateK8sClusterRegistrationGetResponseWithClusterInfo(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)
		// Registration exists. delete it so that we can create it
		t.Log("Delete Cluster registration")
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
	}
}

func TestKarbonCreateClusterRegistrationAndAddonSetInfo(t *testing.T) {
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
	test_k8s_distribution := "Openshift"

	t.Log("Get Cluster registration")
	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		t.Log("Get Cluster registration exists")
		validateK8sClusterRegistrationGetResponse(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)
		// Registration exists. delete it so that we can create it
		t.Log("Delete Cluster registration")
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
	}

	t.Log("Create Cluster registration")
	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              &test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
		K8sDistribution:   &test_k8s_distribution,
	}

	// returns type K8sCreateClusterRegistrationResponse
	responseCreateReg, err := nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	assert.NoError(t, err)
	validateK8sClusterRegistrationCreateResponse(t, test_cluster_name, test_cluster_uuid, responseCreateReg)

	// TODO get task uuid status

	t.Log("Update CSI Info")
	test_addon_name := "CSI"
	test_addon_version := "v3.x"
	test_cluster_addon_info := map[string]string{
		test_addon_name: test_addon_version,
	}
	updateAddonInfoRequest := &K8sUpdateClusterRegistrationAddonInfoRequest{
		ClusterAddonInfo: test_cluster_addon_info,
	}
	responseUpdateAddonInfo, err := nkeClient.ClusterRegistrationOperations.UpdateK8sRegistrationAddonInfo(kctx, test_cluster_uuid, test_addon_name, updateAddonInfoRequest)
	assert.NoError(t, err)
	validateK8sClusterRegistrationAddonInfoResponse(t, test_cluster_name, test_cluster_uuid, responseUpdateAddonInfo)

	t.Log("Update CNDS Info")
	test_addon_name = "CNDS"
	test_addon_version = "v1.x"
	test_cluster_addon_info = map[string]string{
		test_addon_name: test_addon_version,
	}
	updateAddonInfoRequest = &K8sUpdateClusterRegistrationAddonInfoRequest{
		ClusterAddonInfo: test_cluster_addon_info,
	}
	responseUpdateAddonInfo, err = nkeClient.ClusterRegistrationOperations.UpdateK8sRegistrationAddonInfo(kctx, test_cluster_uuid, test_addon_name, updateAddonInfoRequest)
	assert.NoError(t, err)
	validateK8sClusterRegistrationAddonInfoResponse(t, test_cluster_name, test_cluster_uuid, responseUpdateAddonInfo)

	t.Log("Get Cluster registration")
	responseGetReg, err = nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	assert.NoError(t, err)
	validateK8sClusterRegistrationGetResponseWithAddonInfo(t, test_cluster_name, test_cluster_uuid, test_k8s_distribution, responseGetReg)

	t.Log("Delete Cluster registration")
	responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
	assert.NoError(t, err)
	validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
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
