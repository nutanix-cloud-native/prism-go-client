# VmResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumThreadsPerCore** | Pointer to **int32** | Number of logical threads per core. | [optional] 
**IsAgentVm** | Pointer to **bool** | Indicates whether the VM is an agent VM. | [optional] 
**ProtectionPolicyState** | Pointer to [**ProtectionPolicyState**](ProtectionPolicyState.md) |  | [optional] 
**MemorySizeMib** | Pointer to **int32** | Memory size in MiB. | [optional] 
**BootConfig** | Pointer to [**VmBootConfig**](VmBootConfig.md) |  | [optional] 
**DiskList** | Pointer to [**[]VmDiskOutputStatus**](VmDiskOutputStatus.md) | Disks attached to the VM. | [optional] 
**SerialPortList** | Pointer to [**[]SerialPort**](SerialPort.md) | Serial ports configured on the VM. | [optional] 
**GenerationUuid** | Pointer to **string** | Indicates the Generation UUID of the VM. | [optional] 
**BiosUuid** | Pointer to **string** | Indicates the BIOS UUID of the VM. | [optional] 
**VgaConsoleEnabled** | Pointer to **bool** | Indicates whether VGA console has been enabled or not. | [optional] 
**PowerState** | Pointer to **string** | Current power state of the VM. | [optional] 
**RecoveryPlanStateList** | Pointer to [**[]RecoveryPlanPolicyState**](RecoveryPlanPolicyState.md) | Status of the Recovery Plans associated with the VM. | [optional] 
**EffectiveStorageConfig** | Pointer to [**EffectiveStorageConfig**](EffectiveStorageConfig.md) |  | [optional] 
**NumVcpusPerSocket** | Pointer to **int32** | Number of vCPUs per socket. | [optional] 
**NumSockets** | Pointer to **int32** | Number of vCPU sockets. | [optional] 
**HardwareVirtualizationEnabled** | Pointer to **bool** | Indicates whether hardware assisted virtualization should be enabled for the Guest OS. Once enabled, the Guest OS has the ability to deploy a nested hypervisor.  | [optional] 
**StorageConfig** | Pointer to [**VmStorageConfigStatus**](VmStorageConfigStatus.md) |  | [optional] 
**ProtectionType** | Pointer to **string** | The type of protection applied on a VM. PD_PROTECTED indicates a VM protected using Prism Element. RULE_PROTECTED indicates a VM protected using Prism Central.  | [optional] 
**GpuList** | Pointer to [**[]VmGpuOutputStatus**](VmGpuOutputStatus.md) | GPUs attached to the VM. | [optional] 
**MachineType** | Pointer to **string** | Machine type for the VM. Machine type Q35 is required for secure boot and does not support IDE disks.  | [optional] 
**HardwareClockTimezone** | Pointer to **string** | VM&#39;s hardware clock timezone in IANA TZDB format (America/Los_Angeles).  | [optional] 
**GuestCustomization** | Pointer to [**GuestCustomizationStatus**](GuestCustomizationStatus.md) |  | [optional] 
**PowerStateMechanism** | Pointer to [**VmPowerStateMechanism**](VmPowerStateMechanism.md) |  | [optional] 
**IsVcpuHardPinned** | Pointer to **bool** | Indicates whether the vCPUs should be hard pinned to specific pCPUs. | [optional] 
**MemoryOvercommitEnabled** | Pointer to **bool** | Indicates whether memory overcommit feature should be enabled for the VM. If enabled, parts of the VM&#39;s memory may reside outside of the hypervisor physical memory. When enabled, it should be expected that the VM may suffer performance degradation.  | [optional] 
**VnumaConfig** | Pointer to [**VmVnumaConfig**](VmVnumaConfig.md) |  | [optional] 
**NicList** | Pointer to [**[]VmNicOutputStatus**](VmNicOutputStatus.md) | NICs attached to the VM. | [optional] 
**HostReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**GuestOsId** | Pointer to **string** | String that identifies the OS running inside of the guest. Reserved for use by the system. Do not set or modify. | [optional] 
**GuestTools** | Pointer to [**GuestToolsStatus**](GuestToolsStatus.md) |  | [optional] 
**GpuConsoleEnabled** | Pointer to **bool** | Indicates whether vGPU console is enabled or not. | [optional] 
**VtpmConfig** | Pointer to [**VmVtpmStatus**](VmVtpmStatus.md) |  | [optional] 
**EnableCpuPassthrough** | Pointer to **bool** | Indicates whether to passthrough the host&#39;s CPU features to the guest. Enabling this will disable live migration of the VM.  | [optional] 
**ParentReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the hypervisor the VM is hosted on.  | [optional] 
**DisableBranding** | Pointer to **bool** | Indicates whether to remove AHV branding from VM firmware tables.  | [optional] 

