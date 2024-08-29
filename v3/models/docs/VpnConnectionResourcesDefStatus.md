# VpnConnectionResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EbgpStatus** | Pointer to [**VpnComponentStatus**](VpnComponentStatus.md) |  | [optional] 
**IpsecTunnelStatus** | Pointer to [**VpnComponentStatus**](VpnComponentStatus.md) |  | [optional] 
**RemoteGatewayReference** | Pointer to [**VpnGatewayReference**](VpnGatewayReference.md) |  | [optional] 
**PeerRoutePrefixList** | Pointer to [**[]IpSubnetStatus**](IpSubnetStatus.md) | IP prefixes learned from the remote gateway over eBGP. | [optional] 
**QosConfig** | Pointer to [**QosConfigStatus**](QosConfigStatus.md) |  | [optional] 
**LocalGatewayRole** | Pointer to **string** | Local gateway role (acceptor or initiator) in the connection.  | [optional] 
**IpsecConfig** | Pointer to [**IpsecConfigStatus**](IpsecConfigStatus.md) |  | [optional] 
**DynamicRoutePriority** | Pointer to **int32** | Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred.  | [optional] 
**LocalGatewayReference** | Pointer to [**VpnGatewayReference**](VpnGatewayReference.md) |  | [optional] 
**LocalRoutePrefixList** | Pointer to [**[]IpSubnetStatus**](IpSubnetStatus.md) | IP prefixes advertised to the remote gateway over eBGP. | [optional] 
**DpdConfig** | Pointer to [**DpdConfigStatus**](DpdConfigStatus.md) |  | [optional] 

## Methods

### NewVpnConnectionResourcesDefStatus

`func NewVpnConnectionResourcesDefStatus() *VpnConnectionResourcesDefStatus`

NewVpnConnectionResourcesDefStatus instantiates a new VpnConnectionResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionResourcesDefStatusWithDefaults

`func NewVpnConnectionResourcesDefStatusWithDefaults() *VpnConnectionResourcesDefStatus`

NewVpnConnectionResourcesDefStatusWithDefaults instantiates a new VpnConnectionResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEbgpStatus

`func (o *VpnConnectionResourcesDefStatus) GetEbgpStatus() VpnComponentStatus`

GetEbgpStatus returns the EbgpStatus field if non-nil, zero value otherwise.

### GetEbgpStatusOk

`func (o *VpnConnectionResourcesDefStatus) GetEbgpStatusOk() (*VpnComponentStatus, bool)`

GetEbgpStatusOk returns a tuple with the EbgpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpStatus

`func (o *VpnConnectionResourcesDefStatus) SetEbgpStatus(v VpnComponentStatus)`

SetEbgpStatus sets EbgpStatus field to given value.

### HasEbgpStatus

`func (o *VpnConnectionResourcesDefStatus) HasEbgpStatus() bool`

HasEbgpStatus returns a boolean if a field has been set.

### GetIpsecTunnelStatus

`func (o *VpnConnectionResourcesDefStatus) GetIpsecTunnelStatus() VpnComponentStatus`

GetIpsecTunnelStatus returns the IpsecTunnelStatus field if non-nil, zero value otherwise.

### GetIpsecTunnelStatusOk

`func (o *VpnConnectionResourcesDefStatus) GetIpsecTunnelStatusOk() (*VpnComponentStatus, bool)`

GetIpsecTunnelStatusOk returns a tuple with the IpsecTunnelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecTunnelStatus

`func (o *VpnConnectionResourcesDefStatus) SetIpsecTunnelStatus(v VpnComponentStatus)`

SetIpsecTunnelStatus sets IpsecTunnelStatus field to given value.

### HasIpsecTunnelStatus

`func (o *VpnConnectionResourcesDefStatus) HasIpsecTunnelStatus() bool`

HasIpsecTunnelStatus returns a boolean if a field has been set.

### GetRemoteGatewayReference

`func (o *VpnConnectionResourcesDefStatus) GetRemoteGatewayReference() VpnGatewayReference`

GetRemoteGatewayReference returns the RemoteGatewayReference field if non-nil, zero value otherwise.

### GetRemoteGatewayReferenceOk

`func (o *VpnConnectionResourcesDefStatus) GetRemoteGatewayReferenceOk() (*VpnGatewayReference, bool)`

GetRemoteGatewayReferenceOk returns a tuple with the RemoteGatewayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewayReference

`func (o *VpnConnectionResourcesDefStatus) SetRemoteGatewayReference(v VpnGatewayReference)`

SetRemoteGatewayReference sets RemoteGatewayReference field to given value.

### HasRemoteGatewayReference

`func (o *VpnConnectionResourcesDefStatus) HasRemoteGatewayReference() bool`

HasRemoteGatewayReference returns a boolean if a field has been set.

### GetPeerRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) GetPeerRoutePrefixList() []IpSubnetStatus`

GetPeerRoutePrefixList returns the PeerRoutePrefixList field if non-nil, zero value otherwise.

### GetPeerRoutePrefixListOk

`func (o *VpnConnectionResourcesDefStatus) GetPeerRoutePrefixListOk() (*[]IpSubnetStatus, bool)`

GetPeerRoutePrefixListOk returns a tuple with the PeerRoutePrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) SetPeerRoutePrefixList(v []IpSubnetStatus)`

SetPeerRoutePrefixList sets PeerRoutePrefixList field to given value.

### HasPeerRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) HasPeerRoutePrefixList() bool`

