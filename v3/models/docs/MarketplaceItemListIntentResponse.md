# MarketplaceItemListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]MarketplaceItemIntentResource**](MarketplaceItemIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**MarketplaceItemListMetadataOutput**](MarketplaceItemListMetadataOutput.md) |  | 

## Methods

### NewMarketplaceItemListIntentResponse

`func NewMarketplaceItemListIntentResponse(apiVersion string, metadata MarketplaceItemListMetadataOutput, ) *MarketplaceItemListIntentResponse`

NewMarketplaceItemListIntentResponse instantiates a new MarketplaceItemListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemListIntentResponseWithDefaults

`func NewMarketplaceItemListIntentResponseWithDefaults() *MarketplaceItemListIntentResponse`

NewMarketplaceItemListIntentResponseWithDefaults instantiates a new MarketplaceItemListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *MarketplaceItemListIntentResponse) GetEntities() []MarketplaceItemIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *MarketplaceItemListIntentResponse) GetEntitiesOk() (*[]MarketplaceItemIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *MarketplaceItemListIntentResponse) SetEntities(v []MarketplaceItemIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *MarketplaceItemListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *MarketplaceItemListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MarketplaceItemListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MarketplaceItemListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *MarketplaceItemListIntentResponse) GetMetadata() MarketplaceItemListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketplaceItemListIntentResponse) GetMetadataOk() (*MarketplaceItemListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketplaceItemListIntentResponse) SetMetadata(v MarketplaceItemListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


