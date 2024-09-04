# ActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionInstanceIndex** | Pointer to **int64** | Index of action instance in action rule. | [optional] 
**ActionTypeReference** | [**ActionTypeReference**](ActionTypeReference.md) |  | 
**OutputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**ActionTypeDisplayName** | Pointer to **string** | the display name | [optional] 
**ErrorMessageList** | Pointer to **[]string** | A list of error messages if failed, or if timed out.  For actions that have rule&#39;s on_timeout property set to RETRY, each of the retry could have one timeout message.  | [optional] 
**ExecutionStatus** | **string** | The execution status of an action or an action rule. | 
**InfoMessageList** | Pointer to **[]string** | A list of informational messages associated with the actions result.  | [optional] 
**ExecutionStartTime** | **time.Time** | Denotes the timestamp when the action is invoked. | 
**InputParameterValues** | Pointer to **map[string]string** | The trigger or action required input parameter value map, or the output parameters.  | [optional] 
**ExecutionEndTime** | Pointer to **time.Time** | Denotes the timestamp when the action is completed. | [optional] 

## Methods

### NewActionResult

`func NewActionResult(actionTypeReference ActionTypeReference, executionStatus string, executionStartTime time.Time, ) *ActionResult`

NewActionResult instantiates a new ActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionResultWithDefaults

`func NewActionResultWithDefaults() *ActionResult`

NewActionResultWithDefaults instantiates a new ActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionInstanceIndex

`func (o *ActionResult) GetActionInstanceIndex() int64`

GetActionInstanceIndex returns the ActionInstanceIndex field if non-nil, zero value otherwise.

### GetActionInstanceIndexOk

`func (o *ActionResult) GetActionInstanceIndexOk() (*int64, bool)`

GetActionInstanceIndexOk returns a tuple with the ActionInstanceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionInstanceIndex

`func (o *ActionResult) SetActionInstanceIndex(v int64)`

SetActionInstanceIndex sets ActionInstanceIndex field to given value.

### HasActionInstanceIndex

`func (o *ActionResult) HasActionInstanceIndex() bool`

HasActionInstanceIndex returns a boolean if a field has been set.

### GetActionTypeReference

`func (o *ActionResult) GetActionTypeReference() ActionTypeReference`

GetActionTypeReference returns the ActionTypeReference field if non-nil, zero value otherwise.

### GetActionTypeReferenceOk

`func (o *ActionResult) GetActionTypeReferenceOk() (*ActionTypeReference, bool)`

GetActionTypeReferenceOk returns a tuple with the ActionTypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeReference

`func (o *ActionResult) SetActionTypeReference(v ActionTypeReference)`

SetActionTypeReference sets ActionTypeReference field to given value.


### GetOutputParameterValues

`func (o *ActionResult) GetOutputParameterValues() map[string]string`

GetOutputParameterValues returns the OutputParameterValues field if non-nil, zero value otherwise.

### GetOutputParameterValuesOk

`func (o *ActionResult) GetOutputParameterValuesOk() (*map[string]string, bool)`

GetOutputParameterValuesOk returns a tuple with the OutputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameterValues

`func (o *ActionResult) SetOutputParameterValues(v map[string]string)`

SetOutputParameterValues sets OutputParameterValues field to given value.

### HasOutputParameterValues

`func (o *ActionResult) HasOutputParameterValues() bool`

HasOutputParameterValues returns a boolean if a field has been set.

### GetActionTypeDisplayName

`func (o *ActionResult) GetActionTypeDisplayName() string`

GetActionTypeDisplayName returns the ActionTypeDisplayName field if non-nil, zero value otherwise.

### GetActionTypeDisplayNameOk

`func (o *ActionResult) GetActionTypeDisplayNameOk() (*string, bool)`

GetActionTypeDisplayNameOk returns a tuple with the ActionTypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeDisplayName

`func (o *ActionResult) SetActionTypeDisplayName(v string)`

SetActionTypeDisplayName sets ActionTypeDisplayName field to given value.

### HasActionTypeDisplayName

`func (o *ActionResult) HasActionTypeDisplayName() bool`

HasActionTypeDisplayName returns a boolean if a field has been set.

### GetErrorMessageList

`func (o *ActionResult) GetErrorMessageList() []string`

GetErrorMessageList returns the ErrorMessageList field if non-nil, zero value otherwise.

### GetErrorMessageListOk

`func (o *ActionResult) GetErrorMessageListOk() (*[]string, bool)`

GetErrorMessageListOk returns a tuple with the ErrorMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessageList

`func (o *ActionResult) SetErrorMessageList(v []string)`

SetErrorMessageList sets ErrorMessageList field to given value.

### HasErrorMessageList

`func (o *ActionResult) HasErrorMessageList() bool`

HasErrorMessageList returns a boolean if a field has been set.

### GetExecutionStatus

`func (o *ActionResult) GetExecutionStatus() string`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *ActionResult) GetExecutionStatusOk() (*string, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *ActionResult) SetExecutionStatus(v string)`

SetExecutionStatus sets ExecutionStatus field to given value.


### GetInfoMessageList

`func (o *ActionResult) GetInfoMessageList() []string`

GetInfoMessageList returns the InfoMessageList field if non-nil, zero value otherwise.

### GetInfoMessageListOk

`func (o *ActionResult) GetInfoMessageListOk() (*[]string, bool)`

GetInfoMessageListOk returns a tuple with the InfoMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMessageList

`func (o *ActionResult) SetInfoMessageList(v []string)`

SetInfoMessageList sets InfoMessageList field to given value.

### HasInfoMessageList

`func (o *ActionResult) HasInfoMessageList() bool`

HasInfoMessageList returns a boolean if a field has been set.

### GetExecutionStartTime

`func (o *ActionResult) GetExecutionStartTime() time.Time`

GetExecutionStartTime returns the ExecutionStartTime field if non-nil, zero value otherwise.

### GetExecutionStartTimeOk

`func (o *ActionResult) GetExecutionStartTimeOk() (*time.Time, bool)`

GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStartTime

`func (o *ActionResult) SetExecutionStartTime(v time.Time)`

SetExecutionStartTime sets ExecutionStartTime field to given value.


### GetInputParameterValues

`func (o *ActionResult) GetInputParameterValues() map[string]string`

GetInputParameterValues returns the InputParameterValues field if non-nil, zero value otherwise.

### GetInputParameterValuesOk

`func (o *ActionResult) GetInputParameterValuesOk() (*map[string]string, bool)`

GetInputParameterValuesOk returns a tuple with the InputParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterValues

`func (o *ActionResult) SetInputParameterValues(v map[string]string)`

SetInputParameterValues sets InputParameterValues field to given value.

### HasInputParameterValues

`func (o *ActionResult) HasInputParameterValues() bool`

HasInputParameterValues returns a boolean if a field has been set.

### GetExecutionEndTime

`func (o *ActionResult) GetExecutionEndTime() time.Time`

GetExecutionEndTime returns the ExecutionEndTime field if non-nil, zero value otherwise.

### GetExecutionEndTimeOk

`func (o *ActionResult) GetExecutionEndTimeOk() (*time.Time, bool)`

GetExecutionEndTimeOk returns a tuple with the ExecutionEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEndTime

`func (o *ActionResult) SetExecutionEndTime(v time.Time)`

SetExecutionEndTime sets ExecutionEndTime field to given value.

### HasExecutionEndTime

`func (o *ActionResult) HasExecutionEndTime() bool`

HasExecutionEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


