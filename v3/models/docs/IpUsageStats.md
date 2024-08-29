# IpUsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumFreeIps** | Pointer to **int32** | Number of free IPs in the subnet. | [optional] 
**NumMacs** | Pointer to **int32** | Number of MACs associated with the subnet. | [optional] 
**NumAssignedIps** | Pointer to **int32** | Number of assigned IPs in the subnet. | [optional] 
**IpPoolsStats** | Pointer to [**[]PoolStats**](PoolStats.md) |  | [optional] 

## Methods

### NewIpUsageStats

`func NewIpUsageStats() *IpUsageStats`

NewIpUsageStats instantiates a new IpUsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpUsageStatsWithDefaults

`func NewIpUsageStatsWithDefaults() *IpUsageStats`

NewIpUsageStatsWithDefaults instantiates a new IpUsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumFreeIps

`func (o *IpUsageStats) GetNumFreeIps() int32`

GetNumFreeIps returns the NumFreeIps field if non-nil, zero value otherwise.

### GetNumFreeIpsOk

`func (o *IpUsageStats) GetNumFreeIpsOk() (*int32, bool)`

GetNumFreeIpsOk returns a tuple with the NumFreeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFreeIps

`func (o *IpUsageStats) SetNumFreeIps(v int32)`

SetNumFreeIps sets NumFreeIps field to given value.

### HasNumFreeIps

`func (o *IpUsageStats) HasNumFreeIps() bool`

HasNumFreeIps returns a boolean if a field has been set.

### GetNumMacs

`func (o *IpUsageStats) GetNumMacs() int32`

GetNumMacs returns the NumMacs field if non-nil, zero value otherwise.

### GetNumMacsOk

`func (o *IpUsageStats) GetNumMacsOk() (*int32, bool)`

GetNumMacsOk returns a tuple with the NumMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMacs

`func (o *IpUsageStats) SetNumMacs(v int32)`

SetNumMacs sets NumMacs field to given value.

### HasNumMacs

`func (o *IpUsageStats) HasNumMacs() bool`

HasNumMacs returns a boolean if a field has been set.

### GetNumAssignedIps

`func (o *IpUsageStats) GetNumAssignedIps() int32`

GetNumAssignedIps returns the NumAssignedIps field if non-nil, zero value otherwise.

### GetNumAssignedIpsOk

`func (o *IpUsageStats) GetNumAssignedIpsOk() (*int32, bool)`

GetNumAssignedIpsOk returns a tuple with the NumAssignedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssignedIps

`func (o *IpUsageStats) SetNumAssignedIps(v int32)`

SetNumAssignedIps sets NumAssignedIps field to given value.

### HasNumAssignedIps

`func (o *IpUsageStats) HasNumAssignedIps() bool`

HasNumAssignedIps returns a boolean if a field has been set.

### GetIpPoolsStats

`func (o *IpUsageStats) GetIpPoolsStats() []PoolStats`

GetIpPoolsStats returns the IpPoolsStats field if non-nil, zero value otherwise.

### GetIpPoolsStatsOk

`func (o *IpUsageStats) GetIpPoolsStatsOk() (*[]PoolStats, bool)`

GetIpPoolsStatsOk returns a tuple with the IpPoolsStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolsStats

`func (o *IpUsageStats) SetIpPoolsStats(v []PoolStats)`

SetIpPoolsStats sets IpPoolsStats field to given value.

### HasIpPoolsStats

`func (o *IpUsageStats) HasIpPoolsStats() bool`

HasIpPoolsStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


