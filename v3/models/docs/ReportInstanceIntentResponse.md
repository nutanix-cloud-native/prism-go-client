# ReportInstanceIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ReportInstanceDefStatus**](ReportInstanceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ReportInstance**](ReportInstance.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ReportInstanceMetadata**](ReportInstanceMetadata.md) |  | 

## Methods

### NewReportInstanceIntentResponse

`func NewReportInstanceIntentResponse(apiVersion string, metadata ReportInstanceMetadata, ) *ReportInstanceIntentResponse`

NewReportInstanceIntentResponse instantiates a new ReportInstanceIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstanceIntentResponseWithDefaults

`func NewReportInstanceIntentResponseWithDefaults() *ReportInstanceIntentResponse`

NewReportInstanceIntentResponseWithDefaults instantiates a new ReportInstanceIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReportInstanceIntentResponse) GetStatus() ReportInstanceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportInstanceIntentResponse) GetStatusOk() (*ReportInstanceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportInstanceIntentResponse) SetStatus(v ReportInstanceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportInstanceIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ReportInstanceIntentResponse) GetSpec() ReportInstance`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ReportInstanceIntentResponse) GetSpecOk() (*ReportInstance, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ReportInstanceIntentResponse) SetSpec(v ReportInstance)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ReportInstanceIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ReportInstanceIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportInstanceIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportInstanceIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ReportInstanceIntentResponse) GetMetadata() ReportInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportInstanceIntentResponse) GetMetadataOk() (*ReportInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportInstanceIntentResponse) SetMetadata(v ReportInstanceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


