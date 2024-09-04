# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalType** | **string** | Type of schedule interval. For schedule once, the interval_type is set to Once. ONCE is for aperiodic schedules.  | 
**DurationSecs** | Pointer to **int64** | Duration of the event. If set, an event of duration duration_usecs will repeat as per the recurrence defined in interval_type.  | [optional] 
**StartTime** | Pointer to **time.Time** | Start time of the schedule. The time should be specified in this manner. date-fullyear(4digit) \&quot;-\&quot; date-month(2digit) \&quot;-\&quot; date-mday(2digit) \&quot;T\&quot; time-hour(2digit) \&quot;:\&quot; time-minute(2digit) \&quot;:\&quot; time-second(2digit) [.] \&quot;Z\&quot;  or (\&quot;+\&quot; / \&quot;-\&quot;) time-hour(2digit) \&quot;:\&quot; time-minute(2digit) Examples - 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of   April 12th, 1985 in UTC. - 1996-12-19T16:39:57-08:00 This represents 39 minutes and 57 seconds after the 16th hour   of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).   Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.  | [optional] 
**IntervalMultiple** | **int32** | Multiple of interval_type. | 
**DayOfWeek** | Pointer to **[]string** | For schedule on weekly basis,like every Monday and Friday 10:00. To use this field, it is required that start_time (datetime) is set, and the interval_type is set to weekly or NONE.  If interval_type is set to NONE, then, only schedule for Monday and Friday once, in this example.  | [optional] 
**EndTime** | Pointer to **time.Time** | End time of the schedule. The time should be specified in this manner. date-fullyear(4digit) \&quot;-\&quot; date-month(2digit) \&quot;-\&quot; date-mday(2digit) \&quot;T\&quot; time-hour(2digit) \&quot;:\&quot; time-minute(2digit) \&quot;:\&quot; time-second(2digit) [.] \&quot;Z\&quot;  or (\&quot;+\&quot; / \&quot;-\&quot;) time-hour(2digit) \&quot;:\&quot; time-minute(2digit) Examples - 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of   April 12th, 1985 in UTC. - 1996-12-19T16:39:57-08:00 This represents 39 minutes and 57 seconds after the 16th hour   of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).   Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.  | [optional] 
**IsSuspended** | Pointer to **bool** | Whether the schedule is suspended. | [optional] 

## Methods

### NewSchedule

`func NewSchedule(intervalType string, intervalMultiple int32, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalType

`func (o *Schedule) GetIntervalType() string`

GetIntervalType returns the IntervalType field if non-nil, zero value otherwise.

### GetIntervalTypeOk

`func (o *Schedule) GetIntervalTypeOk() (*string, bool)`

GetIntervalTypeOk returns a tuple with the IntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalType

`func (o *Schedule) SetIntervalType(v string)`

SetIntervalType sets IntervalType field to given value.


### GetDurationSecs

`func (o *Schedule) GetDurationSecs() int64`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *Schedule) GetDurationSecsOk() (*int64, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *Schedule) SetDurationSecs(v int64)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *Schedule) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetStartTime

`func (o *Schedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Schedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Schedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Schedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetIntervalMultiple

`func (o *Schedule) GetIntervalMultiple() int32`

GetIntervalMultiple returns the IntervalMultiple field if non-nil, zero value otherwise.

### GetIntervalMultipleOk

`func (o *Schedule) GetIntervalMultipleOk() (*int32, bool)`

GetIntervalMultipleOk returns a tuple with the IntervalMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMultiple

`func (o *Schedule) SetIntervalMultiple(v int32)`

SetIntervalMultiple sets IntervalMultiple field to given value.


### GetDayOfWeek

`func (o *Schedule) GetDayOfWeek() []string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *Schedule) GetDayOfWeekOk() (*[]string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *Schedule) SetDayOfWeek(v []string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *Schedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *Schedule) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Schedule) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Schedule) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Schedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsSuspended

`func (o *Schedule) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *Schedule) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *Schedule) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *Schedule) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


