# AppServiceResponseDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | **string** |  | 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) |  | 
**ElementList** | Pointer to [**[]AppServiceElement**](AppServiceElement.md) |  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**ConfigReference** | Pointer to [**AppServiceReferenceUpload**](AppServiceReferenceUpload.md) |  | [optional] 
**State** | **string** |  | 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**ContainerSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s continaer spec | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for service | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppServiceResponseDownload

`func NewAppServiceResponseDownload(uUID string, actionList []AppActionResponse, name string, state string, variableList []AppVariableResponse, ) *AppServiceResponseDownload`

NewAppServiceResponseDownload instantiates a new AppServiceResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceResponseDownloadWithDefaults

`func NewAppServiceResponseDownloadWithDefaults() *AppServiceResponseDownload`

NewAppServiceResponseDownloadWithDefaults instantiates a new AppServiceResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppServiceResponseDownload) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppServiceResponseDownload) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppServiceResponseDownload) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppServiceResponseDownload) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppServiceResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppServiceResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppServiceResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetActionList

`func (o *AppServiceResponseDownload) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppServiceResponseDownload) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppServiceResponseDownload) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetElementList

`func (o *AppServiceResponseDownload) GetElementList() []AppServiceElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppServiceResponseDownload) GetElementListOk() (*[]AppServiceElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppServiceResponseDownload) SetElementList(v []AppServiceElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppServiceResponseDownload) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppServiceResponseDownload) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppServiceResponseDownload) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppServiceResponseDownload) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppServiceResponseDownload) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppServiceResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetConfigReference

`func (o *AppServiceResponseDownload) GetConfigReference() AppServiceReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppServiceResponseDownload) GetConfigReferenceOk() (*AppServiceReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppServiceResponseDownload) SetConfigReference(v AppServiceReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppServiceResponseDownload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetState

`func (o *AppServiceResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServiceResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServiceResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetPortList

`func (o *AppServiceResponseDownload) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppServiceResponseDownload) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppServiceResponseDownload) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppServiceResponseDownload) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppServiceResponseDownload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppServiceResponseDownload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppServiceResponseDownload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppServiceResponseDownload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppServiceResponseDownload) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppServiceResponseDownload) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppServiceResponseDownload) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppServiceResponseDownload) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetContainerSpec

`func (o *AppServiceResponseDownload) GetContainerSpec() map[string]interface{}`

GetContainerSpec returns the ContainerSpec field if non-nil, zero value otherwise.

### GetContainerSpecOk

`func (o *AppServiceResponseDownload) GetContainerSpecOk() (*map[string]interface{}, bool)`

GetContainerSpecOk returns a tuple with the ContainerSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSpec

`func (o *AppServiceResponseDownload) SetContainerSpec(v map[string]interface{})`

SetContainerSpec sets ContainerSpec field to given value.

### HasContainerSpec

`func (o *AppServiceResponseDownload) HasContainerSpec() bool`

HasContainerSpec returns a boolean if a field has been set.

### GetMessageList

`func (o *AppServiceResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppServiceResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppServiceResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppServiceResponseDownload) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetVariableList

`func (o *AppServiceResponseDownload) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppServiceResponseDownload) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppServiceResponseDownload) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDescription

`func (o *AppServiceResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


