# VmRecoveryPointsCloneInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideSpec** | Pointer to [**VmRecoveryPointsOverrideSpec**](VmRecoveryPointsOverrideSpec.md) |  | [optional] 
**VmRecoveryPointUuid** | Pointer to **string** | UUID of the new vm_recovery_point that will be created.  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**AvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 

## Methods

### NewVmRecoveryPointsCloneInput

`func NewVmRecoveryPointsCloneInput(availabilityZoneReference AvailabilityZoneReference, ) *VmRecoveryPointsCloneInput`

NewVmRecoveryPointsCloneInput instantiates a new VmRecoveryPointsCloneInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointsCloneInputWithDefaults

`func NewVmRecoveryPointsCloneInputWithDefaults() *VmRecoveryPointsCloneInput`

NewVmRecoveryPointsCloneInputWithDefaults instantiates a new VmRecoveryPointsCloneInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideSpec

`func (o *VmRecoveryPointsCloneInput) GetOverrideSpec() VmRecoveryPointsOverrideSpec`

GetOverrideSpec returns the OverrideSpec field if non-nil, zero value otherwise.

### GetOverrideSpecOk

`func (o *VmRecoveryPointsCloneInput) GetOverrideSpecOk() (*VmRecoveryPointsOverrideSpec, bool)`

GetOverrideSpecOk returns a tuple with the OverrideSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSpec

`func (o *VmRecoveryPointsCloneInput) SetOverrideSpec(v VmRecoveryPointsOverrideSpec)`

SetOverrideSpec sets OverrideSpec field to given value.

### HasOverrideSpec

`func (o *VmRecoveryPointsCloneInput) HasOverrideSpec() bool`

HasOverrideSpec returns a boolean if a field has been set.

### GetVmRecoveryPointUuid

`func (o *VmRecoveryPointsCloneInput) GetVmRecoveryPointUuid() string`

GetVmRecoveryPointUuid returns the VmRecoveryPointUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointUuidOk

`func (o *VmRecoveryPointsCloneInput) GetVmRecoveryPointUuidOk() (*string, bool)`

GetVmRecoveryPointUuidOk returns a tuple with the VmRecoveryPointUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointUuid

`func (o *VmRecoveryPointsCloneInput) SetVmRecoveryPointUuid(v string)`

SetVmRecoveryPointUuid sets VmRecoveryPointUuid field to given value.

### HasVmRecoveryPointUuid

`func (o *VmRecoveryPointsCloneInput) HasVmRecoveryPointUuid() bool`

HasVmRecoveryPointUuid returns a boolean if a field has been set.

### GetClusterReference

`func (o *VmRecoveryPointsCloneInput) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VmRecoveryPointsCloneInput) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VmRecoveryPointsCloneInput) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VmRecoveryPointsCloneInput) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *VmRecoveryPointsCloneInput) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VmRecoveryPointsCloneInput) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VmRecoveryPointsCloneInput) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


