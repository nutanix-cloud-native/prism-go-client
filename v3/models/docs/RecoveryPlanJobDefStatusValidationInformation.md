# RecoveryPlanJobDefStatusValidationInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorsList** | Pointer to [**[]RecoveryPlanValidationMessage**](RecoveryPlanValidationMessage.md) | List of errors related to the Recovery Plan. These errors need to be resolved before any action can be executed on the Recovery Plan.  | [optional] 
**WarningsList** | Pointer to [**[]RecoveryPlanValidationMessage**](RecoveryPlanValidationMessage.md) | List of warnings related to the Recovery Plan. These warnings do not prevent actions on the Recovery Plan.  | [optional] 

## Methods

### NewRecoveryPlanJobDefStatusValidationInformation

`func NewRecoveryPlanJobDefStatusValidationInformation() *RecoveryPlanJobDefStatusValidationInformation`

NewRecoveryPlanJobDefStatusValidationInformation instantiates a new RecoveryPlanJobDefStatusValidationInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobDefStatusValidationInformationWithDefaults

`func NewRecoveryPlanJobDefStatusValidationInformationWithDefaults() *RecoveryPlanJobDefStatusValidationInformation`

NewRecoveryPlanJobDefStatusValidationInformationWithDefaults instantiates a new RecoveryPlanJobDefStatusValidationInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) GetErrorsList() []RecoveryPlanValidationMessage`

GetErrorsList returns the ErrorsList field if non-nil, zero value otherwise.

### GetErrorsListOk

`func (o *RecoveryPlanJobDefStatusValidationInformation) GetErrorsListOk() (*[]RecoveryPlanValidationMessage, bool)`

GetErrorsListOk returns a tuple with the ErrorsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) SetErrorsList(v []RecoveryPlanValidationMessage)`

SetErrorsList sets ErrorsList field to given value.

### HasErrorsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) HasErrorsList() bool`

HasErrorsList returns a boolean if a field has been set.

### GetWarningsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) GetWarningsList() []RecoveryPlanValidationMessage`

GetWarningsList returns the WarningsList field if non-nil, zero value otherwise.

### GetWarningsListOk

`func (o *RecoveryPlanJobDefStatusValidationInformation) GetWarningsListOk() (*[]RecoveryPlanValidationMessage, bool)`

GetWarningsListOk returns a tuple with the WarningsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) SetWarningsList(v []RecoveryPlanValidationMessage)`

SetWarningsList sets WarningsList field to given value.

### HasWarningsList

`func (o *RecoveryPlanJobDefStatusValidationInformation) HasWarningsList() bool`

HasWarningsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


