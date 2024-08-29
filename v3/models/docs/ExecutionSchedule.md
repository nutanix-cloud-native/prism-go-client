# ExecutionSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **time.Time** | Start time to execute a request. | 
**TimeZone** | Pointer to **string** | Name of specified time zone. | [optional] 
**TimeoutSecs** | Pointer to **int64** | Time out in seconds for a request execution. | [optional] 

## Methods

### NewExecutionSchedule

`func NewExecutionSchedule(startTime time.Time, ) *ExecutionSchedule`

NewExecutionSchedule instantiates a new ExecutionSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionScheduleWithDefaults

`func NewExecutionScheduleWithDefaults() *ExecutionSchedule`

NewExecutionScheduleWithDefaults instantiates a new ExecutionSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ExecutionSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExecutionSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExecutionSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetTimeZone

`func (o *ExecutionSchedule) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ExecutionSchedule) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ExecutionSchedule) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ExecutionSchedule) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *ExecutionSchedule) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *ExecutionSchedule) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *ExecutionSchedule) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *ExecutionSchedule) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


