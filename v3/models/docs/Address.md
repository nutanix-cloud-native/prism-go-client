# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IPV4 address. | [optional] 
**IsBackup** | Pointer to **bool** | Whether this address is a backup or not. | [optional] 
**Fqdn** | Pointer to **string** | Fully qualified domain name. | [optional] 
**Ip6Range** | Pointer to **string** | IPV6 address range. | [optional] 
**IpRange** | Pointer to **string** | IPV4 address range. | [optional] 
**Ipv6** | Pointer to **string** | IPV6 address. | [optional] 
**Port** | Pointer to **int32** | Port Number | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *Address) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Address) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Address) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Address) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsBackup

`func (o *Address) GetIsBackup() bool`

GetIsBackup returns the IsBackup field if non-nil, zero value otherwise.

### GetIsBackupOk

`func (o *Address) GetIsBackupOk() (*bool, bool)`

GetIsBackupOk returns a tuple with the IsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackup

`func (o *Address) SetIsBackup(v bool)`

SetIsBackup sets IsBackup field to given value.

### HasIsBackup

`func (o *Address) HasIsBackup() bool`

HasIsBackup returns a boolean if a field has been set.

### GetFqdn

`func (o *Address) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Address) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Address) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Address) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetIp6Range

`func (o *Address) GetIp6Range() string`

GetIp6Range returns the Ip6Range field if non-nil, zero value otherwise.

### GetIp6RangeOk

`func (o *Address) GetIp6RangeOk() (*string, bool)`

GetIp6RangeOk returns a tuple with the Ip6Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Range

`func (o *Address) SetIp6Range(v string)`

SetIp6Range sets Ip6Range field to given value.

### HasIp6Range

`func (o *Address) HasIp6Range() bool`

HasIp6Range returns a boolean if a field has been set.

### GetIpRange

`func (o *Address) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *Address) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *Address) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *Address) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetIpv6

`func (o *Address) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Address) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Address) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Address) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetPort

`func (o *Address) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Address) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Address) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Address) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


