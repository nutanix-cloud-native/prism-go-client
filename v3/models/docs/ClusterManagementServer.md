# ClusterManagementServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**DrsEnabled** | Pointer to **bool** | Denotes if DRS is enabled or not. | [optional] 
**StatusList** | Pointer to **[]string** | Array of management server status: - &#39;REGISTERED&#39;: Indicates whether the server is registered with                 Nutanix or not. - &#39;IN_USE&#39;: Indicates whether any host is managed by this server or             not.  | [optional] 
**Type** | **string** |  | 

## Methods

### NewClusterManagementServer

`func NewClusterManagementServer(ip string, type_ string, ) *ClusterManagementServer`

NewClusterManagementServer instantiates a new ClusterManagementServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterManagementServerWithDefaults

`func NewClusterManagementServerWithDefaults() *ClusterManagementServer`

NewClusterManagementServerWithDefaults instantiates a new ClusterManagementServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ClusterManagementServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterManagementServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterManagementServer) SetIp(v string)`

SetIp sets Ip field to given value.


### GetDrsEnabled

`func (o *ClusterManagementServer) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *ClusterManagementServer) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *ClusterManagementServer) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.

### HasDrsEnabled

`func (o *ClusterManagementServer) HasDrsEnabled() bool`

HasDrsEnabled returns a boolean if a field has been set.

### GetStatusList

`func (o *ClusterManagementServer) GetStatusList() []string`

GetStatusList returns the StatusList field if non-nil, zero value otherwise.

### GetStatusListOk

`func (o *ClusterManagementServer) GetStatusListOk() (*[]string, bool)`

GetStatusListOk returns a tuple with the StatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusList

`func (o *ClusterManagementServer) SetStatusList(v []string)`

SetStatusList sets StatusList field to given value.

### HasStatusList

`func (o *ClusterManagementServer) HasStatusList() bool`

HasStatusList returns a boolean if a field has been set.

### GetType

`func (o *ClusterManagementServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterManagementServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterManagementServer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