## Methods

### NewVmResourcesDefStatus

`func NewVmResourcesDefStatus() *VmResourcesDefStatus`

NewVmResourcesDefStatus instantiates a new VmResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmResourcesDefStatusWithDefaults

`func NewVmResourcesDefStatusWithDefaults() *VmResourcesDefStatus`

NewVmResourcesDefStatusWithDefaults instantiates a new VmResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumThreadsPerCore

`func (o *VmResourcesDefStatus) GetNumThreadsPerCore() int32`

GetNumThreadsPerCore returns the NumThreadsPerCore field if non-nil, zero value otherwise.

### GetNumThreadsPerCoreOk

`func (o *VmResourcesDefStatus) GetNumThreadsPerCoreOk() (*int32, bool)`

GetNumThreadsPerCoreOk returns a tuple with the NumThreadsPerCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreadsPerCore

`func (o *VmResourcesDefStatus) SetNumThreadsPerCore(v int32)`

SetNumThreadsPerCore sets NumThreadsPerCore field to given value.

### HasNumThreadsPerCore

`func (o *VmResourcesDefStatus) HasNumThreadsPerCore() bool`

HasNumThreadsPerCore returns a boolean if a field has been set.

### GetIsAgentVm

`func (o *VmResourcesDefStatus) GetIsAgentVm() bool`

GetIsAgentVm returns the IsAgentVm field if non-nil, zero value otherwise.

### GetIsAgentVmOk

`func (o *VmResourcesDefStatus) GetIsAgentVmOk() (*bool, bool)`

GetIsAgentVmOk returns a tuple with the IsAgentVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAgentVm

`func (o *VmResourcesDefStatus) SetIsAgentVm(v bool)`

SetIsAgentVm sets IsAgentVm field to given value.

### HasIsAgentVm

`func (o *VmResourcesDefStatus) HasIsAgentVm() bool`

HasIsAgentVm returns a boolean if a field has been set.

### GetProtectionPolicyState

`func (o *VmResourcesDefStatus) GetProtectionPolicyState() ProtectionPolicyState`

GetProtectionPolicyState returns the ProtectionPolicyState field if non-nil, zero value otherwise.

### GetProtectionPolicyStateOk

`func (o *VmResourcesDefStatus) GetProtectionPolicyStateOk() (*ProtectionPolicyState, bool)`

GetProtectionPolicyStateOk returns a tuple with the ProtectionPolicyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicyState

`func (o *VmResourcesDefStatus) SetProtectionPolicyState(v ProtectionPolicyState)`

SetProtectionPolicyState sets ProtectionPolicyState field to given value.

### HasProtectionPolicyState

`func (o *VmResourcesDefStatus) HasProtectionPolicyState() bool`

HasProtectionPolicyState returns a boolean if a field has been set.

### GetMemorySizeMib

`func (o *VmResourcesDefStatus) GetMemorySizeMib() int32`

GetMemorySizeMib returns the MemorySizeMib field if non-nil, zero value otherwise.

### GetMemorySizeMibOk

`func (o *VmResourcesDefStatus) GetMemorySizeMibOk() (*int32, bool)`

GetMemorySizeMibOk returns a tuple with the MemorySizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMib

`func (o *VmResourcesDefStatus) SetMemorySizeMib(v int32)`

SetMemorySizeMib sets MemorySizeMib field to given value.

### HasMemorySizeMib

`func (o *VmResourcesDefStatus) HasMemorySizeMib() bool`

