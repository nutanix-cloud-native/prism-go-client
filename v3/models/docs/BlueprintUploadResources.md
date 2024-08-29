# BlueprintUploadResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**ServiceDefinitionList** | Pointer to [**[]AppServiceInputUpload**](AppServiceInputUpload.md) | Service definitions for Blueprint. | [optional] 
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateInputUpload**](AppSubstrateInputUpload.md) | Substrate definitions for Blueprint. | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialInputUpload**](AppCredentialInputUpload.md) | Credential definitions for Blueprint. | [optional] 
**PackageDefinitionList** | Pointer to [**[]AppPackageInputUpload**](AppPackageInputUpload.md) | Package definitions for Blueprint. | [optional] 
**AppProfileList** | Pointer to [**[]AppProfileInputUpload**](AppProfileInputUpload.md) | App profile definitions for Blueprint. | [optional] 
**PublishedServiceDefinitionList** | Pointer to [**[]AppPublishedServiceInputUpload**](AppPublishedServiceInputUpload.md) | Published service definitions for Blueprint. | [optional] 
**DefaultCredentialLocalReference** | Pointer to [**AppCredentialReferenceUpload**](AppCredentialReferenceUpload.md) |  | [optional] 
**Type** | Pointer to **string** | Type of blueprint | [optional] 

## Methods

### NewBlueprintUploadResources

`func NewBlueprintUploadResources() *BlueprintUploadResources`

NewBlueprintUploadResources instantiates a new BlueprintUploadResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUploadResourcesWithDefaults

`func NewBlueprintUploadResourcesWithDefaults() *BlueprintUploadResources`

NewBlueprintUploadResourcesWithDefaults instantiates a new BlueprintUploadResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAttrs

`func (o *BlueprintUploadResources) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *BlueprintUploadResources) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *BlueprintUploadResources) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *BlueprintUploadResources) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetServiceDefinitionList

`func (o *BlueprintUploadResources) GetServiceDefinitionList() []AppServiceInputUpload`

GetServiceDefinitionList returns the ServiceDefinitionList field if non-nil, zero value otherwise.

### GetServiceDefinitionListOk

`func (o *BlueprintUploadResources) GetServiceDefinitionListOk() (*[]AppServiceInputUpload, bool)`

GetServiceDefinitionListOk returns a tuple with the ServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefinitionList

`func (o *BlueprintUploadResources) SetServiceDefinitionList(v []AppServiceInputUpload)`

SetServiceDefinitionList sets ServiceDefinitionList field to given value.

### HasServiceDefinitionList

`func (o *BlueprintUploadResources) HasServiceDefinitionList() bool`

HasServiceDefinitionList returns a boolean if a field has been set.

### GetSubstrateDefinitionList

`func (o *BlueprintUploadResources) GetSubstrateDefinitionList() []AppSubstrateInputUpload`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *BlueprintUploadResources) GetSubstrateDefinitionListOk() (*[]AppSubstrateInputUpload, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *BlueprintUploadResources) SetSubstrateDefinitionList(v []AppSubstrateInputUpload)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *BlueprintUploadResources) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *BlueprintUploadResources) GetCredentialDefinitionList() []AppCredentialInputUpload`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *BlueprintUploadResources) GetCredentialDefinitionListOk() (*[]AppCredentialInputUpload, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *BlueprintUploadResources) SetCredentialDefinitionList(v []AppCredentialInputUpload)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *BlueprintUploadResources) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.

### GetPackageDefinitionList

`func (o *BlueprintUploadResources) GetPackageDefinitionList() []AppPackageInputUpload`

GetPackageDefinitionList returns the PackageDefinitionList field if non-nil, zero value otherwise.

### GetPackageDefinitionListOk

`func (o *BlueprintUploadResources) GetPackageDefinitionListOk() (*[]AppPackageInputUpload, bool)`

GetPackageDefinitionListOk returns a tuple with the PackageDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDefinitionList

`func (o *BlueprintUploadResources) SetPackageDefinitionList(v []AppPackageInputUpload)`

SetPackageDefinitionList sets PackageDefinitionList field to given value.

### HasPackageDefinitionList

`func (o *BlueprintUploadResources) HasPackageDefinitionList() bool`

HasPackageDefinitionList returns a boolean if a field has been set.

### GetAppProfileList

`func (o *BlueprintUploadResources) GetAppProfileList() []AppProfileInputUpload`

GetAppProfileList returns the AppProfileList field if non-nil, zero value otherwise.

### GetAppProfileListOk

`func (o *BlueprintUploadResources) GetAppProfileListOk() (*[]AppProfileInputUpload, bool)`

GetAppProfileListOk returns a tuple with the AppProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileList

`func (o *BlueprintUploadResources) SetAppProfileList(v []AppProfileInputUpload)`

SetAppProfileList sets AppProfileList field to given value.

### HasAppProfileList

`func (o *BlueprintUploadResources) HasAppProfileList() bool`

HasAppProfileList returns a boolean if a field has been set.

### GetPublishedServiceDefinitionList

`func (o *BlueprintUploadResources) GetPublishedServiceDefinitionList() []AppPublishedServiceInputUpload`

GetPublishedServiceDefinitionList returns the PublishedServiceDefinitionList field if non-nil, zero value otherwise.

### GetPublishedServiceDefinitionListOk

`func (o *BlueprintUploadResources) GetPublishedServiceDefinitionListOk() (*[]AppPublishedServiceInputUpload, bool)`

GetPublishedServiceDefinitionListOk returns a tuple with the PublishedServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceDefinitionList

`func (o *BlueprintUploadResources) SetPublishedServiceDefinitionList(v []AppPublishedServiceInputUpload)`

SetPublishedServiceDefinitionList sets PublishedServiceDefinitionList field to given value.

### HasPublishedServiceDefinitionList

`func (o *BlueprintUploadResources) HasPublishedServiceDefinitionList() bool`

HasPublishedServiceDefinitionList returns a boolean if a field has been set.

### GetDefaultCredentialLocalReference

`func (o *BlueprintUploadResources) GetDefaultCredentialLocalReference() AppCredentialReferenceUpload`

GetDefaultCredentialLocalReference returns the DefaultCredentialLocalReference field if non-nil, zero value otherwise.

### GetDefaultCredentialLocalReferenceOk

`func (o *BlueprintUploadResources) GetDefaultCredentialLocalReferenceOk() (*AppCredentialReferenceUpload, bool)`

GetDefaultCredentialLocalReferenceOk returns a tuple with the DefaultCredentialLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCredentialLocalReference

`func (o *BlueprintUploadResources) SetDefaultCredentialLocalReference(v AppCredentialReferenceUpload)`

SetDefaultCredentialLocalReference sets DefaultCredentialLocalReference field to given value.

### HasDefaultCredentialLocalReference

`func (o *BlueprintUploadResources) HasDefaultCredentialLocalReference() bool`

HasDefaultCredentialLocalReference returns a boolean if a field has been set.

### GetType

`func (o *BlueprintUploadResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintUploadResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintUploadResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintUploadResources) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


