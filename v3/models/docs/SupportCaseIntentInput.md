# SupportCaseIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**SupportCase**](SupportCase.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SupportCaseMetadata**](SupportCaseMetadata.md) |  | 

## Methods

### NewSupportCaseIntentInput

`func NewSupportCaseIntentInput(spec SupportCase, metadata SupportCaseMetadata, ) *SupportCaseIntentInput`

NewSupportCaseIntentInput instantiates a new SupportCaseIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseIntentInputWithDefaults

`func NewSupportCaseIntentInputWithDefaults() *SupportCaseIntentInput`

NewSupportCaseIntentInputWithDefaults instantiates a new SupportCaseIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *SupportCaseIntentInput) GetSpec() SupportCase`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SupportCaseIntentInput) GetSpecOk() (*SupportCase, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SupportCaseIntentInput) SetSpec(v SupportCase)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *SupportCaseIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SupportCaseIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SupportCaseIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SupportCaseIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SupportCaseIntentInput) GetMetadata() SupportCaseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportCaseIntentInput) GetMetadataOk() (*SupportCaseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportCaseIntentInput) SetMetadata(v SupportCaseMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


