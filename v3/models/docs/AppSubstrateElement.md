# AppSubstrateElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformData** | Pointer to **string** |  | [optional] 
**InstanceAddress** | Pointer to **string** |  | [optional] 
**InstancePowerState** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**UUID** | **string** |  | 
**Name** | **string** |  | 
**InstanceName** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**ReadinessProbe** | Pointer to [**AppSubstrateReadinessProbe**](AppSubstrateReadinessProbe.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**CreateSpec** | Pointer to **map[string]interface{}** | Spec of the substrate | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) | List of variables | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSubstrateElement

`func NewAppSubstrateElement(actionList []AppActionResponse, uUID string, name string, state string, type_ string, variableList []AppVariableResponse, ) *AppSubstrateElement`

NewAppSubstrateElement instantiates a new AppSubstrateElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateElementWithDefaults

`func NewAppSubstrateElementWithDefaults() *AppSubstrateElement`

NewAppSubstrateElementWithDefaults instantiates a new AppSubstrateElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformData

`func (o *AppSubstrateElement) GetPlatformData() string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *AppSubstrateElement) GetPlatformDataOk() (*string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *AppSubstrateElement) SetPlatformData(v string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *AppSubstrateElement) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *AppSubstrateElement) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *AppSubstrateElement) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *AppSubstrateElement) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *AppSubstrateElement) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetInstancePowerState

`func (o *AppSubstrateElement) GetInstancePowerState() string`

GetInstancePowerState returns the InstancePowerState field if non-nil, zero value otherwise.

### GetInstancePowerStateOk

`func (o *AppSubstrateElement) GetInstancePowerStateOk() (*string, bool)`

GetInstancePowerStateOk returns a tuple with the InstancePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePowerState

`func (o *AppSubstrateElement) SetInstancePowerState(v string)`

SetInstancePowerState sets InstancePowerState field to given value.

### HasInstancePowerState

`func (o *AppSubstrateElement) HasInstancePowerState() bool`

HasInstancePowerState returns a boolean if a field has been set.

### GetActionList

`func (o *AppSubstrateElement) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppSubstrateElement) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppSubstrateElement) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetInstanceId

`func (o *AppSubstrateElement) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppSubstrateElement) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppSubstrateElement) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppSubstrateElement) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUUID

`func (o *AppSubstrateElement) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppSubstrateElement) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppSubstrateElement) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetName

`func (o *AppSubstrateElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppSubstrateElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppSubstrateElement) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceName

`func (o *AppSubstrateElement) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AppSubstrateElement) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AppSubstrateElement) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AppSubstrateElement) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetState

`func (o *AppSubstrateElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppSubstrateElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppSubstrateElement) SetState(v string)`

SetState sets State field to given value.


### GetReadinessProbe

`func (o *AppSubstrateElement) GetReadinessProbe() AppSubstrateReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *AppSubstrateElement) GetReadinessProbeOk() (*AppSubstrateReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *AppSubstrateElement) SetReadinessProbe(v AppSubstrateReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *AppSubstrateElement) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetEditables

`func (o *AppSubstrateElement) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppSubstrateElement) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppSubstrateElement) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppSubstrateElement) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetMessageList

`func (o *AppSubstrateElement) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppSubstrateElement) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppSubstrateElement) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppSubstrateElement) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetOsType

`func (o *AppSubstrateElement) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AppSubstrateElement) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AppSubstrateElement) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AppSubstrateElement) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetType

`func (o *AppSubstrateElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubstrateElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubstrateElement) SetType(v string)`

SetType sets Type field to given value.


### GetCreateSpec

`func (o *AppSubstrateElement) GetCreateSpec() map[string]interface{}`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *AppSubstrateElement) GetCreateSpecOk() (*map[string]interface{}, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *AppSubstrateElement) SetCreateSpec(v map[string]interface{})`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *AppSubstrateElement) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppSubstrateElement) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppSubstrateElement) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppSubstrateElement) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetDescription

`func (o *AppSubstrateElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSubstrateElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSubstrateElement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSubstrateElement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


