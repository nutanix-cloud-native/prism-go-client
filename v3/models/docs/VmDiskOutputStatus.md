# VmDiskOutputStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | Pointer to **string** | The device ID which is used to uniquely identify this particular disk.  | [optional] 
**IsEncryptionInProgress** | Pointer to **bool** | Indicates whether encryption is in progress or not. | [optional] [default to false]
**IsMigrationInProgress** | Pointer to **bool** | Indicated if the disk is undergoing migration to another container or not.  | [optional] [default to false]
**DiskSizeBytes** | Pointer to **int64** | Size of the disk in Bytes. | [optional] 
**StorageConfig** | Pointer to [**DiskStorageConfig**](DiskStorageConfig.md) |  | [optional] 
**DeviceProperties** | Pointer to [**VmDiskDeviceProperties**](VmDiskDeviceProperties.md) |  | [optional] 
**DataSourceReference** | Pointer to [**VmDiskOutputStatusDataSourceReference**](VmDiskOutputStatusDataSourceReference.md) |  | [optional] 
**DiskSizeMib** | Pointer to **int32** | Size of the disk in MiB. Must match the size specified in &#39;disk_size_bytes&#39; - rounded up to the nearest MiB -  when that field is present.  | [optional] 
**VolumeGroupReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewVmDiskOutputStatus

`func NewVmDiskOutputStatus() *VmDiskOutputStatus`

NewVmDiskOutputStatus instantiates a new VmDiskOutputStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiskOutputStatusWithDefaults

`func NewVmDiskOutputStatusWithDefaults() *VmDiskOutputStatus`

NewVmDiskOutputStatusWithDefaults instantiates a new VmDiskOutputStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *VmDiskOutputStatus) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmDiskOutputStatus) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmDiskOutputStatus) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmDiskOutputStatus) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetIsEncryptionInProgress

`func (o *VmDiskOutputStatus) GetIsEncryptionInProgress() bool`

GetIsEncryptionInProgress returns the IsEncryptionInProgress field if non-nil, zero value otherwise.

### GetIsEncryptionInProgressOk

`func (o *VmDiskOutputStatus) GetIsEncryptionInProgressOk() (*bool, bool)`

GetIsEncryptionInProgressOk returns a tuple with the IsEncryptionInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryptionInProgress

`func (o *VmDiskOutputStatus) SetIsEncryptionInProgress(v bool)`

SetIsEncryptionInProgress sets IsEncryptionInProgress field to given value.

### HasIsEncryptionInProgress

`func (o *VmDiskOutputStatus) HasIsEncryptionInProgress() bool`

HasIsEncryptionInProgress returns a boolean if a field has been set.

### GetIsMigrationInProgress

`func (o *VmDiskOutputStatus) GetIsMigrationInProgress() bool`

GetIsMigrationInProgress returns the IsMigrationInProgress field if non-nil, zero value otherwise.

### GetIsMigrationInProgressOk

`func (o *VmDiskOutputStatus) GetIsMigrationInProgressOk() (*bool, bool)`

GetIsMigrationInProgressOk returns a tuple with the IsMigrationInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMigrationInProgress

`func (o *VmDiskOutputStatus) SetIsMigrationInProgress(v bool)`

SetIsMigrationInProgress sets IsMigrationInProgress field to given value.

### HasIsMigrationInProgress

`func (o *VmDiskOutputStatus) HasIsMigrationInProgress() bool`

HasIsMigrationInProgress returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *VmDiskOutputStatus) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *VmDiskOutputStatus) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *VmDiskOutputStatus) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *VmDiskOutputStatus) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetStorageConfig

`func (o *VmDiskOutputStatus) GetStorageConfig() DiskStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *VmDiskOutputStatus) GetStorageConfigOk() (*DiskStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *VmDiskOutputStatus) SetStorageConfig(v DiskStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *VmDiskOutputStatus) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetDeviceProperties

`func (o *VmDiskOutputStatus) GetDeviceProperties() VmDiskDeviceProperties`

GetDeviceProperties returns the DeviceProperties field if non-nil, zero value otherwise.

### GetDevicePropertiesOk

`func (o *VmDiskOutputStatus) GetDevicePropertiesOk() (*VmDiskDeviceProperties, bool)`

GetDevicePropertiesOk returns a tuple with the DeviceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProperties

`func (o *VmDiskOutputStatus) SetDeviceProperties(v VmDiskDeviceProperties)`

SetDeviceProperties sets DeviceProperties field to given value.

### HasDeviceProperties

`func (o *VmDiskOutputStatus) HasDeviceProperties() bool`

HasDeviceProperties returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VmDiskOutputStatus) GetDataSourceReference() VmDiskOutputStatusDataSourceReference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VmDiskOutputStatus) GetDataSourceReferenceOk() (*VmDiskOutputStatusDataSourceReference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VmDiskOutputStatus) SetDataSourceReference(v VmDiskOutputStatusDataSourceReference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VmDiskOutputStatus) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetDiskSizeMib

`func (o *VmDiskOutputStatus) GetDiskSizeMib() int32`

GetDiskSizeMib returns the DiskSizeMib field if non-nil, zero value otherwise.

### GetDiskSizeMibOk

`func (o *VmDiskOutputStatus) GetDiskSizeMibOk() (*int32, bool)`

GetDiskSizeMibOk returns a tuple with the DiskSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMib

`func (o *VmDiskOutputStatus) SetDiskSizeMib(v int32)`

SetDiskSizeMib sets DiskSizeMib field to given value.

### HasDiskSizeMib

`func (o *VmDiskOutputStatus) HasDiskSizeMib() bool`

HasDiskSizeMib returns a boolean if a field has been set.

### GetVolumeGroupReference

`func (o *VmDiskOutputStatus) GetVolumeGroupReference() Reference`

GetVolumeGroupReference returns the VolumeGroupReference field if non-nil, zero value otherwise.

### GetVolumeGroupReferenceOk

`func (o *VmDiskOutputStatus) GetVolumeGroupReferenceOk() (*Reference, bool)`

GetVolumeGroupReferenceOk returns a tuple with the VolumeGroupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupReference

`func (o *VmDiskOutputStatus) SetVolumeGroupReference(v Reference)`

SetVolumeGroupReference sets VolumeGroupReference field to given value.

### HasVolumeGroupReference

`func (o *VmDiskOutputStatus) HasVolumeGroupReference() bool`

HasVolumeGroupReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


