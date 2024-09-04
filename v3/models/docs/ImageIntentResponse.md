# ImageIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ImageDefStatus**](ImageDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Image**](Image.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ImageMetadata**](ImageMetadata.md) |  | 

## Methods

### NewImageIntentResponse

`func NewImageIntentResponse(apiVersion string, metadata ImageMetadata, ) *ImageIntentResponse`

NewImageIntentResponse instantiates a new ImageIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageIntentResponseWithDefaults

`func NewImageIntentResponseWithDefaults() *ImageIntentResponse`

NewImageIntentResponseWithDefaults instantiates a new ImageIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ImageIntentResponse) GetStatus() ImageDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageIntentResponse) GetStatusOk() (*ImageDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageIntentResponse) SetStatus(v ImageDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImageIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ImageIntentResponse) GetSpec() Image`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImageIntentResponse) GetSpecOk() (*Image, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImageIntentResponse) SetSpec(v Image)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ImageIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ImageIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImageIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImageIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ImageIntentResponse) GetMetadata() ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageIntentResponse) GetMetadataOk() (*ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageIntentResponse) SetMetadata(v ImageMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


