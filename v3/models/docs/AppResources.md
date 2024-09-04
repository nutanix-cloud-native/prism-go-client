# AppResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionList** | [**[]AppActionInput**](AppActionInput.md) | List of Action for Application. | 
**CredentialList** | [**[]AppCredentialInput**](AppCredentialInput.md) | Credential list for appspec | 
**AppBlueprintConfigReference** | Pointer to [**BlueprintReference**](BlueprintReference.md) |  | [optional] 
**AppBlueprintReference** | Pointer to [**BlueprintReference**](BlueprintReference.md) |  | [optional] 
**AppProfileConfigReference** | Pointer to [**AppProfileReference**](AppProfileReference.md) |  | [optional] 
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**DependencyList** | Pointer to [**[]BlueprintDependencyList**](BlueprintDependencyList.md) | Dependencies or edges between callrunbook tasks formed by usage of macros in child tasks | [optional] 
**VariableList** | [**[]AppVariableInput**](AppVariableInput.md) | List of variables | 
**DeploymentList** | [**[]AppDeploymentInput**](AppDeploymentInput.md) | List of Deployment Spec for Application. | 

## Methods

### NewAppResources

`func NewAppResources(actionList []AppActionInput, credentialList []AppCredentialInput, variableList []AppVariableInput, deploymentList []AppDeploymentInput, ) *AppResources`

NewAppResources instantiates a new AppResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResourcesWithDefaults

`func NewAppResourcesWithDefaults() *AppResources`

NewAppResourcesWithDefaults instantiates a new AppResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionList

`func (o *AppResources) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppResources) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppResources) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.


### GetCredentialList

`func (o *AppResources) GetCredentialList() []AppCredentialInput`

GetCredentialList returns the CredentialList field if non-nil, zero value otherwise.

### GetCredentialListOk

`func (o *AppResources) GetCredentialListOk() (*[]AppCredentialInput, bool)`

GetCredentialListOk returns a tuple with the CredentialList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialList

`func (o *AppResources) SetCredentialList(v []AppCredentialInput)`

SetCredentialList sets CredentialList field to given value.


### GetAppBlueprintConfigReference

`func (o *AppResources) GetAppBlueprintConfigReference() BlueprintReference`

GetAppBlueprintConfigReference returns the AppBlueprintConfigReference field if non-nil, zero value otherwise.

### GetAppBlueprintConfigReferenceOk

`func (o *AppResources) GetAppBlueprintConfigReferenceOk() (*BlueprintReference, bool)`

GetAppBlueprintConfigReferenceOk returns a tuple with the AppBlueprintConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintConfigReference

`func (o *AppResources) SetAppBlueprintConfigReference(v BlueprintReference)`

SetAppBlueprintConfigReference sets AppBlueprintConfigReference field to given value.

### HasAppBlueprintConfigReference

`func (o *AppResources) HasAppBlueprintConfigReference() bool`

HasAppBlueprintConfigReference returns a boolean if a field has been set.

### GetAppBlueprintReference

`func (o *AppResources) GetAppBlueprintReference() BlueprintReference`

GetAppBlueprintReference returns the AppBlueprintReference field if non-nil, zero value otherwise.

### GetAppBlueprintReferenceOk

`func (o *AppResources) GetAppBlueprintReferenceOk() (*BlueprintReference, bool)`

GetAppBlueprintReferenceOk returns a tuple with the AppBlueprintReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintReference

`func (o *AppResources) SetAppBlueprintReference(v BlueprintReference)`

SetAppBlueprintReference sets AppBlueprintReference field to given value.

### HasAppBlueprintReference

`func (o *AppResources) HasAppBlueprintReference() bool`

HasAppBlueprintReference returns a boolean if a field has been set.

### GetAppProfileConfigReference

`func (o *AppResources) GetAppProfileConfigReference() AppProfileReference`

GetAppProfileConfigReference returns the AppProfileConfigReference field if non-nil, zero value otherwise.

### GetAppProfileConfigReferenceOk

`func (o *AppResources) GetAppProfileConfigReferenceOk() (*AppProfileReference, bool)`

GetAppProfileConfigReferenceOk returns a tuple with the AppProfileConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileConfigReference

`func (o *AppResources) SetAppProfileConfigReference(v AppProfileReference)`

SetAppProfileConfigReference sets AppProfileConfigReference field to given value.

### HasAppProfileConfigReference

`func (o *AppResources) HasAppProfileConfigReference() bool`

HasAppProfileConfigReference returns a boolean if a field has been set.

### GetClientAttrs

`func (o *AppResources) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *AppResources) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *AppResources) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *AppResources) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetDependencyList

`func (o *AppResources) GetDependencyList() []BlueprintDependencyList`

GetDependencyList returns the DependencyList field if non-nil, zero value otherwise.

### GetDependencyListOk

`func (o *AppResources) GetDependencyListOk() (*[]BlueprintDependencyList, bool)`

GetDependencyListOk returns a tuple with the DependencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyList

`func (o *AppResources) SetDependencyList(v []BlueprintDependencyList)`

SetDependencyList sets DependencyList field to given value.

### HasDependencyList

`func (o *AppResources) HasDependencyList() bool`

HasDependencyList returns a boolean if a field has been set.

### GetVariableList

`func (o *AppResources) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppResources) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppResources) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.


### GetDeploymentList

`func (o *AppResources) GetDeploymentList() []AppDeploymentInput`

GetDeploymentList returns the DeploymentList field if non-nil, zero value otherwise.

### GetDeploymentListOk

`func (o *AppResources) GetDeploymentListOk() (*[]AppDeploymentInput, bool)`

GetDeploymentListOk returns a tuple with the DeploymentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentList

`func (o *AppResources) SetDeploymentList(v []AppDeploymentInput)`

SetDeploymentList sets DeploymentList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


