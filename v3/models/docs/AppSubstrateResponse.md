# AppSubstrateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformData** | Pointer to **string** |  | [optional] 
**InstanceAddress** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | 
**ElementList** | Pointer to [**[]AppSubstrateElement**](AppSubstrateElement.md) |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**UUID** | **string** |  | 
**AccountReference** | Pointer to [**AccountReference**](AccountReference.md) |  | [optional] 
**Name** | **string** |  | 
**InstanceName** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**ReadinessProbe** | Pointer to [**AppSubstrateReadinessProbe**](AppSubstrateReadinessProbe.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**ConfigReference** | Pointer to [**AppSubstrateReference**](AppSubstrateReference.md) |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**CreateSpec** | Pointer to **map[string]interface{}** | Spec of the substrate | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) | List of variables | 
**InstancePowerState** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSubstrateResponse

`func NewAppSubstrateResponse(actionList []AppActionResponse, uUID string, name string, state string, type_ string, variableList []AppVariableResponse, ) *AppSubstrateResponse`

NewAppSubstrateResponse instantiates a new AppSubstrateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateResponseWithDefaults

`func NewAppSubstrateResponseWithDefaults() *AppSubstrateResponse`

NewAppSubstrateResponseWithDefaults instantiates a new AppSubstrateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformData

`func (o *AppSubstrateResponse) GetPlatformData() string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *AppSubstrateResponse) GetPlatformDataOk() (*string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *AppSubstrateResponse) SetPlatformData(v string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *AppSubstrateResponse) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *AppSubstrateResponse) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *AppSubstrateResponse) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *AppSubstrateResponse) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *AppSubstrateResponse) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetDescription

`func (o *AppSubstrateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSubstrateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSubstrateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSubstrateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppSubstrateResponse) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppSubstrateResponse) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppSubstrateResponse) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetElementList

`func (o *AppSubstrateResponse) GetElementList() []AppSubstrateElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppSubstrateResponse) GetElementListOk() (*[]AppSubstrateElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppSubstrateResponse) SetElementList(v []AppSubstrateElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppSubstrateResponse) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetInstanceId

`func (o *AppSubstrateResponse) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppSubstrateResponse) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppSubstrateResponse) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppSubstrateResponse) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUUID

`func (o *AppSubstrateResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppSubstrateResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppSubstrateResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetAccountReference

`func (o *AppSubstrateResponse) GetAccountReference() AccountReference`

GetAccountReference returns the AccountReference field if non-nil, zero value otherwise.

### GetAccountReferenceOk

`func (o *AppSubstrateResponse) GetAccountReferenceOk() (*AccountReference, bool)`

GetAccountReferenceOk returns a tuple with the AccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReference

`func (o *AppSubstrateResponse) SetAccountReference(v AccountReference)`

SetAccountReference sets AccountReference field to given value.

### HasAccountReference

`func (o *AppSubstrateResponse) HasAccountReference() bool`

HasAccountReference returns a boolean if a field has been set.

### GetName

`func (o *AppSubstrateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppSubstrateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppSubstrateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceName

`func (o *AppSubstrateResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AppSubstrateResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AppSubstrateResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AppSubstrateResponse) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetState

`func (o *AppSubstrateResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppSubstrateResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppSubstrateResponse) SetState(v string)`

SetState sets State field to given value.


### GetReadinessProbe

`func (o *AppSubstrateResponse) GetReadinessProbe() AppSubstrateReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *AppSubstrateResponse) GetReadinessProbeOk() (*AppSubstrateReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *AppSubstrateResponse) SetReadinessProbe(v AppSubstrateReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *AppSubstrateResponse) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetEditables

`func (o *AppSubstrateResponse) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppSubstrateResponse) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppSubstrateResponse) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppSubstrateResponse) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetMessageList

`func (o *AppSubstrateResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppSubstrateResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppSubstrateResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppSubstrateResponse) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppSubstrateResponse) GetConfigReference() AppSubstrateReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppSubstrateResponse) GetConfigReferenceOk() (*AppSubstrateReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppSubstrateResponse) SetConfigReference(v AppSubstrateReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppSubstrateResponse) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetOsType

`func (o *AppSubstrateResponse) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AppSubstrateResponse) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AppSubstrateResponse) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AppSubstrateResponse) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetType

`func (o *AppSubstrateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubstrateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubstrateResponse) SetType(v string)`

SetType sets Type field to given value.


### GetCreateSpec

`func (o *AppSubstrateResponse) GetCreateSpec() map[string]interface{}`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *AppSubstrateResponse) GetCreateSpecOk() (*map[string]interface{}, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *AppSubstrateResponse) SetCreateSpec(v map[string]interface{})`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *AppSubstrateResponse) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppSubstrateResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppSubstrateResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppSubstrateResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetInstancePowerState

`func (o *AppSubstrateResponse) GetInstancePowerState() string`

GetInstancePowerState returns the InstancePowerState field if non-nil, zero value otherwise.

### GetInstancePowerStateOk

`func (o *AppSubstrateResponse) GetInstancePowerStateOk() (*string, bool)`

GetInstancePowerStateOk returns a tuple with the InstancePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePowerState

`func (o *AppSubstrateResponse) SetInstancePowerState(v string)`

SetInstancePowerState sets InstancePowerState field to given value.

### HasInstancePowerState

`func (o *AppSubstrateResponse) HasInstancePowerState() bool`

HasInstancePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


