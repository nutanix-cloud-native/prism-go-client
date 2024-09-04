# BlueprintResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAttrs** | Pointer to **map[string]interface{}** | Data needed for clients. | [optional] 
**ServiceDefinitionList** | Pointer to [**[]AppServiceResponse**](AppServiceResponse.md) | Service definitions for Blueprint. | [optional] 
**SubstrateDefinitionList** | Pointer to [**[]AppSubstrateResponse**](AppSubstrateResponse.md) | Substrate definitions for Blueprint. | [optional] 
**CredentialDefinitionList** | Pointer to [**[]AppCredentialResponse**](AppCredentialResponse.md) | Credential definitions for Blueprint. | [optional] 
**PackageDefinitionList** | Pointer to [**[]AppPackageResponse**](AppPackageResponse.md) | Package definitions for Blueprint. | [optional] 
**AppProfileList** | Pointer to [**[]AppProfileResponse**](AppProfileResponse.md) | App profile definitions for Blueprint. | [optional] 
**PublishedServiceDefinitionList** | Pointer to [**[]AppPublishedServiceResponse**](AppPublishedServiceResponse.md) | Published service definitions for Blueprint. | [optional] 
**DefaultCredentialLocalReference** | Pointer to [**AppCredentialReference**](AppCredentialReference.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type of blueprint | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for app blueprint | [optional] 
**IsCloned** | Pointer to **bool** | Cloned or original blueprint | [optional] [default to false]

## Methods

### NewBlueprintResourcesDefStatus

`func NewBlueprintResourcesDefStatus() *BlueprintResourcesDefStatus`

NewBlueprintResourcesDefStatus instantiates a new BlueprintResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintResourcesDefStatusWithDefaults

`func NewBlueprintResourcesDefStatusWithDefaults() *BlueprintResourcesDefStatus`

NewBlueprintResourcesDefStatusWithDefaults instantiates a new BlueprintResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAttrs

`func (o *BlueprintResourcesDefStatus) GetClientAttrs() map[string]interface{}`

GetClientAttrs returns the ClientAttrs field if non-nil, zero value otherwise.

### GetClientAttrsOk

`func (o *BlueprintResourcesDefStatus) GetClientAttrsOk() (*map[string]interface{}, bool)`

GetClientAttrsOk returns a tuple with the ClientAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttrs

`func (o *BlueprintResourcesDefStatus) SetClientAttrs(v map[string]interface{})`

SetClientAttrs sets ClientAttrs field to given value.

### HasClientAttrs

`func (o *BlueprintResourcesDefStatus) HasClientAttrs() bool`

HasClientAttrs returns a boolean if a field has been set.

### GetServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) GetServiceDefinitionList() []AppServiceResponse`

GetServiceDefinitionList returns the ServiceDefinitionList field if non-nil, zero value otherwise.

### GetServiceDefinitionListOk

`func (o *BlueprintResourcesDefStatus) GetServiceDefinitionListOk() (*[]AppServiceResponse, bool)`

GetServiceDefinitionListOk returns a tuple with the ServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) SetServiceDefinitionList(v []AppServiceResponse)`

SetServiceDefinitionList sets ServiceDefinitionList field to given value.

### HasServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) HasServiceDefinitionList() bool`

HasServiceDefinitionList returns a boolean if a field has been set.

### GetSubstrateDefinitionList

`func (o *BlueprintResourcesDefStatus) GetSubstrateDefinitionList() []AppSubstrateResponse`

GetSubstrateDefinitionList returns the SubstrateDefinitionList field if non-nil, zero value otherwise.

### GetSubstrateDefinitionListOk

`func (o *BlueprintResourcesDefStatus) GetSubstrateDefinitionListOk() (*[]AppSubstrateResponse, bool)`

GetSubstrateDefinitionListOk returns a tuple with the SubstrateDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateDefinitionList

`func (o *BlueprintResourcesDefStatus) SetSubstrateDefinitionList(v []AppSubstrateResponse)`

SetSubstrateDefinitionList sets SubstrateDefinitionList field to given value.

### HasSubstrateDefinitionList

`func (o *BlueprintResourcesDefStatus) HasSubstrateDefinitionList() bool`

HasSubstrateDefinitionList returns a boolean if a field has been set.

### GetCredentialDefinitionList

`func (o *BlueprintResourcesDefStatus) GetCredentialDefinitionList() []AppCredentialResponse`

GetCredentialDefinitionList returns the CredentialDefinitionList field if non-nil, zero value otherwise.

### GetCredentialDefinitionListOk

`func (o *BlueprintResourcesDefStatus) GetCredentialDefinitionListOk() (*[]AppCredentialResponse, bool)`

GetCredentialDefinitionListOk returns a tuple with the CredentialDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDefinitionList

