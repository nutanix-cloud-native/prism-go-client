# RoleDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Role name. | 
**IsSystemDefined** | Pointer to **bool** | Flag identifying if the role is system defined or not. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the role entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**RoleResources1**](RoleResources1.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the role. | [optional] 

## Methods

### NewRoleDefStatus

`func NewRoleDefStatus(name string, resources RoleResources1, ) *RoleDefStatus`

NewRoleDefStatus instantiates a new RoleDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDefStatusWithDefaults

`func NewRoleDefStatusWithDefaults() *RoleDefStatus`

NewRoleDefStatusWithDefaults instantiates a new RoleDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetIsSystemDefined

`func (o *RoleDefStatus) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *RoleDefStatus) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *RoleDefStatus) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *RoleDefStatus) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### GetState

`func (o *RoleDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RoleDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RoleDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RoleDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *RoleDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *RoleDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *RoleDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *RoleDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *RoleDefStatus) GetResources() RoleResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RoleDefStatus) GetResourcesOk() (*RoleResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RoleDefStatus) SetResources(v RoleResources1)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *RoleDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


