# EnvironmentResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectReference** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateResponse**](AppSubstrateResponse.md) | Substrate definitions for Environment. | [optional] 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) | List of variables | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialResponse**](AppCredentialResponse.md) | Credential definitions for Environment. | [optional] 

## Methods

### NewEnvironmentResourcesDefStatus

`func NewEnvironmentResourcesDefStatus() *EnvironmentResourcesDefStatus`

NewEnvironmentResourcesDefStatus instantiates a new EnvironmentResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourcesDefStatusWithDefaults

`func NewEnvironmentResourcesDefStatusWithDefaults() *EnvironmentResourcesDefStatus`

NewEnvironmentResourcesDefStatusWithDefaults instantiates a new EnvironmentResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectReference

`func (o *EnvironmentResourcesDefStatus) GetProjectReference() ProjectReference`

GetProjectReference returns the ProjectReference field if non-nil, zero value otherwise.

### GetProjectReferenceOk

`func (o *EnvironmentResourcesDefStatus) GetProjectReferenceOk() (*ProjectReference, bool)`

GetProjectReferenceOk returns a tuple with the ProjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReference

`func (o *EnvironmentResourcesDefStatus) SetProjectReference(v ProjectReference)`

SetProjectReference sets ProjectReference field to given value.

### HasProjectReference

`func (o *EnvironmentResourcesDefStatus) HasProjectReference() bool`

HasProjectReference returns a boolean if a field has been set.

### GetSubstrateDefinitionList

`func (o *EnvironmentResourcesDefStatus) GetSubstrateDefinitionList() []AppSubstrateResponse`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *EnvironmentResourcesDefStatus) GetSubstrateDefinitionListOk() (*[]AppSubstrateResponse, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *EnvironmentResourcesDefStatus) SetSubstrateDefinitionList(v []AppSubstrateResponse)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *EnvironmentResourcesDefStatus) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetVariableList

`func (o *EnvironmentResourcesDefStatus) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *EnvironmentResourcesDefStatus) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *EnvironmentResourcesDefStatus) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *EnvironmentResourcesDefStatus) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *EnvironmentResourcesDefStatus) GetCredentialDefinitionList() []AppCredentialResponse`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *EnvironmentResourcesDefStatus) GetCredentialDefinitionListOk() (*[]AppCredentialResponse, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *EnvironmentResourcesDefStatus) SetCredentialDefinitionList(v []AppCredentialResponse)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *EnvironmentResourcesDefStatus) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


