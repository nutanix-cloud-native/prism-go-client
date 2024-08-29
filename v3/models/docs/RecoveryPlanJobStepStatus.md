# RecoveryPlanJobStepStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | State of a step. | 
**PercentageComplete** | **int32** | Percentage completed for a step. | 

## Methods

### NewRecoveryPlanJobStepStatus

`func NewRecoveryPlanJobStepStatus(status string, percentageComplete int32, ) *RecoveryPlanJobStepStatus`

NewRecoveryPlanJobStepStatus instantiates a new RecoveryPlanJobStepStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobStepStatusWithDefaults

`func NewRecoveryPlanJobStepStatusWithDefaults() *RecoveryPlanJobStepStatus`

NewRecoveryPlanJobStepStatusWithDefaults instantiates a new RecoveryPlanJobStepStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobStepStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobStepStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobStepStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPercentageComplete

`func (o *RecoveryPlanJobStepStatus) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *RecoveryPlanJobStepStatus) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *RecoveryPlanJobStepStatus) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


