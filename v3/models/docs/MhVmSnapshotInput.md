# MhVmSnapshotInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **time.Time** | The time when this recovery point expires and will be garbage collected. This is in internet date/time format (RFC 3339). If not set, then the recovery point never expires.  | [optional] 
**VmRecoveryPointUuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the recovery point. | [optional] 
**RecoveryPointType** | Pointer to **string** | Crash consistent or Application Consistent recovery point | [optional] 

## Methods

### NewMhVmSnapshotInput

`func NewMhVmSnapshotInput() *MhVmSnapshotInput`

NewMhVmSnapshotInput instantiates a new MhVmSnapshotInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmSnapshotInputWithDefaults

`func NewMhVmSnapshotInputWithDefaults() *MhVmSnapshotInput`

NewMhVmSnapshotInputWithDefaults instantiates a new MhVmSnapshotInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *MhVmSnapshotInput) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *MhVmSnapshotInput) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *MhVmSnapshotInput) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *MhVmSnapshotInput) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetVmRecoveryPointUuid

`func (o *MhVmSnapshotInput) GetVmRecoveryPointUuid() string`

GetVmRecoveryPointUuid returns the VmRecoveryPointUuid field if non-nil, zero value otherwise.

### GetVmRecoveryPointUuidOk

`func (o *MhVmSnapshotInput) GetVmRecoveryPointUuidOk() (*string, bool)`

GetVmRecoveryPointUuidOk returns a tuple with the VmRecoveryPointUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRecoveryPointUuid

`func (o *MhVmSnapshotInput) SetVmRecoveryPointUuid(v string)`

SetVmRecoveryPointUuid sets VmRecoveryPointUuid field to given value.

### HasVmRecoveryPointUuid

`func (o *MhVmSnapshotInput) HasVmRecoveryPointUuid() bool`

HasVmRecoveryPointUuid returns a boolean if a field has been set.

### GetName

`func (o *MhVmSnapshotInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MhVmSnapshotInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MhVmSnapshotInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MhVmSnapshotInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecoveryPointType

`func (o *MhVmSnapshotInput) GetRecoveryPointType() string`

GetRecoveryPointType returns the RecoveryPointType field if non-nil, zero value otherwise.

### GetRecoveryPointTypeOk

`func (o *MhVmSnapshotInput) GetRecoveryPointTypeOk() (*string, bool)`

GetRecoveryPointTypeOk returns a tuple with the RecoveryPointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointType

`func (o *MhVmSnapshotInput) SetRecoveryPointType(v string)`

SetRecoveryPointType sets RecoveryPointType field to given value.

### HasRecoveryPointType

`func (o *MhVmSnapshotInput) HasRecoveryPointType() bool`

HasRecoveryPointType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


