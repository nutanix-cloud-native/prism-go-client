# ActionRuleResultResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionResultList** | [**[]ActionResult**](ActionResult.md) | Ordered list of action execution results. | 
**PostExecutionActionStatus** | Pointer to **string** | The execution status of an action or an action rule. | [optional] 
**ExecutionPlan** | [**RuleExecutionPlan**](RuleExecutionPlan.md) |  | 
**ExecutionStatus** | **string** | The execution status of an action or an action rule. | 
**ActionRuleReference** | [**ActionRuleReference**](ActionRuleReference.md) |  | 
**RuleParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**ExecutionMessages** | Pointer to **[]string** | List of messages associated with rule execution. | [optional] 
**ExecutionStartTime** | Pointer to **time.Time** | The time the first action started execution. | [optional] 
**PostExecutionActionResultList** | Pointer to [**[]ActionResult**](ActionResult.md) | Ordered list of PAF action execution results. | [optional] 
**TriggerInfoList** | [**[]TriggerInfo**](TriggerInfo.md) | The information about the triggers. | 
**ExecutionEndTime** | Pointer to **time.Time** | The time the last action finished execution. | [optional] 

## Methods

### NewActionRuleResultResources

`func NewActionRuleResultResources(actionResultList []ActionResult, executionPlan RuleExecutionPlan, executionStatus string, actionRuleReference ActionRuleReference, triggerInfoList []TriggerInfo, ) *ActionRuleResultResources`

NewActionRuleResultResources instantiates a new ActionRuleResultResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResultResourcesWithDefaults

`func NewActionRuleResultResourcesWithDefaults() *ActionRuleResultResources`

NewActionRuleResultResourcesWithDefaults instantiates a new ActionRuleResultResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionResultList

`func (o *ActionRuleResultResources) GetActionResultList() []ActionResult`

GetActionResultList returns the ActionResultList field if non-nil, zero value otherwise.

### GetActionResultListOk

`func (o *ActionRuleResultResources) GetActionResultListOk() (*[]ActionResult, bool)`

GetActionResultListOk returns a tuple with the ActionResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionResultList

`func (o *ActionRuleResultResources) SetActionResultList(v []ActionResult)`

SetActionResultList sets ActionResultList field to given value.


### GetPostExecutionActionStatus

`func (o *ActionRuleResultResources) GetPostExecutionActionStatus() string`

GetPostExecutionActionStatus returns the PostExecutionActionStatus field if non-nil, zero value otherwise.

### GetPostExecutionActionStatusOk

`func (o *ActionRuleResultResources) GetPostExecutionActionStatusOk() (*string, bool)`

GetPostExecutionActionStatusOk returns a tuple with the PostExecutionActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExecutionActionStatus

`func (o *ActionRuleResultResources) SetPostExecutionActionStatus(v string)`

SetPostExecutionActionStatus sets PostExecutionActionStatus field to given value.

### HasPostExecutionActionStatus

`func (o *ActionRuleResultResources) HasPostExecutionActionStatus() bool`

HasPostExecutionActionStatus returns a boolean if a field has been set.

### GetExecutionPlan

`func (o *ActionRuleResultResources) GetExecutionPlan() RuleExecutionPlan`

GetExecutionPlan returns the ExecutionPlan field if non-nil, zero value otherwise.

### GetExecutionPlanOk

`func (o *ActionRuleResultResources) GetExecutionPlanOk() (*RuleExecutionPlan, bool)`

GetExecutionPlanOk returns a tuple with the ExecutionPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionPlan

`func (o *ActionRuleResultResources) SetExecutionPlan(v RuleExecutionPlan)`

SetExecutionPlan sets ExecutionPlan field to given value.


### GetExecutionStatus

