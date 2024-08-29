# RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | Type of authentication protocol to be used. | [optional] 
**TargetSecret** | Pointer to **string** | Authentication secret provided to iSCSI initiator for CHAP based authentication of the Volume Group.  | [optional] 
**VolumeGroupReference** | [**VolumeGroupReference**](VolumeGroupReference.md) |  | 

## Methods

### NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner

`func NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner(volumeGroupReference VolumeGroupReference, ) *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner`

NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner instantiates a new RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInnerWithDefaults

`func NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInnerWithDefaults() *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner`

NewRecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInnerWithDefaults instantiates a new RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetTargetSecret

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetTargetSecret() string`

GetTargetSecret returns the TargetSecret field if non-nil, zero value otherwise.

### GetTargetSecretOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetTargetSecretOk() (*string, bool)`

GetTargetSecretOk returns a tuple with the TargetSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSecret

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) SetTargetSecret(v string)`

SetTargetSecret sets TargetSecret field to given value.

### HasTargetSecret

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) HasTargetSecret() bool`

HasTargetSecret returns a boolean if a field has been set.

### GetVolumeGroupReference

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetVolumeGroupReference() VolumeGroupReference`

GetVolumeGroupReference returns the VolumeGroupReference field if non-nil, zero value otherwise.

### GetVolumeGroupReferenceOk

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) GetVolumeGroupReferenceOk() (*VolumeGroupReference, bool)`

GetVolumeGroupReferenceOk returns a tuple with the VolumeGroupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupReference

`func (o *RecoveryPlanVolumeGroupRecoveryInfoVolumeGroupConfigInfoListInner) SetVolumeGroupReference(v VolumeGroupReference)`

SetVolumeGroupReference sets VolumeGroupReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


