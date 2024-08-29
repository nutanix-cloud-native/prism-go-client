# VolumeGroupResourcesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlashMode** | Pointer to **string** | Flash Mode, if enabled all volume disks of the VG will be pinned to SSD tier.  | [optional] 
**LoadBalanceVmAttachments** | Pointer to **bool** | Whether to enable volume group load balancing. | [optional] 
**CreatedBy** | Pointer to **string** | Service/user who created this volume group. | [optional] 
**IscsiTargetPrefix** | Pointer to **string** | iSCSI target prefix-name. | [optional] 
**ParentReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**SharingStatus** | Pointer to **string** | Whether the volume group can be shared across multiple iSCSI initiators.  | [optional] 
**AttachmentList** | Pointer to [**[]AttachmentReferenceInput**](AttachmentReferenceInput.md) | VMs attached to volume group. | [optional] 
**UsageType** | Pointer to **string** | Expected usage type for the volume group. | [optional] 
**TargetSecret** | Pointer to **string** | Target Secret in case of CHAP authentication. | [optional] 
**IsHidden** | Pointer to **bool** | Whether the VG is meant to be hidden or not. | [optional] 
**DiskList** | Pointer to [**[]VolumeDiskResourceInput**](VolumeDiskResourceInput.md) | Volume group disk specification. | [optional] 

## Methods

### NewVolumeGroupResourcesInput

`func NewVolumeGroupResourcesInput() *VolumeGroupResourcesInput`

NewVolumeGroupResourcesInput instantiates a new VolumeGroupResourcesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupResourcesInputWithDefaults

`func NewVolumeGroupResourcesInputWithDefaults() *VolumeGroupResourcesInput`

NewVolumeGroupResourcesInputWithDefaults instantiates a new VolumeGroupResourcesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlashMode

`func (o *VolumeGroupResourcesInput) GetFlashMode() string`

GetFlashMode returns the FlashMode field if non-nil, zero value otherwise.

### GetFlashModeOk

`func (o *VolumeGroupResourcesInput) GetFlashModeOk() (*string, bool)`

GetFlashModeOk returns a tuple with the FlashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashMode

`func (o *VolumeGroupResourcesInput) SetFlashMode(v string)`

SetFlashMode sets FlashMode field to given value.

### HasFlashMode

`func (o *VolumeGroupResourcesInput) HasFlashMode() bool`

HasFlashMode returns a boolean if a field has been set.

### GetLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesInput) GetLoadBalanceVmAttachments() bool`

GetLoadBalanceVmAttachments returns the LoadBalanceVmAttachments field if non-nil, zero value otherwise.

### GetLoadBalanceVmAttachmentsOk

`func (o *VolumeGroupResourcesInput) GetLoadBalanceVmAttachmentsOk() (*bool, bool)`

GetLoadBalanceVmAttachmentsOk returns a tuple with the LoadBalanceVmAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesInput) SetLoadBalanceVmAttachments(v bool)`

SetLoadBalanceVmAttachments sets LoadBalanceVmAttachments field to given value.

### HasLoadBalanceVmAttachments

`func (o *VolumeGroupResourcesInput) HasLoadBalanceVmAttachments() bool`

HasLoadBalanceVmAttachments returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeGroupResourcesInput) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeGroupResourcesInput) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeGroupResourcesInput) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeGroupResourcesInput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIscsiTargetPrefix

`func (o *VolumeGroupResourcesInput) GetIscsiTargetPrefix() string`

GetIscsiTargetPrefix returns the IscsiTargetPrefix field if non-nil, zero value otherwise.

### GetIscsiTargetPrefixOk

`func (o *VolumeGroupResourcesInput) GetIscsiTargetPrefixOk() (*string, bool)`

GetIscsiTargetPrefixOk returns a tuple with the IscsiTargetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiTargetPrefix

`func (o *VolumeGroupResourcesInput) SetIscsiTargetPrefix(v string)`

SetIscsiTargetPrefix sets IscsiTargetPrefix field to given value.

### HasIscsiTargetPrefix

`func (o *VolumeGroupResourcesInput) HasIscsiTargetPrefix() bool`

HasIscsiTargetPrefix returns a boolean if a field has been set.

### GetParentReference

`func (o *VolumeGroupResourcesInput) GetParentReference() Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *VolumeGroupResourcesInput) GetParentReferenceOk() (*Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *VolumeGroupResourcesInput) SetParentReference(v Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *VolumeGroupResourcesInput) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetSharingStatus

`func (o *VolumeGroupResourcesInput) GetSharingStatus() string`

GetSharingStatus returns the SharingStatus field if non-nil, zero value otherwise.

### GetSharingStatusOk

`func (o *VolumeGroupResourcesInput) GetSharingStatusOk() (*string, bool)`

GetSharingStatusOk returns a tuple with the SharingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingStatus

`func (o *VolumeGroupResourcesInput) SetSharingStatus(v string)`

SetSharingStatus sets SharingStatus field to given value.

### HasSharingStatus

`func (o *VolumeGroupResourcesInput) HasSharingStatus() bool`

HasSharingStatus returns a boolean if a field has been set.

### GetAttachmentList

`func (o *VolumeGroupResourcesInput) GetAttachmentList() []AttachmentReferenceInput`

GetAttachmentList returns the AttachmentList field if non-nil, zero value otherwise.

### GetAttachmentListOk

`func (o *VolumeGroupResourcesInput) GetAttachmentListOk() (*[]AttachmentReferenceInput, bool)`

GetAttachmentListOk returns a tuple with the AttachmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentList

`func (o *VolumeGroupResourcesInput) SetAttachmentList(v []AttachmentReferenceInput)`

SetAttachmentList sets AttachmentList field to given value.

### HasAttachmentList

`func (o *VolumeGroupResourcesInput) HasAttachmentList() bool`

HasAttachmentList returns a boolean if a field has been set.

### GetUsageType

`func (o *VolumeGroupResourcesInput) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *VolumeGroupResourcesInput) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *VolumeGroupResourcesInput) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *VolumeGroupResourcesInput) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetTargetSecret

`func (o *VolumeGroupResourcesInput) GetTargetSecret() string`

GetTargetSecret returns the TargetSecret field if non-nil, zero value otherwise.

### GetTargetSecretOk

`func (o *VolumeGroupResourcesInput) GetTargetSecretOk() (*string, bool)`

GetTargetSecretOk returns a tuple with the TargetSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecret

`func (o *VolumeGroupResourcesInput) SetTargetSecret(v string)`

SetTargetSecret sets TargetSecret field to given value.

### HasTargetSecret

`func (o *VolumeGroupResourcesInput) HasTargetSecret() bool`

HasTargetSecret returns a boolean if a field has been set.

### GetIsHidden

`func (o *VolumeGroupResourcesInput) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *VolumeGroupResourcesInput) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *VolumeGroupResourcesInput) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *VolumeGroupResourcesInput) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetDiskList

`func (o *VolumeGroupResourcesInput) GetDiskList() []VolumeDiskResourceInput`

GetDiskList returns the DiskList field if non-nil, zero value otherwise.

### GetDiskListOk

`func (o *VolumeGroupResourcesInput) GetDiskListOk() (*[]VolumeDiskResourceInput, bool)`

GetDiskListOk returns a tuple with the DiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskList

`func (o *VolumeGroupResourcesInput) SetDiskList(v []VolumeDiskResourceInput)`

SetDiskList sets DiskList field to given value.

### HasDiskList

`func (o *VolumeGroupResourcesInput) HasDiskList() bool`

HasDiskList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


