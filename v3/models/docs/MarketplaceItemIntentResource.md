# MarketplaceItemIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**MarketplaceItemDefStatus**](MarketplaceItemDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**MarketplaceItem**](MarketplaceItem.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**MarketplaceItemMetadata**](MarketplaceItemMetadata.md) |  | 

## Methods

### NewMarketplaceItemIntentResource

`func NewMarketplaceItemIntentResource(metadata MarketplaceItemMetadata, ) *MarketplaceItemIntentResource`

NewMarketplaceItemIntentResource instantiates a new MarketplaceItemIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemIntentResourceWithDefaults

`func NewMarketplaceItemIntentResourceWithDefaults() *MarketplaceItemIntentResource`

NewMarketplaceItemIntentResourceWithDefaults instantiates a new MarketplaceItemIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MarketplaceItemIntentResource) GetStatus() MarketplaceItemDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceItemIntentResource) GetStatusOk() (*MarketplaceItemDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceItemIntentResource) SetStatus(v MarketplaceItemDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceItemIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *MarketplaceItemIntentResource) GetSpec() MarketplaceItem`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MarketplaceItemIntentResource) GetSpecOk() (*MarketplaceItem, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MarketplaceItemIntentResource) SetSpec(v MarketplaceItem)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MarketplaceItemIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *MarketplaceItemIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MarketplaceItemIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MarketplaceItemIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MarketplaceItemIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketplaceItemIntentResource) GetMetadata() MarketplaceItemMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketplaceItemIntentResource) GetMetadataOk() (*MarketplaceItemMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketplaceItemIntentResource) SetMetadata(v MarketplaceItemMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


