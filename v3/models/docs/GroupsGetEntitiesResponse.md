# GroupsGetEntitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**FilteredGroupCount** | Pointer to **int64** |  | [optional] 
**TotalEntityCount** | Pointer to **int64** |  | [optional] 
**FilteredEntityCount** | Pointer to **int64** |  | [optional] 
**GroupResults** | Pointer to [**[]GroupsGroupResult**](GroupsGroupResult.md) |  | [optional] 
**TotalGroupCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGroupsGetEntitiesResponse

`func NewGroupsGetEntitiesResponse() *GroupsGetEntitiesResponse`

NewGroupsGetEntitiesResponse instantiates a new GroupsGetEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGetEntitiesResponseWithDefaults

`func NewGroupsGetEntitiesResponseWithDefaults() *GroupsGetEntitiesResponse`

NewGroupsGetEntitiesResponseWithDefaults instantiates a new GroupsGetEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *GroupsGetEntitiesResponse) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GroupsGetEntitiesResponse) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GroupsGetEntitiesResponse) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *GroupsGetEntitiesResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFilteredGroupCount

`func (o *GroupsGetEntitiesResponse) GetFilteredGroupCount() int64`

GetFilteredGroupCount returns the FilteredGroupCount field if non-nil, zero value otherwise.

### GetFilteredGroupCountOk

`func (o *GroupsGetEntitiesResponse) GetFilteredGroupCountOk() (*int64, bool)`

GetFilteredGroupCountOk returns a tuple with the FilteredGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredGroupCount

`func (o *GroupsGetEntitiesResponse) SetFilteredGroupCount(v int64)`

SetFilteredGroupCount sets FilteredGroupCount field to given value.

### HasFilteredGroupCount

`func (o *GroupsGetEntitiesResponse) HasFilteredGroupCount() bool`

HasFilteredGroupCount returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *GroupsGetEntitiesResponse) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *GroupsGetEntitiesResponse) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *GroupsGetEntitiesResponse) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *GroupsGetEntitiesResponse) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.

### GetFilteredEntityCount

`func (o *GroupsGetEntitiesResponse) GetFilteredEntityCount() int64`

GetFilteredEntityCount returns the FilteredEntityCount field if non-nil, zero value otherwise.

### GetFilteredEntityCountOk

`func (o *GroupsGetEntitiesResponse) GetFilteredEntityCountOk() (*int64, bool)`

GetFilteredEntityCountOk returns a tuple with the FilteredEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredEntityCount

`func (o *GroupsGetEntitiesResponse) SetFilteredEntityCount(v int64)`

SetFilteredEntityCount sets FilteredEntityCount field to given value.

### HasFilteredEntityCount

`func (o *GroupsGetEntitiesResponse) HasFilteredEntityCount() bool`

HasFilteredEntityCount returns a boolean if a field has been set.

### GetGroupResults

`func (o *GroupsGetEntitiesResponse) GetGroupResults() []GroupsGroupResult`

GetGroupResults returns the GroupResults field if non-nil, zero value otherwise.

### GetGroupResultsOk

`func (o *GroupsGetEntitiesResponse) GetGroupResultsOk() (*[]GroupsGroupResult, bool)`

GetGroupResultsOk returns a tuple with the GroupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupResults

`func (o *GroupsGetEntitiesResponse) SetGroupResults(v []GroupsGroupResult)`

SetGroupResults sets GroupResults field to given value.

### HasGroupResults

`func (o *GroupsGetEntitiesResponse) HasGroupResults() bool`

HasGroupResults returns a boolean if a field has been set.

### GetTotalGroupCount

`func (o *GroupsGetEntitiesResponse) GetTotalGroupCount() int64`

GetTotalGroupCount returns the TotalGroupCount field if non-nil, zero value otherwise.

### GetTotalGroupCountOk

`func (o *GroupsGetEntitiesResponse) GetTotalGroupCountOk() (*int64, bool)`

GetTotalGroupCountOk returns a tuple with the TotalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGroupCount

`func (o *GroupsGetEntitiesResponse) SetTotalGroupCount(v int64)`

SetTotalGroupCount sets TotalGroupCount field to given value.

### HasTotalGroupCount

`func (o *GroupsGetEntitiesResponse) HasTotalGroupCount() bool`

HasTotalGroupCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


