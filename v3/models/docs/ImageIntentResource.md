# ImageIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ImageDefStatus**](ImageDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Image**](Image.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ImageMetadata**](ImageMetadata.md) |  | 

## Methods

### NewImageIntentResource

`func NewImageIntentResource(metadata ImageMetadata, ) *ImageIntentResource`

NewImageIntentResource instantiates a new ImageIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageIntentResourceWithDefaults

`func NewImageIntentResourceWithDefaults() *ImageIntentResource`

NewImageIntentResourceWithDefaults instantiates a new ImageIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ImageIntentResource) GetStatus() ImageDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageIntentResource) GetStatusOk() (*ImageDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageIntentResource) SetStatus(v ImageDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImageIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ImageIntentResource) GetSpec() Image`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImageIntentResource) GetSpecOk() (*Image, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImageIntentResource) SetSpec(v Image)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ImageIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ImageIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImageIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImageIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ImageIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ImageIntentResource) GetMetadata() ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageIntentResource) GetMetadataOk() (*ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageIntentResource) SetMetadata(v ImageMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


