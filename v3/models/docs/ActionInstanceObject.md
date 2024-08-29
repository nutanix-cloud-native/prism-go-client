# ActionInstanceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceUuid** | **string** | Id to uniquely identify action in instance list. | 
**ActionTypeReference** | [**ActionTypeReference**](ActionTypeReference.md) |  | 
**DisplayName** | Pointer to **string** | Each action instance already has a default display name from action type.  However, users could change an action instance with a different display name.  | [optional] 
**Description** | Pointer to **string** | The description of the action instance in the rule.  For example, an action instance is used twice in a rule, each serves different purpose, it could use this field to describe the purpose.  | [optional] 
**ChildActionUuids** | Pointer to **[]string** | List of possible actions that will be executed after this action. | [optional] 
**MaxRetries** | Pointer to **int32** | This field applies to on_timeout enum retry choice.  When this action execution times out, the rule execution will retry the execution until the max retry number is exceeded.  | [optional] 
**OutputParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The action output parameter descriptors. | [optional] 
**InputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**ShouldContinueOnFailure** | **bool** | When this action execution fails, the rule execution should continue to the next action or not.  | 

## Methods

### NewActionInstanceObject

`func NewActionInstanceObject(instanceUuid string, actionTypeReference ActionTypeReference, shouldContinueOnFailure bool, ) *ActionInstanceObject`

NewActionInstanceObject instantiates a new ActionInstanceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionInstanceObjectWithDefaults

`func NewActionInstanceObjectWithDefaults() *ActionInstanceObject`

NewActionInstanceObjectWithDefaults instantiates a new ActionInstanceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceUuid

`func (o *ActionInstanceObject) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *ActionInstanceObject) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *ActionInstanceObject) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.


### GetActionTypeReference

`func (o *ActionInstanceObject) GetActionTypeReference() ActionTypeReference`

GetActionTypeReference returns the ActionTypeReference field if non-nil, zero value otherwise.

### GetActionTypeReferenceOk

`func (o *ActionInstanceObject) GetActionTypeReferenceOk() (*ActionTypeReference, bool)`

GetActionTypeReferenceOk returns a tuple with the ActionTypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeReference

`func (o *ActionInstanceObject) SetActionTypeReference(v ActionTypeReference)`

SetActionTypeReference sets ActionTypeReference field to given value.


### GetDisplayName

`func (o *ActionInstanceObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ActionInstanceObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ActionInstanceObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ActionInstanceObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ActionInstanceObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionInstanceObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionInstanceObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionInstanceObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChildActionUuids

`func (o *ActionInstanceObject) GetChildActionUuids() []string`

GetChildActionUuids returns the ChildActionUuids field if non-nil, zero value otherwise.

### GetChildActionUuidsOk

`func (o *ActionInstanceObject) GetChildActionUuidsOk() (*[]string, bool)`

GetChildActionUuidsOk returns a tuple with the ChildActionUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildActionUuids

`func (o *ActionInstanceObject) SetChildActionUuids(v []string)`

SetChildActionUuids sets ChildActionUuids field to given value.

### HasChildActionUuids

`func (o *ActionInstanceObject) HasChildActionUuids() bool`

HasChildActionUuids returns a boolean if a field has been set.

### GetMaxRetries

`func (o *ActionInstanceObject) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *ActionInstanceObject) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *ActionInstanceObject) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *ActionInstanceObject) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetOutputParameters

`func (o *ActionInstanceObject) GetOutputParameters() map[string]ParamDescriptor`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *ActionInstanceObject) GetOutputParametersOk() (*map[string]ParamDescriptor, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *ActionInstanceObject) SetOutputParameters(v map[string]ParamDescriptor)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *ActionInstanceObject) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetInputParameterValues

`func (o *ActionInstanceObject) GetInputParameterValues() map[string]string`

GetInputParameterValues returns the InputParameterValues field if non-nil, zero value otherwise.

### GetInputParameterValuesOk

`func (o *ActionInstanceObject) GetInputParameterValuesOk() (*map[string]string, bool)`

GetInputParameterValuesOk returns a tuple with the InputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterValues

`func (o *ActionInstanceObject) SetInputParameterValues(v map[string]string)`

SetInputParameterValues sets InputParameterValues field to given value.

### HasInputParameterValues

`func (o *ActionInstanceObject) HasInputParameterValues() bool`

HasInputParameterValues returns a boolean if a field has been set.

### GetShouldContinueOnFailure

`func (o *ActionInstanceObject) GetShouldContinueOnFailure() bool`

GetShouldContinueOnFailure returns the ShouldContinueOnFailure field if non-nil, zero value otherwise.

### GetShouldContinueOnFailureOk

`func (o *ActionInstanceObject) GetShouldContinueOnFailureOk() (*bool, bool)`

GetShouldContinueOnFailureOk returns a tuple with the ShouldContinueOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldContinueOnFailure

`func (o *ActionInstanceObject) SetShouldContinueOnFailure(v bool)`

SetShouldContinueOnFailure sets ShouldContinueOnFailure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


