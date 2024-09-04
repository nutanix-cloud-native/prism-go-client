# VmCloneOverrideSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumThreadsPerCore** | Pointer to **int32** | Number of logical threads per core. | [optional] 
**NicList** | Pointer to [**[]VmNic**](VmNic.md) | If specified, the complete list of NICs attached to the cloned VM.  | [optional] 
**Name** | Pointer to **string** | VM Name. | [optional] 
**NumVcpusPerSocket** | Pointer to **int32** | Number of vCPUs per socket. | [optional] 
**NumSockets** | Pointer to **int32** | Number of vCPU sockets. | [optional] 
**MemorySizeMib** | Pointer to **int32** | Memory size in MiB. | [optional] 
**BootConfig** | Pointer to [**VmBootConfig**](VmBootConfig.md) |  | [optional] 
**GuestCustomization** | Pointer to [**GuestCustomization**](GuestCustomization.md) |  | [optional] 

## Methods

### NewVmCloneOverrideSpec

`func NewVmCloneOverrideSpec() *VmCloneOverrideSpec`

NewVmCloneOverrideSpec instantiates a new VmCloneOverrideSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCloneOverrideSpecWithDefaults

`func NewVmCloneOverrideSpecWithDefaults() *VmCloneOverrideSpec`

NewVmCloneOverrideSpecWithDefaults instantiates a new VmCloneOverrideSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumThreadsPerCore

`func (o *VmCloneOverrideSpec) GetNumThreadsPerCore() int32`

GetNumThreadsPerCore returns the NumThreadsPerCore field if non-nil, zero value otherwise.

### GetNumThreadsPerCoreOk

`func (o *VmCloneOverrideSpec) GetNumThreadsPerCoreOk() (*int32, bool)`

GetNumThreadsPerCoreOk returns a tuple with the NumThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreadsPerCore

`func (o *VmCloneOverrideSpec) SetNumThreadsPerCore(v int32)`

SetNumThreadsPerCore sets NumThreadsPerCore field to given value.

### HasNumThreadsPerCore

`func (o *VmCloneOverrideSpec) HasNumThreadsPerCore() bool`

HasNumThreadsPerCore returns a boolean if a field has been set.

### GetNicList

`func (o *VmCloneOverrideSpec) GetNicList() []VmNic`

GetNicList returns the NicList field if non-nil, zero value otherwise.

### GetNicListOk

`func (o *VmCloneOverrideSpec) GetNicListOk() (*[]VmNic, bool)`

GetNicListOk returns a tuple with the NicList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicList

`func (o *VmCloneOverrideSpec) SetNicList(v []VmNic)`

SetNicList sets NicList field to given value.

### HasNicList

`func (o *VmCloneOverrideSpec) HasNicList() bool`

HasNicList returns a boolean if a field has been set.

### GetName

`func (o *VmCloneOverrideSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmCloneOverrideSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmCloneOverrideSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmCloneOverrideSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumVcpusPerSocket

`func (o *VmCloneOverrideSpec) GetNumVcpusPerSocket() int32`

GetNumVcpusPerSocket returns the NumVcpusPerSocket field if non-nil, zero value otherwise.

### GetNumVcpusPerSocketOk

`func (o *VmCloneOverrideSpec) GetNumVcpusPerSocketOk() (*int32, bool)`

GetNumVcpusPerSocketOk returns a tuple with the NumVcpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVcpusPerSocket

`func (o *VmCloneOverrideSpec) SetNumVcpusPerSocket(v int32)`

SetNumVcpusPerSocket sets NumVcpusPerSocket field to given value.

### HasNumVcpusPerSocket

`func (o *VmCloneOverrideSpec) HasNumVcpusPerSocket() bool`

HasNumVcpusPerSocket returns a boolean if a field has been set.

### GetNumSockets

`func (o *VmCloneOverrideSpec) GetNumSockets() int32`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *VmCloneOverrideSpec) GetNumSocketsOk() (*int32, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *VmCloneOverrideSpec) SetNumSockets(v int32)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *VmCloneOverrideSpec) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.

### GetMemorySizeMib

`func (o *VmCloneOverrideSpec) GetMemorySizeMib() int32`

GetMemorySizeMib returns the MemorySizeMib field if non-nil, zero value otherwise.

### GetMemorySizeMibOk

`func (o *VmCloneOverrideSpec) GetMemorySizeMibOk() (*int32, bool)`

GetMemorySizeMibOk returns a tuple with the MemorySizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMib

`func (o *VmCloneOverrideSpec) SetMemorySizeMib(v int32)`

SetMemorySizeMib sets MemorySizeMib field to given value.

### HasMemorySizeMib

`func (o *VmCloneOverrideSpec) HasMemorySizeMib() bool`

HasMemorySizeMib returns a boolean if a field has been set.

### GetBootConfig

`func (o *VmCloneOverrideSpec) GetBootConfig() VmBootConfig`

GetBootConfig returns the BootConfig field if non-nil, zero value otherwise.

### GetBootConfigOk

`func (o *VmCloneOverrideSpec) GetBootConfigOk() (*VmBootConfig, bool)`

GetBootConfigOk returns a tuple with the BootConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootConfig

`func (o *VmCloneOverrideSpec) SetBootConfig(v VmBootConfig)`

SetBootConfig sets BootConfig field to given value.

### HasBootConfig

`func (o *VmCloneOverrideSpec) HasBootConfig() bool`

HasBootConfig returns a boolean if a field has been set.

### GetGuestCustomization

`func (o *VmCloneOverrideSpec) GetGuestCustomization() GuestCustomization`

GetGuestCustomization returns the GuestCustomization field if non-nil, zero value otherwise.

### GetGuestCustomizationOk

`func (o *VmCloneOverrideSpec) GetGuestCustomizationOk() (*GuestCustomization, bool)`

GetGuestCustomizationOk returns a tuple with the GuestCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomization

`func (o *VmCloneOverrideSpec) SetGuestCustomization(v GuestCustomization)`

SetGuestCustomization sets GuestCustomization field to given value.

### HasGuestCustomization

`func (o *VmCloneOverrideSpec) HasGuestCustomization() bool`

HasGuestCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


