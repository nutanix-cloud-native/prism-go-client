# ReportConfigResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPolicy** | Pointer to [**RetentionPolicy**](RetentionPolicy.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the common report config. | [optional] 
**EndTimeOffsetSecs** | Pointer to **int64** | End time for data collection during report generation. | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 
**TemplateSpecVersion** | Pointer to **string** | Version of the template spec. | [optional] 
**StartTimeOffsetSecs** | Pointer to **int64** | Start time for data collection during report generation. | [optional] 
**GenerationFormat** | Pointer to **[]string** | List specifying the formats in which report is to be created. | [optional] 
**Template** | [**ReportTemplate**](ReportTemplate.md) |  | 
**Timezone** | Pointer to **string** | Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net  | [optional] 
**NotificationPolicy** | Pointer to [**NotificationPolicy**](NotificationPolicy.md) |  | [optional] 

## Methods

### NewReportConfigResources

`func NewReportConfigResources(template ReportTemplate, ) *ReportConfigResources`

NewReportConfigResources instantiates a new ReportConfigResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigResourcesWithDefaults

`func NewReportConfigResourcesWithDefaults() *ReportConfigResources`

NewReportConfigResourcesWithDefaults instantiates a new ReportConfigResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPolicy

`func (o *ReportConfigResources) GetRetentionPolicy() RetentionPolicy`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *ReportConfigResources) GetRetentionPolicyOk() (*RetentionPolicy, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *ReportConfigResources) SetRetentionPolicy(v RetentionPolicy)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *ReportConfigResources) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *ReportConfigResources) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportConfigResources) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportConfigResources) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportConfigResources) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndTimeOffsetSecs

`func (o *ReportConfigResources) GetEndTimeOffsetSecs() int64`

GetEndTimeOffsetSecs returns the EndTimeOffsetSecs field if non-nil, zero value otherwise.

### GetEndTimeOffsetSecsOk

`func (o *ReportConfigResources) GetEndTimeOffsetSecsOk() (*int64, bool)`

GetEndTimeOffsetSecsOk returns a tuple with the EndTimeOffsetSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeOffsetSecs

`func (o *ReportConfigResources) SetEndTimeOffsetSecs(v int64)`

SetEndTimeOffsetSecs sets EndTimeOffsetSecs field to given value.

### HasEndTimeOffsetSecs

`func (o *ReportConfigResources) HasEndTimeOffsetSecs() bool`

HasEndTimeOffsetSecs returns a boolean if a field has been set.

### GetSchedule

`func (o *ReportConfigResources) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReportConfigResources) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReportConfigResources) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ReportConfigResources) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTemplateSpecVersion

`func (o *ReportConfigResources) GetTemplateSpecVersion() string`

GetTemplateSpecVersion returns the TemplateSpecVersion field if non-nil, zero value otherwise.

### GetTemplateSpecVersionOk

`func (o *ReportConfigResources) GetTemplateSpecVersionOk() (*string, bool)`

GetTemplateSpecVersionOk returns a tuple with the TemplateSpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSpecVersion

`func (o *ReportConfigResources) SetTemplateSpecVersion(v string)`

SetTemplateSpecVersion sets TemplateSpecVersion field to given value.

### HasTemplateSpecVersion

`func (o *ReportConfigResources) HasTemplateSpecVersion() bool`

HasTemplateSpecVersion returns a boolean if a field has been set.

### GetStartTimeOffsetSecs

`func (o *ReportConfigResources) GetStartTimeOffsetSecs() int64`

GetStartTimeOffsetSecs returns the StartTimeOffsetSecs field if non-nil, zero value otherwise.

### GetStartTimeOffsetSecsOk

`func (o *ReportConfigResources) GetStartTimeOffsetSecsOk() (*int64, bool)`

GetStartTimeOffsetSecsOk returns a tuple with the StartTimeOffsetSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeOffsetSecs

`func (o *ReportConfigResources) SetStartTimeOffsetSecs(v int64)`

SetStartTimeOffsetSecs sets StartTimeOffsetSecs field to given value.

### HasStartTimeOffsetSecs

`func (o *ReportConfigResources) HasStartTimeOffsetSecs() bool`

HasStartTimeOffsetSecs returns a boolean if a field has been set.

### GetGenerationFormat

`func (o *ReportConfigResources) GetGenerationFormat() []string`

GetGenerationFormat returns the GenerationFormat field if non-nil, zero value otherwise.

### GetGenerationFormatOk

`func (o *ReportConfigResources) GetGenerationFormatOk() (*[]string, bool)`

GetGenerationFormatOk returns a tuple with the GenerationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationFormat

`func (o *ReportConfigResources) SetGenerationFormat(v []string)`

SetGenerationFormat sets GenerationFormat field to given value.

### HasGenerationFormat

`func (o *ReportConfigResources) HasGenerationFormat() bool`

HasGenerationFormat returns a boolean if a field has been set.

### GetTemplate

`func (o *ReportConfigResources) GetTemplate() ReportTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ReportConfigResources) GetTemplateOk() (*ReportTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ReportConfigResources) SetTemplate(v ReportTemplate)`

SetTemplate sets Template field to given value.


### GetTimezone

`func (o *ReportConfigResources) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReportConfigResources) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReportConfigResources) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ReportConfigResources) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetNotificationPolicy

`func (o *ReportConfigResources) GetNotificationPolicy() NotificationPolicy`

GetNotificationPolicy returns the NotificationPolicy field if non-nil, zero value otherwise.

### GetNotificationPolicyOk

`func (o *ReportConfigResources) GetNotificationPolicyOk() (*NotificationPolicy, bool)`

GetNotificationPolicyOk returns a tuple with the NotificationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPolicy

`func (o *ReportConfigResources) SetNotificationPolicy(v NotificationPolicy)`

SetNotificationPolicy sets NotificationPolicy field to given value.

### HasNotificationPolicy

`func (o *ReportConfigResources) HasNotificationPolicy() bool`

HasNotificationPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


