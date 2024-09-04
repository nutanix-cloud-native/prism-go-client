# AppSubstrateResponseDownload

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
**AccountReference** | Pointer to [**AccountReferenceUpload**](AccountReferenceUpload.md) |  | [optional] 
**Name** | **string** |  | 
**InstanceName** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**ReadinessProbe** | Pointer to [**AppSubstrateReadinessProbe**](AppSubstrateReadinessProbe.md) |  | [optional] 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**ConfigReference** | Pointer to [**AppSubstrateReferenceUpload**](AppSubstrateReferenceUpload.md) |  | [optional] 
**OsType** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**CreateSpec** | Pointer to **map[string]interface{}** | Spec of the substrate | [optional] 
**VariableList** | [**[]AppVariableResponse**](AppVariableResponse.md) | List of variables | 
**InstancePowerState** | Pointer to **string** |  | [optional] 

## Methods

### NewAppSubstrateResponseDownload

`func NewAppSubstrateResponseDownload(actionList []AppActionResponse, uUID string, name string, state string, type_ string, variableList []AppVariableResponse, ) *AppSubstrateResponseDownload`

NewAppSubstrateResponseDownload instantiates a new AppSubstrateResponseDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubstrateResponseDownloadWithDefaults

`func NewAppSubstrateResponseDownloadWithDefaults() *AppSubstrateResponseDownload`

NewAppSubstrateResponseDownloadWithDefaults instantiates a new AppSubstrateResponseDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformData

`func (o *AppSubstrateResponseDownload) GetPlatformData() string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *AppSubstrateResponseDownload) GetPlatformDataOk() (*string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *AppSubstrateResponseDownload) SetPlatformData(v string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *AppSubstrateResponseDownload) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *AppSubstrateResponseDownload) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *AppSubstrateResponseDownload) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *AppSubstrateResponseDownload) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *AppSubstrateResponseDownload) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetDescription

`func (o *AppSubstrateResponseDownload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSubstrateResponseDownload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSubstrateResponseDownload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSubstrateResponseDownload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppSubstrateResponseDownload) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppSubstrateResponseDownload) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppSubstrateResponseDownload) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.


### GetElementList

`func (o *AppSubstrateResponseDownload) GetElementList() []AppSubstrateElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppSubstrateResponseDownload) GetElementListOk() (*[]AppSubstrateElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppSubstrateResponseDownload) SetElementList(v []AppSubstrateElement)`

SetElementList sets ElementList field to given value.

### HasElementList

`func (o *AppSubstrateResponseDownload) HasElementList() bool`

HasElementList returns a boolean if a field has been set.

### GetInstanceId

`func (o *AppSubstrateResponseDownload) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AppSubstrateResponseDownload) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AppSubstrateResponseDownload) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AppSubstrateResponseDownload) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetUUID

`func (o *AppSubstrateResponseDownload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppSubstrateResponseDownload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppSubstrateResponseDownload) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetAccountReference

`func (o *AppSubstrateResponseDownload) GetAccountReference() AccountReferenceUpload`

GetAccountReference returns the AccountReference field if non-nil, zero value otherwise.

### GetAccountReferenceOk

`func (o *AppSubstrateResponseDownload) GetAccountReferenceOk() (*AccountReferenceUpload, bool)`

GetAccountReferenceOk returns a tuple with the AccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReference

`func (o *AppSubstrateResponseDownload) SetAccountReference(v AccountReferenceUpload)`

SetAccountReference sets AccountReference field to given value.

### HasAccountReference

`func (o *AppSubstrateResponseDownload) HasAccountReference() bool`

HasAccountReference returns a boolean if a field has been set.

### GetName

`func (o *AppSubstrateResponseDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppSubstrateResponseDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppSubstrateResponseDownload) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceName

`func (o *AppSubstrateResponseDownload) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *AppSubstrateResponseDownload) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *AppSubstrateResponseDownload) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *AppSubstrateResponseDownload) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetState

`func (o *AppSubstrateResponseDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppSubstrateResponseDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppSubstrateResponseDownload) SetState(v string)`

SetState sets State field to given value.


### GetReadinessProbe

`func (o *AppSubstrateResponseDownload) GetReadinessProbe() AppSubstrateReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *AppSubstrateResponseDownload) GetReadinessProbeOk() (*AppSubstrateReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *AppSubstrateResponseDownload) SetReadinessProbe(v AppSubstrateReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *AppSubstrateResponseDownload) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetEditables

`func (o *AppSubstrateResponseDownload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppSubstrateResponseDownload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppSubstrateResponseDownload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppSubstrateResponseDownload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetMessageList

`func (o *AppSubstrateResponseDownload) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppSubstrateResponseDownload) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppSubstrateResponseDownload) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppSubstrateResponseDownload) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppSubstrateResponseDownload) GetConfigReference() AppSubstrateReferenceUpload`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppSubstrateResponseDownload) GetConfigReferenceOk() (*AppSubstrateReferenceUpload, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppSubstrateResponseDownload) SetConfigReference(v AppSubstrateReferenceUpload)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppSubstrateResponseDownload) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetOsType

`func (o *AppSubstrateResponseDownload) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AppSubstrateResponseDownload) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AppSubstrateResponseDownload) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AppSubstrateResponseDownload) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetType

`func (o *AppSubstrateResponseDownload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubstrateResponseDownload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubstrateResponseDownload) SetType(v string)`

SetType sets Type field to given value.


### GetCreateSpec

`func (o *AppSubstrateResponseDownload) GetCreateSpec() map[string]interface{}`

GetCreateSpec returns the CreateSpec field if non-nil, zero value otherwise.

### GetCreateSpecOk

`func (o *AppSubstrateResponseDownload) GetCreateSpecOk() (*map[string]interface{}, bool)`

GetCreateSpecOk returns a tuple with the CreateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpec

`func (o *AppSubstrateResponseDownload) SetCreateSpec(v map[string]interface{})`

SetCreateSpec sets CreateSpec field to given value.

### HasCreateSpec

`func (o *AppSubstrateResponseDownload) HasCreateSpec() bool`

HasCreateSpec returns a boolean if a field has been set.

### GetVariableList

`func (o *AppSubstrateResponseDownload) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppSubstrateResponseDownload) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppSubstrateResponseDownload) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.


### GetInstancePowerState

`func (o *AppSubstrateResponseDownload) GetInstancePowerState() string`

GetInstancePowerState returns the InstancePowerState field if non-nil, zero value otherwise.

### GetInstancePowerStateOk

`func (o *AppSubstrateResponseDownload) GetInstancePowerStateOk() (*string, bool)`

GetInstancePowerStateOk returns a tuple with the InstancePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePowerState

`func (o *AppSubstrateResponseDownload) SetInstancePowerState(v string)`

SetInstancePowerState sets InstancePowerState field to given value.

### HasInstancePowerState

`func (o *AppSubstrateResponseDownload) HasInstancePowerState() bool`

HasInstancePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


