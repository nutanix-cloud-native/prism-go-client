# RouteNexthopStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectConnectVirtualInterfaceReference** | Pointer to [**DirectConnectVirtualInterfaceReference**](DirectConnectVirtualInterfaceReference.md) |  | [optional] 
**ExternalSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**NexthopIpAddress** | Pointer to **string** | Nexthop IP address of this route. | [optional] 
**LocalSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**VpnConnectionReference** | Pointer to [**VpnConnectionReference**](VpnConnectionReference.md) |  | [optional] 

## Methods

### NewRouteNexthopStatus

`func NewRouteNexthopStatus() *RouteNexthopStatus`

NewRouteNexthopStatus instantiates a new RouteNexthopStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteNexthopStatusWithDefaults

`func NewRouteNexthopStatusWithDefaults() *RouteNexthopStatus`

NewRouteNexthopStatusWithDefaults instantiates a new RouteNexthopStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectConnectVirtualInterfaceReference

`func (o *RouteNexthopStatus) GetDirectConnectVirtualInterfaceReference() DirectConnectVirtualInterfaceReference`

GetDirectConnectVirtualInterfaceReference returns the DirectConnectVirtualInterfaceReference field if non-nil, zero value otherwise.

### GetDirectConnectVirtualInterfaceReferenceOk

`func (o *RouteNexthopStatus) GetDirectConnectVirtualInterfaceReferenceOk() (*DirectConnectVirtualInterfaceReference, bool)`

GetDirectConnectVirtualInterfaceReferenceOk returns a tuple with the DirectConnectVirtualInterfaceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectVirtualInterfaceReference

`func (o *RouteNexthopStatus) SetDirectConnectVirtualInterfaceReference(v DirectConnectVirtualInterfaceReference)`

SetDirectConnectVirtualInterfaceReference sets DirectConnectVirtualInterfaceReference field to given value.

### HasDirectConnectVirtualInterfaceReference

`func (o *RouteNexthopStatus) HasDirectConnectVirtualInterfaceReference() bool`

HasDirectConnectVirtualInterfaceReference returns a boolean if a field has been set.

### GetExternalSubnetReference

`func (o *RouteNexthopStatus) GetExternalSubnetReference() SubnetReference`

GetExternalSubnetReference returns the ExternalSubnetReference field if non-nil, zero value otherwise.

### GetExternalSubnetReferenceOk

`func (o *RouteNexthopStatus) GetExternalSubnetReferenceOk() (*SubnetReference, bool)`

GetExternalSubnetReferenceOk returns a tuple with the ExternalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetReference

`func (o *RouteNexthopStatus) SetExternalSubnetReference(v SubnetReference)`

SetExternalSubnetReference sets ExternalSubnetReference field to given value.

### HasExternalSubnetReference

`func (o *RouteNexthopStatus) HasExternalSubnetReference() bool`

HasExternalSubnetReference returns a boolean if a field has been set.

### GetNexthopIpAddress

`func (o *RouteNexthopStatus) GetNexthopIpAddress() string`

GetNexthopIpAddress returns the NexthopIpAddress field if non-nil, zero value otherwise.

### GetNexthopIpAddressOk

`func (o *RouteNexthopStatus) GetNexthopIpAddressOk() (*string, bool)`

GetNexthopIpAddressOk returns a tuple with the NexthopIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopIpAddress

`func (o *RouteNexthopStatus) SetNexthopIpAddress(v string)`

SetNexthopIpAddress sets NexthopIpAddress field to given value.

### HasNexthopIpAddress

`func (o *RouteNexthopStatus) HasNexthopIpAddress() bool`

HasNexthopIpAddress returns a boolean if a field has been set.

### GetLocalSubnetReference

`func (o *RouteNexthopStatus) GetLocalSubnetReference() SubnetReference`

GetLocalSubnetReference returns the LocalSubnetReference field if non-nil, zero value otherwise.

### GetLocalSubnetReferenceOk

`func (o *RouteNexthopStatus) GetLocalSubnetReferenceOk() (*SubnetReference, bool)`

GetLocalSubnetReferenceOk returns a tuple with the LocalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubnetReference

`func (o *RouteNexthopStatus) SetLocalSubnetReference(v SubnetReference)`

SetLocalSubnetReference sets LocalSubnetReference field to given value.

### HasLocalSubnetReference

`func (o *RouteNexthopStatus) HasLocalSubnetReference() bool`

HasLocalSubnetReference returns a boolean if a field has been set.

### GetVpnConnectionReference

`func (o *RouteNexthopStatus) GetVpnConnectionReference() VpnConnectionReference`

GetVpnConnectionReference returns the VpnConnectionReference field if non-nil, zero value otherwise.

### GetVpnConnectionReferenceOk

`func (o *RouteNexthopStatus) GetVpnConnectionReferenceOk() (*VpnConnectionReference, bool)`

GetVpnConnectionReferenceOk returns a tuple with the VpnConnectionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionReference

`func (o *RouteNexthopStatus) SetVpnConnectionReference(v VpnConnectionReference)`

SetVpnConnectionReference sets VpnConnectionReference field to given value.

### HasVpnConnectionReference

`func (o *RouteNexthopStatus) HasVpnConnectionReference() bool`

HasVpnConnectionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


