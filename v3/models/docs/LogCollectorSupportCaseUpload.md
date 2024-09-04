# LogCollectorSupportCaseUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start time of the schedule. The time should be specified in this manner. date-fullyear(4digit) \&quot;-\&quot; date-month(2digit) \&quot;-\&quot; date-mday(2digit) \&quot;T\&quot; time-hour(2digit) \&quot;:\&quot; time-minute(2digit) \&quot;:\&quot; time-second(2digit) [.] \&quot;Z\&quot;  or (\&quot;+\&quot; / \&quot;-\&quot;) time-hour(2digit) \&quot;:\&quot; time-minute(2digit) Examples - 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of   April 12th, 1985 in UTC. - 1996-12-19T16:39:57-08:00 This represents 39 minutes and 57 seconds after the 16th hour   of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).   Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.  | [optional] 
**AnonymizeOutput** | Pointer to **bool** | Flag for anonymized log collection. | [optional] 
**EndTime** | Pointer to **time.Time** | End time of the schedule. The time should be specified in this manner. date-fullyear(4digit) \&quot;-\&quot; date-month(2digit) \&quot;-\&quot; date-mday(2digit) \&quot;T\&quot; time-hour(2digit) \&quot;:\&quot; time-minute(2digit) \&quot;:\&quot; time-second(2digit) [.] \&quot;Z\&quot;  or (\&quot;+\&quot; / \&quot;-\&quot;) time-hour(2digit) \&quot;:\&quot; time-minute(2digit) Examples - 1985-04-12T23:20:50.52Z represents 20 minutes and 50.52 seconds after the 23rd hour of   April 12th, 1985 in UTC. - 1996-12-19T16:39:57-08:00 This represents 39 minutes and 57 seconds after the 16th hour   of December 19th, 1996 with an offset of -08:00 from UTC (Pacific Standard Time).   Note that this is equivalent to 1996-12-20T00:39:57Z in UTC.  | [optional] 
**NumHours** | Pointer to **int32** | Number of hours for which log has to be collected. Starts from current time - no_of_hours till current time. | [optional] 

## Methods

### NewLogCollectorSupportCaseUpload

`func NewLogCollectorSupportCaseUpload() *LogCollectorSupportCaseUpload`

NewLogCollectorSupportCaseUpload instantiates a new LogCollectorSupportCaseUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogCollectorSupportCaseUploadWithDefaults

`func NewLogCollectorSupportCaseUploadWithDefaults() *LogCollectorSupportCaseUpload`

NewLogCollectorSupportCaseUploadWithDefaults instantiates a new LogCollectorSupportCaseUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *LogCollectorSupportCaseUpload) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LogCollectorSupportCaseUpload) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LogCollectorSupportCaseUpload) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LogCollectorSupportCaseUpload) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAnonymizeOutput

`func (o *LogCollectorSupportCaseUpload) GetAnonymizeOutput() bool`

GetAnonymizeOutput returns the AnonymizeOutput field if non-nil, zero value otherwise.

### GetAnonymizeOutputOk

`func (o *LogCollectorSupportCaseUpload) GetAnonymizeOutputOk() (*bool, bool)`

GetAnonymizeOutputOk returns a tuple with the AnonymizeOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeOutput

`func (o *LogCollectorSupportCaseUpload) SetAnonymizeOutput(v bool)`

SetAnonymizeOutput sets AnonymizeOutput field to given value.

### HasAnonymizeOutput

`func (o *LogCollectorSupportCaseUpload) HasAnonymizeOutput() bool`

HasAnonymizeOutput returns a boolean if a field has been set.

### GetEndTime

`func (o *LogCollectorSupportCaseUpload) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LogCollectorSupportCaseUpload) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LogCollectorSupportCaseUpload) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LogCollectorSupportCaseUpload) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetNumHours

`func (o *LogCollectorSupportCaseUpload) GetNumHours() int32`

GetNumHours returns the NumHours field if non-nil, zero value otherwise.

### GetNumHoursOk

`func (o *LogCollectorSupportCaseUpload) GetNumHoursOk() (*int32, bool)`

GetNumHoursOk returns a tuple with the NumHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHours

`func (o *LogCollectorSupportCaseUpload) SetNumHours(v int32)`

SetNumHours sets NumHours field to given value.

### HasNumHours

`func (o *LogCollectorSupportCaseUpload) HasNumHours() bool`

HasNumHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


