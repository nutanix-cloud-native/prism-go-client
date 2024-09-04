# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisambiguatedSearches** | Pointer to [**Perspective**](Perspective.md) |  | [optional] 
**ActivePerspective** | Pointer to [**Perspective**](Perspective.md) |  | [optional] 
**PerspectiveGroupList** | Pointer to [**[]PerspectiveLinkGroup**](PerspectiveLinkGroup.md) | Each result contains a set of perspectives links eg. summary, entities, alerts, metrics, etc.. These are used to populate the left hand navigation panel. These perspective links can be grouped together into different groups. So we send back a list of such perspective groups.  | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisambiguatedSearches

`func (o *Result) GetDisambiguatedSearches() Perspective`

GetDisambiguatedSearches returns the DisambiguatedSearches field if non-nil, zero value otherwise.

### GetDisambiguatedSearchesOk

`func (o *Result) GetDisambiguatedSearchesOk() (*Perspective, bool)`

GetDisambiguatedSearchesOk returns a tuple with the DisambiguatedSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisambiguatedSearches

`func (o *Result) SetDisambiguatedSearches(v Perspective)`

SetDisambiguatedSearches sets DisambiguatedSearches field to given value.

### HasDisambiguatedSearches

`func (o *Result) HasDisambiguatedSearches() bool`

HasDisambiguatedSearches returns a boolean if a field has been set.

### GetActivePerspective

`func (o *Result) GetActivePerspective() Perspective`

GetActivePerspective returns the ActivePerspective field if non-nil, zero value otherwise.

### GetActivePerspectiveOk

`func (o *Result) GetActivePerspectiveOk() (*Perspective, bool)`

GetActivePerspectiveOk returns a tuple with the ActivePerspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePerspective

`func (o *Result) SetActivePerspective(v Perspective)`

SetActivePerspective sets ActivePerspective field to given value.

### HasActivePerspective

`func (o *Result) HasActivePerspective() bool`

HasActivePerspective returns a boolean if a field has been set.

### GetPerspectiveGroupList

`func (o *Result) GetPerspectiveGroupList() []PerspectiveLinkGroup`

GetPerspectiveGroupList returns the PerspectiveGroupList field if non-nil, zero value otherwise.

### GetPerspectiveGroupListOk

`func (o *Result) GetPerspectiveGroupListOk() (*[]PerspectiveLinkGroup, bool)`

GetPerspectiveGroupListOk returns a tuple with the PerspectiveGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspectiveGroupList

`func (o *Result) SetPerspectiveGroupList(v []PerspectiveLinkGroup)`

SetPerspectiveGroupList sets PerspectiveGroupList field to given value.

### HasPerspectiveGroupList

`func (o *Result) HasPerspectiveGroupList() bool`

HasPerspectiveGroupList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


