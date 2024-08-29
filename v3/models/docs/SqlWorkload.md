# SqlWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | Pointer to **string** |  | [optional] 
**BusinessCritical** | Pointer to **bool** |  | [optional] 
**SqlProfileType** | Pointer to **string** |  | [optional] 
**NumDb** | Pointer to **int32** |  | [optional] 

## Methods

### NewSqlWorkload

`func NewSqlWorkload() *SqlWorkload`

NewSqlWorkload instantiates a new SqlWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlWorkloadWithDefaults

`func NewSqlWorkloadWithDefaults() *SqlWorkload`

NewSqlWorkloadWithDefaults instantiates a new SqlWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *SqlWorkload) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *SqlWorkload) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *SqlWorkload) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *SqlWorkload) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetBusinessCritical

`func (o *SqlWorkload) GetBusinessCritical() bool`

GetBusinessCritical returns the BusinessCritical field if non-nil, zero value otherwise.

### GetBusinessCriticalOk

`func (o *SqlWorkload) GetBusinessCriticalOk() (*bool, bool)`

GetBusinessCriticalOk returns a tuple with the BusinessCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCritical

`func (o *SqlWorkload) SetBusinessCritical(v bool)`

SetBusinessCritical sets BusinessCritical field to given value.

### HasBusinessCritical

`func (o *SqlWorkload) HasBusinessCritical() bool`

HasBusinessCritical returns a boolean if a field has been set.

### GetSqlProfileType

`func (o *SqlWorkload) GetSqlProfileType() string`

GetSqlProfileType returns the SqlProfileType field if non-nil, zero value otherwise.

### GetSqlProfileTypeOk

`func (o *SqlWorkload) GetSqlProfileTypeOk() (*string, bool)`

GetSqlProfileTypeOk returns a tuple with the SqlProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlProfileType

`func (o *SqlWorkload) SetSqlProfileType(v string)`

SetSqlProfileType sets SqlProfileType field to given value.

### HasSqlProfileType

`func (o *SqlWorkload) HasSqlProfileType() bool`

HasSqlProfileType returns a boolean if a field has been set.

### GetNumDb

`func (o *SqlWorkload) GetNumDb() int32`

GetNumDb returns the NumDb field if non-nil, zero value otherwise.

### GetNumDbOk

`func (o *SqlWorkload) GetNumDbOk() (*int32, bool)`

GetNumDbOk returns a tuple with the NumDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDb

`func (o *SqlWorkload) SetNumDb(v int32)`

SetNumDb sets NumDb field to given value.

### HasNumDb

`func (o *SqlWorkload) HasNumDb() bool`

HasNumDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


