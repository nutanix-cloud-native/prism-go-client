# RecoveryPlanStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StageWork** | [**RecoveryPlanStageStageWork**](RecoveryPlanStageStageWork.md) |  | 
**StageUuid** | Pointer to **string** | UUID of stage. | [optional] 
**DelayTimeSecs** | Pointer to **int64** | Amount of time in seconds to delay the execution of next stage after execution of current stage.  | [optional] 

## Methods

### NewRecoveryPlanStage

`func NewRecoveryPlanStage(stageWork RecoveryPlanStageStageWork, ) *RecoveryPlanStage`

NewRecoveryPlanStage instantiates a new RecoveryPlanStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanStageWithDefaults

`func NewRecoveryPlanStageWithDefaults() *RecoveryPlanStage`

NewRecoveryPlanStageWithDefaults instantiates a new RecoveryPlanStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStageWork

`func (o *RecoveryPlanStage) GetStageWork() RecoveryPlanStageStageWork`

GetStageWork returns the StageWork field if non-nil, zero value otherwise.

### GetStageWorkOk

`func (o *RecoveryPlanStage) GetStageWorkOk() (*RecoveryPlanStageStageWork, bool)`

GetStageWorkOk returns a tuple with the StageWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageWork

`func (o *RecoveryPlanStage) SetStageWork(v RecoveryPlanStageStageWork)`

SetStageWork sets StageWork field to given value.


### GetStageUuid

`func (o *RecoveryPlanStage) GetStageUuid() string`

GetStageUuid returns the StageUuid field if non-nil, zero value otherwise.

### GetStageUuidOk

`func (o *RecoveryPlanStage) GetStageUuidOk() (*string, bool)`

GetStageUuidOk returns a tuple with the StageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageUuid

`func (o *RecoveryPlanStage) SetStageUuid(v string)`

SetStageUuid sets StageUuid field to given value.

### HasStageUuid

`func (o *RecoveryPlanStage) HasStageUuid() bool`

HasStageUuid returns a boolean if a field has been set.

### GetDelayTimeSecs

`func (o *RecoveryPlanStage) GetDelayTimeSecs() int64`

GetDelayTimeSecs returns the DelayTimeSecs field if non-nil, zero value otherwise.

### GetDelayTimeSecsOk

`func (o *RecoveryPlanStage) GetDelayTimeSecsOk() (*int64, bool)`

GetDelayTimeSecsOk returns a tuple with the DelayTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTimeSecs

`func (o *RecoveryPlanStage) SetDelayTimeSecs(v int64)`

SetDelayTimeSecs sets DelayTimeSecs field to given value.

### HasDelayTimeSecs

`func (o *RecoveryPlanStage) HasDelayTimeSecs() bool`

HasDelayTimeSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


