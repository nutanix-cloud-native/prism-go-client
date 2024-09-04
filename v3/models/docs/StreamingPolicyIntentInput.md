# StreamingPolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**StreamingPolicy**](StreamingPolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**StreamingPolicyMetadata**](StreamingPolicyMetadata.md) |  | 

## Methods

### NewStreamingPolicyIntentInput

`func NewStreamingPolicyIntentInput(spec StreamingPolicy, metadata StreamingPolicyMetadata, ) *StreamingPolicyIntentInput`

NewStreamingPolicyIntentInput instantiates a new StreamingPolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyIntentInputWithDefaults

`func NewStreamingPolicyIntentInputWithDefaults() *StreamingPolicyIntentInput`

NewStreamingPolicyIntentInputWithDefaults instantiates a new StreamingPolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *StreamingPolicyIntentInput) GetSpec() StreamingPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StreamingPolicyIntentInput) GetSpecOk() (*StreamingPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StreamingPolicyIntentInput) SetSpec(v StreamingPolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *StreamingPolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamingPolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamingPolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StreamingPolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *StreamingPolicyIntentInput) GetMetadata() StreamingPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamingPolicyIntentInput) GetMetadataOk() (*StreamingPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamingPolicyIntentInput) SetMetadata(v StreamingPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


