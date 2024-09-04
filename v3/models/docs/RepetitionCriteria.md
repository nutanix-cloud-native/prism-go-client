# RepetitionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepetitionRule** | Pointer to **string** | Rule based on which the widget/section will be repeating. | [optional] 
**EntityType** | **string** | Type of the entity. | 

## Methods

### NewRepetitionCriteria

`func NewRepetitionCriteria(entityType string, ) *RepetitionCriteria`

NewRepetitionCriteria instantiates a new RepetitionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepetitionCriteriaWithDefaults

`func NewRepetitionCriteriaWithDefaults() *RepetitionCriteria`

NewRepetitionCriteriaWithDefaults instantiates a new RepetitionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepetitionRule

`func (o *RepetitionCriteria) GetRepetitionRule() string`

GetRepetitionRule returns the RepetitionRule field if non-nil, zero value otherwise.

### GetRepetitionRuleOk

`func (o *RepetitionCriteria) GetRepetitionRuleOk() (*string, bool)`

GetRepetitionRuleOk returns a tuple with the RepetitionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionRule

`func (o *RepetitionCriteria) SetRepetitionRule(v string)`

SetRepetitionRule sets RepetitionRule field to given value.

### HasRepetitionRule

`func (o *RepetitionCriteria) HasRepetitionRule() bool`

HasRepetitionRule returns a boolean if a field has been set.

### GetEntityType

`func (o *RepetitionCriteria) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RepetitionCriteria) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RepetitionCriteria) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


