# BlueprintDownloadResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**ServiceDefinitionList** | Pointer to [**[]AppServiceResponseDownload**](AppServiceResponseDownload.md) | Service definitions for Blueprint. | [optional] 
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateResponseDownload**](AppSubstrateResponseDownload.md) | Substrate definitions for Blueprint. | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialResponseDownload**](AppCredentialResponseDownload.md) | Credential definitions for Blueprint. | [optional] 
**PackageDefinitionList** | Pointer to [**[]AppPackageResponseDownload**](AppPackageResponseDownload.md) | Package definitions for Blueprint. | [optional] 
**AppProfileList** | Pointer to [**[]AppProfileResponseDownload**](AppProfileResponseDownload.md) | App profile definitions for Blueprint. | [optional] 
**PublishedServiceDefinitionList** | Pointer to [**[]AppPublishedServiceResponseDownload**](AppPublishedServiceResponseDownload.md) | Published service definitions for Blueprint. | [optional] 
**DefaultCredentialLocalReference** | Pointer to [**AppCredentialReferenceUpload**](AppCredentialReferenceUpload.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type of blueprint | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for app blueprint | [optional] 
**IsCloned** | Pointer to **bool** | Cloned or original blueprint | [optional] [default to false]

## Methods

### NewBlueprintDownloadResourcesDefStatus

`func NewBlueprintDownloadResourcesDefStatus() *BlueprintDownloadResourcesDefStatus`

NewBlueprintDownloadResourcesDefStatus instantiates a new BlueprintDownloadResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDownloadResourcesDefStatusWithDefaults

`func NewBlueprintDownloadResourcesDefStatusWithDefaults() *BlueprintDownloadResourcesDefStatus`

NewBlueprintDownloadResourcesDefStatusWithDefaults instantiates a new BlueprintDownloadResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAttrs

`func (o *BlueprintDownloadResourcesDefStatus) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *BlueprintDownloadResourcesDefStatus) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *BlueprintDownloadResourcesDefStatus) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *BlueprintDownloadResourcesDefStatus) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) GetServiceDefinitionList() []AppServiceResponseDownload`

GetServiceDefinitionList returns the ServiceDefinitionList field if non-nil, zero value otherwise.

### GetServiceDefinitionListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetServiceDefinitionListOk() (*[]AppServiceResponseDownload, bool)`

GetServiceDefinitionListOk returns a tuple with the ServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) SetServiceDefinitionList(v []AppServiceResponseDownload)`

SetServiceDefinitionList sets ServiceDefinitionList field to given value.

### HasServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) HasServiceDefinitionList() bool`

HasServiceDefinitionList returns a boolean if a field has been set.

### GetSubstrateDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) GetSubstrateDefinitionList() []AppSubstrateResponseDownload`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetSubstrateDefinitionListOk() (*[]AppSubstrateResponseDownload, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) SetSubstrateDefinitionList(v []AppSubstrateResponseDownload)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) GetCredentialDefinitionList() []AppCredentialResponseDownload`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetCredentialDefinitionListOk() (*[]AppCredentialResponseDownload, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) SetCredentialDefinitionList(v []AppCredentialResponseDownload)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.

### GetPackageDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) GetPackageDefinitionList() []AppPackageResponseDownload`

GetPackageDefinitionList returns the PackageDefinitionList field if non-nil, zero value otherwise.

### GetPackageDefinitionListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetPackageDefinitionListOk() (*[]AppPackageResponseDownload, bool)`

GetPackageDefinitionListOk returns a tuple with the PackageDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) SetPackageDefinitionList(v []AppPackageResponseDownload)`

SetPackageDefinitionList sets PackageDefinitionList field to given value.

### HasPackageDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) HasPackageDefinitionList() bool`

HasPackageDefinitionList returns a boolean if a field has been set.

### GetAppProfileList

`func (o *BlueprintDownloadResourcesDefStatus) GetAppProfileList() []AppProfileResponseDownload`

GetAppProfileList returns the AppProfileList field if non-nil, zero value otherwise.

### GetAppProfileListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetAppProfileListOk() (*[]AppProfileResponseDownload, bool)`

GetAppProfileListOk returns a tuple with the AppProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileList

`func (o *BlueprintDownloadResourcesDefStatus) SetAppProfileList(v []AppProfileResponseDownload)`

SetAppProfileList sets AppProfileList field to given value.

### HasAppProfileList

`func (o *BlueprintDownloadResourcesDefStatus) HasAppProfileList() bool`

HasAppProfileList returns a boolean if a field has been set.

### GetPublishedServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) GetPublishedServiceDefinitionList() []AppPublishedServiceResponseDownload`

GetPublishedServiceDefinitionList returns the PublishedServiceDefinitionList field if non-nil, zero value otherwise.

### GetPublishedServiceDefinitionListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetPublishedServiceDefinitionListOk() (*[]AppPublishedServiceResponseDownload, bool)`

GetPublishedServiceDefinitionListOk returns a tuple with the PublishedServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) SetPublishedServiceDefinitionList(v []AppPublishedServiceResponseDownload)`

SetPublishedServiceDefinitionList sets PublishedServiceDefinitionList field to given value.

### HasPublishedServiceDefinitionList

`func (o *BlueprintDownloadResourcesDefStatus) HasPublishedServiceDefinitionList() bool`

HasPublishedServiceDefinitionList returns a boolean if a field has been set.

### GetDefaultCredentialLocalReference

`func (o *BlueprintDownloadResourcesDefStatus) GetDefaultCredentialLocalReference() AppCredentialReferenceUpload`

GetDefaultCredentialLocalReference returns the DefaultCredentialLocalReference field if non-nil, zero value otherwise.

### GetDefaultCredentialLocalReferenceOk

`func (o *BlueprintDownloadResourcesDefStatus) GetDefaultCredentialLocalReferenceOk() (*AppCredentialReferenceUpload, bool)`

GetDefaultCredentialLocalReferenceOk returns a tuple with the DefaultCredentialLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCredentialLocalReference

`func (o *BlueprintDownloadResourcesDefStatus) SetDefaultCredentialLocalReference(v AppCredentialReferenceUpload)`

SetDefaultCredentialLocalReference sets DefaultCredentialLocalReference field to given value.

### HasDefaultCredentialLocalReference

`func (o *BlueprintDownloadResourcesDefStatus) HasDefaultCredentialLocalReference() bool`

HasDefaultCredentialLocalReference returns a boolean if a field has been set.

### GetState

`func (o *BlueprintDownloadResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintDownloadResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintDownloadResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintDownloadResourcesDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *BlueprintDownloadResourcesDefStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintDownloadResourcesDefStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintDownloadResourcesDefStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintDownloadResourcesDefStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessageList

`func (o *BlueprintDownloadResourcesDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *BlueprintDownloadResourcesDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *BlueprintDownloadResourcesDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *BlueprintDownloadResourcesDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetIsCloned

`func (o *BlueprintDownloadResourcesDefStatus) GetIsCloned() bool`

GetIsCloned returns the IsCloned field if non-nil, zero value otherwise.

### GetIsClonedOk

`func (o *BlueprintDownloadResourcesDefStatus) GetIsClonedOk() (*bool, bool)`

GetIsClonedOk returns a tuple with the IsCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloned

`func (o *BlueprintDownloadResourcesDefStatus) SetIsCloned(v bool)`

SetIsCloned sets IsCloned field to given value.

### HasIsCloned

`func (o *BlueprintDownloadResourcesDefStatus) HasIsCloned() bool`

HasIsCloned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


