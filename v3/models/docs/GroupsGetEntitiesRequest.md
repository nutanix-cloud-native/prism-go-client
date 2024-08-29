# GroupsGetEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfIntervalsForLatestData** | Pointer to **int32** | When retrieving latest values, how far back to look as a multiple of the downsampling interval for the metric.  | [optional] 
**GroupSortAttribute** | Pointer to **string** | The name of the attribute that will be used to sort groups.  | [optional] 
**NumberOfBuckets** | Pointer to **int32** | For grouping, how many groups to return. | [optional] 
**EntityIds** | Pointer to **[]string** | A set of entities that the request will be scoped to. | [optional] 
**GroupMemberAttributes** | Pointer to [**[]GroupsRequestedAttribute**](GroupsRequestedAttribute.md) |  | [optional] 
**GroupMemberSortAttribute** | Pointer to **string** | The name of the attribute that will be used to sort group members.  | [optional] 
**BucketBoundary** | Pointer to **int32** | For grouping, the boundary to snap to when grouping. | [optional] 
**GroupOffset** | Pointer to **int64** | The offset into the total set of groups to return. | [optional] 
**DownsamplingInterval** | Pointer to **int32** | Downsampling interval to apply to query if override is desired.  | [optional] 
**IntervalStartMs** | Pointer to **int64** | For a time-series query, the start of the interval since the epoch in ms. Default is latest value only.  | [optional] [default to 0]
**EntityType** | **string** | The entity type that will be requested. | 
**GroupMemberOffset** | Pointer to **int64** | The offset into the total member set to return per group. | [optional] 
**GroupingAttribute** | Pointer to **string** | Attribute that will be used to perform a group-by if needed.  | [optional] 
**GroupMemberSortDownsamplingFunction** | Pointer to **string** | Downsampling function to take time series data and resolve to one value for sorting purposes.  | [optional] 
**GroupSortOrder** | Pointer to **string** | Sort order for entities and entity groups. | [optional] 
**GroupSortDownsampleFunction** | Pointer to **string** | Downsampling function to take time series data and resolve to one value for sorting purposes.  | [optional] 
**FilterCriteria** | Pointer to **string** | FIQL filter criteria that will be used to filter the returned data.  | [optional] 
**LargeBucketBoundary** | Pointer to **int64** | Same as bucket_boundary but supports larger range of values. | [optional] 
**AvailabilityZoneScope** | Pointer to **string** | The scope of availability zones from which to fetch the  data.  | [optional] [default to "LOCAL"]
**GroupCount** | Pointer to **int64** | The maximum number of groups to return in the result. | [optional] 
**GroupAttributes** | Pointer to [**[]GroupsRequestedAttribute**](GroupsRequestedAttribute.md) |  | [optional] 
**IntervalEndMs** | Pointer to **int64** | For a time-series query, the end of the interval since the epoch in ms. Default is latest value only.  | [optional] [default to 0]
**GroupingAttributeType** | Pointer to **string** | The type of an attribute being used for grouping - may be continuous or discrete.  | [optional] 
**GroupMemberCount** | Pointer to **int64** | The maximum number of members to return per group. | [optional] 
**GroupMemberSortOrder** | Pointer to **string** | Sort order for entities and entity groups. | [optional] 
**QueryName** | Pointer to **string** | A custom name to use for tagging the query when debugging. | [optional] 

## Methods

### NewGroupsGetEntitiesRequest

`func NewGroupsGetEntitiesRequest(entityType string, ) *GroupsGetEntitiesRequest`

NewGroupsGetEntitiesRequest instantiates a new GroupsGetEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGetEntitiesRequestWithDefaults

`func NewGroupsGetEntitiesRequestWithDefaults() *GroupsGetEntitiesRequest`

NewGroupsGetEntitiesRequestWithDefaults instantiates a new GroupsGetEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfIntervalsForLatestData

`func (o *GroupsGetEntitiesRequest) GetNumberOfIntervalsForLatestData() int32`

GetNumberOfIntervalsForLatestData returns the NumberOfIntervalsForLatestData field if non-nil, zero value otherwise.

### GetNumberOfIntervalsForLatestDataOk

`func (o *GroupsGetEntitiesRequest) GetNumberOfIntervalsForLatestDataOk() (*int32, bool)`

GetNumberOfIntervalsForLatestDataOk returns a tuple with the NumberOfIntervalsForLatestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfIntervalsForLatestData

`func (o *GroupsGetEntitiesRequest) SetNumberOfIntervalsForLatestData(v int32)`

