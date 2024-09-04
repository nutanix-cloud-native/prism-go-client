# RecoveryPlanJobDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationInformation** | Pointer to [**RecoveryPlanJobDefStatusValidationInformation**](RecoveryPlanJobDefStatusValidationInformation.md) |  | [optional] 
**Name** | **string** | Recovery Plan Job name. | 
**CleanupStatus** | Pointer to [**RecoveryPlanJobExecutionPhasesStatus**](RecoveryPlanJobExecutionPhasesStatus.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | Time when the Recovery Plan Job was created. | [optional] 
**ExecutionStatus** | Pointer to [**RecoveryPlanJobExecutionPhasesStatus**](RecoveryPlanJobExecutionPhasesStatus.md) |  | [optional] 
**RecoveryPlanSpecification** | Pointer to [**RecoveryPlanJobDefStatusRecoveryPlanSpecification**](RecoveryPlanJobDefStatusRecoveryPlanSpecification.md) |  | [optional] 
**ParentRecoveryPlanJobReference** | Pointer to [**RecoveryPlanJobReference**](RecoveryPlanJobReference.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | Time when the Recovery Plan Job execution ended. | [optional] 
**RootRecoveryPlanJobReference** | Pointer to [**RecoveryPlanJobReference**](RecoveryPlanJobReference.md) |  | [optional] 
**WitnessAddress** | Pointer to **string** | Address of the witness, which has triggered this Recovery Plan Job. This will be same as Availability Zone URL, on which witness is deployed.  | [optional] 
**Resources** | [**RecoveryPlanJobResources**](RecoveryPlanJobResources.md) |  | 

## Methods

### NewRecoveryPlanJobDefStatus

`func NewRecoveryPlanJobDefStatus(name string, resources RecoveryPlanJobResources, ) *RecoveryPlanJobDefStatus`

NewRecoveryPlanJobDefStatus instantiates a new RecoveryPlanJobDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobDefStatusWithDefaults

`func NewRecoveryPlanJobDefStatusWithDefaults() *RecoveryPlanJobDefStatus`

NewRecoveryPlanJobDefStatusWithDefaults instantiates a new RecoveryPlanJobDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationInformation

`func (o *RecoveryPlanJobDefStatus) GetValidationInformation() RecoveryPlanJobDefStatusValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *RecoveryPlanJobDefStatus) GetValidationInformationOk() (*RecoveryPlanJobDefStatusValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *RecoveryPlanJobDefStatus) SetValidationInformation(v RecoveryPlanJobDefStatusValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *RecoveryPlanJobDefStatus) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### GetName

`func (o *RecoveryPlanJobDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryPlanJobDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryPlanJobDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetCleanupStatus

`func (o *RecoveryPlanJobDefStatus) GetCleanupStatus() RecoveryPlanJobExecutionPhasesStatus`

GetCleanupStatus returns the CleanupStatus field if non-nil, zero value otherwise.

### GetCleanupStatusOk

`func (o *RecoveryPlanJobDefStatus) GetCleanupStatusOk() (*RecoveryPlanJobExecutionPhasesStatus, bool)`

GetCleanupStatusOk returns a tuple with the CleanupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupStatus

`func (o *RecoveryPlanJobDefStatus) SetCleanupStatus(v RecoveryPlanJobExecutionPhasesStatus)`

SetCleanupStatus sets CleanupStatus field to given value.

### HasCleanupStatus

`func (o *RecoveryPlanJobDefStatus) HasCleanupStatus() bool`

HasCleanupStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *RecoveryPlanJobDefStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RecoveryPlanJobDefStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RecoveryPlanJobDefStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RecoveryPlanJobDefStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetExecutionStatus

`func (o *RecoveryPlanJobDefStatus) GetExecutionStatus() RecoveryPlanJobExecutionPhasesStatus`

GetExecutionStatus returns the ExecutionStatus field if non-nil, zero value otherwise.

### GetExecutionStatusOk

`func (o *RecoveryPlanJobDefStatus) GetExecutionStatusOk() (*RecoveryPlanJobExecutionPhasesStatus, bool)`

GetExecutionStatusOk returns a tuple with the ExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStatus

`func (o *RecoveryPlanJobDefStatus) SetExecutionStatus(v RecoveryPlanJobExecutionPhasesStatus)`

SetExecutionStatus sets ExecutionStatus field to given value.

### HasExecutionStatus

`func (o *RecoveryPlanJobDefStatus) HasExecutionStatus() bool`

HasExecutionStatus returns a boolean if a field has been set.

### GetRecoveryPlanSpecification

`func (o *RecoveryPlanJobDefStatus) GetRecoveryPlanSpecification() RecoveryPlanJobDefStatusRecoveryPlanSpecification`

GetRecoveryPlanSpecification returns the RecoveryPlanSpecification field if non-nil, zero value otherwise.

### GetRecoveryPlanSpecificationOk

`func (o *RecoveryPlanJobDefStatus) GetRecoveryPlanSpecificationOk() (*RecoveryPlanJobDefStatusRecoveryPlanSpecification, bool)`

GetRecoveryPlanSpecificationOk returns a tuple with the RecoveryPlanSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPlanSpecification

`func (o *RecoveryPlanJobDefStatus) SetRecoveryPlanSpecification(v RecoveryPlanJobDefStatusRecoveryPlanSpecification)`

SetRecoveryPlanSpecification sets RecoveryPlanSpecification field to given value.

### HasRecoveryPlanSpecification

`func (o *RecoveryPlanJobDefStatus) HasRecoveryPlanSpecification() bool`

HasRecoveryPlanSpecification returns a boolean if a field has been set.

### GetParentRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) GetParentRecoveryPlanJobReference() RecoveryPlanJobReference`

GetParentRecoveryPlanJobReference returns the ParentRecoveryPlanJobReference field if non-nil, zero value otherwise.

### GetParentRecoveryPlanJobReferenceOk

`func (o *RecoveryPlanJobDefStatus) GetParentRecoveryPlanJobReferenceOk() (*RecoveryPlanJobReference, bool)`

GetParentRecoveryPlanJobReferenceOk returns a tuple with the ParentRecoveryPlanJobReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) SetParentRecoveryPlanJobReference(v RecoveryPlanJobReference)`

SetParentRecoveryPlanJobReference sets ParentRecoveryPlanJobReference field to given value.

### HasParentRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) HasParentRecoveryPlanJobReference() bool`

HasParentRecoveryPlanJobReference returns a boolean if a field has been set.

### GetEndTime

`func (o *RecoveryPlanJobDefStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RecoveryPlanJobDefStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RecoveryPlanJobDefStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RecoveryPlanJobDefStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRootRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) GetRootRecoveryPlanJobReference() RecoveryPlanJobReference`

GetRootRecoveryPlanJobReference returns the RootRecoveryPlanJobReference field if non-nil, zero value otherwise.

### GetRootRecoveryPlanJobReferenceOk

`func (o *RecoveryPlanJobDefStatus) GetRootRecoveryPlanJobReferenceOk() (*RecoveryPlanJobReference, bool)`

GetRootRecoveryPlanJobReferenceOk returns a tuple with the RootRecoveryPlanJobReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) SetRootRecoveryPlanJobReference(v RecoveryPlanJobReference)`

SetRootRecoveryPlanJobReference sets RootRecoveryPlanJobReference field to given value.

### HasRootRecoveryPlanJobReference

`func (o *RecoveryPlanJobDefStatus) HasRootRecoveryPlanJobReference() bool`

HasRootRecoveryPlanJobReference returns a boolean if a field has been set.

### GetWitnessAddress

`func (o *RecoveryPlanJobDefStatus) GetWitnessAddress() string`

GetWitnessAddress returns the WitnessAddress field if non-nil, zero value otherwise.

### GetWitnessAddressOk

`func (o *RecoveryPlanJobDefStatus) GetWitnessAddressOk() (*string, bool)`

GetWitnessAddressOk returns a tuple with the WitnessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitnessAddress

`func (o *RecoveryPlanJobDefStatus) SetWitnessAddress(v string)`

SetWitnessAddress sets WitnessAddress field to given value.

### HasWitnessAddress

`func (o *RecoveryPlanJobDefStatus) HasWitnessAddress() bool`

HasWitnessAddress returns a boolean if a field has been set.

### GetResources

`func (o *RecoveryPlanJobDefStatus) GetResources() RecoveryPlanJobResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RecoveryPlanJobDefStatus) GetResourcesOk() (*RecoveryPlanJobResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RecoveryPlanJobDefStatus) SetResources(v RecoveryPlanJobResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


