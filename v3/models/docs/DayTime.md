# DayTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | **string** | Day of week the policy has to take effect. | 
**Time** | **string** | Time of the day, the policy will be started. This is in \&quot;&lt;x&gt;h:&lt;y&gt;m\&quot; format. The values must be between 00h:00m and 23h:59m. For example user specified 18h:00m and the current time is 17h:00m then the first snapshot will be captured at 18h:00m. If the current time is 19h:00m then the first snapshot will be captured at 18h:00m next day. If not set, policy will be applicable immediately.  | 

## Methods

### NewDayTime

`func NewDayTime(dayOfWeek string, time string, ) *DayTime`

NewDayTime instantiates a new DayTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayTimeWithDefaults

`func NewDayTimeWithDefaults() *DayTime`

NewDayTimeWithDefaults instantiates a new DayTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *DayTime) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *DayTime) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *DayTime) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.


### GetTime

`func (o *DayTime) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DayTime) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DayTime) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


