# AppRunbookResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskDefinitionList** | Pointer to [**[]AppTaskResponseDownload**](AppTaskResponseDownload.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**MainTaskLocalReference** | Pointer to [**AppTaskReferenceUpload**](AppTaskReferenceUpload.md) |  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list | 
**VariableList** | Pointer to [**[]AppVariableResponseDownload**](AppVariableResponseDownload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppRunbookResponseDownload

`func NewAppRunbookResponseDownload(name string, state string, messageList []MessageResource, ) *AppRunbookResponseDownload`

NewAppRunbookResponseDownload instantiates a new AppRunbookResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunbookResponseDownloadWithDefaults

`func NewAppRunbookResponseDownloadWithDefaults() *AppRunbookResponseDownload`

NewAppRunbookResponseDownloadWithDefaults instantiates a new AppRunbookResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskDefinitionList

`func (o *AppRunbookResponseDownload) GetTaskDefinitionList() []AppTaskResponseDownload`

GetTaskDefinitionList returns the TaskDefinitionList field if non-nil, zero value otherwise.

### GetTaskDefinitionListOk

`func (o *AppRunbookResponseDownload) GetTaskDefinitionListOk() (*[]AppTaskResponseDownload, bool)`

GetTaskDefinitionListOk returns a tuple with the TaskDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionList

`func (o *AppRunbookResponseDownload) SetTaskDefinitionList(v []AppTaskResponseDownload)`

SetTaskDefinitionList sets TaskDefinitionList field to given value.

### HasTaskDefinitionList

`func (o *AppRunbookResponseDownload) HasTaskDefinitionList() bool`

HasTaskDefinitionList returns a boolean if a field has been set.

### GetDescription

`func (o *AppRunbookResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRunbookResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRunbookResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRunbookResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppRunbookResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunbookResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunbookResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppRunbookResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppRunbookResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppRunbookResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetMainTaskLocalReference

`func (o *AppRunbookResponseDownload) GetMainTaskLocalReference() AppTaskReferenceUpload`

GetMainTaskLocalReference returns the MainTaskLocalReference field if non-nil, zero value otherwise.

### GetMainTaskLocalReferenceOk

`func (o *AppRunbookResponseDownload) GetMainTaskLocalReferenceOk() (*AppTaskReferenceUpload, bool)`

GetMainTaskLocalReferenceOk returns a tuple with the MainTaskLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTaskLocalReference

`func (o *AppRunbookResponseDownload) SetMainTaskLocalReference(v AppTaskReferenceUpload)`

SetMainTaskLocalReference sets MainTaskLocalReference field to given value.

### HasMainTaskLocalReference

`func (o *AppRunbookResponseDownload) HasMainTaskLocalReference() bool`

HasMainTaskLocalReference returns a boolean if a field has been set.

### GetMessageList

`func (o *AppRunbookResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppRunbookResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppRunbookResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetVariableList

`func (o *AppRunbookResponseDownload) GetVariableList() []AppVariableResponseDownload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppRunbookResponseDownload) GetVariableListOk() (*[]AppVariableResponseDownload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppRunbookResponseDownload) SetVariableList(v []AppVariableResponseDownload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppRunbookResponseDownload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppRunbookResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppRunbookResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppRunbookResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppRunbookResponseDownload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


