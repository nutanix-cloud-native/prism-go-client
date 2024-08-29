# RoutingPolicyResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceIpCountersReverseDirection** | Pointer to [**[]ServiceIpCounters**](ServiceIpCounters.md) | Policy counters for each service IP for reverse direction reroute routing policy. Applicable only if is_bidirectional is true.  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** | Error message describing why the routing policy is inactive.  | [optional] 
**RoutingPolicyCountersReverseDirection** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 
**Destination** | Pointer to [**NetworkAddressStatus**](NetworkAddressStatus.md) |  | [optional] 
**RerouteFallbackCounters** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 
**ServiceIpCounters** | Pointer to [**[]ServiceIpCounters**](ServiceIpCounters.md) | Policy counters for each service IP. | [optional] 
**IsBidirectional** | Pointer to **bool** | Whether to configure/install policy in reverse direction too (i.e matching traffic from destination to source)  | [optional] [default to false]
**Source** | Pointer to [**NetworkAddressStatus**](NetworkAddressStatus.md) |  | [optional] 
**RoutingPolicyCounters** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 
**ProtocolParameters** | Pointer to [**ProtocolParametersStatus**](ProtocolParametersStatus.md) |  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**RerouteFallbackCountersReverseDirection** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 
**Action** | Pointer to [**RoutingPolicyActionStatus**](RoutingPolicyActionStatus.md) |  | [optional] 
**ProtocolType** | Pointer to **string** | The IP protocol type of traffic that is entering the router.  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewRoutingPolicyResourcesDefStatus

`func NewRoutingPolicyResourcesDefStatus() *RoutingPolicyResourcesDefStatus`

NewRoutingPolicyResourcesDefStatus instantiates a new RoutingPolicyResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyResourcesDefStatusWithDefaults

`func NewRoutingPolicyResourcesDefStatusWithDefaults() *RoutingPolicyResourcesDefStatus`

NewRoutingPolicyResourcesDefStatusWithDefaults instantiates a new RoutingPolicyResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceIpCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) GetServiceIpCountersReverseDirection() []ServiceIpCounters`

GetServiceIpCountersReverseDirection returns the ServiceIpCountersReverseDirection field if non-nil, zero value otherwise.

### GetServiceIpCountersReverseDirectionOk

`func (o *RoutingPolicyResourcesDefStatus) GetServiceIpCountersReverseDirectionOk() (*[]ServiceIpCounters, bool)`

GetServiceIpCountersReverseDirectionOk returns a tuple with the ServiceIpCountersReverseDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) SetServiceIpCountersReverseDirection(v []ServiceIpCounters)`

SetServiceIpCountersReverseDirection sets ServiceIpCountersReverseDirection field to given value.

### HasServiceIpCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) HasServiceIpCountersReverseDirection() bool`

HasServiceIpCountersReverseDirection returns a boolean if a field has been set.

### GetVpcReference

`func (o *RoutingPolicyResourcesDefStatus) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *RoutingPolicyResourcesDefStatus) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *RoutingPolicyResourcesDefStatus) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *RoutingPolicyResourcesDefStatus) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetErrorMessage

`func (o *RoutingPolicyResourcesDefStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *RoutingPolicyResourcesDefStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *RoutingPolicyResourcesDefStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *RoutingPolicyResourcesDefStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRoutingPolicyCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) GetRoutingPolicyCountersReverseDirection() RoutingPolicyCounters`

GetRoutingPolicyCountersReverseDirection returns the RoutingPolicyCountersReverseDirection field if non-nil, zero value otherwise.

### GetRoutingPolicyCountersReverseDirectionOk

`func (o *RoutingPolicyResourcesDefStatus) GetRoutingPolicyCountersReverseDirectionOk() (*RoutingPolicyCounters, bool)`

GetRoutingPolicyCountersReverseDirectionOk returns a tuple with the RoutingPolicyCountersReverseDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) SetRoutingPolicyCountersReverseDirection(v RoutingPolicyCounters)`

SetRoutingPolicyCountersReverseDirection sets RoutingPolicyCountersReverseDirection field to given value.

### HasRoutingPolicyCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) HasRoutingPolicyCountersReverseDirection() bool`

HasRoutingPolicyCountersReverseDirection returns a boolean if a field has been set.

### GetDestination

`func (o *RoutingPolicyResourcesDefStatus) GetDestination() NetworkAddressStatus`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RoutingPolicyResourcesDefStatus) GetDestinationOk() (*NetworkAddressStatus, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RoutingPolicyResourcesDefStatus) SetDestination(v NetworkAddressStatus)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *RoutingPolicyResourcesDefStatus) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetRerouteFallbackCounters

`func (o *RoutingPolicyResourcesDefStatus) GetRerouteFallbackCounters() RoutingPolicyCounters`

GetRerouteFallbackCounters returns the RerouteFallbackCounters field if non-nil, zero value otherwise.

### GetRerouteFallbackCountersOk

`func (o *RoutingPolicyResourcesDefStatus) GetRerouteFallbackCountersOk() (*RoutingPolicyCounters, bool)`

GetRerouteFallbackCountersOk returns a tuple with the RerouteFallbackCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerouteFallbackCounters

`func (o *RoutingPolicyResourcesDefStatus) SetRerouteFallbackCounters(v RoutingPolicyCounters)`

SetRerouteFallbackCounters sets RerouteFallbackCounters field to given value.

### HasRerouteFallbackCounters

`func (o *RoutingPolicyResourcesDefStatus) HasRerouteFallbackCounters() bool`

HasRerouteFallbackCounters returns a boolean if a field has been set.

### GetServiceIpCounters

`func (o *RoutingPolicyResourcesDefStatus) GetServiceIpCounters() []ServiceIpCounters`

