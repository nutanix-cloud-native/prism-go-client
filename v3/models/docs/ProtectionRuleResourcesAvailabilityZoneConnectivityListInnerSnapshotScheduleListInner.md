# ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryPointObjectiveSecs** | **int32** | A recovery point objective (RPO) is the maximum acceptable amount of data loss. RPO is measured in time (in seconds) and then dictates disaster recovery procedures. For example, if the RPO is set to 30 minutes, then a backup of the entity is required to be done every 30 minutes.  | 
**LocalSnapshotRetentionPolicy** | Pointer to [**SnapshotRetentionPolicy**](SnapshotRetentionPolicy.md) |  | [optional] 
**AutoSuspendTimeoutSecs** | Pointer to **int32** | Auto suspend timeout in case of connection failure between the sites. If not set, then policy will not be suspended in case of site connection failure.  | [optional] 
**SnapshotType** | Pointer to **string** | Crash consistent or Application Consistent snapshot  | [optional] 
**RemoteSnapshotRetentionPolicy** | Pointer to [**SnapshotRetentionPolicy**](SnapshotRetentionPolicy.md) |  | [optional] 

## Methods

### NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner

`func NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner(recoveryPointObjectiveSecs int32, ) *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner`

NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner instantiates a new ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInnerWithDefaults

`func NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInnerWithDefaults() *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner`

NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInnerWithDefaults instantiates a new ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryPointObjectiveSecs

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetRecoveryPointObjectiveSecs() int32`

GetRecoveryPointObjectiveSecs returns the RecoveryPointObjectiveSecs field if non-nil, zero value otherwise.

### GetRecoveryPointObjectiveSecsOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetRecoveryPointObjectiveSecsOk() (*int32, bool)`

GetRecoveryPointObjectiveSecsOk returns a tuple with the RecoveryPointObjectiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointObjectiveSecs

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) SetRecoveryPointObjectiveSecs(v int32)`

SetRecoveryPointObjectiveSecs sets RecoveryPointObjectiveSecs field to given value.


### GetLocalSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetLocalSnapshotRetentionPolicy() SnapshotRetentionPolicy`

GetLocalSnapshotRetentionPolicy returns the LocalSnapshotRetentionPolicy field if non-nil, zero value otherwise.

### GetLocalSnapshotRetentionPolicyOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetLocalSnapshotRetentionPolicyOk() (*SnapshotRetentionPolicy, bool)`

GetLocalSnapshotRetentionPolicyOk returns a tuple with the LocalSnapshotRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) SetLocalSnapshotRetentionPolicy(v SnapshotRetentionPolicy)`

SetLocalSnapshotRetentionPolicy sets LocalSnapshotRetentionPolicy field to given value.

### HasLocalSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) HasLocalSnapshotRetentionPolicy() bool`

HasLocalSnapshotRetentionPolicy returns a boolean if a field has been set.

### GetAutoSuspendTimeoutSecs

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetAutoSuspendTimeoutSecs() int32`

GetAutoSuspendTimeoutSecs returns the AutoSuspendTimeoutSecs field if non-nil, zero value otherwise.

### GetAutoSuspendTimeoutSecsOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetAutoSuspendTimeoutSecsOk() (*int32, bool)`

GetAutoSuspendTimeoutSecsOk returns a tuple with the AutoSuspendTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendTimeoutSecs

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) SetAutoSuspendTimeoutSecs(v int32)`

SetAutoSuspendTimeoutSecs sets AutoSuspendTimeoutSecs field to given value.

### HasAutoSuspendTimeoutSecs

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) HasAutoSuspendTimeoutSecs() bool`

HasAutoSuspendTimeoutSecs returns a boolean if a field has been set.

### GetSnapshotType

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetRemoteSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetRemoteSnapshotRetentionPolicy() SnapshotRetentionPolicy`

GetRemoteSnapshotRetentionPolicy returns the RemoteSnapshotRetentionPolicy field if non-nil, zero value otherwise.

### GetRemoteSnapshotRetentionPolicyOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) GetRemoteSnapshotRetentionPolicyOk() (*SnapshotRetentionPolicy, bool)`

GetRemoteSnapshotRetentionPolicyOk returns a tuple with the RemoteSnapshotRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) SetRemoteSnapshotRetentionPolicy(v SnapshotRetentionPolicy)`

SetRemoteSnapshotRetentionPolicy sets RemoteSnapshotRetentionPolicy field to given value.

### HasRemoteSnapshotRetentionPolicy

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner) HasRemoteSnapshotRetentionPolicy() bool`

HasRemoteSnapshotRetentionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


