# RecoveryPointResources1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsistencyGroupUuid** | Pointer to **string** | This field is same for all the entities (irrespective of kind) that were snapshotted together.  | [optional] 
**SourceAvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**SourceClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Recoverability** | Pointer to [**RecoveryPointResources1Recoverability**](RecoveryPointResources1Recoverability.md) |  | [optional] 
**ParentVmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**RecoveryPointType** | Pointer to **string** | Crash consistent or Application Consistent recovery point | [optional] 
**VmSpec** | Pointer to [**Vm**](Vm.md) |  | [optional] 
**VmRecoveryPointLocationAgnosticUuid** | Pointer to **string** | Location agnostic UUID of the recovery point. If a recovery point is replicated to a different clusters, then all the instances of same recovery point will share this UUID.  | [optional] 
**CreationTime** | Pointer to **string** | The time when the the recovery point is created. This is in internet date/time format (RFC 3339). For example, 1985-04-12T23:20:50.52Z, this represents 20 minutes and 50.52 seconds after the 23rd hour of April 12th, 1985 in UTC.  | [optional] 
**VmMetadata** | Pointer to [**VmMetadata**](VmMetadata.md) |  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | The time when this recovery point expires and will be garbage collected. This is in internet date/time format (RFC 3339). For example, 1985-04-12T23:20:50.52Z, this represents 20 minutes and 50.52 seconds after the 23rd hour of April 12th, 1985 in UTC. If not set, then the recovery point never expires.  | [optional] 

## Methods

### NewRecoveryPointResources1

`func NewRecoveryPointResources1() *RecoveryPointResources1`

NewRecoveryPointResources1 instantiates a new RecoveryPointResources1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPointResources1WithDefaults

`func NewRecoveryPointResources1WithDefaults() *RecoveryPointResources1`

NewRecoveryPointResources1WithDefaults instantiates a new RecoveryPointResources1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsistencyGroupUuid

`func (o *RecoveryPointResources1) GetConsistencyGroupUuid() string`

GetConsistencyGroupUuid returns the ConsistencyGroupUuid field if non-nil, zero value otherwise.

### GetConsistencyGroupUuidOk

`func (o *RecoveryPointResources1) GetConsistencyGroupUuidOk() (*string, bool)`

GetConsistencyGroupUuidOk returns a tuple with the ConsistencyGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistencyGroupUuid

`func (o *RecoveryPointResources1) SetConsistencyGroupUuid(v string)`

SetConsistencyGroupUuid sets ConsistencyGroupUuid field to given value.

### HasConsistencyGroupUuid

`func (o *RecoveryPointResources1) HasConsistencyGroupUuid() bool`

HasConsistencyGroupUuid returns a boolean if a field has been set.

### GetSourceAvailabilityZoneReference

`func (o *RecoveryPointResources1) GetSourceAvailabilityZoneReference() AvailabilityZoneReference`

GetSourceAvailabilityZoneReference returns the SourceAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneReferenceOk

