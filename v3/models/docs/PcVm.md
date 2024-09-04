# PcVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmName** | **string** | VM name. | 
**DataDiskSizeBytes** | **int64** | Data disk size in bytes. | 
**DnsServerIpList** | Pointer to **[]string** | List of DNS IP addresses. | [optional] 
**NicList** | [**[]PcVmNicConfiguration**](PcVmNicConfiguration.md) |  | 
**ContainerUuid** | Pointer to **string** | Container uuid. | [optional] 
**NumSockets** | **int64** | Number of sockets allocated per VM. | 
**MemorySizeBytes** | **int64** | Memory in bytes. | 
**Status** | Pointer to **string** | Prism central VM status | [optional] [readonly] 
**PowerState** | Pointer to **string** | The current power state of the VM. | [optional] [readonly] 
**ContainerName** | Pointer to **string** | Container name. | [optional] 
**VmUuid** | Pointer to **string** | VM uuid. | [optional] 
**NtpServerList** | Pointer to **[]string** | List of NTP servers. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewPcVm

`func NewPcVm(vmName string, dataDiskSizeBytes int64, nicList []PcVmNicConfiguration, numSockets int64, memorySizeBytes int64, ) *PcVm`

NewPcVm instantiates a new PcVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcVmWithDefaults

`func NewPcVmWithDefaults() *PcVm`

NewPcVmWithDefaults instantiates a new PcVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmName

`func (o *PcVm) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *PcVm) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *PcVm) SetVmName(v string)`

SetVmName sets VmName field to given value.


### GetDataDiskSizeBytes

`func (o *PcVm) GetDataDiskSizeBytes() int64`

GetDataDiskSizeBytes returns the DataDiskSizeBytes field if non-nil, zero value otherwise.

### GetDataDiskSizeBytesOk

`func (o *PcVm) GetDataDiskSizeBytesOk() (*int64, bool)`

GetDataDiskSizeBytesOk returns a tuple with the DataDiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSizeBytes

`func (o *PcVm) SetDataDiskSizeBytes(v int64)`

SetDataDiskSizeBytes sets DataDiskSizeBytes field to given value.


### GetDnsServerIpList

`func (o *PcVm) GetDnsServerIpList() []string`

GetDnsServerIpList returns the DnsServerIpList field if non-nil, zero value otherwise.

### GetDnsServerIpListOk

`func (o *PcVm) GetDnsServerIpListOk() (*[]string, bool)`

GetDnsServerIpListOk returns a tuple with the DnsServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerIpList

`func (o *PcVm) SetDnsServerIpList(v []string)`

SetDnsServerIpList sets DnsServerIpList field to given value.

### HasDnsServerIpList

`func (o *PcVm) HasDnsServerIpList() bool`

HasDnsServerIpList returns a boolean if a field has been set.

### GetNicList

`func (o *PcVm) GetNicList() []PcVmNicConfiguration`

GetNicList returns the NicList field if non-nil, zero value otherwise.

### GetNicListOk

`func (o *PcVm) GetNicListOk() (*[]PcVmNicConfiguration, bool)`

GetNicListOk returns a tuple with the NicList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicList

`func (o *PcVm) SetNicList(v []PcVmNicConfiguration)`

SetNicList sets NicList field to given value.


### GetContainerUuid

`func (o *PcVm) GetContainerUuid() string`

GetContainerUuid returns the ContainerUuid field if non-nil, zero value otherwise.

### GetContainerUuidOk

`func (o *PcVm) GetContainerUuidOk() (*string, bool)`

GetContainerUuidOk returns a tuple with the ContainerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerUuid

`func (o *PcVm) SetContainerUuid(v string)`

SetContainerUuid sets ContainerUuid field to given value.

### HasContainerUuid

`func (o *PcVm) HasContainerUuid() bool`

HasContainerUuid returns a boolean if a field has been set.

### GetNumSockets

`func (o *PcVm) GetNumSockets() int64`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *PcVm) GetNumSocketsOk() (*int64, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *PcVm) SetNumSockets(v int64)`

SetNumSockets sets NumSockets field to given value.


### GetMemorySizeBytes

`func (o *PcVm) GetMemorySizeBytes() int64`

GetMemorySizeBytes returns the MemorySizeBytes field if non-nil, zero value otherwise.

### GetMemorySizeBytesOk

`func (o *PcVm) GetMemorySizeBytesOk() (*int64, bool)`

GetMemorySizeBytesOk returns a tuple with the MemorySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeBytes

`func (o *PcVm) SetMemorySizeBytes(v int64)`

SetMemorySizeBytes sets MemorySizeBytes field to given value.


### GetStatus

`func (o *PcVm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PcVm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PcVm) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PcVm) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPowerState

`func (o *PcVm) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *PcVm) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *PcVm) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *PcVm) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetContainerName

`func (o *PcVm) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *PcVm) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *PcVm) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *PcVm) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetVmUuid

`func (o *PcVm) GetVmUuid() string`

GetVmUuid returns the VmUuid field if non-nil, zero value otherwise.

### GetVmUuidOk

`func (o *PcVm) GetVmUuidOk() (*string, bool)`

GetVmUuidOk returns a tuple with the VmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUuid

`func (o *PcVm) SetVmUuid(v string)`

SetVmUuid sets VmUuid field to given value.

### HasVmUuid

`func (o *PcVm) HasVmUuid() bool`

HasVmUuid returns a boolean if a field has been set.

### GetNtpServerList

`func (o *PcVm) GetNtpServerList() []string`

GetNtpServerList returns the NtpServerList field if non-nil, zero value otherwise.

### GetNtpServerListOk

`func (o *PcVm) GetNtpServerListOk() (*[]string, bool)`

GetNtpServerListOk returns a tuple with the NtpServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerList

`func (o *PcVm) SetNtpServerList(v []string)`

SetNtpServerList sets NtpServerList field to given value.

### HasNtpServerList

`func (o *PcVm) HasNtpServerList() bool`

HasNtpServerList returns a boolean if a field has been set.

### GetClusterReference

`func (o *PcVm) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *PcVm) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *PcVm) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *PcVm) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


