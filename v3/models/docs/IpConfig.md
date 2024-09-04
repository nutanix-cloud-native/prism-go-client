# IpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGatewayIp** | Pointer to **string** | Default gateway IP address. | [optional] 
**DhcpServerAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**PoolList** | Pointer to [**[]IpPool**](IpPool.md) |  | [optional] 
**PrefixLength** | Pointer to **int32** |  | [optional] 
**SubnetIp** | Pointer to **string** | Subnet IP address. | [optional] 
**DhcpOptions** | Pointer to [**DhcpOptions**](DhcpOptions.md) |  | [optional] 

## Methods

### NewIpConfig

`func NewIpConfig() *IpConfig`

NewIpConfig instantiates a new IpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpConfigWithDefaults

`func NewIpConfigWithDefaults() *IpConfig`

NewIpConfigWithDefaults instantiates a new IpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGatewayIp

`func (o *IpConfig) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *IpConfig) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *IpConfig) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *IpConfig) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### GetDhcpServerAddress

`func (o *IpConfig) GetDhcpServerAddress() Address`

GetDhcpServerAddress returns the DhcpServerAddress field if non-nil, zero value otherwise.

### GetDhcpServerAddressOk

`func (o *IpConfig) GetDhcpServerAddressOk() (*Address, bool)`

GetDhcpServerAddressOk returns a tuple with the DhcpServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerAddress

`func (o *IpConfig) SetDhcpServerAddress(v Address)`

SetDhcpServerAddress sets DhcpServerAddress field to given value.

### HasDhcpServerAddress

`func (o *IpConfig) HasDhcpServerAddress() bool`

HasDhcpServerAddress returns a boolean if a field has been set.

### GetPoolList

`func (o *IpConfig) GetPoolList() []IpPool`

GetPoolList returns the PoolList field if non-nil, zero value otherwise.

### GetPoolListOk

`func (o *IpConfig) GetPoolListOk() (*[]IpPool, bool)`

GetPoolListOk returns a tuple with the PoolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolList

`func (o *IpConfig) SetPoolList(v []IpPool)`

SetPoolList sets PoolList field to given value.

### HasPoolList

`func (o *IpConfig) HasPoolList() bool`

HasPoolList returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpConfig) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpConfig) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpConfig) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *IpConfig) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetSubnetIp

`func (o *IpConfig) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *IpConfig) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *IpConfig) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.

### HasSubnetIp

`func (o *IpConfig) HasSubnetIp() bool`

HasSubnetIp returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *IpConfig) GetDhcpOptions() DhcpOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpConfig) GetDhcpOptionsOk() (*DhcpOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpConfig) SetDhcpOptions(v DhcpOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpConfig) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


