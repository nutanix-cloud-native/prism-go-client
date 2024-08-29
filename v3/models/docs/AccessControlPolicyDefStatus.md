# AccessControlPolicyDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Access Control Policy. | [optional] 
**IsSystemDefined** | Pointer to **bool** | Flag identifying if the ACP is system defined or not. | [optional] 
**State** | Pointer to **string** | The state of the Access Control Policy entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**AccessControlPolicyResources**](AccessControlPolicyResources.md) |  | 
**Description** | Pointer to **string** | The description of the association of a role to a user in a given context. | [optional] 

## Methods

### NewAccessControlPolicyDefStatus

`func NewAccessControlPolicyDefStatus(resources AccessControlPolicyResources, ) *AccessControlPolicyDefStatus`

NewAccessControlPolicyDefStatus instantiates a new AccessControlPolicyDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyDefStatusWithDefaults

`func NewAccessControlPolicyDefStatusWithDefaults() *AccessControlPolicyDefStatus`

NewAccessControlPolicyDefStatusWithDefaults instantiates a new AccessControlPolicyDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessControlPolicyDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessControlPolicyDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessControlPolicyDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessControlPolicyDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsSystemDefined

`func (o *AccessControlPolicyDefStatus) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *AccessControlPolicyDefStatus) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *AccessControlPolicyDefStatus) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.

### HasIsSystemDefined

`func (o *AccessControlPolicyDefStatus) HasIsSystemDefined() bool`

HasIsSystemDefined returns a boolean if a field has been set.

### GetState

`func (o *AccessControlPolicyDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessControlPolicyDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessControlPolicyDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccessControlPolicyDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *AccessControlPolicyDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AccessControlPolicyDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AccessControlPolicyDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AccessControlPolicyDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *AccessControlPolicyDefStatus) GetResources() AccessControlPolicyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AccessControlPolicyDefStatus) GetResourcesOk() (*AccessControlPolicyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AccessControlPolicyDefStatus) SetResources(v AccessControlPolicyResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *AccessControlPolicyDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessControlPolicyDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessControlPolicyDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessControlPolicyDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


