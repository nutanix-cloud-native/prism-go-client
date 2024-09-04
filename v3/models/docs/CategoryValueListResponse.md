# CategoryValueListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]CategoryValueStatus**](CategoryValueStatus.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | Pointer to [**CategoryListMetadata**](CategoryListMetadata.md) |  | [optional] 

## Methods

### NewCategoryValueListResponse

`func NewCategoryValueListResponse() *CategoryValueListResponse`

NewCategoryValueListResponse instantiates a new CategoryValueListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryValueListResponseWithDefaults

`func NewCategoryValueListResponseWithDefaults() *CategoryValueListResponse`

NewCategoryValueListResponseWithDefaults instantiates a new CategoryValueListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *CategoryValueListResponse) GetEntities() []CategoryValueStatus`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CategoryValueListResponse) GetEntitiesOk() (*[]CategoryValueStatus, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *CategoryValueListResponse) SetEntities(v []CategoryValueStatus)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *CategoryValueListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryValueListResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryValueListResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryValueListResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryValueListResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CategoryValueListResponse) GetMetadata() CategoryListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryValueListResponse) GetMetadataOk() (*CategoryListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryValueListResponse) SetMetadata(v CategoryListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CategoryValueListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


