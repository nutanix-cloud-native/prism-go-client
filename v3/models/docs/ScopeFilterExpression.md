# ScopeFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | The operator of the filter expression. | 
**RightHandSide** | [**ScopeRightHandSide**](ScopeRightHandSide.md) |  | 
**LeftHandSide** | **string** | The LHS of the filter expression - the scope type. Following is the list of supported values - * CATEGORY * PROJECT * CLUSTER * VPC * CONNECTIVITY_TYPE  | 

## Methods

### NewScopeFilterExpression

`func NewScopeFilterExpression(operator string, rightHandSide ScopeRightHandSide, leftHandSide string, ) *ScopeFilterExpression`

NewScopeFilterExpression instantiates a new ScopeFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeFilterExpressionWithDefaults

`func NewScopeFilterExpressionWithDefaults() *ScopeFilterExpression`

NewScopeFilterExpressionWithDefaults instantiates a new ScopeFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ScopeFilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ScopeFilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ScopeFilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetRightHandSide

`func (o *ScopeFilterExpression) GetRightHandSide() ScopeRightHandSide`

GetRightHandSide returns the RightHandSide field if non-nil, zero value otherwise.

### GetRightHandSideOk

`func (o *ScopeFilterExpression) GetRightHandSideOk() (*ScopeRightHandSide, bool)`

GetRightHandSideOk returns a tuple with the RightHandSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightHandSide

`func (o *ScopeFilterExpression) SetRightHandSide(v ScopeRightHandSide)`

SetRightHandSide sets RightHandSide field to given value.


### GetLeftHandSide

`func (o *ScopeFilterExpression) GetLeftHandSide() string`

GetLeftHandSide returns the LeftHandSide field if non-nil, zero value otherwise.

### GetLeftHandSideOk

`func (o *ScopeFilterExpression) GetLeftHandSideOk() (*string, bool)`

GetLeftHandSideOk returns a tuple with the LeftHandSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftHandSide

`func (o *ScopeFilterExpression) SetLeftHandSide(v string)`

SetLeftHandSide sets LeftHandSide field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


