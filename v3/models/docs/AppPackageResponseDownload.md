# AppPackageResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s image spec | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | [optional] 
**ElementList** | Pointer to [**[]AppPackageElement**](AppPackageElement.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for package | [optional] 
**ServiceLocalReferenceList** | Pointer to [**[]AppServiceReference**](AppServiceReference.md) | References of the service. | [optional] 
**AccountReference** | Pointer to [**AccountReferenceUpload**](AccountReferenceUpload.md) |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**ServiceElementLocalReferenceList** | Pointer to [**[]AppServiceElement**](AppServiceElement.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**ConfigReference** | Pointer to [**AppPackageReferenceUpload**](AppPackageReferenceUpload.md) |  | [optional] 
**Type** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** | Details based on type of the package. | [optional] 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppPackageResponseDownload

`func NewAppPackageResponseDownload(name string, state string, type_ string, uUID string, ) *AppPackageResponseDownload`

NewAppPackageResponseDownload instantiates a new AppPackageResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPackageResponseDownloadWithDefaults

`func NewAppPackageResponseDownloadWithDefaults() *AppPackageResponseDownload`

NewAppPackageResponseDownloadWithDefaults instantiates a new AppPackageResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageSpec

`func (o *AppPackageResponseDownload) GetImageSpec() map[string]interface{}`

GetImageSpec returns the ImageSpec field if non-nil, zero value otherwise.

### GetImageSpecOk

`func (o *AppPackageResponseDownload) GetImageSpecOk() (*map[string]interface{}, bool)`

GetImageSpecOk returns a tuple with the ImageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSpec

`func (o *AppPackageResponseDownload) SetImageSpec(v map[string]interface{})`

SetImageSpec sets ImageSpec field to given value.

### HasImageSpec

`func (o *AppPackageResponseDownload) HasImageSpec() bool`

HasImageSpec returns a boolean if a field has been set.

### GetDescription

`func (o *AppPackageResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPackageResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPackageResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPackageResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppPackageResponseDownload) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPackageResponseDownload) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPackageResponseDownload) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppPackageResponseDownload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetElementList

`func (o *AppPackageResponseDownload) GetElementList() []AppPackageElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppPackageResponseDownload) GetElementListOk() (*[]AppPackageElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppPackageResponseDownload) SetElementList(v []AppPackageElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppPackageResponseDownload) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetMessageList

`func (o *AppPackageResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppPackageResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppPackageResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppPackageResponseDownload) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetServiceLocalReferenceList

`func (o *AppPackageResponseDownload) GetServiceLocalReferenceList() []AppServiceReference`

GetServiceLocalReferenceList returns the ServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetServiceLocalReferenceListOk

`func (o *AppPackageResponseDownload) GetServiceLocalReferenceListOk() (*[]AppServiceReference, bool)`

GetServiceLocalReferenceListOk returns a tuple with the ServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocalReferenceList

`func (o *AppPackageResponseDownload) SetServiceLocalReferenceList(v []AppServiceReference)`

SetServiceLocalReferenceList sets ServiceLocalReferenceList field to given value.

### HasServiceLocalReferenceList

`func (o *AppPackageResponseDownload) HasServiceLocalReferenceList() bool`

HasServiceLocalReferenceList returns a boolean if a field has been set.

### GetAccountReference

`func (o *AppPackageResponseDownload) GetAccountReference() AccountReferenceUpload`

GetAccountReference returns the AccountReference field if non-nil, zero value otherwise.

### GetAccountReferenceOk

`func (o *AppPackageResponseDownload) GetAccountReferenceOk() (*AccountReferenceUpload, bool)`

GetAccountReferenceOk returns a tuple with the AccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReference

`func (o *AppPackageResponseDownload) SetAccountReference(v AccountReferenceUpload)`

SetAccountReference sets AccountReference field to given value.

### HasAccountReference

`func (o *AppPackageResponseDownload) HasAccountReference() bool`

HasAccountReference returns a boolean if a field has been set.

### GetName

`func (o *AppPackageResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPackageResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPackageResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppPackageResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppPackageResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppPackageResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetVersion

`func (o *AppPackageResponseDownload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppPackageResponseDownload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppPackageResponseDownload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppPackageResponseDownload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServiceElementLocalReferenceList

`func (o *AppPackageResponseDownload) GetServiceElementLocalReferenceList() []AppServiceElement`

GetServiceElementLocalReferenceList returns the ServiceElementLocalReferenceList field if non-nil, zero value otherwise.

### GetServiceElementLocalReferenceListOk

`func (o *AppPackageResponseDownload) GetServiceElementLocalReferenceListOk() (*[]AppServiceElement, bool)`

GetServiceElementLocalReferenceListOk returns a tuple with the ServiceElementLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceElementLocalReferenceList

`func (o *AppPackageResponseDownload) SetServiceElementLocalReferenceList(v []AppServiceElement)`

SetServiceElementLocalReferenceList sets ServiceElementLocalReferenceList field to given value.

### HasServiceElementLocalReferenceList

`func (o *AppPackageResponseDownload) HasServiceElementLocalReferenceList() bool`

HasServiceElementLocalReferenceList returns a boolean if a field has been set.

### GetEditables

`func (o *AppPackageResponseDownload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPackageResponseDownload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPackageResponseDownload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPackageResponseDownload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppPackageResponseDownload) GetConfigReference() AppPackageReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppPackageResponseDownload) GetConfigReferenceOk() (*AppPackageReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppPackageResponseDownload) SetConfigReference(v AppPackageReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppPackageResponseDownload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetType

`func (o *AppPackageResponseDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPackageResponseDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPackageResponseDownload) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *AppPackageResponseDownload) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPackageResponseDownload) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPackageResponseDownload) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPackageResponseDownload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPackageResponseDownload) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPackageResponseDownload) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPackageResponseDownload) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppPackageResponseDownload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppPackageResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPackageResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPackageResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


