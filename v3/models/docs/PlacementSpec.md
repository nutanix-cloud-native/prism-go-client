# PlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitySpecList** | Pointer to [**[]EntitySpec**](EntitySpec.md) | List of entity specs. | [optional] 
**HypervisorType** | Pointer to **string** | Preferred hypervisor type for the entities.  | [optional] 
**ExcludeClusterUuidList** | Pointer to **[]string** | List of UUIDs of the clusters which have to be excluded while auto selecting a target cluster for replication.  | [optional] 
**PreferredClusterUuidList** | Pointer to **[]string** | List of uuids of the cluster which are preferred to be the replication target.  | [optional] 
**PreferredClusterUuid** | Pointer to **string** | Uuid of the cluster which is preferred to be the replication target.  | [optional] 
**ReplicationType** | Pointer to **string** | Type of the replication. It can be asynchronous or synchronous.  | [optional] 
**OperationType** | Pointer to **string** | Operation to be performed for entities. | [optional] 
**SourceAvailabilityZoneUrl** | Pointer to **string** | URL of the source availability zone. Based on this, the mapped load balancer ip address and port of the replication target will be returned.  | [optional] 
**TenantUuid** | Pointer to **string** | UUID of the tenant. | [optional] 

## Methods

### NewPlacementSpec

`func NewPlacementSpec() *PlacementSpec`

NewPlacementSpec instantiates a new PlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementSpecWithDefaults

`func NewPlacementSpecWithDefaults() *PlacementSpec`

NewPlacementSpecWithDefaults instantiates a new PlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitySpecList

`func (o *PlacementSpec) GetEntitySpecList() []EntitySpec`

GetEntitySpecList returns the EntitySpecList field if non-nil, zero value otherwise.

### GetEntitySpecListOk

`func (o *PlacementSpec) GetEntitySpecListOk() (*[]EntitySpec, bool)`

GetEntitySpecListOk returns a tuple with the EntitySpecList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySpecList

`func (o *PlacementSpec) SetEntitySpecList(v []EntitySpec)`

SetEntitySpecList sets EntitySpecList field to given value.

### HasEntitySpecList

`func (o *PlacementSpec) HasEntitySpecList() bool`

HasEntitySpecList returns a boolean if a field has been set.

### GetHypervisorType

`func (o *PlacementSpec) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *PlacementSpec) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *PlacementSpec) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *PlacementSpec) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetExcludeClusterUuidList

`func (o *PlacementSpec) GetExcludeClusterUuidList() []string`

GetExcludeClusterUuidList returns the ExcludeClusterUuidList field if non-nil, zero value otherwise.

### GetExcludeClusterUuidListOk

`func (o *PlacementSpec) GetExcludeClusterUuidListOk() (*[]string, bool)`

GetExcludeClusterUuidListOk returns a tuple with the ExcludeClusterUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeClusterUuidList

`func (o *PlacementSpec) SetExcludeClusterUuidList(v []string)`

SetExcludeClusterUuidList sets ExcludeClusterUuidList field to given value.

### HasExcludeClusterUuidList

`func (o *PlacementSpec) HasExcludeClusterUuidList() bool`

HasExcludeClusterUuidList returns a boolean if a field has been set.

### GetPreferredClusterUuidList

`func (o *PlacementSpec) GetPreferredClusterUuidList() []string`

GetPreferredClusterUuidList returns the PreferredClusterUuidList field if non-nil, zero value otherwise.

### GetPreferredClusterUuidListOk

`func (o *PlacementSpec) GetPreferredClusterUuidListOk() (*[]string, bool)`

GetPreferredClusterUuidListOk returns a tuple with the PreferredClusterUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredClusterUuidList

`func (o *PlacementSpec) SetPreferredClusterUuidList(v []string)`

SetPreferredClusterUuidList sets PreferredClusterUuidList field to given value.

### HasPreferredClusterUuidList

`func (o *PlacementSpec) HasPreferredClusterUuidList() bool`

HasPreferredClusterUuidList returns a boolean if a field has been set.

### GetPreferredClusterUuid

`func (o *PlacementSpec) GetPreferredClusterUuid() string`

GetPreferredClusterUuid returns the PreferredClusterUuid field if non-nil, zero value otherwise.

### GetPreferredClusterUuidOk

`func (o *PlacementSpec) GetPreferredClusterUuidOk() (*string, bool)`

GetPreferredClusterUuidOk returns a tuple with the PreferredClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredClusterUuid

`func (o *PlacementSpec) SetPreferredClusterUuid(v string)`

SetPreferredClusterUuid sets PreferredClusterUuid field to given value.

### HasPreferredClusterUuid

`func (o *PlacementSpec) HasPreferredClusterUuid() bool`

HasPreferredClusterUuid returns a boolean if a field has been set.

### GetReplicationType

`func (o *PlacementSpec) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *PlacementSpec) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *PlacementSpec) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *PlacementSpec) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetOperationType

`func (o *PlacementSpec) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PlacementSpec) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PlacementSpec) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *PlacementSpec) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetSourceAvailabilityZoneUrl

`func (o *PlacementSpec) GetSourceAvailabilityZoneUrl() string`

GetSourceAvailabilityZoneUrl returns the SourceAvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneUrlOk

`func (o *PlacementSpec) GetSourceAvailabilityZoneUrlOk() (*string, bool)`

GetSourceAvailabilityZoneUrlOk returns a tuple with the SourceAvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneUrl

`func (o *PlacementSpec) SetSourceAvailabilityZoneUrl(v string)`

SetSourceAvailabilityZoneUrl sets SourceAvailabilityZoneUrl field to given value.

### HasSourceAvailabilityZoneUrl

`func (o *PlacementSpec) HasSourceAvailabilityZoneUrl() bool`

HasSourceAvailabilityZoneUrl returns a boolean if a field has been set.

### GetTenantUuid

`func (o *PlacementSpec) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *PlacementSpec) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *PlacementSpec) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *PlacementSpec) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


