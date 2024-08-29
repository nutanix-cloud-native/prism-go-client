# AddressGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AddressGroupResponseResource**](AddressGroupResponseResource.md) |  | [optional] 
**Metadata** | Pointer to [**AddressGroupListMetadata**](AddressGroupListMetadata.md) |  | [optional] 

## Methods

### NewAddressGroupListResponse

`func NewAddressGroupListResponse() *AddressGroupListResponse`

NewAddressGroupListResponse instantiates a new AddressGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGroupListResponseWithDefaults

`func NewAddressGroupListResponseWithDefaults() *AddressGroupListResponse`

NewAddressGroupListResponseWithDefaults instantiates a new AddressGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AddressGroupListResponse) GetEntities() []AddressGroupResponseResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AddressGroupListResponse) GetEntitiesOk() (*[]AddressGroupResponseResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AddressGroupListResponse) SetEntities(v []AddressGroupResponseResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AddressGroupListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMetadata

`func (o *AddressGroupListResponse) GetMetadata() AddressGroupListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddressGroupListResponse) GetMetadataOk() (*AddressGroupListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddressGroupListResponse) SetMetadata(v AddressGroupListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddressGroupListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


