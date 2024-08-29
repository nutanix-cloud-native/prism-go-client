# PacketTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StageList** | Pointer to [**[]PacketTraceStage**](PacketTraceStage.md) | Details of stages that dropped/forwarded the packet | [optional] 

## Methods

### NewPacketTrace

`func NewPacketTrace() *PacketTrace`

NewPacketTrace instantiates a new PacketTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketTraceWithDefaults

`func NewPacketTraceWithDefaults() *PacketTrace`

NewPacketTraceWithDefaults instantiates a new PacketTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStageList

`func (o *PacketTrace) GetStageList() []PacketTraceStage`

GetStageList returns the StageList field if non-nil, zero value otherwise.

### GetStageListOk

`func (o *PacketTrace) GetStageListOk() (*[]PacketTraceStage, bool)`

GetStageListOk returns a tuple with the StageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageList

`func (o *PacketTrace) SetStageList(v []PacketTraceStage)`

SetStageList sets StageList field to given value.

### HasStageList

`func (o *PacketTrace) HasStageList() bool`

HasStageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


