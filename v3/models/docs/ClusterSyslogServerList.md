# ClusterSyslogServerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeList** | Pointer to [**[]ClusterSyslogServer**](ClusterSyslogServer.md) |  | [optional] 

## Methods

### NewClusterSyslogServerList

`func NewClusterSyslogServerList() *ClusterSyslogServerList`

NewClusterSyslogServerList instantiates a new ClusterSyslogServerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSyslogServerListWithDefaults

`func NewClusterSyslogServerListWithDefaults() *ClusterSyslogServerList`

NewClusterSyslogServerListWithDefaults instantiates a new ClusterSyslogServerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeList

`func (o *ClusterSyslogServerList) GetPeList() []ClusterSyslogServer`

GetPeList returns the PeList field if non-nil, zero value otherwise.

### GetPeListOk

`func (o *ClusterSyslogServerList) GetPeListOk() (*[]ClusterSyslogServer, bool)`

GetPeListOk returns a tuple with the PeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeList

`func (o *ClusterSyslogServerList) SetPeList(v []ClusterSyslogServer)`

SetPeList sets PeList field to given value.

### HasPeList

`func (o *ClusterSyslogServerList) HasPeList() bool`

HasPeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


