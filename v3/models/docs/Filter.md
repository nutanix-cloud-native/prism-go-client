# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityFilterExpressionList** | [**[]EntityFilterExpression**](EntityFilterExpression.md) | A list of Entity filter expressions. | 
**ScopeFilterExpressionList** | Pointer to [**[]ScopeFilterExpression**](ScopeFilterExpression.md) | A list of Scope filter expressions. | [optional] 

## Methods

### NewFilter

`func NewFilter(entityFilterExpressionList []EntityFilterExpression, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityFilterExpressionList

`func (o *Filter) GetEntityFilterExpressionList() []EntityFilterExpression`

GetEntityFilterExpressionList returns the EntityFilterExpressionList field if non-nil, zero value otherwise.

### GetEntityFilterExpressionListOk

`func (o *Filter) GetEntityFilterExpressionListOk() (*[]EntityFilterExpression, bool)`

GetEntityFilterExpressionListOk returns a tuple with the EntityFilterExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityFilterExpressionList

`func (o *Filter) SetEntityFilterExpressionList(v []EntityFilterExpression)`

SetEntityFilterExpressionList sets EntityFilterExpressionList field to given value.


### GetScopeFilterExpressionList

`func (o *Filter) GetScopeFilterExpressionList() []ScopeFilterExpression`

GetScopeFilterExpressionList returns the ScopeFilterExpressionList field if non-nil, zero value otherwise.

### GetScopeFilterExpressionListOk

`func (o *Filter) GetScopeFilterExpressionListOk() (*[]ScopeFilterExpression, bool)`

GetScopeFilterExpressionListOk returns a tuple with the ScopeFilterExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeFilterExpressionList

`func (o *Filter) SetScopeFilterExpressionList(v []ScopeFilterExpression)`

SetScopeFilterExpressionList sets ScopeFilterExpressionList field to given value.

### HasScopeFilterExpressionList

`func (o *Filter) HasScopeFilterExpressionList() bool`

HasScopeFilterExpressionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


