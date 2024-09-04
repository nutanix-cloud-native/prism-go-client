# DirectConnectVirtualInterfaceResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**PeerIpPrefix** | Pointer to **string** | Customer peering IPv4 /30 or /31 prefix. | [optional] 
**PeerAsn** | Pointer to **int32** | Customer peering autonomous system number (ASN). | [optional] 
**ReceivedPrefixList** | Pointer to [**[]IpSubnetStatus**](IpSubnetStatus.md) | IP prefixes learned from the remote gateway over eBGP. | [optional] 
**SentPrefixList** | Pointer to [**[]IpSubnetStatus**](IpSubnetStatus.md) | IP prefixes advertised to the remote gateway over eBGP. | [optional] 
**NutanixPeerIp** | Pointer to **string** | Nutanix peering IP address. | [optional] 
**DirectConnectReference** | Pointer to [**DirectConnectReference**](DirectConnectReference.md) |  | [optional] 
**DynamicRoutePriority** | Pointer to **int32** | Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred.  | [optional] 
**PeeringStatus** | Pointer to [**DirectConnectVifStatus**](DirectConnectVifStatus.md) |  | [optional] 
**NutanixPeerAsn** | Pointer to **int32** | Nutanix peering autonomous system number. | [optional] 
**Md5** | Pointer to **string** | md5 hash for bgp peering. | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 

## Methods

### NewDirectConnectVirtualInterfaceResourcesDefStatus

`func NewDirectConnectVirtualInterfaceResourcesDefStatus() *DirectConnectVirtualInterfaceResourcesDefStatus`

NewDirectConnectVirtualInterfaceResourcesDefStatus instantiates a new DirectConnectVirtualInterfaceResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceResourcesDefStatusWithDefaults

`func NewDirectConnectVirtualInterfaceResourcesDefStatusWithDefaults() *DirectConnectVirtualInterfaceResourcesDefStatus`

NewDirectConnectVirtualInterfaceResourcesDefStatusWithDefaults instantiates a new DirectConnectVirtualInterfaceResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetPeerIpPrefix

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeerIpPrefix() string`

GetPeerIpPrefix returns the PeerIpPrefix field if non-nil, zero value otherwise.

### GetPeerIpPrefixOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeerIpPrefixOk() (*string, bool)`

GetPeerIpPrefixOk returns a tuple with the PeerIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpPrefix

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetPeerIpPrefix(v string)`

SetPeerIpPrefix sets PeerIpPrefix field to given value.

### HasPeerIpPrefix

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasPeerIpPrefix() bool`

HasPeerIpPrefix returns a boolean if a field has been set.

### GetPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetReceivedPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetReceivedPrefixList() []IpSubnetStatus`

GetReceivedPrefixList returns the ReceivedPrefixList field if non-nil, zero value otherwise.

### GetReceivedPrefixListOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetReceivedPrefixListOk() (*[]IpSubnetStatus, bool)`

GetReceivedPrefixListOk returns a tuple with the ReceivedPrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetReceivedPrefixList(v []IpSubnetStatus)`

SetReceivedPrefixList sets ReceivedPrefixList field to given value.

### HasReceivedPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasReceivedPrefixList() bool`

HasReceivedPrefixList returns a boolean if a field has been set.

### GetSentPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetSentPrefixList() []IpSubnetStatus`

GetSentPrefixList returns the SentPrefixList field if non-nil, zero value otherwise.

### GetSentPrefixListOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetSentPrefixListOk() (*[]IpSubnetStatus, bool)`

GetSentPrefixListOk returns a tuple with the SentPrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetSentPrefixList(v []IpSubnetStatus)`

SetSentPrefixList sets SentPrefixList field to given value.

### HasSentPrefixList

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasSentPrefixList() bool`

HasSentPrefixList returns a boolean if a field has been set.

### GetNutanixPeerIp

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetNutanixPeerIp() string`

GetNutanixPeerIp returns the NutanixPeerIp field if non-nil, zero value otherwise.

### GetNutanixPeerIpOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetNutanixPeerIpOk() (*string, bool)`

GetNutanixPeerIpOk returns a tuple with the NutanixPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNutanixPeerIp

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetNutanixPeerIp(v string)`

SetNutanixPeerIp sets NutanixPeerIp field to given value.

### HasNutanixPeerIp

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasNutanixPeerIp() bool`

HasNutanixPeerIp returns a boolean if a field has been set.

### GetDirectConnectReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetDirectConnectReference() DirectConnectReference`

GetDirectConnectReference returns the DirectConnectReference field if non-nil, zero value otherwise.

### GetDirectConnectReferenceOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetDirectConnectReferenceOk() (*DirectConnectReference, bool)`

GetDirectConnectReferenceOk returns a tuple with the DirectConnectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetDirectConnectReference(v DirectConnectReference)`

SetDirectConnectReference sets DirectConnectReference field to given value.

### HasDirectConnectReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasDirectConnectReference() bool`

HasDirectConnectReference returns a boolean if a field has been set.

### GetDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetDynamicRoutePriority() int32`

GetDynamicRoutePriority returns the DynamicRoutePriority field if non-nil, zero value otherwise.

### GetDynamicRoutePriorityOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetDynamicRoutePriorityOk() (*int32, bool)`

GetDynamicRoutePriorityOk returns a tuple with the DynamicRoutePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetDynamicRoutePriority(v int32)`

SetDynamicRoutePriority sets DynamicRoutePriority field to given value.

### HasDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasDynamicRoutePriority() bool`

HasDynamicRoutePriority returns a boolean if a field has been set.

### GetPeeringStatus

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeeringStatus() DirectConnectVifStatus`

GetPeeringStatus returns the PeeringStatus field if non-nil, zero value otherwise.

### GetPeeringStatusOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetPeeringStatusOk() (*DirectConnectVifStatus, bool)`

GetPeeringStatusOk returns a tuple with the PeeringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeeringStatus

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetPeeringStatus(v DirectConnectVifStatus)`

SetPeeringStatus sets PeeringStatus field to given value.

### HasPeeringStatus

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasPeeringStatus() bool`

HasPeeringStatus returns a boolean if a field has been set.

### GetNutanixPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetNutanixPeerAsn() int32`

GetNutanixPeerAsn returns the NutanixPeerAsn field if non-nil, zero value otherwise.

### GetNutanixPeerAsnOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetNutanixPeerAsnOk() (*int32, bool)`

GetNutanixPeerAsnOk returns a tuple with the NutanixPeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNutanixPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetNutanixPeerAsn(v int32)`

SetNutanixPeerAsn sets NutanixPeerAsn field to given value.

### HasNutanixPeerAsn

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasNutanixPeerAsn() bool`

HasNutanixPeerAsn returns a boolean if a field has been set.

### GetMd5

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResourcesDefStatus) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


