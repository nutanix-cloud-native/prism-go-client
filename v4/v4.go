package v4

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	clustermgmtapi "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/api"
	clustermgmtclient "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	networkingApi "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/api"
	networkingClient "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	prismApi "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/api"
	prismClient "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	vmApi "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/api"
	vmClient "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	volumesApi "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/api"
	volumesClient "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/client"

	prismgoclient "github.com/nutanix-cloud-native/prism-go-client"
)

const (
	defaultEndpointPort = "9440"
	authorizationHeader = "Authorization"
)

// Client manages the V4 API
type Client struct {
	CategoriesApiInstance   *prismApi.CategoriesApi
	ClustersApiInstance     *clustermgmtapi.ClustersApi
	ImagesApiInstance       *vmApi.ImagesApi
	SubnetsApiInstance      *networkingApi.SubnetsApi
	SubnetIPReservationApi  *networkingApi.SubnetIPReservationApi
	StorageContainerAPI     *clustermgmtapi.StorageContainersApi
	TasksApiInstance        *prismApi.TasksApi
	VolumeGroupsApiInstance *volumesApi.VolumeGroupsApi
	VmApiInstance           *vmApi.VmApi
}

type endpointInfo struct {
	host string
	port int
}

// ClientOption is a functional option for the Client
type ClientOption func(*Client) error

// NewV4Client return an internal to operate V4 resources
func NewV4Client(credentials prismgoclient.Credentials, opts ...ClientOption) (*Client, error) {
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

	if err := initPrismApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Tasks API instance: %v", err)
	}

	if err := initStorageApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Storage API instance: %v", err)
	}

	if err := initVolumesApiInstance(v4Client, credentials); err != nil {
		return nil, fmt.Errorf("failed to create Volumes API instance: %v", err)
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
	apiClientInstance := clustermgmtclient.NewApiClient()
	apiClientInstance.VerifySSL = !credentials.Insecure
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.ClustersApiInstance = clustermgmtapi.NewClustersApi(apiClientInstance)
	return nil
}

func initPrismApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
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
	v4Client.TasksApiInstance = prismApi.NewTasksApi(apiClientInstance)
	v4Client.CategoriesApiInstance = prismApi.NewCategoriesApi(apiClientInstance)
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
	v4Client.SubnetsApiInstance = networkingApi.NewSubnetsApi(apiClientInstance)
	v4Client.SubnetIPReservationApi = networkingApi.NewSubnetIPReservationApi(apiClientInstance)
	return nil
}

func initStorageApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := clustermgmtclient.NewApiClient()
	apiClientInstance.SetVerifySSL(!credentials.Insecure)
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.StorageContainerAPI = clustermgmtapi.NewStorageContainersApi(apiClientInstance)
	return nil
}

func initVolumesApiInstance(v4Client *Client, credentials prismgoclient.Credentials) error {
	ep, err := getEndpointInfo(credentials)
	if err != nil {
		return err
	}
	apiClientInstance := volumesClient.NewApiClient()
	apiClientInstance.SetVerifySSL(!credentials.Insecure)
	apiClientInstance.Host = ep.host
	apiClientInstance.Port = ep.port
	apiClientInstance.AddDefaultHeader(
		authorizationHeader, fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))
	v4Client.VolumeGroupsApiInstance = volumesApi.NewVolumeGroupsApi(apiClientInstance)
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
