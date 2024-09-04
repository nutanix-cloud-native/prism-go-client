# MulticlusterConfigSpecResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeSchedule** | Pointer to [**ExecutionSchedule**](ExecutionSchedule.md) |  | [optional] 
**ExecutionPlan** | **string** | Execution plan for multiple requests. | [default to "PARALLEL"]
**SpecList** | [**[]ClusterIntentInput**](ClusterIntentInput.md) |  | 

## Methods

### NewMulticlusterConfigSpecResources

`func NewMulticlusterConfigSpecResources(executionPlan string, specList []ClusterIntentInput, ) *MulticlusterConfigSpecResources`

NewMulticlusterConfigSpecResources instantiates a new MulticlusterConfigSpecResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticlusterConfigSpecResourcesWithDefaults

`func NewMulticlusterConfigSpecResourcesWithDefaults() *MulticlusterConfigSpecResources`

NewMulticlusterConfigSpecResourcesWithDefaults instantiates a new MulticlusterConfigSpecResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeSchedule

`func (o *MulticlusterConfigSpecResources) GetUpgradeSchedule() ExecutionSchedule`

GetUpgradeSchedule returns the UpgradeSchedule field if non-nil, zero value otherwise.

### GetUpgradeScheduleOk

`func (o *MulticlusterConfigSpecResources) GetUpgradeScheduleOk() (*ExecutionSchedule, bool)`

GetUpgradeScheduleOk returns a tuple with the UpgradeSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSchedule

`func (o *MulticlusterConfigSpecResources) SetUpgradeSchedule(v ExecutionSchedule)`

SetUpgradeSchedule sets UpgradeSchedule field to given value.

### HasUpgradeSchedule

`func (o *MulticlusterConfigSpecResources) HasUpgradeSchedule() bool`

HasUpgradeSchedule returns a boolean if a field has been set.

### GetExecutionPlan

`func (o *MulticlusterConfigSpecResources) GetExecutionPlan() string`

GetExecutionPlan returns the ExecutionPlan field if non-nil, zero value otherwise.

### GetExecutionPlanOk

`func (o *MulticlusterConfigSpecResources) GetExecutionPlanOk() (*string, bool)`

GetExecutionPlanOk returns a tuple with the ExecutionPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPlan

`func (o *MulticlusterConfigSpecResources) SetExecutionPlan(v string)`

SetExecutionPlan sets ExecutionPlan field to given value.


### GetSpecList

`func (o *MulticlusterConfigSpecResources) GetSpecList() []ClusterIntentInput`

GetSpecList returns the SpecList field if non-nil, zero value otherwise.

### GetSpecListOk

`func (o *MulticlusterConfigSpecResources) GetSpecListOk() (*[]ClusterIntentInput, bool)`

GetSpecListOk returns a tuple with the SpecList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecList

`func (o *MulticlusterConfigSpecResources) SetSpecList(v []ClusterIntentInput)`

SetSpecList sets SpecList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


