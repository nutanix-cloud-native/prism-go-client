# PerspectiveLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchQuery** | Pointer to [**AutoCompletion**](AutoCompletion.md) |  | [optional] 
**SummaryDataList** | Pointer to **[]map[string]interface{}** | List of summary data. | [optional] 
**Title** | Pointer to **string** | Title of the perspective link. | [optional] 
**ChildLinkList** | Pointer to **[]map[string]interface{}** | List of child perspecitve links. This is list of perspective_link objects.  | [optional] 
**IsSelected** | Pointer to **bool** | Flag to indicate if this perspective link is selected in the navigation panel. Indicates if this is the active perspective.  | [optional] 

## Methods

### NewPerspectiveLink

`func NewPerspectiveLink() *PerspectiveLink`

NewPerspectiveLink instantiates a new PerspectiveLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerspectiveLinkWithDefaults

`func NewPerspectiveLinkWithDefaults() *PerspectiveLink`

NewPerspectiveLinkWithDefaults instantiates a new PerspectiveLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchQuery

`func (o *PerspectiveLink) GetSearchQuery() AutoCompletion`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *PerspectiveLink) GetSearchQueryOk() (*AutoCompletion, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *PerspectiveLink) SetSearchQuery(v AutoCompletion)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *PerspectiveLink) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetSummaryDataList

`func (o *PerspectiveLink) GetSummaryDataList() []map[string]interface{}`

GetSummaryDataList returns the SummaryDataList field if non-nil, zero value otherwise.

### GetSummaryDataListOk

`func (o *PerspectiveLink) GetSummaryDataListOk() (*[]map[string]interface{}, bool)`

GetSummaryDataListOk returns a tuple with the SummaryDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryDataList

`func (o *PerspectiveLink) SetSummaryDataList(v []map[string]interface{})`

SetSummaryDataList sets SummaryDataList field to given value.

### HasSummaryDataList

`func (o *PerspectiveLink) HasSummaryDataList() bool`

HasSummaryDataList returns a boolean if a field has been set.

### GetTitle

`func (o *PerspectiveLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PerspectiveLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PerspectiveLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PerspectiveLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetChildLinkList

`func (o *PerspectiveLink) GetChildLinkList() []map[string]interface{}`

GetChildLinkList returns the ChildLinkList field if non-nil, zero value otherwise.

### GetChildLinkListOk

`func (o *PerspectiveLink) GetChildLinkListOk() (*[]map[string]interface{}, bool)`

GetChildLinkListOk returns a tuple with the ChildLinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildLinkList

`func (o *PerspectiveLink) SetChildLinkList(v []map[string]interface{})`

SetChildLinkList sets ChildLinkList field to given value.

### HasChildLinkList

`func (o *PerspectiveLink) HasChildLinkList() bool`

HasChildLinkList returns a boolean if a field has been set.

### GetIsSelected

`func (o *PerspectiveLink) GetIsSelected() bool`

GetIsSelected returns the IsSelected field if non-nil, zero value otherwise.

### GetIsSelectedOk

`func (o *PerspectiveLink) GetIsSelectedOk() (*bool, bool)`

GetIsSelectedOk returns a tuple with the IsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelected

`func (o *PerspectiveLink) SetIsSelected(v bool)`

SetIsSelected sets IsSelected field to given value.

### HasIsSelected

`func (o *PerspectiveLink) HasIsSelected() bool`

HasIsSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


