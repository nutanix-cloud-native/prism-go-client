# AppTaskResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list. | 
**ProjectReferenceList** | Pointer to [**[]ProjectReference**](ProjectReference.md) | The projects this task has been assigned to | [optional] 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReference**](AppTaskReference.md) |  | [optional] 
**State** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 

## Methods

### NewAppTaskResourcesDefStatus

`func NewAppTaskResourcesDefStatus(messageList []MessageResource, state string, type_ string, ) *AppTaskResourcesDefStatus`

NewAppTaskResourcesDefStatus instantiates a new AppTaskResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskResourcesDefStatusWithDefaults

`func NewAppTaskResourcesDefStatusWithDefaults() *AppTaskResourcesDefStatus`

NewAppTaskResourcesDefStatusWithDefaults instantiates a new AppTaskResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskResourcesDefStatus) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskResourcesDefStatus) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskResourcesDefStatus) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskResourcesDefStatus) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskResourcesDefStatus) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskResourcesDefStatus) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskResourcesDefStatus) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskResourcesDefStatus) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetMessageList

`func (o *AppTaskResourcesDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppTaskResourcesDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppTaskResourcesDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetProjectReferenceList

`func (o *AppTaskResourcesDefStatus) GetProjectReferenceList() []ProjectReference`

GetProjectReferenceList returns the ProjectReferenceList field if non-nil, zero value otherwise.

### GetProjectReferenceListOk

`func (o *AppTaskResourcesDefStatus) GetProjectReferenceListOk() (*[]ProjectReference, bool)`

GetProjectReferenceListOk returns a tuple with the ProjectReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReferenceList

`func (o *AppTaskResourcesDefStatus) SetProjectReferenceList(v []ProjectReference)`

SetProjectReferenceList sets ProjectReferenceList field to given value.

### HasProjectReferenceList

`func (o *AppTaskResourcesDefStatus) HasProjectReferenceList() bool`

HasProjectReferenceList returns a boolean if a field has been set.

### GetChildTasksLocalReferenceList

`func (o *AppTaskResourcesDefStatus) GetChildTasksLocalReferenceList() []AppTaskReference`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskResourcesDefStatus) GetChildTasksLocalReferenceListOk() (*[]AppTaskReference, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskResourcesDefStatus) SetChildTasksLocalReferenceList(v []AppTaskReference)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskResourcesDefStatus) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetState

`func (o *AppTaskResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppTaskResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppTaskResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.


### GetAttrs

`func (o *AppTaskResourcesDefStatus) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskResourcesDefStatus) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskResourcesDefStatus) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskResourcesDefStatus) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskResourcesDefStatus) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskResourcesDefStatus) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskResourcesDefStatus) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskResourcesDefStatus) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskResourcesDefStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskResourcesDefStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskResourcesDefStatus) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskResourcesDefStatus) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskResourcesDefStatus) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskResourcesDefStatus) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskResourcesDefStatus) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


