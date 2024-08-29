# ImageListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ImageIntentResource**](ImageIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ImageListMetadataOutput**](ImageListMetadataOutput.md) |  | 

## Methods

### NewImageListIntentResponse

`func NewImageListIntentResponse(apiVersion string, metadata ImageListMetadataOutput, ) *ImageListIntentResponse`

NewImageListIntentResponse instantiates a new ImageListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageListIntentResponseWithDefaults

`func NewImageListIntentResponseWithDefaults() *ImageListIntentResponse`

NewImageListIntentResponseWithDefaults instantiates a new ImageListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ImageListIntentResponse) GetEntities() []ImageIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ImageListIntentResponse) GetEntitiesOk() (*[]ImageIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ImageListIntentResponse) SetEntities(v []ImageIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ImageListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ImageListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImageListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImageListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ImageListIntentResponse) GetMetadata() ImageListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageListIntentResponse) GetMetadataOk() (*ImageListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageListIntentResponse) SetMetadata(v ImageListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


