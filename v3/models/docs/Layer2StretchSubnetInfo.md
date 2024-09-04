# Layer2StretchSubnetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSubnet** | Pointer to [**IpSubnetStatus**](IpSubnetStatus.md) |  | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**DefaultGatewayIp** | Pointer to **string** |  | [optional] 
**VlanId** | Pointer to **int32** |  | [optional] 

## Methods

### NewLayer2StretchSubnetInfo

`func NewLayer2StretchSubnetInfo() *Layer2StretchSubnetInfo`

NewLayer2StretchSubnetInfo instantiates a new Layer2StretchSubnetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchSubnetInfoWithDefaults

`func NewLayer2StretchSubnetInfoWithDefaults() *Layer2StretchSubnetInfo`

NewLayer2StretchSubnetInfoWithDefaults instantiates a new Layer2StretchSubnetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSubnet

`func (o *Layer2StretchSubnetInfo) GetIpSubnet() IpSubnetStatus`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *Layer2StretchSubnetInfo) GetIpSubnetOk() (*IpSubnetStatus, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *Layer2StretchSubnetInfo) SetIpSubnet(v IpSubnetStatus)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *Layer2StretchSubnetInfo) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetSubnetReference

`func (o *Layer2StretchSubnetInfo) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *Layer2StretchSubnetInfo) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *Layer2StretchSubnetInfo) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *Layer2StretchSubnetInfo) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetVpcReference

`func (o *Layer2StretchSubnetInfo) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *Layer2StretchSubnetInfo) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *Layer2StretchSubnetInfo) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *Layer2StretchSubnetInfo) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetDefaultGatewayIp

`func (o *Layer2StretchSubnetInfo) GetDefaultGatewayIp() string`

GetDefaultGatewayIp returns the DefaultGatewayIp field if non-nil, zero value otherwise.

### GetDefaultGatewayIpOk

`func (o *Layer2StretchSubnetInfo) GetDefaultGatewayIpOk() (*string, bool)`

GetDefaultGatewayIpOk returns a tuple with the DefaultGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayIp

`func (o *Layer2StretchSubnetInfo) SetDefaultGatewayIp(v string)`

SetDefaultGatewayIp sets DefaultGatewayIp field to given value.

### HasDefaultGatewayIp

`func (o *Layer2StretchSubnetInfo) HasDefaultGatewayIp() bool`

HasDefaultGatewayIp returns a boolean if a field has been set.

### GetVlanId

`func (o *Layer2StretchSubnetInfo) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *Layer2StretchSubnetInfo) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *Layer2StretchSubnetInfo) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *Layer2StretchSubnetInfo) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


