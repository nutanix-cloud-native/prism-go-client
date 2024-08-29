# ReportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportCustomization** | Pointer to [**ReportCustomization**](ReportCustomization.md) |  | [optional] 
**TemplateRows** | [**[]TemplateRow**](TemplateRow.md) | List of template rows. | 
**Sections** | Pointer to [**[]Section**](Section.md) | List of sections. | [optional] 
**Name** | Pointer to **string** | Name of the report template. | [optional] 

## Methods

### NewReportTemplate

`func NewReportTemplate(templateRows []TemplateRow, ) *ReportTemplate`

NewReportTemplate instantiates a new ReportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportTemplateWithDefaults

`func NewReportTemplateWithDefaults() *ReportTemplate`

NewReportTemplateWithDefaults instantiates a new ReportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportCustomization

`func (o *ReportTemplate) GetReportCustomization() ReportCustomization`

GetReportCustomization returns the ReportCustomization field if non-nil, zero value otherwise.

### GetReportCustomizationOk

`func (o *ReportTemplate) GetReportCustomizationOk() (*ReportCustomization, bool)`

GetReportCustomizationOk returns a tuple with the ReportCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomization

`func (o *ReportTemplate) SetReportCustomization(v ReportCustomization)`

SetReportCustomization sets ReportCustomization field to given value.

### HasReportCustomization

`func (o *ReportTemplate) HasReportCustomization() bool`

HasReportCustomization returns a boolean if a field has been set.

### GetTemplateRows

`func (o *ReportTemplate) GetTemplateRows() []TemplateRow`

GetTemplateRows returns the TemplateRows field if non-nil, zero value otherwise.

### GetTemplateRowsOk

`func (o *ReportTemplate) GetTemplateRowsOk() (*[]TemplateRow, bool)`

GetTemplateRowsOk returns a tuple with the TemplateRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRows

`func (o *ReportTemplate) SetTemplateRows(v []TemplateRow)`

SetTemplateRows sets TemplateRows field to given value.


### GetSections

`func (o *ReportTemplate) GetSections() []Section`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ReportTemplate) GetSectionsOk() (*[]Section, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ReportTemplate) SetSections(v []Section)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ReportTemplate) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetName

`func (o *ReportTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


