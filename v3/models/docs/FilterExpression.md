# FilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Operator to be used for the value on the entity. Example, &#x3D;, cs, &gt;&#x3D;, etc.  | 
**DisplayForValue** | Pointer to **string** | String corresponding to the value to be displayed on UI. | [optional] 
**Value** | **string** | Value of property. | 
**PropertyName** | **string** | Name of the entity property. | 
**DisplayForOperator** | Pointer to **string** | String corresponding to the operator to be displayed on UI. | [optional] 

## Methods

### NewFilterExpression

`func NewFilterExpression(operator string, value string, propertyName string, ) *FilterExpression`

NewFilterExpression instantiates a new FilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterExpressionWithDefaults

`func NewFilterExpressionWithDefaults() *FilterExpression`

NewFilterExpressionWithDefaults instantiates a new FilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *FilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetDisplayForValue

`func (o *FilterExpression) GetDisplayForValue() string`

GetDisplayForValue returns the DisplayForValue field if non-nil, zero value otherwise.

### GetDisplayForValueOk

`func (o *FilterExpression) GetDisplayForValueOk() (*string, bool)`

GetDisplayForValueOk returns a tuple with the DisplayForValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayForValue

`func (o *FilterExpression) SetDisplayForValue(v string)`

SetDisplayForValue sets DisplayForValue field to given value.

### HasDisplayForValue

`func (o *FilterExpression) HasDisplayForValue() bool`

HasDisplayForValue returns a boolean if a field has been set.

### GetValue

`func (o *FilterExpression) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterExpression) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterExpression) SetValue(v string)`

SetValue sets Value field to given value.


### GetPropertyName

`func (o *FilterExpression) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *FilterExpression) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *FilterExpression) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.


### GetDisplayForOperator

`func (o *FilterExpression) GetDisplayForOperator() string`

GetDisplayForOperator returns the DisplayForOperator field if non-nil, zero value otherwise.

### GetDisplayForOperatorOk

`func (o *FilterExpression) GetDisplayForOperatorOk() (*string, bool)`

GetDisplayForOperatorOk returns a tuple with the DisplayForOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayForOperator

`func (o *FilterExpression) SetDisplayForOperator(v string)`

SetDisplayForOperator sets DisplayForOperator field to given value.

### HasDisplayForOperator

`func (o *FilterExpression) HasDisplayForOperator() bool`

HasDisplayForOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


