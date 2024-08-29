# DockerRegistryResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | User name of the docker registry | 
**IsEnabled** | Pointer to **bool** | Flag to indicate whether the docker registry is enabled or not. If unset, defaults to False.  | [optional] 
**Certificate** | Pointer to **string** | Certificate required for the private docker registry | [optional] 
**Url** | **string** | URL of the docker registry | 
**Password** | **string** | Password of the user for the docker registry | 
**IsSystemDefault** | Pointer to **bool** | Flag to indicate whether the docker registry is a system default  | [optional] [readonly] 

## Methods

### NewDockerRegistryResources

`func NewDockerRegistryResources(username string, url string, password string, ) *DockerRegistryResources`

NewDockerRegistryResources instantiates a new DockerRegistryResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryResourcesWithDefaults

`func NewDockerRegistryResourcesWithDefaults() *DockerRegistryResources`

NewDockerRegistryResourcesWithDefaults instantiates a new DockerRegistryResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DockerRegistryResources) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DockerRegistryResources) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DockerRegistryResources) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetIsEnabled

`func (o *DockerRegistryResources) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DockerRegistryResources) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DockerRegistryResources) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DockerRegistryResources) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetCertificate

`func (o *DockerRegistryResources) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DockerRegistryResources) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DockerRegistryResources) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *DockerRegistryResources) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetUrl

`func (o *DockerRegistryResources) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DockerRegistryResources) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DockerRegistryResources) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPassword

`func (o *DockerRegistryResources) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DockerRegistryResources) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DockerRegistryResources) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetIsSystemDefault

`func (o *DockerRegistryResources) GetIsSystemDefault() bool`

GetIsSystemDefault returns the IsSystemDefault field if non-nil, zero value otherwise.

### GetIsSystemDefaultOk

`func (o *DockerRegistryResources) GetIsSystemDefaultOk() (*bool, bool)`

GetIsSystemDefaultOk returns a tuple with the IsSystemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefault

`func (o *DockerRegistryResources) SetIsSystemDefault(v bool)`

SetIsSystemDefault sets IsSystemDefault field to given value.

### HasIsSystemDefault

`func (o *DockerRegistryResources) HasIsSystemDefault() bool`

HasIsSystemDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


