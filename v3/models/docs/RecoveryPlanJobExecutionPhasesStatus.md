# RecoveryPlanJobExecutionPhasesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | State of execution of the Recovery Plan Job. | 
**PostprocessingStatus** | Pointer to [**RecoveryPlanJobStepStatus**](RecoveryPlanJobStepStatus.md) |  | [optional] 
**OperationStatus** | Pointer to [**RecoveryPlanJobStepStatus**](RecoveryPlanJobStepStatus.md) |  | [optional] 
**PercentageComplete** | **int32** | Percentage completed for Recovery Plan Job. | 
**PreprocessingStatus** | Pointer to [**RecoveryPlanJobStepStatus**](RecoveryPlanJobStepStatus.md) |  | [optional] 

## Methods

### NewRecoveryPlanJobExecutionPhasesStatus

`func NewRecoveryPlanJobExecutionPhasesStatus(status string, percentageComplete int32, ) *RecoveryPlanJobExecutionPhasesStatus`

NewRecoveryPlanJobExecutionPhasesStatus instantiates a new RecoveryPlanJobExecutionPhasesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobExecutionPhasesStatusWithDefaults

`func NewRecoveryPlanJobExecutionPhasesStatusWithDefaults() *RecoveryPlanJobExecutionPhasesStatus`

NewRecoveryPlanJobExecutionPhasesStatusWithDefaults instantiates a new RecoveryPlanJobExecutionPhasesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPostprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPostprocessingStatus() RecoveryPlanJobStepStatus`

GetPostprocessingStatus returns the PostprocessingStatus field if non-nil, zero value otherwise.

### GetPostprocessingStatusOk

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPostprocessingStatusOk() (*RecoveryPlanJobStepStatus, bool)`

GetPostprocessingStatusOk returns a tuple with the PostprocessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) SetPostprocessingStatus(v RecoveryPlanJobStepStatus)`

SetPostprocessingStatus sets PostprocessingStatus field to given value.

### HasPostprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) HasPostprocessingStatus() bool`

HasPostprocessingStatus returns a boolean if a field has been set.

### GetOperationStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetOperationStatus() RecoveryPlanJobStepStatus`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetOperationStatusOk() (*RecoveryPlanJobStepStatus, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) SetOperationStatus(v RecoveryPlanJobStepStatus)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *RecoveryPlanJobExecutionPhasesStatus) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.


### GetPreprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPreprocessingStatus() RecoveryPlanJobStepStatus`

GetPreprocessingStatus returns the PreprocessingStatus field if non-nil, zero value otherwise.

### GetPreprocessingStatusOk

`func (o *RecoveryPlanJobExecutionPhasesStatus) GetPreprocessingStatusOk() (*RecoveryPlanJobStepStatus, bool)`

GetPreprocessingStatusOk returns a tuple with the PreprocessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) SetPreprocessingStatus(v RecoveryPlanJobStepStatus)`

SetPreprocessingStatus sets PreprocessingStatus field to given value.

### HasPreprocessingStatus

`func (o *RecoveryPlanJobExecutionPhasesStatus) HasPreprocessingStatus() bool`

HasPreprocessingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


