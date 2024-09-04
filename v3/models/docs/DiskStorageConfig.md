# DiskStorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlashMode** | Pointer to **string** | State of the storage policy to pin virtual disks to the hot tier. When specified as a VM attribute, the storage policy applies to all virtual disks of the VM unless overridden by the same attribute specified for a virtual disk.  | [optional] 
**StorageContainerReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewDiskStorageConfig

`func NewDiskStorageConfig() *DiskStorageConfig`

NewDiskStorageConfig instantiates a new DiskStorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskStorageConfigWithDefaults

`func NewDiskStorageConfigWithDefaults() *DiskStorageConfig`

NewDiskStorageConfigWithDefaults instantiates a new DiskStorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlashMode

`func (o *DiskStorageConfig) GetFlashMode() string`

GetFlashMode returns the FlashMode field if non-nil, zero value otherwise.

### GetFlashModeOk

`func (o *DiskStorageConfig) GetFlashModeOk() (*string, bool)`

GetFlashModeOk returns a tuple with the FlashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashMode

`func (o *DiskStorageConfig) SetFlashMode(v string)`

SetFlashMode sets FlashMode field to given value.

### HasFlashMode

`func (o *DiskStorageConfig) HasFlashMode() bool`

HasFlashMode returns a boolean if a field has been set.

### GetStorageContainerReference

`func (o *DiskStorageConfig) GetStorageContainerReference() Reference`

GetStorageContainerReference returns the StorageContainerReference field if non-nil, zero value otherwise.

### GetStorageContainerReferenceOk

`func (o *DiskStorageConfig) GetStorageContainerReferenceOk() (*Reference, bool)`

GetStorageContainerReferenceOk returns a tuple with the StorageContainerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerReference

`func (o *DiskStorageConfig) SetStorageContainerReference(v Reference)`

SetStorageContainerReference sets StorageContainerReference field to given value.

### HasStorageContainerReference

`func (o *DiskStorageConfig) HasStorageContainerReference() bool`

HasStorageContainerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


