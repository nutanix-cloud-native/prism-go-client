# Expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | If the term is a filter (LHS operator RHS), then this represents the operator.  | [optional] 
**PropertyType** | Pointer to **string** | Whether the term is an attribute, metric, action. | [optional] 
**Value** | Pointer to **string** | In case of a filter, this represents the RHS. | [optional] 
**PropertyName** | Pointer to **string** | Name of the attribute, metric, action. | [optional] 
**EntityType** | Pointer to **string** | The entity type that the term represents. | [optional] 

## Methods

### NewExpression

`func NewExpression() *Expression`

NewExpression instantiates a new Expression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionWithDefaults

`func NewExpressionWithDefaults() *Expression`

NewExpressionWithDefaults instantiates a new Expression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *Expression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Expression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Expression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Expression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPropertyType

`func (o *Expression) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *Expression) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *Expression) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *Expression) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### GetValue

`func (o *Expression) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Expression) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Expression) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Expression) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPropertyName

`func (o *Expression) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *Expression) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *Expression) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *Expression) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetEntityType

`func (o *Expression) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Expression) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Expression) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Expression) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


