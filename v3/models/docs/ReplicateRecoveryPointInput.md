# ReplicateRecoveryPointInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**TargetClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**TargetAvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 

## Methods

### NewReplicateRecoveryPointInput

`func NewReplicateRecoveryPointInput(sourceAvailabilityZoneReference AvailabilityZoneReference, targetAvailabilityZoneReference AvailabilityZoneReference, ) *ReplicateRecoveryPointInput`

NewReplicateRecoveryPointInput instantiates a new ReplicateRecoveryPointInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicateRecoveryPointInputWithDefaults

`func NewReplicateRecoveryPointInputWithDefaults() *ReplicateRecoveryPointInput`

NewReplicateRecoveryPointInputWithDefaults instantiates a new ReplicateRecoveryPointInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAvailabilityZoneReference

`func (o *ReplicateRecoveryPointInput) GetSourceAvailabilityZoneReference() AvailabilityZoneReference`

GetSourceAvailabilityZoneReference returns the SourceAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneReferenceOk

`func (o *ReplicateRecoveryPointInput) GetSourceAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetSourceAvailabilityZoneReferenceOk returns a tuple with the SourceAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneReference

`func (o *ReplicateRecoveryPointInput) SetSourceAvailabilityZoneReference(v AvailabilityZoneReference)`

SetSourceAvailabilityZoneReference sets SourceAvailabilityZoneReference field to given value.


### GetTargetClusterReference

`func (o *ReplicateRecoveryPointInput) GetTargetClusterReference() ClusterReference`

GetTargetClusterReference returns the TargetClusterReference field if non-nil, zero value otherwise.

### GetTargetClusterReferenceOk

`func (o *ReplicateRecoveryPointInput) GetTargetClusterReferenceOk() (*ClusterReference, bool)`

GetTargetClusterReferenceOk returns a tuple with the TargetClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterReference

`func (o *ReplicateRecoveryPointInput) SetTargetClusterReference(v ClusterReference)`

SetTargetClusterReference sets TargetClusterReference field to given value.

### HasTargetClusterReference

`func (o *ReplicateRecoveryPointInput) HasTargetClusterReference() bool`

HasTargetClusterReference returns a boolean if a field has been set.

### GetTargetAvailabilityZoneReference

`func (o *ReplicateRecoveryPointInput) GetTargetAvailabilityZoneReference() AvailabilityZoneReference`

GetTargetAvailabilityZoneReference returns the TargetAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetTargetAvailabilityZoneReferenceOk

`func (o *ReplicateRecoveryPointInput) GetTargetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetTargetAvailabilityZoneReferenceOk returns a tuple with the TargetAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAvailabilityZoneReference

`func (o *ReplicateRecoveryPointInput) SetTargetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetTargetAvailabilityZoneReference sets TargetAvailabilityZoneReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