SetNumberOfIntervalsForLatestData sets NumberOfIntervalsForLatestData field to given value.

### HasNumberOfIntervalsForLatestData

`func (o *GroupsGetEntitiesRequest) HasNumberOfIntervalsForLatestData() bool`

HasNumberOfIntervalsForLatestData returns a boolean if a field has been set.

### GetGroupSortAttribute

`func (o *GroupsGetEntitiesRequest) GetGroupSortAttribute() string`

GetGroupSortAttribute returns the GroupSortAttribute field if non-nil, zero value otherwise.

### GetGroupSortAttributeOk

`func (o *GroupsGetEntitiesRequest) GetGroupSortAttributeOk() (*string, bool)`

GetGroupSortAttributeOk returns a tuple with the GroupSortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSortAttribute

`func (o *GroupsGetEntitiesRequest) SetGroupSortAttribute(v string)`

SetGroupSortAttribute sets GroupSortAttribute field to given value.

### HasGroupSortAttribute

`func (o *GroupsGetEntitiesRequest) HasGroupSortAttribute() bool`

HasGroupSortAttribute returns a boolean if a field has been set.

### GetNumberOfBuckets

`func (o *GroupsGetEntitiesRequest) GetNumberOfBuckets() int32`

GetNumberOfBuckets returns the NumberOfBuckets field if non-nil, zero value otherwise.

### GetNumberOfBucketsOk

`func (o *GroupsGetEntitiesRequest) GetNumberOfBucketsOk() (*int32, bool)`

GetNumberOfBucketsOk returns a tuple with the NumberOfBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBuckets

`func (o *GroupsGetEntitiesRequest) SetNumberOfBuckets(v int32)`

SetNumberOfBuckets sets NumberOfBuckets field to given value.

### HasNumberOfBuckets

`func (o *GroupsGetEntitiesRequest) HasNumberOfBuckets() bool`

HasNumberOfBuckets returns a boolean if a field has been set.

### GetEntityIds

`func (o *GroupsGetEntitiesRequest) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *GroupsGetEntitiesRequest) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *GroupsGetEntitiesRequest) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *GroupsGetEntitiesRequest) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetGroupMemberAttributes

`func (o *GroupsGetEntitiesRequest) GetGroupMemberAttributes() []GroupsRequestedAttribute`

GetGroupMemberAttributes returns the GroupMemberAttributes field if non-nil, zero value otherwise.

### GetGroupMemberAttributesOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberAttributesOk() (*[]GroupsRequestedAttribute, bool)`

GetGroupMemberAttributesOk returns a tuple with the GroupMemberAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttributes

`func (o *GroupsGetEntitiesRequest) SetGroupMemberAttributes(v []GroupsRequestedAttribute)`

SetGroupMemberAttributes sets GroupMemberAttributes field to given value.

### HasGroupMemberAttributes

`func (o *GroupsGetEntitiesRequest) HasGroupMemberAttributes() bool`

HasGroupMemberAttributes returns a boolean if a field has been set.

### GetGroupMemberSortAttribute

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortAttribute() string`

GetGroupMemberSortAttribute returns the GroupMemberSortAttribute field if non-nil, zero value otherwise.

### GetGroupMemberSortAttributeOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortAttributeOk() (*string, bool)`

GetGroupMemberSortAttributeOk returns a tuple with the GroupMemberSortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberSortAttribute

`func (o *GroupsGetEntitiesRequest) SetGroupMemberSortAttribute(v string)`

SetGroupMemberSortAttribute sets GroupMemberSortAttribute field to given value.

### HasGroupMemberSortAttribute

`func (o *GroupsGetEntitiesRequest) HasGroupMemberSortAttribute() bool`

HasGroupMemberSortAttribute returns a boolean if a field has been set.

### GetBucketBoundary

`func (o *GroupsGetEntitiesRequest) GetBucketBoundary() int32`

GetBucketBoundary returns the BucketBoundary field if non-nil, zero value otherwise.

### GetBucketBoundaryOk

`func (o *GroupsGetEntitiesRequest) GetBucketBoundaryOk() (*int32, bool)`

GetBucketBoundaryOk returns a tuple with the BucketBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBoundary

`func (o *GroupsGetEntitiesRequest) SetBucketBoundary(v int32)`

SetBucketBoundary sets BucketBoundary field to given value.

### HasBucketBoundary

`func (o *GroupsGetEntitiesRequest) HasBucketBoundary() bool`

