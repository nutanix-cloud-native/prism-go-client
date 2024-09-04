# AppVariableResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list for variable | 
**Value** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**Attrs** | Pointer to **map[string]interface{}** |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppVariableResponseDownload

`func NewAppVariableResponseDownload(name string, messageList []MessageResource, state string, ) *AppVariableResponseDownload`

NewAppVariableResponseDownload instantiates a new AppVariableResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVariableResponseDownloadWithDefaults

`func NewAppVariableResponseDownloadWithDefaults() *AppVariableResponseDownload`

NewAppVariableResponseDownloadWithDefaults instantiates a new AppVariableResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValType

`func (o *AppVariableResponseDownload) GetValType() string`

GetValType returns the ValType field if non-nil, zero value otherwise.

### GetValTypeOk

`func (o *AppVariableResponseDownload) GetValTypeOk() (*string, bool)`

GetValTypeOk returns a tuple with the ValType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValType

`func (o *AppVariableResponseDownload) SetValType(v string)`

SetValType sets ValType field to given value.

### HasValType

`func (o *AppVariableResponseDownload) HasValType() bool`

HasValType returns a boolean if a field has been set.

### GetDescription

`func (o *AppVariableResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppVariableResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppVariableResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppVariableResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppVariableResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVariableResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVariableResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetMessageList

`func (o *AppVariableResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppVariableResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppVariableResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetValue

`func (o *AppVariableResponseDownload) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppVariableResponseDownload) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppVariableResponseDownload) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppVariableResponseDownload) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *AppVariableResponseDownload) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppVariableResponseDownload) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppVariableResponseDownload) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppVariableResponseDownload) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetState

`func (o *AppVariableResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppVariableResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppVariableResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetAttrs

`func (o *AppVariableResponseDownload) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppVariableResponseDownload) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppVariableResponseDownload) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppVariableResponseDownload) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetEditables

`func (o *AppVariableResponseDownload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppVariableResponseDownload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppVariableResponseDownload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppVariableResponseDownload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetType

`func (o *AppVariableResponseDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppVariableResponseDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppVariableResponseDownload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppVariableResponseDownload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *AppVariableResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppVariableResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppVariableResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppVariableResponseDownload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


