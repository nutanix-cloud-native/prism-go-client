# AppTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAnyLocalReference** | Pointer to [**EntityReference**](EntityReference.md) |  | [optional] 
**Retries** | Pointer to **string** | Number of retries for the task. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list. | 
**ChildTasksLocalReferenceList** | Pointer to [**[]AppTaskReference**](AppTaskReference.md) |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** | Task attrs for application of type object. | [optional] 
**TimeoutSecs** | Pointer to **string** | task timeout. | [optional] 
**Type** | **string** |  | 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppTaskResponse

`func NewAppTaskResponse(messageList []MessageResource, name string, state string, type_ string, uUID string, ) *AppTaskResponse`

NewAppTaskResponse instantiates a new AppTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskResponseWithDefaults

`func NewAppTaskResponseWithDefaults() *AppTaskResponse`

NewAppTaskResponseWithDefaults instantiates a new AppTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAnyLocalReference

`func (o *AppTaskResponse) GetTargetAnyLocalReference() EntityReference`

GetTargetAnyLocalReference returns the TargetAnyLocalReference field if non-nil, zero value otherwise.

### GetTargetAnyLocalReferenceOk

`func (o *AppTaskResponse) GetTargetAnyLocalReferenceOk() (*EntityReference, bool)`

GetTargetAnyLocalReferenceOk returns a tuple with the TargetAnyLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAnyLocalReference

`func (o *AppTaskResponse) SetTargetAnyLocalReference(v EntityReference)`

SetTargetAnyLocalReference sets TargetAnyLocalReference field to given value.

### HasTargetAnyLocalReference

`func (o *AppTaskResponse) HasTargetAnyLocalReference() bool`

HasTargetAnyLocalReference returns a boolean if a field has been set.

### GetRetries

`func (o *AppTaskResponse) GetRetries() string`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *AppTaskResponse) GetRetriesOk() (*string, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *AppTaskResponse) SetRetries(v string)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *AppTaskResponse) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetDescription

`func (o *AppTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessageList

`func (o *AppTaskResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppTaskResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppTaskResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetChildTasksLocalReferenceList

`func (o *AppTaskResponse) GetChildTasksLocalReferenceList() []AppTaskReference`

GetChildTasksLocalReferenceList returns the ChildTasksLocalReferenceList field if non-nil, zero value otherwise.

### GetChildTasksLocalReferenceListOk

`func (o *AppTaskResponse) GetChildTasksLocalReferenceListOk() (*[]AppTaskReference, bool)`

GetChildTasksLocalReferenceListOk returns a tuple with the ChildTasksLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTasksLocalReferenceList

`func (o *AppTaskResponse) SetChildTasksLocalReferenceList(v []AppTaskReference)`

SetChildTasksLocalReferenceList sets ChildTasksLocalReferenceList field to given value.

### HasChildTasksLocalReferenceList

`func (o *AppTaskResponse) HasChildTasksLocalReferenceList() bool`

HasChildTasksLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppTaskResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppTaskResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppTaskResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppTaskResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppTaskResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppTaskResponse) SetState(v string)`

SetState sets State field to given value.


### GetAttrs

`func (o *AppTaskResponse) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppTaskResponse) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppTaskResponse) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppTaskResponse) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *AppTaskResponse) GetTimeoutSecs() string`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *AppTaskResponse) GetTimeoutSecsOk() (*string, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *AppTaskResponse) SetTimeoutSecs(v string)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *AppTaskResponse) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetType

`func (o *AppTaskResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTaskResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTaskResponse) SetType(v string)`

SetType sets Type field to given value.


### GetVariableList

`func (o *AppTaskResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppTaskResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppTaskResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppTaskResponse) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppTaskResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppTaskResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppTaskResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