HasBucketBoundary returns a boolean if a field has been set.

### GetGroupOffset

`func (o *GroupsGetEntitiesRequest) GetGroupOffset() int64`

GetGroupOffset returns the GroupOffset field if non-nil, zero value otherwise.

### GetGroupOffsetOk

`func (o *GroupsGetEntitiesRequest) GetGroupOffsetOk() (*int64, bool)`

GetGroupOffsetOk returns a tuple with the GroupOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupOffset

`func (o *GroupsGetEntitiesRequest) SetGroupOffset(v int64)`

SetGroupOffset sets GroupOffset field to given value.

### HasGroupOffset

`func (o *GroupsGetEntitiesRequest) HasGroupOffset() bool`

HasGroupOffset returns a boolean if a field has been set.

### GetDownsamplingInterval

`func (o *GroupsGetEntitiesRequest) GetDownsamplingInterval() int32`

GetDownsamplingInterval returns the DownsamplingInterval field if non-nil, zero value otherwise.

### GetDownsamplingIntervalOk

`func (o *GroupsGetEntitiesRequest) GetDownsamplingIntervalOk() (*int32, bool)`

GetDownsamplingIntervalOk returns a tuple with the DownsamplingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownsamplingInterval

`func (o *GroupsGetEntitiesRequest) SetDownsamplingInterval(v int32)`

SetDownsamplingInterval sets DownsamplingInterval field to given value.

### HasDownsamplingInterval

`func (o *GroupsGetEntitiesRequest) HasDownsamplingInterval() bool`

HasDownsamplingInterval returns a boolean if a field has been set.

### GetIntervalStartMs

`func (o *GroupsGetEntitiesRequest) GetIntervalStartMs() int64`

GetIntervalStartMs returns the IntervalStartMs field if non-nil, zero value otherwise.

### GetIntervalStartMsOk

`func (o *GroupsGetEntitiesRequest) GetIntervalStartMsOk() (*int64, bool)`

GetIntervalStartMsOk returns a tuple with the IntervalStartMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalStartMs

`func (o *GroupsGetEntitiesRequest) SetIntervalStartMs(v int64)`

SetIntervalStartMs sets IntervalStartMs field to given value.

### HasIntervalStartMs

`func (o *GroupsGetEntitiesRequest) HasIntervalStartMs() bool`

HasIntervalStartMs returns a boolean if a field has been set.

### GetEntityType

`func (o *GroupsGetEntitiesRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *GroupsGetEntitiesRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *GroupsGetEntitiesRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetGroupMemberOffset

`func (o *GroupsGetEntitiesRequest) GetGroupMemberOffset() int64`

GetGroupMemberOffset returns the GroupMemberOffset field if non-nil, zero value otherwise.

### GetGroupMemberOffsetOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberOffsetOk() (*int64, bool)`

GetGroupMemberOffsetOk returns a tuple with the GroupMemberOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberOffset

`func (o *GroupsGetEntitiesRequest) SetGroupMemberOffset(v int64)`

SetGroupMemberOffset sets GroupMemberOffset field to given value.

### HasGroupMemberOffset

`func (o *GroupsGetEntitiesRequest) HasGroupMemberOffset() bool`

HasGroupMemberOffset returns a boolean if a field has been set.

### GetGroupingAttribute

`func (o *GroupsGetEntitiesRequest) GetGroupingAttribute() string`

GetGroupingAttribute returns the GroupingAttribute field if non-nil, zero value otherwise.

### GetGroupingAttributeOk

`func (o *GroupsGetEntitiesRequest) GetGroupingAttributeOk() (*string, bool)`

GetGroupingAttributeOk returns a tuple with the GroupingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingAttribute

`func (o *GroupsGetEntitiesRequest) SetGroupingAttribute(v string)`

SetGroupingAttribute sets GroupingAttribute field to given value.

### HasGroupingAttribute

`func (o *GroupsGetEntitiesRequest) HasGroupingAttribute() bool`

HasGroupingAttribute returns a boolean if a field has been set.

### GetGroupMemberSortDownsamplingFunction

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortDownsamplingFunction() string`

GetGroupMemberSortDownsamplingFunction returns the GroupMemberSortDownsamplingFunction field if non-nil, zero value otherwise.

### GetGroupMemberSortDownsamplingFunctionOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortDownsamplingFunctionOk() (*string, bool)`

GetGroupMemberSortDownsamplingFunctionOk returns a tuple with the GroupMemberSortDownsamplingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberSortDownsamplingFunction

