# McmConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentSettings** | [**DeploymentSettings**](DeploymentSettings.md) |  | 
**ProductName** | **string** | Type of the deployed component. For example Prism Central or Flow Gateway. | 
**ComponentRegistryBaseUrl** | **string** | API Keys app base URL. | 
**OrchestratorBaseUrl** | **string** | MCM base URL. | 

## Methods

### NewMcmConfig

`func NewMcmConfig(deploymentSettings DeploymentSettings, productName string, componentRegistryBaseUrl string, orchestratorBaseUrl string, ) *McmConfig`

NewMcmConfig instantiates a new McmConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmConfigWithDefaults

`func NewMcmConfigWithDefaults() *McmConfig`

NewMcmConfigWithDefaults instantiates a new McmConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentSettings

`func (o *McmConfig) GetDeploymentSettings() DeploymentSettings`

GetDeploymentSettings returns the DeploymentSettings field if non-nil, zero value otherwise.

### GetDeploymentSettingsOk

`func (o *McmConfig) GetDeploymentSettingsOk() (*DeploymentSettings, bool)`

GetDeploymentSettingsOk returns a tuple with the DeploymentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSettings

`func (o *McmConfig) SetDeploymentSettings(v DeploymentSettings)`

SetDeploymentSettings sets DeploymentSettings field to given value.


### GetProductName

`func (o *McmConfig) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *McmConfig) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *McmConfig) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetComponentRegistryBaseUrl

`func (o *McmConfig) GetComponentRegistryBaseUrl() string`

GetComponentRegistryBaseUrl returns the ComponentRegistryBaseUrl field if non-nil, zero value otherwise.

### GetComponentRegistryBaseUrlOk

`func (o *McmConfig) GetComponentRegistryBaseUrlOk() (*string, bool)`

GetComponentRegistryBaseUrlOk returns a tuple with the ComponentRegistryBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentRegistryBaseUrl

`func (o *McmConfig) SetComponentRegistryBaseUrl(v string)`

SetComponentRegistryBaseUrl sets ComponentRegistryBaseUrl field to given value.


### GetOrchestratorBaseUrl

`func (o *McmConfig) GetOrchestratorBaseUrl() string`

GetOrchestratorBaseUrl returns the OrchestratorBaseUrl field if non-nil, zero value otherwise.

### GetOrchestratorBaseUrlOk

`func (o *McmConfig) GetOrchestratorBaseUrlOk() (*string, bool)`

GetOrchestratorBaseUrlOk returns a tuple with the OrchestratorBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorBaseUrl

`func (o *McmConfig) SetOrchestratorBaseUrl(v string)`

SetOrchestratorBaseUrl sets OrchestratorBaseUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


