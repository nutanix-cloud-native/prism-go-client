# SupportCaseIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SupportCaseDefStatus**](SupportCaseDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**SupportCase**](SupportCase.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SupportCaseMetadata**](SupportCaseMetadata.md) |  | 

## Methods

### NewSupportCaseIntentResource

`func NewSupportCaseIntentResource(metadata SupportCaseMetadata, ) *SupportCaseIntentResource`

NewSupportCaseIntentResource instantiates a new SupportCaseIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseIntentResourceWithDefaults

`func NewSupportCaseIntentResourceWithDefaults() *SupportCaseIntentResource`

NewSupportCaseIntentResourceWithDefaults instantiates a new SupportCaseIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SupportCaseIntentResource) GetStatus() SupportCaseDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportCaseIntentResource) GetStatusOk() (*SupportCaseDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportCaseIntentResource) SetStatus(v SupportCaseDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportCaseIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *SupportCaseIntentResource) GetSpec() SupportCase`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SupportCaseIntentResource) GetSpecOk() (*SupportCase, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SupportCaseIntentResource) SetSpec(v SupportCase)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SupportCaseIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *SupportCaseIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SupportCaseIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SupportCaseIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SupportCaseIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SupportCaseIntentResource) GetMetadata() SupportCaseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportCaseIntentResource) GetMetadataOk() (*SupportCaseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportCaseIntentResource) SetMetadata(v SupportCaseMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


