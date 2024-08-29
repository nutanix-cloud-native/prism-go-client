# IdempotenceIdentifiersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIdentifier** | Pointer to **string** | The client identifier string. | [optional] 
**Count** | **int32** | The number of idempotence identifiers provided. | [default to 1]
**ValidDurationInMinutes** | Pointer to **int32** | Number of minutes from creation time for which idempotence identifier uuid list is valid. | [optional] [default to 527040]

## Methods

### NewIdempotenceIdentifiersInput

`func NewIdempotenceIdentifiersInput(count int32, ) *IdempotenceIdentifiersInput`

NewIdempotenceIdentifiersInput instantiates a new IdempotenceIdentifiersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdempotenceIdentifiersInputWithDefaults

`func NewIdempotenceIdentifiersInputWithDefaults() *IdempotenceIdentifiersInput`

NewIdempotenceIdentifiersInputWithDefaults instantiates a new IdempotenceIdentifiersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIdentifier

`func (o *IdempotenceIdentifiersInput) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *IdempotenceIdentifiersInput) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *IdempotenceIdentifiersInput) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *IdempotenceIdentifiersInput) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetCount

`func (o *IdempotenceIdentifiersInput) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IdempotenceIdentifiersInput) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IdempotenceIdentifiersInput) SetCount(v int32)`

SetCount sets Count field to given value.


### GetValidDurationInMinutes

`func (o *IdempotenceIdentifiersInput) GetValidDurationInMinutes() int32`

GetValidDurationInMinutes returns the ValidDurationInMinutes field if non-nil, zero value otherwise.

### GetValidDurationInMinutesOk

`func (o *IdempotenceIdentifiersInput) GetValidDurationInMinutesOk() (*int32, bool)`

GetValidDurationInMinutesOk returns a tuple with the ValidDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDurationInMinutes

`func (o *IdempotenceIdentifiersInput) SetValidDurationInMinutes(v int32)`

SetValidDurationInMinutes sets ValidDurationInMinutes field to given value.

### HasValidDurationInMinutes

`func (o *IdempotenceIdentifiersInput) HasValidDurationInMinutes() bool`

HasValidDurationInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