GetServiceIpCounters returns the ServiceIpCounters field if non-nil, zero value otherwise.

### GetServiceIpCountersOk

`func (o *RoutingPolicyResourcesDefStatus) GetServiceIpCountersOk() (*[]ServiceIpCounters, bool)`

GetServiceIpCountersOk returns a tuple with the ServiceIpCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpCounters

`func (o *RoutingPolicyResourcesDefStatus) SetServiceIpCounters(v []ServiceIpCounters)`

SetServiceIpCounters sets ServiceIpCounters field to given value.

### HasServiceIpCounters

`func (o *RoutingPolicyResourcesDefStatus) HasServiceIpCounters() bool`

HasServiceIpCounters returns a boolean if a field has been set.

### GetIsBidirectional

`func (o *RoutingPolicyResourcesDefStatus) GetIsBidirectional() bool`

GetIsBidirectional returns the IsBidirectional field if non-nil, zero value otherwise.

### GetIsBidirectionalOk

`func (o *RoutingPolicyResourcesDefStatus) GetIsBidirectionalOk() (*bool, bool)`

GetIsBidirectionalOk returns a tuple with the IsBidirectional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBidirectional

`func (o *RoutingPolicyResourcesDefStatus) SetIsBidirectional(v bool)`

SetIsBidirectional sets IsBidirectional field to given value.

### HasIsBidirectional

`func (o *RoutingPolicyResourcesDefStatus) HasIsBidirectional() bool`

HasIsBidirectional returns a boolean if a field has been set.

### GetSource

`func (o *RoutingPolicyResourcesDefStatus) GetSource() NetworkAddressStatus`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RoutingPolicyResourcesDefStatus) GetSourceOk() (*NetworkAddressStatus, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RoutingPolicyResourcesDefStatus) SetSource(v NetworkAddressStatus)`

SetSource sets Source field to given value.

### HasSource

`func (o *RoutingPolicyResourcesDefStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRoutingPolicyCounters

`func (o *RoutingPolicyResourcesDefStatus) GetRoutingPolicyCounters() RoutingPolicyCounters`

GetRoutingPolicyCounters returns the RoutingPolicyCounters field if non-nil, zero value otherwise.

### GetRoutingPolicyCountersOk

`func (o *RoutingPolicyResourcesDefStatus) GetRoutingPolicyCountersOk() (*RoutingPolicyCounters, bool)`

GetRoutingPolicyCountersOk returns a tuple with the RoutingPolicyCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyCounters

`func (o *RoutingPolicyResourcesDefStatus) SetRoutingPolicyCounters(v RoutingPolicyCounters)`

SetRoutingPolicyCounters sets RoutingPolicyCounters field to given value.

### HasRoutingPolicyCounters

`func (o *RoutingPolicyResourcesDefStatus) HasRoutingPolicyCounters() bool`

HasRoutingPolicyCounters returns a boolean if a field has been set.

### GetProtocolParameters

`func (o *RoutingPolicyResourcesDefStatus) GetProtocolParameters() ProtocolParametersStatus`

GetProtocolParameters returns the ProtocolParameters field if non-nil, zero value otherwise.

### GetProtocolParametersOk

`func (o *RoutingPolicyResourcesDefStatus) GetProtocolParametersOk() (*ProtocolParametersStatus, bool)`

GetProtocolParametersOk returns a tuple with the ProtocolParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolParameters

`func (o *RoutingPolicyResourcesDefStatus) SetProtocolParameters(v ProtocolParametersStatus)`

SetProtocolParameters sets ProtocolParameters field to given value.

### HasProtocolParameters

`func (o *RoutingPolicyResourcesDefStatus) HasProtocolParameters() bool`

HasProtocolParameters returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *RoutingPolicyResourcesDefStatus) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *RoutingPolicyResourcesDefStatus) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *RoutingPolicyResourcesDefStatus) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *RoutingPolicyResourcesDefStatus) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetRerouteFallbackCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) GetRerouteFallbackCountersReverseDirection() RoutingPolicyCounters`

GetRerouteFallbackCountersReverseDirection returns the RerouteFallbackCountersReverseDirection field if non-nil, zero value otherwise.

### GetRerouteFallbackCountersReverseDirectionOk

`func (o *RoutingPolicyResourcesDefStatus) GetRerouteFallbackCountersReverseDirectionOk() (*RoutingPolicyCounters, bool)`

GetRerouteFallbackCountersReverseDirectionOk returns a tuple with the RerouteFallbackCountersReverseDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerouteFallbackCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) SetRerouteFallbackCountersReverseDirection(v RoutingPolicyCounters)`

SetRerouteFallbackCountersReverseDirection sets RerouteFallbackCountersReverseDirection field to given value.

### HasRerouteFallbackCountersReverseDirection

`func (o *RoutingPolicyResourcesDefStatus) HasRerouteFallbackCountersReverseDirection() bool`

HasRerouteFallbackCountersReverseDirection returns a boolean if a field has been set.

### GetAction

`func (o *RoutingPolicyResourcesDefStatus) GetAction() RoutingPolicyActionStatus`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyResourcesDefStatus) GetActionOk() (*RoutingPolicyActionStatus, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyResourcesDefStatus) SetAction(v RoutingPolicyActionStatus)`

SetAction sets Action field to given value.

### HasAction

`func (o *RoutingPolicyResourcesDefStatus) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetProtocolType

`func (o *RoutingPolicyResourcesDefStatus) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *RoutingPolicyResourcesDefStatus) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *RoutingPolicyResourcesDefStatus) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *RoutingPolicyResourcesDefStatus) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetPriority

`func (o *RoutingPolicyResourcesDefStatus) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RoutingPolicyResourcesDefStatus) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RoutingPolicyResourcesDefStatus) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RoutingPolicyResourcesDefStatus) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


