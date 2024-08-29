# VolumeGroupResourcesOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlashMode** | Pointer to **string** | Flash Mode, if enabled all volume disks of the VG will be pinned to SSD tier.  | [optional] 
**IscsiTargetName** | Pointer to **string** | iSCSI target full name | [optional] 
**EnabledAuthentications** | Pointer to **string** | Which authentication is enabled for VG. | [optional] 
**AttachmentList** | Pointer to [**[]AttachmentReferenceOutput**](AttachmentReferenceOutput.md) | VMs attached to volume group. | [optional] 
**CreatedBy** | Pointer to **string** | Service/user who created this Volume Group. | [optional] 
**ParentReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**SharingStatus** | Pointer to **string** | Whether the volume group can be shared across multiple iSCSI initiators.  | [optional] 
**DiskList** | Pointer to [**[]VolumeDiskResourceOutput**](VolumeDiskResourceOutput.md) | Volume group disk specification. | [optional] 
**SizeBytes** | Pointer to **int64** | The total size of the Volume Group in bytes | [optional] 
**UsageType** | Pointer to **string** | Expected usage type for the volume group. | [optional] 
**LoadBalanceVmAttachments** | Pointer to **bool** | Whether volume group load balancing is enabled. | [optional] 
**IsHidden** | Pointer to **bool** | Whether the VG is meant to be hidden or not. | [optional] 
**SizeMib** | Pointer to **int64** | The total size of the Volume Group in mib | [optional] 
**IscsiTargetPrefix** | Pointer to **string** | iSCSI target prefix-name. | [optional] 

## Methods

### NewVolumeGroupResourcesOutput

`func NewVolumeGroupResourcesOutput() *VolumeGroupResourcesOutput`

NewVolumeGroupResourcesOutput instantiates a new VolumeGroupResourcesOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupResourcesOutputWithDefaults

`func NewVolumeGroupResourcesOutputWithDefaults() *VolumeGroupResourcesOutput`

NewVolumeGroupResourcesOutputWithDefaults instantiates a new VolumeGroupResourcesOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlashMode

`func (o *VolumeGroupResourcesOutput) GetFlashMode() string`

GetFlashMode returns the FlashMode field if non-nil, zero value otherwise.

### GetFlashModeOk

`func (o *VolumeGroupResourcesOutput) GetFlashModeOk() (*string, bool)`

GetFlashModeOk returns a tuple with the FlashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashMode

`func (o *VolumeGroupResourcesOutput) SetFlashMode(v string)`

SetFlashMode sets FlashMode field to given value.

### HasFlashMode

`func (o *VolumeGroupResourcesOutput) HasFlashMode() bool`

HasFlashMode returns a boolean if a field has been set.

### GetIscsiTargetName

`func (o *VolumeGroupResourcesOutput) GetIscsiTargetName() string`

GetIscsiTargetName returns the IscsiTargetName field if non-nil, zero value otherwise.

### GetIscsiTargetNameOk

`func (o *VolumeGroupResourcesOutput) GetIscsiTargetNameOk() (*string, bool)`

GetIscsiTargetNameOk returns a tuple with the IscsiTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiTargetName

`func (o *VolumeGroupResourcesOutput) SetIscsiTargetName(v string)`

SetIscsiTargetName sets IscsiTargetName field to given value.

### HasIscsiTargetName

`func (o *VolumeGroupResourcesOutput) HasIscsiTargetName() bool`

HasIscsiTargetName returns a boolean if a field has been set.

### GetEnabledAuthentications

`func (o *VolumeGroupResourcesOutput) GetEnabledAuthentications() string`

GetEnabledAuthentications returns the EnabledAuthentications field if non-nil, zero value otherwise.

### GetEnabledAuthenticationsOk

`func (o *VolumeGroupResourcesOutput) GetEnabledAuthenticationsOk() (*string, bool)`

GetEnabledAuthenticationsOk returns a tuple with the EnabledAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAuthentications

`func (o *VolumeGroupResourcesOutput) SetEnabledAuthentications(v string)`

SetEnabledAuthentications sets EnabledAuthentications field to given value.

### HasEnabledAuthentications

`func (o *VolumeGroupResourcesOutput) HasEnabledAuthentications() bool`

HasEnabledAuthentications returns a boolean if a field has been set.

### GetAttachmentList

`func (o *VolumeGroupResourcesOutput) GetAttachmentList() []AttachmentReferenceOutput`

GetAttachmentList returns the AttachmentList field if non-nil, zero value otherwise.

### GetAttachmentListOk

`func (o *VolumeGroupResourcesOutput) GetAttachmentListOk() (*[]AttachmentReferenceOutput, bool)`

GetAttachmentListOk returns a tuple with the AttachmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentList

`func (o *VolumeGroupResourcesOutput) SetAttachmentList(v []AttachmentReferenceOutput)`

