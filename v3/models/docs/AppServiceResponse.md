# AppServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | **string** |  | 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) |  | 
**ElementList** | Pointer to [**[]AppServiceElement**](AppServiceElement.md) |  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**ConfigReference** | Pointer to [**AppServiceReference**](AppServiceReference.md) |  | [optional] 
**State** | **string** |  | 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**ContainerSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s continaer spec | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for service | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppServiceResponse

`func NewAppServiceResponse(uUID string, actionList []AppActionResponse, name string, state string, variableList []AppVariableResponse, ) *AppServiceResponse`

NewAppServiceResponse instantiates a new AppServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceResponseWithDefaults

`func NewAppServiceResponseWithDefaults() *AppServiceResponse`

NewAppServiceResponseWithDefaults instantiates a new AppServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppServiceResponse) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppServiceResponse) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppServiceResponse) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppServiceResponse) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppServiceResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppServiceResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppServiceResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetActionList

`func (o *AppServiceResponse) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppServiceResponse) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppServiceResponse) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetElementList

`func (o *AppServiceResponse) GetElementList() []AppServiceElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppServiceResponse) GetElementListOk() (*[]AppServiceElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppServiceResponse) SetElementList(v []AppServiceElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppServiceResponse) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppServiceResponse) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppServiceResponse) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppServiceResponse) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppServiceResponse) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetConfigReference

`func (o *AppServiceResponse) GetConfigReference() AppServiceReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppServiceResponse) GetConfigReferenceOk() (*AppServiceReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppServiceResponse) SetConfigReference(v AppServiceReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppServiceResponse) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetState

`func (o *AppServiceResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServiceResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServiceResponse) SetState(v string)`

SetState sets State field to given value.


### GetPortList

`func (o *AppServiceResponse) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppServiceResponse) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppServiceResponse) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppServiceResponse) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppServiceResponse) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppServiceResponse) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppServiceResponse) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppServiceResponse) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppServiceResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppServiceResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppServiceResponse) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppServiceResponse) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetContainerSpec

`func (o *AppServiceResponse) GetContainerSpec() map[string]interface{}`

GetContainerSpec returns the ContainerSpec field if non-nil, zero value otherwise.

### GetContainerSpecOk

`func (o *AppServiceResponse) GetContainerSpecOk() (*map[string]interface{}, bool)`

GetContainerSpecOk returns a tuple with the ContainerSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSpec

`func (o *AppServiceResponse) SetContainerSpec(v map[string]interface{})`

SetContainerSpec sets ContainerSpec field to given value.

### HasContainerSpec

`func (o *AppServiceResponse) HasContainerSpec() bool`

HasContainerSpec returns a boolean if a field has been set.

### GetMessageList

`func (o *AppServiceResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppServiceResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppServiceResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppServiceResponse) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetVariableList

`func (o *AppServiceResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppServiceResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppServiceResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDescription

`func (o *AppServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