HasMemorySizeMib returns a boolean if a field has been set.

### GetBootConfig

`func (o *VmResourcesDefStatus) GetBootConfig() VmBootConfig`

GetBootConfig returns the BootConfig field if non-nil, zero value otherwise.

### GetBootConfigOk

`func (o *VmResourcesDefStatus) GetBootConfigOk() (*VmBootConfig, bool)`

GetBootConfigOk returns a tuple with the BootConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootConfig

`func (o *VmResourcesDefStatus) SetBootConfig(v VmBootConfig)`

SetBootConfig sets BootConfig field to given value.

### HasBootConfig

`func (o *VmResourcesDefStatus) HasBootConfig() bool`

HasBootConfig returns a boolean if a field has been set.

### GetDiskList

`func (o *VmResourcesDefStatus) GetDiskList() []VmDiskOutputStatus`

GetDiskList returns the DiskList field if non-nil, zero value otherwise.

### GetDiskListOk

`func (o *VmResourcesDefStatus) GetDiskListOk() (*[]VmDiskOutputStatus, bool)`

GetDiskListOk returns a tuple with the DiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskList

`func (o *VmResourcesDefStatus) SetDiskList(v []VmDiskOutputStatus)`

SetDiskList sets DiskList field to given value.

### HasDiskList

`func (o *VmResourcesDefStatus) HasDiskList() bool`

HasDiskList returns a boolean if a field has been set.

### GetSerialPortList

`func (o *VmResourcesDefStatus) GetSerialPortList() []SerialPort`

GetSerialPortList returns the SerialPortList field if non-nil, zero value otherwise.

### GetSerialPortListOk

`func (o *VmResourcesDefStatus) GetSerialPortListOk() (*[]SerialPort, bool)`

GetSerialPortListOk returns a tuple with the SerialPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialPortList

`func (o *VmResourcesDefStatus) SetSerialPortList(v []SerialPort)`

SetSerialPortList sets SerialPortList field to given value.

### HasSerialPortList

`func (o *VmResourcesDefStatus) HasSerialPortList() bool`

HasSerialPortList returns a boolean if a field has been set.

### GetGenerationUuid

`func (o *VmResourcesDefStatus) GetGenerationUuid() string`

GetGenerationUuid returns the GenerationUuid field if non-nil, zero value otherwise.

### GetGenerationUuidOk

`func (o *VmResourcesDefStatus) GetGenerationUuidOk() (*string, bool)`

GetGenerationUuidOk returns a tuple with the GenerationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationUuid

`func (o *VmResourcesDefStatus) SetGenerationUuid(v string)`

SetGenerationUuid sets GenerationUuid field to given value.

### HasGenerationUuid

`func (o *VmResourcesDefStatus) HasGenerationUuid() bool`

HasGenerationUuid returns a boolean if a field has been set.

### GetBiosUuid

`func (o *VmResourcesDefStatus) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *VmResourcesDefStatus) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *VmResourcesDefStatus) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.

### HasBiosUuid

`func (o *VmResourcesDefStatus) HasBiosUuid() bool`

HasBiosUuid returns a boolean if a field has been set.

### GetVgaConsoleEnabled

`func (o *VmResourcesDefStatus) GetVgaConsoleEnabled() bool`

GetVgaConsoleEnabled returns the VgaConsoleEnabled field if non-nil, zero value otherwise.

### GetVgaConsoleEnabledOk

`func (o *VmResourcesDefStatus) GetVgaConsoleEnabledOk() (*bool, bool)`

GetVgaConsoleEnabledOk returns a tuple with the VgaConsoleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgaConsoleEnabled

`func (o *VmResourcesDefStatus) SetVgaConsoleEnabled(v bool)`

SetVgaConsoleEnabled sets VgaConsoleEnabled field to given value.

### HasVgaConsoleEnabled

`func (o *VmResourcesDefStatus) HasVgaConsoleEnabled() bool`

HasVgaConsoleEnabled returns a boolean if a field has been set.

### GetPowerState

`func (o *VmResourcesDefStatus) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VmResourcesDefStatus) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VmResourcesDefStatus) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VmResourcesDefStatus) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetRecoveryPlanStateList

