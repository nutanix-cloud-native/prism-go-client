# VmNicOutputStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsIpAddressesList** | Pointer to **[]string** | List of DNS IP addresses. | [optional] 
**NicType** | Pointer to **string** | The type of this NIC. Defaults to NORMAL_NIC. | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**UUID** | Pointer to **string** | The NIC&#39;s UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.  | [optional] 
**IpEndpointList** | Pointer to [**[]IpAddress**](IpAddress.md) | IP endpoints for the adapter. Currently, IPv4 addresses are supported.  | [optional] 
**NetworkFunctionChainReference** | Pointer to [**NetworkFunctionChainReference**](NetworkFunctionChainReference.md) |  | [optional] 
**SecondaryIpAddressList** | Pointer to **[]string** | Secondary IPv4 Addresses for this NIC. | [optional] 
**FloatingIp** | Pointer to **string** | The Floating IP associated with the vnic. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address for the adapter. | [optional] 
**NetworkFunctionNicType** | Pointer to **string** | The type of this Network function NIC. Defaults to INGRESS.  | [optional] 
**VlanMode** | Pointer to **string** | VLAN mode. | [optional] 
**NumQueues** | Pointer to **int32** | The number of tx/rx queue pairs for this NIC.  | [optional] 
**DefaultGatewayAddressList** | Pointer to **[]string** | Default gateway IP addresses. | [optional] 
**DhcpServerIp** | Pointer to **string** | IP address of the DHCP server. | [optional] 
**Model** | Pointer to **string** | The model of this NIC. | [optional] 
**IsConnected** | Pointer to **bool** | Whether or not the NIC is connected. True by default. | [optional] 
**TrunkedVlanList** | Pointer to **[]int32** | List of VLANs to trunk if vlan_mode is TRUNKED. If empty, all VLANs are trunked.  | [optional] 

## Methods

### NewVmNicOutputStatus

`func NewVmNicOutputStatus() *VmNicOutputStatus`

NewVmNicOutputStatus instantiates a new VmNicOutputStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmNicOutputStatusWithDefaults

`func NewVmNicOutputStatusWithDefaults() *VmNicOutputStatus`

NewVmNicOutputStatusWithDefaults instantiates a new VmNicOutputStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsIpAddressesList

`func (o *VmNicOutputStatus) GetDnsIpAddressesList() []string`

GetDnsIpAddressesList returns the DnsIpAddressesList field if non-nil, zero value otherwise.

### GetDnsIpAddressesListOk

`func (o *VmNicOutputStatus) GetDnsIpAddressesListOk() (*[]string, bool)`

GetDnsIpAddressesListOk returns a tuple with the DnsIpAddressesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIpAddressesList

`func (o *VmNicOutputStatus) SetDnsIpAddressesList(v []string)`

SetDnsIpAddressesList sets DnsIpAddressesList field to given value.

### HasDnsIpAddressesList

`func (o *VmNicOutputStatus) HasDnsIpAddressesList() bool`

HasDnsIpAddressesList returns a boolean if a field has been set.

### GetNicType

`func (o *VmNicOutputStatus) GetNicType() string`

GetNicType returns the NicType field if non-nil, zero value otherwise.

### GetNicTypeOk

`func (o *VmNicOutputStatus) GetNicTypeOk() (*string, bool)`

GetNicTypeOk returns a tuple with the NicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicType

`func (o *VmNicOutputStatus) SetNicType(v string)`

SetNicType sets NicType field to given value.

### HasNicType

`func (o *VmNicOutputStatus) HasNicType() bool`

HasNicType returns a boolean if a field has been set.

### GetSubnetReference

`func (o *VmNicOutputStatus) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *VmNicOutputStatus) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *VmNicOutputStatus) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *VmNicOutputStatus) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetUUID

`func (o *VmNicOutputStatus) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmNicOutputStatus) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmNicOutputStatus) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmNicOutputStatus) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetIpEndpointList

`func (o *VmNicOutputStatus) GetIpEndpointList() []IpAddress`

GetIpEndpointList returns the IpEndpointList field if non-nil, zero value otherwise.

### GetIpEndpointListOk

`func (o *VmNicOutputStatus) GetIpEndpointListOk() (*[]IpAddress, bool)`

GetIpEndpointListOk returns a tuple with the IpEndpointList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEndpointList

`func (o *VmNicOutputStatus) SetIpEndpointList(v []IpAddress)`

SetIpEndpointList sets IpEndpointList field to given value.

### HasIpEndpointList

`func (o *VmNicOutputStatus) HasIpEndpointList() bool`

HasIpEndpointList returns a boolean if a field has been set.

### GetNetworkFunctionChainReference

`func (o *VmNicOutputStatus) GetNetworkFunctionChainReference() NetworkFunctionChainReference`

GetNetworkFunctionChainReference returns the NetworkFunctionChainReference field if non-nil, zero value otherwise.

### GetNetworkFunctionChainReferenceOk

`func (o *VmNicOutputStatus) GetNetworkFunctionChainReferenceOk() (*NetworkFunctionChainReference, bool)`

GetNetworkFunctionChainReferenceOk returns a tuple with the NetworkFunctionChainReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionChainReference

`func (o *VmNicOutputStatus) SetNetworkFunctionChainReference(v NetworkFunctionChainReference)`

SetNetworkFunctionChainReference sets NetworkFunctionChainReference field to given value.

### HasNetworkFunctionChainReference

`func (o *VmNicOutputStatus) HasNetworkFunctionChainReference() bool`

HasNetworkFunctionChainReference returns a boolean if a field has been set.

### GetSecondaryIpAddressList

`func (o *VmNicOutputStatus) GetSecondaryIpAddressList() []string`

