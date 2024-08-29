# RuleExecutionPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | Pointer to **string** | Action rule types associated with this rule. | [optional] 
**ExecutionUserReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**XPilotParams** | Pointer to [**ActionRuleXPilotParams**](ActionRuleXPilotParams.md) |  | [optional] 
**ActionList** | [**[]ActionInstanceObject**](ActionInstanceObject.md) | The list of actions to be executed. | 
**PostExecutionActionList** | Pointer to [**[]ActionInstanceObject**](ActionInstanceObject.md) | The list of actions to be executed in the end. | [optional] 

## Methods

### NewRuleExecutionPlan

`func NewRuleExecutionPlan(actionList []ActionInstanceObject, ) *RuleExecutionPlan`

NewRuleExecutionPlan instantiates a new RuleExecutionPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleExecutionPlanWithDefaults

`func NewRuleExecutionPlanWithDefaults() *RuleExecutionPlan`

NewRuleExecutionPlanWithDefaults instantiates a new RuleExecutionPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *RuleExecutionPlan) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *RuleExecutionPlan) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *RuleExecutionPlan) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *RuleExecutionPlan) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetExecutionUserReference

`func (o *RuleExecutionPlan) GetExecutionUserReference() UserReference`

GetExecutionUserReference returns the ExecutionUserReference field if non-nil, zero value otherwise.

### GetExecutionUserReferenceOk

`func (o *RuleExecutionPlan) GetExecutionUserReferenceOk() (*UserReference, bool)`

GetExecutionUserReferenceOk returns a tuple with the ExecutionUserReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionUserReference

`func (o *RuleExecutionPlan) SetExecutionUserReference(v UserReference)`

SetExecutionUserReference sets ExecutionUserReference field to given value.

### HasExecutionUserReference

`func (o *RuleExecutionPlan) HasExecutionUserReference() bool`

HasExecutionUserReference returns a boolean if a field has been set.

### GetXPilotParams

`func (o *RuleExecutionPlan) GetXPilotParams() ActionRuleXPilotParams`

GetXPilotParams returns the XPilotParams field if non-nil, zero value otherwise.

### GetXPilotParamsOk

`func (o *RuleExecutionPlan) GetXPilotParamsOk() (*ActionRuleXPilotParams, bool)`

GetXPilotParamsOk returns a tuple with the XPilotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXPilotParams

`func (o *RuleExecutionPlan) SetXPilotParams(v ActionRuleXPilotParams)`

SetXPilotParams sets XPilotParams field to given value.

### HasXPilotParams

`func (o *RuleExecutionPlan) HasXPilotParams() bool`

HasXPilotParams returns a boolean if a field has been set.

### GetActionList

`func (o *RuleExecutionPlan) GetActionList() []ActionInstanceObject`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *RuleExecutionPlan) GetActionListOk() (*[]ActionInstanceObject, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *RuleExecutionPlan) SetActionList(v []ActionInstanceObject)`

SetActionList sets ActionList field to given value.


### GetPostExecutionActionList

`func (o *RuleExecutionPlan) GetPostExecutionActionList() []ActionInstanceObject`

GetPostExecutionActionList returns the PostExecutionActionList field if non-nil, zero value otherwise.

### GetPostExecutionActionListOk

`func (o *RuleExecutionPlan) GetPostExecutionActionListOk() (*[]ActionInstanceObject, bool)`

GetPostExecutionActionListOk returns a tuple with the PostExecutionActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExecutionActionList

`func (o *RuleExecutionPlan) SetPostExecutionActionList(v []ActionInstanceObject)`

SetPostExecutionActionList sets PostExecutionActionList field to given value.

### HasPostExecutionActionList

`func (o *RuleExecutionPlan) HasPostExecutionActionList() bool`

HasPostExecutionActionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


