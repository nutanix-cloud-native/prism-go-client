# CategoryMappingListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]CategoryMappingIntentResource**](CategoryMappingIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CategoryMappingListMetadataOutput**](CategoryMappingListMetadataOutput.md) |  | 

## Methods

### NewCategoryMappingListIntentResponse

`func NewCategoryMappingListIntentResponse(apiVersion string, metadata CategoryMappingListMetadataOutput, ) *CategoryMappingListIntentResponse`

NewCategoryMappingListIntentResponse instantiates a new CategoryMappingListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingListIntentResponseWithDefaults

`func NewCategoryMappingListIntentResponseWithDefaults() *CategoryMappingListIntentResponse`

NewCategoryMappingListIntentResponseWithDefaults instantiates a new CategoryMappingListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *CategoryMappingListIntentResponse) GetEntities() []CategoryMappingIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CategoryMappingListIntentResponse) GetEntitiesOk() (*[]CategoryMappingIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *CategoryMappingListIntentResponse) SetEntities(v []CategoryMappingIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *CategoryMappingListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryMappingListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryMappingListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryMappingListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CategoryMappingListIntentResponse) GetMetadata() CategoryMappingListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryMappingListIntentResponse) GetMetadataOk() (*CategoryMappingListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryMappingListIntentResponse) SetMetadata(v CategoryMappingListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


