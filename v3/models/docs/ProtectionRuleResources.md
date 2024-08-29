# ProtectionRuleResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **string** | Time of the day, the policy will be started. This is in \&quot;&lt;x&gt;h:&lt;y&gt;m\&quot; format. The values must be between 00h:00m and 23h:59m. For example user specified 18h:00m and the current time is 17h:00m then the first snapshot will be captured at 18h:00m. If the current time is 19h:00m then the first snapshot will be captured at 18h:00m next day. If not set, policy will be applicable immediately.  | [optional] 
**AvailabilityZoneConnectivityList** | [**[]ProtectionRuleResourcesAvailabilityZoneConnectivityListInner**](ProtectionRuleResourcesAvailabilityZoneConnectivityListInner.md) | This encodes the datapipes between various availability zones and the backup policy of the pipes. For example, [1, 2, 3600], [2, 3, 3600], [4, 5, 15000], [2, 6, 4200]. Note 2 here means the entry at index 2 in the ordered_availability_zone_list. And 3600 is the RPO (Recovery Point Objective) in seconds between the two availability zones.  | 
**OrderedAvailabilityZoneList** | [**[]ProtectionRuleResourcesOrderedAvailabilityZoneListInner**](ProtectionRuleResourcesOrderedAvailabilityZoneListInner.md) | A list of availability zones, each of which, receives a replica of the data for the entities protected by this protection rule. The order of the availability zones in the list determines the preference order (highest to lowest) for the entity to run in the case of failure of one or more availability zones. For example, if this list is [A, B, C, D], then the entity will prefer to run on A unless A has failed, in which case, the entity will run on the second choice, B. Failover of the entity is not automatic but this information is used by DR runbook to failover an entity to the desired availability zone.  | 
**PrimaryLocationList** | Pointer to **[]int32** | List of indexes in ordered_availability_zone_list, which constitutes the primary locations of this Protection Rule. This field is mandatory if more than two Availability Zones are provided in the policy.  | [optional] 
**CategoryFilter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 

## Methods

### NewProtectionRuleResources

`func NewProtectionRuleResources(availabilityZoneConnectivityList []ProtectionRuleResourcesAvailabilityZoneConnectivityListInner, orderedAvailabilityZoneList []ProtectionRuleResourcesOrderedAvailabilityZoneListInner, ) *ProtectionRuleResources`

NewProtectionRuleResources instantiates a new ProtectionRuleResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleResourcesWithDefaults

`func NewProtectionRuleResourcesWithDefaults() *ProtectionRuleResources`

NewProtectionRuleResourcesWithDefaults instantiates a new ProtectionRuleResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *ProtectionRuleResources) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProtectionRuleResources) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProtectionRuleResources) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProtectionRuleResources) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAvailabilityZoneConnectivityList

`func (o *ProtectionRuleResources) GetAvailabilityZoneConnectivityList() []ProtectionRuleResourcesAvailabilityZoneConnectivityListInner`

GetAvailabilityZoneConnectivityList returns the AvailabilityZoneConnectivityList field if non-nil, zero value otherwise.

### GetAvailabilityZoneConnectivityListOk

`func (o *ProtectionRuleResources) GetAvailabilityZoneConnectivityListOk() (*[]ProtectionRuleResourcesAvailabilityZoneConnectivityListInner, bool)`

GetAvailabilityZoneConnectivityListOk returns a tuple with the AvailabilityZoneConnectivityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneConnectivityList

`func (o *ProtectionRuleResources) SetAvailabilityZoneConnectivityList(v []ProtectionRuleResourcesAvailabilityZoneConnectivityListInner)`

SetAvailabilityZoneConnectivityList sets AvailabilityZoneConnectivityList field to given value.


### GetOrderedAvailabilityZoneList

`func (o *ProtectionRuleResources) GetOrderedAvailabilityZoneList() []ProtectionRuleResourcesOrderedAvailabilityZoneListInner`

GetOrderedAvailabilityZoneList returns the OrderedAvailabilityZoneList field if non-nil, zero value otherwise.

### GetOrderedAvailabilityZoneListOk

`func (o *ProtectionRuleResources) GetOrderedAvailabilityZoneListOk() (*[]ProtectionRuleResourcesOrderedAvailabilityZoneListInner, bool)`

GetOrderedAvailabilityZoneListOk returns a tuple with the OrderedAvailabilityZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedAvailabilityZoneList

`func (o *ProtectionRuleResources) SetOrderedAvailabilityZoneList(v []ProtectionRuleResourcesOrderedAvailabilityZoneListInner)`

SetOrderedAvailabilityZoneList sets OrderedAvailabilityZoneList field to given value.


### GetPrimaryLocationList

`func (o *ProtectionRuleResources) GetPrimaryLocationList() []int32`

GetPrimaryLocationList returns the PrimaryLocationList field if non-nil, zero value otherwise.

### GetPrimaryLocationListOk

`func (o *ProtectionRuleResources) GetPrimaryLocationListOk() (*[]int32, bool)`

GetPrimaryLocationListOk returns a tuple with the PrimaryLocationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLocationList

`func (o *ProtectionRuleResources) SetPrimaryLocationList(v []int32)`

SetPrimaryLocationList sets PrimaryLocationList field to given value.

### HasPrimaryLocationList

`func (o *ProtectionRuleResources) HasPrimaryLocationList() bool`

HasPrimaryLocationList returns a boolean if a field has been set.

### GetCategoryFilter

`func (o *ProtectionRuleResources) GetCategoryFilter() CategoryFilter`

GetCategoryFilter returns the CategoryFilter field if non-nil, zero value otherwise.

### GetCategoryFilterOk

`func (o *ProtectionRuleResources) GetCategoryFilterOk() (*CategoryFilter, bool)`

GetCategoryFilterOk returns a tuple with the CategoryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFilter

`func (o *ProtectionRuleResources) SetCategoryFilter(v CategoryFilter)`

SetCategoryFilter sets CategoryFilter field to given value.

### HasCategoryFilter

`func (o *ProtectionRuleResources) HasCategoryFilter() bool`

HasCategoryFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


