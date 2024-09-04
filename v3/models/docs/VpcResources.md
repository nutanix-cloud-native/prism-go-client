# VpcResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonDomainNameServerIpList** | Pointer to [**[]Address**](Address.md) | List of domain name server IPs. | [optional] 
**VpcType** | Pointer to **string** | Transit or regular VPC. | [optional] 
**ExternalSubnetList** | Pointer to [**[]ExternalSubnet**](ExternalSubnet.md) | List of external subnets attached to this VPC. | [optional] 
**ExternallyRoutablePrefixList** | Pointer to [**[]IpSubnet**](IpSubnet.md) | CIDR blocks from the VPC which can talk externally without performing NAT. These blocks should be between /16 netmask and /28 netmask and cannot overlap across VPCs. They are effective when the VPC connects to a NAT-less external subnet.  | [optional] 
**AvailabilityZoneReferenceList** | Pointer to [**[]AvailabilityZoneReference**](AvailabilityZoneReference.md) | List of availability zones in Xi from which resources are derived (Only supported on Xi)  | [optional] 

## Methods

### NewVpcResources

`func NewVpcResources() *VpcResources`

NewVpcResources instantiates a new VpcResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcResourcesWithDefaults

`func NewVpcResourcesWithDefaults() *VpcResources`

NewVpcResourcesWithDefaults instantiates a new VpcResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonDomainNameServerIpList

`func (o *VpcResources) GetCommonDomainNameServerIpList() []Address`

GetCommonDomainNameServerIpList returns the CommonDomainNameServerIpList field if non-nil, zero value otherwise.

### GetCommonDomainNameServerIpListOk

`func (o *VpcResources) GetCommonDomainNameServerIpListOk() (*[]Address, bool)`

GetCommonDomainNameServerIpListOk returns a tuple with the CommonDomainNameServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonDomainNameServerIpList

`func (o *VpcResources) SetCommonDomainNameServerIpList(v []Address)`

SetCommonDomainNameServerIpList sets CommonDomainNameServerIpList field to given value.

### HasCommonDomainNameServerIpList

`func (o *VpcResources) HasCommonDomainNameServerIpList() bool`

HasCommonDomainNameServerIpList returns a boolean if a field has been set.

### GetVpcType

`func (o *VpcResources) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *VpcResources) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *VpcResources) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *VpcResources) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.

### GetExternalSubnetList

`func (o *VpcResources) GetExternalSubnetList() []ExternalSubnet`

GetExternalSubnetList returns the ExternalSubnetList field if non-nil, zero value otherwise.

### GetExternalSubnetListOk

`func (o *VpcResources) GetExternalSubnetListOk() (*[]ExternalSubnet, bool)`

GetExternalSubnetListOk returns a tuple with the ExternalSubnetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetList

`func (o *VpcResources) SetExternalSubnetList(v []ExternalSubnet)`

SetExternalSubnetList sets ExternalSubnetList field to given value.

### HasExternalSubnetList

`func (o *VpcResources) HasExternalSubnetList() bool`

HasExternalSubnetList returns a boolean if a field has been set.

### GetExternallyRoutablePrefixList

`func (o *VpcResources) GetExternallyRoutablePrefixList() []IpSubnet`

GetExternallyRoutablePrefixList returns the ExternallyRoutablePrefixList field if non-nil, zero value otherwise.

### GetExternallyRoutablePrefixListOk

`func (o *VpcResources) GetExternallyRoutablePrefixListOk() (*[]IpSubnet, bool)`

GetExternallyRoutablePrefixListOk returns a tuple with the ExternallyRoutablePrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyRoutablePrefixList

`func (o *VpcResources) SetExternallyRoutablePrefixList(v []IpSubnet)`

SetExternallyRoutablePrefixList sets ExternallyRoutablePrefixList field to given value.

### HasExternallyRoutablePrefixList

`func (o *VpcResources) HasExternallyRoutablePrefixList() bool`

HasExternallyRoutablePrefixList returns a boolean if a field has been set.

### GetAvailabilityZoneReferenceList

`func (o *VpcResources) GetAvailabilityZoneReferenceList() []AvailabilityZoneReference`

GetAvailabilityZoneReferenceList returns the AvailabilityZoneReferenceList field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceListOk

`func (o *VpcResources) GetAvailabilityZoneReferenceListOk() (*[]AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceListOk returns a tuple with the AvailabilityZoneReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReferenceList

`func (o *VpcResources) SetAvailabilityZoneReferenceList(v []AvailabilityZoneReference)`

SetAvailabilityZoneReferenceList sets AvailabilityZoneReferenceList field to given value.

### HasAvailabilityZoneReferenceList

`func (o *VpcResources) HasAvailabilityZoneReferenceList() bool`

HasAvailabilityZoneReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


