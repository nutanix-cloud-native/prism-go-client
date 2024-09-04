# MarketplaceItemRenderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputType** | **string** |  | 
**InputSpec** | [**SpecForTheGivenEntityType**](SpecForTheGivenEntityType.md) |  | 
**Name** | **string** | Name of the rendered marketplace item | 
**Description** | Pointer to **string** | A description for the rendered marketplace item | [optional] 

## Methods

### NewMarketplaceItemRenderInput

`func NewMarketplaceItemRenderInput(inputType string, inputSpec SpecForTheGivenEntityType, name string, ) *MarketplaceItemRenderInput`

NewMarketplaceItemRenderInput instantiates a new MarketplaceItemRenderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemRenderInputWithDefaults

`func NewMarketplaceItemRenderInputWithDefaults() *MarketplaceItemRenderInput`

NewMarketplaceItemRenderInputWithDefaults instantiates a new MarketplaceItemRenderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputType

`func (o *MarketplaceItemRenderInput) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *MarketplaceItemRenderInput) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *MarketplaceItemRenderInput) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetInputSpec

`func (o *MarketplaceItemRenderInput) GetInputSpec() SpecForTheGivenEntityType`

GetInputSpec returns the InputSpec field if non-nil, zero value otherwise.

### GetInputSpecOk

`func (o *MarketplaceItemRenderInput) GetInputSpecOk() (*SpecForTheGivenEntityType, bool)`

GetInputSpecOk returns a tuple with the InputSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSpec

`func (o *MarketplaceItemRenderInput) SetInputSpec(v SpecForTheGivenEntityType)`

SetInputSpec sets InputSpec field to given value.


### GetName

`func (o *MarketplaceItemRenderInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceItemRenderInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceItemRenderInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MarketplaceItemRenderInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceItemRenderInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceItemRenderInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceItemRenderInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


