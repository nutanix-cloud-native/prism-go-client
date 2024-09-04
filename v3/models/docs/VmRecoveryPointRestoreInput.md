# VmRecoveryPointRestoreInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmOverrideSpec** | Pointer to [**VmRestoreOverrideSpec**](VmRestoreOverrideSpec.md) |  | [optional] 
**Metadata** | Pointer to [**MetadataOfTheRestoredVm**](MetadataOfTheRestoredVm.md) |  | [optional] 

## Methods

### NewVmRecoveryPointRestoreInput

`func NewVmRecoveryPointRestoreInput() *VmRecoveryPointRestoreInput`

NewVmRecoveryPointRestoreInput instantiates a new VmRecoveryPointRestoreInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointRestoreInputWithDefaults

`func NewVmRecoveryPointRestoreInputWithDefaults() *VmRecoveryPointRestoreInput`

NewVmRecoveryPointRestoreInputWithDefaults instantiates a new VmRecoveryPointRestoreInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmOverrideSpec

`func (o *VmRecoveryPointRestoreInput) GetVmOverrideSpec() VmRestoreOverrideSpec`

GetVmOverrideSpec returns the VmOverrideSpec field if non-nil, zero value otherwise.

### GetVmOverrideSpecOk

`func (o *VmRecoveryPointRestoreInput) GetVmOverrideSpecOk() (*VmRestoreOverrideSpec, bool)`

GetVmOverrideSpecOk returns a tuple with the VmOverrideSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOverrideSpec

`func (o *VmRecoveryPointRestoreInput) SetVmOverrideSpec(v VmRestoreOverrideSpec)`

SetVmOverrideSpec sets VmOverrideSpec field to given value.

### HasVmOverrideSpec

`func (o *VmRecoveryPointRestoreInput) HasVmOverrideSpec() bool`

HasVmOverrideSpec returns a boolean if a field has been set.

### GetMetadata

`func (o *VmRecoveryPointRestoreInput) GetMetadata() MetadataOfTheRestoredVm`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointRestoreInput) GetMetadataOk() (*MetadataOfTheRestoredVm, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointRestoreInput) SetMetadata(v MetadataOfTheRestoredVm)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VmRecoveryPointRestoreInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


