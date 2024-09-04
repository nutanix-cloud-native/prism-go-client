# ReportInstanceIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReportInstanceDefStatus**](ReportInstanceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ReportInstance**](ReportInstance.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ReportInstanceMetadata**](ReportInstanceMetadata.md) |  | 

## Methods

### NewReportInstanceIntentResource

`func NewReportInstanceIntentResource(metadata ReportInstanceMetadata, ) *ReportInstanceIntentResource`

NewReportInstanceIntentResource instantiates a new ReportInstanceIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstanceIntentResourceWithDefaults

`func NewReportInstanceIntentResourceWithDefaults() *ReportInstanceIntentResource`

NewReportInstanceIntentResourceWithDefaults instantiates a new ReportInstanceIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReportInstanceIntentResource) GetStatus() ReportInstanceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportInstanceIntentResource) GetStatusOk() (*ReportInstanceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportInstanceIntentResource) SetStatus(v ReportInstanceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportInstanceIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ReportInstanceIntentResource) GetSpec() ReportInstance`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReportInstanceIntentResource) GetSpecOk() (*ReportInstance, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReportInstanceIntentResource) SetSpec(v ReportInstance)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ReportInstanceIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ReportInstanceIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportInstanceIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportInstanceIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ReportInstanceIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ReportInstanceIntentResource) GetMetadata() ReportInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportInstanceIntentResource) GetMetadataOk() (*ReportInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportInstanceIntentResource) SetMetadata(v ReportInstanceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


