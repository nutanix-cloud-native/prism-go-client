# TimeseriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the timeseries response. | [optional] 
**Result** | Pointer to **map[string]interface{}** | Group query object definition. Making it opaque in the first release. This will have the details of the parameters which are present in group query.  | [optional] 

## Methods

### NewTimeseriesResponse

`func NewTimeseriesResponse() *TimeseriesResponse`

NewTimeseriesResponse instantiates a new TimeseriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesResponseWithDefaults

`func NewTimeseriesResponseWithDefaults() *TimeseriesResponse`

NewTimeseriesResponseWithDefaults instantiates a new TimeseriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeseriesResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeseriesResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeseriesResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeseriesResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *TimeseriesResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TimeseriesResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TimeseriesResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *TimeseriesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


