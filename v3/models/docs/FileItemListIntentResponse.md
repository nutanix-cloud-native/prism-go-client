# FileItemListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]FileItemIntentResource**](FileItemIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**FileItemListMetadataOutput**](FileItemListMetadataOutput.md) |  | 

## Methods

### NewFileItemListIntentResponse

`func NewFileItemListIntentResponse(apiVersion string, metadata FileItemListMetadataOutput, ) *FileItemListIntentResponse`

NewFileItemListIntentResponse instantiates a new FileItemListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileItemListIntentResponseWithDefaults

`func NewFileItemListIntentResponseWithDefaults() *FileItemListIntentResponse`

NewFileItemListIntentResponseWithDefaults instantiates a new FileItemListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *FileItemListIntentResponse) GetEntities() []FileItemIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *FileItemListIntentResponse) GetEntitiesOk() (*[]FileItemIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *FileItemListIntentResponse) SetEntities(v []FileItemIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *FileItemListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *FileItemListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FileItemListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FileItemListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *FileItemListIntentResponse) GetMetadata() FileItemListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileItemListIntentResponse) GetMetadataOk() (*FileItemListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileItemListIntentResponse) SetMetadata(v FileItemListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


