# AppVariableInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Attrs** | Pointer to **map[string]interface{}** |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppVariableInput

`func NewAppVariableInput(name string, uUID string, ) *AppVariableInput`

NewAppVariableInput instantiates a new AppVariableInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVariableInputWithDefaults

`func NewAppVariableInputWithDefaults() *AppVariableInput`

NewAppVariableInputWithDefaults instantiates a new AppVariableInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValType

`func (o *AppVariableInput) GetValType() string`

GetValType returns the ValType field if non-nil, zero value otherwise.

### GetValTypeOk

`func (o *AppVariableInput) GetValTypeOk() (*string, bool)`

GetValTypeOk returns a tuple with the ValType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValType

`func (o *AppVariableInput) SetValType(v string)`

SetValType sets ValType field to given value.

### HasValType

`func (o *AppVariableInput) HasValType() bool`

HasValType returns a boolean if a field has been set.

### GetDescription

`func (o *AppVariableInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppVariableInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppVariableInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppVariableInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppVariableInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVariableInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVariableInput) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppVariableInput) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppVariableInput) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppVariableInput) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppVariableInput) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *AppVariableInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppVariableInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppVariableInput) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppVariableInput) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAttrs

`func (o *AppVariableInput) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppVariableInput) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppVariableInput) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppVariableInput) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetEditables

`func (o *AppVariableInput) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppVariableInput) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppVariableInput) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppVariableInput) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetType

`func (o *AppVariableInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppVariableInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppVariableInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppVariableInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *AppVariableInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppVariableInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppVariableInput) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


