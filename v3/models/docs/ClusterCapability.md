# ClusterCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrecheckList** | Pointer to [**[]ClusterCapabilityPrecheck**](ClusterCapabilityPrecheck.md) | list of any additional precheck for cluster capability | [optional] 
**IsCapable** | Pointer to **bool** | If this cluster is capable for a feature. | [optional] 
**ClusterName** | Pointer to **string** | Name of the cluster. | [optional] 
**HypervisorList** | Pointer to [**[]HypervisorInfo**](HypervisorInfo.md) | Type and version of hypervisors on the cluster. | [optional] 
**ClusterUuid** | Pointer to **string** | Cluster UUID of a cluster managed by Prism Central.  | [optional] 
**AosVersion** | Pointer to **string** | AOS version of the cluster. | [optional] 

## Methods

### NewClusterCapability

`func NewClusterCapability() *ClusterCapability`

NewClusterCapability instantiates a new ClusterCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCapabilityWithDefaults

`func NewClusterCapabilityWithDefaults() *ClusterCapability`

NewClusterCapabilityWithDefaults instantiates a new ClusterCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecheckList

`func (o *ClusterCapability) GetPrecheckList() []ClusterCapabilityPrecheck`

GetPrecheckList returns the PrecheckList field if non-nil, zero value otherwise.

### GetPrecheckListOk

`func (o *ClusterCapability) GetPrecheckListOk() (*[]ClusterCapabilityPrecheck, bool)`

GetPrecheckListOk returns a tuple with the PrecheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckList

`func (o *ClusterCapability) SetPrecheckList(v []ClusterCapabilityPrecheck)`

SetPrecheckList sets PrecheckList field to given value.

### HasPrecheckList

`func (o *ClusterCapability) HasPrecheckList() bool`

HasPrecheckList returns a boolean if a field has been set.

### GetIsCapable

`func (o *ClusterCapability) GetIsCapable() bool`

GetIsCapable returns the IsCapable field if non-nil, zero value otherwise.

### GetIsCapableOk

`func (o *ClusterCapability) GetIsCapableOk() (*bool, bool)`

GetIsCapableOk returns a tuple with the IsCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCapable

`func (o *ClusterCapability) SetIsCapable(v bool)`

SetIsCapable sets IsCapable field to given value.

### HasIsCapable

`func (o *ClusterCapability) HasIsCapable() bool`

HasIsCapable returns a boolean if a field has been set.

### GetClusterName

`func (o *ClusterCapability) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterCapability) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterCapability) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *ClusterCapability) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetHypervisorList

`func (o *ClusterCapability) GetHypervisorList() []HypervisorInfo`

GetHypervisorList returns the HypervisorList field if non-nil, zero value otherwise.

### GetHypervisorListOk

`func (o *ClusterCapability) GetHypervisorListOk() (*[]HypervisorInfo, bool)`

GetHypervisorListOk returns a tuple with the HypervisorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorList

`func (o *ClusterCapability) SetHypervisorList(v []HypervisorInfo)`

SetHypervisorList sets HypervisorList field to given value.

### HasHypervisorList

`func (o *ClusterCapability) HasHypervisorList() bool`

HasHypervisorList returns a boolean if a field has been set.

### GetClusterUuid

`func (o *ClusterCapability) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *ClusterCapability) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *ClusterCapability) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *ClusterCapability) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetAosVersion

`func (o *ClusterCapability) GetAosVersion() string`

GetAosVersion returns the AosVersion field if non-nil, zero value otherwise.

### GetAosVersionOk

`func (o *ClusterCapability) GetAosVersionOk() (*string, bool)`

GetAosVersionOk returns a tuple with the AosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAosVersion

`func (o *ClusterCapability) SetAosVersion(v string)`

SetAosVersion sets AosVersion field to given value.

### HasAosVersion

`func (o *ClusterCapability) HasAosVersion() bool`

HasAosVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


