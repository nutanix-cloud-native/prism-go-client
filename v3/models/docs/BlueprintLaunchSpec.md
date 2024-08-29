# BlueprintLaunchSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** | Application name by which the application need to be created. | [optional] 
**AppProfileReference** | Pointer to [**AppProfileReference**](AppProfileReference.md) |  | [optional] 
**Description** | Pointer to **string** | description for blueprint launch | [optional] 
**Resources** | Pointer to [**BlueprintResources**](BlueprintResources.md) |  | [optional] 

## Methods

### NewBlueprintLaunchSpec

`func NewBlueprintLaunchSpec() *BlueprintLaunchSpec`

NewBlueprintLaunchSpec instantiates a new BlueprintLaunchSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintLaunchSpecWithDefaults

`func NewBlueprintLaunchSpecWithDefaults() *BlueprintLaunchSpec`

NewBlueprintLaunchSpecWithDefaults instantiates a new BlueprintLaunchSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *BlueprintLaunchSpec) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *BlueprintLaunchSpec) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *BlueprintLaunchSpec) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *BlueprintLaunchSpec) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetAppProfileReference

`func (o *BlueprintLaunchSpec) GetAppProfileReference() AppProfileReference`

GetAppProfileReference returns the AppProfileReference field if non-nil, zero value otherwise.

### GetAppProfileReferenceOk

`func (o *BlueprintLaunchSpec) GetAppProfileReferenceOk() (*AppProfileReference, bool)`

GetAppProfileReferenceOk returns a tuple with the AppProfileReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileReference

`func (o *BlueprintLaunchSpec) SetAppProfileReference(v AppProfileReference)`

SetAppProfileReference sets AppProfileReference field to given value.

### HasAppProfileReference

`func (o *BlueprintLaunchSpec) HasAppProfileReference() bool`

HasAppProfileReference returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintLaunchSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintLaunchSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintLaunchSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintLaunchSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *BlueprintLaunchSpec) GetResources() BlueprintResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlueprintLaunchSpec) GetResourcesOk() (*BlueprintResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlueprintLaunchSpec) SetResources(v BlueprintResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BlueprintLaunchSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


