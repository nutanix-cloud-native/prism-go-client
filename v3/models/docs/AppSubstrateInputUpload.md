# AppSubstrateInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformData** | Pointer to **string** |  | [optional] 
**InstanceAddress** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to action  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**InstanceName** | Pointer to **string** |  | [optional] 
**ReadinessProbe** | Pointer to [**AppSubstrateReadinessProbeUpload**](AppSubstrateReadinessProbeUpload.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**ConfigReference** | Pointer to [**AppSubstrateReferenceUpload**](AppSubstrateReferenceUpload.md) |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**CreateSpec** | Pointer to **map[string]interface{}** | Spec of the substrate | [optional] 
**VariableList** | [**[]AppVariableInputUpload**](AppVariableInputUpload.md) | List of variables | 
**InstancePowerState** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSubstrateInputUpload

`func NewAppSubstrateInputUpload(actionList []AppActionInputUpload, name string, type_ string, variableList []AppVariableInputUpload, ) *AppSubstrateInputUpload`

NewAppSubstrateInputUpload instantiates a new AppSubstrateInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateInputUploadWithDefaults

`func NewAppSubstrateInputUploadWithDefaults() *AppSubstrateInputUpload`

NewAppSubstrateInputUploadWithDefaults instantiates a new AppSubstrateInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformData

`func (o *AppSubstrateInputUpload) GetPlatformData() string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *AppSubstrateInputUpload) GetPlatformDataOk() (*string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *AppSubstrateInputUpload) SetPlatformData(v string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *AppSubstrateInputUpload) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *AppSubstrateInputUpload) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *AppSubstrateInputUpload) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *AppSubstrateInputUpload) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *AppSubstrateInputUpload) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetDescription

`func (o *AppSubstrateInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSubstrateInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSubstrateInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSubstrateInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppSubstrateInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppSubstrateInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppSubstrateInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.


### GetInstanceId

`func (o *AppSubstrateInputUpload) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppSubstrateInputUpload) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppSubstrateInputUpload) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppSubstrateInputUpload) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUUID

`func (o *AppSubstrateInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppSubstrateInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppSubstrateInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppSubstrateInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *AppSubstrateInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppSubstrateInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppSubstrateInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceName

`func (o *AppSubstrateInputUpload) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AppSubstrateInputUpload) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AppSubstrateInputUpload) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AppSubstrateInputUpload) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *AppSubstrateInputUpload) GetReadinessProbe() AppSubstrateReadinessProbeUpload`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *AppSubstrateInputUpload) GetReadinessProbeOk() (*AppSubstrateReadinessProbeUpload, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *AppSubstrateInputUpload) SetReadinessProbe(v AppSubstrateReadinessProbeUpload)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *AppSubstrateInputUpload) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetEditables

`func (o *AppSubstrateInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppSubstrateInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppSubstrateInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppSubstrateInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppSubstrateInputUpload) GetConfigReference() AppSubstrateReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppSubstrateInputUpload) GetConfigReferenceOk() (*AppSubstrateReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppSubstrateInputUpload) SetConfigReference(v AppSubstrateReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppSubstrateInputUpload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetOsType

`func (o *AppSubstrateInputUpload) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AppSubstrateInputUpload) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AppSubstrateInputUpload) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AppSubstrateInputUpload) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetType

`func (o *AppSubstrateInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubstrateInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubstrateInputUpload) SetType(v string)`

SetType sets Type field to given value.


### GetCreateSpec

`func (o *AppSubstrateInputUpload) GetCreateSpec() map[string]interface{}`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *AppSubstrateInputUpload) GetCreateSpecOk() (*map[string]interface{}, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *AppSubstrateInputUpload) SetCreateSpec(v map[string]interface{})`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *AppSubstrateInputUpload) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppSubstrateInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppSubstrateInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppSubstrateInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.


### GetInstancePowerState

`func (o *AppSubstrateInputUpload) GetInstancePowerState() string`

GetInstancePowerState returns the InstancePowerState field if non-nil, zero value otherwise.

### GetInstancePowerStateOk

`func (o *AppSubstrateInputUpload) GetInstancePowerStateOk() (*string, bool)`

GetInstancePowerStateOk returns a tuple with the InstancePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePowerState

`func (o *AppSubstrateInputUpload) SetInstancePowerState(v string)`

SetInstancePowerState sets InstancePowerState field to given value.

### HasInstancePowerState

`func (o *AppSubstrateInputUpload) HasInstancePowerState() bool`

HasInstancePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


