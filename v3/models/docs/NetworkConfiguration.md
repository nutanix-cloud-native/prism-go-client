# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsIpAddressesList** | Pointer to **[]string** | List of DNS IP addresses. | [optional] 
**Name** | Pointer to **string** | Name of the network interface. | [optional] 
**MacAddress** | Pointer to **string** | MAC Address of the network interface. | [optional] 
**DhcpServerIp** | Pointer to **string** | IP address of the DHCP server. | [optional] 
**DefaultGatewayAddressList** | Pointer to **[]string** | Default gateway IP addresses. | [optional] 
**IpInfoList** | Pointer to [**[]IpAddress**](IpAddress.md) | List of IP information of the network interface. | [optional] 

## Methods

### NewNetworkConfiguration

`func NewNetworkConfiguration() *NetworkConfiguration`

NewNetworkConfiguration instantiates a new NetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationWithDefaults

`func NewNetworkConfigurationWithDefaults() *NetworkConfiguration`

NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsIpAddressesList

`func (o *NetworkConfiguration) GetDnsIpAddressesList() []string`

GetDnsIpAddressesList returns the DnsIpAddressesList field if non-nil, zero value otherwise.

### GetDnsIpAddressesListOk

`func (o *NetworkConfiguration) GetDnsIpAddressesListOk() (*[]string, bool)`

GetDnsIpAddressesListOk returns a tuple with the DnsIpAddressesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIpAddressesList

`func (o *NetworkConfiguration) SetDnsIpAddressesList(v []string)`

SetDnsIpAddressesList sets DnsIpAddressesList field to given value.

### HasDnsIpAddressesList

`func (o *NetworkConfiguration) HasDnsIpAddressesList() bool`

HasDnsIpAddressesList returns a boolean if a field has been set.

### GetName

`func (o *NetworkConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkConfiguration) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkConfiguration) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkConfiguration) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkConfiguration) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDhcpServerIp

`func (o *NetworkConfiguration) GetDhcpServerIp() string`

GetDhcpServerIp returns the DhcpServerIp field if non-nil, zero value otherwise.

### GetDhcpServerIpOk

`func (o *NetworkConfiguration) GetDhcpServerIpOk() (*string, bool)`

GetDhcpServerIpOk returns a tuple with the DhcpServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIp

`func (o *NetworkConfiguration) SetDhcpServerIp(v string)`

SetDhcpServerIp sets DhcpServerIp field to given value.

### HasDhcpServerIp

`func (o *NetworkConfiguration) HasDhcpServerIp() bool`

HasDhcpServerIp returns a boolean if a field has been set.

### GetDefaultGatewayAddressList

`func (o *NetworkConfiguration) GetDefaultGatewayAddressList() []string`

GetDefaultGatewayAddressList returns the DefaultGatewayAddressList field if non-nil, zero value otherwise.

### GetDefaultGatewayAddressListOk

`func (o *NetworkConfiguration) GetDefaultGatewayAddressListOk() (*[]string, bool)`

GetDefaultGatewayAddressListOk returns a tuple with the DefaultGatewayAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAddressList

`func (o *NetworkConfiguration) SetDefaultGatewayAddressList(v []string)`

SetDefaultGatewayAddressList sets DefaultGatewayAddressList field to given value.

### HasDefaultGatewayAddressList

`func (o *NetworkConfiguration) HasDefaultGatewayAddressList() bool`

HasDefaultGatewayAddressList returns a boolean if a field has been set.

### GetIpInfoList

`func (o *NetworkConfiguration) GetIpInfoList() []IpAddress`

GetIpInfoList returns the IpInfoList field if non-nil, zero value otherwise.

### GetIpInfoListOk

`func (o *NetworkConfiguration) GetIpInfoListOk() (*[]IpAddress, bool)`

GetIpInfoListOk returns a tuple with the IpInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInfoList

`func (o *NetworkConfiguration) SetIpInfoList(v []IpAddress)`

SetIpInfoList sets IpInfoList field to given value.

### HasIpInfoList

`func (o *NetworkConfiguration) HasIpInfoList() bool`

HasIpInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


