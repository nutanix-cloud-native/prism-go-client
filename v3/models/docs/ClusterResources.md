# ClusterResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ClusterConfigSpec**](ClusterConfigSpec.md) |  | [optional] 
**Network** | Pointer to [**ClusterNetwork**](ClusterNetwork.md) |  | [optional] 
**RuntimeStatusList** | Pointer to **[]string** | Cluster onging operations. | [optional] 

## Methods

### NewClusterResources

`func NewClusterResources() *ClusterResources`

NewClusterResources instantiates a new ClusterResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResourcesWithDefaults

`func NewClusterResourcesWithDefaults() *ClusterResources`

NewClusterResourcesWithDefaults instantiates a new ClusterResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ClusterResources) GetConfig() ClusterConfigSpec`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterResources) GetConfigOk() (*ClusterConfigSpec, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterResources) SetConfig(v ClusterConfigSpec)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterResources) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetwork

`func (o *ClusterResources) GetNetwork() ClusterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterResources) GetNetworkOk() (*ClusterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterResources) SetNetwork(v ClusterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterResources) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuntimeStatusList

`func (o *ClusterResources) GetRuntimeStatusList() []string`

GetRuntimeStatusList returns the RuntimeStatusList field if non-nil, zero value otherwise.

### GetRuntimeStatusListOk

`func (o *ClusterResources) GetRuntimeStatusListOk() (*[]string, bool)`

GetRuntimeStatusListOk returns a tuple with the RuntimeStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStatusList

`func (o *ClusterResources) SetRuntimeStatusList(v []string)`

SetRuntimeStatusList sets RuntimeStatusList field to given value.

### HasRuntimeStatusList

`func (o *ClusterResources) HasRuntimeStatusList() bool`

HasRuntimeStatusList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


