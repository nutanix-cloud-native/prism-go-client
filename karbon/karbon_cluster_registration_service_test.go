package karbon

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nutanix-cloud-native/prism-go-client/internal/testhelpers"
)

const (
	noCategoryMappingErrorMsg = "No category mappings provided"
	noUUIDErrorMsg            = "uuid in body is required"
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

func validateK8sClusterRegistrationGetResponseWithClusterInfo(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string,
	test_cluster_info *K8sClusterInfo, responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.NotEmpty(t, responseGetReg.ClusterInfo)
	assert.NotEmpty(t, responseGetReg.ClusterInfo.K8sDistribution)
	assert.NotEmpty(t, responseGetReg.ClusterInfo.K8sVersion)
	assert.Empty(t, responseGetReg.AddonsInfo)
}

func validateK8sClusterRegistrationGetResponseWithAddonInfo(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string,
	test_cluster_addon_info *K8sClusterAddonInfo, responseGetReg *K8sClusterRegistration,
) {
	assert.NotEmpty(t, *responseGetReg.Name)
	assert.Equal(t, expected_k8s_cluster_name, *responseGetReg.Name)
	assert.NotEmpty(t, responseGetReg.Status)
	assert.NotEmpty(t, responseGetReg.UUID)
	assert.Equal(t, expected_k8s_cluster_uuid, responseGetReg.UUID)
	assert.Empty(t, responseGetReg.ClusterInfo)
	assert.NotEmpty(t, responseGetReg.AddonsInfo)
	assert.NotEmpty(t, responseGetReg.AddonsInfo[0].AddonName)
	assert.NotEmpty(t, responseGetReg.AddonsInfo[0].AddonVersion)
	assert.NotEmpty(t, responseGetReg.AddonsInfo[1].AddonName)
	assert.NotEmpty(t, responseGetReg.AddonsInfo[1].AddonVersion)
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

func validateK8sClusterRegistrationInfoResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string, clusterRegInfoResp *K8sUpdateClusterRegistrationInfoResponse) {
	assert.NotEmpty(t, clusterRegInfoResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, *clusterRegInfoResp.ClusterName)
	assert.NotEmpty(t, clusterRegInfoResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, *clusterRegInfoResp.ClusterUUID)
}

func validateK8sClusterRegistrationAddonInfoResponse(t *testing.T, expected_k8s_cluster_name, expected_k8s_cluster_uuid string, clusterRegAddonInfoResp *K8sUpdateClusterRegistrationAddonInfoResponse) {
	assert.NotEmpty(t, clusterRegAddonInfoResp.ClusterName)
	assert.Equal(t, expected_k8s_cluster_name, *clusterRegAddonInfoResp.ClusterName)
	assert.NotEmpty(t, clusterRegAddonInfoResp.ClusterUUID)
	assert.Equal(t, expected_k8s_cluster_uuid, *clusterRegAddonInfoResp.ClusterUUID)
}

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
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
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
	validateK8sClusterRegistrationCreateResponse(t, test_cluster_name, test_cluster_uuid, responseCreateReg)

	// TODO get task uuid status
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
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
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

func TestKarbonCreateClusterRegistrationAndSetInfo(t *testing.T) {
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
		validateK8sClusterRegistrationDeleteResponse(t, test_cluster_name, test_cluster_uuid, responseDelReg)
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
