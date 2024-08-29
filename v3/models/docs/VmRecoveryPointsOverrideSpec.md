# VmRecoveryPointsOverrideSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | **time.Time** | The time when this recovery point expires and will be garbage collected. This is in internet date/time format (RFC 3339). For example, 1985-04-12T23:20:50.52Z, this represents 20 minutes and 50.52 seconds after the 23rd hour of April 12th, 1985 in UTC. If not set, then the recovery point never expires.  | 

## Methods

### NewVmRecoveryPointsOverrideSpec

`func NewVmRecoveryPointsOverrideSpec(expirationTime time.Time, ) *VmRecoveryPointsOverrideSpec`

NewVmRecoveryPointsOverrideSpec instantiates a new VmRecoveryPointsOverrideSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointsOverrideSpecWithDefaults

`func NewVmRecoveryPointsOverrideSpecWithDefaults() *VmRecoveryPointsOverrideSpec`

NewVmRecoveryPointsOverrideSpecWithDefaults instantiates a new VmRecoveryPointsOverrideSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *VmRecoveryPointsOverrideSpec) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *VmRecoveryPointsOverrideSpec) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *VmRecoveryPointsOverrideSpec) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


