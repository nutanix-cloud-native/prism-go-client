# QueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteredEntityCount** | Pointer to **int64** |  | [optional] 
**GroupResults** | Pointer to [**[]QueryGroupResults**](QueryGroupResults.md) |  | [optional] 
**TotalGroupCount** | Pointer to **int64** |  | [optional] 
**TotalEntityCount** | Pointer to **int64** |  | [optional] 
**EntityType** | Pointer to **string** | The entity type for the query response. | [optional] 

## Methods

### NewQueryResponse

`func NewQueryResponse() *QueryResponse`

NewQueryResponse instantiates a new QueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseWithDefaults

`func NewQueryResponseWithDefaults() *QueryResponse`

NewQueryResponseWithDefaults instantiates a new QueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteredEntityCount

`func (o *QueryResponse) GetFilteredEntityCount() int64`

GetFilteredEntityCount returns the FilteredEntityCount field if non-nil, zero value otherwise.

### GetFilteredEntityCountOk

`func (o *QueryResponse) GetFilteredEntityCountOk() (*int64, bool)`

GetFilteredEntityCountOk returns a tuple with the FilteredEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredEntityCount

`func (o *QueryResponse) SetFilteredEntityCount(v int64)`

SetFilteredEntityCount sets FilteredEntityCount field to given value.

### HasFilteredEntityCount

`func (o *QueryResponse) HasFilteredEntityCount() bool`

HasFilteredEntityCount returns a boolean if a field has been set.

### GetGroupResults

`func (o *QueryResponse) GetGroupResults() []QueryGroupResults`

GetGroupResults returns the GroupResults field if non-nil, zero value otherwise.

### GetGroupResultsOk

`func (o *QueryResponse) GetGroupResultsOk() (*[]QueryGroupResults, bool)`

GetGroupResultsOk returns a tuple with the GroupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupResults

`func (o *QueryResponse) SetGroupResults(v []QueryGroupResults)`

SetGroupResults sets GroupResults field to given value.

### HasGroupResults

`func (o *QueryResponse) HasGroupResults() bool`

HasGroupResults returns a boolean if a field has been set.

### GetTotalGroupCount

`func (o *QueryResponse) GetTotalGroupCount() int64`

GetTotalGroupCount returns the TotalGroupCount field if non-nil, zero value otherwise.

### GetTotalGroupCountOk

`func (o *QueryResponse) GetTotalGroupCountOk() (*int64, bool)`

GetTotalGroupCountOk returns a tuple with the TotalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGroupCount

`func (o *QueryResponse) SetTotalGroupCount(v int64)`

SetTotalGroupCount sets TotalGroupCount field to given value.

### HasTotalGroupCount

`func (o *QueryResponse) HasTotalGroupCount() bool`

HasTotalGroupCount returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *QueryResponse) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *QueryResponse) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *QueryResponse) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *QueryResponse) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.

### GetEntityType

`func (o *QueryResponse) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *QueryResponse) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *QueryResponse) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *QueryResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


