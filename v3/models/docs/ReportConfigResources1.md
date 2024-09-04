# ReportConfigResources1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPolicy** | Pointer to [**RetentionPolicy**](RetentionPolicy.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the common report config. | [optional] 
**EndTimeOffsetSecs** | Pointer to **int64** | Offset for end time for data collection during report generation. | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**TemplateSpecVersion** | Pointer to **string** | Version of the template spec. | [optional] 
**GenerationFormat** | Pointer to **[]string** | List specifying the formats in which report is to be created. | [optional] 
**StartTimeOffsetSecs** | Pointer to **int64** | Offset for start time for data collection during report generation. | [optional] 
**Template** | Pointer to [**ReportTemplate**](ReportTemplate.md) |  | [optional] 
**Timezone** | Pointer to **string** | Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net  | [optional] 
**NotificationPolicy** | Pointer to [**NotificationPolicy**](NotificationPolicy.md) |  | [optional] 
**OutOfBoxReport** | Pointer to **bool** | Flag specifying if Report Config is a pre defined report. | [optional] 

## Methods

### NewReportConfigResources1

`func NewReportConfigResources1() *ReportConfigResources1`

NewReportConfigResources1 instantiates a new ReportConfigResources1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigResources1WithDefaults

`func NewReportConfigResources1WithDefaults() *ReportConfigResources1`

NewReportConfigResources1WithDefaults instantiates a new ReportConfigResources1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPolicy

`func (o *ReportConfigResources1) GetRetentionPolicy() RetentionPolicy`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *ReportConfigResources1) GetRetentionPolicyOk() (*RetentionPolicy, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *ReportConfigResources1) SetRetentionPolicy(v RetentionPolicy)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *ReportConfigResources1) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *ReportConfigResources1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportConfigResources1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportConfigResources1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportConfigResources1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndTimeOffsetSecs

`func (o *ReportConfigResources1) GetEndTimeOffsetSecs() int64`

GetEndTimeOffsetSecs returns the EndTimeOffsetSecs field if non-nil, zero value otherwise.

### GetEndTimeOffsetSecsOk

`func (o *ReportConfigResources1) GetEndTimeOffsetSecsOk() (*int64, bool)`

GetEndTimeOffsetSecsOk returns a tuple with the EndTimeOffsetSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeOffsetSecs

`func (o *ReportConfigResources1) SetEndTimeOffsetSecs(v int64)`

SetEndTimeOffsetSecs sets EndTimeOffsetSecs field to given value.

### HasEndTimeOffsetSecs

`func (o *ReportConfigResources1) HasEndTimeOffsetSecs() bool`

HasEndTimeOffsetSecs returns a boolean if a field has been set.

### GetSchedule

`func (o *ReportConfigResources1) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReportConfigResources1) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReportConfigResources1) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ReportConfigResources1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTemplateSpecVersion

`func (o *ReportConfigResources1) GetTemplateSpecVersion() string`

GetTemplateSpecVersion returns the TemplateSpecVersion field if non-nil, zero value otherwise.

### GetTemplateSpecVersionOk

`func (o *ReportConfigResources1) GetTemplateSpecVersionOk() (*string, bool)`

GetTemplateSpecVersionOk returns a tuple with the TemplateSpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSpecVersion

`func (o *ReportConfigResources1) SetTemplateSpecVersion(v string)`

SetTemplateSpecVersion sets TemplateSpecVersion field to given value.

### HasTemplateSpecVersion

`func (o *ReportConfigResources1) HasTemplateSpecVersion() bool`

HasTemplateSpecVersion returns a boolean if a field has been set.

### GetGenerationFormat

`func (o *ReportConfigResources1) GetGenerationFormat() []string`

GetGenerationFormat returns the GenerationFormat field if non-nil, zero value otherwise.

### GetGenerationFormatOk

`func (o *ReportConfigResources1) GetGenerationFormatOk() (*[]string, bool)`

GetGenerationFormatOk returns a tuple with the GenerationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationFormat

`func (o *ReportConfigResources1) SetGenerationFormat(v []string)`

SetGenerationFormat sets GenerationFormat field to given value.

### HasGenerationFormat

`func (o *ReportConfigResources1) HasGenerationFormat() bool`

HasGenerationFormat returns a boolean if a field has been set.

### GetStartTimeOffsetSecs

`func (o *ReportConfigResources1) GetStartTimeOffsetSecs() int64`

GetStartTimeOffsetSecs returns the StartTimeOffsetSecs field if non-nil, zero value otherwise.

### GetStartTimeOffsetSecsOk

`func (o *ReportConfigResources1) GetStartTimeOffsetSecsOk() (*int64, bool)`

GetStartTimeOffsetSecsOk returns a tuple with the StartTimeOffsetSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeOffsetSecs

`func (o *ReportConfigResources1) SetStartTimeOffsetSecs(v int64)`

SetStartTimeOffsetSecs sets StartTimeOffsetSecs field to given value.

### HasStartTimeOffsetSecs

`func (o *ReportConfigResources1) HasStartTimeOffsetSecs() bool`

HasStartTimeOffsetSecs returns a boolean if a field has been set.

### GetTemplate

`func (o *ReportConfigResources1) GetTemplate() ReportTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ReportConfigResources1) GetTemplateOk() (*ReportTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ReportConfigResources1) SetTemplate(v ReportTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ReportConfigResources1) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimezone

`func (o *ReportConfigResources1) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReportConfigResources1) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReportConfigResources1) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ReportConfigResources1) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetNotificationPolicy

`func (o *ReportConfigResources1) GetNotificationPolicy() NotificationPolicy`

GetNotificationPolicy returns the NotificationPolicy field if non-nil, zero value otherwise.

### GetNotificationPolicyOk

`func (o *ReportConfigResources1) GetNotificationPolicyOk() (*NotificationPolicy, bool)`

GetNotificationPolicyOk returns a tuple with the NotificationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPolicy

`func (o *ReportConfigResources1) SetNotificationPolicy(v NotificationPolicy)`

SetNotificationPolicy sets NotificationPolicy field to given value.

### HasNotificationPolicy

`func (o *ReportConfigResources1) HasNotificationPolicy() bool`

HasNotificationPolicy returns a boolean if a field has been set.

### GetOutOfBoxReport

`func (o *ReportConfigResources1) GetOutOfBoxReport() bool`

GetOutOfBoxReport returns the OutOfBoxReport field if non-nil, zero value otherwise.

### GetOutOfBoxReportOk

`func (o *ReportConfigResources1) GetOutOfBoxReportOk() (*bool, bool)`

GetOutOfBoxReportOk returns a tuple with the OutOfBoxReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBoxReport

`func (o *ReportConfigResources1) SetOutOfBoxReport(v bool)`

SetOutOfBoxReport sets OutOfBoxReport field to given value.

### HasOutOfBoxReport

`func (o *ReportConfigResources1) HasOutOfBoxReport() bool`

HasOutOfBoxReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


