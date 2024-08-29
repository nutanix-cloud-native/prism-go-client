# IpSubnetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IPV4 address. | [optional] 
**PrefixLength** | Pointer to **int32** |  | [optional] 

## Methods

### NewIpSubnetStatus

`func NewIpSubnetStatus() *IpSubnetStatus`

NewIpSubnetStatus instantiates a new IpSubnetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpSubnetStatusWithDefaults

`func NewIpSubnetStatusWithDefaults() *IpSubnetStatus`

NewIpSubnetStatusWithDefaults instantiates a new IpSubnetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpSubnetStatus) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpSubnetStatus) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpSubnetStatus) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IpSubnetStatus) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpSubnetStatus) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpSubnetStatus) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpSubnetStatus) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *IpSubnetStatus) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


