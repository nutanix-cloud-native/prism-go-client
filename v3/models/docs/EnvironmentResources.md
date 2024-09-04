# EnvironmentResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateInput**](AppSubstrateInput.md) | Substrate definitions for Environment. | [optional] 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) | List of variables | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialInput**](AppCredentialInput.md) | Credential definitions for Environment. | [optional] 

## Methods

### NewEnvironmentResources

`func NewEnvironmentResources() *EnvironmentResources`

NewEnvironmentResources instantiates a new EnvironmentResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourcesWithDefaults

`func NewEnvironmentResourcesWithDefaults() *EnvironmentResources`

NewEnvironmentResourcesWithDefaults instantiates a new EnvironmentResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubstrateDefinitionList

`func (o *EnvironmentResources) GetSubstrateDefinitionList() []AppSubstrateInput`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *EnvironmentResources) GetSubstrateDefinitionListOk() (*[]AppSubstrateInput, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *EnvironmentResources) SetSubstrateDefinitionList(v []AppSubstrateInput)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *EnvironmentResources) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetVariableList

`func (o *EnvironmentResources) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *EnvironmentResources) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *EnvironmentResources) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *EnvironmentResources) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *EnvironmentResources) GetCredentialDefinitionList() []AppCredentialInput`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *EnvironmentResources) GetCredentialDefinitionListOk() (*[]AppCredentialInput, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *EnvironmentResources) SetCredentialDefinitionList(v []AppCredentialInput)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *EnvironmentResources) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


