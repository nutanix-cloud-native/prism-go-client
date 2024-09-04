# MarketplaceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Marketplace item name | 
**Resources** | [**MarketplaceItemInputResources**](MarketplaceItemInputResources.md) |  | 
**Description** | Pointer to **string** | Marketplace item description | [optional] 

## Methods

### NewMarketplaceItem

`func NewMarketplaceItem(name string, resources MarketplaceItemInputResources, ) *MarketplaceItem`

NewMarketplaceItem instantiates a new MarketplaceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemWithDefaults

`func NewMarketplaceItemWithDefaults() *MarketplaceItem`

NewMarketplaceItemWithDefaults instantiates a new MarketplaceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketplaceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceItem) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *MarketplaceItem) GetResources() MarketplaceItemInputResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MarketplaceItem) GetResourcesOk() (*MarketplaceItemInputResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MarketplaceItem) SetResources(v MarketplaceItemInputResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *MarketplaceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


