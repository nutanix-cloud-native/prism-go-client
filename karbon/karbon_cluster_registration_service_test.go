package karbon

import (
	"errors"
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

var (
	NoCategoryMappingError = errors.New("error: map[Cluster Create:Failed to validate the category mapping. No category mappings provided.]")
	NoUUIDError            = errors.New("error: uuid in body is required")
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
		assert.Equal(t, NoCategoryMappingError, err)
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
		assert.Equal(t, NoUUIDError, err)
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
