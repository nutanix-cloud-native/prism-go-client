# AppServiceInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to service action  | 
**Description** | Pointer to **string** |  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**ConfigReference** | Pointer to [**AppServiceReferenceUpload**](AppServiceReferenceUpload.md) |  | [optional] 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**ContainerSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s continaer spec | [optional] 
**VariableList** | [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | 
**Name** | **string** |  | 

## Methods

### NewAppServiceInputUpload

`func NewAppServiceInputUpload(actionList []AppActionInputUpload, variableList []AppVariableInputUpload, name string, ) *AppServiceInputUpload`

NewAppServiceInputUpload instantiates a new AppServiceInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceInputUploadWithDefaults

`func NewAppServiceInputUploadWithDefaults() *AppServiceInputUpload`

NewAppServiceInputUploadWithDefaults instantiates a new AppServiceInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppServiceInputUpload) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppServiceInputUpload) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppServiceInputUpload) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppServiceInputUpload) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppServiceInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppServiceInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppServiceInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppServiceInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetActionList

`func (o *AppServiceInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppServiceInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppServiceInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.


### GetDescription

`func (o *AppServiceInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppServiceInputUpload) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppServiceInputUpload) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppServiceInputUpload) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppServiceInputUpload) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppServiceInputUpload) GetConfigReference() AppServiceReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppServiceInputUpload) GetConfigReferenceOk() (*AppServiceReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppServiceInputUpload) SetConfigReference(v AppServiceReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppServiceInputUpload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetPortList

`func (o *AppServiceInputUpload) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppServiceInputUpload) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppServiceInputUpload) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppServiceInputUpload) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppServiceInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppServiceInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppServiceInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppServiceInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppServiceInputUpload) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppServiceInputUpload) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppServiceInputUpload) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppServiceInputUpload) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetContainerSpec

`func (o *AppServiceInputUpload) GetContainerSpec() map[string]interface{}`

GetContainerSpec returns the ContainerSpec field if non-nil, zero value otherwise.

### GetContainerSpecOk

`func (o *AppServiceInputUpload) GetContainerSpecOk() (*map[string]interface{}, bool)`

GetContainerSpecOk returns a tuple with the ContainerSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSpec

`func (o *AppServiceInputUpload) SetContainerSpec(v map[string]interface{})`

SetContainerSpec sets ContainerSpec field to given value.

### HasContainerSpec

`func (o *AppServiceInputUpload) HasContainerSpec() bool`

HasContainerSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppServiceInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppServiceInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppServiceInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.


### GetName

`func (o *AppServiceInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceInputUpload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


