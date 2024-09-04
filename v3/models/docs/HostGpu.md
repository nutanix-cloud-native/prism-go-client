# HostGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current status of the physical GPU. | [optional] 
**Assignable** | Pointer to **bool** | Whether this vGPU instance can be allocated to a VM. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the GPU. | [optional] 
**Name** | Pointer to **string** | Name of the host GPU. | [optional] 
**Index** | Pointer to **int32** | The index of the vGPU within physical GPU resource. | [optional] 
**LicenseList** | Pointer to **[]string** | List of license types associated with this GPU. | [optional] 
**NumaNode** | Pointer to **int32** | NUMA node this GPU belongs to. | [optional] 
**MaxResolution** | Pointer to **string** | Maximum resolution per display head. | [optional] 
**NumVgpusAllocated** | Pointer to **int32** | The number of vGPU instances allocated for this physical GPU resource.  | [optional] 
**ConsumerReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**PciAddress** | Pointer to **string** | GPU {segment:bus:device:function} (sbdf) address. | [optional] 
**Fraction** | Pointer to **int32** | Fraction of the physical GPU. | [optional] 
**Mode** | Pointer to **string** | The type of this GPU. | [optional] 
**NumVirtualDisplayHeads** | Pointer to **int32** | Number of supported virtual display heads. | [optional] 
**GuestDriverVersion** | Pointer to **string** | Last determined guest driver version. | [optional] 
**FrameBufferSizeMib** | Pointer to **int64** | GPU frame buffer size in MiB. | [optional] 
**DeviceId** | Pointer to **int32** | The device ID of the GPU. | [optional] 
**MaxInstancesPerVm** | Pointer to **int64** | Maximum number of vGPUs of a GPU type assignable to a VM. | [optional] 
**UUID** | Pointer to **string** | UUID of the GPU. | [optional] 

## Methods

### NewHostGpu

`func NewHostGpu() *HostGpu`

NewHostGpu instantiates a new HostGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGpuWithDefaults

`func NewHostGpuWithDefaults() *HostGpu`

NewHostGpuWithDefaults instantiates a new HostGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostGpu) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostGpu) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostGpu) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostGpu) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignable

`func (o *HostGpu) GetAssignable() bool`

GetAssignable returns the Assignable field if non-nil, zero value otherwise.

### GetAssignableOk

`func (o *HostGpu) GetAssignableOk() (*bool, bool)`

GetAssignableOk returns a tuple with the Assignable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignable

`func (o *HostGpu) SetAssignable(v bool)`

SetAssignable sets Assignable field to given value.

### HasAssignable

`func (o *HostGpu) HasAssignable() bool`

HasAssignable returns a boolean if a field has been set.

### GetVendor

`func (o *HostGpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HostGpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HostGpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *HostGpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetName

`func (o *HostGpu) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostGpu) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostGpu) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostGpu) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIndex

`func (o *HostGpu) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *HostGpu) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *HostGpu) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *HostGpu) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLicenseList

`func (o *HostGpu) GetLicenseList() []string`

GetLicenseList returns the LicenseList field if non-nil, zero value otherwise.

### GetLicenseListOk

`func (o *HostGpu) GetLicenseListOk() (*[]string, bool)`

GetLicenseListOk returns a tuple with the LicenseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseList

`func (o *HostGpu) SetLicenseList(v []string)`

SetLicenseList sets LicenseList field to given value.

### HasLicenseList

`func (o *HostGpu) HasLicenseList() bool`

HasLicenseList returns a boolean if a field has been set.

### GetNumaNode

`func (o *HostGpu) GetNumaNode() int32`

GetNumaNode returns the NumaNode field if non-nil, zero value otherwise.

### GetNumaNodeOk

`func (o *HostGpu) GetNumaNodeOk() (*int32, bool)`

GetNumaNodeOk returns a tuple with the NumaNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumaNode

`func (o *HostGpu) SetNumaNode(v int32)`

SetNumaNode sets NumaNode field to given value.

### HasNumaNode

`func (o *HostGpu) HasNumaNode() bool`

HasNumaNode returns a boolean if a field has been set.

### GetMaxResolution

`func (o *HostGpu) GetMaxResolution() string`

GetMaxResolution returns the MaxResolution field if non-nil, zero value otherwise.

### GetMaxResolutionOk

`func (o *HostGpu) GetMaxResolutionOk() (*string, bool)`

GetMaxResolutionOk returns a tuple with the MaxResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolution

`func (o *HostGpu) SetMaxResolution(v string)`

SetMaxResolution sets MaxResolution field to given value.

### HasMaxResolution

`func (o *HostGpu) HasMaxResolution() bool`

HasMaxResolution returns a boolean if a field has been set.

### GetNumVgpusAllocated

`func (o *HostGpu) GetNumVgpusAllocated() int32`

GetNumVgpusAllocated returns the NumVgpusAllocated field if non-nil, zero value otherwise.

### GetNumVgpusAllocatedOk

`func (o *HostGpu) GetNumVgpusAllocatedOk() (*int32, bool)`

GetNumVgpusAllocatedOk returns a tuple with the NumVgpusAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVgpusAllocated

`func (o *HostGpu) SetNumVgpusAllocated(v int32)`

SetNumVgpusAllocated sets NumVgpusAllocated field to given value.

### HasNumVgpusAllocated

`func (o *HostGpu) HasNumVgpusAllocated() bool`

HasNumVgpusAllocated returns a boolean if a field has been set.

### GetConsumerReference

`func (o *HostGpu) GetConsumerReference() Reference`

GetConsumerReference returns the ConsumerReference field if non-nil, zero value otherwise.

### GetConsumerReferenceOk

`func (o *HostGpu) GetConsumerReferenceOk() (*Reference, bool)`

GetConsumerReferenceOk returns a tuple with the ConsumerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerReference

`func (o *HostGpu) SetConsumerReference(v Reference)`

SetConsumerReference sets ConsumerReference field to given value.

### HasConsumerReference

`func (o *HostGpu) HasConsumerReference() bool`

HasConsumerReference returns a boolean if a field has been set.

### GetPciAddress

`func (o *HostGpu) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *HostGpu) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *HostGpu) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *HostGpu) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetFraction

