# VmGpuOutputStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrameBufferSizeMib** | Pointer to **int64** | GPU frame buffer size in MiB. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the GPU. | [optional] 
**UUID** | Pointer to **string** | UUID of the GPU. | [optional] 
**Name** | Pointer to **string** | Name of the GPU resource. | [optional] 
**PciAddress** | Pointer to **string** | GPU {segment:bus:device:function} (sbdf) address if assigned.  | [optional] 
**Fraction** | Pointer to **int32** | Fraction of the physical GPU assigned. | [optional] 
**Mode** | Pointer to **string** | The mode of this GPU | [optional] 
**NumVirtualDisplayHeads** | Pointer to **int32** | Number of supported virtual display heads. | [optional] 
**GuestDriverVersion** | Pointer to **string** | Last determined guest driver version. | [optional] 
**MaxInstancesPerVm** | Pointer to **int64** | Maximum number of vGPUs of a GPU type assignable to a VM. | [optional] 
**DeviceId** | Pointer to **int32** | The device ID of the GPU. | [optional] 

## Methods

### NewVmGpuOutputStatus

`func NewVmGpuOutputStatus() *VmGpuOutputStatus`

NewVmGpuOutputStatus instantiates a new VmGpuOutputStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGpuOutputStatusWithDefaults

`func NewVmGpuOutputStatusWithDefaults() *VmGpuOutputStatus`

NewVmGpuOutputStatusWithDefaults instantiates a new VmGpuOutputStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrameBufferSizeMib

`func (o *VmGpuOutputStatus) GetFrameBufferSizeMib() int64`

GetFrameBufferSizeMib returns the FrameBufferSizeMib field if non-nil, zero value otherwise.

### GetFrameBufferSizeMibOk

`func (o *VmGpuOutputStatus) GetFrameBufferSizeMibOk() (*int64, bool)`

GetFrameBufferSizeMibOk returns a tuple with the FrameBufferSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeMib

`func (o *VmGpuOutputStatus) SetFrameBufferSizeMib(v int64)`

SetFrameBufferSizeMib sets FrameBufferSizeMib field to given value.

### HasFrameBufferSizeMib

`func (o *VmGpuOutputStatus) HasFrameBufferSizeMib() bool`

HasFrameBufferSizeMib returns a boolean if a field has been set.

### GetVendor

`func (o *VmGpuOutputStatus) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VmGpuOutputStatus) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VmGpuOutputStatus) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VmGpuOutputStatus) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetUUID

`func (o *VmGpuOutputStatus) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmGpuOutputStatus) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmGpuOutputStatus) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmGpuOutputStatus) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *VmGpuOutputStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmGpuOutputStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmGpuOutputStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmGpuOutputStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPciAddress

`func (o *VmGpuOutputStatus) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *VmGpuOutputStatus) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *VmGpuOutputStatus) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *VmGpuOutputStatus) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetFraction

`func (o *VmGpuOutputStatus) GetFraction() int32`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *VmGpuOutputStatus) GetFractionOk() (*int32, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *VmGpuOutputStatus) SetFraction(v int32)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *VmGpuOutputStatus) HasFraction() bool`

HasFraction returns a boolean if a field has been set.

### GetMode

`func (o *VmGpuOutputStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VmGpuOutputStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VmGpuOutputStatus) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VmGpuOutputStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumVirtualDisplayHeads

`func (o *VmGpuOutputStatus) GetNumVirtualDisplayHeads() int32`

GetNumVirtualDisplayHeads returns the NumVirtualDisplayHeads field if non-nil, zero value otherwise.

### GetNumVirtualDisplayHeadsOk

`func (o *VmGpuOutputStatus) GetNumVirtualDisplayHeadsOk() (*int32, bool)`

GetNumVirtualDisplayHeadsOk returns a tuple with the NumVirtualDisplayHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVirtualDisplayHeads

`func (o *VmGpuOutputStatus) SetNumVirtualDisplayHeads(v int32)`

SetNumVirtualDisplayHeads sets NumVirtualDisplayHeads field to given value.

### HasNumVirtualDisplayHeads

`func (o *VmGpuOutputStatus) HasNumVirtualDisplayHeads() bool`

HasNumVirtualDisplayHeads returns a boolean if a field has been set.

### GetGuestDriverVersion

`func (o *VmGpuOutputStatus) GetGuestDriverVersion() string`

GetGuestDriverVersion returns the GuestDriverVersion field if non-nil, zero value otherwise.

### GetGuestDriverVersionOk

`func (o *VmGpuOutputStatus) GetGuestDriverVersionOk() (*string, bool)`

GetGuestDriverVersionOk returns a tuple with the GuestDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestDriverVersion

`func (o *VmGpuOutputStatus) SetGuestDriverVersion(v string)`

SetGuestDriverVersion sets GuestDriverVersion field to given value.

### HasGuestDriverVersion

`func (o *VmGpuOutputStatus) HasGuestDriverVersion() bool`

HasGuestDriverVersion returns a boolean if a field has been set.

### GetMaxInstancesPerVm

`func (o *VmGpuOutputStatus) GetMaxInstancesPerVm() int64`

GetMaxInstancesPerVm returns the MaxInstancesPerVm field if non-nil, zero value otherwise.

### GetMaxInstancesPerVmOk

`func (o *VmGpuOutputStatus) GetMaxInstancesPerVmOk() (*int64, bool)`

GetMaxInstancesPerVmOk returns a tuple with the MaxInstancesPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstancesPerVm

`func (o *VmGpuOutputStatus) SetMaxInstancesPerVm(v int64)`

SetMaxInstancesPerVm sets MaxInstancesPerVm field to given value.

### HasMaxInstancesPerVm

`func (o *VmGpuOutputStatus) HasMaxInstancesPerVm() bool`

HasMaxInstancesPerVm returns a boolean if a field has been set.

### GetDeviceId

`func (o *VmGpuOutputStatus) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *VmGpuOutputStatus) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *VmGpuOutputStatus) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *VmGpuOutputStatus) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


