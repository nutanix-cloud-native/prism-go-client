# ProtocolParametersStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Udp** | Pointer to [**UdpStatus**](UdpStatus.md) |  | [optional] 
**Icmp** | Pointer to [**IcmpStatus**](IcmpStatus.md) |  | [optional] 
**ProtocolNumber** | Pointer to **int32** |  | [optional] 
**Tcp** | Pointer to [**TcpStatus**](TcpStatus.md) |  | [optional] 

## Methods

### NewProtocolParametersStatus

`func NewProtocolParametersStatus() *ProtocolParametersStatus`

NewProtocolParametersStatus instantiates a new ProtocolParametersStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolParametersStatusWithDefaults

`func NewProtocolParametersStatusWithDefaults() *ProtocolParametersStatus`

NewProtocolParametersStatusWithDefaults instantiates a new ProtocolParametersStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUdp

`func (o *ProtocolParametersStatus) GetUdp() UdpStatus`

GetUdp returns the Udp field if non-nil, zero value otherwise.

### GetUdpOk

`func (o *ProtocolParametersStatus) GetUdpOk() (*UdpStatus, bool)`

GetUdpOk returns a tuple with the Udp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdp

`func (o *ProtocolParametersStatus) SetUdp(v UdpStatus)`

SetUdp sets Udp field to given value.

### HasUdp

`func (o *ProtocolParametersStatus) HasUdp() bool`

HasUdp returns a boolean if a field has been set.

### GetIcmp

`func (o *ProtocolParametersStatus) GetIcmp() IcmpStatus`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *ProtocolParametersStatus) GetIcmpOk() (*IcmpStatus, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *ProtocolParametersStatus) SetIcmp(v IcmpStatus)`

SetIcmp sets Icmp field to given value.

### HasIcmp

`func (o *ProtocolParametersStatus) HasIcmp() bool`

HasIcmp returns a boolean if a field has been set.

### GetProtocolNumber

`func (o *ProtocolParametersStatus) GetProtocolNumber() int32`

GetProtocolNumber returns the ProtocolNumber field if non-nil, zero value otherwise.

### GetProtocolNumberOk

`func (o *ProtocolParametersStatus) GetProtocolNumberOk() (*int32, bool)`

GetProtocolNumberOk returns a tuple with the ProtocolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolNumber

`func (o *ProtocolParametersStatus) SetProtocolNumber(v int32)`

SetProtocolNumber sets ProtocolNumber field to given value.

### HasProtocolNumber

`func (o *ProtocolParametersStatus) HasProtocolNumber() bool`

HasProtocolNumber returns a boolean if a field has been set.

### GetTcp

`func (o *ProtocolParametersStatus) GetTcp() TcpStatus`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ProtocolParametersStatus) GetTcpOk() (*TcpStatus, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ProtocolParametersStatus) SetTcp(v TcpStatus)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ProtocolParametersStatus) HasTcp() bool`

HasTcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


