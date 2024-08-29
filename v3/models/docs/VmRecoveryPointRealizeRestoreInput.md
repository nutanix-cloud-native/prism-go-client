# VmRecoveryPointRealizeRestoreInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | UUID of the cluster where recovery point is present. This needs to be provided in case of Self AZ restore.  | [optional] 
**RestoreTargetTime** | **time.Time** | The target time to which the user wishes to restore to. This is in internet date/time format (RFC 3339). This field is explicitly used for time based restore in high frequency snapshotting and will be ignored for a regular restore.  | 
**VmOverrideSpec** | Pointer to [**VmRestoreOverrideSpec**](VmRestoreOverrideSpec.md) |  | [optional] 
**VmUuid** | **string** | UUID of the vm entity for which the recovery point is to be restored.  | 
**Metadata** | Pointer to [**MetadataOfTheRestoredVm**](MetadataOfTheRestoredVm.md) |  | [optional] 

## Methods

### NewVmRecoveryPointRealizeRestoreInput

`func NewVmRecoveryPointRealizeRestoreInput(restoreTargetTime time.Time, vmUuid string, ) *VmRecoveryPointRealizeRestoreInput`

NewVmRecoveryPointRealizeRestoreInput instantiates a new VmRecoveryPointRealizeRestoreInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointRealizeRestoreInputWithDefaults

`func NewVmRecoveryPointRealizeRestoreInputWithDefaults() *VmRecoveryPointRealizeRestoreInput`

NewVmRecoveryPointRealizeRestoreInputWithDefaults instantiates a new VmRecoveryPointRealizeRestoreInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *VmRecoveryPointRealizeRestoreInput) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *VmRecoveryPointRealizeRestoreInput) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *VmRecoveryPointRealizeRestoreInput) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *VmRecoveryPointRealizeRestoreInput) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetRestoreTargetTime

`func (o *VmRecoveryPointRealizeRestoreInput) GetRestoreTargetTime() time.Time`

GetRestoreTargetTime returns the RestoreTargetTime field if non-nil, zero value otherwise.

### GetRestoreTargetTimeOk

`func (o *VmRecoveryPointRealizeRestoreInput) GetRestoreTargetTimeOk() (*time.Time, bool)`

GetRestoreTargetTimeOk returns a tuple with the RestoreTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTargetTime

`func (o *VmRecoveryPointRealizeRestoreInput) SetRestoreTargetTime(v time.Time)`

SetRestoreTargetTime sets RestoreTargetTime field to given value.


### GetVmOverrideSpec

`func (o *VmRecoveryPointRealizeRestoreInput) GetVmOverrideSpec() VmRestoreOverrideSpec`

GetVmOverrideSpec returns the VmOverrideSpec field if non-nil, zero value otherwise.

### GetVmOverrideSpecOk

`func (o *VmRecoveryPointRealizeRestoreInput) GetVmOverrideSpecOk() (*VmRestoreOverrideSpec, bool)`

GetVmOverrideSpecOk returns a tuple with the VmOverrideSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOverrideSpec

`func (o *VmRecoveryPointRealizeRestoreInput) SetVmOverrideSpec(v VmRestoreOverrideSpec)`

SetVmOverrideSpec sets VmOverrideSpec field to given value.

### HasVmOverrideSpec

`func (o *VmRecoveryPointRealizeRestoreInput) HasVmOverrideSpec() bool`

HasVmOverrideSpec returns a boolean if a field has been set.

### GetVmUuid

`func (o *VmRecoveryPointRealizeRestoreInput) GetVmUuid() string`

GetVmUuid returns the VmUuid field if non-nil, zero value otherwise.

### GetVmUuidOk

`func (o *VmRecoveryPointRealizeRestoreInput) GetVmUuidOk() (*string, bool)`

GetVmUuidOk returns a tuple with the VmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUuid

`func (o *VmRecoveryPointRealizeRestoreInput) SetVmUuid(v string)`

SetVmUuid sets VmUuid field to given value.


### GetMetadata

`func (o *VmRecoveryPointRealizeRestoreInput) GetMetadata() MetadataOfTheRestoredVm`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointRealizeRestoreInput) GetMetadataOk() (*MetadataOfTheRestoredVm, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointRealizeRestoreInput) SetMetadata(v MetadataOfTheRestoredVm)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VmRecoveryPointRealizeRestoreInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


