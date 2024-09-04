# DiskDefStatusResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**EnabledFeaturesList** | Pointer to **[]string** | Disk feature flags - &#39;CanAddAsNewDisk&#39;: Flag to indicate if this disk can be added as    new disk. - &#39;CanAddAsOldDisk&#39;: Flag to indicate if the disk can be added as    old disk. - &#39;BootDisk&#39;: Flag to indicate if its a boot disk. - &#39;OnlyBootDisk&#39;: Flag to indicate if the disk is boot only and    no disk operation to be run on it. - &#39;SelfEncryptingEnabled&#39;: Flag to indicate if the disk has self    encryption enabled. - &#39;PasswordProtected&#39;: Flag to indicate if the disk is password    protected. - &#39;SelfManagedNvme&#39;: Flag to indicate if the NVMe disk is self   managed and no host/CVM reboot required.  | [optional] 
**Vendor** | Pointer to **string** | Disk vendor. | [optional] 
**MountPath** | Pointer to **string** | Mount path. | [optional] 
**StoragePoolUuid** | Pointer to **string** | Storage pool uuid. | [optional] 
**StateList** | Pointer to **[]string** | Array of disk states - &#39;DataMigrationInitiated&#39;: Data Migration Initiated. - &#39;MarkedForRemovalButNotDetachable&#39;: Marked for removal, data    migration is in progress. - &#39;ReadyToDetach&#39;: Flag to indicate the disk is detachable. - &#39;DataMigrated&#39;: Flag to indicate if data migration is completed for    this disk. - &#39;MarkedForRemoval&#39;: Flag to indicate if the disk is marked for    removal. - &#39;Online&#39;: Flag to indicate if the disk is online. - &#39;Bad&#39;: Flag to indicate if the disk is bad. - &#39;Mounted&#39;: Flag to indicate if the disk is mounted. - &#39;UnderDiagnosis&#39;: Flag to indicate if the disk is under diagnosis.  | [optional] 
**StorageTierType** | Pointer to **string** | Storage tier type. | [optional] 
**SizeBytes** | Pointer to **int64** | Disk size in bytes. | [optional] 
**SlotNumber** | Pointer to **int32** | Disk location in a node. | [optional] 
**SerialNumber** | Pointer to **string** | Disk serial number. | [optional] 
**Model** | Pointer to **string** | Disk model. | [optional] 
**FirmwareVersion** | Pointer to **string** | Firmware version. | [optional] 

## Methods

### NewDiskDefStatusResources

`func NewDiskDefStatusResources() *DiskDefStatusResources`

NewDiskDefStatusResources instantiates a new DiskDefStatusResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskDefStatusResourcesWithDefaults

`func NewDiskDefStatusResourcesWithDefaults() *DiskDefStatusResources`

NewDiskDefStatusResourcesWithDefaults instantiates a new DiskDefStatusResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostReference

`func (o *DiskDefStatusResources) GetHostReference() Reference`

GetHostReference returns the HostReference field if non-nil, zero value otherwise.

### GetHostReferenceOk

`func (o *DiskDefStatusResources) GetHostReferenceOk() (*Reference, bool)`

GetHostReferenceOk returns a tuple with the HostReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostReference

`func (o *DiskDefStatusResources) SetHostReference(v Reference)`

SetHostReference sets HostReference field to given value.

### HasHostReference

`func (o *DiskDefStatusResources) HasHostReference() bool`

HasHostReference returns a boolean if a field has been set.

### GetEnabledFeaturesList

`func (o *DiskDefStatusResources) GetEnabledFeaturesList() []string`

GetEnabledFeaturesList returns the EnabledFeaturesList field if non-nil, zero value otherwise.

### GetEnabledFeaturesListOk

`func (o *DiskDefStatusResources) GetEnabledFeaturesListOk() (*[]string, bool)`

GetEnabledFeaturesListOk returns a tuple with the EnabledFeaturesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeaturesList

