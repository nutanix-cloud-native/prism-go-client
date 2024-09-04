# PacketTraceStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**RoutingPolicyReference** | Pointer to [**RoutingPolicyReference**](RoutingPolicyReference.md) |  | [optional] 
**DestinationIp** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** | Indicates if the packet was forwarded to a Xi VM, Internet or other data center  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 

## Methods

### NewPacketTraceStage

`func NewPacketTraceStage() *PacketTraceStage`

NewPacketTraceStage instantiates a new PacketTraceStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketTraceStageWithDefaults

`func NewPacketTraceStageWithDefaults() *PacketTraceStage`

NewPacketTraceStageWithDefaults instantiates a new PacketTraceStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PacketTraceStage) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PacketTraceStage) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PacketTraceStage) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PacketTraceStage) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRoutingPolicyReference

`func (o *PacketTraceStage) GetRoutingPolicyReference() RoutingPolicyReference`

GetRoutingPolicyReference returns the RoutingPolicyReference field if non-nil, zero value otherwise.

### GetRoutingPolicyReferenceOk

`func (o *PacketTraceStage) GetRoutingPolicyReferenceOk() (*RoutingPolicyReference, bool)`

GetRoutingPolicyReferenceOk returns a tuple with the RoutingPolicyReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicyReference

`func (o *PacketTraceStage) SetRoutingPolicyReference(v RoutingPolicyReference)`

SetRoutingPolicyReference sets RoutingPolicyReference field to given value.

### HasRoutingPolicyReference

`func (o *PacketTraceStage) HasRoutingPolicyReference() bool`

HasRoutingPolicyReference returns a boolean if a field has been set.

### GetDestinationIp

`func (o *PacketTraceStage) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *PacketTraceStage) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *PacketTraceStage) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *PacketTraceStage) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDestinationType

`func (o *PacketTraceStage) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PacketTraceStage) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PacketTraceStage) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *PacketTraceStage) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetStage

`func (o *PacketTraceStage) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PacketTraceStage) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PacketTraceStage) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PacketTraceStage) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


