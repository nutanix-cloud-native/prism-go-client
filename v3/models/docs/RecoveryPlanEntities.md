# RecoveryPlanEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitiesPerAvailabilityZoneList** | [**[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner**](RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner.md) | List of entities and recovery Availability Zones order list per local Availability Zone. In case of public cloud, entity list and recovery Availability Zones order for each of the Availability Zones will be reported. For example, Let AZ1 and AZ2 be the two Availability Zones for a public cloud and OnPrem be the on-premise Availability Zone. If VMs VM1 and VM2 on OnPrem is protected by a Protection Rule with Availabilty Zone order [OnPrem, AZ1, AZ2], then entities_per_availability_zone_list will be [{AZ1, [[OnPrem, AZ1, AZ2]], [{VM1}, {VM2}]},  {AZ2, [[OnPrem, AZ1, AZ2]], [{VM1}, {VM2}]}]. All the Availability Zones should have same Availability Zone order for a valid Recovery Plan.  | 

## Methods

### NewRecoveryPlanEntities

`func NewRecoveryPlanEntities(entitiesPerAvailabilityZoneList []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner, ) *RecoveryPlanEntities`

NewRecoveryPlanEntities instantiates a new RecoveryPlanEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanEntitiesWithDefaults

`func NewRecoveryPlanEntitiesWithDefaults() *RecoveryPlanEntities`

NewRecoveryPlanEntitiesWithDefaults instantiates a new RecoveryPlanEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitiesPerAvailabilityZoneList

`func (o *RecoveryPlanEntities) GetEntitiesPerAvailabilityZoneList() []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner`

GetEntitiesPerAvailabilityZoneList returns the EntitiesPerAvailabilityZoneList field if non-nil, zero value otherwise.

### GetEntitiesPerAvailabilityZoneListOk

`func (o *RecoveryPlanEntities) GetEntitiesPerAvailabilityZoneListOk() (*[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner, bool)`

GetEntitiesPerAvailabilityZoneListOk returns a tuple with the EntitiesPerAvailabilityZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesPerAvailabilityZoneList

`func (o *RecoveryPlanEntities) SetEntitiesPerAvailabilityZoneList(v []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner)`

SetEntitiesPerAvailabilityZoneList sets EntitiesPerAvailabilityZoneList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


