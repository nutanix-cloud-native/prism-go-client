# VpnGatewayResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**GatewayDeviceVendor** | Pointer to **string** | 3rd Party VPN Vendor. This could be a traditional device vendor (like Cisco ASA, Fortinet etc.), or one of the public cloud providers (eg: AWS).  | [optional] 
**PublicIp** | Pointer to **string** | Public IP address of this gateway. | [optional] 
**GatewayType** | **string** | Whether this is local or remote gateway entity. | 
**Deployment** | Pointer to [**Deployment**](Deployment.md) |  | [optional] 
**InternalRoutingProtocolConfig** | Pointer to [**InternalRoutingProtocolConfig**](InternalRoutingProtocolConfig.md) |  | [optional] 
**EbgpConfig** | Pointer to [**BgpConfig**](BgpConfig.md) |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 

## Methods

### NewVpnGatewayResources

`func NewVpnGatewayResources(gatewayType string, ) *VpnGatewayResources`

NewVpnGatewayResources instantiates a new VpnGatewayResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayResourcesWithDefaults

`func NewVpnGatewayResourcesWithDefaults() *VpnGatewayResources`

NewVpnGatewayResourcesWithDefaults instantiates a new VpnGatewayResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *VpnGatewayResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *VpnGatewayResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *VpnGatewayResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *VpnGatewayResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetGatewayDeviceVendor

`func (o *VpnGatewayResources) GetGatewayDeviceVendor() string`

GetGatewayDeviceVendor returns the GatewayDeviceVendor field if non-nil, zero value otherwise.

### GetGatewayDeviceVendorOk

`func (o *VpnGatewayResources) GetGatewayDeviceVendorOk() (*string, bool)`

GetGatewayDeviceVendorOk returns a tuple with the GatewayDeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDeviceVendor

`func (o *VpnGatewayResources) SetGatewayDeviceVendor(v string)`

SetGatewayDeviceVendor sets GatewayDeviceVendor field to given value.

### HasGatewayDeviceVendor

`func (o *VpnGatewayResources) HasGatewayDeviceVendor() bool`

HasGatewayDeviceVendor returns a boolean if a field has been set.

### GetPublicIp

`func (o *VpnGatewayResources) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VpnGatewayResources) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VpnGatewayResources) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VpnGatewayResources) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetGatewayType

`func (o *VpnGatewayResources) GetGatewayType() string`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *VpnGatewayResources) GetGatewayTypeOk() (*string, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *VpnGatewayResources) SetGatewayType(v string)`

SetGatewayType sets GatewayType field to given value.


### GetDeployment

`func (o *VpnGatewayResources) GetDeployment() Deployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *VpnGatewayResources) GetDeploymentOk() (*Deployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *VpnGatewayResources) SetDeployment(v Deployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *VpnGatewayResources) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetInternalRoutingProtocolConfig

`func (o *VpnGatewayResources) GetInternalRoutingProtocolConfig() InternalRoutingProtocolConfig`

GetInternalRoutingProtocolConfig returns the InternalRoutingProtocolConfig field if non-nil, zero value otherwise.

### GetInternalRoutingProtocolConfigOk

`func (o *VpnGatewayResources) GetInternalRoutingProtocolConfigOk() (*InternalRoutingProtocolConfig, bool)`

GetInternalRoutingProtocolConfigOk returns a tuple with the InternalRoutingProtocolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRoutingProtocolConfig

`func (o *VpnGatewayResources) SetInternalRoutingProtocolConfig(v InternalRoutingProtocolConfig)`

SetInternalRoutingProtocolConfig sets InternalRoutingProtocolConfig field to given value.

### HasInternalRoutingProtocolConfig

`func (o *VpnGatewayResources) HasInternalRoutingProtocolConfig() bool`

HasInternalRoutingProtocolConfig returns a boolean if a field has been set.

### GetEbgpConfig

`func (o *VpnGatewayResources) GetEbgpConfig() BgpConfig`

GetEbgpConfig returns the EbgpConfig field if non-nil, zero value otherwise.

### GetEbgpConfigOk

`func (o *VpnGatewayResources) GetEbgpConfigOk() (*BgpConfig, bool)`

GetEbgpConfigOk returns a tuple with the EbgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpConfig

`func (o *VpnGatewayResources) SetEbgpConfig(v BgpConfig)`

SetEbgpConfig sets EbgpConfig field to given value.

### HasEbgpConfig

`func (o *VpnGatewayResources) HasEbgpConfig() bool`

HasEbgpConfig returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *VpnGatewayResources) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *VpnGatewayResources) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *VpnGatewayResources) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *VpnGatewayResources) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


