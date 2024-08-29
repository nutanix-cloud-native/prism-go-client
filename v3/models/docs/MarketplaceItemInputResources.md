# MarketplaceItemInputResources

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

### NewMarketplaceItemInputResources

`func NewMarketplaceItemInputResources(appBlueprintTemplate MarketplaceItemInputResourcesAppBlueprintTemplate, author string, ) *MarketplaceItemInputResources`

NewMarketplaceItemInputResources instantiates a new MarketplaceItemInputResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceItemInputResourcesWithDefaults

`func NewMarketplaceItemInputResourcesWithDefaults() *MarketplaceItemInputResources`

NewMarketplaceItemInputResourcesWithDefaults instantiates a new MarketplaceItemInputResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppBlueprintTemplate

`func (o *MarketplaceItemInputResources) GetAppBlueprintTemplate() MarketplaceItemInputResourcesAppBlueprintTemplate`

GetAppBlueprintTemplate returns the AppBlueprintTemplate field if non-nil, zero value otherwise.

### GetAppBlueprintTemplateOk

`func (o *MarketplaceItemInputResources) GetAppBlueprintTemplateOk() (*MarketplaceItemInputResourcesAppBlueprintTemplate, bool)`

GetAppBlueprintTemplateOk returns a tuple with the AppBlueprintTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintTemplate

`func (o *MarketplaceItemInputResources) SetAppBlueprintTemplate(v MarketplaceItemInputResourcesAppBlueprintTemplate)`

SetAppBlueprintTemplate sets AppBlueprintTemplate field to given value.


### GetAppState

`func (o *MarketplaceItemInputResources) GetAppState() string`

GetAppState returns the AppState field if non-nil, zero value otherwise.

### GetAppStateOk

`func (o *MarketplaceItemInputResources) GetAppStateOk() (*string, bool)`

GetAppStateOk returns a tuple with the AppState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppState

`func (o *MarketplaceItemInputResources) SetAppState(v string)`

SetAppState sets AppState field to given value.

### HasAppState

`func (o *MarketplaceItemInputResources) HasAppState() bool`

HasAppState returns a boolean if a field has been set.

### GetAuthor

`func (o *MarketplaceItemInputResources) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MarketplaceItemInputResources) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MarketplaceItemInputResources) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetProjectReferenceList

`func (o *MarketplaceItemInputResources) GetProjectReferenceList() []ProjectReference`

GetProjectReferenceList returns the ProjectReferenceList field if non-nil, zero value otherwise.

### GetProjectReferenceListOk

`func (o *MarketplaceItemInputResources) GetProjectReferenceListOk() (*[]ProjectReference, bool)`

GetProjectReferenceListOk returns a tuple with the ProjectReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReferenceList

`func (o *MarketplaceItemInputResources) SetProjectReferenceList(v []ProjectReference)`

SetProjectReferenceList sets ProjectReferenceList field to given value.

### HasProjectReferenceList

`func (o *MarketplaceItemInputResources) HasProjectReferenceList() bool`

HasProjectReferenceList returns a boolean if a field has been set.

### GetIconReferenceList

`func (o *MarketplaceItemInputResources) GetIconReferenceList() []MarketplaceIcon`

GetIconReferenceList returns the IconReferenceList field if non-nil, zero value otherwise.

### GetIconReferenceListOk

`func (o *MarketplaceItemInputResources) GetIconReferenceListOk() (*[]MarketplaceIcon, bool)`

GetIconReferenceListOk returns a tuple with the IconReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconReferenceList

`func (o *MarketplaceItemInputResources) SetIconReferenceList(v []MarketplaceIcon)`

SetIconReferenceList sets IconReferenceList field to given value.

### HasIconReferenceList

`func (o *MarketplaceItemInputResources) HasIconReferenceList() bool`

HasIconReferenceList returns a boolean if a field has been set.

### GetAppGroupUuid

`func (o *MarketplaceItemInputResources) GetAppGroupUuid() string`

GetAppGroupUuid returns the AppGroupUuid field if non-nil, zero value otherwise.

### GetAppGroupUuidOk

`func (o *MarketplaceItemInputResources) GetAppGroupUuidOk() (*string, bool)`

GetAppGroupUuidOk returns a tuple with the AppGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGroupUuid

`func (o *MarketplaceItemInputResources) SetAppGroupUuid(v string)`

SetAppGroupUuid sets AppGroupUuid field to given value.

### HasAppGroupUuid

`func (o *MarketplaceItemInputResources) HasAppGroupUuid() bool`

HasAppGroupUuid returns a boolean if a field has been set.

### GetChangeLog

`func (o *MarketplaceItemInputResources) GetChangeLog() string`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *MarketplaceItemInputResources) GetChangeLogOk() (*string, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *MarketplaceItemInputResources) SetChangeLog(v string)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *MarketplaceItemInputResources) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetVersion

`func (o *MarketplaceItemInputResources) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketplaceItemInputResources) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketplaceItemInputResources) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketplaceItemInputResources) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAppAttributeList

`func (o *MarketplaceItemInputResources) GetAppAttributeList() []string`

GetAppAttributeList returns the AppAttributeList field if non-nil, zero value otherwise.

### GetAppAttributeListOk

`func (o *MarketplaceItemInputResources) GetAppAttributeListOk() (*[]string, bool)`

GetAppAttributeListOk returns a tuple with the AppAttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAttributeList

`func (o *MarketplaceItemInputResources) SetAppAttributeList(v []string)`

SetAppAttributeList sets AppAttributeList field to given value.

### HasAppAttributeList

`func (o *MarketplaceItemInputResources) HasAppAttributeList() bool`

HasAppAttributeList returns a boolean if a field has been set.

### GetAppSource

`func (o *MarketplaceItemInputResources) GetAppSource() string`

GetAppSource returns the AppSource field if non-nil, zero value otherwise.

### GetAppSourceOk

`func (o *MarketplaceItemInputResources) GetAppSourceOk() (*string, bool)`

GetAppSourceOk returns a tuple with the AppSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSource

`func (o *MarketplaceItemInputResources) SetAppSource(v string)`

SetAppSource sets AppSource field to given value.

### HasAppSource

`func (o *MarketplaceItemInputResources) HasAppSource() bool`

HasAppSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


