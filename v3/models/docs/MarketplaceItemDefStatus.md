# MarketplaceItemDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Marketplace item name | 
**Resources** | [**MarketplaceItemOutputResources**](MarketplaceItemOutputResources.md) |  | 
**Description** | Pointer to **string** | Marketplace item description | [optional] 

## Methods

### NewMarketplaceItemDefStatus

`func NewMarketplaceItemDefStatus(name string, resources MarketplaceItemOutputResources, ) *MarketplaceItemDefStatus`

NewMarketplaceItemDefStatus instantiates a new MarketplaceItemDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemDefStatusWithDefaults

`func NewMarketplaceItemDefStatusWithDefaults() *MarketplaceItemDefStatus`

NewMarketplaceItemDefStatusWithDefaults instantiates a new MarketplaceItemDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MarketplaceItemDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MarketplaceItemDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MarketplaceItemDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MarketplaceItemDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *MarketplaceItemDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *MarketplaceItemDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *MarketplaceItemDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *MarketplaceItemDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *MarketplaceItemDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceItemDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceItemDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *MarketplaceItemDefStatus) GetResources() MarketplaceItemOutputResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MarketplaceItemDefStatus) GetResourcesOk() (*MarketplaceItemOutputResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MarketplaceItemDefStatus) SetResources(v MarketplaceItemOutputResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *MarketplaceItemDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceItemDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceItemDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceItemDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


