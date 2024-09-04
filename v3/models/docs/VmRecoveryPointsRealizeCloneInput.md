# VmRecoveryPointsRealizeCloneInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmUuid** | **string** | UUID of the vm entity for which the recovery point is to be cloned.  | 
**AvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**CloneTargetTime** | **time.Time** | The target time of the state that the user wishes to clone to a target site. This is in internet date/time format (RFC 3339). This field is explicitly used for time based replication in high frequency snapshotting and will be ignored for a regular restore.  | 
**VmRecoveryPointUuid** | Pointer to **string** | UUID of the new vm_recovery_point that will be created.  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**OverrideSpec** | Pointer to [**VmRecoveryPointsOverrideSpec**](VmRecoveryPointsOverrideSpec.md) |  | [optional] 

## Methods

### NewVmRecoveryPointsRealizeCloneInput

`func NewVmRecoveryPointsRealizeCloneInput(vmUuid string, availabilityZoneReference AvailabilityZoneReference, cloneTargetTime time.Time, ) *VmRecoveryPointsRealizeCloneInput`

NewVmRecoveryPointsRealizeCloneInput instantiates a new VmRecoveryPointsRealizeCloneInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointsRealizeCloneInputWithDefaults

`func NewVmRecoveryPointsRealizeCloneInputWithDefaults() *VmRecoveryPointsRealizeCloneInput`

NewVmRecoveryPointsRealizeCloneInputWithDefaults instantiates a new VmRecoveryPointsRealizeCloneInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmUuid

`func (o *VmRecoveryPointsRealizeCloneInput) GetVmUuid() string`

GetVmUuid returns the VmUuid field if non-nil, zero value otherwise.

### GetVmUuidOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetVmUuidOk() (*string, bool)`

GetVmUuidOk returns a tuple with the VmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUuid

`func (o *VmRecoveryPointsRealizeCloneInput) SetVmUuid(v string)`

SetVmUuid sets VmUuid field to given value.


### GetAvailabilityZoneReference

`func (o *VmRecoveryPointsRealizeCloneInput) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VmRecoveryPointsRealizeCloneInput) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.


### GetCloneTargetTime

`func (o *VmRecoveryPointsRealizeCloneInput) GetCloneTargetTime() time.Time`

GetCloneTargetTime returns the CloneTargetTime field if non-nil, zero value otherwise.

### GetCloneTargetTimeOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetCloneTargetTimeOk() (*time.Time, bool)`

GetCloneTargetTimeOk returns a tuple with the CloneTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneTargetTime

`func (o *VmRecoveryPointsRealizeCloneInput) SetCloneTargetTime(v time.Time)`

SetCloneTargetTime sets CloneTargetTime field to given value.


### GetVmRecoveryPointUuid

`func (o *VmRecoveryPointsRealizeCloneInput) GetVmRecoveryPointUuid() string`

GetVmRecoveryPointUuid returns the VmRecoveryPointUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointUuidOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetVmRecoveryPointUuidOk() (*string, bool)`

GetVmRecoveryPointUuidOk returns a tuple with the VmRecoveryPointUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointUuid

`func (o *VmRecoveryPointsRealizeCloneInput) SetVmRecoveryPointUuid(v string)`

SetVmRecoveryPointUuid sets VmRecoveryPointUuid field to given value.

### HasVmRecoveryPointUuid

`func (o *VmRecoveryPointsRealizeCloneInput) HasVmRecoveryPointUuid() bool`

HasVmRecoveryPointUuid returns a boolean if a field has been set.

### GetClusterReference

`func (o *VmRecoveryPointsRealizeCloneInput) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VmRecoveryPointsRealizeCloneInput) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VmRecoveryPointsRealizeCloneInput) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetOverrideSpec

`func (o *VmRecoveryPointsRealizeCloneInput) GetOverrideSpec() VmRecoveryPointsOverrideSpec`

GetOverrideSpec returns the OverrideSpec field if non-nil, zero value otherwise.

### GetOverrideSpecOk

`func (o *VmRecoveryPointsRealizeCloneInput) GetOverrideSpecOk() (*VmRecoveryPointsOverrideSpec, bool)`

GetOverrideSpecOk returns a tuple with the OverrideSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSpec

`func (o *VmRecoveryPointsRealizeCloneInput) SetOverrideSpec(v VmRecoveryPointsOverrideSpec)`

SetOverrideSpec sets OverrideSpec field to given value.

### HasOverrideSpec

`func (o *VmRecoveryPointsRealizeCloneInput) HasOverrideSpec() bool`

HasOverrideSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


