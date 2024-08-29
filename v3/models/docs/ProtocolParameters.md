# ProtocolParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udp** | Pointer to [**Udp**](Udp.md) |  | [optional] 
**Icmp** | Pointer to [**Icmp**](Icmp.md) |  | [optional] 
**ProtocolNumber** | Pointer to **int32** |  | [optional] 
**Tcp** | Pointer to [**Tcp**](Tcp.md) |  | [optional] 

## Methods

### NewProtocolParameters

`func NewProtocolParameters() *ProtocolParameters`

NewProtocolParameters instantiates a new ProtocolParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolParametersWithDefaults

`func NewProtocolParametersWithDefaults() *ProtocolParameters`

NewProtocolParametersWithDefaults instantiates a new ProtocolParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdp

`func (o *ProtocolParameters) GetUdp() Udp`

GetUdp returns the Udp field if non-nil, zero value otherwise.

### GetUdpOk

`func (o *ProtocolParameters) GetUdpOk() (*Udp, bool)`

GetUdpOk returns a tuple with the Udp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdp

`func (o *ProtocolParameters) SetUdp(v Udp)`

SetUdp sets Udp field to given value.

### HasUdp

`func (o *ProtocolParameters) HasUdp() bool`

HasUdp returns a boolean if a field has been set.

### GetIcmp

`func (o *ProtocolParameters) GetIcmp() Icmp`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *ProtocolParameters) GetIcmpOk() (*Icmp, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *ProtocolParameters) SetIcmp(v Icmp)`

SetIcmp sets Icmp field to given value.

### HasIcmp

`func (o *ProtocolParameters) HasIcmp() bool`

HasIcmp returns a boolean if a field has been set.

### GetProtocolNumber

`func (o *ProtocolParameters) GetProtocolNumber() int32`

GetProtocolNumber returns the ProtocolNumber field if non-nil, zero value otherwise.

### GetProtocolNumberOk

`func (o *ProtocolParameters) GetProtocolNumberOk() (*int32, bool)`

GetProtocolNumberOk returns a tuple with the ProtocolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolNumber

`func (o *ProtocolParameters) SetProtocolNumber(v int32)`

SetProtocolNumber sets ProtocolNumber field to given value.

### HasProtocolNumber

`func (o *ProtocolParameters) HasProtocolNumber() bool`

HasProtocolNumber returns a boolean if a field has been set.

### GetTcp

`func (o *ProtocolParameters) GetTcp() Tcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ProtocolParameters) GetTcpOk() (*Tcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ProtocolParameters) SetTcp(v Tcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ProtocolParameters) HasTcp() bool`

HasTcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


