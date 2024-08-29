# NetworkAddressStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSubnet** | Pointer to [**IpSubnetStatus**](IpSubnetStatus.md) |  | [optional] 
**AddressType** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkAddressStatus

`func NewNetworkAddressStatus() *NetworkAddressStatus`

NewNetworkAddressStatus instantiates a new NetworkAddressStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAddressStatusWithDefaults

`func NewNetworkAddressStatusWithDefaults() *NetworkAddressStatus`

NewNetworkAddressStatusWithDefaults instantiates a new NetworkAddressStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSubnet

`func (o *NetworkAddressStatus) GetIpSubnet() IpSubnetStatus`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *NetworkAddressStatus) GetIpSubnetOk() (*IpSubnetStatus, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *NetworkAddressStatus) SetIpSubnet(v IpSubnetStatus)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *NetworkAddressStatus) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetAddressType

`func (o *NetworkAddressStatus) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *NetworkAddressStatus) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *NetworkAddressStatus) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *NetworkAddressStatus) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


