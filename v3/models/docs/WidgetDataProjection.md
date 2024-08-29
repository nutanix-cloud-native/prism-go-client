# WidgetDataProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterCriteria** | Pointer to **string** | FIQL criteria that will be used to filter the returned data. | [optional] 
**SortKey** | Pointer to **string** | Aggregation to be used while sorting time-series data. | [optional] 
**SortColumn** | Pointer to **string** | Entity Property based on which the result data is to be sorted. | [optional] 
**SortOrder** | Pointer to **string** | Order of sorting. | [optional] 
**Limit** | Pointer to **int64** | Limit on the maximum number of entities to be represented in the widget. A limit value of more than 10 would not be entertained.  | [optional] 
**CustomKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes. | [optional] 

## Methods

### NewWidgetDataProjection

`func NewWidgetDataProjection() *WidgetDataProjection`

NewWidgetDataProjection instantiates a new WidgetDataProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetDataProjectionWithDefaults

`func NewWidgetDataProjectionWithDefaults() *WidgetDataProjection`

NewWidgetDataProjectionWithDefaults instantiates a new WidgetDataProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterCriteria

`func (o *WidgetDataProjection) GetFilterCriteria() string`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *WidgetDataProjection) GetFilterCriteriaOk() (*string, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *WidgetDataProjection) SetFilterCriteria(v string)`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *WidgetDataProjection) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.

### GetSortKey

`func (o *WidgetDataProjection) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *WidgetDataProjection) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *WidgetDataProjection) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *WidgetDataProjection) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortColumn

`func (o *WidgetDataProjection) GetSortColumn() string`

GetSortColumn returns the SortColumn field if non-nil, zero value otherwise.

### GetSortColumnOk

`func (o *WidgetDataProjection) GetSortColumnOk() (*string, bool)`

GetSortColumnOk returns a tuple with the SortColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumn

`func (o *WidgetDataProjection) SetSortColumn(v string)`

SetSortColumn sets SortColumn field to given value.

### HasSortColumn

`func (o *WidgetDataProjection) HasSortColumn() bool`

HasSortColumn returns a boolean if a field has been set.

### GetSortOrder

`func (o *WidgetDataProjection) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *WidgetDataProjection) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *WidgetDataProjection) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *WidgetDataProjection) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLimit

`func (o *WidgetDataProjection) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WidgetDataProjection) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WidgetDataProjection) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WidgetDataProjection) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCustomKeyValues

`func (o *WidgetDataProjection) GetCustomKeyValues() map[string]string`

GetCustomKeyValues returns the CustomKeyValues field if non-nil, zero value otherwise.

### GetCustomKeyValuesOk

`func (o *WidgetDataProjection) GetCustomKeyValuesOk() (*map[string]string, bool)`

GetCustomKeyValuesOk returns a tuple with the CustomKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyValues

`func (o *WidgetDataProjection) SetCustomKeyValues(v map[string]string)`

SetCustomKeyValues sets CustomKeyValues field to given value.

### HasCustomKeyValues

`func (o *WidgetDataProjection) HasCustomKeyValues() bool`

HasCustomKeyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


