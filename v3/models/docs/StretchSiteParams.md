# StretchSiteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StretchSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**VpnInterfaceIp** | Pointer to **string** |  | [optional] 
**VpnConnectionReference** | Pointer to [**VpnConnectionReference**](VpnConnectionReference.md) |  | [optional] 
**DefaultGatewayIp** | Pointer to **string** |  | [optional] 
**SubnetPrefixLength** | Pointer to **int32** | Prefix length of the subnet being stretched. | [optional] 
**PcClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewStretchSiteParams

`func NewStretchSiteParams() *StretchSiteParams`

NewStretchSiteParams instantiates a new StretchSiteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStretchSiteParamsWithDefaults

`func NewStretchSiteParamsWithDefaults() *StretchSiteParams`

NewStretchSiteParamsWithDefaults instantiates a new StretchSiteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStretchSubnetReference

`func (o *StretchSiteParams) GetStretchSubnetReference() SubnetReference`

GetStretchSubnetReference returns the StretchSubnetReference field if non-nil, zero value otherwise.

### GetStretchSubnetReferenceOk

`func (o *StretchSiteParams) GetStretchSubnetReferenceOk() (*SubnetReference, bool)`

GetStretchSubnetReferenceOk returns a tuple with the StretchSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStretchSubnetReference

`func (o *StretchSiteParams) SetStretchSubnetReference(v SubnetReference)`

SetStretchSubnetReference sets StretchSubnetReference field to given value.

### HasStretchSubnetReference

`func (o *StretchSiteParams) HasStretchSubnetReference() bool`

HasStretchSubnetReference returns a boolean if a field has been set.

### GetVpnInterfaceIp

`func (o *StretchSiteParams) GetVpnInterfaceIp() string`

GetVpnInterfaceIp returns the VpnInterfaceIp field if non-nil, zero value otherwise.

### GetVpnInterfaceIpOk

`func (o *StretchSiteParams) GetVpnInterfaceIpOk() (*string, bool)`

GetVpnInterfaceIpOk returns a tuple with the VpnInterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnInterfaceIp

`func (o *StretchSiteParams) SetVpnInterfaceIp(v string)`

SetVpnInterfaceIp sets VpnInterfaceIp field to given value.

### HasVpnInterfaceIp

`func (o *StretchSiteParams) HasVpnInterfaceIp() bool`

HasVpnInterfaceIp returns a boolean if a field has been set.

### GetVpnConnectionReference

`func (o *StretchSiteParams) GetVpnConnectionReference() VpnConnectionReference`

GetVpnConnectionReference returns the VpnConnectionReference field if non-nil, zero value otherwise.

### GetVpnConnectionReferenceOk

`func (o *StretchSiteParams) GetVpnConnectionReferenceOk() (*VpnConnectionReference, bool)`

GetVpnConnectionReferenceOk returns a tuple with the VpnConnectionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnectionReference

`func (o *StretchSiteParams) SetVpnConnectionReference(v VpnConnectionReference)`

SetVpnConnectionReference sets VpnConnectionReference field to given value.

### HasVpnConnectionReference

`func (o *StretchSiteParams) HasVpnConnectionReference() bool`

HasVpnConnectionReference returns a boolean if a field has been set.

### GetDefaultGatewayIp

`func (o *StretchSiteParams) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *StretchSiteParams) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *StretchSiteParams) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *StretchSiteParams) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### GetSubnetPrefixLength

`func (o *StretchSiteParams) GetSubnetPrefixLength() int32`

GetSubnetPrefixLength returns the SubnetPrefixLength field if non-nil, zero value otherwise.

### GetSubnetPrefixLengthOk

`func (o *StretchSiteParams) GetSubnetPrefixLengthOk() (*int32, bool)`

GetSubnetPrefixLengthOk returns a tuple with the SubnetPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPrefixLength

`func (o *StretchSiteParams) SetSubnetPrefixLength(v int32)`

SetSubnetPrefixLength sets SubnetPrefixLength field to given value.

### HasSubnetPrefixLength

`func (o *StretchSiteParams) HasSubnetPrefixLength() bool`

HasSubnetPrefixLength returns a boolean if a field has been set.

### GetPcClusterReference

`func (o *StretchSiteParams) GetPcClusterReference() ClusterReference`

GetPcClusterReference returns the PcClusterReference field if non-nil, zero value otherwise.

### GetPcClusterReferenceOk

`func (o *StretchSiteParams) GetPcClusterReferenceOk() (*ClusterReference, bool)`

GetPcClusterReferenceOk returns a tuple with the PcClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcClusterReference

`func (o *StretchSiteParams) SetPcClusterReference(v ClusterReference)`

SetPcClusterReference sets PcClusterReference field to given value.

### HasPcClusterReference

`func (o *StretchSiteParams) HasPcClusterReference() bool`

HasPcClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


