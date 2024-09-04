# ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStoreConfig** | Pointer to [**DataStoreConfig**](DataStoreConfig.md) |  | [optional] 
**NodeSpecList** | Pointer to [**[]NodeSpec**](NodeSpec.md) |  | [optional] 
**EffectiveCapacity** | Pointer to [**GenericResourceSpec**](GenericResourceSpec.md) |  | [optional] 
**ResourceList** | Pointer to [**[]ResourceTimeSeries**](ResourceTimeSeries.md) |  | [optional] 

## Methods

### NewClusterSpec

`func NewClusterSpec() *ClusterSpec`

NewClusterSpec instantiates a new ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSpecWithDefaults

`func NewClusterSpecWithDefaults() *ClusterSpec`

NewClusterSpecWithDefaults instantiates a new ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStoreConfig

`func (o *ClusterSpec) GetDataStoreConfig() DataStoreConfig`

GetDataStoreConfig returns the DataStoreConfig field if non-nil, zero value otherwise.

### GetDataStoreConfigOk

`func (o *ClusterSpec) GetDataStoreConfigOk() (*DataStoreConfig, bool)`

GetDataStoreConfigOk returns a tuple with the DataStoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreConfig

`func (o *ClusterSpec) SetDataStoreConfig(v DataStoreConfig)`

SetDataStoreConfig sets DataStoreConfig field to given value.

### HasDataStoreConfig

`func (o *ClusterSpec) HasDataStoreConfig() bool`

HasDataStoreConfig returns a boolean if a field has been set.

### GetNodeSpecList

`func (o *ClusterSpec) GetNodeSpecList() []NodeSpec`

GetNodeSpecList returns the NodeSpecList field if non-nil, zero value otherwise.

### GetNodeSpecListOk

`func (o *ClusterSpec) GetNodeSpecListOk() (*[]NodeSpec, bool)`

GetNodeSpecListOk returns a tuple with the NodeSpecList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSpecList

`func (o *ClusterSpec) SetNodeSpecList(v []NodeSpec)`

SetNodeSpecList sets NodeSpecList field to given value.

### HasNodeSpecList

`func (o *ClusterSpec) HasNodeSpecList() bool`

HasNodeSpecList returns a boolean if a field has been set.

### GetEffectiveCapacity

`func (o *ClusterSpec) GetEffectiveCapacity() GenericResourceSpec`

GetEffectiveCapacity returns the EffectiveCapacity field if non-nil, zero value otherwise.

### GetEffectiveCapacityOk

`func (o *ClusterSpec) GetEffectiveCapacityOk() (*GenericResourceSpec, bool)`

GetEffectiveCapacityOk returns a tuple with the EffectiveCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCapacity

`func (o *ClusterSpec) SetEffectiveCapacity(v GenericResourceSpec)`

SetEffectiveCapacity sets EffectiveCapacity field to given value.

### HasEffectiveCapacity

`func (o *ClusterSpec) HasEffectiveCapacity() bool`

HasEffectiveCapacity returns a boolean if a field has been set.

### GetResourceList

`func (o *ClusterSpec) GetResourceList() []ResourceTimeSeries`

GetResourceList returns the ResourceList field if non-nil, zero value otherwise.

### GetResourceListOk

`func (o *ClusterSpec) GetResourceListOk() (*[]ResourceTimeSeries, bool)`

GetResourceListOk returns a tuple with the ResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceList

`func (o *ClusterSpec) SetResourceList(v []ResourceTimeSeries)`

SetResourceList sets ResourceList field to given value.

### HasResourceList

`func (o *ClusterSpec) HasResourceList() bool`

HasResourceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


