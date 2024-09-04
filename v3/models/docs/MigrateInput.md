# MigrateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**VmList** | [**[]MigrateInputVmListInner**](MigrateInputVmListInner.md) |  | 
**TargetClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**TargetAvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**IsLiveMigration** | Pointer to **bool** | Whether to do live migration of the entity. This is applicable only when the entity is protected with sync protection policy.  | [optional] 

## Methods

### NewMigrateInput

`func NewMigrateInput(sourceAvailabilityZoneReference AvailabilityZoneReference, vmList []MigrateInputVmListInner, targetAvailabilityZoneReference AvailabilityZoneReference, ) *MigrateInput`

NewMigrateInput instantiates a new MigrateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateInputWithDefaults

`func NewMigrateInputWithDefaults() *MigrateInput`

NewMigrateInputWithDefaults instantiates a new MigrateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAvailabilityZoneReference

`func (o *MigrateInput) GetSourceAvailabilityZoneReference() AvailabilityZoneReference`

GetSourceAvailabilityZoneReference returns the SourceAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetSourceAvailabilityZoneReferenceOk

`func (o *MigrateInput) GetSourceAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetSourceAvailabilityZoneReferenceOk returns a tuple with the SourceAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAvailabilityZoneReference

`func (o *MigrateInput) SetSourceAvailabilityZoneReference(v AvailabilityZoneReference)`

SetSourceAvailabilityZoneReference sets SourceAvailabilityZoneReference field to given value.


### GetVmList

`func (o *MigrateInput) GetVmList() []MigrateInputVmListInner`

GetVmList returns the VmList field if non-nil, zero value otherwise.

### GetVmListOk

`func (o *MigrateInput) GetVmListOk() (*[]MigrateInputVmListInner, bool)`

GetVmListOk returns a tuple with the VmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmList

`func (o *MigrateInput) SetVmList(v []MigrateInputVmListInner)`

SetVmList sets VmList field to given value.


### GetTargetClusterReference

`func (o *MigrateInput) GetTargetClusterReference() ClusterReference`

GetTargetClusterReference returns the TargetClusterReference field if non-nil, zero value otherwise.

### GetTargetClusterReferenceOk

`func (o *MigrateInput) GetTargetClusterReferenceOk() (*ClusterReference, bool)`

GetTargetClusterReferenceOk returns a tuple with the TargetClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterReference

`func (o *MigrateInput) SetTargetClusterReference(v ClusterReference)`

SetTargetClusterReference sets TargetClusterReference field to given value.

### HasTargetClusterReference

`func (o *MigrateInput) HasTargetClusterReference() bool`

HasTargetClusterReference returns a boolean if a field has been set.

### GetTargetAvailabilityZoneReference

`func (o *MigrateInput) GetTargetAvailabilityZoneReference() AvailabilityZoneReference`

GetTargetAvailabilityZoneReference returns the TargetAvailabilityZoneReference field if non-nil, zero value otherwise.

### GetTargetAvailabilityZoneReferenceOk

`func (o *MigrateInput) GetTargetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetTargetAvailabilityZoneReferenceOk returns a tuple with the TargetAvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAvailabilityZoneReference

`func (o *MigrateInput) SetTargetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetTargetAvailabilityZoneReference sets TargetAvailabilityZoneReference field to given value.


### GetIsLiveMigration

`func (o *MigrateInput) GetIsLiveMigration() bool`

GetIsLiveMigration returns the IsLiveMigration field if non-nil, zero value otherwise.

### GetIsLiveMigrationOk

`func (o *MigrateInput) GetIsLiveMigrationOk() (*bool, bool)`

GetIsLiveMigrationOk returns a tuple with the IsLiveMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLiveMigration

`func (o *MigrateInput) SetIsLiveMigration(v bool)`

SetIsLiveMigration sets IsLiveMigration field to given value.

### HasIsLiveMigration

`func (o *MigrateInput) HasIsLiveMigration() bool`

HasIsLiveMigration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


