# NetworkRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Select a protocol to allow.  Multiple protocols can be allowed by repeating network_rule object.  If a protocol is not configured in the network_rule object then it is allowed. | [optional] 
**Description** | Pointer to **string** | Description for network security rule that is for inbound or outbound | [optional] 
**IpSubnet** | Pointer to [**IpSubnet**](IpSubnet.md) |  | [optional] 
**TcpPortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of TCP ports that are allowed by this rule. | [optional] 
**UdpPortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of UDP ports that are allowed by this rule. | [optional] 
**PeerSpecificationType** | Pointer to **string** | The set of categories that matching VMs need to have. | [optional] 
**Filter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 
**ServiceGroupList** | Pointer to [**[]ServiceGroupReference**](ServiceGroupReference.md) | List of service groups associated with this rule. The exiting fields for protocol or ports is not recommended for use and will be deprecated for these new fields at the API level. | [optional] 
**IcmpTypeCodeList** | Pointer to [**[]NetworkRuleIcmpTypeCodeListInner**](NetworkRuleIcmpTypeCodeListInner.md) | List of ICMP types and codes allowed by this rule. | [optional] 
**NetworkFunctionChainReference** | Pointer to [**NetworkFunctionChainReference**](NetworkFunctionChainReference.md) |  | [optional] 
**ExpirationTime** | Pointer to **string** | Timestamp of expiration time. | [optional] 
**RuleId** | Pointer to **int32** | Unique identifier for inbound or outbound rule. This is system generated and used internally. User should not set this field while creating a new rule or should not modify it while updating the existing rule. | [optional] 
**AddressGroupInclusionList** | Pointer to [**[]AddressGroupReference**](AddressGroupReference.md) | List of address groups that are allowed access by this rule | [optional] 

## Methods

### NewNetworkRule

`func NewNetworkRule() *NetworkRule`

NewNetworkRule instantiates a new NetworkRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRuleWithDefaults

`func NewNetworkRuleWithDefaults() *NetworkRule`

NewNetworkRuleWithDefaults instantiates a new NetworkRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *NetworkRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworkRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpSubnet

`func (o *NetworkRule) GetIpSubnet() IpSubnet`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *NetworkRule) GetIpSubnetOk() (*IpSubnet, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *NetworkRule) SetIpSubnet(v IpSubnet)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *NetworkRule) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetTcpPortRangeList

`func (o *NetworkRule) GetTcpPortRangeList() []PortRange`

GetTcpPortRangeList returns the TcpPortRangeList field if non-nil, zero value otherwise.

### GetTcpPortRangeListOk

`func (o *NetworkRule) GetTcpPortRangeListOk() (*[]PortRange, bool)`

GetTcpPortRangeListOk returns a tuple with the TcpPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPortRangeList

`func (o *NetworkRule) SetTcpPortRangeList(v []PortRange)`

SetTcpPortRangeList sets TcpPortRangeList field to given value.

### HasTcpPortRangeList

`func (o *NetworkRule) HasTcpPortRangeList() bool`

HasTcpPortRangeList returns a boolean if a field has been set.

### GetUdpPortRangeList

`func (o *NetworkRule) GetUdpPortRangeList() []PortRange`

GetUdpPortRangeList returns the UdpPortRangeList field if non-nil, zero value otherwise.

### GetUdpPortRangeListOk

`func (o *NetworkRule) GetUdpPortRangeListOk() (*[]PortRange, bool)`

GetUdpPortRangeListOk returns a tuple with the UdpPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpPortRangeList

`func (o *NetworkRule) SetUdpPortRangeList(v []PortRange)`

SetUdpPortRangeList sets UdpPortRangeList field to given value.

### HasUdpPortRangeList

`func (o *NetworkRule) HasUdpPortRangeList() bool`

HasUdpPortRangeList returns a boolean if a field has been set.

### GetPeerSpecificationType

`func (o *NetworkRule) GetPeerSpecificationType() string`

GetPeerSpecificationType returns the PeerSpecificationType field if non-nil, zero value otherwise.

### GetPeerSpecificationTypeOk

`func (o *NetworkRule) GetPeerSpecificationTypeOk() (*string, bool)`

GetPeerSpecificationTypeOk returns a tuple with the PeerSpecificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSpecificationType

