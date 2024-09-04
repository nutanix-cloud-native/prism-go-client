# ExternalSubnetDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**ExternalIpList** | Pointer to **[]string** | List of IP addresses for SNAT or Gateway IP Addresses.  | [optional] 
**GatewayNodeUuidList** | Pointer to **[]string** | List of gateway nodes for the external connectivity. | [optional] 
**ActiveGatewayNode** | Pointer to [**Node**](Node.md) |  | [optional] 

## Methods

### NewExternalSubnetDefStatus

`func NewExternalSubnetDefStatus() *ExternalSubnetDefStatus`

NewExternalSubnetDefStatus instantiates a new ExternalSubnetDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalSubnetDefStatusWithDefaults

`func NewExternalSubnetDefStatusWithDefaults() *ExternalSubnetDefStatus`

NewExternalSubnetDefStatusWithDefaults instantiates a new ExternalSubnetDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalSubnetReference

`func (o *ExternalSubnetDefStatus) GetExternalSubnetReference() SubnetReference`

GetExternalSubnetReference returns the ExternalSubnetReference field if non-nil, zero value otherwise.

### GetExternalSubnetReferenceOk

`func (o *ExternalSubnetDefStatus) GetExternalSubnetReferenceOk() (*SubnetReference, bool)`

GetExternalSubnetReferenceOk returns a tuple with the ExternalSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnetReference

`func (o *ExternalSubnetDefStatus) SetExternalSubnetReference(v SubnetReference)`

SetExternalSubnetReference sets ExternalSubnetReference field to given value.

### HasExternalSubnetReference

`func (o *ExternalSubnetDefStatus) HasExternalSubnetReference() bool`

HasExternalSubnetReference returns a boolean if a field has been set.

### GetExternalIpList

`func (o *ExternalSubnetDefStatus) GetExternalIpList() []string`

GetExternalIpList returns the ExternalIpList field if non-nil, zero value otherwise.

### GetExternalIpListOk

`func (o *ExternalSubnetDefStatus) GetExternalIpListOk() (*[]string, bool)`

GetExternalIpListOk returns a tuple with the ExternalIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpList

`func (o *ExternalSubnetDefStatus) SetExternalIpList(v []string)`

SetExternalIpList sets ExternalIpList field to given value.

### HasExternalIpList

`func (o *ExternalSubnetDefStatus) HasExternalIpList() bool`

HasExternalIpList returns a boolean if a field has been set.

### GetGatewayNodeUuidList

`func (o *ExternalSubnetDefStatus) GetGatewayNodeUuidList() []string`

GetGatewayNodeUuidList returns the GatewayNodeUuidList field if non-nil, zero value otherwise.

### GetGatewayNodeUuidListOk

`func (o *ExternalSubnetDefStatus) GetGatewayNodeUuidListOk() (*[]string, bool)`

GetGatewayNodeUuidListOk returns a tuple with the GatewayNodeUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNodeUuidList

`func (o *ExternalSubnetDefStatus) SetGatewayNodeUuidList(v []string)`

SetGatewayNodeUuidList sets GatewayNodeUuidList field to given value.

### HasGatewayNodeUuidList

`func (o *ExternalSubnetDefStatus) HasGatewayNodeUuidList() bool`

HasGatewayNodeUuidList returns a boolean if a field has been set.

### GetActiveGatewayNode

`func (o *ExternalSubnetDefStatus) GetActiveGatewayNode() Node`

GetActiveGatewayNode returns the ActiveGatewayNode field if non-nil, zero value otherwise.

### GetActiveGatewayNodeOk

`func (o *ExternalSubnetDefStatus) GetActiveGatewayNodeOk() (*Node, bool)`

GetActiveGatewayNodeOk returns a tuple with the ActiveGatewayNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGatewayNode

`func (o *ExternalSubnetDefStatus) SetActiveGatewayNode(v Node)`

SetActiveGatewayNode sets ActiveGatewayNode field to given value.

### HasActiveGatewayNode

`func (o *ExternalSubnetDefStatus) HasActiveGatewayNode() bool`

HasActiveGatewayNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


