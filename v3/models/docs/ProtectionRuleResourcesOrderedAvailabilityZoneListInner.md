# ProtectionRuleResourcesOrderedAvailabilityZoneListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | UUID of specific cluster to which we will be replicating. This need to set only if availability zone is of type PC. This can only be specified if target type is AOS_CLUSTER.  | [optional] 
**AvailabilityZoneUrl** | Pointer to **string** | The FQDN or IP address of the availability zone. Every Prism Central deployment acts as an availability zone.  | [optional] 
**ClusterUuidList** | Pointer to **[]string** | List of cluster UUIDs which are source or target for replication. This should be set only if the availability zone is of type PC.  | [optional] 
**TargetType** | Pointer to **string** | Target type for replication. | [optional] 

## Methods

### NewProtectionRuleResourcesOrderedAvailabilityZoneListInner

`func NewProtectionRuleResourcesOrderedAvailabilityZoneListInner() *ProtectionRuleResourcesOrderedAvailabilityZoneListInner`

NewProtectionRuleResourcesOrderedAvailabilityZoneListInner instantiates a new ProtectionRuleResourcesOrderedAvailabilityZoneListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleResourcesOrderedAvailabilityZoneListInnerWithDefaults

`func NewProtectionRuleResourcesOrderedAvailabilityZoneListInnerWithDefaults() *ProtectionRuleResourcesOrderedAvailabilityZoneListInner`

NewProtectionRuleResourcesOrderedAvailabilityZoneListInnerWithDefaults instantiates a new ProtectionRuleResourcesOrderedAvailabilityZoneListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetAvailabilityZoneUrl

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.

### HasAvailabilityZoneUrl

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) HasAvailabilityZoneUrl() bool`

HasAvailabilityZoneUrl returns a boolean if a field has been set.

### GetClusterUuidList

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetClusterUuidList() []string`

GetClusterUuidList returns the ClusterUuidList field if non-nil, zero value otherwise.

### GetClusterUuidListOk

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetClusterUuidListOk() (*[]string, bool)`

GetClusterUuidListOk returns a tuple with the ClusterUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuidList

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) SetClusterUuidList(v []string)`

SetClusterUuidList sets ClusterUuidList field to given value.

### HasClusterUuidList

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) HasClusterUuidList() bool`

HasClusterUuidList returns a boolean if a field has been set.

### GetTargetType

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ProtectionRuleResourcesOrderedAvailabilityZoneListInner) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


