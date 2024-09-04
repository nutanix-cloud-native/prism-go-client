# ReportInstance1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the report. | [optional] 
**RecipientList** | Pointer to [**[]Recipient**](Recipient.md) | Recipients in addition to specified on the report config.  | [optional] 
**ErrorMessage** | Pointer to **string** | Error message when generation failed. | [optional] 
**DataStartTime** | Pointer to **time.Time** | UTC date and time in \&quot;%Y-%m-%d %H:%M:%S\&quot; format for data collection start point.  | [optional] 
**GenerationFormat** | Pointer to **[]string** | List specifying the formats in which report is to be created. | [optional] 
**RecipientFormat** | Pointer to **[]string** | List specifying the formats in which report is to be sent. | [optional] 
**SaveInstance** | Pointer to **bool** | Generated instance saved or not. | [optional] 
**DataEndTime** | Pointer to **time.Time** | UTC date and time in \&quot;%Y-%m-%d %H:%M:%S\&quot; format for data collection end point.  | [optional] 
**Timezone** | Pointer to **string** | Timezone in which report is to be generated. This is the list supported by pytz.all_timezones. For more info, check http://pytz.sourceforge.net  | [optional] 
**RuntimeKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes. | [optional] 
**GenerationStatus** | Pointer to **string** | Generation status of the report. | [optional] 
**ReportConfigReference** | Pointer to [**ReportConfigReference**](ReportConfigReference.md) |  | [optional] 

## Methods

### NewReportInstance1

`func NewReportInstance1() *ReportInstance1`

NewReportInstance1 instantiates a new ReportInstance1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstance1WithDefaults

`func NewReportInstance1WithDefaults() *ReportInstance1`

NewReportInstance1WithDefaults instantiates a new ReportInstance1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReportInstance1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportInstance1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportInstance1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportInstance1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRecipientList

`func (o *ReportInstance1) GetRecipientList() []Recipient`

GetRecipientList returns the RecipientList field if non-nil, zero value otherwise.

### GetRecipientListOk

`func (o *ReportInstance1) GetRecipientListOk() (*[]Recipient, bool)`

GetRecipientListOk returns a tuple with the RecipientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientList

`func (o *ReportInstance1) SetRecipientList(v []Recipient)`

SetRecipientList sets RecipientList field to given value.

### HasRecipientList

`func (o *ReportInstance1) HasRecipientList() bool`

HasRecipientList returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ReportInstance1) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ReportInstance1) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ReportInstance1) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ReportInstance1) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDataStartTime

`func (o *ReportInstance1) GetDataStartTime() time.Time`

GetDataStartTime returns the DataStartTime field if non-nil, zero value otherwise.

### GetDataStartTimeOk

`func (o *ReportInstance1) GetDataStartTimeOk() (*time.Time, bool)`

GetDataStartTimeOk returns a tuple with the DataStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStartTime

`func (o *ReportInstance1) SetDataStartTime(v time.Time)`

SetDataStartTime sets DataStartTime field to given value.

### HasDataStartTime

`func (o *ReportInstance1) HasDataStartTime() bool`

HasDataStartTime returns a boolean if a field has been set.

### GetGenerationFormat

`func (o *ReportInstance1) GetGenerationFormat() []string`

GetGenerationFormat returns the GenerationFormat field if non-nil, zero value otherwise.

### GetGenerationFormatOk

`func (o *ReportInstance1) GetGenerationFormatOk() (*[]string, bool)`

GetGenerationFormatOk returns a tuple with the GenerationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationFormat

`func (o *ReportInstance1) SetGenerationFormat(v []string)`

SetGenerationFormat sets GenerationFormat field to given value.

### HasGenerationFormat

`func (o *ReportInstance1) HasGenerationFormat() bool`

HasGenerationFormat returns a boolean if a field has been set.

### GetRecipientFormat

`func (o *ReportInstance1) GetRecipientFormat() []string`

GetRecipientFormat returns the RecipientFormat field if non-nil, zero value otherwise.

### GetRecipientFormatOk

