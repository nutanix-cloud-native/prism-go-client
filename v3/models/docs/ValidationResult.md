# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPassed** | Pointer to **bool** | Whether the validation passed. | [optional] 
**Reason** | Pointer to **string** | Reason of failed validation. Will only be populated when validation fails.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult() *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPassed

`func (o *ValidationResult) GetHasPassed() bool`

GetHasPassed returns the HasPassed field if non-nil, zero value otherwise.

### GetHasPassedOk

`func (o *ValidationResult) GetHasPassedOk() (*bool, bool)`

GetHasPassedOk returns a tuple with the HasPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassed

`func (o *ValidationResult) SetHasPassed(v bool)`

SetHasPassed sets HasPassed field to given value.

### HasHasPassed

`func (o *ValidationResult) HasHasPassed() bool`

HasHasPassed returns a boolean if a field has been set.

### GetReason

`func (o *ValidationResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ValidationResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ValidationResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ValidationResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetName

`func (o *ValidationResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidationResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidationResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValidationResult) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