`func (o *VmResourcesDefStatus) GetRecoveryPlanStateList() []RecoveryPlanPolicyState`

GetRecoveryPlanStateList returns the RecoveryPlanStateList field if non-nil, zero value otherwise.

### GetRecoveryPlanStateListOk

`func (o *VmResourcesDefStatus) GetRecoveryPlanStateListOk() (*[]RecoveryPlanPolicyState, bool)`

GetRecoveryPlanStateListOk returns a tuple with the RecoveryPlanStateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPlanStateList

`func (o *VmResourcesDefStatus) SetRecoveryPlanStateList(v []RecoveryPlanPolicyState)`

SetRecoveryPlanStateList sets RecoveryPlanStateList field to given value.

### HasRecoveryPlanStateList

`func (o *VmResourcesDefStatus) HasRecoveryPlanStateList() bool`

HasRecoveryPlanStateList returns a boolean if a field has been set.

### GetEffectiveStorageConfig

`func (o *VmResourcesDefStatus) GetEffectiveStorageConfig() EffectiveStorageConfig`

GetEffectiveStorageConfig returns the EffectiveStorageConfig field if non-nil, zero value otherwise.

### GetEffectiveStorageConfigOk

`func (o *VmResourcesDefStatus) GetEffectiveStorageConfigOk() (*EffectiveStorageConfig, bool)`

GetEffectiveStorageConfigOk returns a tuple with the EffectiveStorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStorageConfig

`func (o *VmResourcesDefStatus) SetEffectiveStorageConfig(v EffectiveStorageConfig)`

SetEffectiveStorageConfig sets EffectiveStorageConfig field to given value.

### HasEffectiveStorageConfig

`func (o *VmResourcesDefStatus) HasEffectiveStorageConfig() bool`

HasEffectiveStorageConfig returns a boolean if a field has been set.

### GetNumVcpusPerSocket

`func (o *VmResourcesDefStatus) GetNumVcpusPerSocket() int32`

GetNumVcpusPerSocket returns the NumVcpusPerSocket field if non-nil, zero value otherwise.

### GetNumVcpusPerSocketOk

`func (o *VmResourcesDefStatus) GetNumVcpusPerSocketOk() (*int32, bool)`

GetNumVcpusPerSocketOk returns a tuple with the NumVcpusPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVcpusPerSocket

`func (o *VmResourcesDefStatus) SetNumVcpusPerSocket(v int32)`

SetNumVcpusPerSocket sets NumVcpusPerSocket field to given value.

### HasNumVcpusPerSocket

`func (o *VmResourcesDefStatus) HasNumVcpusPerSocket() bool`

HasNumVcpusPerSocket returns a boolean if a field has been set.

### GetNumSockets

`func (o *VmResourcesDefStatus) GetNumSockets() int32`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *VmResourcesDefStatus) GetNumSocketsOk() (*int32, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *VmResourcesDefStatus) SetNumSockets(v int32)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *VmResourcesDefStatus) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.

### GetHardwareVirtualizationEnabled

`func (o *VmResourcesDefStatus) GetHardwareVirtualizationEnabled() bool`

GetHardwareVirtualizationEnabled returns the HardwareVirtualizationEnabled field if non-nil, zero value otherwise.

### GetHardwareVirtualizationEnabledOk

`func (o *VmResourcesDefStatus) GetHardwareVirtualizationEnabledOk() (*bool, bool)`

GetHardwareVirtualizationEnabledOk returns a tuple with the HardwareVirtualizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVirtualizationEnabled

`func (o *VmResourcesDefStatus) SetHardwareVirtualizationEnabled(v bool)`

SetHardwareVirtualizationEnabled sets HardwareVirtualizationEnabled field to given value.

### HasHardwareVirtualizationEnabled

`func (o *VmResourcesDefStatus) HasHardwareVirtualizationEnabled() bool`

HasHardwareVirtualizationEnabled returns a boolean if a field has been set.

### GetStorageConfig

