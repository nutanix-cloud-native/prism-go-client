# FloatingIpResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**FloatingIp** | Pointer to **string** | The Floating IP associated with the vnic. | [optional] 
**VmNicReference** | Pointer to [**VmNicReference**](VmNicReference.md) |  | [optional] 
**PrivateIp** | Pointer to **string** | Private IP with which the floating IP is associated. | [optional] 

## Methods

### NewFloatingIpResourcesDefStatus

`func NewFloatingIpResourcesDefStatus() *FloatingIpResourcesDefStatus`

NewFloatingIpResourcesDefStatus instantiates a new FloatingIpResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpResourcesDefStatusWithDefaults

`func NewFloatingIpResourcesDefStatusWithDefaults() *FloatingIpResourcesDefStatus`

NewFloatingIpResourcesDefStatusWithDefaults instantiates a new FloatingIpResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSubnetReference

`func (o *FloatingIpResourcesDefStatus) GetExternalSubnetReference() SubnetReference`

GetExternalSubnetReference returns the ExternalSubnetReference field if non-nil, zero value otherwise.

### GetExternalSubnetReferenceOk

`func (o *FloatingIpResourcesDefStatus) GetExternalSubnetReferenceOk() (*SubnetReference, bool)`

GetExternalSubnetReferenceOk returns a tuple with the ExternalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetReference

`func (o *FloatingIpResourcesDefStatus) SetExternalSubnetReference(v SubnetReference)`

SetExternalSubnetReference sets ExternalSubnetReference field to given value.

### HasExternalSubnetReference

`func (o *FloatingIpResourcesDefStatus) HasExternalSubnetReference() bool`

HasExternalSubnetReference returns a boolean if a field has been set.

### GetVpcReference

`func (o *FloatingIpResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *FloatingIpResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *FloatingIpResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *FloatingIpResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetFloatingIp

`func (o *FloatingIpResourcesDefStatus) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *FloatingIpResourcesDefStatus) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *FloatingIpResourcesDefStatus) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *FloatingIpResourcesDefStatus) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetVmNicReference

`func (o *FloatingIpResourcesDefStatus) GetVmNicReference() VmNicReference`

GetVmNicReference returns the VmNicReference field if non-nil, zero value otherwise.

### GetVmNicReferenceOk

`func (o *FloatingIpResourcesDefStatus) GetVmNicReferenceOk() (*VmNicReference, bool)`

GetVmNicReferenceOk returns a tuple with the VmNicReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNicReference

`func (o *FloatingIpResourcesDefStatus) SetVmNicReference(v VmNicReference)`

SetVmNicReference sets VmNicReference field to given value.

### HasVmNicReference

`func (o *FloatingIpResourcesDefStatus) HasVmNicReference() bool`

HasVmNicReference returns a boolean if a field has been set.

### GetPrivateIp

`func (o *FloatingIpResourcesDefStatus) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *FloatingIpResourcesDefStatus) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *FloatingIpResourcesDefStatus) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *FloatingIpResourcesDefStatus) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