GetSecondaryIpAddressList returns the SecondaryIpAddressList field if non-nil, zero value otherwise.

### GetSecondaryIpAddressListOk

`func (o *VmNicOutputStatus) GetSecondaryIpAddressListOk() (*[]string, bool)`

GetSecondaryIpAddressListOk returns a tuple with the SecondaryIpAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIpAddressList

`func (o *VmNicOutputStatus) SetSecondaryIpAddressList(v []string)`

SetSecondaryIpAddressList sets SecondaryIpAddressList field to given value.

### HasSecondaryIpAddressList

`func (o *VmNicOutputStatus) HasSecondaryIpAddressList() bool`

HasSecondaryIpAddressList returns a boolean if a field has been set.

### GetFloatingIp

`func (o *VmNicOutputStatus) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *VmNicOutputStatus) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *VmNicOutputStatus) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *VmNicOutputStatus) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *VmNicOutputStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VmNicOutputStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VmNicOutputStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VmNicOutputStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetworkFunctionNicType

`func (o *VmNicOutputStatus) GetNetworkFunctionNicType() string`

GetNetworkFunctionNicType returns the NetworkFunctionNicType field if non-nil, zero value otherwise.

### GetNetworkFunctionNicTypeOk

`func (o *VmNicOutputStatus) GetNetworkFunctionNicTypeOk() (*string, bool)`

GetNetworkFunctionNicTypeOk returns a tuple with the NetworkFunctionNicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFunctionNicType

`func (o *VmNicOutputStatus) SetNetworkFunctionNicType(v string)`

SetNetworkFunctionNicType sets NetworkFunctionNicType field to given value.

### HasNetworkFunctionNicType

`func (o *VmNicOutputStatus) HasNetworkFunctionNicType() bool`

HasNetworkFunctionNicType returns a boolean if a field has been set.

### GetVlanMode

`func (o *VmNicOutputStatus) GetVlanMode() string`

GetVlanMode returns the VlanMode field if non-nil, zero value otherwise.

### GetVlanModeOk

`func (o *VmNicOutputStatus) GetVlanModeOk() (*string, bool)`

GetVlanModeOk returns a tuple with the VlanMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanMode

`func (o *VmNicOutputStatus) SetVlanMode(v string)`

SetVlanMode sets VlanMode field to given value.

### HasVlanMode

`func (o *VmNicOutputStatus) HasVlanMode() bool`

HasVlanMode returns a boolean if a field has been set.

### GetNumQueues

`func (o *VmNicOutputStatus) GetNumQueues() int32`

GetNumQueues returns the NumQueues field if non-nil, zero value otherwise.

### GetNumQueuesOk

`func (o *VmNicOutputStatus) GetNumQueuesOk() (*int32, bool)`

GetNumQueuesOk returns a tuple with the NumQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueues

`func (o *VmNicOutputStatus) SetNumQueues(v int32)`

SetNumQueues sets NumQueues field to given value.

### HasNumQueues

`func (o *VmNicOutputStatus) HasNumQueues() bool`

HasNumQueues returns a boolean if a field has been set.

### GetDefaultGatewayAddressList

`func (o *VmNicOutputStatus) GetDefaultGatewayAddressList() []string`

GetDefaultGatewayAddressList returns the DefaultGatewayAddressList field if non-nil, zero value otherwise.

### GetDefaultGatewayAddressListOk

`func (o *VmNicOutputStatus) GetDefaultGatewayAddressListOk() (*[]string, bool)`

GetDefaultGatewayAddressListOk returns a tuple with the DefaultGatewayAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayAddressList

`func (o *VmNicOutputStatus) SetDefaultGatewayAddressList(v []string)`

SetDefaultGatewayAddressList sets DefaultGatewayAddressList field to given value.

### HasDefaultGatewayAddressList

`func (o *VmNicOutputStatus) HasDefaultGatewayAddressList() bool`

HasDefaultGatewayAddressList returns a boolean if a field has been set.

### GetDhcpServerIp

`func (o *VmNicOutputStatus) GetDhcpServerIp() string`

GetDhcpServerIp returns the DhcpServerIp field if non-nil, zero value otherwise.

### GetDhcpServerIpOk

`func (o *VmNicOutputStatus) GetDhcpServerIpOk() (*string, bool)`

GetDhcpServerIpOk returns a tuple with the DhcpServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerIp

`func (o *VmNicOutputStatus) SetDhcpServerIp(v string)`

SetDhcpServerIp sets DhcpServerIp field to given value.

### HasDhcpServerIp

`func (o *VmNicOutputStatus) HasDhcpServerIp() bool`

HasDhcpServerIp returns a boolean if a field has been set.

### GetModel

`func (o *VmNicOutputStatus) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VmNicOutputStatus) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VmNicOutputStatus) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VmNicOutputStatus) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetIsConnected

`func (o *VmNicOutputStatus) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *VmNicOutputStatus) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *VmNicOutputStatus) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *VmNicOutputStatus) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetTrunkedVlanList

`func (o *VmNicOutputStatus) GetTrunkedVlanList() []int32`

GetTrunkedVlanList returns the TrunkedVlanList field if non-nil, zero value otherwise.

### GetTrunkedVlanListOk

`func (o *VmNicOutputStatus) GetTrunkedVlanListOk() (*[]int32, bool)`

GetTrunkedVlanListOk returns a tuple with the TrunkedVlanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkedVlanList

`func (o *VmNicOutputStatus) SetTrunkedVlanList(v []int32)`

SetTrunkedVlanList sets TrunkedVlanList field to given value.

### HasTrunkedVlanList

`func (o *VmNicOutputStatus) HasTrunkedVlanList() bool`

HasTrunkedVlanList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


