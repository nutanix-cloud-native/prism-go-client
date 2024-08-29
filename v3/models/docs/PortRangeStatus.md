# PortRangeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPort** | Pointer to **int32** |  | [optional] 
**StartPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewPortRangeStatus

`func NewPortRangeStatus() *PortRangeStatus`

NewPortRangeStatus instantiates a new PortRangeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortRangeStatusWithDefaults

`func NewPortRangeStatusWithDefaults() *PortRangeStatus`

NewPortRangeStatusWithDefaults instantiates a new PortRangeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPort

`func (o *PortRangeStatus) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *PortRangeStatus) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *PortRangeStatus) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *PortRangeStatus) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetStartPort

`func (o *PortRangeStatus) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *PortRangeStatus) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *PortRangeStatus) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *PortRangeStatus) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


