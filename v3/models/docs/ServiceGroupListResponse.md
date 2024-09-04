# ServiceGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ServiceGroupResponseResource**](ServiceGroupResponseResource.md) |  | [optional] 
**Metadata** | Pointer to [**ServiceGroupListMetadata**](ServiceGroupListMetadata.md) |  | [optional] 

## Methods

### NewServiceGroupListResponse

`func NewServiceGroupListResponse() *ServiceGroupListResponse`

NewServiceGroupListResponse instantiates a new ServiceGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupListResponseWithDefaults

`func NewServiceGroupListResponseWithDefaults() *ServiceGroupListResponse`

NewServiceGroupListResponseWithDefaults instantiates a new ServiceGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ServiceGroupListResponse) GetEntities() []ServiceGroupResponseResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ServiceGroupListResponse) GetEntitiesOk() (*[]ServiceGroupResponseResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ServiceGroupListResponse) SetEntities(v []ServiceGroupResponseResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ServiceGroupListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceGroupListResponse) GetMetadata() ServiceGroupListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceGroupListResponse) GetMetadataOk() (*ServiceGroupListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceGroupListResponse) SetMetadata(v ServiceGroupListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceGroupListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


