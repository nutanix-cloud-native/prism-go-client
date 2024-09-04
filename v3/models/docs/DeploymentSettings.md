# DeploymentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemId** | **string** | Unique identifier for the deployment assigned by MCM. | 
**MyNtnxToken** | [**MyNtnxToken**](MyNtnxToken.md) |  | 

## Methods

### NewDeploymentSettings

`func NewDeploymentSettings(systemId string, myNtnxToken MyNtnxToken, ) *DeploymentSettings`

NewDeploymentSettings instantiates a new DeploymentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSettingsWithDefaults

`func NewDeploymentSettingsWithDefaults() *DeploymentSettings`

NewDeploymentSettingsWithDefaults instantiates a new DeploymentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemId

`func (o *DeploymentSettings) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *DeploymentSettings) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *DeploymentSettings) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.


### GetMyNtnxToken

`func (o *DeploymentSettings) GetMyNtnxToken() MyNtnxToken`

GetMyNtnxToken returns the MyNtnxToken field if non-nil, zero value otherwise.

### GetMyNtnxTokenOk

`func (o *DeploymentSettings) GetMyNtnxTokenOk() (*MyNtnxToken, bool)`

GetMyNtnxTokenOk returns a tuple with the MyNtnxToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyNtnxToken

`func (o *DeploymentSettings) SetMyNtnxToken(v MyNtnxToken)`

SetMyNtnxToken sets MyNtnxToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


