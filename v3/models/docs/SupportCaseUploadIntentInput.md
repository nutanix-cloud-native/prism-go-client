# SupportCaseUploadIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**SupportCaseUpload**](SupportCaseUpload.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SupportCaseUploadMetadata**](SupportCaseUploadMetadata.md) |  | 

## Methods

### NewSupportCaseUploadIntentInput

`func NewSupportCaseUploadIntentInput(spec SupportCaseUpload, metadata SupportCaseUploadMetadata, ) *SupportCaseUploadIntentInput`

NewSupportCaseUploadIntentInput instantiates a new SupportCaseUploadIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseUploadIntentInputWithDefaults

`func NewSupportCaseUploadIntentInputWithDefaults() *SupportCaseUploadIntentInput`

NewSupportCaseUploadIntentInputWithDefaults instantiates a new SupportCaseUploadIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *SupportCaseUploadIntentInput) GetSpec() SupportCaseUpload`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SupportCaseUploadIntentInput) GetSpecOk() (*SupportCaseUpload, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SupportCaseUploadIntentInput) SetSpec(v SupportCaseUpload)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *SupportCaseUploadIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SupportCaseUploadIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SupportCaseUploadIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SupportCaseUploadIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SupportCaseUploadIntentInput) GetMetadata() SupportCaseUploadMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportCaseUploadIntentInput) GetMetadataOk() (*SupportCaseUploadMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportCaseUploadIntentInput) SetMetadata(v SupportCaseUploadMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


