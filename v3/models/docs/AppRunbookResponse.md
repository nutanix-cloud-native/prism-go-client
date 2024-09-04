# AppRunbookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskDefinitionList** | Pointer to [**[]AppTaskResponse**](AppTaskResponse.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**MainTaskLocalReference** | Pointer to [**AppTaskReference**](AppTaskReference.md) |  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list | 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppRunbookResponse

`func NewAppRunbookResponse(name string, state string, messageList []MessageResource, uUID string, ) *AppRunbookResponse`

NewAppRunbookResponse instantiates a new AppRunbookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunbookResponseWithDefaults

`func NewAppRunbookResponseWithDefaults() *AppRunbookResponse`

NewAppRunbookResponseWithDefaults instantiates a new AppRunbookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskDefinitionList

`func (o *AppRunbookResponse) GetTaskDefinitionList() []AppTaskResponse`

GetTaskDefinitionList returns the TaskDefinitionList field if non-nil, zero value otherwise.

### GetTaskDefinitionListOk

`func (o *AppRunbookResponse) GetTaskDefinitionListOk() (*[]AppTaskResponse, bool)`

GetTaskDefinitionListOk returns a tuple with the TaskDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionList

`func (o *AppRunbookResponse) SetTaskDefinitionList(v []AppTaskResponse)`

SetTaskDefinitionList sets TaskDefinitionList field to given value.

### HasTaskDefinitionList

`func (o *AppRunbookResponse) HasTaskDefinitionList() bool`

HasTaskDefinitionList returns a boolean if a field has been set.

### GetDescription

`func (o *AppRunbookResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRunbookResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRunbookResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppRunbookResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppRunbookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunbookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunbookResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppRunbookResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppRunbookResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppRunbookResponse) SetState(v string)`

SetState sets State field to given value.


### GetMainTaskLocalReference

`func (o *AppRunbookResponse) GetMainTaskLocalReference() AppTaskReference`

GetMainTaskLocalReference returns the MainTaskLocalReference field if non-nil, zero value otherwise.

### GetMainTaskLocalReferenceOk

`func (o *AppRunbookResponse) GetMainTaskLocalReferenceOk() (*AppTaskReference, bool)`

GetMainTaskLocalReferenceOk returns a tuple with the MainTaskLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTaskLocalReference

`func (o *AppRunbookResponse) SetMainTaskLocalReference(v AppTaskReference)`

SetMainTaskLocalReference sets MainTaskLocalReference field to given value.

### HasMainTaskLocalReference

`func (o *AppRunbookResponse) HasMainTaskLocalReference() bool`

HasMainTaskLocalReference returns a boolean if a field has been set.

### GetMessageList

`func (o *AppRunbookResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppRunbookResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppRunbookResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetVariableList

`func (o *AppRunbookResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppRunbookResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppRunbookResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppRunbookResponse) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppRunbookResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppRunbookResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppRunbookResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


