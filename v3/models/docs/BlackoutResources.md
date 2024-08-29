# BlackoutResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeEntityList** | [**[]BlackoutResourcesScopeEntityListInner**](BlackoutResourcesScopeEntityListInner.md) | The list of 1) entity type, or 2) entity type + uuid_list, will define the targeted entities.  Any entities contained inside the targeted entities are affected by this blackout schedule. If it is case 1), where the uuid_list is empty, then, it means the blackout applies to all the entities of this entity_type.  For example, blackout for all clusters, the entity_type &#x3D; \&quot;cluster\&quot;, the uuid_list will be empty.  If the uuid_list is not empty, then, the blackout is applied to only these entities in the uuid_list.  | 
**ScheduleList** | [**[]Schedule**](Schedule.md) | A list of time schedules for the blackout.  For example, if the blackout is for every Monday and Friday, 10:00 - 11:00 am, 3:00-5:00 pm, then, there will be 2 items:   schedule1: day_of_week &#x3D; Monday,Friday, Time &#x3D; 10:00 - 11:00 am   schedule2: day_of_week &#x3D; Monday,Friday, 3:00 - 5:00 pm  | 
**FeatureList** | **[]string** | A list of features that this blackout is applied to. | 

## Methods

### NewBlackoutResources

`func NewBlackoutResources(scopeEntityList []BlackoutResourcesScopeEntityListInner, scheduleList []Schedule, featureList []string, ) *BlackoutResources`

NewBlackoutResources instantiates a new BlackoutResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutResourcesWithDefaults

`func NewBlackoutResourcesWithDefaults() *BlackoutResources`

NewBlackoutResourcesWithDefaults instantiates a new BlackoutResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeEntityList

`func (o *BlackoutResources) GetScopeEntityList() []BlackoutResourcesScopeEntityListInner`

GetScopeEntityList returns the ScopeEntityList field if non-nil, zero value otherwise.

### GetScopeEntityListOk

`func (o *BlackoutResources) GetScopeEntityListOk() (*[]BlackoutResourcesScopeEntityListInner, bool)`

GetScopeEntityListOk returns a tuple with the ScopeEntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEntityList

`func (o *BlackoutResources) SetScopeEntityList(v []BlackoutResourcesScopeEntityListInner)`

SetScopeEntityList sets ScopeEntityList field to given value.


### GetScheduleList

`func (o *BlackoutResources) GetScheduleList() []Schedule`

GetScheduleList returns the ScheduleList field if non-nil, zero value otherwise.

### GetScheduleListOk

`func (o *BlackoutResources) GetScheduleListOk() (*[]Schedule, bool)`

GetScheduleListOk returns a tuple with the ScheduleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleList

`func (o *BlackoutResources) SetScheduleList(v []Schedule)`

SetScheduleList sets ScheduleList field to given value.


### GetFeatureList

`func (o *BlackoutResources) GetFeatureList() []string`

GetFeatureList returns the FeatureList field if non-nil, zero value otherwise.

### GetFeatureListOk

`func (o *BlackoutResources) GetFeatureListOk() (*[]string, bool)`

GetFeatureListOk returns a tuple with the FeatureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureList

`func (o *BlackoutResources) SetFeatureList(v []string)`

SetFeatureList sets FeatureList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


