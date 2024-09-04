# PermissionDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the permission entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Permission name. | 
**Resources** | [**PermissionResources**](PermissionResources.md) |  | 
**Description** | Pointer to **string** | Human readable description of the permission. | [optional] 

## Methods

### NewPermissionDefStatus

`func NewPermissionDefStatus(name string, resources PermissionResources, ) *PermissionDefStatus`

NewPermissionDefStatus instantiates a new PermissionDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDefStatusWithDefaults

`func NewPermissionDefStatusWithDefaults() *PermissionDefStatus`

NewPermissionDefStatusWithDefaults instantiates a new PermissionDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PermissionDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PermissionDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PermissionDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PermissionDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *PermissionDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *PermissionDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *PermissionDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *PermissionDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *PermissionDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *PermissionDefStatus) GetResources() PermissionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionDefStatus) GetResourcesOk() (*PermissionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionDefStatus) SetResources(v PermissionResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *PermissionDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


