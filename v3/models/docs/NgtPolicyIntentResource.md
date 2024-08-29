# NgtPolicyIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NgtPolicyDefStatus**](NgtPolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**NgtPolicy**](NgtPolicy.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**NgtPolicyMetadata**](NgtPolicyMetadata.md) |  | 

## Methods

### NewNgtPolicyIntentResource

`func NewNgtPolicyIntentResource(metadata NgtPolicyMetadata, ) *NgtPolicyIntentResource`

NewNgtPolicyIntentResource instantiates a new NgtPolicyIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyIntentResourceWithDefaults

`func NewNgtPolicyIntentResourceWithDefaults() *NgtPolicyIntentResource`

NewNgtPolicyIntentResourceWithDefaults instantiates a new NgtPolicyIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NgtPolicyIntentResource) GetStatus() NgtPolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NgtPolicyIntentResource) GetStatusOk() (*NgtPolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NgtPolicyIntentResource) SetStatus(v NgtPolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NgtPolicyIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *NgtPolicyIntentResource) GetSpec() NgtPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NgtPolicyIntentResource) GetSpecOk() (*NgtPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NgtPolicyIntentResource) SetSpec(v NgtPolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NgtPolicyIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *NgtPolicyIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NgtPolicyIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NgtPolicyIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NgtPolicyIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NgtPolicyIntentResource) GetMetadata() NgtPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NgtPolicyIntentResource) GetMetadataOk() (*NgtPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NgtPolicyIntentResource) SetMetadata(v NgtPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


