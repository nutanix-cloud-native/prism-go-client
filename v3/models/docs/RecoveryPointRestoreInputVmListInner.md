# RecoveryPointRestoreInputVmListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmRecoveryPointReference** | [**VmRecoveryPointReference**](VmRecoveryPointReference.md) |  | 
**MhVmSpec** | Pointer to [**MhVmSpecOverride**](MhVmSpecOverride.md) |  | [optional] 
**VmSpec** | Pointer to [**VmSpecOverride**](VmSpecOverride.md) |  | [optional] 
**Metadata** | Pointer to [**VmMetadataOverride**](VmMetadataOverride.md) |  | [optional] 

## Methods

### NewRecoveryPointRestoreInputVmListInner

`func NewRecoveryPointRestoreInputVmListInner(vmRecoveryPointReference VmRecoveryPointReference, ) *RecoveryPointRestoreInputVmListInner`

NewRecoveryPointRestoreInputVmListInner instantiates a new RecoveryPointRestoreInputVmListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPointRestoreInputVmListInnerWithDefaults

`func NewRecoveryPointRestoreInputVmListInnerWithDefaults() *RecoveryPointRestoreInputVmListInner`

NewRecoveryPointRestoreInputVmListInnerWithDefaults instantiates a new RecoveryPointRestoreInputVmListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmRecoveryPointReference

`func (o *RecoveryPointRestoreInputVmListInner) GetVmRecoveryPointReference() VmRecoveryPointReference`

GetVmRecoveryPointReference returns the VmRecoveryPointReference field if non-nil, zero value otherwise.

### GetVmRecoveryPointReferenceOk

`func (o *RecoveryPointRestoreInputVmListInner) GetVmRecoveryPointReferenceOk() (*VmRecoveryPointReference, bool)`

GetVmRecoveryPointReferenceOk returns a tuple with the VmRecoveryPointReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointReference

`func (o *RecoveryPointRestoreInputVmListInner) SetVmRecoveryPointReference(v VmRecoveryPointReference)`

SetVmRecoveryPointReference sets VmRecoveryPointReference field to given value.


### GetMhVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) GetMhVmSpec() MhVmSpecOverride`

GetMhVmSpec returns the MhVmSpec field if non-nil, zero value otherwise.

### GetMhVmSpecOk

`func (o *RecoveryPointRestoreInputVmListInner) GetMhVmSpecOk() (*MhVmSpecOverride, bool)`

GetMhVmSpecOk returns a tuple with the MhVmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMhVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) SetMhVmSpec(v MhVmSpecOverride)`

SetMhVmSpec sets MhVmSpec field to given value.

### HasMhVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) HasMhVmSpec() bool`

HasMhVmSpec returns a boolean if a field has been set.

### GetVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) GetVmSpec() VmSpecOverride`

GetVmSpec returns the VmSpec field if non-nil, zero value otherwise.

### GetVmSpecOk

`func (o *RecoveryPointRestoreInputVmListInner) GetVmSpecOk() (*VmSpecOverride, bool)`

GetVmSpecOk returns a tuple with the VmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) SetVmSpec(v VmSpecOverride)`

SetVmSpec sets VmSpec field to given value.

### HasVmSpec

`func (o *RecoveryPointRestoreInputVmListInner) HasVmSpec() bool`

HasVmSpec returns a boolean if a field has been set.

### GetMetadata

`func (o *RecoveryPointRestoreInputVmListInner) GetMetadata() VmMetadataOverride`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPointRestoreInputVmListInner) GetMetadataOk() (*VmMetadataOverride, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPointRestoreInputVmListInner) SetMetadata(v VmMetadataOverride)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RecoveryPointRestoreInputVmListInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


