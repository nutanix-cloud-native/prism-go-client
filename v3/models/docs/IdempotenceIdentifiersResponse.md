# IdempotenceIdentifiersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIdentifier** | Pointer to **string** | The client identifier string. | [optional] 
**Count** | **int32** | The number of idempotence identifiers provided. | [default to 1]
**ExpirationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format of the expiration time (with reference to system time). Value is creation time + valid_duration | [optional] 
**UuidList** | **[]string** |  | 

## Methods

### NewIdempotenceIdentifiersResponse

`func NewIdempotenceIdentifiersResponse(count int32, uuidList []string, ) *IdempotenceIdentifiersResponse`

NewIdempotenceIdentifiersResponse instantiates a new IdempotenceIdentifiersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdempotenceIdentifiersResponseWithDefaults

`func NewIdempotenceIdentifiersResponseWithDefaults() *IdempotenceIdentifiersResponse`

NewIdempotenceIdentifiersResponseWithDefaults instantiates a new IdempotenceIdentifiersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIdentifier

`func (o *IdempotenceIdentifiersResponse) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *IdempotenceIdentifiersResponse) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *IdempotenceIdentifiersResponse) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *IdempotenceIdentifiersResponse) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetCount

`func (o *IdempotenceIdentifiersResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IdempotenceIdentifiersResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IdempotenceIdentifiersResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetExpirationTime

`func (o *IdempotenceIdentifiersResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *IdempotenceIdentifiersResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *IdempotenceIdentifiersResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *IdempotenceIdentifiersResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetUuidList

`func (o *IdempotenceIdentifiersResponse) GetUuidList() []string`

GetUuidList returns the UuidList field if non-nil, zero value otherwise.

### GetUuidListOk

`func (o *IdempotenceIdentifiersResponse) GetUuidListOk() (*[]string, bool)`

GetUuidListOk returns a tuple with the UuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidList

`func (o *IdempotenceIdentifiersResponse) SetUuidList(v []string)`

SetUuidList sets UuidList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


