# AppProfileInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCreateList** | Pointer to [**[]BlueprintDeploymentInput**](BlueprintDeploymentInput.md) |  | [optional] 
**UUID** | **string** |  | 
**ActionList** | Pointer to [**[]AppActionInput**](AppActionInput.md) | List of references to action  | [optional] 
**Name** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppProfileInput

`func NewAppProfileInput(uUID string, name string, ) *AppProfileInput`

NewAppProfileInput instantiates a new AppProfileInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProfileInputWithDefaults

`func NewAppProfileInputWithDefaults() *AppProfileInput`

NewAppProfileInputWithDefaults instantiates a new AppProfileInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCreateList

`func (o *AppProfileInput) GetDeploymentCreateList() []BlueprintDeploymentInput`

GetDeploymentCreateList returns the DeploymentCreateList field if non-nil, zero value otherwise.

### GetDeploymentCreateListOk

`func (o *AppProfileInput) GetDeploymentCreateListOk() (*[]BlueprintDeploymentInput, bool)`

GetDeploymentCreateListOk returns a tuple with the DeploymentCreateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCreateList

`func (o *AppProfileInput) SetDeploymentCreateList(v []BlueprintDeploymentInput)`

SetDeploymentCreateList sets DeploymentCreateList field to given value.

### HasDeploymentCreateList

`func (o *AppProfileInput) HasDeploymentCreateList() bool`

HasDeploymentCreateList returns a boolean if a field has been set.

### GetUUID

`func (o *AppProfileInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppProfileInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppProfileInput) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetActionList

`func (o *AppProfileInput) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppProfileInput) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppProfileInput) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppProfileInput) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetName

`func (o *AppProfileInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProfileInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProfileInput) SetName(v string)`

SetName sets Name field to given value.


### GetEditables

`func (o *AppProfileInput) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppProfileInput) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppProfileInput) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppProfileInput) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetVariableList

`func (o *AppProfileInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppProfileInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppProfileInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppProfileInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetDescription

`func (o *AppProfileInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppProfileInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppProfileInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppProfileInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


