# IpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Address string. | [optional] 
**Type** | Pointer to **string** | Address type. It can only be \&quot;ASSIGNED\&quot; in the spec. If no type is specified in the spec, the default type is set to \&quot;ASSIGNED\&quot;.  | [optional] 
**GatewayAddressList** | Pointer to **[]string** | Gateway IP addresses matching the subnet. | [optional] 
**PrefixLength** | Pointer to **int32** | Prefix length for the IP address. | [optional] 
**IpType** | Pointer to **string** | Indicates whether IP address is DHCP or Static. | [optional] 

## Methods

### NewIpAddress

`func NewIpAddress() *IpAddress`

NewIpAddress instantiates a new IpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressWithDefaults

`func NewIpAddressWithDefaults() *IpAddress`

NewIpAddressWithDefaults instantiates a new IpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpAddress) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpAddress) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpAddress) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IpAddress) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetType

`func (o *IpAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGatewayAddressList

`func (o *IpAddress) GetGatewayAddressList() []string`

GetGatewayAddressList returns the GatewayAddressList field if non-nil, zero value otherwise.

### GetGatewayAddressListOk

`func (o *IpAddress) GetGatewayAddressListOk() (*[]string, bool)`

GetGatewayAddressListOk returns a tuple with the GatewayAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddressList

`func (o *IpAddress) SetGatewayAddressList(v []string)`

SetGatewayAddressList sets GatewayAddressList field to given value.

### HasGatewayAddressList

`func (o *IpAddress) HasGatewayAddressList() bool`

HasGatewayAddressList returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpAddress) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpAddress) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpAddress) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *IpAddress) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetIpType

`func (o *IpAddress) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *IpAddress) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *IpAddress) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *IpAddress) HasIpType() bool`

HasIpType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


