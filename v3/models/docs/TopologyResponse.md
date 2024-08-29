# TopologyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VertexList** | Pointer to [**[]Vertex**](Vertex.md) | List containing vertices information in the topology visualization in Security Monitoring.  | [optional] 
**EdgeList** | Pointer to [**[]Edge**](Edge.md) | List containing edges information in the topology visualization in Security Monitoring.  | [optional] 

## Methods

### NewTopologyResponse

`func NewTopologyResponse() *TopologyResponse`

NewTopologyResponse instantiates a new TopologyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyResponseWithDefaults

`func NewTopologyResponseWithDefaults() *TopologyResponse`

NewTopologyResponseWithDefaults instantiates a new TopologyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVertexList

`func (o *TopologyResponse) GetVertexList() []Vertex`

GetVertexList returns the VertexList field if non-nil, zero value otherwise.

### GetVertexListOk

`func (o *TopologyResponse) GetVertexListOk() (*[]Vertex, bool)`

GetVertexListOk returns a tuple with the VertexList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertexList

`func (o *TopologyResponse) SetVertexList(v []Vertex)`

SetVertexList sets VertexList field to given value.

### HasVertexList

`func (o *TopologyResponse) HasVertexList() bool`

HasVertexList returns a boolean if a field has been set.

### GetEdgeList

`func (o *TopologyResponse) GetEdgeList() []Edge`

GetEdgeList returns the EdgeList field if non-nil, zero value otherwise.

### GetEdgeListOk

`func (o *TopologyResponse) GetEdgeListOk() (*[]Edge, bool)`

GetEdgeListOk returns a tuple with the EdgeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeList

`func (o *TopologyResponse) SetEdgeList(v []Edge)`

SetEdgeList sets EdgeList field to given value.

### HasEdgeList

`func (o *TopologyResponse) HasEdgeList() bool`

HasEdgeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


