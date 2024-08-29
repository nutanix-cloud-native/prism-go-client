# VolumeDiskResourceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Index of the volume disk in the group. | [optional] 
**StorageContainerUuid** | Pointer to **string** | Storage container UUID on which the disk has been created. | [optional] 
**DiskSizeMib** | Pointer to **int64** | Size of the disk in MiB. | [optional] 
**DiskSizeBytes** | Pointer to **int64** | Size of the disk in Bytes. | [optional] 
**UUID** | Pointer to **string** | The device ID which is used to uniquely identify this particular disk.  | [optional] 

## Methods

### NewVolumeDiskResourceOutput

`func NewVolumeDiskResourceOutput() *VolumeDiskResourceOutput`

NewVolumeDiskResourceOutput instantiates a new VolumeDiskResourceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDiskResourceOutputWithDefaults

`func NewVolumeDiskResourceOutputWithDefaults() *VolumeDiskResourceOutput`

NewVolumeDiskResourceOutputWithDefaults instantiates a new VolumeDiskResourceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *VolumeDiskResourceOutput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VolumeDiskResourceOutput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VolumeDiskResourceOutput) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VolumeDiskResourceOutput) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetStorageContainerUuid

`func (o *VolumeDiskResourceOutput) GetStorageContainerUuid() string`

GetStorageContainerUuid returns the StorageContainerUuid field if non-nil, zero value otherwise.

### GetStorageContainerUuidOk

`func (o *VolumeDiskResourceOutput) GetStorageContainerUuidOk() (*string, bool)`

GetStorageContainerUuidOk returns a tuple with the StorageContainerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerUuid

`func (o *VolumeDiskResourceOutput) SetStorageContainerUuid(v string)`

SetStorageContainerUuid sets StorageContainerUuid field to given value.

### HasStorageContainerUuid

`func (o *VolumeDiskResourceOutput) HasStorageContainerUuid() bool`

HasStorageContainerUuid returns a boolean if a field has been set.

### GetDiskSizeMib

`func (o *VolumeDiskResourceOutput) GetDiskSizeMib() int64`

GetDiskSizeMib returns the DiskSizeMib field if non-nil, zero value otherwise.

### GetDiskSizeMibOk

`func (o *VolumeDiskResourceOutput) GetDiskSizeMibOk() (*int64, bool)`

GetDiskSizeMibOk returns a tuple with the DiskSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMib

`func (o *VolumeDiskResourceOutput) SetDiskSizeMib(v int64)`

SetDiskSizeMib sets DiskSizeMib field to given value.

### HasDiskSizeMib

`func (o *VolumeDiskResourceOutput) HasDiskSizeMib() bool`

HasDiskSizeMib returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *VolumeDiskResourceOutput) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *VolumeDiskResourceOutput) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *VolumeDiskResourceOutput) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *VolumeDiskResourceOutput) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetUUID

`func (o *VolumeDiskResourceOutput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VolumeDiskResourceOutput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VolumeDiskResourceOutput) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VolumeDiskResourceOutput) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


