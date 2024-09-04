# QueryTimevaluePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**QueryValue**](QueryValue.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryTimevaluePair

`func NewQueryTimevaluePair() *QueryTimevaluePair`

NewQueryTimevaluePair instantiates a new QueryTimevaluePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTimevaluePairWithDefaults

`func NewQueryTimevaluePairWithDefaults() *QueryTimevaluePair`

NewQueryTimevaluePairWithDefaults instantiates a new QueryTimevaluePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *QueryTimevaluePair) GetValue() QueryValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryTimevaluePair) GetValueOk() (*QueryValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryTimevaluePair) SetValue(v QueryValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *QueryTimevaluePair) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTime

`func (o *QueryTimevaluePair) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryTimevaluePair) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryTimevaluePair) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryTimevaluePair) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