`func (o *VmResourcesDefStatus) GetStorageConfig() VmStorageConfigStatus`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *VmResourcesDefStatus) GetStorageConfigOk() (*VmStorageConfigStatus, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *VmResourcesDefStatus) SetStorageConfig(v VmStorageConfigStatus)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *VmResourcesDefStatus) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetProtectionType

`func (o *VmResourcesDefStatus) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *VmResourcesDefStatus) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *VmResourcesDefStatus) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *VmResourcesDefStatus) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetGpuList

`func (o *VmResourcesDefStatus) GetGpuList() []VmGpuOutputStatus`

GetGpuList returns the GpuList field if non-nil, zero value otherwise.

### GetGpuListOk

`func (o *VmResourcesDefStatus) GetGpuListOk() (*[]VmGpuOutputStatus, bool)`

GetGpuListOk returns a tuple with the GpuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuList

`func (o *VmResourcesDefStatus) SetGpuList(v []VmGpuOutputStatus)`

SetGpuList sets GpuList field to given value.

### HasGpuList

`func (o *VmResourcesDefStatus) HasGpuList() bool`

HasGpuList returns a boolean if a field has been set.

### GetMachineType

`func (o *VmResourcesDefStatus) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *VmResourcesDefStatus) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *VmResourcesDefStatus) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *VmResourcesDefStatus) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetHardwareClockTimezone

`func (o *VmResourcesDefStatus) GetHardwareClockTimezone() string`

GetHardwareClockTimezone returns the HardwareClockTimezone field if non-nil, zero value otherwise.

### GetHardwareClockTimezoneOk

`func (o *VmResourcesDefStatus) GetHardwareClockTimezoneOk() (*string, bool)`

GetHardwareClockTimezoneOk returns a tuple with the HardwareClockTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareClockTimezone

`func (o *VmResourcesDefStatus) SetHardwareClockTimezone(v string)`

SetHardwareClockTimezone sets HardwareClockTimezone field to given value.

### HasHardwareClockTimezone

`func (o *VmResourcesDefStatus) HasHardwareClockTimezone() bool`

HasHardwareClockTimezone returns a boolean if a field has been set.

### GetGuestCustomization

`func (o *VmResourcesDefStatus) GetGuestCustomization() GuestCustomizationStatus`

GetGuestCustomization returns the GuestCustomization field if non-nil, zero value otherwise.

### GetGuestCustomizationOk

`func (o *VmResourcesDefStatus) GetGuestCustomizationOk() (*GuestCustomizationStatus, bool)`

GetGuestCustomizationOk returns a tuple with the GuestCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomization

`func (o *VmResourcesDefStatus) SetGuestCustomization(v GuestCustomizationStatus)`

SetGuestCustomization sets GuestCustomization field to given value.

### HasGuestCustomization

`func (o *VmResourcesDefStatus) HasGuestCustomization() bool`

HasGuestCustomization returns a boolean if a field has been set.

### GetPowerStateMechanism

`func (o *VmResourcesDefStatus) GetPowerStateMechanism() VmPowerStateMechanism`

GetPowerStateMechanism returns the PowerStateMechanism field if non-nil, zero value otherwise.

### GetPowerStateMechanismOk

`func (o *VmResourcesDefStatus) GetPowerStateMechanismOk() (*VmPowerStateMechanism, bool)`

GetPowerStateMechanismOk returns a tuple with the PowerStateMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStateMechanism

`func (o *VmResourcesDefStatus) SetPowerStateMechanism(v VmPowerStateMechanism)`

SetPowerStateMechanism sets PowerStateMechanism field to given value.

### HasPowerStateMechanism

`func (o *VmResourcesDefStatus) HasPowerStateMechanism() bool`

HasPowerStateMechanism returns a boolean if a field has been set.

### GetIsVcpuHardPinned

`func (o *VmResourcesDefStatus) GetIsVcpuHardPinned() bool`

GetIsVcpuHardPinned returns the IsVcpuHardPinned field if non-nil, zero value otherwise.

### GetIsVcpuHardPinnedOk

`func (o *VmResourcesDefStatus) GetIsVcpuHardPinnedOk() (*bool, bool)`

GetIsVcpuHardPinnedOk returns a tuple with the IsVcpuHardPinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVcpuHardPinned

