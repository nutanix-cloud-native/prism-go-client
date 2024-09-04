# SubnetResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetType** | **string** |  | 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**VirtualSwitchUuid** | Pointer to **string** | Reference to virtual switch | [optional] 
**IsExternal** | Pointer to **bool** | Whether the subnet is external subnet or not. | [optional] 
**VswitchName** | Pointer to **string** |  | [optional] 
**NetworkFunctionChainReference** | Pointer to [**NetworkFunctionChainReference**](NetworkFunctionChainReference.md) |  | [optional] 
**ReservedIpAddressList** | Pointer to **[]string** | List of IPs that are not considered while allocating IP addresses to Atlas ports.  | [optional] 
**AvailabilityZoneReferenceList** | Pointer to [**[]AvailabilityZoneReference**](AvailabilityZoneReference.md) | List of availability zones from which resources are derived (Only supported on Xi).  | [optional] 
**AdvancedNetworking** | Pointer to **bool** | Whether the subnet should be realized on OVN stack or not. | [optional] 
**IpConfig** | Pointer to [**IpConfig**](IpConfig.md) |  | [optional] 
**EnableNat** | Pointer to **bool** | Whether NAT should be performed for VPCs attaching to the subnet. This field is supported only for external subnets. NAT is enabled by default on external subnets.  | [optional] 
**ExternalConnectivityState** | Pointer to **string** | External connectivity state (Only supported on Xi) | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 

## Methods

### NewSubnetResources

`func NewSubnetResources(subnetType string, ) *SubnetResources`

NewSubnetResources instantiates a new SubnetResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetResourcesWithDefaults

`func NewSubnetResourcesWithDefaults() *SubnetResources`

NewSubnetResourcesWithDefaults instantiates a new SubnetResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetType

`func (o *SubnetResources) GetSubnetType() string`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *SubnetResources) GetSubnetTypeOk() (*string, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *SubnetResources) SetSubnetType(v string)`

SetSubnetType sets SubnetType field to given value.


### GetVpcReference

`func (o *SubnetResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *SubnetResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *SubnetResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *SubnetResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetVirtualSwitchUuid

`func (o *SubnetResources) GetVirtualSwitchUuid() string`

GetVirtualSwitchUuid returns the VirtualSwitchUuid field if non-nil, zero value otherwise.

### GetVirtualSwitchUuidOk

`func (o *SubnetResources) GetVirtualSwitchUuidOk() (*string, bool)`

GetVirtualSwitchUuidOk returns a tuple with the VirtualSwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitchUuid

`func (o *SubnetResources) SetVirtualSwitchUuid(v string)`

SetVirtualSwitchUuid sets VirtualSwitchUuid field to given value.

### HasVirtualSwitchUuid

`func (o *SubnetResources) HasVirtualSwitchUuid() bool`

HasVirtualSwitchUuid returns a boolean if a field has been set.

### GetIsExternal

`func (o *SubnetResources) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *SubnetResources) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *SubnetResources) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *SubnetResources) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetVswitchName

`func (o *SubnetResources) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *SubnetResources) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *SubnetResources) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *SubnetResources) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetNetworkFunctionChainReference

`func (o *SubnetResources) GetNetworkFunctionChainReference() NetworkFunctionChainReference`

GetNetworkFunctionChainReference returns the NetworkFunctionChainReference field if non-nil, zero value otherwise.

### GetNetworkFunctionChainReferenceOk

`func (o *SubnetResources) GetNetworkFunctionChainReferenceOk() (*NetworkFunctionChainReference, bool)`

GetNetworkFunctionChainReferenceOk returns a tuple with the NetworkFunctionChainReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionChainReference

`func (o *SubnetResources) SetNetworkFunctionChainReference(v NetworkFunctionChainReference)`

SetNetworkFunctionChainReference sets NetworkFunctionChainReference field to given value.

### HasNetworkFunctionChainReference

`func (o *SubnetResources) HasNetworkFunctionChainReference() bool`

HasNetworkFunctionChainReference returns a boolean if a field has been set.

