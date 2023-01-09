package karbon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/keploy/go-sdk/integrations/khttpclient"
	"github.com/keploy/go-sdk/keploy"
	"github.com/keploy/go-sdk/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func printK8sClusterRegistration(k8sClusterReg *K8sClusterRegistration) {
	fmt.Printf("Type: K8sClusterRegistration\n")
	fmt.Printf("|-Name: %s\n", *k8sClusterReg.Name)
	fmt.Printf("|-Status: %s\n", k8sClusterReg.Status)
	fmt.Printf("|-UUID: %s\n", k8sClusterReg.UUID)
	// fmt.Printf("|-K8s Identity Name: %s\n", k8sClusterReg.Identity.Name)
	// fmt.Printf("|-K8s Identity UUID: %s\n", *k8sClusterReg.Identity.UUID)
	// fmt.Printf("|-K8s Identity Kind: %s\n", *k8sClusterReg.Identity.Kind)
	fmt.Printf("|-Category Mapping: %#v\n", k8sClusterReg.CategoriesMapping)
}

func printK8sClusterRegistrationDeleteResponse(delResp *K8sClusterRegistrationDeleteResponse) {
	fmt.Printf("Type: K8sClusterRegistrationDeleteResponse\n")
	fmt.Printf("|-Cluster Name: %s\n", delResp.ClusterName)
	fmt.Printf("|-Cluster UUID: %s\n", delResp.ClusterUUID)
	fmt.Printf("|-Task UUID: %s\n", *delResp.TaskUUID)
}

func printK8sClusterRegistrationCreateResponse(createResp *K8sCreateClusterRegistrationResponse) {
	fmt.Printf("Type: K8sCreateClusterRegistrationResponse\n")
	fmt.Printf("|-Cluster Name: %s\n", *createResp.ClusterName)
	fmt.Printf("|-Cluster UUID: %s\n", *createResp.ClusterUUID)
	fmt.Printf("|-Task UUID: %s\n", *createResp.TaskUUID)
}

func printK8sClusterRegistrationList(clusterRegList *K8sClusterRegistrationList) {
	for _, k8sRegistration := range *clusterRegList {
		printK8sClusterRegistration(k8sRegistration)
	}
}

func TestKarbonCreateClusterRegistration(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testCredsFromEnv(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})

	test_cluster_name := "cluster1"
	test_cluster_uuid := "634F5852-CE1B-465C-A95A-9B8DFFBDDE42"
	test_category_mapping := map[string]string{
		fmt.Sprintf("kubernetes-io-cluster-%s", test_cluster_name): "owned",
		"KubernetesClusterName": test_cluster_name,
		"KubernetesClusterUUID": test_cluster_uuid,
	}
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	fmt.Println("Get Cluster registration")
	responseGetReg, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistration(kctx, test_cluster_uuid)
	if err == nil {
		fmt.Println("Get Cluster registration exists")
		printK8sClusterRegistration(responseGetReg)
		// Registration exists. delete it so that we can create it
		fmt.Println("Delete Cluster registration")
		responseDelReg, err := nkeClient.ClusterRegistrationOperations.DeleteK8sRegistration(kctx, test_cluster_uuid)
		assert.NoError(t, err)
		printK8sClusterRegistrationDeleteResponse(responseDelReg)
	}

	fmt.Println("Create Cluster registration")
	createRequest := &K8sCreateClusterRegistrationRequest{
		Name:              &test_cluster_name,
		UUID:              test_cluster_uuid,
		CategoriesMapping: test_category_mapping,
		Metadata:          test_metadata,
	}

	// returns type K8sCreateClusterRegistrationResponse
	responseCreateReg, err := nkeClient.ClusterRegistrationOperations.CreateK8sRegistration(kctx, createRequest)
	assert.NoError(t, err)
	printK8sClusterRegistrationCreateResponse(responseCreateReg)

	// TODO get task uuid status
}

func TestKarbonGetK8sRegistrationList(t *testing.T) {
	interceptor := khttpclient.NewInterceptor(http.DefaultTransport)
	creds := testCredsFromEnv(t)
	nkeClient, err := NewKarbonAPIClient(creds, WithRoundTripper(interceptor))
	require.NoError(t, err)

	kctx := mock.NewContext(mock.Config{
		Mode: keploy.MODE_TEST,
		Name: t.Name(),
	})
	fmt.Println("Get Cluster registration List")
	// returns type K8sCreateClusterRegistrationResponse
	response, err := nkeClient.ClusterRegistrationOperations.GetK8sRegistrationList(kctx)
	assert.NoError(t, err)
	printK8sClusterRegistrationList(response)
}
