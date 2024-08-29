# ClusterSynchronousReplicationCapableResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RttMsecs** | Pointer to **string** | Round trip time in milliseconds. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewClusterSynchronousReplicationCapableResponseInner

`func NewClusterSynchronousReplicationCapableResponseInner() *ClusterSynchronousReplicationCapableResponseInner`

NewClusterSynchronousReplicationCapableResponseInner instantiates a new ClusterSynchronousReplicationCapableResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSynchronousReplicationCapableResponseInnerWithDefaults

`func NewClusterSynchronousReplicationCapableResponseInnerWithDefaults() *ClusterSynchronousReplicationCapableResponseInner`

NewClusterSynchronousReplicationCapableResponseInnerWithDefaults instantiates a new ClusterSynchronousReplicationCapableResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRttMsecs

`func (o *ClusterSynchronousReplicationCapableResponseInner) GetRttMsecs() string`

GetRttMsecs returns the RttMsecs field if non-nil, zero value otherwise.

### GetRttMsecsOk

`func (o *ClusterSynchronousReplicationCapableResponseInner) GetRttMsecsOk() (*string, bool)`

GetRttMsecsOk returns a tuple with the RttMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttMsecs

`func (o *ClusterSynchronousReplicationCapableResponseInner) SetRttMsecs(v string)`

SetRttMsecs sets RttMsecs field to given value.

### HasRttMsecs

`func (o *ClusterSynchronousReplicationCapableResponseInner) HasRttMsecs() bool`

HasRttMsecs returns a boolean if a field has been set.

### GetClusterReference

`func (o *ClusterSynchronousReplicationCapableResponseInner) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *ClusterSynchronousReplicationCapableResponseInner) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *ClusterSynchronousReplicationCapableResponseInner) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *ClusterSynchronousReplicationCapableResponseInner) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


