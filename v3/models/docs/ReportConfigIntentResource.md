# ReportConfigIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReportConfigDefStatus**](ReportConfigDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ReportConfig**](ReportConfig.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ReportConfigMetadata**](ReportConfigMetadata.md) |  | 

## Methods

### NewReportConfigIntentResource

`func NewReportConfigIntentResource(metadata ReportConfigMetadata, ) *ReportConfigIntentResource`

NewReportConfigIntentResource instantiates a new ReportConfigIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigIntentResourceWithDefaults

`func NewReportConfigIntentResourceWithDefaults() *ReportConfigIntentResource`

NewReportConfigIntentResourceWithDefaults instantiates a new ReportConfigIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReportConfigIntentResource) GetStatus() ReportConfigDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportConfigIntentResource) GetStatusOk() (*ReportConfigDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportConfigIntentResource) SetStatus(v ReportConfigDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportConfigIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ReportConfigIntentResource) GetSpec() ReportConfig`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReportConfigIntentResource) GetSpecOk() (*ReportConfig, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReportConfigIntentResource) SetSpec(v ReportConfig)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ReportConfigIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ReportConfigIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportConfigIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportConfigIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ReportConfigIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ReportConfigIntentResource) GetMetadata() ReportConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportConfigIntentResource) GetMetadataOk() (*ReportConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportConfigIntentResource) SetMetadata(v ReportConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