`func (o *NetworkRule) SetPeerSpecificationType(v string)`

SetPeerSpecificationType sets PeerSpecificationType field to given value.

### HasPeerSpecificationType

`func (o *NetworkRule) HasPeerSpecificationType() bool`

HasPeerSpecificationType returns a boolean if a field has been set.

### GetFilter

`func (o *NetworkRule) GetFilter() CategoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NetworkRule) GetFilterOk() (*CategoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NetworkRule) SetFilter(v CategoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *NetworkRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetServiceGroupList

`func (o *NetworkRule) GetServiceGroupList() []ServiceGroupReference`

GetServiceGroupList returns the ServiceGroupList field if non-nil, zero value otherwise.

### GetServiceGroupListOk

`func (o *NetworkRule) GetServiceGroupListOk() (*[]ServiceGroupReference, bool)`

GetServiceGroupListOk returns a tuple with the ServiceGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroupList

`func (o *NetworkRule) SetServiceGroupList(v []ServiceGroupReference)`

SetServiceGroupList sets ServiceGroupList field to given value.

### HasServiceGroupList

`func (o *NetworkRule) HasServiceGroupList() bool`

HasServiceGroupList returns a boolean if a field has been set.

### GetIcmpTypeCodeList

`func (o *NetworkRule) GetIcmpTypeCodeList() []NetworkRuleIcmpTypeCodeListInner`

GetIcmpTypeCodeList returns the IcmpTypeCodeList field if non-nil, zero value otherwise.

### GetIcmpTypeCodeListOk

`func (o *NetworkRule) GetIcmpTypeCodeListOk() (*[]NetworkRuleIcmpTypeCodeListInner, bool)`

GetIcmpTypeCodeListOk returns a tuple with the IcmpTypeCodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpTypeCodeList

`func (o *NetworkRule) SetIcmpTypeCodeList(v []NetworkRuleIcmpTypeCodeListInner)`

SetIcmpTypeCodeList sets IcmpTypeCodeList field to given value.

### HasIcmpTypeCodeList

`func (o *NetworkRule) HasIcmpTypeCodeList() bool`

HasIcmpTypeCodeList returns a boolean if a field has been set.

### GetNetworkFunctionChainReference

`func (o *NetworkRule) GetNetworkFunctionChainReference() NetworkFunctionChainReference`

GetNetworkFunctionChainReference returns the NetworkFunctionChainReference field if non-nil, zero value otherwise.

### GetNetworkFunctionChainReferenceOk

`func (o *NetworkRule) GetNetworkFunctionChainReferenceOk() (*NetworkFunctionChainReference, bool)`

GetNetworkFunctionChainReferenceOk returns a tuple with the NetworkFunctionChainReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionChainReference

`func (o *NetworkRule) SetNetworkFunctionChainReference(v NetworkFunctionChainReference)`

SetNetworkFunctionChainReference sets NetworkFunctionChainReference field to given value.

### HasNetworkFunctionChainReference

`func (o *NetworkRule) HasNetworkFunctionChainReference() bool`

HasNetworkFunctionChainReference returns a boolean if a field has been set.

### GetExpirationTime

`func (o *NetworkRule) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *NetworkRule) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *NetworkRule) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *NetworkRule) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetRuleId

`func (o *NetworkRule) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *NetworkRule) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *NetworkRule) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *NetworkRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetAddressGroupInclusionList

`func (o *NetworkRule) GetAddressGroupInclusionList() []AddressGroupReference`

GetAddressGroupInclusionList returns the AddressGroupInclusionList field if non-nil, zero value otherwise.

### GetAddressGroupInclusionListOk

`func (o *NetworkRule) GetAddressGroupInclusionListOk() (*[]AddressGroupReference, bool)`

GetAddressGroupInclusionListOk returns a tuple with the AddressGroupInclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressGroupInclusionList

`func (o *NetworkRule) SetAddressGroupInclusionList(v []AddressGroupReference)`

SetAddressGroupInclusionList sets AddressGroupInclusionList field to given value.

### HasAddressGroupInclusionList

`func (o *NetworkRule) HasAddressGroupInclusionList() bool`

HasAddressGroupInclusionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


