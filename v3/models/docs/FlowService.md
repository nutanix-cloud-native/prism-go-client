# FlowService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TcpPortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of TCP ports in the service | [optional] 
**Protocol** | Pointer to **string** | protocol for the service | [optional] 
**UdpPortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of UDP ports in the service | [optional] 
**IcmpTypeCodeList** | Pointer to [**[]NetworkRuleStatusIcmpTypeCodeListInner**](NetworkRuleStatusIcmpTypeCodeListInner.md) | List of ICMP types and codes in the service | [optional] 

## Methods

### NewFlowService

`func NewFlowService() *FlowService`

NewFlowService instantiates a new FlowService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowServiceWithDefaults

`func NewFlowServiceWithDefaults() *FlowService`

NewFlowServiceWithDefaults instantiates a new FlowService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcpPortRangeList

`func (o *FlowService) GetTcpPortRangeList() []PortRange`

GetTcpPortRangeList returns the TcpPortRangeList field if non-nil, zero value otherwise.

### GetTcpPortRangeListOk

`func (o *FlowService) GetTcpPortRangeListOk() (*[]PortRange, bool)`

GetTcpPortRangeListOk returns a tuple with the TcpPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPortRangeList

`func (o *FlowService) SetTcpPortRangeList(v []PortRange)`

SetTcpPortRangeList sets TcpPortRangeList field to given value.

### HasTcpPortRangeList

`func (o *FlowService) HasTcpPortRangeList() bool`

HasTcpPortRangeList returns a boolean if a field has been set.

### GetProtocol

`func (o *FlowService) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FlowService) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FlowService) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FlowService) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUdpPortRangeList

`func (o *FlowService) GetUdpPortRangeList() []PortRange`

GetUdpPortRangeList returns the UdpPortRangeList field if non-nil, zero value otherwise.

### GetUdpPortRangeListOk

`func (o *FlowService) GetUdpPortRangeListOk() (*[]PortRange, bool)`

GetUdpPortRangeListOk returns a tuple with the UdpPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpPortRangeList

`func (o *FlowService) SetUdpPortRangeList(v []PortRange)`

SetUdpPortRangeList sets UdpPortRangeList field to given value.

### HasUdpPortRangeList

`func (o *FlowService) HasUdpPortRangeList() bool`

HasUdpPortRangeList returns a boolean if a field has been set.

### GetIcmpTypeCodeList

`func (o *FlowService) GetIcmpTypeCodeList() []NetworkRuleStatusIcmpTypeCodeListInner`

GetIcmpTypeCodeList returns the IcmpTypeCodeList field if non-nil, zero value otherwise.

### GetIcmpTypeCodeListOk

`func (o *FlowService) GetIcmpTypeCodeListOk() (*[]NetworkRuleStatusIcmpTypeCodeListInner, bool)`

GetIcmpTypeCodeListOk returns a tuple with the IcmpTypeCodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpTypeCodeList

`func (o *FlowService) SetIcmpTypeCodeList(v []NetworkRuleStatusIcmpTypeCodeListInner)`

SetIcmpTypeCodeList sets IcmpTypeCodeList field to given value.

### HasIcmpTypeCodeList

`func (o *FlowService) HasIcmpTypeCodeList() bool`

HasIcmpTypeCodeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


