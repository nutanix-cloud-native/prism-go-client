# GroupsGroupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityResults** | Pointer to [**[]GroupsEntity**](GroupsEntity.md) |  | [optional] 
**GroupByColumnValue** | Pointer to **string** |  | [optional] 
**TotalEntityCount** | Pointer to **int64** |  | [optional] 
**GroupSummaries** | Pointer to [**map[string]GroupsFieldData**](GroupsFieldData.md) | Group Summary Info Map. | [optional] 

## Methods

### NewGroupsGroupResult

`func NewGroupsGroupResult() *GroupsGroupResult`

NewGroupsGroupResult instantiates a new GroupsGroupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGroupResultWithDefaults

`func NewGroupsGroupResultWithDefaults() *GroupsGroupResult`

NewGroupsGroupResultWithDefaults instantiates a new GroupsGroupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityResults

`func (o *GroupsGroupResult) GetEntityResults() []GroupsEntity`

GetEntityResults returns the EntityResults field if non-nil, zero value otherwise.

### GetEntityResultsOk

`func (o *GroupsGroupResult) GetEntityResultsOk() (*[]GroupsEntity, bool)`

GetEntityResultsOk returns a tuple with the EntityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityResults

`func (o *GroupsGroupResult) SetEntityResults(v []GroupsEntity)`

SetEntityResults sets EntityResults field to given value.

### HasEntityResults

`func (o *GroupsGroupResult) HasEntityResults() bool`

HasEntityResults returns a boolean if a field has been set.

### GetGroupByColumnValue

`func (o *GroupsGroupResult) GetGroupByColumnValue() string`

GetGroupByColumnValue returns the GroupByColumnValue field if non-nil, zero value otherwise.

### GetGroupByColumnValueOk

`func (o *GroupsGroupResult) GetGroupByColumnValueOk() (*string, bool)`

GetGroupByColumnValueOk returns a tuple with the GroupByColumnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByColumnValue

`func (o *GroupsGroupResult) SetGroupByColumnValue(v string)`

SetGroupByColumnValue sets GroupByColumnValue field to given value.

### HasGroupByColumnValue

`func (o *GroupsGroupResult) HasGroupByColumnValue() bool`

HasGroupByColumnValue returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *GroupsGroupResult) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *GroupsGroupResult) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *GroupsGroupResult) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *GroupsGroupResult) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.

### GetGroupSummaries

`func (o *GroupsGroupResult) GetGroupSummaries() map[string]GroupsFieldData`

GetGroupSummaries returns the GroupSummaries field if non-nil, zero value otherwise.

### GetGroupSummariesOk

`func (o *GroupsGroupResult) GetGroupSummariesOk() (*map[string]GroupsFieldData, bool)`

GetGroupSummariesOk returns a tuple with the GroupSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSummaries

`func (o *GroupsGroupResult) SetGroupSummaries(v map[string]GroupsFieldData)`

SetGroupSummaries sets GroupSummaries field to given value.

### HasGroupSummaries

`func (o *GroupsGroupResult) HasGroupSummaries() bool`

HasGroupSummaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


