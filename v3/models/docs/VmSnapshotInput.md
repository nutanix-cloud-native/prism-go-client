# VmSnapshotInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The target time at which the user wishes to have a fully usable recovery point. This is in internet date/time format (RFC 3339). For example, 1985-04-12T22:10:10Z, this represents 10 minutes and 10.10 seconds after the 22nd hour of April 12th, 1985 in UTC. The backend will generate a complete recovery point out of the closest available partial backend recovery point that is created on or before the provided timestamp. This field is explicitly used for the realization semantics for high frequency snapshotting. If this field is set, none of the other fields in this parameter are valid.  | [optional] 
**ExpirationTime** | Pointer to **time.Time** | The time when this recovery point expires and will be garbage collected. This is in internet date/time format (RFC 3339). If not set, then the recovery point never expires.  | [optional] 
**VmRecoveryPointUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the recovery point. | [optional] 
**RecoveryPointType** | Pointer to **string** | Crash consistent or Application Consistent recovery point | [optional] 

## Methods

### NewVmSnapshotInput

`func NewVmSnapshotInput() *VmSnapshotInput`

NewVmSnapshotInput instantiates a new VmSnapshotInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSnapshotInputWithDefaults

`func NewVmSnapshotInputWithDefaults() *VmSnapshotInput`

NewVmSnapshotInputWithDefaults instantiates a new VmSnapshotInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *VmSnapshotInput) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VmSnapshotInput) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VmSnapshotInput) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VmSnapshotInput) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetExpirationTime

`func (o *VmSnapshotInput) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VmSnapshotInput) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VmSnapshotInput) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *VmSnapshotInput) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetVmRecoveryPointUuid

`func (o *VmSnapshotInput) GetVmRecoveryPointUuid() string`

GetVmRecoveryPointUuid returns the VmRecoveryPointUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointUuidOk

`func (o *VmSnapshotInput) GetVmRecoveryPointUuidOk() (*string, bool)`

GetVmRecoveryPointUuidOk returns a tuple with the VmRecoveryPointUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointUuid

`func (o *VmSnapshotInput) SetVmRecoveryPointUuid(v string)`

SetVmRecoveryPointUuid sets VmRecoveryPointUuid field to given value.

### HasVmRecoveryPointUuid

`func (o *VmSnapshotInput) HasVmRecoveryPointUuid() bool`

HasVmRecoveryPointUuid returns a boolean if a field has been set.

### GetName

`func (o *VmSnapshotInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmSnapshotInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmSnapshotInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmSnapshotInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryPointType

`func (o *VmSnapshotInput) GetRecoveryPointType() string`

GetRecoveryPointType returns the RecoveryPointType field if non-nil, zero value otherwise.

### GetRecoveryPointTypeOk

`func (o *VmSnapshotInput) GetRecoveryPointTypeOk() (*string, bool)`

GetRecoveryPointTypeOk returns a tuple with the RecoveryPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointType

`func (o *VmSnapshotInput) SetRecoveryPointType(v string)`

SetRecoveryPointType sets RecoveryPointType field to given value.

### HasRecoveryPointType

`func (o *VmSnapshotInput) HasRecoveryPointType() bool`

HasRecoveryPointType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