### GetReservedIpAddressList

`func (o *SubnetResources) GetReservedIpAddressList() []string`

GetReservedIpAddressList returns the ReservedIpAddressList field if non-nil, zero value otherwise.

### GetReservedIpAddressListOk

`func (o *SubnetResources) GetReservedIpAddressListOk() (*[]string, bool)`

GetReservedIpAddressListOk returns a tuple with the ReservedIpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpAddressList

`func (o *SubnetResources) SetReservedIpAddressList(v []string)`

SetReservedIpAddressList sets ReservedIpAddressList field to given value.

### HasReservedIpAddressList

`func (o *SubnetResources) HasReservedIpAddressList() bool`

HasReservedIpAddressList returns a boolean if a field has been set.

### GetAvailabilityZoneReferenceList

`func (o *SubnetResources) GetAvailabilityZoneReferenceList() []AvailabilityZoneReference`

GetAvailabilityZoneReferenceList returns the AvailabilityZoneReferenceList field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceListOk

`func (o *SubnetResources) GetAvailabilityZoneReferenceListOk() (*[]AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceListOk returns a tuple with the AvailabilityZoneReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReferenceList

`func (o *SubnetResources) SetAvailabilityZoneReferenceList(v []AvailabilityZoneReference)`

SetAvailabilityZoneReferenceList sets AvailabilityZoneReferenceList field to given value.

### HasAvailabilityZoneReferenceList

`func (o *SubnetResources) HasAvailabilityZoneReferenceList() bool`

HasAvailabilityZoneReferenceList returns a boolean if a field has been set.

### GetAdvancedNetworking

`func (o *SubnetResources) GetAdvancedNetworking() bool`

GetAdvancedNetworking returns the AdvancedNetworking field if non-nil, zero value otherwise.

### GetAdvancedNetworkingOk

`func (o *SubnetResources) GetAdvancedNetworkingOk() (*bool, bool)`

GetAdvancedNetworkingOk returns a tuple with the AdvancedNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedNetworking

`func (o *SubnetResources) SetAdvancedNetworking(v bool)`

SetAdvancedNetworking sets AdvancedNetworking field to given value.

### HasAdvancedNetworking

`func (o *SubnetResources) HasAdvancedNetworking() bool`

HasAdvancedNetworking returns a boolean if a field has been set.

### GetIpConfig

`func (o *SubnetResources) GetIpConfig() IpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *SubnetResources) GetIpConfigOk() (*IpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *SubnetResources) SetIpConfig(v IpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *SubnetResources) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetEnableNat

`func (o *SubnetResources) GetEnableNat() bool`

GetEnableNat returns the EnableNat field if non-nil, zero value otherwise.

### GetEnableNatOk

`func (o *SubnetResources) GetEnableNatOk() (*bool, bool)`

GetEnableNatOk returns a tuple with the EnableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNat

`func (o *SubnetResources) SetEnableNat(v bool)`

SetEnableNat sets EnableNat field to given value.

### HasEnableNat

`func (o *SubnetResources) HasEnableNat() bool`

HasEnableNat returns a boolean if a field has been set.

### GetExternalConnectivityState

`func (o *SubnetResources) GetExternalConnectivityState() string`

GetExternalConnectivityState returns the ExternalConnectivityState field if non-nil, zero value otherwise.

### GetExternalConnectivityStateOk

`func (o *SubnetResources) GetExternalConnectivityStateOk() (*string, bool)`

GetExternalConnectivityStateOk returns a tuple with the ExternalConnectivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivityState

`func (o *SubnetResources) SetExternalConnectivityState(v string)`

SetExternalConnectivityState sets ExternalConnectivityState field to given value.

### HasExternalConnectivityState

`func (o *SubnetResources) HasExternalConnectivityState() bool`

HasExternalConnectivityState returns a boolean if a field has been set.

### GetVlanId

`func (o *SubnetResources) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SubnetResources) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SubnetResources) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SubnetResources) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *SubnetResources) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *SubnetResources) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *SubnetResources) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *SubnetResources) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


