# CmspNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network ID for CMSP. UUID for AHV or name for ESXi.  | [optional] 
**SubnetMask** | **string** | Subnet mask IP address. | 
**Type** | Pointer to **string** | Network type. | [optional] [default to "kFull"]
**NetworkName** | Pointer to **string** | Network name for cmsp cluster. | [optional] 
**DefaultGateway** | **string** | Gateway IP address. | 

## Methods

### NewCmspNetworkConfig

`func NewCmspNetworkConfig(subnetMask string, defaultGateway string, ) *CmspNetworkConfig`

NewCmspNetworkConfig instantiates a new CmspNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmspNetworkConfigWithDefaults

`func NewCmspNetworkConfigWithDefaults() *CmspNetworkConfig`

NewCmspNetworkConfigWithDefaults instantiates a new CmspNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *CmspNetworkConfig) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CmspNetworkConfig) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CmspNetworkConfig) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CmspNetworkConfig) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubnetMask

`func (o *CmspNetworkConfig) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *CmspNetworkConfig) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *CmspNetworkConfig) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.


### GetType

`func (o *CmspNetworkConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CmspNetworkConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CmspNetworkConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CmspNetworkConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkName

`func (o *CmspNetworkConfig) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CmspNetworkConfig) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CmspNetworkConfig) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *CmspNetworkConfig) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *CmspNetworkConfig) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *CmspNetworkConfig) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *CmspNetworkConfig) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


