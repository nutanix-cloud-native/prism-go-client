# ClusterSyslogServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RsyslogServerList** | Pointer to [**[]SyslogServerResources**](SyslogServerResources.md) |  | [optional] 
**ClusterVersion** | Pointer to **string** |  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewClusterSyslogServer

`func NewClusterSyslogServer() *ClusterSyslogServer`

NewClusterSyslogServer instantiates a new ClusterSyslogServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSyslogServerWithDefaults

`func NewClusterSyslogServerWithDefaults() *ClusterSyslogServer`

NewClusterSyslogServerWithDefaults instantiates a new ClusterSyslogServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsyslogServerList

`func (o *ClusterSyslogServer) GetRsyslogServerList() []SyslogServerResources`

GetRsyslogServerList returns the RsyslogServerList field if non-nil, zero value otherwise.

### GetRsyslogServerListOk

`func (o *ClusterSyslogServer) GetRsyslogServerListOk() (*[]SyslogServerResources, bool)`

GetRsyslogServerListOk returns a tuple with the RsyslogServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsyslogServerList

`func (o *ClusterSyslogServer) SetRsyslogServerList(v []SyslogServerResources)`

SetRsyslogServerList sets RsyslogServerList field to given value.

### HasRsyslogServerList

`func (o *ClusterSyslogServer) HasRsyslogServerList() bool`

HasRsyslogServerList returns a boolean if a field has been set.

### GetClusterVersion

`func (o *ClusterSyslogServer) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *ClusterSyslogServer) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *ClusterSyslogServer) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *ClusterSyslogServer) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.

### GetClusterReference

`func (o *ClusterSyslogServer) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *ClusterSyslogServer) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *ClusterSyslogServer) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *ClusterSyslogServer) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


