# VmHostAffinityPolicyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostsCount** | Pointer to **int64** | Total number of applicable Hosts. | [optional] 
**VmsCount** | Pointer to **int64** | Total number of applicable VMs. | [optional] 
**Config** | Pointer to [**VmHostAffinityPolicyConfig**](VmHostAffinityPolicyConfig.md) |  | [optional] 
**NonCompliantVmsCount** | Pointer to **int64** | Total number of VMs which are non-compliant with the policy. | [optional] 
**CompliantVmsCount** | Pointer to **int64** | Total number of VMs which are compliant with the policy. | [optional] 

## Methods

### NewVmHostAffinityPolicyInfo

`func NewVmHostAffinityPolicyInfo() *VmHostAffinityPolicyInfo`

NewVmHostAffinityPolicyInfo instantiates a new VmHostAffinityPolicyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyInfoWithDefaults

`func NewVmHostAffinityPolicyInfoWithDefaults() *VmHostAffinityPolicyInfo`

NewVmHostAffinityPolicyInfoWithDefaults instantiates a new VmHostAffinityPolicyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostsCount

`func (o *VmHostAffinityPolicyInfo) GetHostsCount() int64`

GetHostsCount returns the HostsCount field if non-nil, zero value otherwise.

### GetHostsCountOk

`func (o *VmHostAffinityPolicyInfo) GetHostsCountOk() (*int64, bool)`

GetHostsCountOk returns a tuple with the HostsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsCount

`func (o *VmHostAffinityPolicyInfo) SetHostsCount(v int64)`

SetHostsCount sets HostsCount field to given value.

### HasHostsCount

`func (o *VmHostAffinityPolicyInfo) HasHostsCount() bool`

HasHostsCount returns a boolean if a field has been set.

### GetVmsCount

`func (o *VmHostAffinityPolicyInfo) GetVmsCount() int64`

GetVmsCount returns the VmsCount field if non-nil, zero value otherwise.

### GetVmsCountOk

`func (o *VmHostAffinityPolicyInfo) GetVmsCountOk() (*int64, bool)`

GetVmsCountOk returns a tuple with the VmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmsCount

`func (o *VmHostAffinityPolicyInfo) SetVmsCount(v int64)`

SetVmsCount sets VmsCount field to given value.

### HasVmsCount

`func (o *VmHostAffinityPolicyInfo) HasVmsCount() bool`

HasVmsCount returns a boolean if a field has been set.

### GetConfig

`func (o *VmHostAffinityPolicyInfo) GetConfig() VmHostAffinityPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VmHostAffinityPolicyInfo) GetConfigOk() (*VmHostAffinityPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VmHostAffinityPolicyInfo) SetConfig(v VmHostAffinityPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VmHostAffinityPolicyInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNonCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) GetNonCompliantVmsCount() int64`

GetNonCompliantVmsCount returns the NonCompliantVmsCount field if non-nil, zero value otherwise.

### GetNonCompliantVmsCountOk

`func (o *VmHostAffinityPolicyInfo) GetNonCompliantVmsCountOk() (*int64, bool)`

GetNonCompliantVmsCountOk returns a tuple with the NonCompliantVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) SetNonCompliantVmsCount(v int64)`

SetNonCompliantVmsCount sets NonCompliantVmsCount field to given value.

### HasNonCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) HasNonCompliantVmsCount() bool`

HasNonCompliantVmsCount returns a boolean if a field has been set.

### GetCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) GetCompliantVmsCount() int64`

GetCompliantVmsCount returns the CompliantVmsCount field if non-nil, zero value otherwise.

### GetCompliantVmsCountOk

`func (o *VmHostAffinityPolicyInfo) GetCompliantVmsCountOk() (*int64, bool)`

GetCompliantVmsCountOk returns a tuple with the CompliantVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) SetCompliantVmsCount(v int64)`

SetCompliantVmsCount sets CompliantVmsCount field to given value.

### HasCompliantVmsCount

`func (o *VmHostAffinityPolicyInfo) HasCompliantVmsCount() bool`

HasCompliantVmsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


