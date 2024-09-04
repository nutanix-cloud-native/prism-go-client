# Layer2StretchIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Layer2Stretch**](Layer2Stretch.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**Layer2StretchMetadata**](Layer2StretchMetadata.md) |  | 

## Methods

### NewLayer2StretchIntentInput

`func NewLayer2StretchIntentInput(spec Layer2Stretch, metadata Layer2StretchMetadata, ) *Layer2StretchIntentInput`

NewLayer2StretchIntentInput instantiates a new Layer2StretchIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchIntentInputWithDefaults

`func NewLayer2StretchIntentInputWithDefaults() *Layer2StretchIntentInput`

NewLayer2StretchIntentInputWithDefaults instantiates a new Layer2StretchIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *Layer2StretchIntentInput) GetSpec() Layer2Stretch`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Layer2StretchIntentInput) GetSpecOk() (*Layer2Stretch, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Layer2StretchIntentInput) SetSpec(v Layer2Stretch)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *Layer2StretchIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Layer2StretchIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Layer2StretchIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Layer2StretchIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *Layer2StretchIntentInput) GetMetadata() Layer2StretchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Layer2StretchIntentInput) GetMetadataOk() (*Layer2StretchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Layer2StretchIntentInput) SetMetadata(v Layer2StretchMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


