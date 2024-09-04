# VmNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NicType** | Pointer to **string** | The type of this NIC. Defaults to NORMAL_NIC. | [optional] 
**UUID** | Pointer to **string** | The NIC&#39;s UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.  | [optional] 
**IpEndpointList** | Pointer to [**[]IpAddress**](IpAddress.md) | IP endpoints for the adapter. Currently, IPv4 addresses are supported.  | [optional] 
**NumQueues** | Pointer to **int32** | The number of tx/rx queue pairs for this NIC.  | [optional] 
**SecondaryIpAddressList** | Pointer to **[]string** | Secondary IPv4 Addresses for this NIC. | [optional] 
**NetworkFunctionNicType** | Pointer to **string** | The type of this Network function NIC. Defaults to INGRESS.  | [optional] 
**NetworkFunctionChainReference** | Pointer to [**NetworkFunctionChainReference**](NetworkFunctionChainReference.md) |  | [optional] 
**VlanMode** | Pointer to **string** | By default, all virtual NICs are created in ACCESS mode, which permits only one VLAN per virtual network. TRUNKED mode allows multiple VLANs on a single VM NIC for network-aware user VMs.  | [optional] 
**MacAddress** | Pointer to **string** | The MAC address for the adapter. | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**Model** | Pointer to **string** | The model of this NIC. | [optional] 
**IsConnected** | Pointer to **bool** | Whether or not the NIC is connected. True by default. | [optional] 
**TrunkedVlanList** | Pointer to **[]int32** | List of networks to trunk if vlan_mode is TRUNKED. If empty and VLAN mode is TRUNKED, all VLANs are trunked.  | [optional] 

## Methods

### NewVmNic

`func NewVmNic() *VmNic`

NewVmNic instantiates a new VmNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmNicWithDefaults

`func NewVmNicWithDefaults() *VmNic`

NewVmNicWithDefaults instantiates a new VmNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNicType

`func (o *VmNic) GetNicType() string`

GetNicType returns the NicType field if non-nil, zero value otherwise.

### GetNicTypeOk

`func (o *VmNic) GetNicTypeOk() (*string, bool)`

GetNicTypeOk returns a tuple with the NicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicType

`func (o *VmNic) SetNicType(v string)`

SetNicType sets NicType field to given value.

### HasNicType

`func (o *VmNic) HasNicType() bool`

HasNicType returns a boolean if a field has been set.

### GetUUID

