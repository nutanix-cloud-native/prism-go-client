# AppPublishedServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Singleton** | Pointer to **bool** | If True, then this service can only be in a deployment with replica 1  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) |  | 
**ElementList** | Pointer to [**[]AppPublishedServiceElement**](AppPublishedServiceElement.md) |  | [optional] 
**Type** | Pointer to **string** | Type of published service | [optional] [default to "K8S_SERVICE"]
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**Name** | **string** |  | 
**ConfigReference** | Pointer to [**AppPublishedServiceReference**](AppPublishedServiceReference.md) |  | [optional] 
**State** | **string** |  | 
**PortList** | Pointer to [**[]AppServicePort**](AppServicePort.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Tier** | Pointer to **string** | Service tier name | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for service | [optional] 
**Options** | Pointer to **map[string]interface{}** | Additional published service options | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) |  | 
**UUID** | **string** |  | 

## Methods

### NewAppPublishedServiceResponse

`func NewAppPublishedServiceResponse(actionList []AppActionResponse, name string, state string, variableList []AppVariableResponse, uUID string, ) *AppPublishedServiceResponse`

NewAppPublishedServiceResponse instantiates a new AppPublishedServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPublishedServiceResponseWithDefaults

`func NewAppPublishedServiceResponseWithDefaults() *AppPublishedServiceResponse`

NewAppPublishedServiceResponseWithDefaults instantiates a new AppPublishedServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleton

`func (o *AppPublishedServiceResponse) GetSingleton() bool`

GetSingleton returns the Singleton field if non-nil, zero value otherwise.

### GetSingletonOk

`func (o *AppPublishedServiceResponse) GetSingletonOk() (*bool, bool)`

GetSingletonOk returns a tuple with the Singleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleton

`func (o *AppPublishedServiceResponse) SetSingleton(v bool)`

SetSingleton sets Singleton field to given value.

### HasSingleton

`func (o *AppPublishedServiceResponse) HasSingleton() bool`

HasSingleton returns a boolean if a field has been set.

### GetDescription

`func (o *AppPublishedServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPublishedServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPublishedServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPublishedServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppPublishedServiceResponse) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppPublishedServiceResponse) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppPublishedServiceResponse) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetElementList

`func (o *AppPublishedServiceResponse) GetElementList() []AppPublishedServiceElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppPublishedServiceResponse) GetElementListOk() (*[]AppPublishedServiceElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppPublishedServiceResponse) SetElementList(v []AppPublishedServiceElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppPublishedServiceResponse) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetType

`func (o *AppPublishedServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPublishedServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPublishedServiceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppPublishedServiceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppPublishedServiceResponse) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppPublishedServiceResponse) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppPublishedServiceResponse) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppPublishedServiceResponse) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetName

`func (o *AppPublishedServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPublishedServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPublishedServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetConfigReference

`func (o *AppPublishedServiceResponse) GetConfigReference() AppPublishedServiceReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppPublishedServiceResponse) GetConfigReferenceOk() (*AppPublishedServiceReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppPublishedServiceResponse) SetConfigReference(v AppPublishedServiceReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppPublishedServiceResponse) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetState

`func (o *AppPublishedServiceResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppPublishedServiceResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppPublishedServiceResponse) SetState(v string)`

SetState sets State field to given value.


### GetPortList

`func (o *AppPublishedServiceResponse) GetPortList() []AppServicePort`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *AppPublishedServiceResponse) GetPortListOk() (*[]AppServicePort, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *AppPublishedServiceResponse) SetPortList(v []AppServicePort)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *AppPublishedServiceResponse) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetEditables

`func (o *AppPublishedServiceResponse) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppPublishedServiceResponse) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppPublishedServiceResponse) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppPublishedServiceResponse) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetTier

`func (o *AppPublishedServiceResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *AppPublishedServiceResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *AppPublishedServiceResponse) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *AppPublishedServiceResponse) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetMessageList

`func (o *AppPublishedServiceResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppPublishedServiceResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppPublishedServiceResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppPublishedServiceResponse) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetOptions

`func (o *AppPublishedServiceResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppPublishedServiceResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppPublishedServiceResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppPublishedServiceResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariableList

`func (o *AppPublishedServiceResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppPublishedServiceResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppPublishedServiceResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetUUID

`func (o *AppPublishedServiceResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppPublishedServiceResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppPublishedServiceResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


