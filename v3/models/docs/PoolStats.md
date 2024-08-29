# PoolStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumTotalIps** | Pointer to **int32** | Number of total IPs in the pool. | [optional] 
**Range** | Pointer to **string** | Range of IPs (example: 10.0.0.9 10.0.0.19).  | [optional] 
**NumFreeIps** | Pointer to **int32** | Number of free IPs in the pool. | [optional] 

## Methods

### NewPoolStats

`func NewPoolStats() *PoolStats`

NewPoolStats instantiates a new PoolStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolStatsWithDefaults

`func NewPoolStatsWithDefaults() *PoolStats`

NewPoolStatsWithDefaults instantiates a new PoolStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumTotalIps

`func (o *PoolStats) GetNumTotalIps() int32`

GetNumTotalIps returns the NumTotalIps field if non-nil, zero value otherwise.

### GetNumTotalIpsOk

`func (o *PoolStats) GetNumTotalIpsOk() (*int32, bool)`

GetNumTotalIpsOk returns a tuple with the NumTotalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalIps

`func (o *PoolStats) SetNumTotalIps(v int32)`

SetNumTotalIps sets NumTotalIps field to given value.

### HasNumTotalIps

`func (o *PoolStats) HasNumTotalIps() bool`

HasNumTotalIps returns a boolean if a field has been set.

### GetRange

`func (o *PoolStats) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *PoolStats) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *PoolStats) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *PoolStats) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetNumFreeIps

`func (o *PoolStats) GetNumFreeIps() int32`

GetNumFreeIps returns the NumFreeIps field if non-nil, zero value otherwise.

### GetNumFreeIpsOk

`func (o *PoolStats) GetNumFreeIpsOk() (*int32, bool)`

GetNumFreeIpsOk returns a tuple with the NumFreeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFreeIps

`func (o *PoolStats) SetNumFreeIps(v int32)`

SetNumFreeIps sets NumFreeIps field to given value.

### HasNumFreeIps

`func (o *PoolStats) HasNumFreeIps() bool`

HasNumFreeIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


