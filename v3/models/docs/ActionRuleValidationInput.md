# ActionRuleValidationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionRuleReference** | Pointer to [**ActionRuleReference**](ActionRuleReference.md) |  | [optional] 
**Resources** | [**ActionRuleResources**](ActionRuleResources.md) |  | 

## Methods

### NewActionRuleValidationInput

`func NewActionRuleValidationInput(resources ActionRuleResources, ) *ActionRuleValidationInput`

NewActionRuleValidationInput instantiates a new ActionRuleValidationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleValidationInputWithDefaults

`func NewActionRuleValidationInputWithDefaults() *ActionRuleValidationInput`

NewActionRuleValidationInputWithDefaults instantiates a new ActionRuleValidationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionRuleReference

`func (o *ActionRuleValidationInput) GetActionRuleReference() ActionRuleReference`

GetActionRuleReference returns the ActionRuleReference field if non-nil, zero value otherwise.

### GetActionRuleReferenceOk

`func (o *ActionRuleValidationInput) GetActionRuleReferenceOk() (*ActionRuleReference, bool)`

GetActionRuleReferenceOk returns a tuple with the ActionRuleReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRuleReference

`func (o *ActionRuleValidationInput) SetActionRuleReference(v ActionRuleReference)`

SetActionRuleReference sets ActionRuleReference field to given value.

### HasActionRuleReference

`func (o *ActionRuleValidationInput) HasActionRuleReference() bool`

HasActionRuleReference returns a boolean if a field has been set.

### GetResources

`func (o *ActionRuleValidationInput) GetResources() ActionRuleResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ActionRuleValidationInput) GetResourcesOk() (*ActionRuleResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ActionRuleValidationInput) SetResources(v ActionRuleResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