`func (o *VmResourcesDefStatus) SetIsVcpuHardPinned(v bool)`

SetIsVcpuHardPinned sets IsVcpuHardPinned field to given value.

### HasIsVcpuHardPinned

`func (o *VmResourcesDefStatus) HasIsVcpuHardPinned() bool`

HasIsVcpuHardPinned returns a boolean if a field has been set.

### GetMemoryOvercommitEnabled

`func (o *VmResourcesDefStatus) GetMemoryOvercommitEnabled() bool`

GetMemoryOvercommitEnabled returns the MemoryOvercommitEnabled field if non-nil, zero value otherwise.

### GetMemoryOvercommitEnabledOk

`func (o *VmResourcesDefStatus) GetMemoryOvercommitEnabledOk() (*bool, bool)`

GetMemoryOvercommitEnabledOk returns a tuple with the MemoryOvercommitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOvercommitEnabled

`func (o *VmResourcesDefStatus) SetMemoryOvercommitEnabled(v bool)`

SetMemoryOvercommitEnabled sets MemoryOvercommitEnabled field to given value.

### HasMemoryOvercommitEnabled

`func (o *VmResourcesDefStatus) HasMemoryOvercommitEnabled() bool`

HasMemoryOvercommitEnabled returns a boolean if a field has been set.

### GetVnumaConfig

`func (o *VmResourcesDefStatus) GetVnumaConfig() VmVnumaConfig`

GetVnumaConfig returns the VnumaConfig field if non-nil, zero value otherwise.

### GetVnumaConfigOk

`func (o *VmResourcesDefStatus) GetVnumaConfigOk() (*VmVnumaConfig, bool)`

GetVnumaConfigOk returns a tuple with the VnumaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnumaConfig

`func (o *VmResourcesDefStatus) SetVnumaConfig(v VmVnumaConfig)`

SetVnumaConfig sets VnumaConfig field to given value.

### HasVnumaConfig

`func (o *VmResourcesDefStatus) HasVnumaConfig() bool`

HasVnumaConfig returns a boolean if a field has been set.

### GetNicList

`func (o *VmResourcesDefStatus) GetNicList() []VmNicOutputStatus`

GetNicList returns the NicList field if non-nil, zero value otherwise.

### GetNicListOk

`func (o *VmResourcesDefStatus) GetNicListOk() (*[]VmNicOutputStatus, bool)`

GetNicListOk returns a tuple with the NicList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicList

`func (o *VmResourcesDefStatus) SetNicList(v []VmNicOutputStatus)`

SetNicList sets NicList field to given value.

### HasNicList

`func (o *VmResourcesDefStatus) HasNicList() bool`

HasNicList returns a boolean if a field has been set.

### GetHostReference

`func (o *VmResourcesDefStatus) GetHostReference() Reference`

GetHostReference returns the HostReference field if non-nil, zero value otherwise.

### GetHostReferenceOk

`func (o *VmResourcesDefStatus) GetHostReferenceOk() (*Reference, bool)`

GetHostReferenceOk returns a tuple with the HostReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostReference

`func (o *VmResourcesDefStatus) SetHostReference(v Reference)`

SetHostReference sets HostReference field to given value.

### HasHostReference

`func (o *VmResourcesDefStatus) HasHostReference() bool`

HasHostReference returns a boolean if a field has been set.

### GetGuestOsId

`func (o *VmResourcesDefStatus) GetGuestOsId() string`

GetGuestOsId returns the GuestOsId field if non-nil, zero value otherwise.

### GetGuestOsIdOk

`func (o *VmResourcesDefStatus) GetGuestOsIdOk() (*string, bool)`

GetGuestOsIdOk returns a tuple with the GuestOsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsId

`func (o *VmResourcesDefStatus) SetGuestOsId(v string)`

SetGuestOsId sets GuestOsId field to given value.

### HasGuestOsId

`func (o *VmResourcesDefStatus) HasGuestOsId() bool`

HasGuestOsId returns a boolean if a field has been set.

### GetGuestTools

`func (o *VmResourcesDefStatus) GetGuestTools() GuestToolsStatus`

