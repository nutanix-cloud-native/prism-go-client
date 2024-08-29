# MarketplaceItemOutputResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppBlueprintTemplate** | [**MarketplaceItemInputResourcesAppBlueprintTemplate**](MarketplaceItemInputResourcesAppBlueprintTemplate.md) |  | 
**AppState** | Pointer to **string** | State indicating if marketplace item is approved, pending or rejected  | [optional] [default to "PENDING"]
**Author** | **string** | Person or company which developed the app | 
**ProjectReferenceList** | Pointer to [**[]ProjectReference**](ProjectReference.md) | The projects this marketplace item has been assigned to | [optional] 
**IconReferenceList** | Pointer to [**[]MarketplaceIcon**](MarketplaceIcon.md) |  | [optional] 
**AppGroupUuid** | Pointer to **string** | UUID for the group of MPIs which are versions of the same App  | [optional] 
**ChangeLog** | Pointer to **string** | Change log for this version of the app  | [optional] 
**Version** | Pointer to **string** | Indicates version of the App that this marketplace item represents  | [optional] 
**AppAttributeList** | Pointer to **[]string** | Attributes of this app. | [optional] 
**AppSource** | Pointer to **string** | Indicates whether the app is Global(Marketplace item) or Local(User created)  | [optional] [default to "LOCAL"]

## Methods

### NewMarketplaceItemOutputResources

`func NewMarketplaceItemOutputResources(appBlueprintTemplate MarketplaceItemInputResourcesAppBlueprintTemplate, author string, ) *MarketplaceItemOutputResources`

NewMarketplaceItemOutputResources instantiates a new MarketplaceItemOutputResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemOutputResourcesWithDefaults

`func NewMarketplaceItemOutputResourcesWithDefaults() *MarketplaceItemOutputResources`

NewMarketplaceItemOutputResourcesWithDefaults instantiates a new MarketplaceItemOutputResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppBlueprintTemplate

`func (o *MarketplaceItemOutputResources) GetAppBlueprintTemplate() MarketplaceItemInputResourcesAppBlueprintTemplate`

GetAppBlueprintTemplate returns the AppBlueprintTemplate field if non-nil, zero value otherwise.

### GetAppBlueprintTemplateOk

`func (o *MarketplaceItemOutputResources) GetAppBlueprintTemplateOk() (*MarketplaceItemInputResourcesAppBlueprintTemplate, bool)`

GetAppBlueprintTemplateOk returns a tuple with the AppBlueprintTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintTemplate

`func (o *MarketplaceItemOutputResources) SetAppBlueprintTemplate(v MarketplaceItemInputResourcesAppBlueprintTemplate)`

SetAppBlueprintTemplate sets AppBlueprintTemplate field to given value.


### GetAppState

`func (o *MarketplaceItemOutputResources) GetAppState() string`

GetAppState returns the AppState field if non-nil, zero value otherwise.

### GetAppStateOk

`func (o *MarketplaceItemOutputResources) GetAppStateOk() (*string, bool)`

GetAppStateOk returns a tuple with the AppState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppState

`func (o *MarketplaceItemOutputResources) SetAppState(v string)`

SetAppState sets AppState field to given value.

### HasAppState

`func (o *MarketplaceItemOutputResources) HasAppState() bool`

HasAppState returns a boolean if a field has been set.

### GetAuthor

`func (o *MarketplaceItemOutputResources) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MarketplaceItemOutputResources) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MarketplaceItemOutputResources) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetProjectReferenceList

`func (o *MarketplaceItemOutputResources) GetProjectReferenceList() []ProjectReference`

GetProjectReferenceList returns the ProjectReferenceList field if non-nil, zero value otherwise.

### GetProjectReferenceListOk

`func (o *MarketplaceItemOutputResources) GetProjectReferenceListOk() (*[]ProjectReference, bool)`

GetProjectReferenceListOk returns a tuple with the ProjectReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReferenceList

`func (o *MarketplaceItemOutputResources) SetProjectReferenceList(v []ProjectReference)`

SetProjectReferenceList sets ProjectReferenceList field to given value.

### HasProjectReferenceList

`func (o *MarketplaceItemOutputResources) HasProjectReferenceList() bool`

HasProjectReferenceList returns a boolean if a field has been set.

### GetIconReferenceList

`func (o *MarketplaceItemOutputResources) GetIconReferenceList() []MarketplaceIcon`

GetIconReferenceList returns the IconReferenceList field if non-nil, zero value otherwise.

### GetIconReferenceListOk

`func (o *MarketplaceItemOutputResources) GetIconReferenceListOk() (*[]MarketplaceIcon, bool)`

GetIconReferenceListOk returns a tuple with the IconReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconReferenceList

`func (o *MarketplaceItemOutputResources) SetIconReferenceList(v []MarketplaceIcon)`

SetIconReferenceList sets IconReferenceList field to given value.

### HasIconReferenceList

`func (o *MarketplaceItemOutputResources) HasIconReferenceList() bool`

HasIconReferenceList returns a boolean if a field has been set.

### GetAppGroupUuid

`func (o *MarketplaceItemOutputResources) GetAppGroupUuid() string`

GetAppGroupUuid returns the AppGroupUuid field if non-nil, zero value otherwise.

### GetAppGroupUuidOk

`func (o *MarketplaceItemOutputResources) GetAppGroupUuidOk() (*string, bool)`

GetAppGroupUuidOk returns a tuple with the AppGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGroupUuid

`func (o *MarketplaceItemOutputResources) SetAppGroupUuid(v string)`

SetAppGroupUuid sets AppGroupUuid field to given value.

### HasAppGroupUuid

`func (o *MarketplaceItemOutputResources) HasAppGroupUuid() bool`

HasAppGroupUuid returns a boolean if a field has been set.

### GetChangeLog

`func (o *MarketplaceItemOutputResources) GetChangeLog() string`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *MarketplaceItemOutputResources) GetChangeLogOk() (*string, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *MarketplaceItemOutputResources) SetChangeLog(v string)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *MarketplaceItemOutputResources) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetVersion

`func (o *MarketplaceItemOutputResources) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceItemOutputResources) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceItemOutputResources) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceItemOutputResources) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAppAttributeList

`func (o *MarketplaceItemOutputResources) GetAppAttributeList() []string`

GetAppAttributeList returns the AppAttributeList field if non-nil, zero value otherwise.

### GetAppAttributeListOk

`func (o *MarketplaceItemOutputResources) GetAppAttributeListOk() (*[]string, bool)`

GetAppAttributeListOk returns a tuple with the AppAttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAttributeList

`func (o *MarketplaceItemOutputResources) SetAppAttributeList(v []string)`

SetAppAttributeList sets AppAttributeList field to given value.

### HasAppAttributeList

`func (o *MarketplaceItemOutputResources) HasAppAttributeList() bool`

HasAppAttributeList returns a boolean if a field has been set.

### GetAppSource

`func (o *MarketplaceItemOutputResources) GetAppSource() string`

GetAppSource returns the AppSource field if non-nil, zero value otherwise.

### GetAppSourceOk

`func (o *MarketplaceItemOutputResources) GetAppSourceOk() (*string, bool)`

GetAppSourceOk returns a tuple with the AppSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSource

`func (o *MarketplaceItemOutputResources) SetAppSource(v string)`

SetAppSource sets AppSource field to given value.

### HasAppSource

`func (o *MarketplaceItemOutputResources) HasAppSource() bool`

HasAppSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


