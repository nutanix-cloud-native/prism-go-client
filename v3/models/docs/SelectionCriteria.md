# SelectionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpressionList** | [**[]FilterExpression**](FilterExpression.md) |  | 
**EntityType** | **string** | Entity type which has to be selected | 

## Methods

### NewSelectionCriteria

`func NewSelectionCriteria(expressionList []FilterExpression, entityType string, ) *SelectionCriteria`

NewSelectionCriteria instantiates a new SelectionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectionCriteriaWithDefaults

`func NewSelectionCriteriaWithDefaults() *SelectionCriteria`

NewSelectionCriteriaWithDefaults instantiates a new SelectionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressionList

`func (o *SelectionCriteria) GetExpressionList() []FilterExpression`

GetExpressionList returns the ExpressionList field if non-nil, zero value otherwise.

### GetExpressionListOk

`func (o *SelectionCriteria) GetExpressionListOk() (*[]FilterExpression, bool)`

GetExpressionListOk returns a tuple with the ExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionList

`func (o *SelectionCriteria) SetExpressionList(v []FilterExpression)`

SetExpressionList sets ExpressionList field to given value.


### GetEntityType

`func (o *SelectionCriteria) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *SelectionCriteria) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *SelectionCriteria) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


