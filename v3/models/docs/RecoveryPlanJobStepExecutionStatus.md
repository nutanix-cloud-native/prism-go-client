# RecoveryPlanJobStepExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of this step. | 
**ErrorDetail** | Pointer to **string** | The error detail for the step failure. | [optional] 
**AnyEntityReferenceList** | Pointer to [**[]Reference**](Reference.md) | Reference to the list of the entities on which the step is performing action.  | [optional] 
**StartTime** | Pointer to **time.Time** | Time when this step started. | [optional] 
**RecoveredEntityInfoList** | Pointer to [**[]RecoveredEntityInformation**](RecoveredEntityInformation.md) | List containing the details about the recovered entities.  | [optional] 
**PercentageComplete** | Pointer to **int32** | Percentage of step completed. | [optional] 
**StepSequenceNumber** | Pointer to **int64** | Sequence number of the step among its siblings. This can be used for ordering the sub-steps for a step.  | [optional] 
**EndTime** | Pointer to **time.Time** | Time when this step ended. | [optional] 
**OperationType** | **string** | Type of operation being performed. | 
**ParentStepUuid** | Pointer to **string** | UUID of the parent step.  | [optional] 
**Message** | Pointer to **string** | User readable message for the action being done for the step.  | [optional] 
**ErrorCode** | Pointer to **string** | The error code for the step failure. | [optional] 
**StepUuid** | **string** | UUID of a step. | 

## Methods

### NewRecoveryPlanJobStepExecutionStatus

`func NewRecoveryPlanJobStepExecutionStatus(status string, operationType string, stepUuid string, ) *RecoveryPlanJobStepExecutionStatus`

NewRecoveryPlanJobStepExecutionStatus instantiates a new RecoveryPlanJobStepExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobStepExecutionStatusWithDefaults

`func NewRecoveryPlanJobStepExecutionStatusWithDefaults() *RecoveryPlanJobStepExecutionStatus`

NewRecoveryPlanJobStepExecutionStatusWithDefaults instantiates a new RecoveryPlanJobStepExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobStepExecutionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobStepExecutionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorDetail

`func (o *RecoveryPlanJobStepExecutionStatus) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *RecoveryPlanJobStepExecutionStatus) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *RecoveryPlanJobStepExecutionStatus) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetAnyEntityReferenceList

`func (o *RecoveryPlanJobStepExecutionStatus) GetAnyEntityReferenceList() []Reference`

GetAnyEntityReferenceList returns the AnyEntityReferenceList field if non-nil, zero value otherwise.

### GetAnyEntityReferenceListOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetAnyEntityReferenceListOk() (*[]Reference, bool)`

GetAnyEntityReferenceListOk returns a tuple with the AnyEntityReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEntityReferenceList

`func (o *RecoveryPlanJobStepExecutionStatus) SetAnyEntityReferenceList(v []Reference)`

SetAnyEntityReferenceList sets AnyEntityReferenceList field to given value.

### HasAnyEntityReferenceList

`func (o *RecoveryPlanJobStepExecutionStatus) HasAnyEntityReferenceList() bool`

HasAnyEntityReferenceList returns a boolean if a field has been set.

### GetStartTime

`func (o *RecoveryPlanJobStepExecutionStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RecoveryPlanJobStepExecutionStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RecoveryPlanJobStepExecutionStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRecoveredEntityInfoList

`func (o *RecoveryPlanJobStepExecutionStatus) GetRecoveredEntityInfoList() []RecoveredEntityInformation`

GetRecoveredEntityInfoList returns the RecoveredEntityInfoList field if non-nil, zero value otherwise.

### GetRecoveredEntityInfoListOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetRecoveredEntityInfoListOk() (*[]RecoveredEntityInformation, bool)`

GetRecoveredEntityInfoListOk returns a tuple with the RecoveredEntityInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveredEntityInfoList

`func (o *RecoveryPlanJobStepExecutionStatus) SetRecoveredEntityInfoList(v []RecoveredEntityInformation)`

SetRecoveredEntityInfoList sets RecoveredEntityInfoList field to given value.

### HasRecoveredEntityInfoList

`func (o *RecoveryPlanJobStepExecutionStatus) HasRecoveredEntityInfoList() bool`

HasRecoveredEntityInfoList returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *RecoveryPlanJobStepExecutionStatus) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *RecoveryPlanJobStepExecutionStatus) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.

### HasPercentageComplete

`func (o *RecoveryPlanJobStepExecutionStatus) HasPercentageComplete() bool`

HasPercentageComplete returns a boolean if a field has been set.

### GetStepSequenceNumber

`func (o *RecoveryPlanJobStepExecutionStatus) GetStepSequenceNumber() int64`

GetStepSequenceNumber returns the StepSequenceNumber field if non-nil, zero value otherwise.

### GetStepSequenceNumberOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetStepSequenceNumberOk() (*int64, bool)`

GetStepSequenceNumberOk returns a tuple with the StepSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSequenceNumber

`func (o *RecoveryPlanJobStepExecutionStatus) SetStepSequenceNumber(v int64)`

SetStepSequenceNumber sets StepSequenceNumber field to given value.

### HasStepSequenceNumber

`func (o *RecoveryPlanJobStepExecutionStatus) HasStepSequenceNumber() bool`

HasStepSequenceNumber returns a boolean if a field has been set.

### GetEndTime

`func (o *RecoveryPlanJobStepExecutionStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RecoveryPlanJobStepExecutionStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RecoveryPlanJobStepExecutionStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetOperationType

`func (o *RecoveryPlanJobStepExecutionStatus) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RecoveryPlanJobStepExecutionStatus) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetParentStepUuid

`func (o *RecoveryPlanJobStepExecutionStatus) GetParentStepUuid() string`

GetParentStepUuid returns the ParentStepUuid field if non-nil, zero value otherwise.

### GetParentStepUuidOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetParentStepUuidOk() (*string, bool)`

GetParentStepUuidOk returns a tuple with the ParentStepUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepUuid

`func (o *RecoveryPlanJobStepExecutionStatus) SetParentStepUuid(v string)`

SetParentStepUuid sets ParentStepUuid field to given value.

### HasParentStepUuid

`func (o *RecoveryPlanJobStepExecutionStatus) HasParentStepUuid() bool`

HasParentStepUuid returns a boolean if a field has been set.

### GetMessage

`func (o *RecoveryPlanJobStepExecutionStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecoveryPlanJobStepExecutionStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RecoveryPlanJobStepExecutionStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *RecoveryPlanJobStepExecutionStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RecoveryPlanJobStepExecutionStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *RecoveryPlanJobStepExecutionStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetStepUuid

`func (o *RecoveryPlanJobStepExecutionStatus) GetStepUuid() string`

GetStepUuid returns the StepUuid field if non-nil, zero value otherwise.

### GetStepUuidOk

`func (o *RecoveryPlanJobStepExecutionStatus) GetStepUuidOk() (*string, bool)`

GetStepUuidOk returns a tuple with the StepUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUuid

`func (o *RecoveryPlanJobStepExecutionStatus) SetStepUuid(v string)`

SetStepUuid sets StepUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


