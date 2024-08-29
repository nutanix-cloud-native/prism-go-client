# SupportCaseUploadIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SupportCaseUploadDefStatus**](SupportCaseUploadDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**SupportCaseUpload**](SupportCaseUpload.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SupportCaseUploadMetadata**](SupportCaseUploadMetadata.md) |  | 

## Methods

### NewSupportCaseUploadIntentResponse

`func NewSupportCaseUploadIntentResponse(apiVersion string, metadata SupportCaseUploadMetadata, ) *SupportCaseUploadIntentResponse`

NewSupportCaseUploadIntentResponse instantiates a new SupportCaseUploadIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseUploadIntentResponseWithDefaults

`func NewSupportCaseUploadIntentResponseWithDefaults() *SupportCaseUploadIntentResponse`

NewSupportCaseUploadIntentResponseWithDefaults instantiates a new SupportCaseUploadIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SupportCaseUploadIntentResponse) GetStatus() SupportCaseUploadDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportCaseUploadIntentResponse) GetStatusOk() (*SupportCaseUploadDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportCaseUploadIntentResponse) SetStatus(v SupportCaseUploadDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportCaseUploadIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *SupportCaseUploadIntentResponse) GetSpec() SupportCaseUpload`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SupportCaseUploadIntentResponse) GetSpecOk() (*SupportCaseUpload, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SupportCaseUploadIntentResponse) SetSpec(v SupportCaseUpload)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SupportCaseUploadIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *SupportCaseUploadIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SupportCaseUploadIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SupportCaseUploadIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SupportCaseUploadIntentResponse) GetMetadata() SupportCaseUploadMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportCaseUploadIntentResponse) GetMetadataOk() (*SupportCaseUploadMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportCaseUploadIntentResponse) SetMetadata(v SupportCaseUploadMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


