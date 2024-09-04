# UserGroupDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**UserGroupOutputResource**](UserGroupOutputResource.md) |  | 

## Methods

### NewUserGroupDefStatus

`func NewUserGroupDefStatus(resources UserGroupOutputResource, ) *UserGroupDefStatus`

NewUserGroupDefStatus instantiates a new UserGroupDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupDefStatusWithDefaults

`func NewUserGroupDefStatusWithDefaults() *UserGroupDefStatus`

NewUserGroupDefStatusWithDefaults instantiates a new UserGroupDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *UserGroupDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserGroupDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserGroupDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserGroupDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *UserGroupDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *UserGroupDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *UserGroupDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *UserGroupDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *UserGroupDefStatus) GetResources() UserGroupOutputResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserGroupDefStatus) GetResourcesOk() (*UserGroupOutputResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserGroupDefStatus) SetResources(v UserGroupOutputResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