GetGuestTools returns the GuestTools field if non-nil, zero value otherwise.

### GetGuestToolsOk

`func (o *VmResourcesDefStatus) GetGuestToolsOk() (*GuestToolsStatus, bool)`

GetGuestToolsOk returns a tuple with the GuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTools

`func (o *VmResourcesDefStatus) SetGuestTools(v GuestToolsStatus)`

SetGuestTools sets GuestTools field to given value.

### HasGuestTools

`func (o *VmResourcesDefStatus) HasGuestTools() bool`

HasGuestTools returns a boolean if a field has been set.

### GetGpuConsoleEnabled

`func (o *VmResourcesDefStatus) GetGpuConsoleEnabled() bool`

GetGpuConsoleEnabled returns the GpuConsoleEnabled field if non-nil, zero value otherwise.

### GetGpuConsoleEnabledOk

`func (o *VmResourcesDefStatus) GetGpuConsoleEnabledOk() (*bool, bool)`

GetGpuConsoleEnabledOk returns a tuple with the GpuConsoleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuConsoleEnabled

`func (o *VmResourcesDefStatus) SetGpuConsoleEnabled(v bool)`

SetGpuConsoleEnabled sets GpuConsoleEnabled field to given value.

### HasGpuConsoleEnabled

`func (o *VmResourcesDefStatus) HasGpuConsoleEnabled() bool`

HasGpuConsoleEnabled returns a boolean if a field has been set.

### GetVtpmConfig

`func (o *VmResourcesDefStatus) GetVtpmConfig() VmVtpmStatus`

GetVtpmConfig returns the VtpmConfig field if non-nil, zero value otherwise.

### GetVtpmConfigOk

`func (o *VmResourcesDefStatus) GetVtpmConfigOk() (*VmVtpmStatus, bool)`

GetVtpmConfigOk returns a tuple with the VtpmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmConfig

`func (o *VmResourcesDefStatus) SetVtpmConfig(v VmVtpmStatus)`

SetVtpmConfig sets VtpmConfig field to given value.

### HasVtpmConfig

`func (o *VmResourcesDefStatus) HasVtpmConfig() bool`

HasVtpmConfig returns a boolean if a field has been set.

### GetEnableCpuPassthrough

`func (o *VmResourcesDefStatus) GetEnableCpuPassthrough() bool`

GetEnableCpuPassthrough returns the EnableCpuPassthrough field if non-nil, zero value otherwise.

### GetEnableCpuPassthroughOk

`func (o *VmResourcesDefStatus) GetEnableCpuPassthroughOk() (*bool, bool)`

GetEnableCpuPassthroughOk returns a tuple with the EnableCpuPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCpuPassthrough

`func (o *VmResourcesDefStatus) SetEnableCpuPassthrough(v bool)`

SetEnableCpuPassthrough sets EnableCpuPassthrough field to given value.

### HasEnableCpuPassthrough

`func (o *VmResourcesDefStatus) HasEnableCpuPassthrough() bool`

HasEnableCpuPassthrough returns a boolean if a field has been set.

### GetParentReference

`func (o *VmResourcesDefStatus) GetParentReference() Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *VmResourcesDefStatus) GetParentReferenceOk() (*Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *VmResourcesDefStatus) SetParentReference(v Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *VmResourcesDefStatus) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetHypervisorType

`func (o *VmResourcesDefStatus) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *VmResourcesDefStatus) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *VmResourcesDefStatus) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *VmResourcesDefStatus) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetDisableBranding

`func (o *VmResourcesDefStatus) GetDisableBranding() bool`

GetDisableBranding returns the DisableBranding field if non-nil, zero value otherwise.

### GetDisableBrandingOk

`func (o *VmResourcesDefStatus) GetDisableBrandingOk() (*bool, bool)`

GetDisableBrandingOk returns a tuple with the DisableBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBranding

`func (o *VmResourcesDefStatus) SetDisableBranding(v bool)`

SetDisableBranding sets DisableBranding field to given value.

### HasDisableBranding

`func (o *VmResourcesDefStatus) HasDisableBranding() bool`

HasDisableBranding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


