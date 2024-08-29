# EntitySyncRuleDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Entity sync rule name. | 
**Resources** | [**EntitySyncRuleResources**](EntitySyncRuleResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the entity sync rule. | [optional] 

## Methods

### NewEntitySyncRuleDefStatus

`func NewEntitySyncRuleDefStatus(name string, resources EntitySyncRuleResources, ) *EntitySyncRuleDefStatus`

NewEntitySyncRuleDefStatus instantiates a new EntitySyncRuleDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySyncRuleDefStatusWithDefaults

`func NewEntitySyncRuleDefStatusWithDefaults() *EntitySyncRuleDefStatus`

NewEntitySyncRuleDefStatusWithDefaults instantiates a new EntitySyncRuleDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntitySyncRuleDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitySyncRuleDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitySyncRuleDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *EntitySyncRuleDefStatus) GetResources() EntitySyncRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntitySyncRuleDefStatus) GetResourcesOk() (*EntitySyncRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntitySyncRuleDefStatus) SetResources(v EntitySyncRuleResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *EntitySyncRuleDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitySyncRuleDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitySyncRuleDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitySyncRuleDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


