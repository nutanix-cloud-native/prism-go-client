# TriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTriggerTypeReference** | [**ActionTriggerTypeReference**](ActionTriggerTypeReference.md) |  | 
**TriggerTypeDisplayName** | Pointer to **string** | The trigger display name.  Default to trigger type display name.  | [optional] 
**TriggerTime** | **time.Time** | The time that this trigger happened | 
**InputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**OutputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 

## Methods

### NewTriggerInfo

`func NewTriggerInfo(actionTriggerTypeReference ActionTriggerTypeReference, triggerTime time.Time, ) *TriggerInfo`

NewTriggerInfo instantiates a new TriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerInfoWithDefaults

`func NewTriggerInfoWithDefaults() *TriggerInfo`

NewTriggerInfoWithDefaults instantiates a new TriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTriggerTypeReference

`func (o *TriggerInfo) GetActionTriggerTypeReference() ActionTriggerTypeReference`

GetActionTriggerTypeReference returns the ActionTriggerTypeReference field if non-nil, zero value otherwise.

### GetActionTriggerTypeReferenceOk

`func (o *TriggerInfo) GetActionTriggerTypeReferenceOk() (*ActionTriggerTypeReference, bool)`

GetActionTriggerTypeReferenceOk returns a tuple with the ActionTriggerTypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTriggerTypeReference

`func (o *TriggerInfo) SetActionTriggerTypeReference(v ActionTriggerTypeReference)`

SetActionTriggerTypeReference sets ActionTriggerTypeReference field to given value.


### GetTriggerTypeDisplayName

`func (o *TriggerInfo) GetTriggerTypeDisplayName() string`

GetTriggerTypeDisplayName returns the TriggerTypeDisplayName field if non-nil, zero value otherwise.

### GetTriggerTypeDisplayNameOk

`func (o *TriggerInfo) GetTriggerTypeDisplayNameOk() (*string, bool)`

GetTriggerTypeDisplayNameOk returns a tuple with the TriggerTypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTypeDisplayName

`func (o *TriggerInfo) SetTriggerTypeDisplayName(v string)`

SetTriggerTypeDisplayName sets TriggerTypeDisplayName field to given value.

### HasTriggerTypeDisplayName

`func (o *TriggerInfo) HasTriggerTypeDisplayName() bool`

HasTriggerTypeDisplayName returns a boolean if a field has been set.

### GetTriggerTime

`func (o *TriggerInfo) GetTriggerTime() time.Time`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *TriggerInfo) GetTriggerTimeOk() (*time.Time, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *TriggerInfo) SetTriggerTime(v time.Time)`

SetTriggerTime sets TriggerTime field to given value.


### GetInputParameterValues

`func (o *TriggerInfo) GetInputParameterValues() map[string]string`

GetInputParameterValues returns the InputParameterValues field if non-nil, zero value otherwise.

### GetInputParameterValuesOk

`func (o *TriggerInfo) GetInputParameterValuesOk() (*map[string]string, bool)`

GetInputParameterValuesOk returns a tuple with the InputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterValues

`func (o *TriggerInfo) SetInputParameterValues(v map[string]string)`

SetInputParameterValues sets InputParameterValues field to given value.

### HasInputParameterValues

`func (o *TriggerInfo) HasInputParameterValues() bool`

HasInputParameterValues returns a boolean if a field has been set.

### GetOutputParameterValues

`func (o *TriggerInfo) GetOutputParameterValues() map[string]string`

GetOutputParameterValues returns the OutputParameterValues field if non-nil, zero value otherwise.

### GetOutputParameterValuesOk

`func (o *TriggerInfo) GetOutputParameterValuesOk() (*map[string]string, bool)`

GetOutputParameterValuesOk returns a tuple with the OutputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameterValues

`func (o *TriggerInfo) SetOutputParameterValues(v map[string]string)`

SetOutputParameterValues sets OutputParameterValues field to given value.

### HasOutputParameterValues

`func (o *TriggerInfo) HasOutputParameterValues() bool`

HasOutputParameterValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


