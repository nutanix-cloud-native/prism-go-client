# VmRecoveryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the recovery point. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | Pointer to [**RecoveryPointResources**](RecoveryPointResources.md) |  | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 

## Methods

### NewVmRecoveryPoint

`func NewVmRecoveryPoint() *VmRecoveryPoint`

NewVmRecoveryPoint instantiates a new VmRecoveryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointWithDefaults

`func NewVmRecoveryPointWithDefaults() *VmRecoveryPoint`

NewVmRecoveryPointWithDefaults instantiates a new VmRecoveryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmRecoveryPoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmRecoveryPoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmRecoveryPoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmRecoveryPoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClusterReference

`func (o *VmRecoveryPoint) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VmRecoveryPoint) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VmRecoveryPoint) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VmRecoveryPoint) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *VmRecoveryPoint) GetResources() RecoveryPointResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VmRecoveryPoint) GetResourcesOk() (*RecoveryPointResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VmRecoveryPoint) SetResources(v RecoveryPointResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VmRecoveryPoint) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *VmRecoveryPoint) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VmRecoveryPoint) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VmRecoveryPoint) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *VmRecoveryPoint) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


