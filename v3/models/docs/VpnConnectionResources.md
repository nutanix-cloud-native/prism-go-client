# VpnConnectionResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalGatewayRole** | **string** | Local gateway role (acceptor or initiator) in the connection.  | 
**QosConfig** | Pointer to [**QosConfig**](QosConfig.md) |  | [optional] 
**RemoteGatewayReference** | [**VpnGatewayReference**](VpnGatewayReference.md) |  | 
**IpsecConfig** | [**IpsecConfig**](IpsecConfig.md) |  | 
**DynamicRoutePriority** | Pointer to **int32** | Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred.  | [optional] 
**LocalGatewayReference** | [**VpnGatewayReference**](VpnGatewayReference.md) |  | 
**DpdConfig** | Pointer to [**DpdConfig**](DpdConfig.md) |  | [optional] 

## Methods

### NewVpnConnectionResources

`func NewVpnConnectionResources(localGatewayRole string, remoteGatewayReference VpnGatewayReference, ipsecConfig IpsecConfig, localGatewayReference VpnGatewayReference, ) *VpnConnectionResources`

NewVpnConnectionResources instantiates a new VpnConnectionResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionResourcesWithDefaults

`func NewVpnConnectionResourcesWithDefaults() *VpnConnectionResources`

NewVpnConnectionResourcesWithDefaults instantiates a new VpnConnectionResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalGatewayRole

`func (o *VpnConnectionResources) GetLocalGatewayRole() string`

GetLocalGatewayRole returns the LocalGatewayRole field if non-nil, zero value otherwise.

### GetLocalGatewayRoleOk

`func (o *VpnConnectionResources) GetLocalGatewayRoleOk() (*string, bool)`

GetLocalGatewayRoleOk returns a tuple with the LocalGatewayRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayRole

`func (o *VpnConnectionResources) SetLocalGatewayRole(v string)`

SetLocalGatewayRole sets LocalGatewayRole field to given value.


### GetQosConfig

`func (o *VpnConnectionResources) GetQosConfig() QosConfig`

GetQosConfig returns the QosConfig field if non-nil, zero value otherwise.

### GetQosConfigOk

`func (o *VpnConnectionResources) GetQosConfigOk() (*QosConfig, bool)`

GetQosConfigOk returns a tuple with the QosConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosConfig

`func (o *VpnConnectionResources) SetQosConfig(v QosConfig)`

SetQosConfig sets QosConfig field to given value.

### HasQosConfig

`func (o *VpnConnectionResources) HasQosConfig() bool`

HasQosConfig returns a boolean if a field has been set.

### GetRemoteGatewayReference

`func (o *VpnConnectionResources) GetRemoteGatewayReference() VpnGatewayReference`

GetRemoteGatewayReference returns the RemoteGatewayReference field if non-nil, zero value otherwise.

### GetRemoteGatewayReferenceOk

`func (o *VpnConnectionResources) GetRemoteGatewayReferenceOk() (*VpnGatewayReference, bool)`

GetRemoteGatewayReferenceOk returns a tuple with the RemoteGatewayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewayReference

`func (o *VpnConnectionResources) SetRemoteGatewayReference(v VpnGatewayReference)`

SetRemoteGatewayReference sets RemoteGatewayReference field to given value.


### GetIpsecConfig

`func (o *VpnConnectionResources) GetIpsecConfig() IpsecConfig`

GetIpsecConfig returns the IpsecConfig field if non-nil, zero value otherwise.

### GetIpsecConfigOk

`func (o *VpnConnectionResources) GetIpsecConfigOk() (*IpsecConfig, bool)`

GetIpsecConfigOk returns a tuple with the IpsecConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecConfig

`func (o *VpnConnectionResources) SetIpsecConfig(v IpsecConfig)`

SetIpsecConfig sets IpsecConfig field to given value.


### GetDynamicRoutePriority

`func (o *VpnConnectionResources) GetDynamicRoutePriority() int32`

GetDynamicRoutePriority returns the DynamicRoutePriority field if non-nil, zero value otherwise.

### GetDynamicRoutePriorityOk

`func (o *VpnConnectionResources) GetDynamicRoutePriorityOk() (*int32, bool)`

GetDynamicRoutePriorityOk returns a tuple with the DynamicRoutePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRoutePriority

`func (o *VpnConnectionResources) SetDynamicRoutePriority(v int32)`

SetDynamicRoutePriority sets DynamicRoutePriority field to given value.

### HasDynamicRoutePriority

`func (o *VpnConnectionResources) HasDynamicRoutePriority() bool`

HasDynamicRoutePriority returns a boolean if a field has been set.

### GetLocalGatewayReference

`func (o *VpnConnectionResources) GetLocalGatewayReference() VpnGatewayReference`

GetLocalGatewayReference returns the LocalGatewayReference field if non-nil, zero value otherwise.

### GetLocalGatewayReferenceOk

`func (o *VpnConnectionResources) GetLocalGatewayReferenceOk() (*VpnGatewayReference, bool)`

GetLocalGatewayReferenceOk returns a tuple with the LocalGatewayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGatewayReference

`func (o *VpnConnectionResources) SetLocalGatewayReference(v VpnGatewayReference)`

SetLocalGatewayReference sets LocalGatewayReference field to given value.


### GetDpdConfig

`func (o *VpnConnectionResources) GetDpdConfig() DpdConfig`

GetDpdConfig returns the DpdConfig field if non-nil, zero value otherwise.

### GetDpdConfigOk

`func (o *VpnConnectionResources) GetDpdConfigOk() (*DpdConfig, bool)`

GetDpdConfigOk returns a tuple with the DpdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdConfig

`func (o *VpnConnectionResources) SetDpdConfig(v DpdConfig)`

SetDpdConfig sets DpdConfig field to given value.

### HasDpdConfig

`func (o *VpnConnectionResources) HasDpdConfig() bool`

HasDpdConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