`func (o *ReportInstance1) GetRecipientFormatOk() (*[]string, bool)`

GetRecipientFormatOk returns a tuple with the RecipientFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFormat

`func (o *ReportInstance1) SetRecipientFormat(v []string)`

SetRecipientFormat sets RecipientFormat field to given value.

### HasRecipientFormat

`func (o *ReportInstance1) HasRecipientFormat() bool`

HasRecipientFormat returns a boolean if a field has been set.

### GetSaveInstance

`func (o *ReportInstance1) GetSaveInstance() bool`

GetSaveInstance returns the SaveInstance field if non-nil, zero value otherwise.

### GetSaveInstanceOk

`func (o *ReportInstance1) GetSaveInstanceOk() (*bool, bool)`

GetSaveInstanceOk returns a tuple with the SaveInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveInstance

`func (o *ReportInstance1) SetSaveInstance(v bool)`

SetSaveInstance sets SaveInstance field to given value.

### HasSaveInstance

`func (o *ReportInstance1) HasSaveInstance() bool`

HasSaveInstance returns a boolean if a field has been set.

### GetDataEndTime

`func (o *ReportInstance1) GetDataEndTime() time.Time`

GetDataEndTime returns the DataEndTime field if non-nil, zero value otherwise.

### GetDataEndTimeOk

`func (o *ReportInstance1) GetDataEndTimeOk() (*time.Time, bool)`

GetDataEndTimeOk returns a tuple with the DataEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEndTime

`func (o *ReportInstance1) SetDataEndTime(v time.Time)`

SetDataEndTime sets DataEndTime field to given value.

### HasDataEndTime

`func (o *ReportInstance1) HasDataEndTime() bool`

HasDataEndTime returns a boolean if a field has been set.

### GetTimezone

`func (o *ReportInstance1) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ReportInstance1) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ReportInstance1) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ReportInstance1) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetRuntimeKeyValues

`func (o *ReportInstance1) GetRuntimeKeyValues() map[string]string`

GetRuntimeKeyValues returns the RuntimeKeyValues field if non-nil, zero value otherwise.

### GetRuntimeKeyValuesOk

`func (o *ReportInstance1) GetRuntimeKeyValuesOk() (*map[string]string, bool)`

GetRuntimeKeyValuesOk returns a tuple with the RuntimeKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeKeyValues

`func (o *ReportInstance1) SetRuntimeKeyValues(v map[string]string)`

SetRuntimeKeyValues sets RuntimeKeyValues field to given value.

### HasRuntimeKeyValues

`func (o *ReportInstance1) HasRuntimeKeyValues() bool`

HasRuntimeKeyValues returns a boolean if a field has been set.

### GetGenerationStatus

`func (o *ReportInstance1) GetGenerationStatus() string`

GetGenerationStatus returns the GenerationStatus field if non-nil, zero value otherwise.

### GetGenerationStatusOk

`func (o *ReportInstance1) GetGenerationStatusOk() (*string, bool)`

GetGenerationStatusOk returns a tuple with the GenerationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationStatus

`func (o *ReportInstance1) SetGenerationStatus(v string)`

SetGenerationStatus sets GenerationStatus field to given value.

### HasGenerationStatus

`func (o *ReportInstance1) HasGenerationStatus() bool`

HasGenerationStatus returns a boolean if a field has been set.

### GetReportConfigReference

`func (o *ReportInstance1) GetReportConfigReference() ReportConfigReference`

GetReportConfigReference returns the ReportConfigReference field if non-nil, zero value otherwise.

### GetReportConfigReferenceOk

`func (o *ReportInstance1) GetReportConfigReferenceOk() (*ReportConfigReference, bool)`

GetReportConfigReferenceOk returns a tuple with the ReportConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportConfigReference

`func (o *ReportInstance1) SetReportConfigReference(v ReportConfigReference)`

SetReportConfigReference sets ReportConfigReference field to given value.

### HasReportConfigReference

`func (o *ReportInstance1) HasReportConfigReference() bool`

HasReportConfigReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


