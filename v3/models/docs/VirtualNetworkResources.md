# VirtualNetworkResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonDomainNameServerIpList** | Pointer to [**[]Address**](Address.md) | List of domain name server IPs. | [optional] 
**ExternalSubnetList** | Pointer to [**[]ExternalSubnet**](ExternalSubnet.md) | List of external subnets attached to this VPC. | [optional] 
**VpnConfig** | Pointer to **string** | Per region providing secure connection from on-prem to Xi (Only supported on Xi)  | [optional] 
**AvailabilityZoneReferenceList** | Pointer to [**[]AvailabilityZoneReference**](AvailabilityZoneReference.md) | List of availability zones in Xi from which resources are derived (Only supported on Xi)  | [optional] 
**VpcType** | Pointer to **string** | Transit or regular VPC. | [optional] 
**ExternallyRoutablePrefixList** | Pointer to [**[]IpSubnet**](IpSubnet.md) | CIDR blocks from the VPC which can talk externally without performing NAT. These blocks should be between /16 netmask and /28 netmask and cannot overlap across VPCs. They are effective when the VPC connects to a NAT-less external subnet.  | [optional] 

## Methods

### NewVirtualNetworkResources

`func NewVirtualNetworkResources() *VirtualNetworkResources`

NewVirtualNetworkResources instantiates a new VirtualNetworkResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkResourcesWithDefaults

`func NewVirtualNetworkResourcesWithDefaults() *VirtualNetworkResources`

NewVirtualNetworkResourcesWithDefaults instantiates a new VirtualNetworkResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonDomainNameServerIpList

`func (o *VirtualNetworkResources) GetCommonDomainNameServerIpList() []Address`

GetCommonDomainNameServerIpList returns the CommonDomainNameServerIpList field if non-nil, zero value otherwise.

### GetCommonDomainNameServerIpListOk

`func (o *VirtualNetworkResources) GetCommonDomainNameServerIpListOk() (*[]Address, bool)`

GetCommonDomainNameServerIpListOk returns a tuple with the CommonDomainNameServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonDomainNameServerIpList

`func (o *VirtualNetworkResources) SetCommonDomainNameServerIpList(v []Address)`

SetCommonDomainNameServerIpList sets CommonDomainNameServerIpList field to given value.

### HasCommonDomainNameServerIpList

`func (o *VirtualNetworkResources) HasCommonDomainNameServerIpList() bool`

HasCommonDomainNameServerIpList returns a boolean if a field has been set.

### GetExternalSubnetList

`func (o *VirtualNetworkResources) GetExternalSubnetList() []ExternalSubnet`

GetExternalSubnetList returns the ExternalSubnetList field if non-nil, zero value otherwise.

### GetExternalSubnetListOk

`func (o *VirtualNetworkResources) GetExternalSubnetListOk() (*[]ExternalSubnet, bool)`

GetExternalSubnetListOk returns a tuple with the ExternalSubnetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetList

`func (o *VirtualNetworkResources) SetExternalSubnetList(v []ExternalSubnet)`

SetExternalSubnetList sets ExternalSubnetList field to given value.

### HasExternalSubnetList

`func (o *VirtualNetworkResources) HasExternalSubnetList() bool`

HasExternalSubnetList returns a boolean if a field has been set.

### GetVpnConfig

`func (o *VirtualNetworkResources) GetVpnConfig() string`

GetVpnConfig returns the VpnConfig field if non-nil, zero value otherwise.

### GetVpnConfigOk

`func (o *VirtualNetworkResources) GetVpnConfigOk() (*string, bool)`

GetVpnConfigOk returns a tuple with the VpnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConfig

`func (o *VirtualNetworkResources) SetVpnConfig(v string)`

SetVpnConfig sets VpnConfig field to given value.

### HasVpnConfig

`func (o *VirtualNetworkResources) HasVpnConfig() bool`

HasVpnConfig returns a boolean if a field has been set.

### GetAvailabilityZoneReferenceList

`func (o *VirtualNetworkResources) GetAvailabilityZoneReferenceList() []AvailabilityZoneReference`

GetAvailabilityZoneReferenceList returns the AvailabilityZoneReferenceList field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceListOk

`func (o *VirtualNetworkResources) GetAvailabilityZoneReferenceListOk() (*[]AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceListOk returns a tuple with the AvailabilityZoneReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReferenceList

`func (o *VirtualNetworkResources) SetAvailabilityZoneReferenceList(v []AvailabilityZoneReference)`

SetAvailabilityZoneReferenceList sets AvailabilityZoneReferenceList field to given value.

### HasAvailabilityZoneReferenceList

`func (o *VirtualNetworkResources) HasAvailabilityZoneReferenceList() bool`

HasAvailabilityZoneReferenceList returns a boolean if a field has been set.

### GetVpcType

`func (o *VirtualNetworkResources) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *VirtualNetworkResources) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *VirtualNetworkResources) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *VirtualNetworkResources) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.

### GetExternallyRoutablePrefixList

`func (o *VirtualNetworkResources) GetExternallyRoutablePrefixList() []IpSubnet`

GetExternallyRoutablePrefixList returns the ExternallyRoutablePrefixList field if non-nil, zero value otherwise.

### GetExternallyRoutablePrefixListOk

`func (o *VirtualNetworkResources) GetExternallyRoutablePrefixListOk() (*[]IpSubnet, bool)`

GetExternallyRoutablePrefixListOk returns a tuple with the ExternallyRoutablePrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyRoutablePrefixList

`func (o *VirtualNetworkResources) SetExternallyRoutablePrefixList(v []IpSubnet)`

SetExternallyRoutablePrefixList sets ExternallyRoutablePrefixList field to given value.

### HasExternallyRoutablePrefixList

`func (o *VirtualNetworkResources) HasExternallyRoutablePrefixList() bool`

HasExternallyRoutablePrefixList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


