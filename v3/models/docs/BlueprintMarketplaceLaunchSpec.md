# BlueprintMarketplaceLaunchSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description for the marketplace item | [optional] 
**EnvironmentUuid** | Pointer to **string** | Environment uuid. | [optional] 
**AppBlueprintName** | **string** | Name of the app blueprint to be created. | 
**SourceMarketplaceName** | Pointer to **string** | Name of the source marketplace item of the app blueprint  | [optional] 
**SourceMarketplaceVersion** | Pointer to **string** | Indicates version of the source marketplace item of the app blueprint  | [optional] 
**Resources** | [**BlueprintUploadResources**](BlueprintUploadResources.md) |  | 

## Methods

### NewBlueprintMarketplaceLaunchSpec

`func NewBlueprintMarketplaceLaunchSpec(appBlueprintName string, resources BlueprintUploadResources, ) *BlueprintMarketplaceLaunchSpec`

NewBlueprintMarketplaceLaunchSpec instantiates a new BlueprintMarketplaceLaunchSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMarketplaceLaunchSpecWithDefaults

`func NewBlueprintMarketplaceLaunchSpecWithDefaults() *BlueprintMarketplaceLaunchSpec`

NewBlueprintMarketplaceLaunchSpecWithDefaults instantiates a new BlueprintMarketplaceLaunchSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BlueprintMarketplaceLaunchSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintMarketplaceLaunchSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintMarketplaceLaunchSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintMarketplaceLaunchSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentUuid

`func (o *BlueprintMarketplaceLaunchSpec) GetEnvironmentUuid() string`

GetEnvironmentUuid returns the EnvironmentUuid field if non-nil, zero value otherwise.

### GetEnvironmentUuidOk

`func (o *BlueprintMarketplaceLaunchSpec) GetEnvironmentUuidOk() (*string, bool)`

GetEnvironmentUuidOk returns a tuple with the EnvironmentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUuid

`func (o *BlueprintMarketplaceLaunchSpec) SetEnvironmentUuid(v string)`

SetEnvironmentUuid sets EnvironmentUuid field to given value.

### HasEnvironmentUuid

`func (o *BlueprintMarketplaceLaunchSpec) HasEnvironmentUuid() bool`

HasEnvironmentUuid returns a boolean if a field has been set.

### GetAppBlueprintName

`func (o *BlueprintMarketplaceLaunchSpec) GetAppBlueprintName() string`

GetAppBlueprintName returns the AppBlueprintName field if non-nil, zero value otherwise.

### GetAppBlueprintNameOk

`func (o *BlueprintMarketplaceLaunchSpec) GetAppBlueprintNameOk() (*string, bool)`

GetAppBlueprintNameOk returns a tuple with the AppBlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintName

`func (o *BlueprintMarketplaceLaunchSpec) SetAppBlueprintName(v string)`

SetAppBlueprintName sets AppBlueprintName field to given value.


### GetSourceMarketplaceName

`func (o *BlueprintMarketplaceLaunchSpec) GetSourceMarketplaceName() string`

GetSourceMarketplaceName returns the SourceMarketplaceName field if non-nil, zero value otherwise.

### GetSourceMarketplaceNameOk

`func (o *BlueprintMarketplaceLaunchSpec) GetSourceMarketplaceNameOk() (*string, bool)`

GetSourceMarketplaceNameOk returns a tuple with the SourceMarketplaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMarketplaceName

`func (o *BlueprintMarketplaceLaunchSpec) SetSourceMarketplaceName(v string)`

SetSourceMarketplaceName sets SourceMarketplaceName field to given value.

### HasSourceMarketplaceName

`func (o *BlueprintMarketplaceLaunchSpec) HasSourceMarketplaceName() bool`

HasSourceMarketplaceName returns a boolean if a field has been set.

### GetSourceMarketplaceVersion

`func (o *BlueprintMarketplaceLaunchSpec) GetSourceMarketplaceVersion() string`

GetSourceMarketplaceVersion returns the SourceMarketplaceVersion field if non-nil, zero value otherwise.

### GetSourceMarketplaceVersionOk

`func (o *BlueprintMarketplaceLaunchSpec) GetSourceMarketplaceVersionOk() (*string, bool)`

GetSourceMarketplaceVersionOk returns a tuple with the SourceMarketplaceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMarketplaceVersion

`func (o *BlueprintMarketplaceLaunchSpec) SetSourceMarketplaceVersion(v string)`

SetSourceMarketplaceVersion sets SourceMarketplaceVersion field to given value.

### HasSourceMarketplaceVersion

`func (o *BlueprintMarketplaceLaunchSpec) HasSourceMarketplaceVersion() bool`

HasSourceMarketplaceVersion returns a boolean if a field has been set.

### GetResources

`func (o *BlueprintMarketplaceLaunchSpec) GetResources() BlueprintUploadResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintMarketplaceLaunchSpec) GetResourcesOk() (*BlueprintUploadResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintMarketplaceLaunchSpec) SetResources(v BlueprintUploadResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


