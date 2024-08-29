# TopologyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableOneHopTraffic** | Pointer to **bool** | Toggle to view one hop traffic | [optional] 
**FilterCriteria** | Pointer to **string** | FIQL filter criteria that will be used to form the filter field for the backend queries.  | [optional] 
**IntervalEndMs** | Pointer to **int64** | For a time-series query, the end of the interval since epoch in ms.  | [optional] 
**IntervalStartMs** | Pointer to **int64** | For a time-series query, the start of the interval since epoch in ms.  | [optional] 
**GroupBy** | Pointer to **string** | GroupBy field | [optional] 

## Methods

### NewTopologyRequest

`func NewTopologyRequest() *TopologyRequest`

NewTopologyRequest instantiates a new TopologyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyRequestWithDefaults

`func NewTopologyRequestWithDefaults() *TopologyRequest`

NewTopologyRequestWithDefaults instantiates a new TopologyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableOneHopTraffic

`func (o *TopologyRequest) GetEnableOneHopTraffic() bool`

GetEnableOneHopTraffic returns the EnableOneHopTraffic field if non-nil, zero value otherwise.

### GetEnableOneHopTrafficOk

`func (o *TopologyRequest) GetEnableOneHopTrafficOk() (*bool, bool)`

GetEnableOneHopTrafficOk returns a tuple with the EnableOneHopTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOneHopTraffic

`func (o *TopologyRequest) SetEnableOneHopTraffic(v bool)`

SetEnableOneHopTraffic sets EnableOneHopTraffic field to given value.

### HasEnableOneHopTraffic

`func (o *TopologyRequest) HasEnableOneHopTraffic() bool`

HasEnableOneHopTraffic returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *TopologyRequest) GetFilterCriteria() string`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *TopologyRequest) GetFilterCriteriaOk() (*string, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *TopologyRequest) SetFilterCriteria(v string)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *TopologyRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetIntervalEndMs

`func (o *TopologyRequest) GetIntervalEndMs() int64`

GetIntervalEndMs returns the IntervalEndMs field if non-nil, zero value otherwise.

### GetIntervalEndMsOk

`func (o *TopologyRequest) GetIntervalEndMsOk() (*int64, bool)`

GetIntervalEndMsOk returns a tuple with the IntervalEndMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalEndMs

`func (o *TopologyRequest) SetIntervalEndMs(v int64)`

SetIntervalEndMs sets IntervalEndMs field to given value.

### HasIntervalEndMs

`func (o *TopologyRequest) HasIntervalEndMs() bool`

HasIntervalEndMs returns a boolean if a field has been set.

### GetIntervalStartMs

`func (o *TopologyRequest) GetIntervalStartMs() int64`

GetIntervalStartMs returns the IntervalStartMs field if non-nil, zero value otherwise.

### GetIntervalStartMsOk

`func (o *TopologyRequest) GetIntervalStartMsOk() (*int64, bool)`

GetIntervalStartMsOk returns a tuple with the IntervalStartMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalStartMs

`func (o *TopologyRequest) SetIntervalStartMs(v int64)`

SetIntervalStartMs sets IntervalStartMs field to given value.

### HasIntervalStartMs

`func (o *TopologyRequest) HasIntervalStartMs() bool`

HasIntervalStartMs returns a boolean if a field has been set.

### GetGroupBy

`func (o *TopologyRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *TopologyRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *TopologyRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *TopologyRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


