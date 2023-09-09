package v4

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
	clusterApi "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/api"
	clusterClient "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	networkingApi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/api"
	networkingClient "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	prismApi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/api"
	prismClient "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	storageApi "github.com/nutanix/ntnx-api-golang-clients/storage-go-client/v4/api"
	storageClient "github.com/nutanix/ntnx-api-golang-clients/storage-go-client/v4/client"
	vmApi "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/api"
	vmClient "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
)

const (
	defaultEndpointPort = "9440"
	authorizationHeader = "Authorization"
)

// Client manages the V4 API
type Client struct {
	VmApiInstance                       *vmApi.VmApi
	ImagesApiInstance                   *vmApi.ImagesApi
	SubnetApiInstance                   *networkingApi.SubnetApi
	SubnetReserveUnreserveIPAPIInstance *networkingApi.SubnetReserveUnreserveIpApi
	ClusterApiInstance                  *clusterApi.ClusterApi
	TasksApiInstance                    *prismApi.TaskApi
	StorageContainerAPI                 *storageApi.StorageContainerApi
}

type endpointInfo struct {
	host string
	port int
}

// NewV4Client return a internal to operate V4 resources
func NewV4Client(credentials prismgoclient.Credentials) (*Client, error) {
	if credentials.Username == "" || credentials.Password == "" || credentials.Endpoint == "" {
		return nil, fmt.Errorf("username, password and endpoint are required")
	}

	v4Client := &Client{}

	if err := initVmApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create VM API instance: %v", err)
	}

	if err := initSubnetApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Subnet API instance: %v", err)
	}

	if err := initClusterApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Cluster API instance: %v", err)
	}

	if err := initTasksApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Tasks API instance: %v", err)
	}

	if err := initStorageApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Storage API instance: %v", err)
	}

	return v4Client, nil
}

func initVmApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := vmClient.NewApiClient()
	apiClientInstance.VerifySSL = !credentials.Insecure
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.VmApiInstance = vmApi.NewVmApi(apiClientInstance)
	v4Client.ImagesApiInstance = vmApi.NewImagesApi(apiClientInstance)
	return nil
}

func initClusterApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := clusterClient.NewApiClient()
	apiClientInstance.VerifySSL = !credentials.Insecure
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.ClusterApiInstance = clusterApi.NewClusterApi(apiClientInstance)
	return nil
}

func initTasksApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := prismClient.NewApiClient()
	apiClientInstance.VerifySSL = !credentials.Insecure
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.TasksApiInstance = prismApi.NewTaskApi(apiClientInstance)
	return nil
}

func initSubnetApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := networkingClient.NewApiClient()
	apiClientInstance.SetVerifySSL(!credentials.Insecure)
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.SubnetApiInstance = networkingApi.NewSubnetApi(apiClientInstance)
	v4Client.SubnetReserveUnreserveIPAPIInstance = networkingApi.NewSubnetReserveUnreserveIpApi(apiClientInstance)
	return nil
}

func initStorageApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := storageClient.NewApiClient()
	apiClientInstance.SetVerifySSL(!credentials.Insecure)
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.StorageContainerAPI = storageApi.NewStorageContainerApi(apiClientInstance)
	return nil
}

func getEndpointInfo(credentials prismgoclient.Credentials) (*endpointInfo, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s", credentials.Endpoint))
	if err != nil {
		return nil, err
	}
	h := u.Hostname()
	port := u.Port()
	if port == "" {
		port = defaultEndpointPort
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to convert port number to int: %v", err)
	}
	return &endpointInfo{
		host: h,
		port: p,
	}, nil
}
