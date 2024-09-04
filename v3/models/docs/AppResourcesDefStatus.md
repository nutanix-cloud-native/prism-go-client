# AppResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppState** | **string** |  | 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) | List of Action for Application. | 
**SourceMarketplaceName** | Pointer to **string** | Name of the marketplace item used to create this app | [optional] 
**CredentialList** | [**[]AppCredentialResponse**](AppCredentialResponse.md) | Credential list for appspec | 
**AppBlueprintConfigReference** | Pointer to [**BlueprintReference**](BlueprintReference.md) |  | [optional] 
**AppBlueprintReference** | [**BlueprintReference**](BlueprintReference.md) |  | 
**AppProfileConfigReference** | Pointer to [**AppProfileReference**](AppProfileReference.md) |  | [optional] 
**SourceMarketplaceVersion** | Pointer to **string** | Version of the marketplace item used to create this app | [optional] 
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**DependencyList** | Pointer to [**[]BlueprintDependencyList**](BlueprintDependencyList.md) | Dependencies or edges between callrunbook tasks formed by usage of macros in child tasks | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) | List of variables | 
**DeploymentList** | [**[]AppDeploymentResponse**](AppDeploymentResponse.md) | List of Deployment Spec for Application. | 

## Methods

### NewAppResourcesDefStatus

`func NewAppResourcesDefStatus(appState string, actionList []AppActionResponse, credentialList []AppCredentialResponse, appBlueprintReference BlueprintReference, variableList []AppVariableResponse, deploymentList []AppDeploymentResponse, ) *AppResourcesDefStatus`

NewAppResourcesDefStatus instantiates a new AppResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResourcesDefStatusWithDefaults

`func NewAppResourcesDefStatusWithDefaults() *AppResourcesDefStatus`

NewAppResourcesDefStatusWithDefaults instantiates a new AppResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppState

`func (o *AppResourcesDefStatus) GetAppState() string`

GetAppState returns the AppState field if non-nil, zero value otherwise.

### GetAppStateOk

`func (o *AppResourcesDefStatus) GetAppStateOk() (*string, bool)`

GetAppStateOk returns a tuple with the AppState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppState

`func (o *AppResourcesDefStatus) SetAppState(v string)`

SetAppState sets AppState field to given value.


### GetActionList

