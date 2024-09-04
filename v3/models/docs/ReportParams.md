# ReportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** |  | [optional] [default to "en-US"]
**ScenarioUuid** | **string** |  | 

## Methods

### NewReportParams

`func NewReportParams(scenarioUuid string, ) *ReportParams`

NewReportParams instantiates a new ReportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportParamsWithDefaults

`func NewReportParamsWithDefaults() *ReportParams`

NewReportParamsWithDefaults instantiates a new ReportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *ReportParams) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ReportParams) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ReportParams) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ReportParams) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetScenarioUuid

`func (o *ReportParams) GetScenarioUuid() string`

GetScenarioUuid returns the ScenarioUuid field if non-nil, zero value otherwise.

### GetScenarioUuidOk

`func (o *ReportParams) GetScenarioUuidOk() (*string, bool)`

GetScenarioUuidOk returns a tuple with the ScenarioUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenarioUuid

`func (o *ReportParams) SetScenarioUuid(v string)`

SetScenarioUuid sets ScenarioUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


