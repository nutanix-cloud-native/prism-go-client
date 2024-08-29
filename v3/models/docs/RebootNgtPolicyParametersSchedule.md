# RebootNgtPolicyParametersSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start time of the schedule. The time should be specified in this manner. date-fullyear(4digit) \&quot;-\&quot; date-month(2digit) \&quot;-\&quot; date-mday(2digit) \&quot;T\&quot; time-hour(2digit) \&quot;:\&quot; time-minute(2digit) \&quot;:\&quot; time-second (2digit) [.] \&quot;Z\&quot;  or (\&quot;+\&quot; / \&quot;-\&quot;) time-hour(2digit) \&quot;:\&quot; time-minute(2digit) Examples - 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds   after the 23rd hour of April 12th, 1985 in UTC. - 1996-12-19T16:39:57-08:00 This represents 39 minutes and 57   seconds after the 16th hour of December 19th, 1996 with an offset   of -08:00 from UTC (Pacific Standard Time).   Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.  | [optional] 
**DayTime** | Pointer to [**DayTime**](DayTime.md) |  | [optional] 

## Methods

### NewRebootNgtPolicyParametersSchedule

`func NewRebootNgtPolicyParametersSchedule() *RebootNgtPolicyParametersSchedule`

NewRebootNgtPolicyParametersSchedule instantiates a new RebootNgtPolicyParametersSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootNgtPolicyParametersScheduleWithDefaults

`func NewRebootNgtPolicyParametersScheduleWithDefaults() *RebootNgtPolicyParametersSchedule`

NewRebootNgtPolicyParametersScheduleWithDefaults instantiates a new RebootNgtPolicyParametersSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *RebootNgtPolicyParametersSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RebootNgtPolicyParametersSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RebootNgtPolicyParametersSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RebootNgtPolicyParametersSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDayTime

`func (o *RebootNgtPolicyParametersSchedule) GetDayTime() DayTime`

GetDayTime returns the DayTime field if non-nil, zero value otherwise.

### GetDayTimeOk

`func (o *RebootNgtPolicyParametersSchedule) GetDayTimeOk() (*DayTime, bool)`

GetDayTimeOk returns a tuple with the DayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTime

`func (o *RebootNgtPolicyParametersSchedule) SetDayTime(v DayTime)`

SetDayTime sets DayTime field to given value.

### HasDayTime

`func (o *RebootNgtPolicyParametersSchedule) HasDayTime() bool`

HasDayTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


