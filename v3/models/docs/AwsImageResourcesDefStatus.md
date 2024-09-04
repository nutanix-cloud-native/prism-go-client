# AwsImageResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootDeviceType** | Pointer to **string** | Root device type (e.g. ebs, instance-store) | [optional] 
**KernelId** | Pointer to **string** | Kernal AWS ID of the image. | [optional] 
**Hypervisor** | Pointer to **string** | The supported hypervisor. | [optional] 
**Id** | Pointer to **string** | The AWS ID of the AMI. | [optional] 
**Platform** | Pointer to **string** | Platform of the image (e.g. Windows) | [optional] 
**Architecture** | Pointer to **string** | Architecture of the AWS image (e.g. i386, x86_64) | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**VirtualizationType** | Pointer to **string** | Type of virtualization supported. | [optional] 
**RootDeviceName** | Pointer to **string** | Root device name (e.g. /dev/sda1, /dev/sda2) | [optional] 

## Methods

### NewAwsImageResourcesDefStatus

`func NewAwsImageResourcesDefStatus() *AwsImageResourcesDefStatus`

NewAwsImageResourcesDefStatus instantiates a new AwsImageResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsImageResourcesDefStatusWithDefaults

`func NewAwsImageResourcesDefStatusWithDefaults() *AwsImageResourcesDefStatus`

NewAwsImageResourcesDefStatusWithDefaults instantiates a new AwsImageResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootDeviceType

`func (o *AwsImageResourcesDefStatus) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *AwsImageResourcesDefStatus) GetRootDeviceTypeOk() (*string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceType

`func (o *AwsImageResourcesDefStatus) SetRootDeviceType(v string)`

SetRootDeviceType sets RootDeviceType field to given value.

### HasRootDeviceType

`func (o *AwsImageResourcesDefStatus) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### GetKernelId

`func (o *AwsImageResourcesDefStatus) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *AwsImageResourcesDefStatus) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *AwsImageResourcesDefStatus) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *AwsImageResourcesDefStatus) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### GetHypervisor

`func (o *AwsImageResourcesDefStatus) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *AwsImageResourcesDefStatus) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *AwsImageResourcesDefStatus) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *AwsImageResourcesDefStatus) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetId

`func (o *AwsImageResourcesDefStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsImageResourcesDefStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsImageResourcesDefStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsImageResourcesDefStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *AwsImageResourcesDefStatus) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AwsImageResourcesDefStatus) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AwsImageResourcesDefStatus) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AwsImageResourcesDefStatus) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetArchitecture

`func (o *AwsImageResourcesDefStatus) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *AwsImageResourcesDefStatus) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *AwsImageResourcesDefStatus) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *AwsImageResourcesDefStatus) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetTagList

`func (o *AwsImageResourcesDefStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsImageResourcesDefStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsImageResourcesDefStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsImageResourcesDefStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetVirtualizationType

`func (o *AwsImageResourcesDefStatus) GetVirtualizationType() string`

GetVirtualizationType returns the VirtualizationType field if non-nil, zero value otherwise.

### GetVirtualizationTypeOk

`func (o *AwsImageResourcesDefStatus) GetVirtualizationTypeOk() (*string, bool)`

GetVirtualizationTypeOk returns a tuple with the VirtualizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualizationType

`func (o *AwsImageResourcesDefStatus) SetVirtualizationType(v string)`

SetVirtualizationType sets VirtualizationType field to given value.

### HasVirtualizationType

`func (o *AwsImageResourcesDefStatus) HasVirtualizationType() bool`

HasVirtualizationType returns a boolean if a field has been set.

### GetRootDeviceName

`func (o *AwsImageResourcesDefStatus) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *AwsImageResourcesDefStatus) GetRootDeviceNameOk() (*string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceName

`func (o *AwsImageResourcesDefStatus) SetRootDeviceName(v string)`

SetRootDeviceName sets RootDeviceName field to given value.

### HasRootDeviceName

`func (o *AwsImageResourcesDefStatus) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


