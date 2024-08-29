# RoutingPolicyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBidirectional** | Pointer to **bool** | Whether to configure/install policy in reverse direction too (i.e matching traffic from destination to source)  | [optional] [default to false]
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**Destination** | [**NetworkAddress**](NetworkAddress.md) |  | 
**Priority** | **int32** |  | 
**Source** | [**NetworkAddress**](NetworkAddress.md) |  | 
**ProtocolParameters** | Pointer to [**ProtocolParameters**](ProtocolParameters.md) |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**Action** | [**RoutingPolicyAction**](RoutingPolicyAction.md) |  | 
**ProtocolType** | **string** |  | 

## Methods

### NewRoutingPolicyResources

`func NewRoutingPolicyResources(destination NetworkAddress, priority int32, source NetworkAddress, action RoutingPolicyAction, protocolType string, ) *RoutingPolicyResources`

NewRoutingPolicyResources instantiates a new RoutingPolicyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyResourcesWithDefaults

`func NewRoutingPolicyResourcesWithDefaults() *RoutingPolicyResources`

NewRoutingPolicyResourcesWithDefaults instantiates a new RoutingPolicyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBidirectional

`func (o *RoutingPolicyResources) GetIsBidirectional() bool`

GetIsBidirectional returns the IsBidirectional field if non-nil, zero value otherwise.

### GetIsBidirectionalOk

`func (o *RoutingPolicyResources) GetIsBidirectionalOk() (*bool, bool)`

GetIsBidirectionalOk returns a tuple with the IsBidirectional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBidirectional

`func (o *RoutingPolicyResources) SetIsBidirectional(v bool)`

SetIsBidirectional sets IsBidirectional field to given value.

### HasIsBidirectional

`func (o *RoutingPolicyResources) HasIsBidirectional() bool`

HasIsBidirectional returns a boolean if a field has been set.

### GetVpcReference

`func (o *RoutingPolicyResources) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *RoutingPolicyResources) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *RoutingPolicyResources) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *RoutingPolicyResources) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetDestination

`func (o *RoutingPolicyResources) GetDestination() NetworkAddress`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RoutingPolicyResources) GetDestinationOk() (*NetworkAddress, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RoutingPolicyResources) SetDestination(v NetworkAddress)`

SetDestination sets Destination field to given value.


### GetPriority

`func (o *RoutingPolicyResources) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RoutingPolicyResources) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RoutingPolicyResources) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSource

`func (o *RoutingPolicyResources) GetSource() NetworkAddress`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RoutingPolicyResources) GetSourceOk() (*NetworkAddress, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RoutingPolicyResources) SetSource(v NetworkAddress)`

SetSource sets Source field to given value.


### GetProtocolParameters

`func (o *RoutingPolicyResources) GetProtocolParameters() ProtocolParameters`

GetProtocolParameters returns the ProtocolParameters field if non-nil, zero value otherwise.

### GetProtocolParametersOk

`func (o *RoutingPolicyResources) GetProtocolParametersOk() (*ProtocolParameters, bool)`

GetProtocolParametersOk returns a tuple with the ProtocolParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolParameters

`func (o *RoutingPolicyResources) SetProtocolParameters(v ProtocolParameters)`

SetProtocolParameters sets ProtocolParameters field to given value.

### HasProtocolParameters

`func (o *RoutingPolicyResources) HasProtocolParameters() bool`

HasProtocolParameters returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *RoutingPolicyResources) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *RoutingPolicyResources) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *RoutingPolicyResources) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *RoutingPolicyResources) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetAction

`func (o *RoutingPolicyResources) GetAction() RoutingPolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyResources) GetActionOk() (*RoutingPolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyResources) SetAction(v RoutingPolicyAction)`

SetAction sets Action field to given value.


### GetProtocolType

`func (o *RoutingPolicyResources) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *RoutingPolicyResources) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *RoutingPolicyResources) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


