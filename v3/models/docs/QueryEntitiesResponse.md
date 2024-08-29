# QueryEntitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResults** | Pointer to **int32** | Total number of entities. | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset. | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list. | [optional] 
**EntityList** | Pointer to [**[]QueryEntitiesResponseEntityListInner**](QueryEntitiesResponseEntityListInner.md) |  | [optional] 

## Methods

### NewQueryEntitiesResponse

`func NewQueryEntitiesResponse() *QueryEntitiesResponse`

NewQueryEntitiesResponse instantiates a new QueryEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryEntitiesResponseWithDefaults

`func NewQueryEntitiesResponseWithDefaults() *QueryEntitiesResponse`

NewQueryEntitiesResponseWithDefaults instantiates a new QueryEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResults

`func (o *QueryEntitiesResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *QueryEntitiesResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *QueryEntitiesResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *QueryEntitiesResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetLength

`func (o *QueryEntitiesResponse) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *QueryEntitiesResponse) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *QueryEntitiesResponse) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *QueryEntitiesResponse) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetOffset

`func (o *QueryEntitiesResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryEntitiesResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryEntitiesResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryEntitiesResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetEntityList

`func (o *QueryEntitiesResponse) GetEntityList() []QueryEntitiesResponseEntityListInner`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *QueryEntitiesResponse) GetEntityListOk() (*[]QueryEntitiesResponseEntityListInner, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *QueryEntitiesResponse) SetEntityList(v []QueryEntitiesResponseEntityListInner)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *QueryEntitiesResponse) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