HasPeerRoutePrefixList returns a boolean if a field has been set.

### GetQosConfig

`func (o *VpnConnectionResourcesDefStatus) GetQosConfig() QosConfigStatus`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *VpnConnectionResourcesDefStatus) GetQosConfigOk() (*QosConfigStatus, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *VpnConnectionResourcesDefStatus) SetQosConfig(v QosConfigStatus)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *VpnConnectionResourcesDefStatus) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetLocalGatewayRole

`func (o *VpnConnectionResourcesDefStatus) GetLocalGatewayRole() string`

GetLocalGatewayRole returns the LocalGatewayRole field if non-nil, zero value otherwise.

### GetLocalGatewayRoleOk

`func (o *VpnConnectionResourcesDefStatus) GetLocalGatewayRoleOk() (*string, bool)`

GetLocalGatewayRoleOk returns a tuple with the LocalGatewayRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayRole

`func (o *VpnConnectionResourcesDefStatus) SetLocalGatewayRole(v string)`

SetLocalGatewayRole sets LocalGatewayRole field to given value.

### HasLocalGatewayRole

`func (o *VpnConnectionResourcesDefStatus) HasLocalGatewayRole() bool`

HasLocalGatewayRole returns a boolean if a field has been set.

### GetIpsecConfig

`func (o *VpnConnectionResourcesDefStatus) GetIpsecConfig() IpsecConfigStatus`

GetIpsecConfig returns the IpsecConfig field if non-nil, zero value otherwise.

### GetIpsecConfigOk

`func (o *VpnConnectionResourcesDefStatus) GetIpsecConfigOk() (*IpsecConfigStatus, bool)`

GetIpsecConfigOk returns a tuple with the IpsecConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecConfig

`func (o *VpnConnectionResourcesDefStatus) SetIpsecConfig(v IpsecConfigStatus)`

SetIpsecConfig sets IpsecConfig field to given value.

### HasIpsecConfig

`func (o *VpnConnectionResourcesDefStatus) HasIpsecConfig() bool`

HasIpsecConfig returns a boolean if a field has been set.

### GetDynamicRoutePriority

`func (o *VpnConnectionResourcesDefStatus) GetDynamicRoutePriority() int32`

GetDynamicRoutePriority returns the DynamicRoutePriority field if non-nil, zero value otherwise.

### GetDynamicRoutePriorityOk

`func (o *VpnConnectionResourcesDefStatus) GetDynamicRoutePriorityOk() (*int32, bool)`

GetDynamicRoutePriorityOk returns a tuple with the DynamicRoutePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRoutePriority

`func (o *VpnConnectionResourcesDefStatus) SetDynamicRoutePriority(v int32)`

SetDynamicRoutePriority sets DynamicRoutePriority field to given value.

### HasDynamicRoutePriority

`func (o *VpnConnectionResourcesDefStatus) HasDynamicRoutePriority() bool`

HasDynamicRoutePriority returns a boolean if a field has been set.

### GetLocalGatewayReference

`func (o *VpnConnectionResourcesDefStatus) GetLocalGatewayReference() VpnGatewayReference`

GetLocalGatewayReference returns the LocalGatewayReference field if non-nil, zero value otherwise.

### GetLocalGatewayReferenceOk

`func (o *VpnConnectionResourcesDefStatus) GetLocalGatewayReferenceOk() (*VpnGatewayReference, bool)`

GetLocalGatewayReferenceOk returns a tuple with the LocalGatewayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayReference

`func (o *VpnConnectionResourcesDefStatus) SetLocalGatewayReference(v VpnGatewayReference)`

SetLocalGatewayReference sets LocalGatewayReference field to given value.

### HasLocalGatewayReference

`func (o *VpnConnectionResourcesDefStatus) HasLocalGatewayReference() bool`

HasLocalGatewayReference returns a boolean if a field has been set.

### GetLocalRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) GetLocalRoutePrefixList() []IpSubnetStatus`

GetLocalRoutePrefixList returns the LocalRoutePrefixList field if non-nil, zero value otherwise.

### GetLocalRoutePrefixListOk

`func (o *VpnConnectionResourcesDefStatus) GetLocalRoutePrefixListOk() (*[]IpSubnetStatus, bool)`

GetLocalRoutePrefixListOk returns a tuple with the LocalRoutePrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) SetLocalRoutePrefixList(v []IpSubnetStatus)`

SetLocalRoutePrefixList sets LocalRoutePrefixList field to given value.

### HasLocalRoutePrefixList

`func (o *VpnConnectionResourcesDefStatus) HasLocalRoutePrefixList() bool`

HasLocalRoutePrefixList returns a boolean if a field has been set.

### GetDpdConfig

`func (o *VpnConnectionResourcesDefStatus) GetDpdConfig() DpdConfigStatus`

GetDpdConfig returns the DpdConfig field if non-nil, zero value otherwise.

### GetDpdConfigOk

`func (o *VpnConnectionResourcesDefStatus) GetDpdConfigOk() (*DpdConfigStatus, bool)`

GetDpdConfigOk returns a tuple with the DpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdConfig

`func (o *VpnConnectionResourcesDefStatus) SetDpdConfig(v DpdConfigStatus)`

SetDpdConfig sets DpdConfig field to given value.

### HasDpdConfig

`func (o *VpnConnectionResourcesDefStatus) HasDpdConfig() bool`

HasDpdConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