`func (o *ActionRuleResultResources) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *ActionRuleResultResources) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *ActionRuleResultResources) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.


### GetActionRuleReference

`func (o *ActionRuleResultResources) GetActionRuleReference() ActionRuleReference`

GetActionRuleReference returns the ActionRuleReference field if non-nil, zero value otherwise.

### GetActionRuleReferenceOk

`func (o *ActionRuleResultResources) GetActionRuleReferenceOk() (*ActionRuleReference, bool)`

GetActionRuleReferenceOk returns a tuple with the ActionRuleReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRuleReference

`func (o *ActionRuleResultResources) SetActionRuleReference(v ActionRuleReference)`

SetActionRuleReference sets ActionRuleReference field to given value.


### GetRuleParameterValues

`func (o *ActionRuleResultResources) GetRuleParameterValues() map[string]string`

GetRuleParameterValues returns the RuleParameterValues field if non-nil, zero value otherwise.

### GetRuleParameterValuesOk

`func (o *ActionRuleResultResources) GetRuleParameterValuesOk() (*map[string]string, bool)`

GetRuleParameterValuesOk returns a tuple with the RuleParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleParameterValues

`func (o *ActionRuleResultResources) SetRuleParameterValues(v map[string]string)`

SetRuleParameterValues sets RuleParameterValues field to given value.

### HasRuleParameterValues

`func (o *ActionRuleResultResources) HasRuleParameterValues() bool`

HasRuleParameterValues returns a boolean if a field has been set.

### GetExecutionMessages

`func (o *ActionRuleResultResources) GetExecutionMessages() []string`

GetExecutionMessages returns the ExecutionMessages field if non-nil, zero value otherwise.

### GetExecutionMessagesOk

`func (o *ActionRuleResultResources) GetExecutionMessagesOk() (*[]string, bool)`

GetExecutionMessagesOk returns a tuple with the ExecutionMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMessages

`func (o *ActionRuleResultResources) SetExecutionMessages(v []string)`

SetExecutionMessages sets ExecutionMessages field to given value.

### HasExecutionMessages

`func (o *ActionRuleResultResources) HasExecutionMessages() bool`

HasExecutionMessages returns a boolean if a field has been set.

### GetExecutionStartTime

`func (o *ActionRuleResultResources) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *ActionRuleResultResources) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *ActionRuleResultResources) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.

### HasExecutionStartTime

`func (o *ActionRuleResultResources) HasExecutionStartTime() bool`

HasExecutionStartTime returns a boolean if a field has been set.

### GetPostExecutionActionResultList

`func (o *ActionRuleResultResources) GetPostExecutionActionResultList() []ActionResult`

GetPostExecutionActionResultList returns the PostExecutionActionResultList field if non-nil, zero value otherwise.

### GetPostExecutionActionResultListOk

`func (o *ActionRuleResultResources) GetPostExecutionActionResultListOk() (*[]ActionResult, bool)`

GetPostExecutionActionResultListOk returns a tuple with the PostExecutionActionResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExecutionActionResultList

`func (o *ActionRuleResultResources) SetPostExecutionActionResultList(v []ActionResult)`

SetPostExecutionActionResultList sets PostExecutionActionResultList field to given value.

### HasPostExecutionActionResultList

`func (o *ActionRuleResultResources) HasPostExecutionActionResultList() bool`

HasPostExecutionActionResultList returns a boolean if a field has been set.

### GetTriggerInfoList

`func (o *ActionRuleResultResources) GetTriggerInfoList() []TriggerInfo`

GetTriggerInfoList returns the TriggerInfoList field if non-nil, zero value otherwise.

### GetTriggerInfoListOk

`func (o *ActionRuleResultResources) GetTriggerInfoListOk() (*[]TriggerInfo, bool)`

GetTriggerInfoListOk returns a tuple with the TriggerInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerInfoList

`func (o *ActionRuleResultResources) SetTriggerInfoList(v []TriggerInfo)`

SetTriggerInfoList sets TriggerInfoList field to given value.


### GetExecutionEndTime

`func (o *ActionRuleResultResources) GetExecutionEndTime() time.Time`

GetExecutionEndTime returns the ExecutionEndTime field if non-nil, zero value otherwise.

### GetExecutionEndTimeOk

`func (o *ActionRuleResultResources) GetExecutionEndTimeOk() (*time.Time, bool)`

GetExecutionEndTimeOk returns a tuple with the ExecutionEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEndTime

`func (o *ActionRuleResultResources) SetExecutionEndTime(v time.Time)`

SetExecutionEndTime sets ExecutionEndTime field to given value.

### HasExecutionEndTime

`func (o *ActionRuleResultResources) HasExecutionEndTime() bool`

HasExecutionEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


