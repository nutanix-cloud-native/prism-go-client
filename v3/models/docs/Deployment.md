# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallLbRoute** | Pointer to **bool** | Boolean flag indicating user opt-in for installing Xi LB route in on-prem PC and PE CVMs provided on-prem PC, PE and VPN VM are in the same subnet  | [optional] 
**VcenterDeploymentDetails** | Pointer to [**VcenterDeploymentDetails**](VcenterDeploymentDetails.md) |  | [optional] 
**IpPrefixLength** | Pointer to **int32** | IP prefix length of the subnet that the gateway VM is on. | [optional] 
**StaticIp** | Pointer to **string** | Static IP address of the VPN gateway VM. | [optional] 
**DefaultGatewayIp** | Pointer to **string** | Default gateway IP address. | [optional] 
**ImageSourceUrl** | Pointer to **string** | The software image to install on the gateway appliance. If set, \&quot;installed_software_version\&quot; must be omitted.  | [optional] 
**PeClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**InstalledSoftwareVersion** | Pointer to **string** | The software version to install on the gateway appliance. If set, \&quot;image_source_url\&quot; must be omitted.  | [optional] 
**VlanId** | Pointer to **int32** | The on-prem VLAN to deploy the VPN gateway on. This is not needed if the subnet_reference is provided.  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallLbRoute

`func (o *Deployment) GetInstallLbRoute() bool`

GetInstallLbRoute returns the InstallLbRoute field if non-nil, zero value otherwise.

### GetInstallLbRouteOk

`func (o *Deployment) GetInstallLbRouteOk() (*bool, bool)`

GetInstallLbRouteOk returns a tuple with the InstallLbRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallLbRoute

`func (o *Deployment) SetInstallLbRoute(v bool)`

SetInstallLbRoute sets InstallLbRoute field to given value.

### HasInstallLbRoute

`func (o *Deployment) HasInstallLbRoute() bool`

HasInstallLbRoute returns a boolean if a field has been set.

### GetVcenterDeploymentDetails

`func (o *Deployment) GetVcenterDeploymentDetails() VcenterDeploymentDetails`

GetVcenterDeploymentDetails returns the VcenterDeploymentDetails field if non-nil, zero value otherwise.

### GetVcenterDeploymentDetailsOk

`func (o *Deployment) GetVcenterDeploymentDetailsOk() (*VcenterDeploymentDetails, bool)`

GetVcenterDeploymentDetailsOk returns a tuple with the VcenterDeploymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDeploymentDetails

`func (o *Deployment) SetVcenterDeploymentDetails(v VcenterDeploymentDetails)`

SetVcenterDeploymentDetails sets VcenterDeploymentDetails field to given value.

### HasVcenterDeploymentDetails

`func (o *Deployment) HasVcenterDeploymentDetails() bool`

HasVcenterDeploymentDetails returns a boolean if a field has been set.

### GetIpPrefixLength

`func (o *Deployment) GetIpPrefixLength() int32`

GetIpPrefixLength returns the IpPrefixLength field if non-nil, zero value otherwise.

### GetIpPrefixLengthOk

`func (o *Deployment) GetIpPrefixLengthOk() (*int32, bool)`

GetIpPrefixLengthOk returns a tuple with the IpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPrefixLength

`func (o *Deployment) SetIpPrefixLength(v int32)`

SetIpPrefixLength sets IpPrefixLength field to given value.

### HasIpPrefixLength

`func (o *Deployment) HasIpPrefixLength() bool`

HasIpPrefixLength returns a boolean if a field has been set.

### GetStaticIp

`func (o *Deployment) GetStaticIp() string`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *Deployment) GetStaticIpOk() (*string, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *Deployment) SetStaticIp(v string)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *Deployment) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### GetDefaultGatewayIp

`func (o *Deployment) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *Deployment) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *Deployment) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *Deployment) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### GetImageSourceUrl

`func (o *Deployment) GetImageSourceUrl() string`

GetImageSourceUrl returns the ImageSourceUrl field if non-nil, zero value otherwise.

### GetImageSourceUrlOk

`func (o *Deployment) GetImageSourceUrlOk() (*string, bool)`

GetImageSourceUrlOk returns a tuple with the ImageSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSourceUrl

`func (o *Deployment) SetImageSourceUrl(v string)`

SetImageSourceUrl sets ImageSourceUrl field to given value.

### HasImageSourceUrl

`func (o *Deployment) HasImageSourceUrl() bool`

HasImageSourceUrl returns a boolean if a field has been set.

### GetPeClusterReference

`func (o *Deployment) GetPeClusterReference() ClusterReference`

GetPeClusterReference returns the PeClusterReference field if non-nil, zero value otherwise.

### GetPeClusterReferenceOk

`func (o *Deployment) GetPeClusterReferenceOk() (*ClusterReference, bool)`

GetPeClusterReferenceOk returns a tuple with the PeClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeClusterReference

`func (o *Deployment) SetPeClusterReference(v ClusterReference)`

SetPeClusterReference sets PeClusterReference field to given value.

### HasPeClusterReference

`func (o *Deployment) HasPeClusterReference() bool`

HasPeClusterReference returns a boolean if a field has been set.

### GetSubnetReference

`func (o *Deployment) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *Deployment) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *Deployment) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *Deployment) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetInstalledSoftwareVersion

`func (o *Deployment) GetInstalledSoftwareVersion() string`

GetInstalledSoftwareVersion returns the InstalledSoftwareVersion field if non-nil, zero value otherwise.

### GetInstalledSoftwareVersionOk

`func (o *Deployment) GetInstalledSoftwareVersionOk() (*string, bool)`

GetInstalledSoftwareVersionOk returns a tuple with the InstalledSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSoftwareVersion

`func (o *Deployment) SetInstalledSoftwareVersion(v string)`

SetInstalledSoftwareVersion sets InstalledSoftwareVersion field to given value.

### HasInstalledSoftwareVersion

`func (o *Deployment) HasInstalledSoftwareVersion() bool`

HasInstalledSoftwareVersion returns a boolean if a field has been set.

### GetVlanId

`func (o *Deployment) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Deployment) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Deployment) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Deployment) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


