# PcVmNicConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkConfiguration** | Pointer to [**NetworkConfig**](NetworkConfig.md) |  | [optional] 
**IpList** | Pointer to **[]string** | Network IP address. | [optional] 

## Methods

### NewPcVmNicConfiguration

`func NewPcVmNicConfiguration() *PcVmNicConfiguration`

NewPcVmNicConfiguration instantiates a new PcVmNicConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcVmNicConfigurationWithDefaults

`func NewPcVmNicConfigurationWithDefaults() *PcVmNicConfiguration`

NewPcVmNicConfigurationWithDefaults instantiates a new PcVmNicConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkConfiguration

`func (o *PcVmNicConfiguration) GetNetworkConfiguration() NetworkConfig`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *PcVmNicConfiguration) GetNetworkConfigurationOk() (*NetworkConfig, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *PcVmNicConfiguration) SetNetworkConfiguration(v NetworkConfig)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *PcVmNicConfiguration) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### GetIpList

`func (o *PcVmNicConfiguration) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *PcVmNicConfiguration) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *PcVmNicConfiguration) SetIpList(v []string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *PcVmNicConfiguration) HasIpList() bool`

HasIpList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


