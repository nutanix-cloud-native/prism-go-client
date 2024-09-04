# MhVmResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorSpecificId** | Pointer to **string** | The hypervisor specific ID of the VM.  | [optional] 
**HostReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**GuestTools** | Pointer to [**GuestToolsStatus**](GuestToolsStatus.md) |  | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the hypervisor the VM is hosted on.  | [optional] 
**StorageConfig** | Pointer to [**MhVmStorageConfigStatus**](MhVmStorageConfigStatus.md) |  | [optional] 
**ProtectionType** | Pointer to **string** | The type of protection applied on a VM. PD_PROTECTED indicates a VM protected using Prism Element. RULE_PROTECTED indicates a VM protected using Prism Central.  | [optional] 
**ParentReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**RecoveryPlanStateList** | Pointer to [**[]RecoveryPlanPolicyState**](RecoveryPlanPolicyState.md) | Status of the Recovery Plans associated with the VM. | [optional] 
**ProtectionPolicyState** | Pointer to [**ProtectionPolicyState**](ProtectionPolicyState.md) |  | [optional] 

## Methods

### NewMhVmResourcesDefStatus

`func NewMhVmResourcesDefStatus() *MhVmResourcesDefStatus`

NewMhVmResourcesDefStatus instantiates a new MhVmResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmResourcesDefStatusWithDefaults

`func NewMhVmResourcesDefStatusWithDefaults() *MhVmResourcesDefStatus`

NewMhVmResourcesDefStatusWithDefaults instantiates a new MhVmResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervisorSpecificId

`func (o *MhVmResourcesDefStatus) GetHypervisorSpecificId() string`

GetHypervisorSpecificId returns the HypervisorSpecificId field if non-nil, zero value otherwise.

### GetHypervisorSpecificIdOk

`func (o *MhVmResourcesDefStatus) GetHypervisorSpecificIdOk() (*string, bool)`

GetHypervisorSpecificIdOk returns a tuple with the HypervisorSpecificId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorSpecificId

`func (o *MhVmResourcesDefStatus) SetHypervisorSpecificId(v string)`

SetHypervisorSpecificId sets HypervisorSpecificId field to given value.

### HasHypervisorSpecificId

`func (o *MhVmResourcesDefStatus) HasHypervisorSpecificId() bool`

HasHypervisorSpecificId returns a boolean if a field has been set.

### GetHostReference

`func (o *MhVmResourcesDefStatus) GetHostReference() Reference`

GetHostReference returns the HostReference field if non-nil, zero value otherwise.

### GetHostReferenceOk

`func (o *MhVmResourcesDefStatus) GetHostReferenceOk() (*Reference, bool)`

GetHostReferenceOk returns a tuple with the HostReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostReference

`func (o *MhVmResourcesDefStatus) SetHostReference(v Reference)`

SetHostReference sets HostReference field to given value.

### HasHostReference

`func (o *MhVmResourcesDefStatus) HasHostReference() bool`

HasHostReference returns a boolean if a field has been set.

### GetGuestTools

`func (o *MhVmResourcesDefStatus) GetGuestTools() GuestToolsStatus`

GetGuestTools returns the GuestTools field if non-nil, zero value otherwise.

### GetGuestToolsOk

`func (o *MhVmResourcesDefStatus) GetGuestToolsOk() (*GuestToolsStatus, bool)`

GetGuestToolsOk returns a tuple with the GuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTools

`func (o *MhVmResourcesDefStatus) SetGuestTools(v GuestToolsStatus)`

SetGuestTools sets GuestTools field to given value.

### HasGuestTools

`func (o *MhVmResourcesDefStatus) HasGuestTools() bool`

HasGuestTools returns a boolean if a field has been set.

### GetHypervisorType

`func (o *MhVmResourcesDefStatus) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *MhVmResourcesDefStatus) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *MhVmResourcesDefStatus) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *MhVmResourcesDefStatus) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetStorageConfig

`func (o *MhVmResourcesDefStatus) GetStorageConfig() MhVmStorageConfigStatus`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *MhVmResourcesDefStatus) GetStorageConfigOk() (*MhVmStorageConfigStatus, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *MhVmResourcesDefStatus) SetStorageConfig(v MhVmStorageConfigStatus)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *MhVmResourcesDefStatus) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetProtectionType

`func (o *MhVmResourcesDefStatus) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *MhVmResourcesDefStatus) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *MhVmResourcesDefStatus) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *MhVmResourcesDefStatus) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetParentReference

`func (o *MhVmResourcesDefStatus) GetParentReference() Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MhVmResourcesDefStatus) GetParentReferenceOk() (*Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MhVmResourcesDefStatus) SetParentReference(v Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MhVmResourcesDefStatus) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetRecoveryPlanStateList

`func (o *MhVmResourcesDefStatus) GetRecoveryPlanStateList() []RecoveryPlanPolicyState`

GetRecoveryPlanStateList returns the RecoveryPlanStateList field if non-nil, zero value otherwise.

### GetRecoveryPlanStateListOk

`func (o *MhVmResourcesDefStatus) GetRecoveryPlanStateListOk() (*[]RecoveryPlanPolicyState, bool)`

GetRecoveryPlanStateListOk returns a tuple with the RecoveryPlanStateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPlanStateList

`func (o *MhVmResourcesDefStatus) SetRecoveryPlanStateList(v []RecoveryPlanPolicyState)`

SetRecoveryPlanStateList sets RecoveryPlanStateList field to given value.

### HasRecoveryPlanStateList

`func (o *MhVmResourcesDefStatus) HasRecoveryPlanStateList() bool`

HasRecoveryPlanStateList returns a boolean if a field has been set.

### GetProtectionPolicyState

`func (o *MhVmResourcesDefStatus) GetProtectionPolicyState() ProtectionPolicyState`

GetProtectionPolicyState returns the ProtectionPolicyState field if non-nil, zero value otherwise.

### GetProtectionPolicyStateOk

`func (o *MhVmResourcesDefStatus) GetProtectionPolicyStateOk() (*ProtectionPolicyState, bool)`

GetProtectionPolicyStateOk returns a tuple with the ProtectionPolicyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionPolicyState

`func (o *MhVmResourcesDefStatus) SetProtectionPolicyState(v ProtectionPolicyState)`

SetProtectionPolicyState sets ProtectionPolicyState field to given value.

### HasProtectionPolicyState

`func (o *MhVmResourcesDefStatus) HasProtectionPolicyState() bool`

HasProtectionPolicyState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


