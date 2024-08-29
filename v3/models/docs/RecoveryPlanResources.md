# RecoveryPlanResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeGroupRecoveryInfoList** | Pointer to [**[]RecoveryPlanVolumeGroupRecoveryInfo**](RecoveryPlanVolumeGroupRecoveryInfo.md) | Information about Volume Groups to be recovered. | [optional] 
**StageList** | Pointer to [**[]RecoveryPlanStage**](RecoveryPlanStage.md) | Input for the stages of the Recovery Plan. Each stage will perform a predefined type of task. For example, a stage can perform the recovery of the entities specified in a stage.  | [optional] 
**Parameters** | Pointer to [**RecoveryPlanResourcesParameters**](RecoveryPlanResourcesParameters.md) |  | [optional] 

## Methods

### NewRecoveryPlanResources

`func NewRecoveryPlanResources() *RecoveryPlanResources`

NewRecoveryPlanResources instantiates a new RecoveryPlanResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesWithDefaults

`func NewRecoveryPlanResourcesWithDefaults() *RecoveryPlanResources`

NewRecoveryPlanResourcesWithDefaults instantiates a new RecoveryPlanResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeGroupRecoveryInfoList

`func (o *RecoveryPlanResources) GetVolumeGroupRecoveryInfoList() []RecoveryPlanVolumeGroupRecoveryInfo`

GetVolumeGroupRecoveryInfoList returns the VolumeGroupRecoveryInfoList field if non-nil, zero value otherwise.

### GetVolumeGroupRecoveryInfoListOk

`func (o *RecoveryPlanResources) GetVolumeGroupRecoveryInfoListOk() (*[]RecoveryPlanVolumeGroupRecoveryInfo, bool)`

GetVolumeGroupRecoveryInfoListOk returns a tuple with the VolumeGroupRecoveryInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupRecoveryInfoList

`func (o *RecoveryPlanResources) SetVolumeGroupRecoveryInfoList(v []RecoveryPlanVolumeGroupRecoveryInfo)`

SetVolumeGroupRecoveryInfoList sets VolumeGroupRecoveryInfoList field to given value.

### HasVolumeGroupRecoveryInfoList

`func (o *RecoveryPlanResources) HasVolumeGroupRecoveryInfoList() bool`

HasVolumeGroupRecoveryInfoList returns a boolean if a field has been set.

### GetStageList

`func (o *RecoveryPlanResources) GetStageList() []RecoveryPlanStage`

GetStageList returns the StageList field if non-nil, zero value otherwise.

### GetStageListOk

`func (o *RecoveryPlanResources) GetStageListOk() (*[]RecoveryPlanStage, bool)`

GetStageListOk returns a tuple with the StageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageList

`func (o *RecoveryPlanResources) SetStageList(v []RecoveryPlanStage)`

SetStageList sets StageList field to given value.

### HasStageList

`func (o *RecoveryPlanResources) HasStageList() bool`

HasStageList returns a boolean if a field has been set.

### GetParameters

`func (o *RecoveryPlanResources) GetParameters() RecoveryPlanResourcesParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RecoveryPlanResources) GetParametersOk() (*RecoveryPlanResourcesParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RecoveryPlanResources) SetParameters(v RecoveryPlanResourcesParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RecoveryPlanResources) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


