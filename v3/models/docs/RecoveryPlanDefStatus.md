# RecoveryPlanDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Recovery Plan name | 
**RecoveryAvailabilityZoneOrderList** | [**[]RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner**](RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner.md) | List of recovery Availability Zones order list. Each unique Availability Zone order list will be reported. A cluster might also be specified for each Availability Zone in the Availability Zones order list, in case the entity is protected for replication to/from a cluster in the Protection Policy.  | 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**State** | Pointer to **string** | The state of the Recovery Plan entity. | [optional] 
**LatestValidationTime** | Pointer to **time.Time** | Time when latest validation was done for Recovery Plan. | [optional] 
**LatestTestTime** | Pointer to **time.Time** | Time when latest test was done for Recovery Plan. | [optional] 
**Resources** | [**RecoveryPlanResources**](RecoveryPlanResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the Recovery Plan. | [optional] 

## Methods

### NewRecoveryPlanDefStatus

`func NewRecoveryPlanDefStatus(name string, recoveryAvailabilityZoneOrderList []RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner, resources RecoveryPlanResources, ) *RecoveryPlanDefStatus`

NewRecoveryPlanDefStatus instantiates a new RecoveryPlanDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanDefStatusWithDefaults

`func NewRecoveryPlanDefStatusWithDefaults() *RecoveryPlanDefStatus`

NewRecoveryPlanDefStatusWithDefaults instantiates a new RecoveryPlanDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryPlanDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryPlanDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryPlanDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetRecoveryAvailabilityZoneOrderList

`func (o *RecoveryPlanDefStatus) GetRecoveryAvailabilityZoneOrderList() []RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner`

GetRecoveryAvailabilityZoneOrderList returns the RecoveryAvailabilityZoneOrderList field if non-nil, zero value otherwise.

### GetRecoveryAvailabilityZoneOrderListOk

`func (o *RecoveryPlanDefStatus) GetRecoveryAvailabilityZoneOrderListOk() (*[]RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner, bool)`

GetRecoveryAvailabilityZoneOrderListOk returns a tuple with the RecoveryAvailabilityZoneOrderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAvailabilityZoneOrderList

`func (o *RecoveryPlanDefStatus) SetRecoveryAvailabilityZoneOrderList(v []RecoveryPlanDefStatusRecoveryAvailabilityZoneOrderListInner)`

SetRecoveryAvailabilityZoneOrderList sets RecoveryAvailabilityZoneOrderList field to given value.


### GetMessageList

`func (o *RecoveryPlanDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *RecoveryPlanDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *RecoveryPlanDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *RecoveryPlanDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *RecoveryPlanDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecoveryPlanDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecoveryPlanDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RecoveryPlanDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLatestValidationTime

`func (o *RecoveryPlanDefStatus) GetLatestValidationTime() time.Time`

GetLatestValidationTime returns the LatestValidationTime field if non-nil, zero value otherwise.

### GetLatestValidationTimeOk

`func (o *RecoveryPlanDefStatus) GetLatestValidationTimeOk() (*time.Time, bool)`

GetLatestValidationTimeOk returns a tuple with the LatestValidationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValidationTime

`func (o *RecoveryPlanDefStatus) SetLatestValidationTime(v time.Time)`

SetLatestValidationTime sets LatestValidationTime field to given value.

### HasLatestValidationTime

`func (o *RecoveryPlanDefStatus) HasLatestValidationTime() bool`

HasLatestValidationTime returns a boolean if a field has been set.

### GetLatestTestTime

`func (o *RecoveryPlanDefStatus) GetLatestTestTime() time.Time`

GetLatestTestTime returns the LatestTestTime field if non-nil, zero value otherwise.

### GetLatestTestTimeOk

`func (o *RecoveryPlanDefStatus) GetLatestTestTimeOk() (*time.Time, bool)`

GetLatestTestTimeOk returns a tuple with the LatestTestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTestTime

`func (o *RecoveryPlanDefStatus) SetLatestTestTime(v time.Time)`

SetLatestTestTime sets LatestTestTime field to given value.

### HasLatestTestTime

`func (o *RecoveryPlanDefStatus) HasLatestTestTime() bool`

HasLatestTestTime returns a boolean if a field has been set.

### GetResources

`func (o *RecoveryPlanDefStatus) GetResources() RecoveryPlanResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RecoveryPlanDefStatus) GetResourcesOk() (*RecoveryPlanResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RecoveryPlanDefStatus) SetResources(v RecoveryPlanResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *RecoveryPlanDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecoveryPlanDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecoveryPlanDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecoveryPlanDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


