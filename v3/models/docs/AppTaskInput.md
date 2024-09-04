# AppTaskInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReference**](AppTaskReference.md) |  | [optional] 
**Name** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppTaskInput

`func NewAppTaskInput(name string, type_ string, uUID string, ) *AppTaskInput`

NewAppTaskInput instantiates a new AppTaskInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskInputWithDefaults

`func NewAppTaskInputWithDefaults() *AppTaskInput`

NewAppTaskInputWithDefaults instantiates a new AppTaskInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskInput) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskInput) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskInput) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskInput) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskInput) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskInput) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskInput) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskInput) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetDescription

`func (o *AppTaskInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppTaskInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppTaskInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppTaskInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChildTasksLocalReferenceList

`func (o *AppTaskInput) GetChildTasksLocalReferenceList() []AppTaskReference`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskInput) GetChildTasksLocalReferenceListOk() (*[]AppTaskReference, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskInput) SetChildTasksLocalReferenceList(v []AppTaskReference)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskInput) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppTaskInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppTaskInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppTaskInput) SetName(v string)`

SetName sets Name field to given value.


### GetAttrs

`func (o *AppTaskInput) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskInput) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskInput) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskInput) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskInput) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskInput) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskInput) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskInput) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskInput) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppTaskInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppTaskInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppTaskInput) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


