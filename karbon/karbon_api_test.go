package karbon

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
)

func TestNewKarbonAPIClient(t *testing.T) {
	// verifies positive httpClient creation
	cred := prismgoclient.Credentials{
		URL:                "foo.com",
		Username:           "username",
		Password:           "password",
		Port:               "",
		Endpoint:           "0.0.0.0",
		Insecure:           true,
		FoundationEndpoint: "10.0.0.0",
		FoundationPort:     "8000",
		RequiredFields:     nil,
	}
	_, err := NewKarbonAPIClient(cred)
	assert.NoError(t, err)

	// verify missing httpClient scenario
	cred = prismgoclient.Credentials{
		URL:      "foo.com",
		Insecure: true,
		RequiredFields: map[string][]string{
			"karbon": {"username", "password", "endpoint"},
		},
	}
	_, err = NewKarbonAPIClient(cred)
	assert.Error(t, err)
}

func getCredsForTesting() *prismgoclient.Credentials {
	test_endpoint := "x.x.x.x"
	test_port := "9440"
	test_url := fmt.Sprintf("https://%s:%s", test_endpoint, test_port)
	test_insecure := true
	test_username := "xx"
	test_password := "xx"
	return &prismgoclient.Credentials{
		URL:      test_url,
		Username: test_username,
		Password: test_password,
		Port:     test_port,
		Endpoint: test_endpoint,
		Insecure: test_insecure,
	}
}

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
	cred := getCredsForTesting()
	nkeClient, err := NewKarbonAPIClient(*cred)
	assert.NoError(t, err)

	test_cluster_name := "cluster1"
	test_cluster_uuid := "634F5852-CE1B-465C-A95A-9B8DFFBDDE42"
	test_category_mapping := map[string]string{
		"kubernetes-io-cluster-cluster1": "owned",
		"KubernetesClusterName":          test_cluster_name,
		"KubernetesClusterUUID":          test_cluster_uuid,
	}
	test_metadata_apiversion := "v1.1.0"
	test_metadata := &Metadata{APIVersion: &test_metadata_apiversion}

	fmt.Println("Get Cluster registration")
	responseGetReg, err := nkeClient.Cluster.GetK8sRegistration(context.Background(), test_cluster_uuid)
	if err == nil {
		fmt.Println("Get Cluster registration exists")
		printK8sClusterRegistration(responseGetReg)
		// Registration exists. delete it so that we can create it
		fmt.Println("Delete Cluster registration")
		responseDelReg, err := nkeClient.Cluster.DeleteK8sRegistration(context.Background(), test_cluster_uuid)
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
	responseCreateReg, err := nkeClient.Cluster.CreateK8sRegistration(context.Background(), createRequest)
	assert.NoError(t, err)
	printK8sClusterRegistrationCreateResponse(responseCreateReg)

	// TODO get task uuid status
}

func TestKarbonGetK8sRegistrationList(t *testing.T) {
	cred := getCredsForTesting()
	nkeClient, err := NewKarbonAPIClient(*cred)
	assert.NoError(t, err)

	fmt.Println("Get Cluster registration List")
	// returns type K8sCreateClusterRegistrationResponse
	response, err := nkeClient.Cluster.GetK8sRegistrationList(context.Background())
	assert.NoError(t, err)
	printK8sClusterRegistrationList(response)
}
