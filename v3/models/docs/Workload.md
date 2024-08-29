# Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedWorkload** | Pointer to [**AdjustedWorkload**](AdjustedWorkload.md) |  | [optional] 
**SqlWorkload** | Pointer to [**SqlWorkload**](SqlWorkload.md) |  | [optional] 
**WorkloadName** | Pointer to **string** | Workload name. | [optional] 
**ExchangeWorkload** | Pointer to [**ExchangeWorkload**](ExchangeWorkload.md) |  | [optional] 
**ToRemove** | Pointer to **bool** | The variable to indicated if the workload is used as removed workload. | [optional] [default to false]
**Enabled** | Pointer to **bool** | The variable to indicate if the workload is enabled. | [optional] [default to true]
**SplunkWorkload** | Pointer to [**SplunkWorkload**](SplunkWorkload.md) |  | [optional] 
**VdiWorkload** | Pointer to [**VdiWorkload**](VdiWorkload.md) |  | [optional] 
**VmWorkload** | Pointer to [**VmWorkload**](VmWorkload.md) |  | [optional] 
**WorkloadType** | Pointer to **string** | The type of workload. | [optional] 
**ResourceRequirement** | Pointer to [**GenericResourceSpec**](GenericResourceSpec.md) |  | [optional] 
**ScheduleTimestampSec** | Pointer to **int32** | The scheduled timestamp in seconds. | [optional] 
**XenWorkload** | Pointer to [**XenWorkload**](XenWorkload.md) |  | [optional] 
**VirtualServerWorkload** | Pointer to [**VirtualServerWorkload**](VirtualServerWorkload.md) |  | [optional] 
**VmCategoryWorkload** | Pointer to [**VmCategoryWorkload**](VmCategoryWorkload.md) |  | [optional] 

## Methods

### NewWorkload

`func NewWorkload() *Workload`

NewWorkload instantiates a new Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWithDefaults

`func NewWorkloadWithDefaults() *Workload`

NewWorkloadWithDefaults instantiates a new Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedWorkload

`func (o *Workload) GetAdjustedWorkload() AdjustedWorkload`

GetAdjustedWorkload returns the AdjustedWorkload field if non-nil, zero value otherwise.

### GetAdjustedWorkloadOk

`func (o *Workload) GetAdjustedWorkloadOk() (*AdjustedWorkload, bool)`

GetAdjustedWorkloadOk returns a tuple with the AdjustedWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedWorkload

`func (o *Workload) SetAdjustedWorkload(v AdjustedWorkload)`

SetAdjustedWorkload sets AdjustedWorkload field to given value.

### HasAdjustedWorkload

`func (o *Workload) HasAdjustedWorkload() bool`

HasAdjustedWorkload returns a boolean if a field has been set.

### GetSqlWorkload

`func (o *Workload) GetSqlWorkload() SqlWorkload`

GetSqlWorkload returns the SqlWorkload field if non-nil, zero value otherwise.

### GetSqlWorkloadOk

`func (o *Workload) GetSqlWorkloadOk() (*SqlWorkload, bool)`

GetSqlWorkloadOk returns a tuple with the SqlWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWorkload

`func (o *Workload) SetSqlWorkload(v SqlWorkload)`

SetSqlWorkload sets SqlWorkload field to given value.

### HasSqlWorkload

`func (o *Workload) HasSqlWorkload() bool`

HasSqlWorkload returns a boolean if a field has been set.

### GetWorkloadName

`func (o *Workload) GetWorkloadName() string`

GetWorkloadName returns the WorkloadName field if non-nil, zero value otherwise.

### GetWorkloadNameOk

`func (o *Workload) GetWorkloadNameOk() (*string, bool)`

GetWorkloadNameOk returns a tuple with the WorkloadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadName

`func (o *Workload) SetWorkloadName(v string)`

SetWorkloadName sets WorkloadName field to given value.

### HasWorkloadName

`func (o *Workload) HasWorkloadName() bool`

HasWorkloadName returns a boolean if a field has been set.

### GetExchangeWorkload

`func (o *Workload) GetExchangeWorkload() ExchangeWorkload`

GetExchangeWorkload returns the ExchangeWorkload field if non-nil, zero value otherwise.

### GetExchangeWorkloadOk

`func (o *Workload) GetExchangeWorkloadOk() (*ExchangeWorkload, bool)`

GetExchangeWorkloadOk returns a tuple with the ExchangeWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeWorkload

`func (o *Workload) SetExchangeWorkload(v ExchangeWorkload)`

SetExchangeWorkload sets ExchangeWorkload field to given value.

### HasExchangeWorkload

`func (o *Workload) HasExchangeWorkload() bool`

HasExchangeWorkload returns a boolean if a field has been set.

### GetToRemove

`func (o *Workload) GetToRemove() bool`

GetToRemove returns the ToRemove field if non-nil, zero value otherwise.

### GetToRemoveOk

`func (o *Workload) GetToRemoveOk() (*bool, bool)`

GetToRemoveOk returns a tuple with the ToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRemove

`func (o *Workload) SetToRemove(v bool)`

SetToRemove sets ToRemove field to given value.

### HasToRemove

`func (o *Workload) HasToRemove() bool`

HasToRemove returns a boolean if a field has been set.

### GetEnabled

