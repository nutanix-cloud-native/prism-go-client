# Edge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationVertexId** | Pointer to **string** | ID of destination vertex | [optional] 
**SourceVertexId** | Pointer to **string** | ID of source vertex | [optional] 
**TrafficList** | Pointer to [**[]Traffic**](Traffic.md) | Traffic information of the edge. | [optional] 

## Methods

### NewEdge

`func NewEdge() *Edge`

NewEdge instantiates a new Edge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeWithDefaults

`func NewEdgeWithDefaults() *Edge`

NewEdgeWithDefaults instantiates a new Edge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationVertexId

`func (o *Edge) GetDestinationVertexId() string`

GetDestinationVertexId returns the DestinationVertexId field if non-nil, zero value otherwise.

### GetDestinationVertexIdOk

`func (o *Edge) GetDestinationVertexIdOk() (*string, bool)`

GetDestinationVertexIdOk returns a tuple with the DestinationVertexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVertexId

`func (o *Edge) SetDestinationVertexId(v string)`

SetDestinationVertexId sets DestinationVertexId field to given value.

### HasDestinationVertexId

`func (o *Edge) HasDestinationVertexId() bool`

HasDestinationVertexId returns a boolean if a field has been set.

### GetSourceVertexId

`func (o *Edge) GetSourceVertexId() string`

GetSourceVertexId returns the SourceVertexId field if non-nil, zero value otherwise.

### GetSourceVertexIdOk

`func (o *Edge) GetSourceVertexIdOk() (*string, bool)`

GetSourceVertexIdOk returns a tuple with the SourceVertexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVertexId

`func (o *Edge) SetSourceVertexId(v string)`

SetSourceVertexId sets SourceVertexId field to given value.

### HasSourceVertexId

`func (o *Edge) HasSourceVertexId() bool`

HasSourceVertexId returns a boolean if a field has been set.

### GetTrafficList

`func (o *Edge) GetTrafficList() []Traffic`

GetTrafficList returns the TrafficList field if non-nil, zero value otherwise.

### GetTrafficListOk

`func (o *Edge) GetTrafficListOk() (*[]Traffic, bool)`

GetTrafficListOk returns a tuple with the TrafficList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficList

`func (o *Edge) SetTrafficList(v []Traffic)`

SetTrafficList sets TrafficList field to given value.

### HasTrafficList

`func (o *Edge) HasTrafficList() bool`

HasTrafficList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


