# DeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallLbRoute** | Pointer to **bool** | Boolean flag indicating user opt-in for installing Xi LB route in on-prem PC and PE CVMs provided on-prem PC, PE and VPN VM are in the same subnet  | [optional] 
**VmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**VcenterDeploymentDetails** | Pointer to [**VcenterDeploymentDetails**](VcenterDeploymentDetails.md) |  | [optional] 
**IpPrefixLength** | Pointer to **int32** | IP prefix length of the subnet that the gateway VM is on. | [optional] 
**StaticIp** | Pointer to **string** | Static IP address of the VPN gateway VM. | [optional] 
**DefaultGatewayIp** | Pointer to **string** | Default gateway IP address. | [optional] 
**ImageSourceUrl** | Pointer to **string** | The software image installed on the gateway appliance. | [optional] 
**PeClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**SupportedSoftwareVersion** | Pointer to **string** | The supported gateway appliance version. | [optional] 
**InstalledSoftwareVersion** | Pointer to **string** | The software version installed on the gateway appliance. | [optional] 
**VlanId** | Pointer to **int32** | The on-prem VLAN to deploy the VPN gateway on. | [optional] 

## Methods

### NewDeploymentStatus

`func NewDeploymentStatus() *DeploymentStatus`

NewDeploymentStatus instantiates a new DeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStatusWithDefaults

`func NewDeploymentStatusWithDefaults() *DeploymentStatus`

NewDeploymentStatusWithDefaults instantiates a new DeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallLbRoute

`func (o *DeploymentStatus) GetInstallLbRoute() bool`

GetInstallLbRoute returns the InstallLbRoute field if non-nil, zero value otherwise.

### GetInstallLbRouteOk

`func (o *DeploymentStatus) GetInstallLbRouteOk() (*bool, bool)`

GetInstallLbRouteOk returns a tuple with the InstallLbRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallLbRoute

`func (o *DeploymentStatus) SetInstallLbRoute(v bool)`

SetInstallLbRoute sets InstallLbRoute field to given value.

### HasInstallLbRoute

`func (o *DeploymentStatus) HasInstallLbRoute() bool`

HasInstallLbRoute returns a boolean if a field has been set.

### GetVmReference

`func (o *DeploymentStatus) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *DeploymentStatus) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *DeploymentStatus) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.

### HasVmReference

`func (o *DeploymentStatus) HasVmReference() bool`

HasVmReference returns a boolean if a field has been set.

### GetVcenterDeploymentDetails

`func (o *DeploymentStatus) GetVcenterDeploymentDetails() VcenterDeploymentDetails`

GetVcenterDeploymentDetails returns the VcenterDeploymentDetails field if non-nil, zero value otherwise.

### GetVcenterDeploymentDetailsOk

`func (o *DeploymentStatus) GetVcenterDeploymentDetailsOk() (*VcenterDeploymentDetails, bool)`

GetVcenterDeploymentDetailsOk returns a tuple with the VcenterDeploymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDeploymentDetails

`func (o *DeploymentStatus) SetVcenterDeploymentDetails(v VcenterDeploymentDetails)`

SetVcenterDeploymentDetails sets VcenterDeploymentDetails field to given value.

### HasVcenterDeploymentDetails

`func (o *DeploymentStatus) HasVcenterDeploymentDetails() bool`

HasVcenterDeploymentDetails returns a boolean if a field has been set.

### GetIpPrefixLength

`func (o *DeploymentStatus) GetIpPrefixLength() int32`

GetIpPrefixLength returns the IpPrefixLength field if non-nil, zero value otherwise.

### GetIpPrefixLengthOk

`func (o *DeploymentStatus) GetIpPrefixLengthOk() (*int32, bool)`

GetIpPrefixLengthOk returns a tuple with the IpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixLength

`func (o *DeploymentStatus) SetIpPrefixLength(v int32)`

SetIpPrefixLength sets IpPrefixLength field to given value.

### HasIpPrefixLength

`func (o *DeploymentStatus) HasIpPrefixLength() bool`

HasIpPrefixLength returns a boolean if a field has been set.

### GetStaticIp

`func (o *DeploymentStatus) GetStaticIp() string`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *DeploymentStatus) GetStaticIpOk() (*string, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *DeploymentStatus) SetStaticIp(v string)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *DeploymentStatus) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### GetDefaultGatewayIp

`func (o *DeploymentStatus) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *DeploymentStatus) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *DeploymentStatus) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *DeploymentStatus) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### GetImageSourceUrl

