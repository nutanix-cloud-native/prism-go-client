# ClusterSynchronousReplicationCapableInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) |  | [optional] 
**RemoteAvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**RemoteClusterReference** | [**ClusterReference**](ClusterReference.md) |  | 

## Methods

### NewClusterSynchronousReplicationCapableInput

`func NewClusterSynchronousReplicationCapableInput(remoteAvailabilityZoneReference AvailabilityZoneReference, remoteClusterReference ClusterReference, ) *ClusterSynchronousReplicationCapableInput`

NewClusterSynchronousReplicationCapableInput instantiates a new ClusterSynchronousReplicationCapableInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSynchronousReplicationCapableInputWithDefaults

`func NewClusterSynchronousReplicationCapableInputWithDefaults() *ClusterSynchronousReplicationCapableInput`

NewClusterSynchronousReplicationCapableInputWithDefaults instantiates a new ClusterSynchronousReplicationCapableInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceClusterReferenceList

`func (o *ClusterSynchronousReplicationCapableInput) GetSourceClusterReferenceList() []ClusterReference`

GetSourceClusterReferenceList returns the SourceClusterReferenceList field if non-nil, zero value otherwise.

### GetSourceClusterReferenceListOk

`func (o *ClusterSynchronousReplicationCapableInput) GetSourceClusterReferenceListOk() (*[]ClusterReference, bool)`

GetSourceClusterReferenceListOk returns a tuple with the SourceClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterReferenceList

`func (o *ClusterSynchronousReplicationCapableInput) SetSourceClusterReferenceList(v []ClusterReference)`

SetSourceClusterReferenceList sets SourceClusterReferenceList field to given value.

### HasSourceClusterReferenceList

`func (o *ClusterSynchronousReplicationCapableInput) HasSourceClusterReferenceList() bool`

HasSourceClusterReferenceList returns a boolean if a field has been set.

### GetRemoteAvailabilityZoneReference

`func (o *ClusterSynchronousReplicationCapableInput) GetRemoteAvailabilityZoneReference() AvailabilityZoneReference`

GetRemoteAvailabilityZoneReference returns the RemoteAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetRemoteAvailabilityZoneReferenceOk

`func (o *ClusterSynchronousReplicationCapableInput) GetRemoteAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetRemoteAvailabilityZoneReferenceOk returns a tuple with the RemoteAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAvailabilityZoneReference

`func (o *ClusterSynchronousReplicationCapableInput) SetRemoteAvailabilityZoneReference(v AvailabilityZoneReference)`

SetRemoteAvailabilityZoneReference sets RemoteAvailabilityZoneReference field to given value.


### GetRemoteClusterReference

`func (o *ClusterSynchronousReplicationCapableInput) GetRemoteClusterReference() ClusterReference`

GetRemoteClusterReference returns the RemoteClusterReference field if non-nil, zero value otherwise.

### GetRemoteClusterReferenceOk

`func (o *ClusterSynchronousReplicationCapableInput) GetRemoteClusterReferenceOk() (*ClusterReference, bool)`

GetRemoteClusterReferenceOk returns a tuple with the RemoteClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterReference

`func (o *ClusterSynchronousReplicationCapableInput) SetRemoteClusterReference(v ClusterReference)`

SetRemoteClusterReference sets RemoteClusterReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


