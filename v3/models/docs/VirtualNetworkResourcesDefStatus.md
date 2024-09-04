# VirtualNetworkResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonDomainNameServerIpList** | Pointer to [**[]Address**](Address.md) | List of domain name server IPs. | [optional] 
**ExternalSubnetList** | Pointer to [**[]ExternalSubnetDefStatus**](ExternalSubnetDefStatus.md) | List of external subnets attached to this VPC. | [optional] 
**VpnConfig** | Pointer to **string** | Per region providing secure connection from on-prem to Xi (Only supported on Xi)  | [optional] 
**NatIpList** | Pointer to **[]string** | List of IP addresses used for SNAT. | [optional] 
**AvailabilityZoneReferenceList** | Pointer to [**[]AvailabilityZoneReference**](AvailabilityZoneReference.md) | List of availability zones in Xi from which resources are derived (Only supported on Xi)  | [optional] 
**VpcType** | Pointer to **string** | Transit or regular VPC. | [optional] 
**ExternallyRoutablePrefixList** | Pointer to [**[]IpSubnet**](IpSubnet.md) | CIDR blocks from the VPC which can talk externally without performing NAT. These blocks should be between /16 netmask and /28 netmask and cannot overlap across VPCs. They are effective when the VPC connects to a NAT-less external subnet.  | [optional] 

## Methods

### NewVirtualNetworkResourcesDefStatus

`func NewVirtualNetworkResourcesDefStatus() *VirtualNetworkResourcesDefStatus`

NewVirtualNetworkResourcesDefStatus instantiates a new VirtualNetworkResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkResourcesDefStatusWithDefaults

`func NewVirtualNetworkResourcesDefStatusWithDefaults() *VirtualNetworkResourcesDefStatus`

NewVirtualNetworkResourcesDefStatusWithDefaults instantiates a new VirtualNetworkResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonDomainNameServerIpList

`func (o *VirtualNetworkResourcesDefStatus) GetCommonDomainNameServerIpList() []Address`

GetCommonDomainNameServerIpList returns the CommonDomainNameServerIpList field if non-nil, zero value otherwise.

### GetCommonDomainNameServerIpListOk

`func (o *VirtualNetworkResourcesDefStatus) GetCommonDomainNameServerIpListOk() (*[]Address, bool)`

GetCommonDomainNameServerIpListOk returns a tuple with the CommonDomainNameServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonDomainNameServerIpList

`func (o *VirtualNetworkResourcesDefStatus) SetCommonDomainNameServerIpList(v []Address)`

SetCommonDomainNameServerIpList sets CommonDomainNameServerIpList field to given value.

### HasCommonDomainNameServerIpList

`func (o *VirtualNetworkResourcesDefStatus) HasCommonDomainNameServerIpList() bool`

HasCommonDomainNameServerIpList returns a boolean if a field has been set.

### GetExternalSubnetList

`func (o *VirtualNetworkResourcesDefStatus) GetExternalSubnetList() []ExternalSubnetDefStatus`

GetExternalSubnetList returns the ExternalSubnetList field if non-nil, zero value otherwise.

### GetExternalSubnetListOk

`func (o *VirtualNetworkResourcesDefStatus) GetExternalSubnetListOk() (*[]ExternalSubnetDefStatus, bool)`

GetExternalSubnetListOk returns a tuple with the ExternalSubnetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetList

`func (o *VirtualNetworkResourcesDefStatus) SetExternalSubnetList(v []ExternalSubnetDefStatus)`

SetExternalSubnetList sets ExternalSubnetList field to given value.

### HasExternalSubnetList

`func (o *VirtualNetworkResourcesDefStatus) HasExternalSubnetList() bool`

HasExternalSubnetList returns a boolean if a field has been set.

### GetVpnConfig

`func (o *VirtualNetworkResourcesDefStatus) GetVpnConfig() string`

GetVpnConfig returns the VpnConfig field if non-nil, zero value otherwise.

### GetVpnConfigOk

`func (o *VirtualNetworkResourcesDefStatus) GetVpnConfigOk() (*string, bool)`

GetVpnConfigOk returns a tuple with the VpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfig

`func (o *VirtualNetworkResourcesDefStatus) SetVpnConfig(v string)`

SetVpnConfig sets VpnConfig field to given value.

### HasVpnConfig

`func (o *VirtualNetworkResourcesDefStatus) HasVpnConfig() bool`

HasVpnConfig returns a boolean if a field has been set.

### GetNatIpList

`func (o *VirtualNetworkResourcesDefStatus) GetNatIpList() []string`

GetNatIpList returns the NatIpList field if non-nil, zero value otherwise.

### GetNatIpListOk

`func (o *VirtualNetworkResourcesDefStatus) GetNatIpListOk() (*[]string, bool)`

GetNatIpListOk returns a tuple with the NatIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIpList

`func (o *VirtualNetworkResourcesDefStatus) SetNatIpList(v []string)`

SetNatIpList sets NatIpList field to given value.

### HasNatIpList

`func (o *VirtualNetworkResourcesDefStatus) HasNatIpList() bool`

HasNatIpList returns a boolean if a field has been set.

### GetAvailabilityZoneReferenceList

`func (o *VirtualNetworkResourcesDefStatus) GetAvailabilityZoneReferenceList() []AvailabilityZoneReference`

GetAvailabilityZoneReferenceList returns the AvailabilityZoneReferenceList field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceListOk

`func (o *VirtualNetworkResourcesDefStatus) GetAvailabilityZoneReferenceListOk() (*[]AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceListOk returns a tuple with the AvailabilityZoneReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReferenceList

`func (o *VirtualNetworkResourcesDefStatus) SetAvailabilityZoneReferenceList(v []AvailabilityZoneReference)`

SetAvailabilityZoneReferenceList sets AvailabilityZoneReferenceList field to given value.

### HasAvailabilityZoneReferenceList

`func (o *VirtualNetworkResourcesDefStatus) HasAvailabilityZoneReferenceList() bool`

HasAvailabilityZoneReferenceList returns a boolean if a field has been set.

### GetVpcType

`func (o *VirtualNetworkResourcesDefStatus) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *VirtualNetworkResourcesDefStatus) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *VirtualNetworkResourcesDefStatus) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *VirtualNetworkResourcesDefStatus) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.

### GetExternallyRoutablePrefixList

`func (o *VirtualNetworkResourcesDefStatus) GetExternallyRoutablePrefixList() []IpSubnet`

GetExternallyRoutablePrefixList returns the ExternallyRoutablePrefixList field if non-nil, zero value otherwise.

### GetExternallyRoutablePrefixListOk

`func (o *VirtualNetworkResourcesDefStatus) GetExternallyRoutablePrefixListOk() (*[]IpSubnet, bool)`

GetExternallyRoutablePrefixListOk returns a tuple with the ExternallyRoutablePrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyRoutablePrefixList

`func (o *VirtualNetworkResourcesDefStatus) SetExternallyRoutablePrefixList(v []IpSubnet)`

SetExternallyRoutablePrefixList sets ExternallyRoutablePrefixList field to given value.

### HasExternallyRoutablePrefixList

`func (o *VirtualNetworkResourcesDefStatus) HasExternallyRoutablePrefixList() bool`

HasExternallyRoutablePrefixList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


