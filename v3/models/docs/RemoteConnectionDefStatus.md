# RemoteConnectionDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the remote connection entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Remote connection name | 
**Resources** | [**RemoteConnectionResources**](RemoteConnectionResources.md) |  | 
**Description** | **string** | Remote connection description | 

## Methods

### NewRemoteConnectionDefStatus

`func NewRemoteConnectionDefStatus(name string, resources RemoteConnectionResources, description string, ) *RemoteConnectionDefStatus`

NewRemoteConnectionDefStatus instantiates a new RemoteConnectionDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionDefStatusWithDefaults

`func NewRemoteConnectionDefStatusWithDefaults() *RemoteConnectionDefStatus`

NewRemoteConnectionDefStatusWithDefaults instantiates a new RemoteConnectionDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RemoteConnectionDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RemoteConnectionDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RemoteConnectionDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RemoteConnectionDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *RemoteConnectionDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *RemoteConnectionDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *RemoteConnectionDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *RemoteConnectionDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *RemoteConnectionDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteConnectionDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteConnectionDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *RemoteConnectionDefStatus) GetResources() RemoteConnectionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RemoteConnectionDefStatus) GetResourcesOk() (*RemoteConnectionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RemoteConnectionDefStatus) SetResources(v RemoteConnectionResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *RemoteConnectionDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteConnectionDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteConnectionDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


