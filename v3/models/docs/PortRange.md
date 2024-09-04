# PortRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPort** | Pointer to **int32** |  | [optional] 
**StartPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewPortRange

`func NewPortRange() *PortRange`

NewPortRange instantiates a new PortRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortRangeWithDefaults

`func NewPortRangeWithDefaults() *PortRange`

NewPortRangeWithDefaults instantiates a new PortRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPort

`func (o *PortRange) GetEndPort() int32`

GetEndPort returns the EndPort field if non-nil, zero value otherwise.

### GetEndPortOk

`func (o *PortRange) GetEndPortOk() (*int32, bool)`

GetEndPortOk returns a tuple with the EndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPort

`func (o *PortRange) SetEndPort(v int32)`

SetEndPort sets EndPort field to given value.

### HasEndPort

`func (o *PortRange) HasEndPort() bool`

HasEndPort returns a boolean if a field has been set.

### GetStartPort

`func (o *PortRange) GetStartPort() int32`

GetStartPort returns the StartPort field if non-nil, zero value otherwise.

### GetStartPortOk

`func (o *PortRange) GetStartPortOk() (*int32, bool)`

GetStartPortOk returns a tuple with the StartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPort

`func (o *PortRange) SetStartPort(v int32)`

SetStartPort sets StartPort field to given value.

### HasStartPort

`func (o *PortRange) HasStartPort() bool`

HasStartPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


