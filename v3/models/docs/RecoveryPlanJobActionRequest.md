# RecoveryPlanJobActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldContinueRerunOnValidationFailure** | Pointer to **bool** | Whether to continue rerun execution if warnings are detected during recovery validations.  | [optional] [default to false]
**RerunRecoveryPlanJobUuid** | Pointer to **string** | UUID for referencing the new Recovery Plan Job created for running the failed and incomplete operations. If not specified system generated one will be used. Reference to this will also be populated in entity_reference_list of the task returned in the response.  | [optional] 

## Methods

### NewRecoveryPlanJobActionRequest

`func NewRecoveryPlanJobActionRequest() *RecoveryPlanJobActionRequest`

NewRecoveryPlanJobActionRequest instantiates a new RecoveryPlanJobActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobActionRequestWithDefaults

`func NewRecoveryPlanJobActionRequestWithDefaults() *RecoveryPlanJobActionRequest`

NewRecoveryPlanJobActionRequestWithDefaults instantiates a new RecoveryPlanJobActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldContinueRerunOnValidationFailure

`func (o *RecoveryPlanJobActionRequest) GetShouldContinueRerunOnValidationFailure() bool`

GetShouldContinueRerunOnValidationFailure returns the ShouldContinueRerunOnValidationFailure field if non-nil, zero value otherwise.

### GetShouldContinueRerunOnValidationFailureOk

`func (o *RecoveryPlanJobActionRequest) GetShouldContinueRerunOnValidationFailureOk() (*bool, bool)`

GetShouldContinueRerunOnValidationFailureOk returns a tuple with the ShouldContinueRerunOnValidationFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldContinueRerunOnValidationFailure

`func (o *RecoveryPlanJobActionRequest) SetShouldContinueRerunOnValidationFailure(v bool)`

SetShouldContinueRerunOnValidationFailure sets ShouldContinueRerunOnValidationFailure field to given value.

### HasShouldContinueRerunOnValidationFailure

`func (o *RecoveryPlanJobActionRequest) HasShouldContinueRerunOnValidationFailure() bool`

HasShouldContinueRerunOnValidationFailure returns a boolean if a field has been set.

### GetRerunRecoveryPlanJobUuid

`func (o *RecoveryPlanJobActionRequest) GetRerunRecoveryPlanJobUuid() string`

GetRerunRecoveryPlanJobUuid returns the RerunRecoveryPlanJobUuid field if non-nil, zero value otherwise.

### GetRerunRecoveryPlanJobUuidOk

`func (o *RecoveryPlanJobActionRequest) GetRerunRecoveryPlanJobUuidOk() (*string, bool)`

GetRerunRecoveryPlanJobUuidOk returns a tuple with the RerunRecoveryPlanJobUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunRecoveryPlanJobUuid

`func (o *RecoveryPlanJobActionRequest) SetRerunRecoveryPlanJobUuid(v string)`

SetRerunRecoveryPlanJobUuid sets RerunRecoveryPlanJobUuid field to given value.

### HasRerunRecoveryPlanJobUuid

`func (o *RecoveryPlanJobActionRequest) HasRerunRecoveryPlanJobUuid() bool`

HasRerunRecoveryPlanJobUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


