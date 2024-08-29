# NetworkSecurityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Resources** | [**NetworkSecurityRuleResources**](NetworkSecurityRuleResources.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkSecurityRule

`func NewNetworkSecurityRule(name string, resources NetworkSecurityRuleResources, ) *NetworkSecurityRule`

NewNetworkSecurityRule instantiates a new NetworkSecurityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleWithDefaults

`func NewNetworkSecurityRuleWithDefaults() *NetworkSecurityRule`

NewNetworkSecurityRuleWithDefaults instantiates a new NetworkSecurityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkSecurityRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkSecurityRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkSecurityRule) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *NetworkSecurityRule) GetResources() NetworkSecurityRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkSecurityRule) GetResourcesOk() (*NetworkSecurityRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkSecurityRule) SetResources(v NetworkSecurityRuleResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *NetworkSecurityRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkSecurityRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkSecurityRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkSecurityRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