`func (o *GroupsGetEntitiesRequest) SetGroupMemberSortDownsamplingFunction(v string)`

SetGroupMemberSortDownsamplingFunction sets GroupMemberSortDownsamplingFunction field to given value.

### HasGroupMemberSortDownsamplingFunction

`func (o *GroupsGetEntitiesRequest) HasGroupMemberSortDownsamplingFunction() bool`

HasGroupMemberSortDownsamplingFunction returns a boolean if a field has been set.

### GetGroupSortOrder

`func (o *GroupsGetEntitiesRequest) GetGroupSortOrder() string`

GetGroupSortOrder returns the GroupSortOrder field if non-nil, zero value otherwise.

### GetGroupSortOrderOk

`func (o *GroupsGetEntitiesRequest) GetGroupSortOrderOk() (*string, bool)`

GetGroupSortOrderOk returns a tuple with the GroupSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSortOrder

`func (o *GroupsGetEntitiesRequest) SetGroupSortOrder(v string)`

SetGroupSortOrder sets GroupSortOrder field to given value.

### HasGroupSortOrder

`func (o *GroupsGetEntitiesRequest) HasGroupSortOrder() bool`

HasGroupSortOrder returns a boolean if a field has been set.

### GetGroupSortDownsampleFunction

`func (o *GroupsGetEntitiesRequest) GetGroupSortDownsampleFunction() string`

GetGroupSortDownsampleFunction returns the GroupSortDownsampleFunction field if non-nil, zero value otherwise.

### GetGroupSortDownsampleFunctionOk

`func (o *GroupsGetEntitiesRequest) GetGroupSortDownsampleFunctionOk() (*string, bool)`

GetGroupSortDownsampleFunctionOk returns a tuple with the GroupSortDownsampleFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSortDownsampleFunction

`func (o *GroupsGetEntitiesRequest) SetGroupSortDownsampleFunction(v string)`

SetGroupSortDownsampleFunction sets GroupSortDownsampleFunction field to given value.

### HasGroupSortDownsampleFunction

`func (o *GroupsGetEntitiesRequest) HasGroupSortDownsampleFunction() bool`

