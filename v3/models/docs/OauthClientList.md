# OauthClientList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]OauthClientResponse**](OauthClientResponse.md) |  | 
**Metadata** | [**OauthClientListMetadata**](OauthClientListMetadata.md) |  | 

## Methods

### NewOauthClientList

`func NewOauthClientList(entities []OauthClientResponse, metadata OauthClientListMetadata, ) *OauthClientList`

NewOauthClientList instantiates a new OauthClientList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthClientListWithDefaults

`func NewOauthClientListWithDefaults() *OauthClientList`

NewOauthClientListWithDefaults instantiates a new OauthClientList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *OauthClientList) GetEntities() []OauthClientResponse`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *OauthClientList) GetEntitiesOk() (*[]OauthClientResponse, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *OauthClientList) SetEntities(v []OauthClientResponse)`

SetEntities sets Entities field to given value.


### GetMetadata

`func (o *OauthClientList) GetMetadata() OauthClientListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OauthClientList) GetMetadataOk() (*OauthClientListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OauthClientList) SetMetadata(v OauthClientListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


