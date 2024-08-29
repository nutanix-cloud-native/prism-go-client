# RecoveryPlanScriptConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableScriptExec** | **bool** | Indicates whether to execute script.  | 
**Timeout** | Pointer to **int32** | The timeout for the script (seconds).  | [optional] 

## Methods

### NewRecoveryPlanScriptConfig

`func NewRecoveryPlanScriptConfig(enableScriptExec bool, ) *RecoveryPlanScriptConfig`

NewRecoveryPlanScriptConfig instantiates a new RecoveryPlanScriptConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanScriptConfigWithDefaults

`func NewRecoveryPlanScriptConfigWithDefaults() *RecoveryPlanScriptConfig`

NewRecoveryPlanScriptConfigWithDefaults instantiates a new RecoveryPlanScriptConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableScriptExec

`func (o *RecoveryPlanScriptConfig) GetEnableScriptExec() bool`

GetEnableScriptExec returns the EnableScriptExec field if non-nil, zero value otherwise.

### GetEnableScriptExecOk

`func (o *RecoveryPlanScriptConfig) GetEnableScriptExecOk() (*bool, bool)`

GetEnableScriptExecOk returns a tuple with the EnableScriptExec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScriptExec

`func (o *RecoveryPlanScriptConfig) SetEnableScriptExec(v bool)`

SetEnableScriptExec sets EnableScriptExec field to given value.


### GetTimeout

`func (o *RecoveryPlanScriptConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RecoveryPlanScriptConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RecoveryPlanScriptConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RecoveryPlanScriptConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


