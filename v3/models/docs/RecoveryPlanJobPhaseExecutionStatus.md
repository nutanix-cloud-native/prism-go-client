# RecoveryPlanJobPhaseExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Execution state of a phase of Recovery Plan Job execution. | [optional] 
**PercentageComplete** | Pointer to **int32** | Percentage completed for a phase of the Recovery Plan Job execution.  | [optional] 
**StepExecutionStatusList** | Pointer to [**[]RecoveryPlanJobStepExecutionStatus**](RecoveryPlanJobStepExecutionStatus.md) | List of execution status of steps for a phase of the Recovery Plan Job execution.  | [optional] 

## Methods

### NewRecoveryPlanJobPhaseExecutionStatus

`func NewRecoveryPlanJobPhaseExecutionStatus() *RecoveryPlanJobPhaseExecutionStatus`

NewRecoveryPlanJobPhaseExecutionStatus instantiates a new RecoveryPlanJobPhaseExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobPhaseExecutionStatusWithDefaults

`func NewRecoveryPlanJobPhaseExecutionStatusWithDefaults() *RecoveryPlanJobPhaseExecutionStatus`

NewRecoveryPlanJobPhaseExecutionStatusWithDefaults instantiates a new RecoveryPlanJobPhaseExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobPhaseExecutionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecoveryPlanJobPhaseExecutionStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *RecoveryPlanJobPhaseExecutionStatus) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.

### HasPercentageComplete

`func (o *RecoveryPlanJobPhaseExecutionStatus) HasPercentageComplete() bool`

HasPercentageComplete returns a boolean if a field has been set.

### GetStepExecutionStatusList

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetStepExecutionStatusList() []RecoveryPlanJobStepExecutionStatus`

GetStepExecutionStatusList returns the StepExecutionStatusList field if non-nil, zero value otherwise.

### GetStepExecutionStatusListOk

`func (o *RecoveryPlanJobPhaseExecutionStatus) GetStepExecutionStatusListOk() (*[]RecoveryPlanJobStepExecutionStatus, bool)`

GetStepExecutionStatusListOk returns a tuple with the StepExecutionStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutionStatusList

`func (o *RecoveryPlanJobPhaseExecutionStatus) SetStepExecutionStatusList(v []RecoveryPlanJobStepExecutionStatus)`

SetStepExecutionStatusList sets StepExecutionStatusList field to given value.

### HasStepExecutionStatusList

`func (o *RecoveryPlanJobPhaseExecutionStatus) HasStepExecutionStatusList() bool`

HasStepExecutionStatusList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


