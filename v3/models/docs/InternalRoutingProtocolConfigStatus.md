# InternalRoutingProtocolConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPrefixList** | Pointer to [**[]IpSubnetStatus**](IpSubnetStatus.md) | list of local prefixes to be advertised over eBGP | [optional] 
**IbgpConfigList** | Pointer to [**[]BgpConfigStatus**](BgpConfigStatus.md) | iBGP configuration. | [optional] 
**OspfConfig** | Pointer to [**OspfConfigStatus**](OspfConfigStatus.md) |  | [optional] 

## Methods

### NewInternalRoutingProtocolConfigStatus

`func NewInternalRoutingProtocolConfigStatus() *InternalRoutingProtocolConfigStatus`

NewInternalRoutingProtocolConfigStatus instantiates a new InternalRoutingProtocolConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalRoutingProtocolConfigStatusWithDefaults

`func NewInternalRoutingProtocolConfigStatusWithDefaults() *InternalRoutingProtocolConfigStatus`

NewInternalRoutingProtocolConfigStatusWithDefaults instantiates a new InternalRoutingProtocolConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPrefixList

`func (o *InternalRoutingProtocolConfigStatus) GetLocalPrefixList() []IpSubnetStatus`

GetLocalPrefixList returns the LocalPrefixList field if non-nil, zero value otherwise.

### GetLocalPrefixListOk

`func (o *InternalRoutingProtocolConfigStatus) GetLocalPrefixListOk() (*[]IpSubnetStatus, bool)`

GetLocalPrefixListOk returns a tuple with the LocalPrefixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPrefixList

`func (o *InternalRoutingProtocolConfigStatus) SetLocalPrefixList(v []IpSubnetStatus)`

SetLocalPrefixList sets LocalPrefixList field to given value.

### HasLocalPrefixList

`func (o *InternalRoutingProtocolConfigStatus) HasLocalPrefixList() bool`

HasLocalPrefixList returns a boolean if a field has been set.

### GetIbgpConfigList

`func (o *InternalRoutingProtocolConfigStatus) GetIbgpConfigList() []BgpConfigStatus`

GetIbgpConfigList returns the IbgpConfigList field if non-nil, zero value otherwise.

### GetIbgpConfigListOk

`func (o *InternalRoutingProtocolConfigStatus) GetIbgpConfigListOk() (*[]BgpConfigStatus, bool)`

GetIbgpConfigListOk returns a tuple with the IbgpConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpConfigList

`func (o *InternalRoutingProtocolConfigStatus) SetIbgpConfigList(v []BgpConfigStatus)`

SetIbgpConfigList sets IbgpConfigList field to given value.

### HasIbgpConfigList

`func (o *InternalRoutingProtocolConfigStatus) HasIbgpConfigList() bool`

HasIbgpConfigList returns a boolean if a field has been set.

### GetOspfConfig

`func (o *InternalRoutingProtocolConfigStatus) GetOspfConfig() OspfConfigStatus`

GetOspfConfig returns the OspfConfig field if non-nil, zero value otherwise.

### GetOspfConfigOk

`func (o *InternalRoutingProtocolConfigStatus) GetOspfConfigOk() (*OspfConfigStatus, bool)`

GetOspfConfigOk returns a tuple with the OspfConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfConfig

`func (o *InternalRoutingProtocolConfigStatus) SetOspfConfig(v OspfConfigStatus)`

SetOspfConfig sets OspfConfig field to given value.

### HasOspfConfig

`func (o *InternalRoutingProtocolConfigStatus) HasOspfConfig() bool`

HasOspfConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


