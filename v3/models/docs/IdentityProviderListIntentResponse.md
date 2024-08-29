# IdentityProviderListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]IdentityProviderIntentResource**](IdentityProviderIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**IdentityProviderListMetadataOutput**](IdentityProviderListMetadataOutput.md) |  | 

## Methods

### NewIdentityProviderListIntentResponse

`func NewIdentityProviderListIntentResponse(apiVersion string, metadata IdentityProviderListMetadataOutput, ) *IdentityProviderListIntentResponse`

NewIdentityProviderListIntentResponse instantiates a new IdentityProviderListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderListIntentResponseWithDefaults

`func NewIdentityProviderListIntentResponseWithDefaults() *IdentityProviderListIntentResponse`

NewIdentityProviderListIntentResponseWithDefaults instantiates a new IdentityProviderListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *IdentityProviderListIntentResponse) GetEntities() []IdentityProviderIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *IdentityProviderListIntentResponse) GetEntitiesOk() (*[]IdentityProviderIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *IdentityProviderListIntentResponse) SetEntities(v []IdentityProviderIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *IdentityProviderListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *IdentityProviderListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IdentityProviderListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IdentityProviderListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *IdentityProviderListIntentResponse) GetMetadata() IdentityProviderListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityProviderListIntentResponse) GetMetadataOk() (*IdentityProviderListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityProviderListIntentResponse) SetMetadata(v IdentityProviderListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


