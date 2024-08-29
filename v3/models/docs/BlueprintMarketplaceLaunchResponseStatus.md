# BlueprintMarketplaceLaunchResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppBlueprintUuid** | Pointer to **string** | AppBlueprint uuid. | [optional] 
**Description** | Pointer to **string** | Description for the marketplace item | [optional] 
**Resources** | Pointer to [**BlueprintResourcesDefStatus**](BlueprintResourcesDefStatus.md) |  | [optional] 

## Methods

### NewBlueprintMarketplaceLaunchResponseStatus

`func NewBlueprintMarketplaceLaunchResponseStatus() *BlueprintMarketplaceLaunchResponseStatus`

NewBlueprintMarketplaceLaunchResponseStatus instantiates a new BlueprintMarketplaceLaunchResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMarketplaceLaunchResponseStatusWithDefaults

`func NewBlueprintMarketplaceLaunchResponseStatusWithDefaults() *BlueprintMarketplaceLaunchResponseStatus`

NewBlueprintMarketplaceLaunchResponseStatusWithDefaults instantiates a new BlueprintMarketplaceLaunchResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppBlueprintUuid

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetAppBlueprintUuid() string`

GetAppBlueprintUuid returns the AppBlueprintUuid field if non-nil, zero value otherwise.

### GetAppBlueprintUuidOk

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetAppBlueprintUuidOk() (*string, bool)`

GetAppBlueprintUuidOk returns a tuple with the AppBlueprintUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintUuid

`func (o *BlueprintMarketplaceLaunchResponseStatus) SetAppBlueprintUuid(v string)`

SetAppBlueprintUuid sets AppBlueprintUuid field to given value.

### HasAppBlueprintUuid

`func (o *BlueprintMarketplaceLaunchResponseStatus) HasAppBlueprintUuid() bool`

HasAppBlueprintUuid returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintMarketplaceLaunchResponseStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintMarketplaceLaunchResponseStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetResources() BlueprintResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintMarketplaceLaunchResponseStatus) GetResourcesOk() (*BlueprintResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintMarketplaceLaunchResponseStatus) SetResources(v BlueprintResourcesDefStatus)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintMarketplaceLaunchResponseStatus) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


