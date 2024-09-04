# NgtPolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**NgtPolicy**](NgtPolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**NgtPolicyMetadata**](NgtPolicyMetadata.md) |  | 

## Methods

### NewNgtPolicyIntentInput

`func NewNgtPolicyIntentInput(spec NgtPolicy, metadata NgtPolicyMetadata, ) *NgtPolicyIntentInput`

NewNgtPolicyIntentInput instantiates a new NgtPolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyIntentInputWithDefaults

`func NewNgtPolicyIntentInputWithDefaults() *NgtPolicyIntentInput`

NewNgtPolicyIntentInputWithDefaults instantiates a new NgtPolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *NgtPolicyIntentInput) GetSpec() NgtPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NgtPolicyIntentInput) GetSpecOk() (*NgtPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NgtPolicyIntentInput) SetSpec(v NgtPolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *NgtPolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NgtPolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NgtPolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NgtPolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NgtPolicyIntentInput) GetMetadata() NgtPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NgtPolicyIntentInput) GetMetadataOk() (*NgtPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NgtPolicyIntentInput) SetMetadata(v NgtPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


