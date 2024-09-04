# AppPackageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s image spec | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInput**](AppActionInput.md) | List of references to action  | [optional] 
**ServiceLocalReferenceList** | Pointer to [**[]AppServiceReference**](AppServiceReference.md) | References of the service. | [optional] 
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**ConfigReference** | Pointer to [**AppPackageReference**](AppPackageReference.md) |  | [optional] 
**Type** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** | Details based on type of the package. | [optional] 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppPackageInput

`func NewAppPackageInput(name string, type_ string, uUID string, ) *AppPackageInput`

NewAppPackageInput instantiates a new AppPackageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPackageInputWithDefaults

`func NewAppPackageInputWithDefaults() *AppPackageInput`

NewAppPackageInputWithDefaults instantiates a new AppPackageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageSpec

`func (o *AppPackageInput) GetImageSpec() map[string]interface{}`

GetImageSpec returns the ImageSpec field if non-nil, zero value otherwise.

### GetImageSpecOk

`func (o *AppPackageInput) GetImageSpecOk() (*map[string]interface{}, bool)`

GetImageSpecOk returns a tuple with the ImageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSpec

`func (o *AppPackageInput) SetImageSpec(v map[string]interface{})`

SetImageSpec sets ImageSpec field to given value.

### HasImageSpec

`func (o *AppPackageInput) HasImageSpec() bool`

HasImageSpec returns a boolean if a field has been set.

### GetDescription

`func (o *AppPackageInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPackageInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPackageInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPackageInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppPackageInput) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPackageInput) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPackageInput) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppPackageInput) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetServiceLocalReferenceList

`func (o *AppPackageInput) GetServiceLocalReferenceList() []AppServiceReference`

GetServiceLocalReferenceList returns the ServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetServiceLocalReferenceListOk

`func (o *AppPackageInput) GetServiceLocalReferenceListOk() (*[]AppServiceReference, bool)`

GetServiceLocalReferenceListOk returns a tuple with the ServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocalReferenceList

`func (o *AppPackageInput) SetServiceLocalReferenceList(v []AppServiceReference)`

SetServiceLocalReferenceList sets ServiceLocalReferenceList field to given value.

### HasServiceLocalReferenceList

`func (o *AppPackageInput) HasServiceLocalReferenceList() bool`

HasServiceLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppPackageInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPackageInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPackageInput) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *AppPackageInput) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppPackageInput) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppPackageInput) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppPackageInput) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEditables

`func (o *AppPackageInput) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPackageInput) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPackageInput) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPackageInput) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppPackageInput) GetConfigReference() AppPackageReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppPackageInput) GetConfigReferenceOk() (*AppPackageReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppPackageInput) SetConfigReference(v AppPackageReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppPackageInput) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetType

`func (o *AppPackageInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPackageInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPackageInput) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *AppPackageInput) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPackageInput) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPackageInput) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPackageInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPackageInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPackageInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPackageInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppPackageInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppPackageInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPackageInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPackageInput) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


