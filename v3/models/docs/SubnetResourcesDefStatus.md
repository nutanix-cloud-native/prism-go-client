# SubnetResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetType** | Pointer to **string** |  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**VirtualSwitchUuid** | Pointer to **string** | Reference to virtual switch | [optional] 
**VswitchName** | Pointer to **string** |  | [optional] 
**NetworkFunctionChainReference** | Pointer to [**NetworkFunctionChainReference**](NetworkFunctionChainReference.md) |  | [optional] 
**IpUsageStats** | Pointer to [**IpUsageStats**](IpUsageStats.md) |  | [optional] 
**IpConfig** | Pointer to [**IpConfig**](IpConfig.md) |  | [optional] 
**ReservedIpAddressList** | Pointer to **[]string** | List of IPs that are not considered while allocating IP addresses to Atlas ports.  | [optional] 
**EnableNat** | Pointer to **bool** | Whether NAT should be performed for VPCs attaching to the subnet. This field is supported only for external subnets. NAT is enabled by default on external subnets.  | [optional] 
**AdvancedNetworking** | Pointer to **bool** | Whether the subnet should be realized on OVN stack or not. | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**MigrationState** | Pointer to **string** |  | [optional] 
**AvailabilityZoneReferenceList** | Pointer to [**[]AvailabilityZoneReference**](AvailabilityZoneReference.md) | List of availability zones from which resources are derived (Only supported on Xi).  | [optional] 
**ExternalConnectivityState** | Pointer to **string** | External connectivity state (Only supported on Xi) | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 
**IsExternal** | Pointer to **bool** | Whether the subnet is external subnet or not. | [optional] 

## Methods

### NewSubnetResourcesDefStatus

`func NewSubnetResourcesDefStatus() *SubnetResourcesDefStatus`

NewSubnetResourcesDefStatus instantiates a new SubnetResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetResourcesDefStatusWithDefaults

`func NewSubnetResourcesDefStatusWithDefaults() *SubnetResourcesDefStatus`

NewSubnetResourcesDefStatusWithDefaults instantiates a new SubnetResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetType

`func (o *SubnetResourcesDefStatus) GetSubnetType() string`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *SubnetResourcesDefStatus) GetSubnetTypeOk() (*string, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *SubnetResourcesDefStatus) SetSubnetType(v string)`

SetSubnetType sets SubnetType field to given value.

### HasSubnetType

`func (o *SubnetResourcesDefStatus) HasSubnetType() bool`

HasSubnetType returns a boolean if a field has been set.

### GetVpcReference

`func (o *SubnetResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *SubnetResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *SubnetResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *SubnetResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetVirtualSwitchUuid

`func (o *SubnetResourcesDefStatus) GetVirtualSwitchUuid() string`

GetVirtualSwitchUuid returns the VirtualSwitchUuid field if non-nil, zero value otherwise.

### GetVirtualSwitchUuidOk

`func (o *SubnetResourcesDefStatus) GetVirtualSwitchUuidOk() (*string, bool)`

GetVirtualSwitchUuidOk returns a tuple with the VirtualSwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSwitchUuid

`func (o *SubnetResourcesDefStatus) SetVirtualSwitchUuid(v string)`

SetVirtualSwitchUuid sets VirtualSwitchUuid field to given value.

### HasVirtualSwitchUuid

`func (o *SubnetResourcesDefStatus) HasVirtualSwitchUuid() bool`

HasVirtualSwitchUuid returns a boolean if a field has been set.

### GetVswitchName

`func (o *SubnetResourcesDefStatus) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *SubnetResourcesDefStatus) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *SubnetResourcesDefStatus) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *SubnetResourcesDefStatus) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetNetworkFunctionChainReference

`func (o *SubnetResourcesDefStatus) GetNetworkFunctionChainReference() NetworkFunctionChainReference`

GetNetworkFunctionChainReference returns the NetworkFunctionChainReference field if non-nil, zero value otherwise.

### GetNetworkFunctionChainReferenceOk

`func (o *SubnetResourcesDefStatus) GetNetworkFunctionChainReferenceOk() (*NetworkFunctionChainReference, bool)`

GetNetworkFunctionChainReferenceOk returns a tuple with the NetworkFunctionChainReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionChainReference

`func (o *SubnetResourcesDefStatus) SetNetworkFunctionChainReference(v NetworkFunctionChainReference)`

SetNetworkFunctionChainReference sets NetworkFunctionChainReference field to given value.

### HasNetworkFunctionChainReference

`func (o *SubnetResourcesDefStatus) HasNetworkFunctionChainReference() bool`

HasNetworkFunctionChainReference returns a boolean if a field has been set.

### GetIpUsageStats

`func (o *SubnetResourcesDefStatus) GetIpUsageStats() IpUsageStats`

GetIpUsageStats returns the IpUsageStats field if non-nil, zero value otherwise.

### GetIpUsageStatsOk

`func (o *SubnetResourcesDefStatus) GetIpUsageStatsOk() (*IpUsageStats, bool)`

GetIpUsageStatsOk returns a tuple with the IpUsageStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpUsageStats

`func (o *SubnetResourcesDefStatus) SetIpUsageStats(v IpUsageStats)`

SetIpUsageStats sets IpUsageStats field to given value.

### HasIpUsageStats

`func (o *SubnetResourcesDefStatus) HasIpUsageStats() bool`

HasIpUsageStats returns a boolean if a field has been set.

### GetIpConfig

`func (o *SubnetResourcesDefStatus) GetIpConfig() IpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *SubnetResourcesDefStatus) GetIpConfigOk() (*IpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *SubnetResourcesDefStatus) SetIpConfig(v IpConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *SubnetResourcesDefStatus) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetReservedIpAddressList

