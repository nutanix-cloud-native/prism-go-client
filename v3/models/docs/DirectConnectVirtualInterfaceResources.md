# DirectConnectVirtualInterfaceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**DirectConnectReference** | [**DirectConnectReference**](DirectConnectReference.md) |  | 
**DynamicRoutePriority** | Pointer to **int32** | Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred.  | [optional] 
**PeerAsn** | **int32** | Autonomous system number | 
**PeerIpPrefix** | **string** | Customer peering IPv4 /30 or /31 prefix. | 
**Md5** | Pointer to **string** | md5 hash for bgp peering. | [optional] 

## Methods

### NewDirectConnectVirtualInterfaceResources

`func NewDirectConnectVirtualInterfaceResources(directConnectReference DirectConnectReference, peerAsn int32, peerIpPrefix string, ) *DirectConnectVirtualInterfaceResources`

NewDirectConnectVirtualInterfaceResources instantiates a new DirectConnectVirtualInterfaceResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceResourcesWithDefaults

`func NewDirectConnectVirtualInterfaceResourcesWithDefaults() *DirectConnectVirtualInterfaceResources`

NewDirectConnectVirtualInterfaceResourcesWithDefaults instantiates a new DirectConnectVirtualInterfaceResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *DirectConnectVirtualInterfaceResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *DirectConnectVirtualInterfaceResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *DirectConnectVirtualInterfaceResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *DirectConnectVirtualInterfaceResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResources) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *DirectConnectVirtualInterfaceResources) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResources) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *DirectConnectVirtualInterfaceResources) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetDirectConnectReference

`func (o *DirectConnectVirtualInterfaceResources) GetDirectConnectReference() DirectConnectReference`

GetDirectConnectReference returns the DirectConnectReference field if non-nil, zero value otherwise.

### GetDirectConnectReferenceOk

`func (o *DirectConnectVirtualInterfaceResources) GetDirectConnectReferenceOk() (*DirectConnectReference, bool)`

GetDirectConnectReferenceOk returns a tuple with the DirectConnectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectReference

`func (o *DirectConnectVirtualInterfaceResources) SetDirectConnectReference(v DirectConnectReference)`

SetDirectConnectReference sets DirectConnectReference field to given value.


### GetDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResources) GetDynamicRoutePriority() int32`

GetDynamicRoutePriority returns the DynamicRoutePriority field if non-nil, zero value otherwise.

### GetDynamicRoutePriorityOk

`func (o *DirectConnectVirtualInterfaceResources) GetDynamicRoutePriorityOk() (*int32, bool)`

GetDynamicRoutePriorityOk returns a tuple with the DynamicRoutePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResources) SetDynamicRoutePriority(v int32)`

SetDynamicRoutePriority sets DynamicRoutePriority field to given value.

### HasDynamicRoutePriority

`func (o *DirectConnectVirtualInterfaceResources) HasDynamicRoutePriority() bool`

HasDynamicRoutePriority returns a boolean if a field has been set.

### GetPeerAsn

`func (o *DirectConnectVirtualInterfaceResources) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *DirectConnectVirtualInterfaceResources) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *DirectConnectVirtualInterfaceResources) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.


### GetPeerIpPrefix

`func (o *DirectConnectVirtualInterfaceResources) GetPeerIpPrefix() string`

GetPeerIpPrefix returns the PeerIpPrefix field if non-nil, zero value otherwise.

### GetPeerIpPrefixOk

`func (o *DirectConnectVirtualInterfaceResources) GetPeerIpPrefixOk() (*string, bool)`

GetPeerIpPrefixOk returns a tuple with the PeerIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIpPrefix

`func (o *DirectConnectVirtualInterfaceResources) SetPeerIpPrefix(v string)`

SetPeerIpPrefix sets PeerIpPrefix field to given value.


### GetMd5

`func (o *DirectConnectVirtualInterfaceResources) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *DirectConnectVirtualInterfaceResources) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *DirectConnectVirtualInterfaceResources) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *DirectConnectVirtualInterfaceResources) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


