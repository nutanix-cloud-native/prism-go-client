# RecoveryPlanVolumeGroupRecoveryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeGroupConfigInfoList** | Pointer to [**[]RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner**](RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner.md) | List of configuration information for each Volume Group specified explicitly or in the category filter. Each entry in this list will contain reference to the Volume Group, authentication protocol and the target secret to be used for authenticating the Volume Group.  | [optional] 
**CategoryFilter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 
**VolumeGroupReference** | Pointer to [**VolumeGroupReference**](VolumeGroupReference.md) |  | [optional] 

## Methods

### NewRecoveryPlanVolumeGroupRecoveryInfo

`func NewRecoveryPlanVolumeGroupRecoveryInfo() *RecoveryPlanVolumeGroupRecoveryInfo`

NewRecoveryPlanVolumeGroupRecoveryInfo instantiates a new RecoveryPlanVolumeGroupRecoveryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanVolumeGroupRecoveryInfoWithDefaults

`func NewRecoveryPlanVolumeGroupRecoveryInfoWithDefaults() *RecoveryPlanVolumeGroupRecoveryInfo`

NewRecoveryPlanVolumeGroupRecoveryInfoWithDefaults instantiates a new RecoveryPlanVolumeGroupRecoveryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeGroupConfigInfoList

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetVolumeGroupConfigInfoList() []RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner`

GetVolumeGroupConfigInfoList returns the VolumeGroupConfigInfoList field if non-nil, zero value otherwise.

### GetVolumeGroupConfigInfoListOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetVolumeGroupConfigInfoListOk() (*[]RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner, bool)`

GetVolumeGroupConfigInfoListOk returns a tuple with the VolumeGroupConfigInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupConfigInfoList

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) SetVolumeGroupConfigInfoList(v []RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner)`

SetVolumeGroupConfigInfoList sets VolumeGroupConfigInfoList field to given value.

### HasVolumeGroupConfigInfoList

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) HasVolumeGroupConfigInfoList() bool`

HasVolumeGroupConfigInfoList returns a boolean if a field has been set.

### GetCategoryFilter

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetCategoryFilter() CategoryFilter`

GetCategoryFilter returns the CategoryFilter field if non-nil, zero value otherwise.

### GetCategoryFilterOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetCategoryFilterOk() (*CategoryFilter, bool)`

GetCategoryFilterOk returns a tuple with the CategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFilter

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) SetCategoryFilter(v CategoryFilter)`

SetCategoryFilter sets CategoryFilter field to given value.

### HasCategoryFilter

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) HasCategoryFilter() bool`

HasCategoryFilter returns a boolean if a field has been set.

### GetVolumeGroupReference

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetVolumeGroupReference() VolumeGroupReference`

GetVolumeGroupReference returns the VolumeGroupReference field if non-nil, zero value otherwise.

### GetVolumeGroupReferenceOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) GetVolumeGroupReferenceOk() (*VolumeGroupReference, bool)`

GetVolumeGroupReferenceOk returns a tuple with the VolumeGroupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupReference

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) SetVolumeGroupReference(v VolumeGroupReference)`

SetVolumeGroupReference sets VolumeGroupReference field to given value.

### HasVolumeGroupReference

`func (o *RecoveryPlanVolumeGroupRecoveryInfo) HasVolumeGroupReference() bool`

HasVolumeGroupReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