`func (o *SubnetResourcesDefStatus) GetReservedIpAddressList() []string`

GetReservedIpAddressList returns the ReservedIpAddressList field if non-nil, zero value otherwise.

### GetReservedIpAddressListOk

`func (o *SubnetResourcesDefStatus) GetReservedIpAddressListOk() (*[]string, bool)`

GetReservedIpAddressListOk returns a tuple with the ReservedIpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpAddressList

`func (o *SubnetResourcesDefStatus) SetReservedIpAddressList(v []string)`

SetReservedIpAddressList sets ReservedIpAddressList field to given value.

### HasReservedIpAddressList

`func (o *SubnetResourcesDefStatus) HasReservedIpAddressList() bool`

HasReservedIpAddressList returns a boolean if a field has been set.

### GetEnableNat

`func (o *SubnetResourcesDefStatus) GetEnableNat() bool`

GetEnableNat returns the EnableNat field if non-nil, zero value otherwise.

### GetEnableNatOk

`func (o *SubnetResourcesDefStatus) GetEnableNatOk() (*bool, bool)`

GetEnableNatOk returns a tuple with the EnableNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNat

`func (o *SubnetResourcesDefStatus) SetEnableNat(v bool)`

SetEnableNat sets EnableNat field to given value.

### HasEnableNat

`func (o *SubnetResourcesDefStatus) HasEnableNat() bool`

HasEnableNat returns a boolean if a field has been set.

### GetAdvancedNetworking

`func (o *SubnetResourcesDefStatus) GetAdvancedNetworking() bool`

GetAdvancedNetworking returns the AdvancedNetworking field if non-nil, zero value otherwise.

### GetAdvancedNetworkingOk

`func (o *SubnetResourcesDefStatus) GetAdvancedNetworkingOk() (*bool, bool)`

GetAdvancedNetworkingOk returns a tuple with the AdvancedNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedNetworking

`func (o *SubnetResourcesDefStatus) SetAdvancedNetworking(v bool)`

SetAdvancedNetworking sets AdvancedNetworking field to given value.

### HasAdvancedNetworking

`func (o *SubnetResourcesDefStatus) HasAdvancedNetworking() bool`

HasAdvancedNetworking returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *SubnetResourcesDefStatus) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *SubnetResourcesDefStatus) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *SubnetResourcesDefStatus) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *SubnetResourcesDefStatus) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetMigrationState

`func (o *SubnetResourcesDefStatus) GetMigrationState() string`

GetMigrationState returns the MigrationState field if non-nil, zero value otherwise.

### GetMigrationStateOk

`func (o *SubnetResourcesDefStatus) GetMigrationStateOk() (*string, bool)`

GetMigrationStateOk returns a tuple with the MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationState

`func (o *SubnetResourcesDefStatus) SetMigrationState(v string)`

SetMigrationState sets MigrationState field to given value.

### HasMigrationState

`func (o *SubnetResourcesDefStatus) HasMigrationState() bool`

HasMigrationState returns a boolean if a field has been set.

### GetAvailabilityZoneReferenceList

`func (o *SubnetResourcesDefStatus) GetAvailabilityZoneReferenceList() []AvailabilityZoneReference`

GetAvailabilityZoneReferenceList returns the AvailabilityZoneReferenceList field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceListOk

`func (o *SubnetResourcesDefStatus) GetAvailabilityZoneReferenceListOk() (*[]AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceListOk returns a tuple with the AvailabilityZoneReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReferenceList

`func (o *SubnetResourcesDefStatus) SetAvailabilityZoneReferenceList(v []AvailabilityZoneReference)`

SetAvailabilityZoneReferenceList sets AvailabilityZoneReferenceList field to given value.

### HasAvailabilityZoneReferenceList

`func (o *SubnetResourcesDefStatus) HasAvailabilityZoneReferenceList() bool`

HasAvailabilityZoneReferenceList returns a boolean if a field has been set.

### GetExternalConnectivityState

`func (o *SubnetResourcesDefStatus) GetExternalConnectivityState() string`

GetExternalConnectivityState returns the ExternalConnectivityState field if non-nil, zero value otherwise.

### GetExternalConnectivityStateOk

`func (o *SubnetResourcesDefStatus) GetExternalConnectivityStateOk() (*string, bool)`

GetExternalConnectivityStateOk returns a tuple with the ExternalConnectivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivityState

`func (o *SubnetResourcesDefStatus) SetExternalConnectivityState(v string)`

SetExternalConnectivityState sets ExternalConnectivityState field to given value.

### HasExternalConnectivityState

`func (o *SubnetResourcesDefStatus) HasExternalConnectivityState() bool`

HasExternalConnectivityState returns a boolean if a field has been set.

### GetVlanId

`func (o *SubnetResourcesDefStatus) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SubnetResourcesDefStatus) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SubnetResourcesDefStatus) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SubnetResourcesDefStatus) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetIsExternal

`func (o *SubnetResourcesDefStatus) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *SubnetResourcesDefStatus) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *SubnetResourcesDefStatus) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *SubnetResourcesDefStatus) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


