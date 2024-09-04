# EntitySyncRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Entity sync rule name. | 
**Resources** | [**EntitySyncRuleResources**](EntitySyncRuleResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the entity sync rule. | [optional] 

## Methods

### NewEntitySyncRule

`func NewEntitySyncRule(name string, resources EntitySyncRuleResources, ) *EntitySyncRule`

NewEntitySyncRule instantiates a new EntitySyncRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySyncRuleWithDefaults

`func NewEntitySyncRuleWithDefaults() *EntitySyncRule`

NewEntitySyncRuleWithDefaults instantiates a new EntitySyncRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntitySyncRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitySyncRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitySyncRule) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *EntitySyncRule) GetResources() EntitySyncRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntitySyncRule) GetResourcesOk() (*EntitySyncRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntitySyncRule) SetResources(v EntitySyncRuleResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *EntitySyncRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitySyncRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitySyncRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitySyncRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


