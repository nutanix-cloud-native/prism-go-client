# DockerRegistryDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Name of the docker registry | 
**Resources** | [**DockerRegistryResources**](DockerRegistryResources.md) |  | 

## Methods

### NewDockerRegistryDefStatus

`func NewDockerRegistryDefStatus(name string, resources DockerRegistryResources, ) *DockerRegistryDefStatus`

NewDockerRegistryDefStatus instantiates a new DockerRegistryDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryDefStatusWithDefaults

`func NewDockerRegistryDefStatusWithDefaults() *DockerRegistryDefStatus`

NewDockerRegistryDefStatusWithDefaults instantiates a new DockerRegistryDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DockerRegistryDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DockerRegistryDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DockerRegistryDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DockerRegistryDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *DockerRegistryDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *DockerRegistryDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *DockerRegistryDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *DockerRegistryDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *DockerRegistryDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DockerRegistryDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DockerRegistryDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *DockerRegistryDefStatus) GetResources() DockerRegistryResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DockerRegistryDefStatus) GetResourcesOk() (*DockerRegistryResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DockerRegistryDefStatus) SetResources(v DockerRegistryResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