`func (o *AppResourcesDefStatus) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppResourcesDefStatus) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppResourcesDefStatus) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetSourceMarketplaceName

`func (o *AppResourcesDefStatus) GetSourceMarketplaceName() string`

GetSourceMarketplaceName returns the SourceMarketplaceName field if non-nil, zero value otherwise.

### GetSourceMarketplaceNameOk

`func (o *AppResourcesDefStatus) GetSourceMarketplaceNameOk() (*string, bool)`

GetSourceMarketplaceNameOk returns a tuple with the SourceMarketplaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMarketplaceName

`func (o *AppResourcesDefStatus) SetSourceMarketplaceName(v string)`

SetSourceMarketplaceName sets SourceMarketplaceName field to given value.

### HasSourceMarketplaceName

`func (o *AppResourcesDefStatus) HasSourceMarketplaceName() bool`

HasSourceMarketplaceName returns a boolean if a field has been set.

### GetCredentialList

`func (o *AppResourcesDefStatus) GetCredentialList() []AppCredentialResponse`

GetCredentialList returns the CredentialList field if non-nil, zero value otherwise.

### GetCredentialListOk

`func (o *AppResourcesDefStatus) GetCredentialListOk() (*[]AppCredentialResponse, bool)`

GetCredentialListOk returns a tuple with the CredentialList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialList

`func (o *AppResourcesDefStatus) SetCredentialList(v []AppCredentialResponse)`

SetCredentialList sets CredentialList field to given value.


### GetAppBlueprintConfigReference

`func (o *AppResourcesDefStatus) GetAppBlueprintConfigReference() BlueprintReference`

GetAppBlueprintConfigReference returns the AppBlueprintConfigReference field if non-nil, zero value otherwise.

### GetAppBlueprintConfigReferenceOk

`func (o *AppResourcesDefStatus) GetAppBlueprintConfigReferenceOk() (*BlueprintReference, bool)`

GetAppBlueprintConfigReferenceOk returns a tuple with the AppBlueprintConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintConfigReference

`func (o *AppResourcesDefStatus) SetAppBlueprintConfigReference(v BlueprintReference)`

SetAppBlueprintConfigReference sets AppBlueprintConfigReference field to given value.

### HasAppBlueprintConfigReference

`func (o *AppResourcesDefStatus) HasAppBlueprintConfigReference() bool`

HasAppBlueprintConfigReference returns a boolean if a field has been set.

### GetAppBlueprintReference

`func (o *AppResourcesDefStatus) GetAppBlueprintReference() BlueprintReference`

GetAppBlueprintReference returns the AppBlueprintReference field if non-nil, zero value otherwise.

### GetAppBlueprintReferenceOk

`func (o *AppResourcesDefStatus) GetAppBlueprintReferenceOk() (*BlueprintReference, bool)`

GetAppBlueprintReferenceOk returns a tuple with the AppBlueprintReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBlueprintReference

`func (o *AppResourcesDefStatus) SetAppBlueprintReference(v BlueprintReference)`

SetAppBlueprintReference sets AppBlueprintReference field to given value.


### GetAppProfileConfigReference

`func (o *AppResourcesDefStatus) GetAppProfileConfigReference() AppProfileReference`

GetAppProfileConfigReference returns the AppProfileConfigReference field if non-nil, zero value otherwise.

### GetAppProfileConfigReferenceOk

`func (o *AppResourcesDefStatus) GetAppProfileConfigReferenceOk() (*AppProfileReference, bool)`

GetAppProfileConfigReferenceOk returns a tuple with the AppProfileConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileConfigReference

`func (o *AppResourcesDefStatus) SetAppProfileConfigReference(v AppProfileReference)`

SetAppProfileConfigReference sets AppProfileConfigReference field to given value.

### HasAppProfileConfigReference

`func (o *AppResourcesDefStatus) HasAppProfileConfigReference() bool`

HasAppProfileConfigReference returns a boolean if a field has been set.

### GetSourceMarketplaceVersion

`func (o *AppResourcesDefStatus) GetSourceMarketplaceVersion() string`

GetSourceMarketplaceVersion returns the SourceMarketplaceVersion field if non-nil, zero value otherwise.

### GetSourceMarketplaceVersionOk

`func (o *AppResourcesDefStatus) GetSourceMarketplaceVersionOk() (*string, bool)`

GetSourceMarketplaceVersionOk returns a tuple with the SourceMarketplaceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMarketplaceVersion

`func (o *AppResourcesDefStatus) SetSourceMarketplaceVersion(v string)`

SetSourceMarketplaceVersion sets SourceMarketplaceVersion field to given value.

### HasSourceMarketplaceVersion

`func (o *AppResourcesDefStatus) HasSourceMarketplaceVersion() bool`

HasSourceMarketplaceVersion returns a boolean if a field has been set.

### GetClientAttrs

`func (o *AppResourcesDefStatus) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *AppResourcesDefStatus) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *AppResourcesDefStatus) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *AppResourcesDefStatus) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetDependencyList

`func (o *AppResourcesDefStatus) GetDependencyList() []BlueprintDependencyList`

GetDependencyList returns the DependencyList field if non-nil, zero value otherwise.

### GetDependencyListOk

`func (o *AppResourcesDefStatus) GetDependencyListOk() (*[]BlueprintDependencyList, bool)`

GetDependencyListOk returns a tuple with the DependencyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyList

`func (o *AppResourcesDefStatus) SetDependencyList(v []BlueprintDependencyList)`

SetDependencyList sets DependencyList field to given value.

### HasDependencyList

`func (o *AppResourcesDefStatus) HasDependencyList() bool`

HasDependencyList returns a boolean if a field has been set.

### GetVariableList

`func (o *AppResourcesDefStatus) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppResourcesDefStatus) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppResourcesDefStatus) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDeploymentList

`func (o *AppResourcesDefStatus) GetDeploymentList() []AppDeploymentResponse`

GetDeploymentList returns the DeploymentList field if non-nil, zero value otherwise.

### GetDeploymentListOk

`func (o *AppResourcesDefStatus) GetDeploymentListOk() (*[]AppDeploymentResponse, bool)`

GetDeploymentListOk returns a tuple with the DeploymentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentList

`func (o *AppResourcesDefStatus) SetDeploymentList(v []AppDeploymentResponse)`

SetDeploymentList sets DeploymentList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


