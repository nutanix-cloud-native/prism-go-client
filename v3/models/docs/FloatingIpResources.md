# FloatingIpResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**PrivateIp** | Pointer to **string** | Private IP with which the floating IP is associated. | [optional] 
**VmNicReference** | Pointer to [**VmNicReference**](VmNicReference.md) |  | [optional] 
**FloatingIp** | Pointer to **string** | The Floating IP Address. | [optional] 

## Methods

### NewFloatingIpResources

`func NewFloatingIpResources() *FloatingIpResources`

NewFloatingIpResources instantiates a new FloatingIpResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpResourcesWithDefaults

`func NewFloatingIpResourcesWithDefaults() *FloatingIpResources`

NewFloatingIpResourcesWithDefaults instantiates a new FloatingIpResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSubnetReference

`func (o *FloatingIpResources) GetExternalSubnetReference() SubnetReference`

GetExternalSubnetReference returns the ExternalSubnetReference field if non-nil, zero value otherwise.

### GetExternalSubnetReferenceOk

`func (o *FloatingIpResources) GetExternalSubnetReferenceOk() (*SubnetReference, bool)`

GetExternalSubnetReferenceOk returns a tuple with the ExternalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetReference

`func (o *FloatingIpResources) SetExternalSubnetReference(v SubnetReference)`

SetExternalSubnetReference sets ExternalSubnetReference field to given value.

### HasExternalSubnetReference

`func (o *FloatingIpResources) HasExternalSubnetReference() bool`

HasExternalSubnetReference returns a boolean if a field has been set.

### GetVpcReference

`func (o *FloatingIpResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *FloatingIpResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *FloatingIpResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *FloatingIpResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetPrivateIp

`func (o *FloatingIpResources) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *FloatingIpResources) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *FloatingIpResources) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *FloatingIpResources) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetVmNicReference

`func (o *FloatingIpResources) GetVmNicReference() VmNicReference`

GetVmNicReference returns the VmNicReference field if non-nil, zero value otherwise.

### GetVmNicReferenceOk

`func (o *FloatingIpResources) GetVmNicReferenceOk() (*VmNicReference, bool)`

GetVmNicReferenceOk returns a tuple with the VmNicReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNicReference

`func (o *FloatingIpResources) SetVmNicReference(v VmNicReference)`

SetVmNicReference sets VmNicReference field to given value.

### HasVmNicReference

`func (o *FloatingIpResources) HasVmNicReference() bool`

HasVmNicReference returns a boolean if a field has been set.

### GetFloatingIp

`func (o *FloatingIpResources) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *FloatingIpResources) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *FloatingIpResources) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *FloatingIpResources) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


