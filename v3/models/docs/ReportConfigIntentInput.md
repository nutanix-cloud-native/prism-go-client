# ReportConfigIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ReportConfig**](ReportConfig.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ReportConfigMetadata**](ReportConfigMetadata.md) |  | 

## Methods

### NewReportConfigIntentInput

`func NewReportConfigIntentInput(spec ReportConfig, metadata ReportConfigMetadata, ) *ReportConfigIntentInput`

NewReportConfigIntentInput instantiates a new ReportConfigIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigIntentInputWithDefaults

`func NewReportConfigIntentInputWithDefaults() *ReportConfigIntentInput`

NewReportConfigIntentInputWithDefaults instantiates a new ReportConfigIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ReportConfigIntentInput) GetSpec() ReportConfig`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReportConfigIntentInput) GetSpecOk() (*ReportConfig, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReportConfigIntentInput) SetSpec(v ReportConfig)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ReportConfigIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportConfigIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportConfigIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ReportConfigIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ReportConfigIntentInput) GetMetadata() ReportConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportConfigIntentInput) GetMetadataOk() (*ReportConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportConfigIntentInput) SetMetadata(v ReportConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


