# NetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetMask** | Pointer to **string** | Subnet mask IP address. | [optional] 
**NetworkUuid** | Pointer to **string** | Network uuid. | [optional] 
**NetworkName** | Pointer to **string** | Network name. | [optional] 
**DefaultGateway** | Pointer to **string** | Gateway IP address. | [optional] 

## Methods

### NewNetworkConfig

`func NewNetworkConfig() *NetworkConfig`

NewNetworkConfig instantiates a new NetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigWithDefaults

`func NewNetworkConfigWithDefaults() *NetworkConfig`

NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetMask

`func (o *NetworkConfig) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *NetworkConfig) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *NetworkConfig) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *NetworkConfig) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetNetworkUuid

`func (o *NetworkConfig) GetNetworkUuid() string`

GetNetworkUuid returns the NetworkUuid field if non-nil, zero value otherwise.

### GetNetworkUuidOk

`func (o *NetworkConfig) GetNetworkUuidOk() (*string, bool)`

GetNetworkUuidOk returns a tuple with the NetworkUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUuid

`func (o *NetworkConfig) SetNetworkUuid(v string)`

SetNetworkUuid sets NetworkUuid field to given value.

### HasNetworkUuid

`func (o *NetworkConfig) HasNetworkUuid() bool`

HasNetworkUuid returns a boolean if a field has been set.

### GetNetworkName

`func (o *NetworkConfig) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *NetworkConfig) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *NetworkConfig) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *NetworkConfig) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *NetworkConfig) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *NetworkConfig) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *NetworkConfig) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *NetworkConfig) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


