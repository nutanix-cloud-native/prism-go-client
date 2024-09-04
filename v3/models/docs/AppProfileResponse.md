# AppProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCreateList** | Pointer to [**[]BlueprintDeploymentResponse**](BlueprintDeploymentResponse.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list for blueprint deployment | 
**DependencyList** | [**[]BlueprintDependencyList**](BlueprintDependencyList.md) |  | 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppProfileResponse

`func NewAppProfileResponse(name string, state string, messageList []MessageResource, dependencyList []BlueprintDependencyList, uUID string, ) *AppProfileResponse`

NewAppProfileResponse instantiates a new AppProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProfileResponseWithDefaults

`func NewAppProfileResponseWithDefaults() *AppProfileResponse`

NewAppProfileResponseWithDefaults instantiates a new AppProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCreateList

`func (o *AppProfileResponse) GetDeploymentCreateList() []BlueprintDeploymentResponse`

GetDeploymentCreateList returns the DeploymentCreateList field if non-nil, zero value otherwise.

### GetDeploymentCreateListOk

`func (o *AppProfileResponse) GetDeploymentCreateListOk() (*[]BlueprintDeploymentResponse, bool)`

GetDeploymentCreateListOk returns a tuple with the DeploymentCreateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCreateList

`func (o *AppProfileResponse) SetDeploymentCreateList(v []BlueprintDeploymentResponse)`

SetDeploymentCreateList sets DeploymentCreateList field to given value.

### HasDeploymentCreateList

`func (o *AppProfileResponse) HasDeploymentCreateList() bool`

HasDeploymentCreateList returns a boolean if a field has been set.

### GetDescription

`func (o *AppProfileResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppProfileResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppProfileResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppProfileResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppProfileResponse) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppProfileResponse) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppProfileResponse) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppProfileResponse) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetName

`func (o *AppProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppProfileResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppProfileResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppProfileResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppProfileResponse) SetState(v string)`

SetState sets State field to given value.


### GetEditables

`func (o *AppProfileResponse) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppProfileResponse) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppProfileResponse) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppProfileResponse) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetMessageList

`func (o *AppProfileResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppProfileResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppProfileResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetDependencyList

`func (o *AppProfileResponse) GetDependencyList() []BlueprintDependencyList`

GetDependencyList returns the DependencyList field if non-nil, zero value otherwise.

### GetDependencyListOk

`func (o *AppProfileResponse) GetDependencyListOk() (*[]BlueprintDependencyList, bool)`

GetDependencyListOk returns a tuple with the DependencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyList

`func (o *AppProfileResponse) SetDependencyList(v []BlueprintDependencyList)`

SetDependencyList sets DependencyList field to given value.


### GetVariableList

`func (o *AppProfileResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppProfileResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppProfileResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppProfileResponse) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppProfileResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppProfileResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppProfileResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


