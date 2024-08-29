# AppProfileResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCreateList** | Pointer to [**[]BlueprintDeploymentResponseDownload**](BlueprintDeploymentResponseDownload.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionResponseDownload**](AppActionResponseDownload.md) | List of references to action  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list for blueprint deployment | 
**DependencyList** | [**[]BlueprintDependencyList**](BlueprintDependencyList.md) |  | 
**VariableList** | Pointer to [**[]AppVariableResponseDownload**](AppVariableResponseDownload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppProfileResponseDownload

`func NewAppProfileResponseDownload(name string, state string, messageList []MessageResource, dependencyList []BlueprintDependencyList, ) *AppProfileResponseDownload`

NewAppProfileResponseDownload instantiates a new AppProfileResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProfileResponseDownloadWithDefaults

`func NewAppProfileResponseDownloadWithDefaults() *AppProfileResponseDownload`

NewAppProfileResponseDownloadWithDefaults instantiates a new AppProfileResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCreateList

`func (o *AppProfileResponseDownload) GetDeploymentCreateList() []BlueprintDeploymentResponseDownload`

GetDeploymentCreateList returns the DeploymentCreateList field if non-nil, zero value otherwise.

### GetDeploymentCreateListOk

`func (o *AppProfileResponseDownload) GetDeploymentCreateListOk() (*[]BlueprintDeploymentResponseDownload, bool)`

GetDeploymentCreateListOk returns a tuple with the DeploymentCreateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCreateList

`func (o *AppProfileResponseDownload) SetDeploymentCreateList(v []BlueprintDeploymentResponseDownload)`

SetDeploymentCreateList sets DeploymentCreateList field to given value.

### HasDeploymentCreateList

`func (o *AppProfileResponseDownload) HasDeploymentCreateList() bool`

HasDeploymentCreateList returns a boolean if a field has been set.

### GetDescription

`func (o *AppProfileResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppProfileResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppProfileResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppProfileResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppProfileResponseDownload) GetActionList() []AppActionResponseDownload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppProfileResponseDownload) GetActionListOk() (*[]AppActionResponseDownload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppProfileResponseDownload) SetActionList(v []AppActionResponseDownload)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppProfileResponseDownload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetName

`func (o *AppProfileResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProfileResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProfileResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppProfileResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppProfileResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppProfileResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetEditables

`func (o *AppProfileResponseDownload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppProfileResponseDownload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppProfileResponseDownload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppProfileResponseDownload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetMessageList

`func (o *AppProfileResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppProfileResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppProfileResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetDependencyList

`func (o *AppProfileResponseDownload) GetDependencyList() []BlueprintDependencyList`

GetDependencyList returns the DependencyList field if non-nil, zero value otherwise.

### GetDependencyListOk

`func (o *AppProfileResponseDownload) GetDependencyListOk() (*[]BlueprintDependencyList, bool)`

GetDependencyListOk returns a tuple with the DependencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyList

`func (o *AppProfileResponseDownload) SetDependencyList(v []BlueprintDependencyList)`

SetDependencyList sets DependencyList field to given value.


### GetVariableList

`func (o *AppProfileResponseDownload) GetVariableList() []AppVariableResponseDownload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppProfileResponseDownload) GetVariableListOk() (*[]AppVariableResponseDownload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppProfileResponseDownload) SetVariableList(v []AppVariableResponseDownload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppProfileResponseDownload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppProfileResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppProfileResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppProfileResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppProfileResponseDownload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


