# AppServiceElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | **string** |  | 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) |  | 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for service | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppServiceElement

`func NewAppServiceElement(uUID string, actionList []AppActionResponse, name string, state string, variableList []AppVariableResponse, ) *AppServiceElement`

NewAppServiceElement instantiates a new AppServiceElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceElementWithDefaults

`func NewAppServiceElementWithDefaults() *AppServiceElement`

NewAppServiceElementWithDefaults instantiates a new AppServiceElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppServiceElement) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppServiceElement) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppServiceElement) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppServiceElement) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppServiceElement) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppServiceElement) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppServiceElement) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetActionList

`func (o *AppServiceElement) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppServiceElement) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppServiceElement) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetDependsOnList

`func (o *AppServiceElement) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppServiceElement) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppServiceElement) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppServiceElement) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppServiceElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceElement) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppServiceElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServiceElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServiceElement) SetState(v string)`

SetState sets State field to given value.


### GetPortList

`func (o *AppServiceElement) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppServiceElement) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppServiceElement) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppServiceElement) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppServiceElement) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppServiceElement) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppServiceElement) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppServiceElement) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppServiceElement) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppServiceElement) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppServiceElement) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppServiceElement) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetMessageList

`func (o *AppServiceElement) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppServiceElement) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppServiceElement) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppServiceElement) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetVariableList

`func (o *AppServiceElement) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppServiceElement) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppServiceElement) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDescription

`func (o *AppServiceElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


