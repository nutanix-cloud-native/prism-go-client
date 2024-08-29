# RecoveryPointResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**SourceAvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**ParentVmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**RecoveryPointType** | Pointer to **string** | Crash consistent or Application Consistent recovery point | [optional] 
**CreationTime** | Pointer to **time.Time** | The time when the the recovery point is created. This is in internet date/time format (RFC 3339). For example, 1985-04-12T23:20:50.52Z, this represents 20 minutes and 50.52 seconds after the 23rd hour of April 12th, 1985 in UTC.  | [optional] 
**VmRecoveryPointLocationAgnosticUuid** | Pointer to **string** | Location agnostic UUID of the recovery point. If a recovery point is replicated to a different clusters, then all the instances of same recovery point will share this UUID.  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | The time when this recovery point expires and will be garbage collected. This is in internet date/time format (RFC 3339). For example, 1985-04-12T23:20:50.52Z, this represents 20 minutes and 50.52 seconds after the 23rd hour of April 12th, 1985 in UTC. If not set, then the recovery point never expires.  | [optional] 

## Methods

### NewRecoveryPointResources

`func NewRecoveryPointResources() *RecoveryPointResources`

NewRecoveryPointResources instantiates a new RecoveryPointResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPointResourcesWithDefaults

`func NewRecoveryPointResourcesWithDefaults() *RecoveryPointResources`

NewRecoveryPointResourcesWithDefaults instantiates a new RecoveryPointResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceClusterReference

`func (o *RecoveryPointResources) GetSourceClusterReference() ClusterReference`

GetSourceClusterReference returns the SourceClusterReference field if non-nil, zero value otherwise.

### GetSourceClusterReferenceOk

`func (o *RecoveryPointResources) GetSourceClusterReferenceOk() (*ClusterReference, bool)`

GetSourceClusterReferenceOk returns a tuple with the SourceClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterReference

`func (o *RecoveryPointResources) SetSourceClusterReference(v ClusterReference)`

SetSourceClusterReference sets SourceClusterReference field to given value.

### HasSourceClusterReference

`func (o *RecoveryPointResources) HasSourceClusterReference() bool`

HasSourceClusterReference returns a boolean if a field has been set.

### GetSourceAvailabilityZoneReference

`func (o *RecoveryPointResources) GetSourceAvailabilityZoneReference() AvailabilityZoneReference`

GetSourceAvailabilityZoneReference returns the SourceAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneReferenceOk

`func (o *RecoveryPointResources) GetSourceAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetSourceAvailabilityZoneReferenceOk returns a tuple with the SourceAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneReference

`func (o *RecoveryPointResources) SetSourceAvailabilityZoneReference(v AvailabilityZoneReference)`

SetSourceAvailabilityZoneReference sets SourceAvailabilityZoneReference field to given value.

### HasSourceAvailabilityZoneReference

`func (o *RecoveryPointResources) HasSourceAvailabilityZoneReference() bool`

HasSourceAvailabilityZoneReference returns a boolean if a field has been set.

### GetParentVmReference

`func (o *RecoveryPointResources) GetParentVmReference() VmReference`

GetParentVmReference returns the ParentVmReference field if non-nil, zero value otherwise.

### GetParentVmReferenceOk

`func (o *RecoveryPointResources) GetParentVmReferenceOk() (*VmReference, bool)`

GetParentVmReferenceOk returns a tuple with the ParentVmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmReference

`func (o *RecoveryPointResources) SetParentVmReference(v VmReference)`

SetParentVmReference sets ParentVmReference field to given value.

### HasParentVmReference

`func (o *RecoveryPointResources) HasParentVmReference() bool`

HasParentVmReference returns a boolean if a field has been set.

### GetRecoveryPointType

`func (o *RecoveryPointResources) GetRecoveryPointType() string`

GetRecoveryPointType returns the RecoveryPointType field if non-nil, zero value otherwise.

### GetRecoveryPointTypeOk

`func (o *RecoveryPointResources) GetRecoveryPointTypeOk() (*string, bool)`

GetRecoveryPointTypeOk returns a tuple with the RecoveryPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointType

`func (o *RecoveryPointResources) SetRecoveryPointType(v string)`

SetRecoveryPointType sets RecoveryPointType field to given value.

### HasRecoveryPointType

`func (o *RecoveryPointResources) HasRecoveryPointType() bool`

HasRecoveryPointType returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecoveryPointResources) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecoveryPointResources) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecoveryPointResources) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecoveryPointResources) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources) GetVmRecoveryPointLocationAgnosticUuid() string`

GetVmRecoveryPointLocationAgnosticUuid returns the VmRecoveryPointLocationAgnosticUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointLocationAgnosticUuidOk

`func (o *RecoveryPointResources) GetVmRecoveryPointLocationAgnosticUuidOk() (*string, bool)`

GetVmRecoveryPointLocationAgnosticUuidOk returns a tuple with the VmRecoveryPointLocationAgnosticUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources) SetVmRecoveryPointLocationAgnosticUuid(v string)`

SetVmRecoveryPointLocationAgnosticUuid sets VmRecoveryPointLocationAgnosticUuid field to given value.

### HasVmRecoveryPointLocationAgnosticUuid

`func (o *RecoveryPointResources) HasVmRecoveryPointLocationAgnosticUuid() bool`

HasVmRecoveryPointLocationAgnosticUuid returns a boolean if a field has been set.

### GetExpirationTime

`func (o *RecoveryPointResources) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *RecoveryPointResources) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *RecoveryPointResources) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *RecoveryPointResources) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


