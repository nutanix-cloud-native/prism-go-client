# MigrationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**AvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 

## Methods

### NewMigrationTarget

`func NewMigrationTarget(availabilityZoneReference AvailabilityZoneReference, ) *MigrationTarget`

NewMigrationTarget instantiates a new MigrationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationTargetWithDefaults

`func NewMigrationTargetWithDefaults() *MigrationTarget`

NewMigrationTargetWithDefaults instantiates a new MigrationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterReference

`func (o *MigrationTarget) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *MigrationTarget) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *MigrationTarget) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *MigrationTarget) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *MigrationTarget) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *MigrationTarget) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *MigrationTarget) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


