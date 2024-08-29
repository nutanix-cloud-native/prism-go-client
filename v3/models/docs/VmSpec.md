# VmSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | **string** | The UUID which is used to uniquely identify this VM.  | 
**NumCoresPerVcpu** | Pointer to **int32** | Number of cores per vCPU. | [optional] 
**VmDiskList** | Pointer to [**[]DiskSpec**](DiskSpec.md) | List of associated VM virtual disks. | [optional] 
**VmRecoveryPointReference** | Pointer to [**VmRecoveryPointReference**](VmRecoveryPointReference.md) |  | [optional] 
**NumVcpus** | Pointer to **int32** | Number of vCPUs needed. | [optional] 
**MemoryBytes** | Pointer to **int32** | Amount of memory needed in bytes. | [optional] 
**VmCapabilityList** | Pointer to **[]string** | List of capabilities VM is configured with. | [optional] 
**VirtualHardwareVersion** | Pointer to **int32** | Virtual hardware version of the VM. | [optional] 
**SourceVmUuid** | Pointer to **string** | Source VM UUID of which this VM is clone of.  | [optional] 

## Methods

### NewVmSpec

`func NewVmSpec(uUID string, ) *VmSpec`

NewVmSpec instantiates a new VmSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSpecWithDefaults

`func NewVmSpecWithDefaults() *VmSpec`

NewVmSpecWithDefaults instantiates a new VmSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *VmSpec) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmSpec) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmSpec) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetNumCoresPerVcpu

`func (o *VmSpec) GetNumCoresPerVcpu() int32`

GetNumCoresPerVcpu returns the NumCoresPerVcpu field if non-nil, zero value otherwise.

### GetNumCoresPerVcpuOk

`func (o *VmSpec) GetNumCoresPerVcpuOk() (*int32, bool)`

GetNumCoresPerVcpuOk returns a tuple with the NumCoresPerVcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCoresPerVcpu

`func (o *VmSpec) SetNumCoresPerVcpu(v int32)`

SetNumCoresPerVcpu sets NumCoresPerVcpu field to given value.

### HasNumCoresPerVcpu

`func (o *VmSpec) HasNumCoresPerVcpu() bool`

HasNumCoresPerVcpu returns a boolean if a field has been set.

### GetVmDiskList

`func (o *VmSpec) GetVmDiskList() []DiskSpec`

GetVmDiskList returns the VmDiskList field if non-nil, zero value otherwise.

### GetVmDiskListOk

`func (o *VmSpec) GetVmDiskListOk() (*[]DiskSpec, bool)`

GetVmDiskListOk returns a tuple with the VmDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskList

`func (o *VmSpec) SetVmDiskList(v []DiskSpec)`

SetVmDiskList sets VmDiskList field to given value.

### HasVmDiskList

`func (o *VmSpec) HasVmDiskList() bool`

HasVmDiskList returns a boolean if a field has been set.

### GetVmRecoveryPointReference

`func (o *VmSpec) GetVmRecoveryPointReference() VmRecoveryPointReference`

GetVmRecoveryPointReference returns the VmRecoveryPointReference field if non-nil, zero value otherwise.

### GetVmRecoveryPointReferenceOk

`func (o *VmSpec) GetVmRecoveryPointReferenceOk() (*VmRecoveryPointReference, bool)`

GetVmRecoveryPointReferenceOk returns a tuple with the VmRecoveryPointReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointReference

`func (o *VmSpec) SetVmRecoveryPointReference(v VmRecoveryPointReference)`

SetVmRecoveryPointReference sets VmRecoveryPointReference field to given value.

### HasVmRecoveryPointReference

`func (o *VmSpec) HasVmRecoveryPointReference() bool`

HasVmRecoveryPointReference returns a boolean if a field has been set.

### GetNumVcpus

`func (o *VmSpec) GetNumVcpus() int32`

GetNumVcpus returns the NumVcpus field if non-nil, zero value otherwise.

### GetNumVcpusOk

`func (o *VmSpec) GetNumVcpusOk() (*int32, bool)`

GetNumVcpusOk returns a tuple with the NumVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVcpus

`func (o *VmSpec) SetNumVcpus(v int32)`

SetNumVcpus sets NumVcpus field to given value.

### HasNumVcpus

`func (o *VmSpec) HasNumVcpus() bool`

HasNumVcpus returns a boolean if a field has been set.

### GetMemoryBytes

`func (o *VmSpec) GetMemoryBytes() int32`

GetMemoryBytes returns the MemoryBytes field if non-nil, zero value otherwise.

### GetMemoryBytesOk

`func (o *VmSpec) GetMemoryBytesOk() (*int32, bool)`

GetMemoryBytesOk returns a tuple with the MemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBytes

`func (o *VmSpec) SetMemoryBytes(v int32)`

SetMemoryBytes sets MemoryBytes field to given value.

### HasMemoryBytes

`func (o *VmSpec) HasMemoryBytes() bool`

HasMemoryBytes returns a boolean if a field has been set.

### GetVmCapabilityList

`func (o *VmSpec) GetVmCapabilityList() []string`

GetVmCapabilityList returns the VmCapabilityList field if non-nil, zero value otherwise.

### GetVmCapabilityListOk

`func (o *VmSpec) GetVmCapabilityListOk() (*[]string, bool)`

GetVmCapabilityListOk returns a tuple with the VmCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCapabilityList

`func (o *VmSpec) SetVmCapabilityList(v []string)`

SetVmCapabilityList sets VmCapabilityList field to given value.

### HasVmCapabilityList

`func (o *VmSpec) HasVmCapabilityList() bool`

HasVmCapabilityList returns a boolean if a field has been set.

### GetVirtualHardwareVersion

`func (o *VmSpec) GetVirtualHardwareVersion() int32`

GetVirtualHardwareVersion returns the VirtualHardwareVersion field if non-nil, zero value otherwise.

### GetVirtualHardwareVersionOk

`func (o *VmSpec) GetVirtualHardwareVersionOk() (*int32, bool)`

GetVirtualHardwareVersionOk returns a tuple with the VirtualHardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHardwareVersion

`func (o *VmSpec) SetVirtualHardwareVersion(v int32)`

SetVirtualHardwareVersion sets VirtualHardwareVersion field to given value.

### HasVirtualHardwareVersion

`func (o *VmSpec) HasVirtualHardwareVersion() bool`

HasVirtualHardwareVersion returns a boolean if a field has been set.

### GetSourceVmUuid

`func (o *VmSpec) GetSourceVmUuid() string`

GetSourceVmUuid returns the SourceVmUuid field if non-nil, zero value otherwise.

### GetSourceVmUuidOk

`func (o *VmSpec) GetSourceVmUuidOk() (*string, bool)`

GetSourceVmUuidOk returns a tuple with the SourceVmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVmUuid

`func (o *VmSpec) SetSourceVmUuid(v string)`

SetSourceVmUuid sets SourceVmUuid field to given value.

### HasSourceVmUuid

`func (o *VmSpec) HasSourceVmUuid() bool`

HasSourceVmUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


