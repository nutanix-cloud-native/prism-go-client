# RecoveryPlanJobResourcesExecutionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAvailabilityZoneList** | [**[]AvailabilityZoneInformation**](AvailabilityZoneInformation.md) | Availability Zones wherein entities need to be recovered.  | 
**FailedAvailabilityZoneList** | [**[]AvailabilityZoneInformation**](AvailabilityZoneInformation.md) | Availability Zones that have failed.  | 
**RecoveryReferenceTime** | Pointer to **time.Time** | Time with respect to which Recovery Plan Job has to be executed. This time will be used as reference time with respect to which latest snapshot will have to be restored in case of failover. For example, if failover is required to be done using snapshot created on or before yesterday &#39;2:00&#39; PM, then recovery_reference_time will be set to this time.  | [optional] 
**ActionType** | **string** | Type of action performed by the Recovery Plan Job. VALIDATE - Performs the validation of the Recovery Plan.            The validation includes checks for the presence of            entities, networks, categories etc. referenced in the            Recovery Plan. MIGRATE - VM would be powered off on the sourece before migrating           it to the recovery Availability Zone. FAILOVER - Restore the entity from the recovery points on the            recovery Availability Zone. TEST_FAILOVER - Same as FAILOVER but on a test network. LIVE_MIGRATE - Migrate without powering off the VM.  | 
**ShouldContinueOnValidationFailure** | Pointer to **bool** | Whether to ignore the validation failures(e.g. Network mapping is missing for some networks on failed Availability Zone, Virtual network missing.) for the Recovery Plan actions MIGRATE, FAILOVER, TEST_FAILOVER and execute the Recovery Plan.  | [optional] [default to false]

## Methods

### NewRecoveryPlanJobResourcesExecutionParameters

`func NewRecoveryPlanJobResourcesExecutionParameters(recoveryAvailabilityZoneList []AvailabilityZoneInformation, failedAvailabilityZoneList []AvailabilityZoneInformation, actionType string, ) *RecoveryPlanJobResourcesExecutionParameters`

NewRecoveryPlanJobResourcesExecutionParameters instantiates a new RecoveryPlanJobResourcesExecutionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobResourcesExecutionParametersWithDefaults

`func NewRecoveryPlanJobResourcesExecutionParametersWithDefaults() *RecoveryPlanJobResourcesExecutionParameters`

NewRecoveryPlanJobResourcesExecutionParametersWithDefaults instantiates a new RecoveryPlanJobResourcesExecutionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAvailabilityZoneList

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetRecoveryAvailabilityZoneList() []AvailabilityZoneInformation`

GetRecoveryAvailabilityZoneList returns the RecoveryAvailabilityZoneList field if non-nil, zero value otherwise.

### GetRecoveryAvailabilityZoneListOk

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetRecoveryAvailabilityZoneListOk() (*[]AvailabilityZoneInformation, bool)`

GetRecoveryAvailabilityZoneListOk returns a tuple with the RecoveryAvailabilityZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAvailabilityZoneList

`func (o *RecoveryPlanJobResourcesExecutionParameters) SetRecoveryAvailabilityZoneList(v []AvailabilityZoneInformation)`

SetRecoveryAvailabilityZoneList sets RecoveryAvailabilityZoneList field to given value.


### GetFailedAvailabilityZoneList

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetFailedAvailabilityZoneList() []AvailabilityZoneInformation`

GetFailedAvailabilityZoneList returns the FailedAvailabilityZoneList field if non-nil, zero value otherwise.

### GetFailedAvailabilityZoneListOk

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetFailedAvailabilityZoneListOk() (*[]AvailabilityZoneInformation, bool)`

GetFailedAvailabilityZoneListOk returns a tuple with the FailedAvailabilityZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAvailabilityZoneList

`func (o *RecoveryPlanJobResourcesExecutionParameters) SetFailedAvailabilityZoneList(v []AvailabilityZoneInformation)`

SetFailedAvailabilityZoneList sets FailedAvailabilityZoneList field to given value.


### GetRecoveryReferenceTime

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetRecoveryReferenceTime() time.Time`

GetRecoveryReferenceTime returns the RecoveryReferenceTime field if non-nil, zero value otherwise.

### GetRecoveryReferenceTimeOk

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetRecoveryReferenceTimeOk() (*time.Time, bool)`

GetRecoveryReferenceTimeOk returns a tuple with the RecoveryReferenceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryReferenceTime

`func (o *RecoveryPlanJobResourcesExecutionParameters) SetRecoveryReferenceTime(v time.Time)`

SetRecoveryReferenceTime sets RecoveryReferenceTime field to given value.

### HasRecoveryReferenceTime

`func (o *RecoveryPlanJobResourcesExecutionParameters) HasRecoveryReferenceTime() bool`

HasRecoveryReferenceTime returns a boolean if a field has been set.

### GetActionType

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *RecoveryPlanJobResourcesExecutionParameters) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetShouldContinueOnValidationFailure

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetShouldContinueOnValidationFailure() bool`

GetShouldContinueOnValidationFailure returns the ShouldContinueOnValidationFailure field if non-nil, zero value otherwise.

### GetShouldContinueOnValidationFailureOk

`func (o *RecoveryPlanJobResourcesExecutionParameters) GetShouldContinueOnValidationFailureOk() (*bool, bool)`

GetShouldContinueOnValidationFailureOk returns a tuple with the ShouldContinueOnValidationFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldContinueOnValidationFailure

`func (o *RecoveryPlanJobResourcesExecutionParameters) SetShouldContinueOnValidationFailure(v bool)`

SetShouldContinueOnValidationFailure sets ShouldContinueOnValidationFailure field to given value.

### HasShouldContinueOnValidationFailure

`func (o *RecoveryPlanJobResourcesExecutionParameters) HasShouldContinueOnValidationFailure() bool`

HasShouldContinueOnValidationFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