HasGroupSortDownsampleFunction returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *GroupsGetEntitiesRequest) GetFilterCriteria() string`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *GroupsGetEntitiesRequest) GetFilterCriteriaOk() (*string, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *GroupsGetEntitiesRequest) SetFilterCriteria(v string)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *GroupsGetEntitiesRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetLargeBucketBoundary

`func (o *GroupsGetEntitiesRequest) GetLargeBucketBoundary() int64`

GetLargeBucketBoundary returns the LargeBucketBoundary field if non-nil, zero value otherwise.

### GetLargeBucketBoundaryOk

`func (o *GroupsGetEntitiesRequest) GetLargeBucketBoundaryOk() (*int64, bool)`

GetLargeBucketBoundaryOk returns a tuple with the LargeBucketBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeBucketBoundary

`func (o *GroupsGetEntitiesRequest) SetLargeBucketBoundary(v int64)`

SetLargeBucketBoundary sets LargeBucketBoundary field to given value.

### HasLargeBucketBoundary

`func (o *GroupsGetEntitiesRequest) HasLargeBucketBoundary() bool`

HasLargeBucketBoundary returns a boolean if a field has been set.

### GetAvailabilityZoneScope

`func (o *GroupsGetEntitiesRequest) GetAvailabilityZoneScope() string`

GetAvailabilityZoneScope returns the AvailabilityZoneScope field if non-nil, zero value otherwise.

### GetAvailabilityZoneScopeOk

`func (o *GroupsGetEntitiesRequest) GetAvailabilityZoneScopeOk() (*string, bool)`

GetAvailabilityZoneScopeOk returns a tuple with the AvailabilityZoneScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneScope

`func (o *GroupsGetEntitiesRequest) SetAvailabilityZoneScope(v string)`

SetAvailabilityZoneScope sets AvailabilityZoneScope field to given value.

### HasAvailabilityZoneScope

`func (o *GroupsGetEntitiesRequest) HasAvailabilityZoneScope() bool`

HasAvailabilityZoneScope returns a boolean if a field has been set.

### GetGroupCount

`func (o *GroupsGetEntitiesRequest) GetGroupCount() int64`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *GroupsGetEntitiesRequest) GetGroupCountOk() (*int64, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *GroupsGetEntitiesRequest) SetGroupCount(v int64)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *GroupsGetEntitiesRequest) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### GetGroupAttributes

`func (o *GroupsGetEntitiesRequest) GetGroupAttributes() []GroupsRequestedAttribute`

GetGroupAttributes returns the GroupAttributes field if non-nil, zero value otherwise.

### GetGroupAttributesOk

`func (o *GroupsGetEntitiesRequest) GetGroupAttributesOk() (*[]GroupsRequestedAttribute, bool)`

GetGroupAttributesOk returns a tuple with the GroupAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributes

`func (o *GroupsGetEntitiesRequest) SetGroupAttributes(v []GroupsRequestedAttribute)`

SetGroupAttributes sets GroupAttributes field to given value.

### HasGroupAttributes

`func (o *GroupsGetEntitiesRequest) HasGroupAttributes() bool`

HasGroupAttributes returns a boolean if a field has been set.

### GetIntervalEndMs

`func (o *GroupsGetEntitiesRequest) GetIntervalEndMs() int64`

GetIntervalEndMs returns the IntervalEndMs field if non-nil, zero value otherwise.

### GetIntervalEndMsOk

`func (o *GroupsGetEntitiesRequest) GetIntervalEndMsOk() (*int64, bool)`

GetIntervalEndMsOk returns a tuple with the IntervalEndMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEndMs

`func (o *GroupsGetEntitiesRequest) SetIntervalEndMs(v int64)`

SetIntervalEndMs sets IntervalEndMs field to given value.

### HasIntervalEndMs

`func (o *GroupsGetEntitiesRequest) HasIntervalEndMs() bool`

HasIntervalEndMs returns a boolean if a field has been set.

### GetGroupingAttributeType

`func (o *GroupsGetEntitiesRequest) GetGroupingAttributeType() string`

GetGroupingAttributeType returns the GroupingAttributeType field if non-nil, zero value otherwise.

### GetGroupingAttributeTypeOk

`func (o *GroupsGetEntitiesRequest) GetGroupingAttributeTypeOk() (*string, bool)`

GetGroupingAttributeTypeOk returns a tuple with the GroupingAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingAttributeType

`func (o *GroupsGetEntitiesRequest) SetGroupingAttributeType(v string)`

SetGroupingAttributeType sets GroupingAttributeType field to given value.

### HasGroupingAttributeType

`func (o *GroupsGetEntitiesRequest) HasGroupingAttributeType() bool`

HasGroupingAttributeType returns a boolean if a field has been set.

### GetGroupMemberCount

`func (o *GroupsGetEntitiesRequest) GetGroupMemberCount() int64`

GetGroupMemberCount returns the GroupMemberCount field if non-nil, zero value otherwise.

### GetGroupMemberCountOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberCountOk() (*int64, bool)`

GetGroupMemberCountOk returns a tuple with the GroupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberCount

`func (o *GroupsGetEntitiesRequest) SetGroupMemberCount(v int64)`

SetGroupMemberCount sets GroupMemberCount field to given value.

### HasGroupMemberCount

`func (o *GroupsGetEntitiesRequest) HasGroupMemberCount() bool`

HasGroupMemberCount returns a boolean if a field has been set.

### GetGroupMemberSortOrder

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortOrder() string`

GetGroupMemberSortOrder returns the GroupMemberSortOrder field if non-nil, zero value otherwise.

### GetGroupMemberSortOrderOk

`func (o *GroupsGetEntitiesRequest) GetGroupMemberSortOrderOk() (*string, bool)`

GetGroupMemberSortOrderOk returns a tuple with the GroupMemberSortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberSortOrder

`func (o *GroupsGetEntitiesRequest) SetGroupMemberSortOrder(v string)`

SetGroupMemberSortOrder sets GroupMemberSortOrder field to given value.

### HasGroupMemberSortOrder

`func (o *GroupsGetEntitiesRequest) HasGroupMemberSortOrder() bool`

HasGroupMemberSortOrder returns a boolean if a field has been set.

### GetQueryName

`func (o *GroupsGetEntitiesRequest) GetQueryName() string`

GetQueryName returns the QueryName field if non-nil, zero value otherwise.

### GetQueryNameOk

`func (o *GroupsGetEntitiesRequest) GetQueryNameOk() (*string, bool)`

GetQueryNameOk returns a tuple with the QueryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryName

`func (o *GroupsGetEntitiesRequest) SetQueryName(v string)`

SetQueryName sets QueryName field to given value.

### HasQueryName

`func (o *GroupsGetEntitiesRequest) HasQueryName() bool`

HasQueryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


