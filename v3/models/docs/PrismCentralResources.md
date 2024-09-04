# PrismCentralResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldAutoRegister** | Pointer to **bool** | Indicates if the new prism central should be automatically register to the cluster.  | [optional] [default to false]
**InitialPassword** | Pointer to **string** | Initial boot up password for new created Prism Central. | [optional] 
**IsRegisteredToHostingPe** | Pointer to **bool** | If this prism central vm is registered to hosting PE. | [optional] [default to false]
**CmspConfig** | Pointer to [**CmspConfig**](CmspConfig.md) |  | [optional] 
**McmConfig** | Pointer to [**McmConfig**](McmConfig.md) |  | [optional] 
**PcVmList** | [**[]PcVm**](PcVm.md) |  | 
**Environment** | Pointer to **string** | The MCM environment config. | [optional] 
**Version** | **string** | The desired version of Prism Central. | 
**VirtualIp** | Pointer to **string** | The desired virtual IP address of Prism Central cluster.  | [optional] 
**ClusterUuid** | Pointer to **string** | The Cluster UUID for the Prism Central Deployment. | [optional] 
**Type** | Pointer to **string** | The type of the Prism Central cluster. | [optional] [default to "PC"]

## Methods

### NewPrismCentralResources

`func NewPrismCentralResources(pcVmList []PcVm, version string, ) *PrismCentralResources`

NewPrismCentralResources instantiates a new PrismCentralResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismCentralResourcesWithDefaults

`func NewPrismCentralResourcesWithDefaults() *PrismCentralResources`

NewPrismCentralResourcesWithDefaults instantiates a new PrismCentralResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldAutoRegister

`func (o *PrismCentralResources) GetShouldAutoRegister() bool`

GetShouldAutoRegister returns the ShouldAutoRegister field if non-nil, zero value otherwise.

### GetShouldAutoRegisterOk

`func (o *PrismCentralResources) GetShouldAutoRegisterOk() (*bool, bool)`

GetShouldAutoRegisterOk returns a tuple with the ShouldAutoRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAutoRegister

`func (o *PrismCentralResources) SetShouldAutoRegister(v bool)`

SetShouldAutoRegister sets ShouldAutoRegister field to given value.

### HasShouldAutoRegister

`func (o *PrismCentralResources) HasShouldAutoRegister() bool`

HasShouldAutoRegister returns a boolean if a field has been set.

### GetInitialPassword

`func (o *PrismCentralResources) GetInitialPassword() string`

GetInitialPassword returns the InitialPassword field if non-nil, zero value otherwise.

### GetInitialPasswordOk

`func (o *PrismCentralResources) GetInitialPasswordOk() (*string, bool)`

GetInitialPasswordOk returns a tuple with the InitialPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPassword

`func (o *PrismCentralResources) SetInitialPassword(v string)`

SetInitialPassword sets InitialPassword field to given value.

### HasInitialPassword

`func (o *PrismCentralResources) HasInitialPassword() bool`

HasInitialPassword returns a boolean if a field has been set.

### GetIsRegisteredToHostingPe

`func (o *PrismCentralResources) GetIsRegisteredToHostingPe() bool`

GetIsRegisteredToHostingPe returns the IsRegisteredToHostingPe field if non-nil, zero value otherwise.

### GetIsRegisteredToHostingPeOk

`func (o *PrismCentralResources) GetIsRegisteredToHostingPeOk() (*bool, bool)`

GetIsRegisteredToHostingPeOk returns a tuple with the IsRegisteredToHostingPe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisteredToHostingPe

`func (o *PrismCentralResources) SetIsRegisteredToHostingPe(v bool)`

SetIsRegisteredToHostingPe sets IsRegisteredToHostingPe field to given value.

### HasIsRegisteredToHostingPe

`func (o *PrismCentralResources) HasIsRegisteredToHostingPe() bool`

HasIsRegisteredToHostingPe returns a boolean if a field has been set.

### GetCmspConfig

`func (o *PrismCentralResources) GetCmspConfig() CmspConfig`

GetCmspConfig returns the CmspConfig field if non-nil, zero value otherwise.

### GetCmspConfigOk

`func (o *PrismCentralResources) GetCmspConfigOk() (*CmspConfig, bool)`

GetCmspConfigOk returns a tuple with the CmspConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmspConfig

`func (o *PrismCentralResources) SetCmspConfig(v CmspConfig)`

SetCmspConfig sets CmspConfig field to given value.

### HasCmspConfig

`func (o *PrismCentralResources) HasCmspConfig() bool`

HasCmspConfig returns a boolean if a field has been set.

### GetMcmConfig

`func (o *PrismCentralResources) GetMcmConfig() McmConfig`

GetMcmConfig returns the McmConfig field if non-nil, zero value otherwise.

### GetMcmConfigOk

`func (o *PrismCentralResources) GetMcmConfigOk() (*McmConfig, bool)`

GetMcmConfigOk returns a tuple with the McmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcmConfig

`func (o *PrismCentralResources) SetMcmConfig(v McmConfig)`

SetMcmConfig sets McmConfig field to given value.

### HasMcmConfig

`func (o *PrismCentralResources) HasMcmConfig() bool`

HasMcmConfig returns a boolean if a field has been set.

### GetPcVmList

`func (o *PrismCentralResources) GetPcVmList() []PcVm`

GetPcVmList returns the PcVmList field if non-nil, zero value otherwise.

### GetPcVmListOk

`func (o *PrismCentralResources) GetPcVmListOk() (*[]PcVm, bool)`

GetPcVmListOk returns a tuple with the PcVmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcVmList

`func (o *PrismCentralResources) SetPcVmList(v []PcVm)`

SetPcVmList sets PcVmList field to given value.


### GetEnvironment

`func (o *PrismCentralResources) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PrismCentralResources) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PrismCentralResources) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PrismCentralResources) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetVersion

`func (o *PrismCentralResources) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PrismCentralResources) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PrismCentralResources) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVirtualIp

`func (o *PrismCentralResources) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *PrismCentralResources) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *PrismCentralResources) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *PrismCentralResources) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetClusterUuid

`func (o *PrismCentralResources) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *PrismCentralResources) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *PrismCentralResources) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *PrismCentralResources) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetType

`func (o *PrismCentralResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrismCentralResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrismCentralResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrismCentralResources) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


