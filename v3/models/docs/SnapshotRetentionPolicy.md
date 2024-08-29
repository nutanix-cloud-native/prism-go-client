# SnapshotRetentionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSnapshots** | Pointer to **int32** | Number of snapshots need to be retained. This will be set in case of linear snapshot retention.  | [optional] 
**RollupRetentionPolicy** | Pointer to [**SnapshotRetentionPolicyForRollupSchedules**](SnapshotRetentionPolicyForRollupSchedules.md) |  | [optional] 

## Methods

### NewSnapshotRetentionPolicy

`func NewSnapshotRetentionPolicy() *SnapshotRetentionPolicy`

NewSnapshotRetentionPolicy instantiates a new SnapshotRetentionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetentionPolicyWithDefaults

`func NewSnapshotRetentionPolicyWithDefaults() *SnapshotRetentionPolicy`

NewSnapshotRetentionPolicyWithDefaults instantiates a new SnapshotRetentionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumSnapshots

`func (o *SnapshotRetentionPolicy) GetNumSnapshots() int32`

GetNumSnapshots returns the NumSnapshots field if non-nil, zero value otherwise.

### GetNumSnapshotsOk

`func (o *SnapshotRetentionPolicy) GetNumSnapshotsOk() (*int32, bool)`

GetNumSnapshotsOk returns a tuple with the NumSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapshots

`func (o *SnapshotRetentionPolicy) SetNumSnapshots(v int32)`

SetNumSnapshots sets NumSnapshots field to given value.

### HasNumSnapshots

`func (o *SnapshotRetentionPolicy) HasNumSnapshots() bool`

HasNumSnapshots returns a boolean if a field has been set.

### GetRollupRetentionPolicy

`func (o *SnapshotRetentionPolicy) GetRollupRetentionPolicy() SnapshotRetentionPolicyForRollupSchedules`

GetRollupRetentionPolicy returns the RollupRetentionPolicy field if non-nil, zero value otherwise.

### GetRollupRetentionPolicyOk

`func (o *SnapshotRetentionPolicy) GetRollupRetentionPolicyOk() (*SnapshotRetentionPolicyForRollupSchedules, bool)`

GetRollupRetentionPolicyOk returns a tuple with the RollupRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupRetentionPolicy

`func (o *SnapshotRetentionPolicy) SetRollupRetentionPolicy(v SnapshotRetentionPolicyForRollupSchedules)`

SetRollupRetentionPolicy sets RollupRetentionPolicy field to given value.

### HasRollupRetentionPolicy

`func (o *SnapshotRetentionPolicy) HasRollupRetentionPolicy() bool`

HasRollupRetentionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


