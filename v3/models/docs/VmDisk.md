# VmDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | Pointer to **string** | The device ID which is used to uniquely identify this particular disk.  | [optional] 
**DiskSizeBytes** | Pointer to **int64** | Size of the disk in Bytes. | [optional] 
**StorageConfig** | Pointer to [**DiskStorageConfig**](DiskStorageConfig.md) |  | [optional] 
**DeviceProperties** | Pointer to [**VmDiskDeviceProperties**](VmDiskDeviceProperties.md) |  | [optional] 
**DataSourceReference** | Pointer to [**VmDiskDataSourceReference**](VmDiskDataSourceReference.md) |  | [optional] 
**DiskSizeMib** | Pointer to **int32** | Size of the disk in MiB. Must match the size specified in &#39;disk_size_bytes&#39; - rounded up to the nearest MiB -  when that field is present.  | [optional] 
**VolumeGroupReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewVmDisk

`func NewVmDisk() *VmDisk`

NewVmDisk instantiates a new VmDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiskWithDefaults

`func NewVmDiskWithDefaults() *VmDisk`

NewVmDiskWithDefaults instantiates a new VmDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *VmDisk) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmDisk) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmDisk) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmDisk) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *VmDisk) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *VmDisk) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *VmDisk) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *VmDisk) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetStorageConfig

`func (o *VmDisk) GetStorageConfig() DiskStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *VmDisk) GetStorageConfigOk() (*DiskStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *VmDisk) SetStorageConfig(v DiskStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *VmDisk) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetDeviceProperties

`func (o *VmDisk) GetDeviceProperties() VmDiskDeviceProperties`

GetDeviceProperties returns the DeviceProperties field if non-nil, zero value otherwise.

### GetDevicePropertiesOk

`func (o *VmDisk) GetDevicePropertiesOk() (*VmDiskDeviceProperties, bool)`

GetDevicePropertiesOk returns a tuple with the DeviceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProperties

`func (o *VmDisk) SetDeviceProperties(v VmDiskDeviceProperties)`

SetDeviceProperties sets DeviceProperties field to given value.

### HasDeviceProperties

`func (o *VmDisk) HasDeviceProperties() bool`

HasDeviceProperties returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VmDisk) GetDataSourceReference() VmDiskDataSourceReference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VmDisk) GetDataSourceReferenceOk() (*VmDiskDataSourceReference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VmDisk) SetDataSourceReference(v VmDiskDataSourceReference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VmDisk) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetDiskSizeMib

`func (o *VmDisk) GetDiskSizeMib() int32`

GetDiskSizeMib returns the DiskSizeMib field if non-nil, zero value otherwise.

### GetDiskSizeMibOk

`func (o *VmDisk) GetDiskSizeMibOk() (*int32, bool)`

GetDiskSizeMibOk returns a tuple with the DiskSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMib

`func (o *VmDisk) SetDiskSizeMib(v int32)`

SetDiskSizeMib sets DiskSizeMib field to given value.

### HasDiskSizeMib

`func (o *VmDisk) HasDiskSizeMib() bool`

HasDiskSizeMib returns a boolean if a field has been set.

### GetVolumeGroupReference

`func (o *VmDisk) GetVolumeGroupReference() Reference`

GetVolumeGroupReference returns the VolumeGroupReference field if non-nil, zero value otherwise.

### GetVolumeGroupReferenceOk

`func (o *VmDisk) GetVolumeGroupReferenceOk() (*Reference, bool)`

GetVolumeGroupReferenceOk returns a tuple with the VolumeGroupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupReference

`func (o *VmDisk) SetVolumeGroupReference(v Reference)`

SetVolumeGroupReference sets VolumeGroupReference field to given value.

### HasVolumeGroupReference

`func (o *VmDisk) HasVolumeGroupReference() bool`

HasVolumeGroupReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


