# ProtectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Protection Rule name | 
**Resources** | [**ProtectionRuleResources**](ProtectionRuleResources.md) |  | 
**Description** | Pointer to **string** | A description for the protection rule. | [optional] 

## Methods

### NewProtectionRule

`func NewProtectionRule(name string, resources ProtectionRuleResources, ) *ProtectionRule`

NewProtectionRule instantiates a new ProtectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleWithDefaults

`func NewProtectionRuleWithDefaults() *ProtectionRule`

NewProtectionRuleWithDefaults instantiates a new ProtectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProtectionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionRule) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ProtectionRule) GetResources() ProtectionRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ProtectionRule) GetResourcesOk() (*ProtectionRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ProtectionRule) SetResources(v ProtectionRuleResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ProtectionRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


