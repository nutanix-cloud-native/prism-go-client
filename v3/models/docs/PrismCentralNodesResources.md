# PrismCentralNodesResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmspConfig** | Pointer to [**CmspExpansionConfig**](CmspExpansionConfig.md) |  | [optional] 
**PcVmList** | [**[]PcVm**](PcVm.md) |  | 

## Methods

### NewPrismCentralNodesResources

`func NewPrismCentralNodesResources(pcVmList []PcVm, ) *PrismCentralNodesResources`

NewPrismCentralNodesResources instantiates a new PrismCentralNodesResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismCentralNodesResourcesWithDefaults

`func NewPrismCentralNodesResourcesWithDefaults() *PrismCentralNodesResources`

NewPrismCentralNodesResourcesWithDefaults instantiates a new PrismCentralNodesResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmspConfig

`func (o *PrismCentralNodesResources) GetCmspConfig() CmspExpansionConfig`

GetCmspConfig returns the CmspConfig field if non-nil, zero value otherwise.

### GetCmspConfigOk

`func (o *PrismCentralNodesResources) GetCmspConfigOk() (*CmspExpansionConfig, bool)`

GetCmspConfigOk returns a tuple with the CmspConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmspConfig

`func (o *PrismCentralNodesResources) SetCmspConfig(v CmspExpansionConfig)`

SetCmspConfig sets CmspConfig field to given value.

### HasCmspConfig

`func (o *PrismCentralNodesResources) HasCmspConfig() bool`

HasCmspConfig returns a boolean if a field has been set.

### GetPcVmList

`func (o *PrismCentralNodesResources) GetPcVmList() []PcVm`

GetPcVmList returns the PcVmList field if non-nil, zero value otherwise.

### GetPcVmListOk

`func (o *PrismCentralNodesResources) GetPcVmListOk() (*[]PcVm, bool)`

GetPcVmListOk returns a tuple with the PcVmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcVmList

`func (o *PrismCentralNodesResources) SetPcVmList(v []PcVm)`

SetPcVmList sets PcVmList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


