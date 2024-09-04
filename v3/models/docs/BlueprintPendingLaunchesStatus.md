# BlueprintPendingLaunchesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationUuid** | Pointer to **string** | application uuid present after the status moves to success | [optional] 
**AppName** | Pointer to **string** | name of the application | [optional] 
**State** | Pointer to **string** | Status of the launch | [optional] 
**BlueprintUuid** | Pointer to **string** | blueprint uuid | [optional] 
**Details** | Pointer to **string** | additional details about the status of launch | [optional] 
**Milestone** | Pointer to **string** | gives more granularity in status | [optional] 
**BpName** | Pointer to **string** | name of the blueprint | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list for app blueprint | [optional] 

## Methods

### NewBlueprintPendingLaunchesStatus

`func NewBlueprintPendingLaunchesStatus() *BlueprintPendingLaunchesStatus`

NewBlueprintPendingLaunchesStatus instantiates a new BlueprintPendingLaunchesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPendingLaunchesStatusWithDefaults

`func NewBlueprintPendingLaunchesStatusWithDefaults() *BlueprintPendingLaunchesStatus`

NewBlueprintPendingLaunchesStatusWithDefaults instantiates a new BlueprintPendingLaunchesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationUuid

`func (o *BlueprintPendingLaunchesStatus) GetApplicationUuid() string`

GetApplicationUuid returns the ApplicationUuid field if non-nil, zero value otherwise.

### GetApplicationUuidOk

`func (o *BlueprintPendingLaunchesStatus) GetApplicationUuidOk() (*string, bool)`

GetApplicationUuidOk returns a tuple with the ApplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUuid

`func (o *BlueprintPendingLaunchesStatus) SetApplicationUuid(v string)`

SetApplicationUuid sets ApplicationUuid field to given value.

### HasApplicationUuid

`func (o *BlueprintPendingLaunchesStatus) HasApplicationUuid() bool`

HasApplicationUuid returns a boolean if a field has been set.

### GetAppName

`func (o *BlueprintPendingLaunchesStatus) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *BlueprintPendingLaunchesStatus) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *BlueprintPendingLaunchesStatus) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *BlueprintPendingLaunchesStatus) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetState

`func (o *BlueprintPendingLaunchesStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintPendingLaunchesStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintPendingLaunchesStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintPendingLaunchesStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBlueprintUuid

`func (o *BlueprintPendingLaunchesStatus) GetBlueprintUuid() string`

GetBlueprintUuid returns the BlueprintUuid field if non-nil, zero value otherwise.

### GetBlueprintUuidOk

`func (o *BlueprintPendingLaunchesStatus) GetBlueprintUuidOk() (*string, bool)`

GetBlueprintUuidOk returns a tuple with the BlueprintUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintUuid

`func (o *BlueprintPendingLaunchesStatus) SetBlueprintUuid(v string)`

SetBlueprintUuid sets BlueprintUuid field to given value.

### HasBlueprintUuid

`func (o *BlueprintPendingLaunchesStatus) HasBlueprintUuid() bool`

HasBlueprintUuid returns a boolean if a field has been set.

### GetDetails

`func (o *BlueprintPendingLaunchesStatus) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BlueprintPendingLaunchesStatus) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BlueprintPendingLaunchesStatus) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BlueprintPendingLaunchesStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMilestone

`func (o *BlueprintPendingLaunchesStatus) GetMilestone() string`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *BlueprintPendingLaunchesStatus) GetMilestoneOk() (*string, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *BlueprintPendingLaunchesStatus) SetMilestone(v string)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *BlueprintPendingLaunchesStatus) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetBpName

`func (o *BlueprintPendingLaunchesStatus) GetBpName() string`

GetBpName returns the BpName field if non-nil, zero value otherwise.

### GetBpNameOk

`func (o *BlueprintPendingLaunchesStatus) GetBpNameOk() (*string, bool)`

GetBpNameOk returns a tuple with the BpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBpName

`func (o *BlueprintPendingLaunchesStatus) SetBpName(v string)`

SetBpName sets BpName field to given value.

### HasBpName

`func (o *BlueprintPendingLaunchesStatus) HasBpName() bool`

HasBpName returns a boolean if a field has been set.

### GetMessageList

`func (o *BlueprintPendingLaunchesStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *BlueprintPendingLaunchesStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *BlueprintPendingLaunchesStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *BlueprintPendingLaunchesStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


