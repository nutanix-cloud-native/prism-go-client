# ProtectionRuleDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the protection rule. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the protection rule, if in an error state.  | [optional] 
**Name** | **string** | Protection Rule name | 
**Resources** | [**ProtectionRuleResources**](ProtectionRuleResources.md) |  | 
**Description** | Pointer to **string** | A description for the protection rule. | [optional] 

## Methods

### NewProtectionRuleDefStatus

`func NewProtectionRuleDefStatus(name string, resources ProtectionRuleResources, ) *ProtectionRuleDefStatus`

NewProtectionRuleDefStatus instantiates a new ProtectionRuleDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleDefStatusWithDefaults

`func NewProtectionRuleDefStatusWithDefaults() *ProtectionRuleDefStatus`

NewProtectionRuleDefStatusWithDefaults instantiates a new ProtectionRuleDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ProtectionRuleDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProtectionRuleDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProtectionRuleDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProtectionRuleDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ProtectionRuleDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ProtectionRuleDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ProtectionRuleDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ProtectionRuleDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *ProtectionRuleDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionRuleDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionRuleDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ProtectionRuleDefStatus) GetResources() ProtectionRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProtectionRuleDefStatus) GetResourcesOk() (*ProtectionRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProtectionRuleDefStatus) SetResources(v ProtectionRuleResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ProtectionRuleDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionRuleDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionRuleDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionRuleDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


