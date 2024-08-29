# AppActionResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for action | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Critical** | Pointer to **bool** | action critical flag | [optional] [default to false]
**Attrs** | Pointer to **map[string]interface{}** | action attrs | [optional] 
**Runbook** | Pointer to [**AppRunbookResponseDownload**](AppRunbookResponseDownload.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppActionResponseDownload

`func NewAppActionResponseDownload() *AppActionResponseDownload`

NewAppActionResponseDownload instantiates a new AppActionResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppActionResponseDownloadWithDefaults

`func NewAppActionResponseDownloadWithDefaults() *AppActionResponseDownload`

NewAppActionResponseDownloadWithDefaults instantiates a new AppActionResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppActionResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppActionResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppActionResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppActionResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessageList

`func (o *AppActionResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppActionResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppActionResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppActionResponseDownload) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *AppActionResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppActionResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppActionResponseDownload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppActionResponseDownload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *AppActionResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppActionResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppActionResponseDownload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppActionResponseDownload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCritical

`func (o *AppActionResponseDownload) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *AppActionResponseDownload) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *AppActionResponseDownload) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *AppActionResponseDownload) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetAttrs

`func (o *AppActionResponseDownload) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppActionResponseDownload) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppActionResponseDownload) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppActionResponseDownload) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetRunbook

`func (o *AppActionResponseDownload) GetRunbook() AppRunbookResponseDownload`

GetRunbook returns the Runbook field if non-nil, zero value otherwise.

### GetRunbookOk

`func (o *AppActionResponseDownload) GetRunbookOk() (*AppRunbookResponseDownload, bool)`

GetRunbookOk returns a tuple with the Runbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunbook

`func (o *AppActionResponseDownload) SetRunbook(v AppRunbookResponseDownload)`

SetRunbook sets Runbook field to given value.

### HasRunbook

`func (o *AppActionResponseDownload) HasRunbook() bool`

HasRunbook returns a boolean if a field has been set.

### GetType

`func (o *AppActionResponseDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppActionResponseDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppActionResponseDownload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppActionResponseDownload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *AppActionResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppActionResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppActionResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppActionResponseDownload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


