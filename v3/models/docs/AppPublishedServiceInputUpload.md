# AppPublishedServiceInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to service action  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**ConfigReference** | Pointer to [**AppPublishedServiceReferenceUpload**](AppPublishedServiceReferenceUpload.md) |  | [optional] 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**Type** | Pointer to **string** | Type of published service | [optional] [default to "K8S_SERVICE"]
**Options** | Pointer to **map[string]interface{}** | Additional published service options | [optional] 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppPublishedServiceInputUpload

`func NewAppPublishedServiceInputUpload(name string, ) *AppPublishedServiceInputUpload`

NewAppPublishedServiceInputUpload instantiates a new AppPublishedServiceInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPublishedServiceInputUploadWithDefaults

`func NewAppPublishedServiceInputUploadWithDefaults() *AppPublishedServiceInputUpload`

NewAppPublishedServiceInputUploadWithDefaults instantiates a new AppPublishedServiceInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppPublishedServiceInputUpload) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppPublishedServiceInputUpload) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppPublishedServiceInputUpload) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppPublishedServiceInputUpload) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppPublishedServiceInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPublishedServiceInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPublishedServiceInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppPublishedServiceInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetActionList

`func (o *AppPublishedServiceInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPublishedServiceInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPublishedServiceInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppPublishedServiceInputUpload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppPublishedServiceInputUpload) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppPublishedServiceInputUpload) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppPublishedServiceInputUpload) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppPublishedServiceInputUpload) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppPublishedServiceInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPublishedServiceInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPublishedServiceInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetConfigReference

`func (o *AppPublishedServiceInputUpload) GetConfigReference() AppPublishedServiceReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppPublishedServiceInputUpload) GetConfigReferenceOk() (*AppPublishedServiceReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppPublishedServiceInputUpload) SetConfigReference(v AppPublishedServiceReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppPublishedServiceInputUpload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetPortList

`func (o *AppPublishedServiceInputUpload) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppPublishedServiceInputUpload) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppPublishedServiceInputUpload) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppPublishedServiceInputUpload) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppPublishedServiceInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPublishedServiceInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPublishedServiceInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPublishedServiceInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppPublishedServiceInputUpload) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppPublishedServiceInputUpload) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppPublishedServiceInputUpload) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppPublishedServiceInputUpload) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetType

`func (o *AppPublishedServiceInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPublishedServiceInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPublishedServiceInputUpload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppPublishedServiceInputUpload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *AppPublishedServiceInputUpload) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPublishedServiceInputUpload) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPublishedServiceInputUpload) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPublishedServiceInputUpload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPublishedServiceInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPublishedServiceInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPublishedServiceInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppPublishedServiceInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetDescription

`func (o *AppPublishedServiceInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPublishedServiceInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPublishedServiceInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPublishedServiceInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


