# VolumeDiskResourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Index of the volume disk in the group. | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**DiskSizeMib** | Pointer to **int64** | Size of the disk in MiB. | [optional] 
**DiskSizeBytes** | Pointer to **int64** | Size of the disk in Bytes. | [optional] 
**StorageContainerUuid** | Pointer to **string** | Storage container UUID on which to create the disk. | [optional] 

## Methods

### NewVolumeDiskResourceInput

`func NewVolumeDiskResourceInput() *VolumeDiskResourceInput`

NewVolumeDiskResourceInput instantiates a new VolumeDiskResourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDiskResourceInputWithDefaults

`func NewVolumeDiskResourceInputWithDefaults() *VolumeDiskResourceInput`

NewVolumeDiskResourceInputWithDefaults instantiates a new VolumeDiskResourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *VolumeDiskResourceInput) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VolumeDiskResourceInput) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VolumeDiskResourceInput) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VolumeDiskResourceInput) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VolumeDiskResourceInput) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VolumeDiskResourceInput) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VolumeDiskResourceInput) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VolumeDiskResourceInput) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetDiskSizeMib

`func (o *VolumeDiskResourceInput) GetDiskSizeMib() int64`

GetDiskSizeMib returns the DiskSizeMib field if non-nil, zero value otherwise.

### GetDiskSizeMibOk

`func (o *VolumeDiskResourceInput) GetDiskSizeMibOk() (*int64, bool)`

GetDiskSizeMibOk returns a tuple with the DiskSizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMib

`func (o *VolumeDiskResourceInput) SetDiskSizeMib(v int64)`

SetDiskSizeMib sets DiskSizeMib field to given value.

### HasDiskSizeMib

`func (o *VolumeDiskResourceInput) HasDiskSizeMib() bool`

HasDiskSizeMib returns a boolean if a field has been set.

### GetDiskSizeBytes

`func (o *VolumeDiskResourceInput) GetDiskSizeBytes() int64`

GetDiskSizeBytes returns the DiskSizeBytes field if non-nil, zero value otherwise.

### GetDiskSizeBytesOk

`func (o *VolumeDiskResourceInput) GetDiskSizeBytesOk() (*int64, bool)`

GetDiskSizeBytesOk returns a tuple with the DiskSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeBytes

`func (o *VolumeDiskResourceInput) SetDiskSizeBytes(v int64)`

SetDiskSizeBytes sets DiskSizeBytes field to given value.

### HasDiskSizeBytes

`func (o *VolumeDiskResourceInput) HasDiskSizeBytes() bool`

HasDiskSizeBytes returns a boolean if a field has been set.

### GetStorageContainerUuid

`func (o *VolumeDiskResourceInput) GetStorageContainerUuid() string`

GetStorageContainerUuid returns the StorageContainerUuid field if non-nil, zero value otherwise.

### GetStorageContainerUuidOk

`func (o *VolumeDiskResourceInput) GetStorageContainerUuidOk() (*string, bool)`

GetStorageContainerUuidOk returns a tuple with the StorageContainerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerUuid

`func (o *VolumeDiskResourceInput) SetStorageContainerUuid(v string)`

SetStorageContainerUuid sets StorageContainerUuid field to given value.

### HasStorageContainerUuid

`func (o *VolumeDiskResourceInput) HasStorageContainerUuid() bool`

HasStorageContainerUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


