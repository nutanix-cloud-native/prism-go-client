# ExternalSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSubnetReference** | [**SubnetReference**](SubnetReference.md) |  | 
**ExternalIpList** | Pointer to **[]string** | List of IP addresses for SNAT or Gateway IP Addresses.  | [optional] 
**GatewayNodeUuidList** | Pointer to **[]string** | List of gateway nodes for the external connectivity. | [optional] 

## Methods

### NewExternalSubnet

`func NewExternalSubnet(externalSubnetReference SubnetReference, ) *ExternalSubnet`

NewExternalSubnet instantiates a new ExternalSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSubnetWithDefaults

`func NewExternalSubnetWithDefaults() *ExternalSubnet`

NewExternalSubnetWithDefaults instantiates a new ExternalSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSubnetReference

`func (o *ExternalSubnet) GetExternalSubnetReference() SubnetReference`

GetExternalSubnetReference returns the ExternalSubnetReference field if non-nil, zero value otherwise.

### GetExternalSubnetReferenceOk

`func (o *ExternalSubnet) GetExternalSubnetReferenceOk() (*SubnetReference, bool)`

GetExternalSubnetReferenceOk returns a tuple with the ExternalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetReference

`func (o *ExternalSubnet) SetExternalSubnetReference(v SubnetReference)`

SetExternalSubnetReference sets ExternalSubnetReference field to given value.


### GetExternalIpList

`func (o *ExternalSubnet) GetExternalIpList() []string`

GetExternalIpList returns the ExternalIpList field if non-nil, zero value otherwise.

### GetExternalIpListOk

`func (o *ExternalSubnet) GetExternalIpListOk() (*[]string, bool)`

GetExternalIpListOk returns a tuple with the ExternalIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpList

`func (o *ExternalSubnet) SetExternalIpList(v []string)`

SetExternalIpList sets ExternalIpList field to given value.

### HasExternalIpList

`func (o *ExternalSubnet) HasExternalIpList() bool`

HasExternalIpList returns a boolean if a field has been set.

### GetGatewayNodeUuidList

`func (o *ExternalSubnet) GetGatewayNodeUuidList() []string`

GetGatewayNodeUuidList returns the GatewayNodeUuidList field if non-nil, zero value otherwise.

### GetGatewayNodeUuidListOk

`func (o *ExternalSubnet) GetGatewayNodeUuidListOk() (*[]string, bool)`

GetGatewayNodeUuidListOk returns a tuple with the GatewayNodeUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNodeUuidList

`func (o *ExternalSubnet) SetGatewayNodeUuidList(v []string)`

SetGatewayNodeUuidList sets GatewayNodeUuidList field to given value.

### HasGatewayNodeUuidList

`func (o *ExternalSubnet) HasGatewayNodeUuidList() bool`

HasGatewayNodeUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


