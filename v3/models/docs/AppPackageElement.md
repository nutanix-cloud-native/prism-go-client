# AppPackageElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s image spec | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for package | [optional] 
**ServiceLocalReferenceList** | Pointer to [**[]AppServiceReference**](AppServiceReference.md) | References of the service. | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Type** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** | Details based on type of the package. | [optional] 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppPackageElement

`func NewAppPackageElement(name string, state string, type_ string, uUID string, ) *AppPackageElement`

NewAppPackageElement instantiates a new AppPackageElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPackageElementWithDefaults

`func NewAppPackageElementWithDefaults() *AppPackageElement`

NewAppPackageElementWithDefaults instantiates a new AppPackageElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageSpec

`func (o *AppPackageElement) GetImageSpec() map[string]interface{}`

GetImageSpec returns the ImageSpec field if non-nil, zero value otherwise.

### GetImageSpecOk

`func (o *AppPackageElement) GetImageSpecOk() (*map[string]interface{}, bool)`

GetImageSpecOk returns a tuple with the ImageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSpec

`func (o *AppPackageElement) SetImageSpec(v map[string]interface{})`

SetImageSpec sets ImageSpec field to given value.

### HasImageSpec

`func (o *AppPackageElement) HasImageSpec() bool`

HasImageSpec returns a boolean if a field has been set.

### GetDescription

`func (o *AppPackageElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPackageElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPackageElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPackageElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppPackageElement) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPackageElement) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPackageElement) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppPackageElement) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetMessageList

`func (o *AppPackageElement) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppPackageElement) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppPackageElement) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppPackageElement) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetServiceLocalReferenceList

`func (o *AppPackageElement) GetServiceLocalReferenceList() []AppServiceReference`

GetServiceLocalReferenceList returns the ServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetServiceLocalReferenceListOk

`func (o *AppPackageElement) GetServiceLocalReferenceListOk() (*[]AppServiceReference, bool)`

GetServiceLocalReferenceListOk returns a tuple with the ServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLocalReferenceList

`func (o *AppPackageElement) SetServiceLocalReferenceList(v []AppServiceReference)`

SetServiceLocalReferenceList sets ServiceLocalReferenceList field to given value.

### HasServiceLocalReferenceList

`func (o *AppPackageElement) HasServiceLocalReferenceList() bool`

HasServiceLocalReferenceList returns a boolean if a field has been set.

### GetName

`func (o *AppPackageElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPackageElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPackageElement) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppPackageElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppPackageElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppPackageElement) SetState(v string)`

SetState sets State field to given value.


### GetVersion

`func (o *AppPackageElement) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppPackageElement) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppPackageElement) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppPackageElement) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEditables

`func (o *AppPackageElement) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPackageElement) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPackageElement) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPackageElement) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetType

`func (o *AppPackageElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPackageElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPackageElement) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *AppPackageElement) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPackageElement) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPackageElement) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPackageElement) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPackageElement) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPackageElement) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPackageElement) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppPackageElement) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetUUID

`func (o *AppPackageElement) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPackageElement) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPackageElement) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


