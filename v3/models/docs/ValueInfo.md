# ValueInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | A valid regex expression. | [optional] 
**Minimum** | Pointer to [**ValueInfoMinimum**](ValueInfoMinimum.md) |  | [optional] 
**Maximum** | Pointer to [**ValueInfoMaximum**](ValueInfoMaximum.md) |  | [optional] 
**Unit** | Pointer to **string** | The unit of measurement. | [optional] 

## Methods

### NewValueInfo

`func NewValueInfo() *ValueInfo`

NewValueInfo instantiates a new ValueInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueInfoWithDefaults

`func NewValueInfoWithDefaults() *ValueInfo`

NewValueInfoWithDefaults instantiates a new ValueInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *ValueInfo) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ValueInfo) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ValueInfo) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ValueInfo) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetMinimum

`func (o *ValueInfo) GetMinimum() ValueInfoMinimum`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *ValueInfo) GetMinimumOk() (*ValueInfoMinimum, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *ValueInfo) SetMinimum(v ValueInfoMinimum)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *ValueInfo) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMaximum

`func (o *ValueInfo) GetMaximum() ValueInfoMaximum`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *ValueInfo) GetMaximumOk() (*ValueInfoMaximum, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *ValueInfo) SetMaximum(v ValueInfoMaximum)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *ValueInfo) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetUnit

`func (o *ValueInfo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ValueInfo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ValueInfo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ValueInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


