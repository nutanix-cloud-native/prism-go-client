# AppTaskResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list. | 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReferenceUpload**](AppTaskReferenceUpload.md) |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableResponseDownload**](AppVariableResponseDownload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppTaskResponseDownload

`func NewAppTaskResponseDownload(messageList []MessageResource, name string, state string, type_ string, ) *AppTaskResponseDownload`

NewAppTaskResponseDownload instantiates a new AppTaskResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskResponseDownloadWithDefaults

`func NewAppTaskResponseDownloadWithDefaults() *AppTaskResponseDownload`

NewAppTaskResponseDownloadWithDefaults instantiates a new AppTaskResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskResponseDownload) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskResponseDownload) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskResponseDownload) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskResponseDownload) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskResponseDownload) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskResponseDownload) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskResponseDownload) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskResponseDownload) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetDescription

`func (o *AppTaskResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppTaskResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppTaskResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppTaskResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessageList

`func (o *AppTaskResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppTaskResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppTaskResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetChildTasksLocalReferenceList

`func (o *AppTaskResponseDownload) GetChildTasksLocalReferenceList() []AppTaskReferenceUpload`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskResponseDownload) GetChildTasksLocalReferenceListOk() (*[]AppTaskReferenceUpload, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskResponseDownload) SetChildTasksLocalReferenceList(v []AppTaskReferenceUpload)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskResponseDownload) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppTaskResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppTaskResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppTaskResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppTaskResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppTaskResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppTaskResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetAttrs

`func (o *AppTaskResponseDownload) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskResponseDownload) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskResponseDownload) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskResponseDownload) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskResponseDownload) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskResponseDownload) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskResponseDownload) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskResponseDownload) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskResponseDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskResponseDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskResponseDownload) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskResponseDownload) GetVariableList() []AppVariableResponseDownload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskResponseDownload) GetVariableListOk() (*[]AppVariableResponseDownload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskResponseDownload) SetVariableList(v []AppVariableResponseDownload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskResponseDownload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppTaskResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppTaskResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppTaskResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppTaskResponseDownload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


