# CmspConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformIpBlockList** | **[]string** | IP range blocks for CMSP. | 
**PlatformNetworkConfiguration** | [**CmspNetworkConfig**](CmspNetworkConfig.md) |  | 
**CmspArgs** | Pointer to **string** | A serialized json containing additional arguments to be passed to CMSP. | [optional] 
**PcDomainName** | **string** | The domain name for CMSP. | 

## Methods

### NewCmspConfig

`func NewCmspConfig(platformIpBlockList []string, platformNetworkConfiguration CmspNetworkConfig, pcDomainName string, ) *CmspConfig`

NewCmspConfig instantiates a new CmspConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmspConfigWithDefaults

`func NewCmspConfigWithDefaults() *CmspConfig`

NewCmspConfigWithDefaults instantiates a new CmspConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformIpBlockList

`func (o *CmspConfig) GetPlatformIpBlockList() []string`

GetPlatformIpBlockList returns the PlatformIpBlockList field if non-nil, zero value otherwise.

### GetPlatformIpBlockListOk

`func (o *CmspConfig) GetPlatformIpBlockListOk() (*[]string, bool)`

GetPlatformIpBlockListOk returns a tuple with the PlatformIpBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIpBlockList

`func (o *CmspConfig) SetPlatformIpBlockList(v []string)`

SetPlatformIpBlockList sets PlatformIpBlockList field to given value.


### GetPlatformNetworkConfiguration

`func (o *CmspConfig) GetPlatformNetworkConfiguration() CmspNetworkConfig`

GetPlatformNetworkConfiguration returns the PlatformNetworkConfiguration field if non-nil, zero value otherwise.

### GetPlatformNetworkConfigurationOk

`func (o *CmspConfig) GetPlatformNetworkConfigurationOk() (*CmspNetworkConfig, bool)`

GetPlatformNetworkConfigurationOk returns a tuple with the PlatformNetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformNetworkConfiguration

`func (o *CmspConfig) SetPlatformNetworkConfiguration(v CmspNetworkConfig)`

SetPlatformNetworkConfiguration sets PlatformNetworkConfiguration field to given value.


### GetCmspArgs

`func (o *CmspConfig) GetCmspArgs() string`

GetCmspArgs returns the CmspArgs field if non-nil, zero value otherwise.

### GetCmspArgsOk

`func (o *CmspConfig) GetCmspArgsOk() (*string, bool)`

GetCmspArgsOk returns a tuple with the CmspArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmspArgs

`func (o *CmspConfig) SetCmspArgs(v string)`

SetCmspArgs sets CmspArgs field to given value.

### HasCmspArgs

`func (o *CmspConfig) HasCmspArgs() bool`

HasCmspArgs returns a boolean if a field has been set.

### GetPcDomainName

`func (o *CmspConfig) GetPcDomainName() string`

GetPcDomainName returns the PcDomainName field if non-nil, zero value otherwise.

### GetPcDomainNameOk

`func (o *CmspConfig) GetPcDomainNameOk() (*string, bool)`

GetPcDomainNameOk returns a tuple with the PcDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcDomainName

`func (o *CmspConfig) SetPcDomainName(v string)`

SetPcDomainName sets PcDomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


