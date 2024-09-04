# PacketTraceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IcmpCode** | Pointer to **int32** | ICMP code | [optional] 
**DestinationIp** | **string** | Destination IP Address | 
**Protocol** | Pointer to **int32** | L4 Protocol | [optional] 
**SourceIp** | **string** | Source IP Address | 
**SourcePort** | Pointer to **int32** | Source port | [optional] 
**DestinationPort** | Pointer to **int32** | Destination port | [optional] 
**IcmpType** | Pointer to **int32** | ICMP type | [optional] 

## Methods

### NewPacketTraceInput

`func NewPacketTraceInput(destinationIp string, sourceIp string, ) *PacketTraceInput`

NewPacketTraceInput instantiates a new PacketTraceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketTraceInputWithDefaults

`func NewPacketTraceInputWithDefaults() *PacketTraceInput`

NewPacketTraceInputWithDefaults instantiates a new PacketTraceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcmpCode

`func (o *PacketTraceInput) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *PacketTraceInput) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *PacketTraceInput) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *PacketTraceInput) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### GetDestinationIp

`func (o *PacketTraceInput) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *PacketTraceInput) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *PacketTraceInput) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.


### GetProtocol

`func (o *PacketTraceInput) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PacketTraceInput) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PacketTraceInput) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PacketTraceInput) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceIp

`func (o *PacketTraceInput) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *PacketTraceInput) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *PacketTraceInput) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.


### GetSourcePort

`func (o *PacketTraceInput) GetSourcePort() int32`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *PacketTraceInput) GetSourcePortOk() (*int32, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *PacketTraceInput) SetSourcePort(v int32)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *PacketTraceInput) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetDestinationPort

`func (o *PacketTraceInput) GetDestinationPort() int32`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *PacketTraceInput) GetDestinationPortOk() (*int32, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *PacketTraceInput) SetDestinationPort(v int32)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *PacketTraceInput) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetIcmpType

`func (o *PacketTraceInput) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *PacketTraceInput) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *PacketTraceInput) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *PacketTraceInput) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


