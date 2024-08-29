# SnapshotRetentionPolicyForRollupSchedules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Multiple** | Pointer to **int32** | Multiplier to &#39;snapshot_interval_type&#39;. For example if &#39;snapshot_interval_type&#39; is \&quot;YEARLY\&quot; and &#39;multiple&#39; is 5, then 5 years worth of rollup snapshots will be retained.  | [optional] 
**SnapshotIntervalType** | Pointer to **string** | Snapshot interval period.  | [optional] 

## Methods

### NewSnapshotRetentionPolicyForRollupSchedules

`func NewSnapshotRetentionPolicyForRollupSchedules() *SnapshotRetentionPolicyForRollupSchedules`

NewSnapshotRetentionPolicyForRollupSchedules instantiates a new SnapshotRetentionPolicyForRollupSchedules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetentionPolicyForRollupSchedulesWithDefaults

`func NewSnapshotRetentionPolicyForRollupSchedulesWithDefaults() *SnapshotRetentionPolicyForRollupSchedules`

NewSnapshotRetentionPolicyForRollupSchedulesWithDefaults instantiates a new SnapshotRetentionPolicyForRollupSchedules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiple

`func (o *SnapshotRetentionPolicyForRollupSchedules) GetMultiple() int32`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *SnapshotRetentionPolicyForRollupSchedules) GetMultipleOk() (*int32, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *SnapshotRetentionPolicyForRollupSchedules) SetMultiple(v int32)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *SnapshotRetentionPolicyForRollupSchedules) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetSnapshotIntervalType

`func (o *SnapshotRetentionPolicyForRollupSchedules) GetSnapshotIntervalType() string`

GetSnapshotIntervalType returns the SnapshotIntervalType field if non-nil, zero value otherwise.

### GetSnapshotIntervalTypeOk

`func (o *SnapshotRetentionPolicyForRollupSchedules) GetSnapshotIntervalTypeOk() (*string, bool)`

GetSnapshotIntervalTypeOk returns a tuple with the SnapshotIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIntervalType

`func (o *SnapshotRetentionPolicyForRollupSchedules) SetSnapshotIntervalType(v string)`

SetSnapshotIntervalType sets SnapshotIntervalType field to given value.

### HasSnapshotIntervalType

`func (o *SnapshotRetentionPolicyForRollupSchedules) HasSnapshotIntervalType() bool`

HasSnapshotIntervalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


