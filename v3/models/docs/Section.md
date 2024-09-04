# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepetitionCriteria** | Pointer to [**RepetitionCriteria**](RepetitionCriteria.md) |  | [optional] 
**TemplateRows** | Pointer to [**[]TemplateRow**](TemplateRow.md) | List of template rows. | [optional] 
**Description** | Pointer to **string** | Description of the section. | [optional] 
**Name** | Pointer to **string** | Name of the section. | [optional] 
**SectionId** | **string** | Identifier for a section. This should be unique in a report config. | 

## Methods

### NewSection

`func NewSection(sectionId string, ) *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepetitionCriteria

`func (o *Section) GetRepetitionCriteria() RepetitionCriteria`

GetRepetitionCriteria returns the RepetitionCriteria field if non-nil, zero value otherwise.

### GetRepetitionCriteriaOk

`func (o *Section) GetRepetitionCriteriaOk() (*RepetitionCriteria, bool)`

GetRepetitionCriteriaOk returns a tuple with the RepetitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionCriteria

`func (o *Section) SetRepetitionCriteria(v RepetitionCriteria)`

SetRepetitionCriteria sets RepetitionCriteria field to given value.

### HasRepetitionCriteria

`func (o *Section) HasRepetitionCriteria() bool`

HasRepetitionCriteria returns a boolean if a field has been set.

### GetTemplateRows

`func (o *Section) GetTemplateRows() []TemplateRow`

GetTemplateRows returns the TemplateRows field if non-nil, zero value otherwise.

### GetTemplateRowsOk

`func (o *Section) GetTemplateRowsOk() (*[]TemplateRow, bool)`

GetTemplateRowsOk returns a tuple with the TemplateRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRows

`func (o *Section) SetTemplateRows(v []TemplateRow)`

SetTemplateRows sets TemplateRows field to given value.

### HasTemplateRows

`func (o *Section) HasTemplateRows() bool`

HasTemplateRows returns a boolean if a field has been set.

### GetDescription

`func (o *Section) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Section) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Section) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Section) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Section) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Section) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Section) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Section) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSectionId

`func (o *Section) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *Section) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *Section) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


