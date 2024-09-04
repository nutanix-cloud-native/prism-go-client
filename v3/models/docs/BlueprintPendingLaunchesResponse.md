# BlueprintPendingLaunchesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationUuid** | Pointer to **string** | application uuid present after the status moves to success | [optional] 
**AppName** | **string** | name of the application | 
**State** | **string** | state of launch which can be pending, running, success, failure | 
**BlueprintUuid** | Pointer to **string** | blueprint uuid | [optional] 
**Details** | Pointer to **string** | additional details about the status of launch | [optional] 
**Milestone** | **string** | gives more granularity in status | 
**BpName** | **string** | name of the blueprint | 

## Methods

### NewBlueprintPendingLaunchesResponse

`func NewBlueprintPendingLaunchesResponse(appName string, state string, milestone string, bpName string, ) *BlueprintPendingLaunchesResponse`

NewBlueprintPendingLaunchesResponse instantiates a new BlueprintPendingLaunchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPendingLaunchesResponseWithDefaults

`func NewBlueprintPendingLaunchesResponseWithDefaults() *BlueprintPendingLaunchesResponse`

NewBlueprintPendingLaunchesResponseWithDefaults instantiates a new BlueprintPendingLaunchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationUuid

`func (o *BlueprintPendingLaunchesResponse) GetApplicationUuid() string`

GetApplicationUuid returns the ApplicationUuid field if non-nil, zero value otherwise.

### GetApplicationUuidOk

`func (o *BlueprintPendingLaunchesResponse) GetApplicationUuidOk() (*string, bool)`

GetApplicationUuidOk returns a tuple with the ApplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUuid

`func (o *BlueprintPendingLaunchesResponse) SetApplicationUuid(v string)`

SetApplicationUuid sets ApplicationUuid field to given value.

### HasApplicationUuid

`func (o *BlueprintPendingLaunchesResponse) HasApplicationUuid() bool`

HasApplicationUuid returns a boolean if a field has been set.

### GetAppName

`func (o *BlueprintPendingLaunchesResponse) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *BlueprintPendingLaunchesResponse) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *BlueprintPendingLaunchesResponse) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetState

`func (o *BlueprintPendingLaunchesResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintPendingLaunchesResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintPendingLaunchesResponse) SetState(v string)`

SetState sets State field to given value.


### GetBlueprintUuid

`func (o *BlueprintPendingLaunchesResponse) GetBlueprintUuid() string`

GetBlueprintUuid returns the BlueprintUuid field if non-nil, zero value otherwise.

### GetBlueprintUuidOk

`func (o *BlueprintPendingLaunchesResponse) GetBlueprintUuidOk() (*string, bool)`

GetBlueprintUuidOk returns a tuple with the BlueprintUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintUuid

`func (o *BlueprintPendingLaunchesResponse) SetBlueprintUuid(v string)`

SetBlueprintUuid sets BlueprintUuid field to given value.

### HasBlueprintUuid

`func (o *BlueprintPendingLaunchesResponse) HasBlueprintUuid() bool`

HasBlueprintUuid returns a boolean if a field has been set.

### GetDetails

`func (o *BlueprintPendingLaunchesResponse) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BlueprintPendingLaunchesResponse) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BlueprintPendingLaunchesResponse) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BlueprintPendingLaunchesResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMilestone

`func (o *BlueprintPendingLaunchesResponse) GetMilestone() string`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *BlueprintPendingLaunchesResponse) GetMilestoneOk() (*string, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *BlueprintPendingLaunchesResponse) SetMilestone(v string)`

SetMilestone sets Milestone field to given value.


### GetBpName

`func (o *BlueprintPendingLaunchesResponse) GetBpName() string`

GetBpName returns the BpName field if non-nil, zero value otherwise.

### GetBpNameOk

`func (o *BlueprintPendingLaunchesResponse) GetBpNameOk() (*string, bool)`

GetBpNameOk returns a tuple with the BpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpName

`func (o *BlueprintPendingLaunchesResponse) SetBpName(v string)`

SetBpName sets BpName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


