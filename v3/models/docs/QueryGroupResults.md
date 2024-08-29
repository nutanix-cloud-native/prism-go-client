# QueryGroupResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityResults** | Pointer to [**[]QueryGroupsEntity**](QueryGroupsEntity.md) |  | [optional] 
**TotalEntityCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryGroupResults

`func NewQueryGroupResults() *QueryGroupResults`

NewQueryGroupResults instantiates a new QueryGroupResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGroupResultsWithDefaults

`func NewQueryGroupResultsWithDefaults() *QueryGroupResults`

NewQueryGroupResultsWithDefaults instantiates a new QueryGroupResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityResults

`func (o *QueryGroupResults) GetEntityResults() []QueryGroupsEntity`

GetEntityResults returns the EntityResults field if non-nil, zero value otherwise.

### GetEntityResultsOk

`func (o *QueryGroupResults) GetEntityResultsOk() (*[]QueryGroupsEntity, bool)`

GetEntityResultsOk returns a tuple with the EntityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityResults

`func (o *QueryGroupResults) SetEntityResults(v []QueryGroupsEntity)`

SetEntityResults sets EntityResults field to given value.

### HasEntityResults

`func (o *QueryGroupResults) HasEntityResults() bool`

HasEntityResults returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *QueryGroupResults) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *QueryGroupResults) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *QueryGroupResults) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *QueryGroupResults) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


