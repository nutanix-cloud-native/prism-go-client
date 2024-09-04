# DockerRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the docker registry | 
**Resources** | [**DockerRegistryResources**](DockerRegistryResources.md) |  | 

## Methods

### NewDockerRegistry

`func NewDockerRegistry(name string, resources DockerRegistryResources, ) *DockerRegistry`

NewDockerRegistry instantiates a new DockerRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryWithDefaults

`func NewDockerRegistryWithDefaults() *DockerRegistry`

NewDockerRegistryWithDefaults instantiates a new DockerRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DockerRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerRegistry) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *DockerRegistry) GetResources() DockerRegistryResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DockerRegistry) GetResourcesOk() (*DockerRegistryResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DockerRegistry) SetResources(v DockerRegistryResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


