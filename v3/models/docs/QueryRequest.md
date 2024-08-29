# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterCriteria** | Pointer to **string** | FIQL filter criteria that will be used to form the filter field for the backend queries.  | [optional] 
**IntervalStartMs** | Pointer to **int64** | For a time-series query, the start of the interval since epoch in ms.  | [optional] 
**EntityType** | Pointer to **string** | Helps identify the query to be executed in Security Monitoring.  | [optional] 
**GroupMemberOffset** | Pointer to **int64** | The offset into the total member set to return per group. | [optional] 
**GroupMemberCount** | Pointer to **int64** | The maximum number of members to return per group. | [optional] 
**GroupMemberSortOrder** | Pointer to **string** | Sort order for entities and entity groups. | [optional] 
**QueryName** | Pointer to **string** | A custom name to use for tagging the query when debugging. | [optional] 
**GroupMemberAttributes** | Pointer to [**[]GroupRequestedAttribute**](GroupRequestedAttribute.md) |  | [optional] 
**GroupMemberSortAttribute** | Pointer to **string** | The name of the attribute that will be used to sort group members.  | [optional] 
**IntervalEndMs** | Pointer to **int64** | For a time-series query, the end of the interval since epoch in ms.  | [optional] 

## Methods

### NewQueryRequest

`func NewQueryRequest() *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterCriteria

`func (o *QueryRequest) GetFilterCriteria() string`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *QueryRequest) GetFilterCriteriaOk() (*string, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *QueryRequest) SetFilterCriteria(v string)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *QueryRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetIntervalStartMs

`func (o *QueryRequest) GetIntervalStartMs() int64`

GetIntervalStartMs returns the IntervalStartMs field if non-nil, zero value otherwise.

### GetIntervalStartMsOk

`func (o *QueryRequest) GetIntervalStartMsOk() (*int64, bool)`

GetIntervalStartMsOk returns a tuple with the IntervalStartMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalStartMs

`func (o *QueryRequest) SetIntervalStartMs(v int64)`

SetIntervalStartMs sets IntervalStartMs field to given value.

### HasIntervalStartMs

`func (o *QueryRequest) HasIntervalStartMs() bool`

HasIntervalStartMs returns a boolean if a field has been set.

### GetEntityType

`func (o *QueryRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *QueryRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *QueryRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *QueryRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetGroupMemberOffset

`func (o *QueryRequest) GetGroupMemberOffset() int64`

GetGroupMemberOffset returns the GroupMemberOffset field if non-nil, zero value otherwise.

### GetGroupMemberOffsetOk

`func (o *QueryRequest) GetGroupMemberOffsetOk() (*int64, bool)`

GetGroupMemberOffsetOk returns a tuple with the GroupMemberOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberOffset

`func (o *QueryRequest) SetGroupMemberOffset(v int64)`

SetGroupMemberOffset sets GroupMemberOffset field to given value.

### HasGroupMemberOffset

`func (o *QueryRequest) HasGroupMemberOffset() bool`

HasGroupMemberOffset returns a boolean if a field has been set.

### GetGroupMemberCount

`func (o *QueryRequest) GetGroupMemberCount() int64`

GetGroupMemberCount returns the GroupMemberCount field if non-nil, zero value otherwise.

### GetGroupMemberCountOk

`func (o *QueryRequest) GetGroupMemberCountOk() (*int64, bool)`

GetGroupMemberCountOk returns a tuple with the GroupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberCount

`func (o *QueryRequest) SetGroupMemberCount(v int64)`

SetGroupMemberCount sets GroupMemberCount field to given value.

### HasGroupMemberCount

`func (o *QueryRequest) HasGroupMemberCount() bool`

HasGroupMemberCount returns a boolean if a field has been set.

### GetGroupMemberSortOrder

`func (o *QueryRequest) GetGroupMemberSortOrder() string`

GetGroupMemberSortOrder returns the GroupMemberSortOrder field if non-nil, zero value otherwise.

### GetGroupMemberSortOrderOk

`func (o *QueryRequest) GetGroupMemberSortOrderOk() (*string, bool)`

GetGroupMemberSortOrderOk returns a tuple with the GroupMemberSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberSortOrder

`func (o *QueryRequest) SetGroupMemberSortOrder(v string)`

SetGroupMemberSortOrder sets GroupMemberSortOrder field to given value.

### HasGroupMemberSortOrder

`func (o *QueryRequest) HasGroupMemberSortOrder() bool`

HasGroupMemberSortOrder returns a boolean if a field has been set.

### GetQueryName

`func (o *QueryRequest) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *QueryRequest) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *QueryRequest) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *QueryRequest) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.

### GetGroupMemberAttributes

`func (o *QueryRequest) GetGroupMemberAttributes() []GroupRequestedAttribute`

GetGroupMemberAttributes returns the GroupMemberAttributes field if non-nil, zero value otherwise.

### GetGroupMemberAttributesOk

`func (o *QueryRequest) GetGroupMemberAttributesOk() (*[]GroupRequestedAttribute, bool)`

GetGroupMemberAttributesOk returns a tuple with the GroupMemberAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttributes

`func (o *QueryRequest) SetGroupMemberAttributes(v []GroupRequestedAttribute)`

SetGroupMemberAttributes sets GroupMemberAttributes field to given value.

### HasGroupMemberAttributes

`func (o *QueryRequest) HasGroupMemberAttributes() bool`

HasGroupMemberAttributes returns a boolean if a field has been set.

### GetGroupMemberSortAttribute

`func (o *QueryRequest) GetGroupMemberSortAttribute() string`

GetGroupMemberSortAttribute returns the GroupMemberSortAttribute field if non-nil, zero value otherwise.

### GetGroupMemberSortAttributeOk

`func (o *QueryRequest) GetGroupMemberSortAttributeOk() (*string, bool)`

GetGroupMemberSortAttributeOk returns a tuple with the GroupMemberSortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberSortAttribute

`func (o *QueryRequest) SetGroupMemberSortAttribute(v string)`

SetGroupMemberSortAttribute sets GroupMemberSortAttribute field to given value.

### HasGroupMemberSortAttribute

`func (o *QueryRequest) HasGroupMemberSortAttribute() bool`

HasGroupMemberSortAttribute returns a boolean if a field has been set.

### GetIntervalEndMs

`func (o *QueryRequest) GetIntervalEndMs() int64`

GetIntervalEndMs returns the IntervalEndMs field if non-nil, zero value otherwise.

### GetIntervalEndMsOk

`func (o *QueryRequest) GetIntervalEndMsOk() (*int64, bool)`

GetIntervalEndMsOk returns a tuple with the IntervalEndMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEndMs

`func (o *QueryRequest) SetIntervalEndMs(v int64)`

SetIntervalEndMs sets IntervalEndMs field to given value.

### HasIntervalEndMs

`func (o *QueryRequest) HasIntervalEndMs() bool`

HasIntervalEndMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


