# AppProfileInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCreateList** | Pointer to [**[]BlueprintDeploymentInputUpload**](BlueprintDeploymentInputUpload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to action  | [optional] 
**Name** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppProfileInputUpload

`func NewAppProfileInputUpload(name string, ) *AppProfileInputUpload`

NewAppProfileInputUpload instantiates a new AppProfileInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProfileInputUploadWithDefaults

`func NewAppProfileInputUploadWithDefaults() *AppProfileInputUpload`

NewAppProfileInputUploadWithDefaults instantiates a new AppProfileInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCreateList

`func (o *AppProfileInputUpload) GetDeploymentCreateList() []BlueprintDeploymentInputUpload`

GetDeploymentCreateList returns the DeploymentCreateList field if non-nil, zero value otherwise.

### GetDeploymentCreateListOk

`func (o *AppProfileInputUpload) GetDeploymentCreateListOk() (*[]BlueprintDeploymentInputUpload, bool)`

GetDeploymentCreateListOk returns a tuple with the DeploymentCreateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCreateList

`func (o *AppProfileInputUpload) SetDeploymentCreateList(v []BlueprintDeploymentInputUpload)`

SetDeploymentCreateList sets DeploymentCreateList field to given value.

### HasDeploymentCreateList

`func (o *AppProfileInputUpload) HasDeploymentCreateList() bool`

HasDeploymentCreateList returns a boolean if a field has been set.

### GetUUID

`func (o *AppProfileInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppProfileInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppProfileInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppProfileInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetActionList

`func (o *AppProfileInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppProfileInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppProfileInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppProfileInputUpload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetName

`func (o *AppProfileInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProfileInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProfileInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetEditables

`func (o *AppProfileInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppProfileInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppProfileInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppProfileInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetVariableList

`func (o *AppProfileInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppProfileInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppProfileInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppProfileInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetDescription

`func (o *AppProfileInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppProfileInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppProfileInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppProfileInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


