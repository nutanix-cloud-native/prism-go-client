# AppVariableInputUpload

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
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppVariableInputUpload

`func NewAppVariableInputUpload(name string, ) *AppVariableInputUpload`

NewAppVariableInputUpload instantiates a new AppVariableInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppVariableInputUploadWithDefaults

`func NewAppVariableInputUploadWithDefaults() *AppVariableInputUpload`

NewAppVariableInputUploadWithDefaults instantiates a new AppVariableInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValType

`func (o *AppVariableInputUpload) GetValType() string`

GetValType returns the ValType field if non-nil, zero value otherwise.

### GetValTypeOk

`func (o *AppVariableInputUpload) GetValTypeOk() (*string, bool)`

GetValTypeOk returns a tuple with the ValType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValType

`func (o *AppVariableInputUpload) SetValType(v string)`

SetValType sets ValType field to given value.

### HasValType

`func (o *AppVariableInputUpload) HasValType() bool`

HasValType returns a boolean if a field has been set.

### GetDescription

`func (o *AppVariableInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppVariableInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppVariableInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppVariableInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppVariableInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppVariableInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppVariableInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AppVariableInputUpload) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppVariableInputUpload) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppVariableInputUpload) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppVariableInputUpload) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *AppVariableInputUpload) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppVariableInputUpload) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppVariableInputUpload) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppVariableInputUpload) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAttrs

`func (o *AppVariableInputUpload) GetAttrs() map[string]interface{}`

GetAttrs returns the Attrs field if non-nil, zero value otherwise.

### GetAttrsOk

`func (o *AppVariableInputUpload) GetAttrsOk() (*map[string]interface{}, bool)`

GetAttrsOk returns a tuple with the Attrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrs

`func (o *AppVariableInputUpload) SetAttrs(v map[string]interface{})`

SetAttrs sets Attrs field to given value.

### HasAttrs

`func (o *AppVariableInputUpload) HasAttrs() bool`

HasAttrs returns a boolean if a field has been set.

### GetEditables

`func (o *AppVariableInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppVariableInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppVariableInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppVariableInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetType

`func (o *AppVariableInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppVariableInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppVariableInputUpload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppVariableInputUpload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *AppVariableInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppVariableInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppVariableInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppVariableInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


