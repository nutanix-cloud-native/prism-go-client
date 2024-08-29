# VpnGatewayResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**OperationalStatus** | Pointer to [**VpnComponentStatus**](VpnComponentStatus.md) |  | [optional] 
**GatewayDeviceVendor** | Pointer to **string** | 3rd Party VPN Vendor. This could be a traditional device vendor (like Cisco ASA, Fortinet etc.) or one of the public cloud providers (eg: AWS).  | [optional] 
**PublicIp** | Pointer to **string** | Public IP address of this gateway. | [optional] 
**GatewayType** | Pointer to **string** | Whether this is local or remote gateway entity. | [optional] 
**Deployment** | Pointer to [**DeploymentStatus**](DeploymentStatus.md) |  | [optional] 
**InternalRoutingProtocolConfig** | Pointer to [**InternalRoutingProtocolConfigStatus**](InternalRoutingProtocolConfigStatus.md) |  | [optional] 
**EbgpConfig** | Pointer to [**BgpConfigStatus**](BgpConfigStatus.md) |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 

## Methods

### NewVpnGatewayResourcesDefStatus

`func NewVpnGatewayResourcesDefStatus() *VpnGatewayResourcesDefStatus`

NewVpnGatewayResourcesDefStatus instantiates a new VpnGatewayResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayResourcesDefStatusWithDefaults

`func NewVpnGatewayResourcesDefStatusWithDefaults() *VpnGatewayResourcesDefStatus`

NewVpnGatewayResourcesDefStatusWithDefaults instantiates a new VpnGatewayResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcReference

`func (o *VpnGatewayResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *VpnGatewayResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *VpnGatewayResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *VpnGatewayResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *VpnGatewayResourcesDefStatus) GetOperationalStatus() VpnComponentStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *VpnGatewayResourcesDefStatus) GetOperationalStatusOk() (*VpnComponentStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *VpnGatewayResourcesDefStatus) SetOperationalStatus(v VpnComponentStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *VpnGatewayResourcesDefStatus) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetGatewayDeviceVendor

`func (o *VpnGatewayResourcesDefStatus) GetGatewayDeviceVendor() string`

GetGatewayDeviceVendor returns the GatewayDeviceVendor field if non-nil, zero value otherwise.

### GetGatewayDeviceVendorOk

`func (o *VpnGatewayResourcesDefStatus) GetGatewayDeviceVendorOk() (*string, bool)`

GetGatewayDeviceVendorOk returns a tuple with the GatewayDeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDeviceVendor

`func (o *VpnGatewayResourcesDefStatus) SetGatewayDeviceVendor(v string)`

SetGatewayDeviceVendor sets GatewayDeviceVendor field to given value.

### HasGatewayDeviceVendor

`func (o *VpnGatewayResourcesDefStatus) HasGatewayDeviceVendor() bool`

HasGatewayDeviceVendor returns a boolean if a field has been set.

### GetPublicIp

`func (o *VpnGatewayResourcesDefStatus) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *VpnGatewayResourcesDefStatus) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *VpnGatewayResourcesDefStatus) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *VpnGatewayResourcesDefStatus) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetGatewayType

`func (o *VpnGatewayResourcesDefStatus) GetGatewayType() string`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *VpnGatewayResourcesDefStatus) GetGatewayTypeOk() (*string, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *VpnGatewayResourcesDefStatus) SetGatewayType(v string)`

SetGatewayType sets GatewayType field to given value.

### HasGatewayType

`func (o *VpnGatewayResourcesDefStatus) HasGatewayType() bool`

HasGatewayType returns a boolean if a field has been set.

### GetDeployment

`func (o *VpnGatewayResourcesDefStatus) GetDeployment() DeploymentStatus`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *VpnGatewayResourcesDefStatus) GetDeploymentOk() (*DeploymentStatus, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *VpnGatewayResourcesDefStatus) SetDeployment(v DeploymentStatus)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *VpnGatewayResourcesDefStatus) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetInternalRoutingProtocolConfig

`func (o *VpnGatewayResourcesDefStatus) GetInternalRoutingProtocolConfig() InternalRoutingProtocolConfigStatus`

GetInternalRoutingProtocolConfig returns the InternalRoutingProtocolConfig field if non-nil, zero value otherwise.

### GetInternalRoutingProtocolConfigOk

`func (o *VpnGatewayResourcesDefStatus) GetInternalRoutingProtocolConfigOk() (*InternalRoutingProtocolConfigStatus, bool)`

GetInternalRoutingProtocolConfigOk returns a tuple with the InternalRoutingProtocolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRoutingProtocolConfig

`func (o *VpnGatewayResourcesDefStatus) SetInternalRoutingProtocolConfig(v InternalRoutingProtocolConfigStatus)`

SetInternalRoutingProtocolConfig sets InternalRoutingProtocolConfig field to given value.

### HasInternalRoutingProtocolConfig

`func (o *VpnGatewayResourcesDefStatus) HasInternalRoutingProtocolConfig() bool`

HasInternalRoutingProtocolConfig returns a boolean if a field has been set.

### GetEbgpConfig

`func (o *VpnGatewayResourcesDefStatus) GetEbgpConfig() BgpConfigStatus`

GetEbgpConfig returns the EbgpConfig field if non-nil, zero value otherwise.

### GetEbgpConfigOk

`func (o *VpnGatewayResourcesDefStatus) GetEbgpConfigOk() (*BgpConfigStatus, bool)`

GetEbgpConfigOk returns a tuple with the EbgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpConfig

`func (o *VpnGatewayResourcesDefStatus) SetEbgpConfig(v BgpConfigStatus)`

SetEbgpConfig sets EbgpConfig field to given value.

### HasEbgpConfig

`func (o *VpnGatewayResourcesDefStatus) HasEbgpConfig() bool`

HasEbgpConfig returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *VpnGatewayResourcesDefStatus) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *VpnGatewayResourcesDefStatus) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *VpnGatewayResourcesDefStatus) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *VpnGatewayResourcesDefStatus) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


