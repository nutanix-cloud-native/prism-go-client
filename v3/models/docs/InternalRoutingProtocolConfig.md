# InternalRoutingProtocolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPrefixList** | Pointer to [**[]IpSubnet**](IpSubnet.md) | list of local prefixes to be advertised over eBGP | [optional] 
**IbgpConfigList** | Pointer to [**[]BgpConfig**](BgpConfig.md) | iBGP configuration. | [optional] 
**OspfConfig** | Pointer to [**OspfConfig**](OspfConfig.md) |  | [optional] 

## Methods

### NewInternalRoutingProtocolConfig

`func NewInternalRoutingProtocolConfig() *InternalRoutingProtocolConfig`

NewInternalRoutingProtocolConfig instantiates a new InternalRoutingProtocolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalRoutingProtocolConfigWithDefaults

`func NewInternalRoutingProtocolConfigWithDefaults() *InternalRoutingProtocolConfig`

NewInternalRoutingProtocolConfigWithDefaults instantiates a new InternalRoutingProtocolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPrefixList

`func (o *InternalRoutingProtocolConfig) GetLocalPrefixList() []IpSubnet`

GetLocalPrefixList returns the LocalPrefixList field if non-nil, zero value otherwise.

### GetLocalPrefixListOk

`func (o *InternalRoutingProtocolConfig) GetLocalPrefixListOk() (*[]IpSubnet, bool)`

GetLocalPrefixListOk returns a tuple with the LocalPrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPrefixList

`func (o *InternalRoutingProtocolConfig) SetLocalPrefixList(v []IpSubnet)`

SetLocalPrefixList sets LocalPrefixList field to given value.

### HasLocalPrefixList

`func (o *InternalRoutingProtocolConfig) HasLocalPrefixList() bool`

HasLocalPrefixList returns a boolean if a field has been set.

### GetIbgpConfigList

`func (o *InternalRoutingProtocolConfig) GetIbgpConfigList() []BgpConfig`

GetIbgpConfigList returns the IbgpConfigList field if non-nil, zero value otherwise.

### GetIbgpConfigListOk

`func (o *InternalRoutingProtocolConfig) GetIbgpConfigListOk() (*[]BgpConfig, bool)`

GetIbgpConfigListOk returns a tuple with the IbgpConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpConfigList

`func (o *InternalRoutingProtocolConfig) SetIbgpConfigList(v []BgpConfig)`

SetIbgpConfigList sets IbgpConfigList field to given value.

### HasIbgpConfigList

`func (o *InternalRoutingProtocolConfig) HasIbgpConfigList() bool`

HasIbgpConfigList returns a boolean if a field has been set.

### GetOspfConfig

`func (o *InternalRoutingProtocolConfig) GetOspfConfig() OspfConfig`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *InternalRoutingProtocolConfig) GetOspfConfigOk() (*OspfConfig, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *InternalRoutingProtocolConfig) SetOspfConfig(v OspfConfig)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *InternalRoutingProtocolConfig) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