`func (o *DiskDefStatusResources) SetEnabledFeaturesList(v []string)`

SetEnabledFeaturesList sets EnabledFeaturesList field to given value.

### HasEnabledFeaturesList

`func (o *DiskDefStatusResources) HasEnabledFeaturesList() bool`

HasEnabledFeaturesList returns a boolean if a field has been set.

### GetVendor

`func (o *DiskDefStatusResources) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DiskDefStatusResources) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DiskDefStatusResources) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DiskDefStatusResources) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetMountPath

`func (o *DiskDefStatusResources) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *DiskDefStatusResources) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *DiskDefStatusResources) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *DiskDefStatusResources) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetStoragePoolUuid

`func (o *DiskDefStatusResources) GetStoragePoolUuid() string`

GetStoragePoolUuid returns the StoragePoolUuid field if non-nil, zero value otherwise.

### GetStoragePoolUuidOk

`func (o *DiskDefStatusResources) GetStoragePoolUuidOk() (*string, bool)`

GetStoragePoolUuidOk returns a tuple with the StoragePoolUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolUuid

`func (o *DiskDefStatusResources) SetStoragePoolUuid(v string)`

SetStoragePoolUuid sets StoragePoolUuid field to given value.

### HasStoragePoolUuid

`func (o *DiskDefStatusResources) HasStoragePoolUuid() bool`

HasStoragePoolUuid returns a boolean if a field has been set.

### GetStateList

`func (o *DiskDefStatusResources) GetStateList() []string`

GetStateList returns the StateList field if non-nil, zero value otherwise.

### GetStateListOk

`func (o *DiskDefStatusResources) GetStateListOk() (*[]string, bool)`

GetStateListOk returns a tuple with the StateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateList

`func (o *DiskDefStatusResources) SetStateList(v []string)`

SetStateList sets StateList field to given value.

### HasStateList

`func (o *DiskDefStatusResources) HasStateList() bool`

HasStateList returns a boolean if a field has been set.

### GetStorageTierType

`func (o *DiskDefStatusResources) GetStorageTierType() string`

GetStorageTierType returns the StorageTierType field if non-nil, zero value otherwise.

### GetStorageTierTypeOk

`func (o *DiskDefStatusResources) GetStorageTierTypeOk() (*string, bool)`

GetStorageTierTypeOk returns a tuple with the StorageTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTierType

`func (o *DiskDefStatusResources) SetStorageTierType(v string)`

SetStorageTierType sets StorageTierType field to given value.

### HasStorageTierType

`func (o *DiskDefStatusResources) HasStorageTierType() bool`

HasStorageTierType returns a boolean if a field has been set.

### GetSizeBytes

`func (o *DiskDefStatusResources) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *DiskDefStatusResources) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *DiskDefStatusResources) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *DiskDefStatusResources) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetSlotNumber

`func (o *DiskDefStatusResources) GetSlotNumber() int32`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *DiskDefStatusResources) GetSlotNumberOk() (*int32, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *DiskDefStatusResources) SetSlotNumber(v int32)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *DiskDefStatusResources) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DiskDefStatusResources) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DiskDefStatusResources) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DiskDefStatusResources) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DiskDefStatusResources) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetModel

`func (o *DiskDefStatusResources) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiskDefStatusResources) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiskDefStatusResources) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiskDefStatusResources) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetFirmwareVersion

`func (o *DiskDefStatusResources) GetFirmwareVersion() string`

GetFirmwareVersion returns the FirmwareVersion field if non-nil, zero value otherwise.

### GetFirmwareVersionOk

`func (o *DiskDefStatusResources) GetFirmwareVersionOk() (*string, bool)`

GetFirmwareVersionOk returns a tuple with the FirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareVersion

`func (o *DiskDefStatusResources) SetFirmwareVersion(v string)`

SetFirmwareVersion sets FirmwareVersion field to given value.

### HasFirmwareVersion

`func (o *DiskDefStatusResources) HasFirmwareVersion() bool`

HasFirmwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


