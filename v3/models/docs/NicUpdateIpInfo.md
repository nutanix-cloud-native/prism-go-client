# NicUpdateIpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | If specified, try to request this IP address. Otherwise request any IP address from the network pool.  | [optional] 
**UUID** | **string** | UUID of the NIC whose IP is being requested. | 

## Methods

### NewNicUpdateIpInfo

`func NewNicUpdateIpInfo(uUID string, ) *NicUpdateIpInfo`

NewNicUpdateIpInfo instantiates a new NicUpdateIpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicUpdateIpInfoWithDefaults

`func NewNicUpdateIpInfoWithDefaults() *NicUpdateIpInfo`

NewNicUpdateIpInfoWithDefaults instantiates a new NicUpdateIpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NicUpdateIpInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NicUpdateIpInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NicUpdateIpInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NicUpdateIpInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUUID

`func (o *NicUpdateIpInfo) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *NicUpdateIpInfo) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *NicUpdateIpInfo) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


