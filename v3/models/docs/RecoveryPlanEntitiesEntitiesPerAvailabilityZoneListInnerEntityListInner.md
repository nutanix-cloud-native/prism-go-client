# RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionStatus** | Pointer to **string** | Protection status of the entity. | [optional] 
**Recoverability** | Pointer to [**RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerRecoverability**](RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerRecoverability.md) |  | [optional] 
**RecoveryAvailabilityZoneOrderIndex** | Pointer to **int32** | Index of recovery Availability Zone order for this entity in availability_zone_order_list. Index starts at 0.  | [optional] 
**IsRecoveryPoint** | **bool** | On the Availability Zone, whether the live entity is present or the Recovery Points for the entity are present. This will be set to true, only if the Recovery Points for the entity are present on the Availability Zone and not the live entity.  | 
**ReplicationStatus** | Pointer to **string** | Replication status of the entity for which synchronous replication is enabled.  | [optional] 
**AnyEntityReference** | [**Reference**](Reference.md) |  | 

## Methods

### NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner

`func NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner(isRecoveryPoint bool, anyEntityReference Reference, ) *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner`

NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner instantiates a new RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerWithDefaults

`func NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerWithDefaults() *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner`

NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerWithDefaults instantiates a new RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetProtectionStatus() string`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetProtectionStatusOk() (*string, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetProtectionStatus(v string)`

SetProtectionStatus sets ProtectionStatus field to given value.

### HasProtectionStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) HasProtectionStatus() bool`

HasProtectionStatus returns a boolean if a field has been set.

### GetRecoverability

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetRecoverability() RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerRecoverability`

GetRecoverability returns the Recoverability field if non-nil, zero value otherwise.

### GetRecoverabilityOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetRecoverabilityOk() (*RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerRecoverability, bool)`

GetRecoverabilityOk returns a tuple with the Recoverability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverability

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetRecoverability(v RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInnerRecoverability)`

SetRecoverability sets Recoverability field to given value.

### HasRecoverability

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) HasRecoverability() bool`

HasRecoverability returns a boolean if a field has been set.

### GetRecoveryAvailabilityZoneOrderIndex

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetRecoveryAvailabilityZoneOrderIndex() int32`

GetRecoveryAvailabilityZoneOrderIndex returns the RecoveryAvailabilityZoneOrderIndex field if non-nil, zero value otherwise.

### GetRecoveryAvailabilityZoneOrderIndexOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetRecoveryAvailabilityZoneOrderIndexOk() (*int32, bool)`

GetRecoveryAvailabilityZoneOrderIndexOk returns a tuple with the RecoveryAvailabilityZoneOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAvailabilityZoneOrderIndex

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetRecoveryAvailabilityZoneOrderIndex(v int32)`

SetRecoveryAvailabilityZoneOrderIndex sets RecoveryAvailabilityZoneOrderIndex field to given value.

### HasRecoveryAvailabilityZoneOrderIndex

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) HasRecoveryAvailabilityZoneOrderIndex() bool`

HasRecoveryAvailabilityZoneOrderIndex returns a boolean if a field has been set.

### GetIsRecoveryPoint

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetIsRecoveryPoint() bool`

GetIsRecoveryPoint returns the IsRecoveryPoint field if non-nil, zero value otherwise.

### GetIsRecoveryPointOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetIsRecoveryPointOk() (*bool, bool)`

GetIsRecoveryPointOk returns a tuple with the IsRecoveryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecoveryPoint

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetIsRecoveryPoint(v bool)`

SetIsRecoveryPoint sets IsRecoveryPoint field to given value.


### GetReplicationStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetAnyEntityReference

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetAnyEntityReference() Reference`

GetAnyEntityReference returns the AnyEntityReference field if non-nil, zero value otherwise.

### GetAnyEntityReferenceOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) GetAnyEntityReferenceOk() (*Reference, bool)`

GetAnyEntityReferenceOk returns a tuple with the AnyEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEntityReference

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner) SetAnyEntityReference(v Reference)`

SetAnyEntityReference sets AnyEntityReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


