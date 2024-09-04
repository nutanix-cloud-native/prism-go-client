# ConfigurationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalMemoryBytes** | Pointer to **float64** | Original memory (in bytes) | [optional] 
**AdditionalVcpu** | Pointer to **int64** | Additional vcpu required to enable the service | [optional] 
**AdditionalMemory** | Pointer to **int64** | Additional memory (in GiB) required to enable the service | [optional] 
**AdditionalMemoryBytes** | Pointer to **float64** | Additional memory (in bytes) required to enable the service | [optional] 

## Methods

### NewConfigurationInfo

`func NewConfigurationInfo() *ConfigurationInfo`

NewConfigurationInfo instantiates a new ConfigurationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationInfoWithDefaults

`func NewConfigurationInfoWithDefaults() *ConfigurationInfo`

NewConfigurationInfoWithDefaults instantiates a new ConfigurationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalMemoryBytes

`func (o *ConfigurationInfo) GetOriginalMemoryBytes() float64`

GetOriginalMemoryBytes returns the OriginalMemoryBytes field if non-nil, zero value otherwise.

### GetOriginalMemoryBytesOk

`func (o *ConfigurationInfo) GetOriginalMemoryBytesOk() (*float64, bool)`

GetOriginalMemoryBytesOk returns a tuple with the OriginalMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMemoryBytes

`func (o *ConfigurationInfo) SetOriginalMemoryBytes(v float64)`

SetOriginalMemoryBytes sets OriginalMemoryBytes field to given value.

### HasOriginalMemoryBytes

`func (o *ConfigurationInfo) HasOriginalMemoryBytes() bool`

HasOriginalMemoryBytes returns a boolean if a field has been set.

### GetAdditionalVcpu

`func (o *ConfigurationInfo) GetAdditionalVcpu() int64`

GetAdditionalVcpu returns the AdditionalVcpu field if non-nil, zero value otherwise.

### GetAdditionalVcpuOk

`func (o *ConfigurationInfo) GetAdditionalVcpuOk() (*int64, bool)`

GetAdditionalVcpuOk returns a tuple with the AdditionalVcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalVcpu

`func (o *ConfigurationInfo) SetAdditionalVcpu(v int64)`

SetAdditionalVcpu sets AdditionalVcpu field to given value.

### HasAdditionalVcpu

`func (o *ConfigurationInfo) HasAdditionalVcpu() bool`

HasAdditionalVcpu returns a boolean if a field has been set.

### GetAdditionalMemory

`func (o *ConfigurationInfo) GetAdditionalMemory() int64`

GetAdditionalMemory returns the AdditionalMemory field if non-nil, zero value otherwise.

### GetAdditionalMemoryOk

`func (o *ConfigurationInfo) GetAdditionalMemoryOk() (*int64, bool)`

GetAdditionalMemoryOk returns a tuple with the AdditionalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMemory

`func (o *ConfigurationInfo) SetAdditionalMemory(v int64)`

SetAdditionalMemory sets AdditionalMemory field to given value.

### HasAdditionalMemory

`func (o *ConfigurationInfo) HasAdditionalMemory() bool`

HasAdditionalMemory returns a boolean if a field has been set.

### GetAdditionalMemoryBytes

`func (o *ConfigurationInfo) GetAdditionalMemoryBytes() float64`

GetAdditionalMemoryBytes returns the AdditionalMemoryBytes field if non-nil, zero value otherwise.

### GetAdditionalMemoryBytesOk

`func (o *ConfigurationInfo) GetAdditionalMemoryBytesOk() (*float64, bool)`

GetAdditionalMemoryBytesOk returns a tuple with the AdditionalMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMemoryBytes

`func (o *ConfigurationInfo) SetAdditionalMemoryBytes(v float64)`

SetAdditionalMemoryBytes sets AdditionalMemoryBytes field to given value.

### HasAdditionalMemoryBytes

`func (o *ConfigurationInfo) HasAdditionalMemoryBytes() bool`

HasAdditionalMemoryBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


