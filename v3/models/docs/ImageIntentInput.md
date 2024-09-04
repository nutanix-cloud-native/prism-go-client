# ImageIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Image**](Image.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ImageMetadata**](ImageMetadata.md) |  | 

## Methods

### NewImageIntentInput

`func NewImageIntentInput(spec Image, metadata ImageMetadata, ) *ImageIntentInput`

NewImageIntentInput instantiates a new ImageIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageIntentInputWithDefaults

`func NewImageIntentInputWithDefaults() *ImageIntentInput`

NewImageIntentInputWithDefaults instantiates a new ImageIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ImageIntentInput) GetSpec() Image`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImageIntentInput) GetSpecOk() (*Image, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImageIntentInput) SetSpec(v Image)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ImageIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImageIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImageIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ImageIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ImageIntentInput) GetMetadata() ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageIntentInput) GetMetadataOk() (*ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageIntentInput) SetMetadata(v ImageMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