`func (o *DeploymentStatus) GetImageSourceUrl() string`

GetImageSourceUrl returns the ImageSourceUrl field if non-nil, zero value otherwise.

### GetImageSourceUrlOk

`func (o *DeploymentStatus) GetImageSourceUrlOk() (*string, bool)`

GetImageSourceUrlOk returns a tuple with the ImageSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSourceUrl

`func (o *DeploymentStatus) SetImageSourceUrl(v string)`

SetImageSourceUrl sets ImageSourceUrl field to given value.

### HasImageSourceUrl

`func (o *DeploymentStatus) HasImageSourceUrl() bool`

HasImageSourceUrl returns a boolean if a field has been set.

### GetPeClusterReference

`func (o *DeploymentStatus) GetPeClusterReference() ClusterReference`

GetPeClusterReference returns the PeClusterReference field if non-nil, zero value otherwise.

### GetPeClusterReferenceOk

`func (o *DeploymentStatus) GetPeClusterReferenceOk() (*ClusterReference, bool)`

GetPeClusterReferenceOk returns a tuple with the PeClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeClusterReference

`func (o *DeploymentStatus) SetPeClusterReference(v ClusterReference)`

SetPeClusterReference sets PeClusterReference field to given value.

### HasPeClusterReference

`func (o *DeploymentStatus) HasPeClusterReference() bool`

HasPeClusterReference returns a boolean if a field has been set.

### GetSubnetReference

`func (o *DeploymentStatus) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *DeploymentStatus) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *DeploymentStatus) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *DeploymentStatus) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetSupportedSoftwareVersion

`func (o *DeploymentStatus) GetSupportedSoftwareVersion() string`

GetSupportedSoftwareVersion returns the SupportedSoftwareVersion field if non-nil, zero value otherwise.

### GetSupportedSoftwareVersionOk

`func (o *DeploymentStatus) GetSupportedSoftwareVersionOk() (*string, bool)`

GetSupportedSoftwareVersionOk returns a tuple with the SupportedSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSoftwareVersion

`func (o *DeploymentStatus) SetSupportedSoftwareVersion(v string)`

SetSupportedSoftwareVersion sets SupportedSoftwareVersion field to given value.

### HasSupportedSoftwareVersion

`func (o *DeploymentStatus) HasSupportedSoftwareVersion() bool`

HasSupportedSoftwareVersion returns a boolean if a field has been set.

### GetInstalledSoftwareVersion

`func (o *DeploymentStatus) GetInstalledSoftwareVersion() string`

GetInstalledSoftwareVersion returns the InstalledSoftwareVersion field if non-nil, zero value otherwise.

### GetInstalledSoftwareVersionOk

`func (o *DeploymentStatus) GetInstalledSoftwareVersionOk() (*string, bool)`

GetInstalledSoftwareVersionOk returns a tuple with the InstalledSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSoftwareVersion

`func (o *DeploymentStatus) SetInstalledSoftwareVersion(v string)`

SetInstalledSoftwareVersion sets InstalledSoftwareVersion field to given value.

### HasInstalledSoftwareVersion

`func (o *DeploymentStatus) HasInstalledSoftwareVersion() bool`

HasInstalledSoftwareVersion returns a boolean if a field has been set.

### GetVlanId

`func (o *DeploymentStatus) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DeploymentStatus) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DeploymentStatus) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DeploymentStatus) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


