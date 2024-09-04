# TriggerInstanceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceUuid** | Pointer to **string** | Id to uniquely identify trigger in instance list. | [optional] 
**DisplayName** | Pointer to **string** | Trigger display name.  This is for display name, hence ready-only.  | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the trigger instance in the rule.  | [optional] 
**ActionTriggerTypeReference** | [**ActionTriggerTypeReference**](ActionTriggerTypeReference.md) |  | 
**OutputParameters** | Pointer to [**map[string]ParamDescriptor**](ParamDescriptor.md) | The trigger output parameter descriptors.  Came from trigger type.  | [optional] [readonly] 
**InputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 

## Methods

### NewTriggerInstanceObject

`func NewTriggerInstanceObject(actionTriggerTypeReference ActionTriggerTypeReference, ) *TriggerInstanceObject`

NewTriggerInstanceObject instantiates a new TriggerInstanceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInstanceObjectWithDefaults

`func NewTriggerInstanceObjectWithDefaults() *TriggerInstanceObject`

NewTriggerInstanceObjectWithDefaults instantiates a new TriggerInstanceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceUuid

`func (o *TriggerInstanceObject) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *TriggerInstanceObject) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *TriggerInstanceObject) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *TriggerInstanceObject) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetDisplayName

`func (o *TriggerInstanceObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TriggerInstanceObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TriggerInstanceObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TriggerInstanceObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *TriggerInstanceObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TriggerInstanceObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TriggerInstanceObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TriggerInstanceObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionTriggerTypeReference

`func (o *TriggerInstanceObject) GetActionTriggerTypeReference() ActionTriggerTypeReference`

GetActionTriggerTypeReference returns the ActionTriggerTypeReference field if non-nil, zero value otherwise.

### GetActionTriggerTypeReferenceOk

`func (o *TriggerInstanceObject) GetActionTriggerTypeReferenceOk() (*ActionTriggerTypeReference, bool)`

GetActionTriggerTypeReferenceOk returns a tuple with the ActionTriggerTypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTriggerTypeReference

`func (o *TriggerInstanceObject) SetActionTriggerTypeReference(v ActionTriggerTypeReference)`

SetActionTriggerTypeReference sets ActionTriggerTypeReference field to given value.


### GetOutputParameters

`func (o *TriggerInstanceObject) GetOutputParameters() map[string]ParamDescriptor`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *TriggerInstanceObject) GetOutputParametersOk() (*map[string]ParamDescriptor, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *TriggerInstanceObject) SetOutputParameters(v map[string]ParamDescriptor)`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *TriggerInstanceObject) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### GetInputParameterValues

`func (o *TriggerInstanceObject) GetInputParameterValues() map[string]string`

GetInputParameterValues returns the InputParameterValues field if non-nil, zero value otherwise.

### GetInputParameterValuesOk

`func (o *TriggerInstanceObject) GetInputParameterValuesOk() (*map[string]string, bool)`

GetInputParameterValuesOk returns a tuple with the InputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterValues

`func (o *TriggerInstanceObject) SetInputParameterValues(v map[string]string)`

SetInputParameterValues sets InputParameterValues field to given value.

### HasInputParameterValues

`func (o *TriggerInstanceObject) HasInputParameterValues() bool`

HasInputParameterValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


