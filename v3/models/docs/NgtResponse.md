# NgtResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInfoList** | Pointer to [**[]NetworkConfiguration**](NetworkConfiguration.md) | List of network configuration of the VMs. | [optional] 
**UUID** | **string** | UUID | 
**VmUuid** | **string** | VM UUID | 
**ScriptsExecutable** | Pointer to **[]bool** | Executable scripts in the VM. | [optional] 
**ScriptList** | Pointer to **[]string** | Scripts present in the VM. | [optional] 

## Methods

### NewNgtResponse

`func NewNgtResponse(uUID string, vmUuid string, ) *NgtResponse`

NewNgtResponse instantiates a new NgtResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtResponseWithDefaults

`func NewNgtResponseWithDefaults() *NgtResponse`

NewNgtResponseWithDefaults instantiates a new NgtResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInfoList

`func (o *NgtResponse) GetNetworkInfoList() []NetworkConfiguration`

GetNetworkInfoList returns the NetworkInfoList field if non-nil, zero value otherwise.

### GetNetworkInfoListOk

`func (o *NgtResponse) GetNetworkInfoListOk() (*[]NetworkConfiguration, bool)`

GetNetworkInfoListOk returns a tuple with the NetworkInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfoList

`func (o *NgtResponse) SetNetworkInfoList(v []NetworkConfiguration)`

SetNetworkInfoList sets NetworkInfoList field to given value.

### HasNetworkInfoList

`func (o *NgtResponse) HasNetworkInfoList() bool`

HasNetworkInfoList returns a boolean if a field has been set.

### GetUUID

`func (o *NgtResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *NgtResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *NgtResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetVmUuid

`func (o *NgtResponse) GetVmUuid() string`

GetVmUuid returns the VmUuid field if non-nil, zero value otherwise.

### GetVmUuidOk

`func (o *NgtResponse) GetVmUuidOk() (*string, bool)`

GetVmUuidOk returns a tuple with the VmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUuid

`func (o *NgtResponse) SetVmUuid(v string)`

SetVmUuid sets VmUuid field to given value.


### GetScriptsExecutable

`func (o *NgtResponse) GetScriptsExecutable() []bool`

GetScriptsExecutable returns the ScriptsExecutable field if non-nil, zero value otherwise.

### GetScriptsExecutableOk

`func (o *NgtResponse) GetScriptsExecutableOk() (*[]bool, bool)`

GetScriptsExecutableOk returns a tuple with the ScriptsExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptsExecutable

`func (o *NgtResponse) SetScriptsExecutable(v []bool)`

SetScriptsExecutable sets ScriptsExecutable field to given value.

### HasScriptsExecutable

`func (o *NgtResponse) HasScriptsExecutable() bool`

HasScriptsExecutable returns a boolean if a field has been set.

### GetScriptList

`func (o *NgtResponse) GetScriptList() []string`

GetScriptList returns the ScriptList field if non-nil, zero value otherwise.

### GetScriptListOk

`func (o *NgtResponse) GetScriptListOk() (*[]string, bool)`

GetScriptListOk returns a tuple with the ScriptList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptList

`func (o *NgtResponse) SetScriptList(v []string)`

SetScriptList sets ScriptList field to given value.

### HasScriptList

`func (o *NgtResponse) HasScriptList() bool`

HasScriptList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


