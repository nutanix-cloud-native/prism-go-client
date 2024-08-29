# ProtectionRuleResourcesAvailabilityZoneConnectivityListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAvailabilityZoneIndex** | Pointer to **int32** | Index of the availability zone in ordered_availability_zone_list. This represents the availability zone where the entity needs to be replicated to. Index starts at 0.  | [optional] 
**SourceAvailabilityZoneIndex** | Pointer to **int32** | Index of the availability zone in ordered_availability_zone_list. This represents the source availability zone where the entity is running. Index starts at 0.  | [optional] 
**SnapshotScheduleList** | Pointer to [**[]ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner**](ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner.md) | Snapshot schedules for the pair of the availability zones.  | [optional] 

## Methods

### NewProtectionRuleResourcesAvailabilityZoneConnectivityListInner

`func NewProtectionRuleResourcesAvailabilityZoneConnectivityListInner() *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner`

NewProtectionRuleResourcesAvailabilityZoneConnectivityListInner instantiates a new ProtectionRuleResourcesAvailabilityZoneConnectivityListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerWithDefaults

`func NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerWithDefaults() *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner`

NewProtectionRuleResourcesAvailabilityZoneConnectivityListInnerWithDefaults instantiates a new ProtectionRuleResourcesAvailabilityZoneConnectivityListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetDestinationAvailabilityZoneIndex() int32`

GetDestinationAvailabilityZoneIndex returns the DestinationAvailabilityZoneIndex field if non-nil, zero value otherwise.

### GetDestinationAvailabilityZoneIndexOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetDestinationAvailabilityZoneIndexOk() (*int32, bool)`

GetDestinationAvailabilityZoneIndexOk returns a tuple with the DestinationAvailabilityZoneIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) SetDestinationAvailabilityZoneIndex(v int32)`

SetDestinationAvailabilityZoneIndex sets DestinationAvailabilityZoneIndex field to given value.

### HasDestinationAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) HasDestinationAvailabilityZoneIndex() bool`

HasDestinationAvailabilityZoneIndex returns a boolean if a field has been set.

### GetSourceAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetSourceAvailabilityZoneIndex() int32`

GetSourceAvailabilityZoneIndex returns the SourceAvailabilityZoneIndex field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneIndexOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetSourceAvailabilityZoneIndexOk() (*int32, bool)`

GetSourceAvailabilityZoneIndexOk returns a tuple with the SourceAvailabilityZoneIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) SetSourceAvailabilityZoneIndex(v int32)`

SetSourceAvailabilityZoneIndex sets SourceAvailabilityZoneIndex field to given value.

### HasSourceAvailabilityZoneIndex

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) HasSourceAvailabilityZoneIndex() bool`

HasSourceAvailabilityZoneIndex returns a boolean if a field has been set.

### GetSnapshotScheduleList

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetSnapshotScheduleList() []ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner`

GetSnapshotScheduleList returns the SnapshotScheduleList field if non-nil, zero value otherwise.

### GetSnapshotScheduleListOk

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) GetSnapshotScheduleListOk() (*[]ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner, bool)`

GetSnapshotScheduleListOk returns a tuple with the SnapshotScheduleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotScheduleList

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) SetSnapshotScheduleList(v []ProtectionRuleResourcesAvailabilityZoneConnectivityListInnerSnapshotScheduleListInner)`

SetSnapshotScheduleList sets SnapshotScheduleList field to given value.

### HasSnapshotScheduleList

`func (o *ProtectionRuleResourcesAvailabilityZoneConnectivityListInner) HasSnapshotScheduleList() bool`

HasSnapshotScheduleList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