`func (o *VmNic) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmNic) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmNic) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmNic) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetIpEndpointList

`func (o *VmNic) GetIpEndpointList() []IpAddress`

GetIpEndpointList returns the IpEndpointList field if non-nil, zero value otherwise.

### GetIpEndpointListOk

`func (o *VmNic) GetIpEndpointListOk() (*[]IpAddress, bool)`

GetIpEndpointListOk returns a tuple with the IpEndpointList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEndpointList

`func (o *VmNic) SetIpEndpointList(v []IpAddress)`

SetIpEndpointList sets IpEndpointList field to given value.

### HasIpEndpointList

`func (o *VmNic) HasIpEndpointList() bool`

HasIpEndpointList returns a boolean if a field has been set.

### GetNumQueues

`func (o *VmNic) GetNumQueues() int32`

GetNumQueues returns the NumQueues field if non-nil, zero value otherwise.

### GetNumQueuesOk

`func (o *VmNic) GetNumQueuesOk() (*int32, bool)`

GetNumQueuesOk returns a tuple with the NumQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueues

`func (o *VmNic) SetNumQueues(v int32)`

SetNumQueues sets NumQueues field to given value.

### HasNumQueues

`func (o *VmNic) HasNumQueues() bool`

HasNumQueues returns a boolean if a field has been set.

### GetSecondaryIpAddressList

`func (o *VmNic) GetSecondaryIpAddressList() []string`

GetSecondaryIpAddressList returns the SecondaryIpAddressList field if non-nil, zero value otherwise.

### GetSecondaryIpAddressListOk

`func (o *VmNic) GetSecondaryIpAddressListOk() (*[]string, bool)`

GetSecondaryIpAddressListOk returns a tuple with the SecondaryIpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIpAddressList

`func (o *VmNic) SetSecondaryIpAddressList(v []string)`

SetSecondaryIpAddressList sets SecondaryIpAddressList field to given value.

### HasSecondaryIpAddressList

`func (o *VmNic) HasSecondaryIpAddressList() bool`

HasSecondaryIpAddressList returns a boolean if a field has been set.

### GetNetworkFunctionNicType

`func (o *VmNic) GetNetworkFunctionNicType() string`

GetNetworkFunctionNicType returns the NetworkFunctionNicType field if non-nil, zero value otherwise.

### GetNetworkFunctionNicTypeOk

`func (o *VmNic) GetNetworkFunctionNicTypeOk() (*string, bool)`

GetNetworkFunctionNicTypeOk returns a tuple with the NetworkFunctionNicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionNicType

`func (o *VmNic) SetNetworkFunctionNicType(v string)`

SetNetworkFunctionNicType sets NetworkFunctionNicType field to given value.

### HasNetworkFunctionNicType

`func (o *VmNic) HasNetworkFunctionNicType() bool`

HasNetworkFunctionNicType returns a boolean if a field has been set.

### GetNetworkFunctionChainReference

`func (o *VmNic) GetNetworkFunctionChainReference() NetworkFunctionChainReference`

GetNetworkFunctionChainReference returns the NetworkFunctionChainReference field if non-nil, zero value otherwise.

### GetNetworkFunctionChainReferenceOk

`func (o *VmNic) GetNetworkFunctionChainReferenceOk() (*NetworkFunctionChainReference, bool)`

GetNetworkFunctionChainReferenceOk returns a tuple with the NetworkFunctionChainReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionChainReference

`func (o *VmNic) SetNetworkFunctionChainReference(v NetworkFunctionChainReference)`

SetNetworkFunctionChainReference sets NetworkFunctionChainReference field to given value.

### HasNetworkFunctionChainReference

`func (o *VmNic) HasNetworkFunctionChainReference() bool`

HasNetworkFunctionChainReference returns a boolean if a field has been set.

### GetVlanMode

`func (o *VmNic) GetVlanMode() string`

GetVlanMode returns the VlanMode field if non-nil, zero value otherwise.

### GetVlanModeOk

`func (o *VmNic) GetVlanModeOk() (*string, bool)`

GetVlanModeOk returns a tuple with the VlanMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanMode

`func (o *VmNic) SetVlanMode(v string)`

SetVlanMode sets VlanMode field to given value.

### HasVlanMode

`func (o *VmNic) HasVlanMode() bool`

HasVlanMode returns a boolean if a field has been set.

### GetMacAddress

`func (o *VmNic) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VmNic) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VmNic) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VmNic) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetSubnetReference

`func (o *VmNic) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *VmNic) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *VmNic) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *VmNic) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetModel

`func (o *VmNic) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VmNic) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VmNic) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VmNic) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetIsConnected

`func (o *VmNic) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *VmNic) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *VmNic) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *VmNic) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetTrunkedVlanList

`func (o *VmNic) GetTrunkedVlanList() []int32`

GetTrunkedVlanList returns the TrunkedVlanList field if non-nil, zero value otherwise.

### GetTrunkedVlanListOk

`func (o *VmNic) GetTrunkedVlanListOk() (*[]int32, bool)`

GetTrunkedVlanListOk returns a tuple with the TrunkedVlanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkedVlanList

`func (o *VmNic) SetTrunkedVlanList(v []int32)`

SetTrunkedVlanList sets TrunkedVlanList field to given value.

### HasTrunkedVlanList

`func (o *VmNic) HasTrunkedVlanList() bool`

HasTrunkedVlanList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


