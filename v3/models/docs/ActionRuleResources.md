# ActionRuleResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | is the rule enabled or disabled. | [optional] 
**ShouldValidate** | Pointer to **bool** | The rule should be validated or not.  If True, then, the rule will be validated before saving.  If the validation failed, the spec status message list will have errors.  | [optional] 
**TriggerList** | [**[]TriggerInstanceObject**](TriggerInstanceObject.md) | The only trigger that the rule has. | 
**Name** | Pointer to **string** | The action rule name | [optional] 
**ExecutionUserReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**RuleType** | Pointer to **string** | Action rule types associated with this rule. | [optional] 
**PostExecutionActionList** | Pointer to [**[]ActionInstanceObject**](ActionInstanceObject.md) | Ordered list of the actions to be executed in the end. | [optional] 
**ActionList** | [**[]ActionInstanceObject**](ActionInstanceObject.md) | Ordered list of the actions to execute in this rule. | 
**CheckTriggerValidity** | Pointer to **bool** | If it is true, the rule action execution at the schedueled time will check if the original trigger is still valid.  If not set, default to True.  | [optional] 
**GlobalParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The rule level global parameter descriptors. This is provided by the system.  | [optional] [readonly] 
**XPilotParams** | Pointer to [**ActionRuleXPilotParams**](ActionRuleXPilotParams.md) |  | [optional] 
**ErrorList** | Pointer to [**[]ComponentError**](ComponentError.md) | Errors identified in this rule. | [optional] 
**Validated** | Pointer to **bool** | Is this been validated to trur or not. | [optional] [readonly] 
**Description** | Pointer to **string** | The rule description | [optional] 

## Methods

### NewActionRuleResources

`func NewActionRuleResources(triggerList []TriggerInstanceObject, actionList []ActionInstanceObject, ) *ActionRuleResources`

NewActionRuleResources instantiates a new ActionRuleResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResourcesWithDefaults

`func NewActionRuleResourcesWithDefaults() *ActionRuleResources`

NewActionRuleResourcesWithDefaults instantiates a new ActionRuleResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ActionRuleResources) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ActionRuleResources) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ActionRuleResources) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ActionRuleResources) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetShouldValidate

`func (o *ActionRuleResources) GetShouldValidate() bool`

GetShouldValidate returns the ShouldValidate field if non-nil, zero value otherwise.

### GetShouldValidateOk

`func (o *ActionRuleResources) GetShouldValidateOk() (*bool, bool)`

GetShouldValidateOk returns a tuple with the ShouldValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldValidate

`func (o *ActionRuleResources) SetShouldValidate(v bool)`

SetShouldValidate sets ShouldValidate field to given value.

### HasShouldValidate

`func (o *ActionRuleResources) HasShouldValidate() bool`

HasShouldValidate returns a boolean if a field has been set.

### GetTriggerList

`func (o *ActionRuleResources) GetTriggerList() []TriggerInstanceObject`

GetTriggerList returns the TriggerList field if non-nil, zero value otherwise.

### GetTriggerListOk

`func (o *ActionRuleResources) GetTriggerListOk() (*[]TriggerInstanceObject, bool)`

GetTriggerListOk returns a tuple with the TriggerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerList

`func (o *ActionRuleResources) SetTriggerList(v []TriggerInstanceObject)`

SetTriggerList sets TriggerList field to given value.


### GetName

`func (o *ActionRuleResources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionRuleResources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionRuleResources) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionRuleResources) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionUserReference

`func (o *ActionRuleResources) GetExecutionUserReference() UserReference`

GetExecutionUserReference returns the ExecutionUserReference field if non-nil, zero value otherwise.

### GetExecutionUserReferenceOk

`func (o *ActionRuleResources) GetExecutionUserReferenceOk() (*UserReference, bool)`

GetExecutionUserReferenceOk returns a tuple with the ExecutionUserReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionUserReference

`func (o *ActionRuleResources) SetExecutionUserReference(v UserReference)`

SetExecutionUserReference sets ExecutionUserReference field to given value.

### HasExecutionUserReference

`func (o *ActionRuleResources) HasExecutionUserReference() bool`

HasExecutionUserReference returns a boolean if a field has been set.

### GetRuleType

`func (o *ActionRuleResources) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ActionRuleResources) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ActionRuleResources) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *ActionRuleResources) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPostExecutionActionList

