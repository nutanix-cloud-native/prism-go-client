# AppPublishedServiceElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**UUID** | **string** |  | 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) |  | 
**Type** | Pointer to **string** | Type of published service | [optional] [default to "K8S_SERVICE"]
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**State** | **string** |  | 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for service | [optional] 
**Options** | Pointer to **map[string]interface{}** | Additional published service options | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppPublishedServiceElement

`func NewAppPublishedServiceElement(uUID string, actionList []AppActionResponse, name string, state string, variableList []AppVariableResponse, ) *AppPublishedServiceElement`

NewAppPublishedServiceElement instantiates a new AppPublishedServiceElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPublishedServiceElementWithDefaults

`func NewAppPublishedServiceElementWithDefaults() *AppPublishedServiceElement`

NewAppPublishedServiceElementWithDefaults instantiates a new AppPublishedServiceElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppPublishedServiceElement) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppPublishedServiceElement) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppPublishedServiceElement) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppPublishedServiceElement) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetUUID

`func (o *AppPublishedServiceElement) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPublishedServiceElement) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPublishedServiceElement) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetActionList

`func (o *AppPublishedServiceElement) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPublishedServiceElement) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPublishedServiceElement) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetType

`func (o *AppPublishedServiceElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPublishedServiceElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPublishedServiceElement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppPublishedServiceElement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppPublishedServiceElement) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppPublishedServiceElement) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppPublishedServiceElement) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppPublishedServiceElement) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppPublishedServiceElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPublishedServiceElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPublishedServiceElement) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppPublishedServiceElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppPublishedServiceElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppPublishedServiceElement) SetState(v string)`

SetState sets State field to given value.


### GetPortList

`func (o *AppPublishedServiceElement) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppPublishedServiceElement) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppPublishedServiceElement) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppPublishedServiceElement) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppPublishedServiceElement) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPublishedServiceElement) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPublishedServiceElement) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPublishedServiceElement) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppPublishedServiceElement) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppPublishedServiceElement) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppPublishedServiceElement) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppPublishedServiceElement) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetMessageList

`func (o *AppPublishedServiceElement) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppPublishedServiceElement) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppPublishedServiceElement) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppPublishedServiceElement) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetOptions

`func (o *AppPublishedServiceElement) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPublishedServiceElement) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPublishedServiceElement) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPublishedServiceElement) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPublishedServiceElement) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPublishedServiceElement) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPublishedServiceElement) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDescription

`func (o *AppPublishedServiceElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPublishedServiceElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPublishedServiceElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPublishedServiceElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