`func (o *RecoveryPointResources1) GetSourceAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetSourceAvailabilityZoneReferenceOk returns a tuple with the SourceAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneReference

`func (o *RecoveryPointResources1) SetSourceAvailabilityZoneReference(v AvailabilityZoneReference)`

SetSourceAvailabilityZoneReference sets SourceAvailabilityZoneReference field to given value.

### HasSourceAvailabilityZoneReference

`func (o *RecoveryPointResources1) HasSourceAvailabilityZoneReference() bool`

HasSourceAvailabilityZoneReference returns a boolean if a field has been set.

### GetSourceClusterReference

`func (o *RecoveryPointResources1) GetSourceClusterReference() ClusterReference`

GetSourceClusterReference returns the SourceClusterReference field if non-nil, zero value otherwise.

### GetSourceClusterReferenceOk

`func (o *RecoveryPointResources1) GetSourceClusterReferenceOk() (*ClusterReference, bool)`

GetSourceClusterReferenceOk returns a tuple with the SourceClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterReference

`func (o *RecoveryPointResources1) SetSourceClusterReference(v ClusterReference)`

SetSourceClusterReference sets SourceClusterReference field to given value.

### HasSourceClusterReference

`func (o *RecoveryPointResources1) HasSourceClusterReference() bool`

HasSourceClusterReference returns a boolean if a field has been set.

### GetRecoverability

`func (o *RecoveryPointResources1) GetRecoverability() RecoveryPointResources1Recoverability`

GetRecoverability returns the Recoverability field if non-nil, zero value otherwise.

### GetRecoverabilityOk

`func (o *RecoveryPointResources1) GetRecoverabilityOk() (*RecoveryPointResources1Recoverability, bool)`

GetRecoverabilityOk returns a tuple with the Recoverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverability

`func (o *RecoveryPointResources1) SetRecoverability(v RecoveryPointResources1Recoverability)`

SetRecoverability sets Recoverability field to given value.

### HasRecoverability

`func (o *RecoveryPointResources1) HasRecoverability() bool`

HasRecoverability returns a boolean if a field has been set.

### GetParentVmReference

`func (o *RecoveryPointResources1) GetParentVmReference() VmReference`

GetParentVmReference returns the ParentVmReference field if non-nil, zero value otherwise.

### GetParentVmReferenceOk

`func (o *RecoveryPointResources1) GetParentVmReferenceOk() (*VmReference, bool)`

GetParentVmReferenceOk returns a tuple with the ParentVmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmReference

`func (o *RecoveryPointResources1) SetParentVmReference(v VmReference)`

SetParentVmReference sets ParentVmReference field to given value.

### HasParentVmReference

`func (o *RecoveryPointResources1) HasParentVmReference() bool`

HasParentVmReference returns a boolean if a field has been set.

### GetRecoveryPointType

`func (o *RecoveryPointResources1) GetRecoveryPointType() string`

GetRecoveryPointType returns the RecoveryPointType field if non-nil, zero value otherwise.

### GetRecoveryPointTypeOk

`func (o *RecoveryPointResources1) GetRecoveryPointTypeOk() (*string, bool)`

GetRecoveryPointTypeOk returns a tuple with the RecoveryPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointType

`func (o *RecoveryPointResources1) SetRecoveryPointType(v string)`

SetRecoveryPointType sets RecoveryPointType field to given value.

### HasRecoveryPointType

`func (o *RecoveryPointResources1) HasRecoveryPointType() bool`

HasRecoveryPointType returns a boolean if a field has been set.

### GetVmSpec

`func (o *RecoveryPointResources1) GetVmSpec() Vm`

GetVmSpec returns the VmSpec field if non-nil, zero value otherwise.

### GetVmSpecOk

`func (o *RecoveryPointResources1) GetVmSpecOk() (*Vm, bool)`

GetVmSpecOk returns a tuple with the VmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpec

`func (o *RecoveryPointResources1) SetVmSpec(v Vm)`

SetVmSpec sets VmSpec field to given value.

### HasVmSpec

`func (o *RecoveryPointResources1) HasVmSpec() bool`

HasVmSpec returns a boolean if a field has been set.

### GetVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources1) GetVmRecoveryPointLocationAgnosticUuid() string`

GetVmRecoveryPointLocationAgnosticUuid returns the VmRecoveryPointLocationAgnosticUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointLocationAgnosticUuidOk

`func (o *RecoveryPointResources1) GetVmRecoveryPointLocationAgnosticUuidOk() (*string, bool)`

GetVmRecoveryPointLocationAgnosticUuidOk returns a tuple with the VmRecoveryPointLocationAgnosticUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources1) SetVmRecoveryPointLocationAgnosticUuid(v string)`

SetVmRecoveryPointLocationAgnosticUuid sets VmRecoveryPointLocationAgnosticUuid field to given value.

### HasVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources1) HasVmRecoveryPointLocationAgnosticUuid() bool`

HasVmRecoveryPointLocationAgnosticUuid returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecoveryPointResources1) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecoveryPointResources1) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecoveryPointResources1) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecoveryPointResources1) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetVmMetadata

`func (o *RecoveryPointResources1) GetVmMetadata() VmMetadata`

GetVmMetadata returns the VmMetadata field if non-nil, zero value otherwise.

### GetVmMetadataOk

`func (o *RecoveryPointResources1) GetVmMetadataOk() (*VmMetadata, bool)`

GetVmMetadataOk returns a tuple with the VmMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMetadata

`func (o *RecoveryPointResources1) SetVmMetadata(v VmMetadata)`

SetVmMetadata sets VmMetadata field to given value.

### HasVmMetadata

`func (o *RecoveryPointResources1) HasVmMetadata() bool`

HasVmMetadata returns a boolean if a field has been set.

### GetExpirationTime

`func (o *RecoveryPointResources1) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *RecoveryPointResources1) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *RecoveryPointResources1) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *RecoveryPointResources1) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