`func (o *ActionRuleResources) GetPostExecutionActionList() []ActionInstanceObject`

GetPostExecutionActionList returns the PostExecutionActionList field if non-nil, zero value otherwise.

### GetPostExecutionActionListOk

`func (o *ActionRuleResources) GetPostExecutionActionListOk() (*[]ActionInstanceObject, bool)`

GetPostExecutionActionListOk returns a tuple with the PostExecutionActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostExecutionActionList

`func (o *ActionRuleResources) SetPostExecutionActionList(v []ActionInstanceObject)`

SetPostExecutionActionList sets PostExecutionActionList field to given value.

### HasPostExecutionActionList

`func (o *ActionRuleResources) HasPostExecutionActionList() bool`

HasPostExecutionActionList returns a boolean if a field has been set.

### GetActionList

`func (o *ActionRuleResources) GetActionList() []ActionInstanceObject`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *ActionRuleResources) GetActionListOk() (*[]ActionInstanceObject, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *ActionRuleResources) SetActionList(v []ActionInstanceObject)`

SetActionList sets ActionList field to given value.


### GetCheckTriggerValidity

`func (o *ActionRuleResources) GetCheckTriggerValidity() bool`

GetCheckTriggerValidity returns the CheckTriggerValidity field if non-nil, zero value otherwise.

### GetCheckTriggerValidityOk

`func (o *ActionRuleResources) GetCheckTriggerValidityOk() (*bool, bool)`

GetCheckTriggerValidityOk returns a tuple with the CheckTriggerValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTriggerValidity

`func (o *ActionRuleResources) SetCheckTriggerValidity(v bool)`

SetCheckTriggerValidity sets CheckTriggerValidity field to given value.

### HasCheckTriggerValidity

`func (o *ActionRuleResources) HasCheckTriggerValidity() bool`

HasCheckTriggerValidity returns a boolean if a field has been set.

### GetGlobalParameters

`func (o *ActionRuleResources) GetGlobalParameters() map[string]ParamDescriptor`

GetGlobalParameters returns the GlobalParameters field if non-nil, zero value otherwise.

### GetGlobalParametersOk

`func (o *ActionRuleResources) GetGlobalParametersOk() (*map[string]ParamDescriptor, bool)`

GetGlobalParametersOk returns a tuple with the GlobalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalParameters

`func (o *ActionRuleResources) SetGlobalParameters(v map[string]ParamDescriptor)`

SetGlobalParameters sets GlobalParameters field to given value.

### HasGlobalParameters

`func (o *ActionRuleResources) HasGlobalParameters() bool`

HasGlobalParameters returns a boolean if a field has been set.

### GetXPilotParams

`func (o *ActionRuleResources) GetXPilotParams() ActionRuleXPilotParams`

GetXPilotParams returns the XPilotParams field if non-nil, zero value otherwise.

### GetXPilotParamsOk

`func (o *ActionRuleResources) GetXPilotParamsOk() (*ActionRuleXPilotParams, bool)`

GetXPilotParamsOk returns a tuple with the XPilotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXPilotParams

`func (o *ActionRuleResources) SetXPilotParams(v ActionRuleXPilotParams)`

SetXPilotParams sets XPilotParams field to given value.

### HasXPilotParams

`func (o *ActionRuleResources) HasXPilotParams() bool`

HasXPilotParams returns a boolean if a field has been set.

### GetErrorList

`func (o *ActionRuleResources) GetErrorList() []ComponentError`

GetErrorList returns the ErrorList field if non-nil, zero value otherwise.

### GetErrorListOk

`func (o *ActionRuleResources) GetErrorListOk() (*[]ComponentError, bool)`

GetErrorListOk returns a tuple with the ErrorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorList

`func (o *ActionRuleResources) SetErrorList(v []ComponentError)`

SetErrorList sets ErrorList field to given value.

### HasErrorList

`func (o *ActionRuleResources) HasErrorList() bool`

HasErrorList returns a boolean if a field has been set.

### GetValidated

`func (o *ActionRuleResources) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *ActionRuleResources) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *ActionRuleResources) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *ActionRuleResources) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetDescription

`func (o *ActionRuleResources) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionRuleResources) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionRuleResources) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionRuleResources) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


