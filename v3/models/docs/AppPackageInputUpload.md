# AppPackageInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s image spec | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to action  | [optional] 
**ServiceLocalReferenceList** | Pointer to [**[]AppServiceReferenceUpload**](AppServiceReferenceUpload.md) | References of the service. | [optional] 
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**ConfigReference** | Pointer to [**AppPackageReferenceUpload**](AppPackageReferenceUpload.md) |  | [optional] 
**Type** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** | Details based on type of the package. | [optional] 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppPackageInputUpload

`func NewAppPackageInputUpload(name string, type_ string, ) *AppPackageInputUpload`

NewAppPackageInputUpload instantiates a new AppPackageInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPackageInputUploadWithDefaults

`func NewAppPackageInputUploadWithDefaults() *AppPackageInputUpload`

NewAppPackageInputUploadWithDefaults instantiates a new AppPackageInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageSpec

`func (o *AppPackageInputUpload) GetImageSpec() map[string]interface{}`

GetImageSpec returns the ImageSpec field if non-nil, zero value otherwise.

### GetImageSpecOk

`func (o *AppPackageInputUpload) GetImageSpecOk() (*map[string]interface{}, bool)`

GetImageSpecOk returns a tuple with the ImageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSpec

`func (o *AppPackageInputUpload) SetImageSpec(v map[string]interface{})`

SetImageSpec sets ImageSpec field to given value.

### HasImageSpec

`func (o *AppPackageInputUpload) HasImageSpec() bool`

HasImageSpec returns a boolean if a field has been set.

### GetDescription

`func (o *AppPackageInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPackageInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPackageInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPackageInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppPackageInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPackageInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPackageInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppPackageInputUpload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetServiceLocalReferenceList

`func (o *AppPackageInputUpload) GetServiceLocalReferenceList() []AppServiceReferenceUpload`

GetServiceLocalReferenceList returns the ServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetServiceLocalReferenceListOk

`func (o *AppPackageInputUpload) GetServiceLocalReferenceListOk() (*[]AppServiceReferenceUpload, bool)`

GetServiceLocalReferenceListOk returns a tuple with the ServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocalReferenceList

`func (o *AppPackageInputUpload) SetServiceLocalReferenceList(v []AppServiceReferenceUpload)`

SetServiceLocalReferenceList sets ServiceLocalReferenceList field to given value.

### HasServiceLocalReferenceList

`func (o *AppPackageInputUpload) HasServiceLocalReferenceList() bool`

HasServiceLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppPackageInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPackageInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPackageInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *AppPackageInputUpload) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppPackageInputUpload) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppPackageInputUpload) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppPackageInputUpload) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEditables

`func (o *AppPackageInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPackageInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPackageInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPackageInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppPackageInputUpload) GetConfigReference() AppPackageReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppPackageInputUpload) GetConfigReferenceOk() (*AppPackageReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppPackageInputUpload) SetConfigReference(v AppPackageReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppPackageInputUpload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetType

`func (o *AppPackageInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPackageInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPackageInputUpload) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *AppPackageInputUpload) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPackageInputUpload) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPackageInputUpload) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPackageInputUpload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPackageInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPackageInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPackageInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppPackageInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppPackageInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPackageInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPackageInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppPackageInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


