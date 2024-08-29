# EntityFilterExpressionLeftHandSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** | The representation of the possible value of an LHS, in this case it is an entity type.A certain entity type e.g. VM Note. To express all entity types, use ALL  | [optional] 

## Methods

### NewEntityFilterExpressionLeftHandSide

`func NewEntityFilterExpressionLeftHandSide() *EntityFilterExpressionLeftHandSide`

NewEntityFilterExpressionLeftHandSide instantiates a new EntityFilterExpressionLeftHandSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityFilterExpressionLeftHandSideWithDefaults

`func NewEntityFilterExpressionLeftHandSideWithDefaults() *EntityFilterExpressionLeftHandSide`

NewEntityFilterExpressionLeftHandSideWithDefaults instantiates a new EntityFilterExpressionLeftHandSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *EntityFilterExpressionLeftHandSide) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityFilterExpressionLeftHandSide) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityFilterExpressionLeftHandSide) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EntityFilterExpressionLeftHandSide) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


