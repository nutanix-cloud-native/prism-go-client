# AppActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for action | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Critical** | Pointer to **bool** | action critical flag | [optional] [default to false]
**Attrs** | Pointer to **map[string]interface{}** | action attrs | [optional] 
**Runbook** | Pointer to [**AppRunbookResponse**](AppRunbookResponse.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppActionResponse

`func NewAppActionResponse() *AppActionResponse`

NewAppActionResponse instantiates a new AppActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppActionResponseWithDefaults

`func NewAppActionResponseWithDefaults() *AppActionResponse`

NewAppActionResponseWithDefaults instantiates a new AppActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessageList

`func (o *AppActionResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppActionResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppActionResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppActionResponse) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *AppActionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppActionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppActionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppActionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *AppActionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppActionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppActionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppActionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCritical

`func (o *AppActionResponse) GetCritical() bool`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *AppActionResponse) GetCriticalOk() (*bool, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *AppActionResponse) SetCritical(v bool)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *AppActionResponse) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetAttrs

`func (o *AppActionResponse) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppActionResponse) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppActionResponse) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppActionResponse) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetRunbook

`func (o *AppActionResponse) GetRunbook() AppRunbookResponse`

GetRunbook returns the Runbook field if non-nil, zero value otherwise.

### GetRunbookOk

`func (o *AppActionResponse) GetRunbookOk() (*AppRunbookResponse, bool)`

GetRunbookOk returns a tuple with the Runbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunbook

`func (o *AppActionResponse) SetRunbook(v AppRunbookResponse)`

SetRunbook sets Runbook field to given value.

### HasRunbook

`func (o *AppActionResponse) HasRunbook() bool`

HasRunbook returns a boolean if a field has been set.

### GetType

`func (o *AppActionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppActionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppActionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *AppActionResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppActionResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppActionResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppActionResponse) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


