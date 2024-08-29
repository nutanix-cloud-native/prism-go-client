# ReportInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the report. This will be part of generated report.  | [optional] 
**RecipientList** | Pointer to [**[]Recipient**](Recipient.md) | Recipients in addition to specified on the report config. | [optional] 
**DataStartTime** | Pointer to **time.Time** | UTC date and time in \&quot;%Y-%m-%d %H:%M:%S\&quot; format for data collection start point.  | [optional] 
**GenerationFormat** | Pointer to **[]string** | List specifying the formats in which report is to be created. | [optional] 
**RecipientFormat** | Pointer to **[]string** | List specifying the formats in which report is to be sent. | [optional] 
**SaveInstance** | Pointer to **bool** | Generated instance saved or not. | [optional] 
**DataEndTime** | Pointer to **time.Time** | UTC date and time in \&quot;%Y-%m-%d %H:%M:%S\&quot; format for data collection end point.  | [optional] 
**Timezone** | Pointer to **string** | Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net  | [optional] 
**RuntimeKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes. | [optional] 
**ReportConfigReference** | [**ReportConfigReference**](ReportConfigReference.md) |  | 

## Methods

### NewReportInstance

`func NewReportInstance(reportConfigReference ReportConfigReference, ) *ReportInstance`

NewReportInstance instantiates a new ReportInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstanceWithDefaults

`func NewReportInstanceWithDefaults() *ReportInstance`

NewReportInstanceWithDefaults instantiates a new ReportInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReportInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecipientList

`func (o *ReportInstance) GetRecipientList() []Recipient`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *ReportInstance) GetRecipientListOk() (*[]Recipient, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *ReportInstance) SetRecipientList(v []Recipient)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *ReportInstance) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.

### GetDataStartTime

`func (o *ReportInstance) GetDataStartTime() time.Time`

GetDataStartTime returns the DataStartTime field if non-nil, zero value otherwise.

### GetDataStartTimeOk

`func (o *ReportInstance) GetDataStartTimeOk() (*time.Time, bool)`

GetDataStartTimeOk returns a tuple with the DataStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartTime

`func (o *ReportInstance) SetDataStartTime(v time.Time)`

SetDataStartTime sets DataStartTime field to given value.

### HasDataStartTime

`func (o *ReportInstance) HasDataStartTime() bool`

HasDataStartTime returns a boolean if a field has been set.

### GetGenerationFormat

`func (o *ReportInstance) GetGenerationFormat() []string`

GetGenerationFormat returns the GenerationFormat field if non-nil, zero value otherwise.

### GetGenerationFormatOk

`func (o *ReportInstance) GetGenerationFormatOk() (*[]string, bool)`

GetGenerationFormatOk returns a tuple with the GenerationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationFormat

`func (o *ReportInstance) SetGenerationFormat(v []string)`

SetGenerationFormat sets GenerationFormat field to given value.

### HasGenerationFormat

`func (o *ReportInstance) HasGenerationFormat() bool`

HasGenerationFormat returns a boolean if a field has been set.

### GetRecipientFormat

`func (o *ReportInstance) GetRecipientFormat() []string`

GetRecipientFormat returns the RecipientFormat field if non-nil, zero value otherwise.

### GetRecipientFormatOk

`func (o *ReportInstance) GetRecipientFormatOk() (*[]string, bool)`

GetRecipientFormatOk returns a tuple with the RecipientFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFormat

`func (o *ReportInstance) SetRecipientFormat(v []string)`

SetRecipientFormat sets RecipientFormat field to given value.

### HasRecipientFormat

`func (o *ReportInstance) HasRecipientFormat() bool`

HasRecipientFormat returns a boolean if a field has been set.

### GetSaveInstance

`func (o *ReportInstance) GetSaveInstance() bool`

GetSaveInstance returns the SaveInstance field if non-nil, zero value otherwise.

### GetSaveInstanceOk

`func (o *ReportInstance) GetSaveInstanceOk() (*bool, bool)`

GetSaveInstanceOk returns a tuple with the SaveInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveInstance

`func (o *ReportInstance) SetSaveInstance(v bool)`

SetSaveInstance sets SaveInstance field to given value.

### HasSaveInstance

`func (o *ReportInstance) HasSaveInstance() bool`

HasSaveInstance returns a boolean if a field has been set.

### GetDataEndTime

`func (o *ReportInstance) GetDataEndTime() time.Time`

GetDataEndTime returns the DataEndTime field if non-nil, zero value otherwise.

### GetDataEndTimeOk

`func (o *ReportInstance) GetDataEndTimeOk() (*time.Time, bool)`

GetDataEndTimeOk returns a tuple with the DataEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndTime

`func (o *ReportInstance) SetDataEndTime(v time.Time)`

SetDataEndTime sets DataEndTime field to given value.

### HasDataEndTime

`func (o *ReportInstance) HasDataEndTime() bool`

HasDataEndTime returns a boolean if a field has been set.

### GetTimezone

`func (o *ReportInstance) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReportInstance) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReportInstance) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ReportInstance) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetRuntimeKeyValues

`func (o *ReportInstance) GetRuntimeKeyValues() map[string]string`

GetRuntimeKeyValues returns the RuntimeKeyValues field if non-nil, zero value otherwise.

### GetRuntimeKeyValuesOk

`func (o *ReportInstance) GetRuntimeKeyValuesOk() (*map[string]string, bool)`

GetRuntimeKeyValuesOk returns a tuple with the RuntimeKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeKeyValues

`func (o *ReportInstance) SetRuntimeKeyValues(v map[string]string)`

SetRuntimeKeyValues sets RuntimeKeyValues field to given value.

### HasRuntimeKeyValues

`func (o *ReportInstance) HasRuntimeKeyValues() bool`

HasRuntimeKeyValues returns a boolean if a field has been set.

### GetReportConfigReference

`func (o *ReportInstance) GetReportConfigReference() ReportConfigReference`

GetReportConfigReference returns the ReportConfigReference field if non-nil, zero value otherwise.

### GetReportConfigReferenceOk

`func (o *ReportInstance) GetReportConfigReferenceOk() (*ReportConfigReference, bool)`

GetReportConfigReferenceOk returns a tuple with the ReportConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfigReference

`func (o *ReportInstance) SetReportConfigReference(v ReportConfigReference)`

SetReportConfigReference sets ReportConfigReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