`func (o *BlueprintResourcesDefStatus) SetCredentialDefinitionList(v []AppCredentialResponse)`

SetCredentialDefinitionList sets CredentialDefinitionList field to given value.

### HasCredentialDefinitionList

`func (o *BlueprintResourcesDefStatus) HasCredentialDefinitionList() bool`

HasCredentialDefinitionList returns a boolean if a field has been set.

### GetPackageDefinitionList

`func (o *BlueprintResourcesDefStatus) GetPackageDefinitionList() []AppPackageResponse`

GetPackageDefinitionList returns the PackageDefinitionList field if non-nil, zero value otherwise.

### GetPackageDefinitionListOk

`func (o *BlueprintResourcesDefStatus) GetPackageDefinitionListOk() (*[]AppPackageResponse, bool)`

GetPackageDefinitionListOk returns a tuple with the PackageDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDefinitionList

`func (o *BlueprintResourcesDefStatus) SetPackageDefinitionList(v []AppPackageResponse)`

SetPackageDefinitionList sets PackageDefinitionList field to given value.

### HasPackageDefinitionList

`func (o *BlueprintResourcesDefStatus) HasPackageDefinitionList() bool`

HasPackageDefinitionList returns a boolean if a field has been set.

### GetAppProfileList

`func (o *BlueprintResourcesDefStatus) GetAppProfileList() []AppProfileResponse`

GetAppProfileList returns the AppProfileList field if non-nil, zero value otherwise.

### GetAppProfileListOk

`func (o *BlueprintResourcesDefStatus) GetAppProfileListOk() (*[]AppProfileResponse, bool)`

GetAppProfileListOk returns a tuple with the AppProfileList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfileList

`func (o *BlueprintResourcesDefStatus) SetAppProfileList(v []AppProfileResponse)`

SetAppProfileList sets AppProfileList field to given value.

### HasAppProfileList

`func (o *BlueprintResourcesDefStatus) HasAppProfileList() bool`

HasAppProfileList returns a boolean if a field has been set.

### GetPublishedServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) GetPublishedServiceDefinitionList() []AppPublishedServiceResponse`

GetPublishedServiceDefinitionList returns the PublishedServiceDefinitionList field if non-nil, zero value otherwise.

### GetPublishedServiceDefinitionListOk

`func (o *BlueprintResourcesDefStatus) GetPublishedServiceDefinitionListOk() (*[]AppPublishedServiceResponse, bool)`

GetPublishedServiceDefinitionListOk returns a tuple with the PublishedServiceDefinitionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) SetPublishedServiceDefinitionList(v []AppPublishedServiceResponse)`

SetPublishedServiceDefinitionList sets PublishedServiceDefinitionList field to given value.

### HasPublishedServiceDefinitionList

`func (o *BlueprintResourcesDefStatus) HasPublishedServiceDefinitionList() bool`

HasPublishedServiceDefinitionList returns a boolean if a field has been set.

### GetDefaultCredentialLocalReference

`func (o *BlueprintResourcesDefStatus) GetDefaultCredentialLocalReference() AppCredentialReference`

GetDefaultCredentialLocalReference returns the DefaultCredentialLocalReference field if non-nil, zero value otherwise.

### GetDefaultCredentialLocalReferenceOk

`func (o *BlueprintResourcesDefStatus) GetDefaultCredentialLocalReferenceOk() (*AppCredentialReference, bool)`

GetDefaultCredentialLocalReferenceOk returns a tuple with the DefaultCredentialLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCredentialLocalReference

`func (o *BlueprintResourcesDefStatus) SetDefaultCredentialLocalReference(v AppCredentialReference)`

SetDefaultCredentialLocalReference sets DefaultCredentialLocalReference field to given value.

### HasDefaultCredentialLocalReference

`func (o *BlueprintResourcesDefStatus) HasDefaultCredentialLocalReference() bool`

HasDefaultCredentialLocalReference returns a boolean if a field has been set.

### GetState

`func (o *BlueprintResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintResourcesDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *BlueprintResourcesDefStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintResourcesDefStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintResourcesDefStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintResourcesDefStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessageList

`func (o *BlueprintResourcesDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *BlueprintResourcesDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *BlueprintResourcesDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *BlueprintResourcesDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetIsCloned

`func (o *BlueprintResourcesDefStatus) GetIsCloned() bool`

GetIsCloned returns the IsCloned field if non-nil, zero value otherwise.

### GetIsClonedOk

`func (o *BlueprintResourcesDefStatus) GetIsClonedOk() (*bool, bool)`

GetIsClonedOk returns a tuple with the IsCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloned

`func (o *BlueprintResourcesDefStatus) SetIsCloned(v bool)`

SetIsCloned sets IsCloned field to given value.

### HasIsCloned

`func (o *BlueprintResourcesDefStatus) HasIsCloned() bool`

HasIsCloned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


