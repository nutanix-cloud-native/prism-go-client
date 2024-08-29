# AppSubstrateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformData** | Pointer to **string** |  | [optional] 
**InstanceAddress** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionInput**](AppActionInput.md) | List of references to action  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**UUID** | **string** |  | 
**Name** | **string** |  | 
**InstanceName** | Pointer to **string** |  | [optional] 
**ReadinessProbe** | Pointer to [**AppSubstrateReadinessProbe**](AppSubstrateReadinessProbe.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**ConfigReference** | Pointer to [**AppSubstrateReference**](AppSubstrateReference.md) |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**CreateSpec** | Pointer to **map[string]interface{}** | Spec of the substrate | [optional] 
**VariableList** | [**[]AppVariableInput**](AppVariableInput.md) | List of variables | 
**InstancePowerState** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSubstrateInput

`func NewAppSubstrateInput(actionList []AppActionInput, uUID string, name string, type_ string, variableList []AppVariableInput, ) *AppSubstrateInput`

NewAppSubstrateInput instantiates a new AppSubstrateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateInputWithDefaults

`func NewAppSubstrateInputWithDefaults() *AppSubstrateInput`

NewAppSubstrateInputWithDefaults instantiates a new AppSubstrateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformData

`func (o *AppSubstrateInput) GetPlatformData() string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *AppSubstrateInput) GetPlatformDataOk() (*string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *AppSubstrateInput) SetPlatformData(v string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *AppSubstrateInput) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *AppSubstrateInput) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *AppSubstrateInput) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *AppSubstrateInput) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *AppSubstrateInput) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetDescription

`func (o *AppSubstrateInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSubstrateInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSubstrateInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSubstrateInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppSubstrateInput) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppSubstrateInput) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppSubstrateInput) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.


### GetInstanceId

`func (o *AppSubstrateInput) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppSubstrateInput) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppSubstrateInput) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppSubstrateInput) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUUID

`func (o *AppSubstrateInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppSubstrateInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppSubstrateInput) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetName

`func (o *AppSubstrateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppSubstrateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppSubstrateInput) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceName

`func (o *AppSubstrateInput) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AppSubstrateInput) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AppSubstrateInput) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AppSubstrateInput) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *AppSubstrateInput) GetReadinessProbe() AppSubstrateReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *AppSubstrateInput) GetReadinessProbeOk() (*AppSubstrateReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *AppSubstrateInput) SetReadinessProbe(v AppSubstrateReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *AppSubstrateInput) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetEditables

`func (o *AppSubstrateInput) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppSubstrateInput) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppSubstrateInput) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppSubstrateInput) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppSubstrateInput) GetConfigReference() AppSubstrateReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppSubstrateInput) GetConfigReferenceOk() (*AppSubstrateReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppSubstrateInput) SetConfigReference(v AppSubstrateReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppSubstrateInput) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetOsType

`func (o *AppSubstrateInput) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AppSubstrateInput) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AppSubstrateInput) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AppSubstrateInput) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetType

`func (o *AppSubstrateInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubstrateInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubstrateInput) SetType(v string)`

SetType sets Type field to given value.


### GetCreateSpec

`func (o *AppSubstrateInput) GetCreateSpec() map[string]interface{}`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *AppSubstrateInput) GetCreateSpecOk() (*map[string]interface{}, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *AppSubstrateInput) SetCreateSpec(v map[string]interface{})`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *AppSubstrateInput) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppSubstrateInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppSubstrateInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppSubstrateInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.


### GetInstancePowerState

`func (o *AppSubstrateInput) GetInstancePowerState() string`

GetInstancePowerState returns the InstancePowerState field if non-nil, zero value otherwise.

### GetInstancePowerStateOk

`func (o *AppSubstrateInput) GetInstancePowerStateOk() (*string, bool)`

GetInstancePowerStateOk returns a tuple with the InstancePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePowerState

`func (o *AppSubstrateInput) SetInstancePowerState(v string)`

SetInstancePowerState sets InstancePowerState field to given value.

### HasInstancePowerState

`func (o *AppSubstrateInput) HasInstancePowerState() bool`

HasInstancePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


