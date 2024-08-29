# Layer2StretchIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Layer2StretchDefStatus**](Layer2StretchDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Layer2Stretch**](Layer2Stretch.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**Layer2StretchMetadata**](Layer2StretchMetadata.md) |  | 

## Methods

### NewLayer2StretchIntentResource

`func NewLayer2StretchIntentResource(metadata Layer2StretchMetadata, ) *Layer2StretchIntentResource`

NewLayer2StretchIntentResource instantiates a new Layer2StretchIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchIntentResourceWithDefaults

`func NewLayer2StretchIntentResourceWithDefaults() *Layer2StretchIntentResource`

NewLayer2StretchIntentResourceWithDefaults instantiates a new Layer2StretchIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Layer2StretchIntentResource) GetStatus() Layer2StretchDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Layer2StretchIntentResource) GetStatusOk() (*Layer2StretchDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Layer2StretchIntentResource) SetStatus(v Layer2StretchDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Layer2StretchIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *Layer2StretchIntentResource) GetSpec() Layer2Stretch`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Layer2StretchIntentResource) GetSpecOk() (*Layer2Stretch, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Layer2StretchIntentResource) SetSpec(v Layer2Stretch)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Layer2StretchIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *Layer2StretchIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Layer2StretchIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Layer2StretchIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Layer2StretchIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *Layer2StretchIntentResource) GetMetadata() Layer2StretchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Layer2StretchIntentResource) GetMetadataOk() (*Layer2StretchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Layer2StretchIntentResource) SetMetadata(v Layer2StretchMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