`func (o *HostGpu) GetFraction() int32`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *HostGpu) GetFractionOk() (*int32, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *HostGpu) SetFraction(v int32)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *HostGpu) HasFraction() bool`

HasFraction returns a boolean if a field has been set.

### GetMode

`func (o *HostGpu) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HostGpu) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HostGpu) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HostGpu) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNumVirtualDisplayHeads

`func (o *HostGpu) GetNumVirtualDisplayHeads() int32`

GetNumVirtualDisplayHeads returns the NumVirtualDisplayHeads field if non-nil, zero value otherwise.

### GetNumVirtualDisplayHeadsOk

`func (o *HostGpu) GetNumVirtualDisplayHeadsOk() (*int32, bool)`

GetNumVirtualDisplayHeadsOk returns a tuple with the NumVirtualDisplayHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVirtualDisplayHeads

`func (o *HostGpu) SetNumVirtualDisplayHeads(v int32)`

SetNumVirtualDisplayHeads sets NumVirtualDisplayHeads field to given value.

### HasNumVirtualDisplayHeads

`func (o *HostGpu) HasNumVirtualDisplayHeads() bool`

HasNumVirtualDisplayHeads returns a boolean if a field has been set.

### GetGuestDriverVersion

`func (o *HostGpu) GetGuestDriverVersion() string`

GetGuestDriverVersion returns the GuestDriverVersion field if non-nil, zero value otherwise.

### GetGuestDriverVersionOk

`func (o *HostGpu) GetGuestDriverVersionOk() (*string, bool)`

GetGuestDriverVersionOk returns a tuple with the GuestDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestDriverVersion

`func (o *HostGpu) SetGuestDriverVersion(v string)`

SetGuestDriverVersion sets GuestDriverVersion field to given value.

### HasGuestDriverVersion

`func (o *HostGpu) HasGuestDriverVersion() bool`

HasGuestDriverVersion returns a boolean if a field has been set.

### GetFrameBufferSizeMib

`func (o *HostGpu) GetFrameBufferSizeMib() int64`

GetFrameBufferSizeMib returns the FrameBufferSizeMib field if non-nil, zero value otherwise.

### GetFrameBufferSizeMibOk

`func (o *HostGpu) GetFrameBufferSizeMibOk() (*int64, bool)`

GetFrameBufferSizeMibOk returns a tuple with the FrameBufferSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameBufferSizeMib

`func (o *HostGpu) SetFrameBufferSizeMib(v int64)`

SetFrameBufferSizeMib sets FrameBufferSizeMib field to given value.

### HasFrameBufferSizeMib

`func (o *HostGpu) HasFrameBufferSizeMib() bool`

HasFrameBufferSizeMib returns a boolean if a field has been set.

### GetDeviceId

`func (o *HostGpu) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HostGpu) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HostGpu) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HostGpu) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetMaxInstancesPerVm

`func (o *HostGpu) GetMaxInstancesPerVm() int64`

GetMaxInstancesPerVm returns the MaxInstancesPerVm field if non-nil, zero value otherwise.

### GetMaxInstancesPerVmOk

`func (o *HostGpu) GetMaxInstancesPerVmOk() (*int64, bool)`

GetMaxInstancesPerVmOk returns a tuple with the MaxInstancesPerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstancesPerVm

`func (o *HostGpu) SetMaxInstancesPerVm(v int64)`

SetMaxInstancesPerVm sets MaxInstancesPerVm field to given value.

### HasMaxInstancesPerVm

`func (o *HostGpu) HasMaxInstancesPerVm() bool`

HasMaxInstancesPerVm returns a boolean if a field has been set.

### GetUUID

`func (o *HostGpu) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *HostGpu) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *HostGpu) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *HostGpu) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


