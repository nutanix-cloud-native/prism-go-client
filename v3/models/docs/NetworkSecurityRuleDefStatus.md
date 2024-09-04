# NetworkSecurityRuleDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**NetworkSecurityRuleResourcesStatus**](NetworkSecurityRuleResourcesStatus.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkSecurityRuleDefStatus

`func NewNetworkSecurityRuleDefStatus() *NetworkSecurityRuleDefStatus`

NewNetworkSecurityRuleDefStatus instantiates a new NetworkSecurityRuleDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleDefStatusWithDefaults

`func NewNetworkSecurityRuleDefStatusWithDefaults() *NetworkSecurityRuleDefStatus`

NewNetworkSecurityRuleDefStatusWithDefaults instantiates a new NetworkSecurityRuleDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *NetworkSecurityRuleDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetworkSecurityRuleDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetworkSecurityRuleDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetworkSecurityRuleDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *NetworkSecurityRuleDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *NetworkSecurityRuleDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *NetworkSecurityRuleDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *NetworkSecurityRuleDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *NetworkSecurityRuleDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSecurityRuleDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSecurityRuleDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkSecurityRuleDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *NetworkSecurityRuleDefStatus) GetResources() NetworkSecurityRuleResourcesStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkSecurityRuleDefStatus) GetResourcesOk() (*NetworkSecurityRuleResourcesStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkSecurityRuleDefStatus) SetResources(v NetworkSecurityRuleResourcesStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *NetworkSecurityRuleDefStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkSecurityRuleDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkSecurityRuleDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkSecurityRuleDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkSecurityRuleDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