SetAttachmentList sets AttachmentList field to given value.

### HasAttachmentList

`func (o *VolumeGroupResourcesOutput) HasAttachmentList() bool`

HasAttachmentList returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeGroupResourcesOutput) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeGroupResourcesOutput) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeGroupResourcesOutput) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeGroupResourcesOutput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetParentReference

`func (o *VolumeGroupResourcesOutput) GetParentReference() Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *VolumeGroupResourcesOutput) GetParentReferenceOk() (*Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *VolumeGroupResourcesOutput) SetParentReference(v Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *VolumeGroupResourcesOutput) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetSharingStatus

`func (o *VolumeGroupResourcesOutput) GetSharingStatus() string`

GetSharingStatus returns the SharingStatus field if non-nil, zero value otherwise.

### GetSharingStatusOk

`func (o *VolumeGroupResourcesOutput) GetSharingStatusOk() (*string, bool)`

GetSharingStatusOk returns a tuple with the SharingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingStatus

`func (o *VolumeGroupResourcesOutput) SetSharingStatus(v string)`

SetSharingStatus sets SharingStatus field to given value.

### HasSharingStatus

`func (o *VolumeGroupResourcesOutput) HasSharingStatus() bool`

HasSharingStatus returns a boolean if a field has been set.

### GetDiskList

`func (o *VolumeGroupResourcesOutput) GetDiskList() []VolumeDiskResourceOutput`

GetDiskList returns the DiskList field if non-nil, zero value otherwise.

### GetDiskListOk

`func (o *VolumeGroupResourcesOutput) GetDiskListOk() (*[]VolumeDiskResourceOutput, bool)`

GetDiskListOk returns a tuple with the DiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskList

`func (o *VolumeGroupResourcesOutput) SetDiskList(v []VolumeDiskResourceOutput)`

SetDiskList sets DiskList field to given value.

### HasDiskList

`func (o *VolumeGroupResourcesOutput) HasDiskList() bool`

HasDiskList returns a boolean if a field has been set.

### GetSizeBytes

`func (o *VolumeGroupResourcesOutput) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *VolumeGroupResourcesOutput) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *VolumeGroupResourcesOutput) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *VolumeGroupResourcesOutput) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetUsageType

`func (o *VolumeGroupResourcesOutput) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *VolumeGroupResourcesOutput) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *VolumeGroupResourcesOutput) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *VolumeGroupResourcesOutput) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesOutput) GetLoadBalanceVmAttachments() bool`

GetLoadBalanceVmAttachments returns the LoadBalanceVmAttachments field if non-nil, zero value otherwise.

### GetLoadBalanceVmAttachmentsOk

`func (o *VolumeGroupResourcesOutput) GetLoadBalanceVmAttachmentsOk() (*bool, bool)`

GetLoadBalanceVmAttachmentsOk returns a tuple with the LoadBalanceVmAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesOutput) SetLoadBalanceVmAttachments(v bool)`

SetLoadBalanceVmAttachments sets LoadBalanceVmAttachments field to given value.

### HasLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesOutput) HasLoadBalanceVmAttachments() bool`

HasLoadBalanceVmAttachments returns a boolean if a field has been set.

### GetIsHidden

`func (o *VolumeGroupResourcesOutput) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *VolumeGroupResourcesOutput) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *VolumeGroupResourcesOutput) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *VolumeGroupResourcesOutput) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetSizeMib

`func (o *VolumeGroupResourcesOutput) GetSizeMib() int64`

GetSizeMib returns the SizeMib field if non-nil, zero value otherwise.

### GetSizeMibOk

`func (o *VolumeGroupResourcesOutput) GetSizeMibOk() (*int64, bool)`

GetSizeMibOk returns a tuple with the SizeMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMib

`func (o *VolumeGroupResourcesOutput) SetSizeMib(v int64)`

SetSizeMib sets SizeMib field to given value.

### HasSizeMib

`func (o *VolumeGroupResourcesOutput) HasSizeMib() bool`

HasSizeMib returns a boolean if a field has been set.

### GetIscsiTargetPrefix

`func (o *VolumeGroupResourcesOutput) GetIscsiTargetPrefix() string`

GetIscsiTargetPrefix returns the IscsiTargetPrefix field if non-nil, zero value otherwise.

### GetIscsiTargetPrefixOk

`func (o *VolumeGroupResourcesOutput) GetIscsiTargetPrefixOk() (*string, bool)`

GetIscsiTargetPrefixOk returns a tuple with the IscsiTargetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiTargetPrefix

`func (o *VolumeGroupResourcesOutput) SetIscsiTargetPrefix(v string)`

SetIscsiTargetPrefix sets IscsiTargetPrefix field to given value.

### HasIscsiTargetPrefix

`func (o *VolumeGroupResourcesOutput) HasIscsiTargetPrefix() bool`

HasIscsiTargetPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


