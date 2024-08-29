# RecoveryPlanJobExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostprocessingStatus** | Pointer to [**RecoveryPlanJobPhaseExecutionStatus**](RecoveryPlanJobPhaseExecutionStatus.md) |  | [optional] 
**OperationStatus** | Pointer to [**RecoveryPlanJobPhaseExecutionStatus**](RecoveryPlanJobPhaseExecutionStatus.md) |  | [optional] 
**PreprocessingStatus** | Pointer to [**RecoveryPlanJobPhaseExecutionStatus**](RecoveryPlanJobPhaseExecutionStatus.md) |  | [optional] 

## Methods

### NewRecoveryPlanJobExecutionStatus

`func NewRecoveryPlanJobExecutionStatus() *RecoveryPlanJobExecutionStatus`

NewRecoveryPlanJobExecutionStatus instantiates a new RecoveryPlanJobExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobExecutionStatusWithDefaults

`func NewRecoveryPlanJobExecutionStatusWithDefaults() *RecoveryPlanJobExecutionStatus`

NewRecoveryPlanJobExecutionStatusWithDefaults instantiates a new RecoveryPlanJobExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) GetPostprocessingStatus() RecoveryPlanJobPhaseExecutionStatus`

GetPostprocessingStatus returns the PostprocessingStatus field if non-nil, zero value otherwise.

### GetPostprocessingStatusOk

`func (o *RecoveryPlanJobExecutionStatus) GetPostprocessingStatusOk() (*RecoveryPlanJobPhaseExecutionStatus, bool)`

GetPostprocessingStatusOk returns a tuple with the PostprocessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) SetPostprocessingStatus(v RecoveryPlanJobPhaseExecutionStatus)`

SetPostprocessingStatus sets PostprocessingStatus field to given value.

### HasPostprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) HasPostprocessingStatus() bool`

HasPostprocessingStatus returns a boolean if a field has been set.

### GetOperationStatus

`func (o *RecoveryPlanJobExecutionStatus) GetOperationStatus() RecoveryPlanJobPhaseExecutionStatus`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *RecoveryPlanJobExecutionStatus) GetOperationStatusOk() (*RecoveryPlanJobPhaseExecutionStatus, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *RecoveryPlanJobExecutionStatus) SetOperationStatus(v RecoveryPlanJobPhaseExecutionStatus)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *RecoveryPlanJobExecutionStatus) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetPreprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) GetPreprocessingStatus() RecoveryPlanJobPhaseExecutionStatus`

GetPreprocessingStatus returns the PreprocessingStatus field if non-nil, zero value otherwise.

### GetPreprocessingStatusOk

`func (o *RecoveryPlanJobExecutionStatus) GetPreprocessingStatusOk() (*RecoveryPlanJobPhaseExecutionStatus, bool)`

GetPreprocessingStatusOk returns a tuple with the PreprocessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) SetPreprocessingStatus(v RecoveryPlanJobPhaseExecutionStatus)`

SetPreprocessingStatus sets PreprocessingStatus field to given value.

### HasPreprocessingStatus

`func (o *RecoveryPlanJobExecutionStatus) HasPreprocessingStatus() bool`

HasPreprocessingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


