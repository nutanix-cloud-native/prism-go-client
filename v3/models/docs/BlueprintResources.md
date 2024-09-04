# BlueprintResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**ServiceDefinitionList** | Pointer to [**[]AppServiceInput**](AppServiceInput.md) | Service definitions for Blueprint. | [optional] 
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateInput**](AppSubstrateInput.md) | Substrate definitions for Blueprint. | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialInput**](AppCredentialInput.md) | Credential definitions for Blueprint. | [optional] 
**PackageDefinitionList** | Pointer to [**[]AppPackageInput**](AppPackageInput.md) | Package definitions for Blueprint. | [optional] 
**AppProfileList** | Pointer to [**[]AppProfileInput**](AppProfileInput.md) | App profile definitions for Blueprint. | [optional] 
**PublishedServiceDefinitionList** | Pointer to [**[]AppPublishedServiceInput**](AppPublishedServiceInput.md) | Published service definitions for Blueprint. | [optional] 
**DefaultCredentialLocalReference** | Pointer to [**AppCredentialReference**](AppCredentialReference.md) |  | [optional] 
**Type** | Pointer to **string** | Type of blueprint | [optional] 

## Methods

### NewBlueprintResources

`func NewBlueprintResources() *BlueprintResources`

NewBlueprintResources instantiates a new BlueprintResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintResourcesWithDefaults

`func NewBlueprintResourcesWithDefaults() *BlueprintResources`

NewBlueprintResourcesWithDefaults instantiates a new BlueprintResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAttrs

`func (o *BlueprintResources) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *BlueprintResources) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *BlueprintResources) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *BlueprintResources) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetServiceDefinitionList

`func (o *BlueprintResources) GetServiceDefinitionList() []AppServiceInput`

GetServiceDefinitionList returns the ServiceDefinitionList field if non-nil, zero value otherwise.

### GetServiceDefinitionListOk

`func (o *BlueprintResources) GetServiceDefinitionListOk() (*[]AppServiceInput, bool)`

GetServiceDefinitionListOk returns a tuple with the ServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefinitionList

`func (o *BlueprintResources) SetServiceDefinitionList(v []AppServiceInput)`

SetServiceDefinitionList sets ServiceDefinitionList field to given value.

### HasServiceDefinitionList

`func (o *BlueprintResources) HasServiceDefinitionList() bool`

HasServiceDefinitionList returns a boolean if a field has been set.

### GetSubstrateDefinitionList

`func (o *BlueprintResources) GetSubstrateDefinitionList() []AppSubstrateInput`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *BlueprintResources) GetSubstrateDefinitionListOk() (*[]AppSubstrateInput, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *BlueprintResources) SetSubstrateDefinitionList(v []AppSubstrateInput)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *BlueprintResources) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *BlueprintResources) GetCredentialDefinitionList() []AppCredentialInput`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *BlueprintResources) GetCredentialDefinitionListOk() (*[]AppCredentialInput, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *BlueprintResources) SetCredentialDefinitionList(v []AppCredentialInput)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *BlueprintResources) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.

### GetPackageDefinitionList

`func (o *BlueprintResources) GetPackageDefinitionList() []AppPackageInput`

GetPackageDefinitionList returns the PackageDefinitionList field if non-nil, zero value otherwise.

### GetPackageDefinitionListOk

`func (o *BlueprintResources) GetPackageDefinitionListOk() (*[]AppPackageInput, bool)`

GetPackageDefinitionListOk returns a tuple with the PackageDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDefinitionList

`func (o *BlueprintResources) SetPackageDefinitionList(v []AppPackageInput)`

SetPackageDefinitionList sets PackageDefinitionList field to given value.

### HasPackageDefinitionList

`func (o *BlueprintResources) HasPackageDefinitionList() bool`

HasPackageDefinitionList returns a boolean if a field has been set.

### GetAppProfileList

`func (o *BlueprintResources) GetAppProfileList() []AppProfileInput`

GetAppProfileList returns the AppProfileList field if non-nil, zero value otherwise.

### GetAppProfileListOk

`func (o *BlueprintResources) GetAppProfileListOk() (*[]AppProfileInput, bool)`

GetAppProfileListOk returns a tuple with the AppProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileList

`func (o *BlueprintResources) SetAppProfileList(v []AppProfileInput)`

SetAppProfileList sets AppProfileList field to given value.

### HasAppProfileList

`func (o *BlueprintResources) HasAppProfileList() bool`

HasAppProfileList returns a boolean if a field has been set.

### GetPublishedServiceDefinitionList

`func (o *BlueprintResources) GetPublishedServiceDefinitionList() []AppPublishedServiceInput`

GetPublishedServiceDefinitionList returns the PublishedServiceDefinitionList field if non-nil, zero value otherwise.

### GetPublishedServiceDefinitionListOk

`func (o *BlueprintResources) GetPublishedServiceDefinitionListOk() (*[]AppPublishedServiceInput, bool)`

GetPublishedServiceDefinitionListOk returns a tuple with the PublishedServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceDefinitionList

`func (o *BlueprintResources) SetPublishedServiceDefinitionList(v []AppPublishedServiceInput)`

SetPublishedServiceDefinitionList sets PublishedServiceDefinitionList field to given value.

### HasPublishedServiceDefinitionList

`func (o *BlueprintResources) HasPublishedServiceDefinitionList() bool`

HasPublishedServiceDefinitionList returns a boolean if a field has been set.

### GetDefaultCredentialLocalReference

`func (o *BlueprintResources) GetDefaultCredentialLocalReference() AppCredentialReference`

GetDefaultCredentialLocalReference returns the DefaultCredentialLocalReference field if non-nil, zero value otherwise.

### GetDefaultCredentialLocalReferenceOk

`func (o *BlueprintResources) GetDefaultCredentialLocalReferenceOk() (*AppCredentialReference, bool)`

GetDefaultCredentialLocalReferenceOk returns a tuple with the DefaultCredentialLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCredentialLocalReference

`func (o *BlueprintResources) SetDefaultCredentialLocalReference(v AppCredentialReference)`

SetDefaultCredentialLocalReference sets DefaultCredentialLocalReference field to given value.

### HasDefaultCredentialLocalReference

`func (o *BlueprintResources) HasDefaultCredentialLocalReference() bool`

HasDefaultCredentialLocalReference returns a boolean if a field has been set.

### GetType

`func (o *BlueprintResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintResources) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


