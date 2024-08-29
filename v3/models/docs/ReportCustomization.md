# ReportCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FooterHtml** | Pointer to **string** | Custom footer HTML for the report. | [optional] 
**CssStyleSheet** | Pointer to **string** | Global cascadable style for the report. | [optional] 
**HeaderHtml** | Pointer to **string** | Custom header HTML for the report. | [optional] 
**LogoImageUuid** | Pointer to **string** | Custom logo for the report as selected by the user. | [optional] 
**OverridableStyleSheet** | Pointer to **string** | Global overridable style for the report in the form of a serialized JSON. This will be used for page number style in header.  | [optional] 

## Methods

### NewReportCustomization

`func NewReportCustomization() *ReportCustomization`

NewReportCustomization instantiates a new ReportCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCustomizationWithDefaults

`func NewReportCustomizationWithDefaults() *ReportCustomization`

NewReportCustomizationWithDefaults instantiates a new ReportCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFooterHtml

`func (o *ReportCustomization) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *ReportCustomization) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *ReportCustomization) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *ReportCustomization) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### GetCssStyleSheet

`func (o *ReportCustomization) GetCssStyleSheet() string`

GetCssStyleSheet returns the CssStyleSheet field if non-nil, zero value otherwise.

### GetCssStyleSheetOk

`func (o *ReportCustomization) GetCssStyleSheetOk() (*string, bool)`

GetCssStyleSheetOk returns a tuple with the CssStyleSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssStyleSheet

`func (o *ReportCustomization) SetCssStyleSheet(v string)`

SetCssStyleSheet sets CssStyleSheet field to given value.

### HasCssStyleSheet

`func (o *ReportCustomization) HasCssStyleSheet() bool`

HasCssStyleSheet returns a boolean if a field has been set.

### GetHeaderHtml

`func (o *ReportCustomization) GetHeaderHtml() string`

GetHeaderHtml returns the HeaderHtml field if non-nil, zero value otherwise.

### GetHeaderHtmlOk

`func (o *ReportCustomization) GetHeaderHtmlOk() (*string, bool)`

GetHeaderHtmlOk returns a tuple with the HeaderHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHtml

`func (o *ReportCustomization) SetHeaderHtml(v string)`

SetHeaderHtml sets HeaderHtml field to given value.

### HasHeaderHtml

`func (o *ReportCustomization) HasHeaderHtml() bool`

HasHeaderHtml returns a boolean if a field has been set.

### GetLogoImageUuid

`func (o *ReportCustomization) GetLogoImageUuid() string`

GetLogoImageUuid returns the LogoImageUuid field if non-nil, zero value otherwise.

### GetLogoImageUuidOk

`func (o *ReportCustomization) GetLogoImageUuidOk() (*string, bool)`

GetLogoImageUuidOk returns a tuple with the LogoImageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImageUuid

`func (o *ReportCustomization) SetLogoImageUuid(v string)`

SetLogoImageUuid sets LogoImageUuid field to given value.

### HasLogoImageUuid

`func (o *ReportCustomization) HasLogoImageUuid() bool`

HasLogoImageUuid returns a boolean if a field has been set.

### GetOverridableStyleSheet

`func (o *ReportCustomization) GetOverridableStyleSheet() string`

GetOverridableStyleSheet returns the OverridableStyleSheet field if non-nil, zero value otherwise.

### GetOverridableStyleSheetOk

`func (o *ReportCustomization) GetOverridableStyleSheetOk() (*string, bool)`

GetOverridableStyleSheetOk returns a tuple with the OverridableStyleSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridableStyleSheet

`func (o *ReportCustomization) SetOverridableStyleSheet(v string)`

SetOverridableStyleSheet sets OverridableStyleSheet field to given value.

### HasOverridableStyleSheet

`func (o *ReportCustomization) HasOverridableStyleSheet() bool`

HasOverridableStyleSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


