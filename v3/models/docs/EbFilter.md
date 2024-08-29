# EbFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rhs** | Pointer to **string** | Right hand side of the filter expression. | [optional] 
**AliasForLhs** | Pointer to **string** | Alias for LHS of the filter expression. Used for Backend to tell UI if the filter name UI use to do group does not equal to actual one showing search bar and filter panel.  | [optional] 
**DisplayForOperator** | Pointer to **string** | Display string for the operator. | [optional] 
**Lhs** | Pointer to **string** | Left hand side of the filter expression. | [optional] 
**DisplayForRhs** | Pointer to **string** | Display for RHS value of the filter expression. | [optional] 
**FilterName** | Pointer to **string** | Filter Display name. | [optional] 
**Operator** | Pointer to **string** | Operator that is being used in filter. | [optional] 

## Methods

### NewEbFilter

`func NewEbFilter() *EbFilter`

NewEbFilter instantiates a new EbFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEbFilterWithDefaults

`func NewEbFilterWithDefaults() *EbFilter`

NewEbFilterWithDefaults instantiates a new EbFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRhs

`func (o *EbFilter) GetRhs() string`

GetRhs returns the Rhs field if non-nil, zero value otherwise.

### GetRhsOk

`func (o *EbFilter) GetRhsOk() (*string, bool)`

GetRhsOk returns a tuple with the Rhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhs

`func (o *EbFilter) SetRhs(v string)`

SetRhs sets Rhs field to given value.

### HasRhs

`func (o *EbFilter) HasRhs() bool`

HasRhs returns a boolean if a field has been set.

### GetAliasForLhs

`func (o *EbFilter) GetAliasForLhs() string`

GetAliasForLhs returns the AliasForLhs field if non-nil, zero value otherwise.

### GetAliasForLhsOk

`func (o *EbFilter) GetAliasForLhsOk() (*string, bool)`

GetAliasForLhsOk returns a tuple with the AliasForLhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasForLhs

`func (o *EbFilter) SetAliasForLhs(v string)`

SetAliasForLhs sets AliasForLhs field to given value.

### HasAliasForLhs

`func (o *EbFilter) HasAliasForLhs() bool`

HasAliasForLhs returns a boolean if a field has been set.

### GetDisplayForOperator

`func (o *EbFilter) GetDisplayForOperator() string`

GetDisplayForOperator returns the DisplayForOperator field if non-nil, zero value otherwise.

### GetDisplayForOperatorOk

`func (o *EbFilter) GetDisplayForOperatorOk() (*string, bool)`

GetDisplayForOperatorOk returns a tuple with the DisplayForOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayForOperator

`func (o *EbFilter) SetDisplayForOperator(v string)`

SetDisplayForOperator sets DisplayForOperator field to given value.

### HasDisplayForOperator

`func (o *EbFilter) HasDisplayForOperator() bool`

HasDisplayForOperator returns a boolean if a field has been set.

### GetLhs

`func (o *EbFilter) GetLhs() string`

GetLhs returns the Lhs field if non-nil, zero value otherwise.

### GetLhsOk

`func (o *EbFilter) GetLhsOk() (*string, bool)`

GetLhsOk returns a tuple with the Lhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhs

`func (o *EbFilter) SetLhs(v string)`

SetLhs sets Lhs field to given value.

### HasLhs

`func (o *EbFilter) HasLhs() bool`

HasLhs returns a boolean if a field has been set.

### GetDisplayForRhs

`func (o *EbFilter) GetDisplayForRhs() string`

GetDisplayForRhs returns the DisplayForRhs field if non-nil, zero value otherwise.

### GetDisplayForRhsOk

`func (o *EbFilter) GetDisplayForRhsOk() (*string, bool)`

GetDisplayForRhsOk returns a tuple with the DisplayForRhs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayForRhs

`func (o *EbFilter) SetDisplayForRhs(v string)`

SetDisplayForRhs sets DisplayForRhs field to given value.

### HasDisplayForRhs

`func (o *EbFilter) HasDisplayForRhs() bool`

HasDisplayForRhs returns a boolean if a field has been set.

### GetFilterName

`func (o *EbFilter) GetFilterName() string`

GetFilterName returns the FilterName field if non-nil, zero value otherwise.

### GetFilterNameOk

`func (o *EbFilter) GetFilterNameOk() (*string, bool)`

GetFilterNameOk returns a tuple with the FilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterName

`func (o *EbFilter) SetFilterName(v string)`

SetFilterName sets FilterName field to given value.

### HasFilterName

`func (o *EbFilter) HasFilterName() bool`

HasFilterName returns a boolean if a field has been set.

### GetOperator

`func (o *EbFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EbFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EbFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *EbFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


