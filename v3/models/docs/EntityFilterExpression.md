# EntityFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | The operator in the filter expression. | 
**RightHandSide** | [**RightHandSide**](RightHandSide.md) |  | 
**LeftHandSide** | [**EntityFilterExpressionLeftHandSide**](EntityFilterExpressionLeftHandSide.md) |  | 

## Methods

### NewEntityFilterExpression

`func NewEntityFilterExpression(operator string, rightHandSide RightHandSide, leftHandSide EntityFilterExpressionLeftHandSide, ) *EntityFilterExpression`

NewEntityFilterExpression instantiates a new EntityFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityFilterExpressionWithDefaults

`func NewEntityFilterExpressionWithDefaults() *EntityFilterExpression`

NewEntityFilterExpressionWithDefaults instantiates a new EntityFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *EntityFilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EntityFilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EntityFilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetRightHandSide

`func (o *EntityFilterExpression) GetRightHandSide() RightHandSide`

GetRightHandSide returns the RightHandSide field if non-nil, zero value otherwise.

### GetRightHandSideOk

`func (o *EntityFilterExpression) GetRightHandSideOk() (*RightHandSide, bool)`

GetRightHandSideOk returns a tuple with the RightHandSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightHandSide

`func (o *EntityFilterExpression) SetRightHandSide(v RightHandSide)`

SetRightHandSide sets RightHandSide field to given value.


### GetLeftHandSide

`func (o *EntityFilterExpression) GetLeftHandSide() EntityFilterExpressionLeftHandSide`

GetLeftHandSide returns the LeftHandSide field if non-nil, zero value otherwise.

### GetLeftHandSideOk

`func (o *EntityFilterExpression) GetLeftHandSideOk() (*EntityFilterExpressionLeftHandSide, bool)`

GetLeftHandSideOk returns a tuple with the LeftHandSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftHandSide

`func (o *EntityFilterExpression) SetLeftHandSide(v EntityFilterExpressionLeftHandSide)`

SetLeftHandSide sets LeftHandSide field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


