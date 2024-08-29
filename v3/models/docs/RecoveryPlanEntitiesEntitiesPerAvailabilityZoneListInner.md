# RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneUrl** | **string** | URL of the Availability Zone.  | 
**AvailabilityZoneOrderList** | Pointer to [**[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerAvailabilityZoneOrderListInner**](RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerAvailabilityZoneOrderListInner.md) | List of recovery Availability Zone orders for entities in the Recovery Plan. More than one entry in this list indicates entities in Recovery Plan are protected by the Protection Rules with different recovery Availability Zone order and makes Recovery Plan invalid.  | [optional] 
**EntityList** | Pointer to [**[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner**](RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner.md) | List of protected entities associated with Recovery Plan. This list includes the entities which are live and the entities for which Recovery Points are present on the Availability Zone.  | [optional] 

## Methods

### NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner

`func NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner(availabilityZoneUrl string, ) *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner`

NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner instantiates a new RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerWithDefaults

`func NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerWithDefaults() *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner`

NewRecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerWithDefaults instantiates a new RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneUrl

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.


### GetAvailabilityZoneOrderList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetAvailabilityZoneOrderList() []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerAvailabilityZoneOrderListInner`

GetAvailabilityZoneOrderList returns the AvailabilityZoneOrderList field if non-nil, zero value otherwise.

### GetAvailabilityZoneOrderListOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetAvailabilityZoneOrderListOk() (*[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerAvailabilityZoneOrderListInner, bool)`

GetAvailabilityZoneOrderListOk returns a tuple with the AvailabilityZoneOrderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneOrderList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) SetAvailabilityZoneOrderList(v []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerAvailabilityZoneOrderListInner)`

SetAvailabilityZoneOrderList sets AvailabilityZoneOrderList field to given value.

### HasAvailabilityZoneOrderList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) HasAvailabilityZoneOrderList() bool`

HasAvailabilityZoneOrderList returns a boolean if a field has been set.

### GetEntityList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetEntityList() []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) GetEntityListOk() (*[]RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) SetEntityList(v []RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInnerEntityListInner)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *RecoveryPlanEntitiesEntitiesPerAvailabilityZoneListInner) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