`func (o *Workload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Workload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Workload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Workload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSplunkWorkload

`func (o *Workload) GetSplunkWorkload() SplunkWorkload`

GetSplunkWorkload returns the SplunkWorkload field if non-nil, zero value otherwise.

### GetSplunkWorkloadOk

`func (o *Workload) GetSplunkWorkloadOk() (*SplunkWorkload, bool)`

GetSplunkWorkloadOk returns a tuple with the SplunkWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkWorkload

`func (o *Workload) SetSplunkWorkload(v SplunkWorkload)`

SetSplunkWorkload sets SplunkWorkload field to given value.

### HasSplunkWorkload

`func (o *Workload) HasSplunkWorkload() bool`

HasSplunkWorkload returns a boolean if a field has been set.

### GetVdiWorkload

`func (o *Workload) GetVdiWorkload() VdiWorkload`

GetVdiWorkload returns the VdiWorkload field if non-nil, zero value otherwise.

### GetVdiWorkloadOk

`func (o *Workload) GetVdiWorkloadOk() (*VdiWorkload, bool)`

GetVdiWorkloadOk returns a tuple with the VdiWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiWorkload

`func (o *Workload) SetVdiWorkload(v VdiWorkload)`

SetVdiWorkload sets VdiWorkload field to given value.

### HasVdiWorkload

`func (o *Workload) HasVdiWorkload() bool`

HasVdiWorkload returns a boolean if a field has been set.

### GetVmWorkload

`func (o *Workload) GetVmWorkload() VmWorkload`

GetVmWorkload returns the VmWorkload field if non-nil, zero value otherwise.

### GetVmWorkloadOk

`func (o *Workload) GetVmWorkloadOk() (*VmWorkload, bool)`

GetVmWorkloadOk returns a tuple with the VmWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmWorkload

`func (o *Workload) SetVmWorkload(v VmWorkload)`

SetVmWorkload sets VmWorkload field to given value.

### HasVmWorkload

`func (o *Workload) HasVmWorkload() bool`

HasVmWorkload returns a boolean if a field has been set.

### GetWorkloadType

`func (o *Workload) GetWorkloadType() string`

GetWorkloadType returns the WorkloadType field if non-nil, zero value otherwise.

### GetWorkloadTypeOk

`func (o *Workload) GetWorkloadTypeOk() (*string, bool)`

GetWorkloadTypeOk returns a tuple with the WorkloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadType

`func (o *Workload) SetWorkloadType(v string)`

SetWorkloadType sets WorkloadType field to given value.

### HasWorkloadType

`func (o *Workload) HasWorkloadType() bool`

HasWorkloadType returns a boolean if a field has been set.

### GetResourceRequirement

`func (o *Workload) GetResourceRequirement() GenericResourceSpec`

GetResourceRequirement returns the ResourceRequirement field if non-nil, zero value otherwise.

### GetResourceRequirementOk

`func (o *Workload) GetResourceRequirementOk() (*GenericResourceSpec, bool)`

GetResourceRequirementOk returns a tuple with the ResourceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRequirement

`func (o *Workload) SetResourceRequirement(v GenericResourceSpec)`

SetResourceRequirement sets ResourceRequirement field to given value.

### HasResourceRequirement

`func (o *Workload) HasResourceRequirement() bool`

HasResourceRequirement returns a boolean if a field has been set.

### GetScheduleTimestampSec

`func (o *Workload) GetScheduleTimestampSec() int32`

GetScheduleTimestampSec returns the ScheduleTimestampSec field if non-nil, zero value otherwise.

### GetScheduleTimestampSecOk

`func (o *Workload) GetScheduleTimestampSecOk() (*int32, bool)`

GetScheduleTimestampSecOk returns a tuple with the ScheduleTimestampSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimestampSec

`func (o *Workload) SetScheduleTimestampSec(v int32)`

SetScheduleTimestampSec sets ScheduleTimestampSec field to given value.

### HasScheduleTimestampSec

`func (o *Workload) HasScheduleTimestampSec() bool`

HasScheduleTimestampSec returns a boolean if a field has been set.

### GetXenWorkload

`func (o *Workload) GetXenWorkload() XenWorkload`

GetXenWorkload returns the XenWorkload field if non-nil, zero value otherwise.

### GetXenWorkloadOk

`func (o *Workload) GetXenWorkloadOk() (*XenWorkload, bool)`

GetXenWorkloadOk returns a tuple with the XenWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXenWorkload

`func (o *Workload) SetXenWorkload(v XenWorkload)`

SetXenWorkload sets XenWorkload field to given value.

### HasXenWorkload

`func (o *Workload) HasXenWorkload() bool`

HasXenWorkload returns a boolean if a field has been set.

### GetVirtualServerWorkload

`func (o *Workload) GetVirtualServerWorkload() VirtualServerWorkload`

GetVirtualServerWorkload returns the VirtualServerWorkload field if non-nil, zero value otherwise.

### GetVirtualServerWorkloadOk

`func (o *Workload) GetVirtualServerWorkloadOk() (*VirtualServerWorkload, bool)`

GetVirtualServerWorkloadOk returns a tuple with the VirtualServerWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServerWorkload

`func (o *Workload) SetVirtualServerWorkload(v VirtualServerWorkload)`

SetVirtualServerWorkload sets VirtualServerWorkload field to given value.

### HasVirtualServerWorkload

`func (o *Workload) HasVirtualServerWorkload() bool`

HasVirtualServerWorkload returns a boolean if a field has been set.

### GetVmCategoryWorkload

`func (o *Workload) GetVmCategoryWorkload() VmCategoryWorkload`

GetVmCategoryWorkload returns the VmCategoryWorkload field if non-nil, zero value otherwise.

### GetVmCategoryWorkloadOk

`func (o *Workload) GetVmCategoryWorkloadOk() (*VmCategoryWorkload, bool)`

GetVmCategoryWorkloadOk returns a tuple with the VmCategoryWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCategoryWorkload

`func (o *Workload) SetVmCategoryWorkload(v VmCategoryWorkload)`

SetVmCategoryWorkload sets VmCategoryWorkload field to given value.

### HasVmCategoryWorkload

`func (o *Workload) HasVmCategoryWorkload() bool`

HasVmCategoryWorkload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


