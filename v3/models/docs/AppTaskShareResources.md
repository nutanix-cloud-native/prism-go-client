# AppTaskShareResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**ProjectReferenceList** | Pointer to [**[]ProjectReference**](ProjectReference.md) | The projects this task has been assigned to | [optional] 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReference**](AppTaskReference.md) |  | [optional] 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 

## Methods

### NewAppTaskShareResources

`func NewAppTaskShareResources(type_ string, ) *AppTaskShareResources`

NewAppTaskShareResources instantiates a new AppTaskShareResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskShareResourcesWithDefaults

`func NewAppTaskShareResourcesWithDefaults() *AppTaskShareResources`

NewAppTaskShareResourcesWithDefaults instantiates a new AppTaskShareResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskShareResources) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskShareResources) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskShareResources) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskShareResources) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskShareResources) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskShareResources) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskShareResources) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskShareResources) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetProjectReferenceList

`func (o *AppTaskShareResources) GetProjectReferenceList() []ProjectReference`

GetProjectReferenceList returns the ProjectReferenceList field if non-nil, zero value otherwise.

### GetProjectReferenceListOk

`func (o *AppTaskShareResources) GetProjectReferenceListOk() (*[]ProjectReference, bool)`

GetProjectReferenceListOk returns a tuple with the ProjectReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReferenceList

`func (o *AppTaskShareResources) SetProjectReferenceList(v []ProjectReference)`

SetProjectReferenceList sets ProjectReferenceList field to given value.

### HasProjectReferenceList

`func (o *AppTaskShareResources) HasProjectReferenceList() bool`

HasProjectReferenceList returns a boolean if a field has been set.

### GetChildTasksLocalReferenceList

`func (o *AppTaskShareResources) GetChildTasksLocalReferenceList() []AppTaskReference`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskShareResources) GetChildTasksLocalReferenceListOk() (*[]AppTaskReference, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskShareResources) SetChildTasksLocalReferenceList(v []AppTaskReference)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskShareResources) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetAttrs

`func (o *AppTaskShareResources) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskShareResources) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskShareResources) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskShareResources) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskShareResources) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskShareResources) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskShareResources) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskShareResources) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskShareResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskShareResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskShareResources) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskShareResources) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskShareResources) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskShareResources) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskShareResources) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


