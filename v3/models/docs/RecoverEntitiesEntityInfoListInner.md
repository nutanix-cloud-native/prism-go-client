# RecoverEntitiesEntityInfoListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmPowerState** | Pointer to **string** | Power state of the VM(s) after recovery. | [optional] 
**AnyEntityReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**VolumeGroupAttachmentList** | Pointer to [**[]RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInner**](RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInner.md) | List containing the VMs to Volume Group attachment information.  | [optional] 
**Categories** | Pointer to **map[string]string** | Categories for filtering entities. | [optional] 
**ScriptList** | Pointer to [**[]RecoveryPlanScriptConfig**](RecoveryPlanScriptConfig.md) | List of scripts to be executed inside the guest VMs after recovery.  | [optional] 

## Methods

### NewRecoverEntitiesEntityInfoListInner

`func NewRecoverEntitiesEntityInfoListInner() *RecoverEntitiesEntityInfoListInner`

NewRecoverEntitiesEntityInfoListInner instantiates a new RecoverEntitiesEntityInfoListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverEntitiesEntityInfoListInnerWithDefaults

`func NewRecoverEntitiesEntityInfoListInnerWithDefaults() *RecoverEntitiesEntityInfoListInner`

NewRecoverEntitiesEntityInfoListInnerWithDefaults instantiates a new RecoverEntitiesEntityInfoListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmPowerState

`func (o *RecoverEntitiesEntityInfoListInner) GetVmPowerState() string`

GetVmPowerState returns the VmPowerState field if non-nil, zero value otherwise.

### GetVmPowerStateOk

`func (o *RecoverEntitiesEntityInfoListInner) GetVmPowerStateOk() (*string, bool)`

GetVmPowerStateOk returns a tuple with the VmPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPowerState

`func (o *RecoverEntitiesEntityInfoListInner) SetVmPowerState(v string)`

SetVmPowerState sets VmPowerState field to given value.

### HasVmPowerState

`func (o *RecoverEntitiesEntityInfoListInner) HasVmPowerState() bool`

HasVmPowerState returns a boolean if a field has been set.

### GetAnyEntityReference

`func (o *RecoverEntitiesEntityInfoListInner) GetAnyEntityReference() Reference`

GetAnyEntityReference returns the AnyEntityReference field if non-nil, zero value otherwise.

### GetAnyEntityReferenceOk

`func (o *RecoverEntitiesEntityInfoListInner) GetAnyEntityReferenceOk() (*Reference, bool)`

GetAnyEntityReferenceOk returns a tuple with the AnyEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEntityReference

`func (o *RecoverEntitiesEntityInfoListInner) SetAnyEntityReference(v Reference)`

SetAnyEntityReference sets AnyEntityReference field to given value.

### HasAnyEntityReference

`func (o *RecoverEntitiesEntityInfoListInner) HasAnyEntityReference() bool`

HasAnyEntityReference returns a boolean if a field has been set.

### GetVolumeGroupAttachmentList

`func (o *RecoverEntitiesEntityInfoListInner) GetVolumeGroupAttachmentList() []RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInner`

GetVolumeGroupAttachmentList returns the VolumeGroupAttachmentList field if non-nil, zero value otherwise.

### GetVolumeGroupAttachmentListOk

`func (o *RecoverEntitiesEntityInfoListInner) GetVolumeGroupAttachmentListOk() (*[]RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInner, bool)`

GetVolumeGroupAttachmentListOk returns a tuple with the VolumeGroupAttachmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupAttachmentList

`func (o *RecoverEntitiesEntityInfoListInner) SetVolumeGroupAttachmentList(v []RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInner)`

SetVolumeGroupAttachmentList sets VolumeGroupAttachmentList field to given value.

### HasVolumeGroupAttachmentList

`func (o *RecoverEntitiesEntityInfoListInner) HasVolumeGroupAttachmentList() bool`

HasVolumeGroupAttachmentList returns a boolean if a field has been set.

### GetCategories

`func (o *RecoverEntitiesEntityInfoListInner) GetCategories() map[string]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *RecoverEntitiesEntityInfoListInner) GetCategoriesOk() (*map[string]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *RecoverEntitiesEntityInfoListInner) SetCategories(v map[string]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *RecoverEntitiesEntityInfoListInner) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetScriptList

`func (o *RecoverEntitiesEntityInfoListInner) GetScriptList() []RecoveryPlanScriptConfig`

GetScriptList returns the ScriptList field if non-nil, zero value otherwise.

### GetScriptListOk

`func (o *RecoverEntitiesEntityInfoListInner) GetScriptListOk() (*[]RecoveryPlanScriptConfig, bool)`

GetScriptListOk returns a tuple with the ScriptList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptList

`func (o *RecoverEntitiesEntityInfoListInner) SetScriptList(v []RecoveryPlanScriptConfig)`

SetScriptList sets ScriptList field to given value.

### HasScriptList

`func (o *RecoverEntitiesEntityInfoListInner) HasScriptList() bool`

HasScriptList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


