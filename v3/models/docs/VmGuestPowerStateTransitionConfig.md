# VmGuestPowerStateTransitionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldFailOnScriptFailure** | Pointer to **bool** | Indicates whether to abort ngt shutdown/reboot if script fails. | [optional] 
**EnableScriptExec** | Pointer to **bool** | Indicates whether to execute set script before ngt shutdown/reboot. | [optional] 

## Methods

### NewVmGuestPowerStateTransitionConfig

`func NewVmGuestPowerStateTransitionConfig() *VmGuestPowerStateTransitionConfig`

NewVmGuestPowerStateTransitionConfig instantiates a new VmGuestPowerStateTransitionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGuestPowerStateTransitionConfigWithDefaults

`func NewVmGuestPowerStateTransitionConfigWithDefaults() *VmGuestPowerStateTransitionConfig`

NewVmGuestPowerStateTransitionConfigWithDefaults instantiates a new VmGuestPowerStateTransitionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldFailOnScriptFailure

`func (o *VmGuestPowerStateTransitionConfig) GetShouldFailOnScriptFailure() bool`

GetShouldFailOnScriptFailure returns the ShouldFailOnScriptFailure field if non-nil, zero value otherwise.

### GetShouldFailOnScriptFailureOk

`func (o *VmGuestPowerStateTransitionConfig) GetShouldFailOnScriptFailureOk() (*bool, bool)`

GetShouldFailOnScriptFailureOk returns a tuple with the ShouldFailOnScriptFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldFailOnScriptFailure

`func (o *VmGuestPowerStateTransitionConfig) SetShouldFailOnScriptFailure(v bool)`

SetShouldFailOnScriptFailure sets ShouldFailOnScriptFailure field to given value.

### HasShouldFailOnScriptFailure

`func (o *VmGuestPowerStateTransitionConfig) HasShouldFailOnScriptFailure() bool`

HasShouldFailOnScriptFailure returns a boolean if a field has been set.

### GetEnableScriptExec

`func (o *VmGuestPowerStateTransitionConfig) GetEnableScriptExec() bool`

GetEnableScriptExec returns the EnableScriptExec field if non-nil, zero value otherwise.

### GetEnableScriptExecOk

`func (o *VmGuestPowerStateTransitionConfig) GetEnableScriptExecOk() (*bool, bool)`

GetEnableScriptExecOk returns a tuple with the EnableScriptExec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScriptExec

`func (o *VmGuestPowerStateTransitionConfig) SetEnableScriptExec(v bool)`

SetEnableScriptExec sets EnableScriptExec field to given value.

### HasEnableScriptExec

`func (o *VmGuestPowerStateTransitionConfig) HasEnableScriptExec() bool`

HasEnableScriptExec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


