# Vertex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOneHop** | Pointer to **bool** | Denotes if this is a one-hop node. | [optional] 
**EntityType** | Pointer to **string** | Type of the entity represented by the vertex | [optional] 
**Label** | Pointer to **string** | Vertex label | [optional] 
**VertexId** | Pointer to **string** | ID of vertex. | [optional] 
**EntityUuid** | Pointer to **string** | UUID of the entity represented by the vertex. | [optional] 
**InstanceCount** | Pointer to **int32** | Number of instances within the node denoting a group | [optional] 

## Methods

### NewVertex

`func NewVertex() *Vertex`

NewVertex instantiates a new Vertex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVertexWithDefaults

`func NewVertexWithDefaults() *Vertex`

NewVertexWithDefaults instantiates a new Vertex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOneHop

`func (o *Vertex) GetIsOneHop() bool`

GetIsOneHop returns the IsOneHop field if non-nil, zero value otherwise.

### GetIsOneHopOk

`func (o *Vertex) GetIsOneHopOk() (*bool, bool)`

GetIsOneHopOk returns a tuple with the IsOneHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOneHop

`func (o *Vertex) SetIsOneHop(v bool)`

SetIsOneHop sets IsOneHop field to given value.

### HasIsOneHop

`func (o *Vertex) HasIsOneHop() bool`

HasIsOneHop returns a boolean if a field has been set.

### GetEntityType

`func (o *Vertex) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Vertex) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Vertex) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Vertex) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetLabel

`func (o *Vertex) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Vertex) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Vertex) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Vertex) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetVertexId

`func (o *Vertex) GetVertexId() string`

GetVertexId returns the VertexId field if non-nil, zero value otherwise.

### GetVertexIdOk

`func (o *Vertex) GetVertexIdOk() (*string, bool)`

GetVertexIdOk returns a tuple with the VertexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexId

`func (o *Vertex) SetVertexId(v string)`

SetVertexId sets VertexId field to given value.

### HasVertexId

`func (o *Vertex) HasVertexId() bool`

HasVertexId returns a boolean if a field has been set.

### GetEntityUuid

`func (o *Vertex) GetEntityUuid() string`

GetEntityUuid returns the EntityUuid field if non-nil, zero value otherwise.

### GetEntityUuidOk

`func (o *Vertex) GetEntityUuidOk() (*string, bool)`

GetEntityUuidOk returns a tuple with the EntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuid

`func (o *Vertex) SetEntityUuid(v string)`

SetEntityUuid sets EntityUuid field to given value.

### HasEntityUuid

`func (o *Vertex) HasEntityUuid() bool`

HasEntityUuid returns a boolean if a field has been set.

### GetInstanceCount

`func (o *Vertex) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *Vertex) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *Vertex) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *Vertex) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


