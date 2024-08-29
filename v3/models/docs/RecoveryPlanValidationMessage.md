# RecoveryPlanValidationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Message describing validation error or warning. | 
**ValidationType** | **string** | Type of validation. | 
**AffectedAnyReferenceList** | Pointer to [**[]Reference**](Reference.md) | List of affected entities for this validation message. | [optional] 
**ImpactMessageList** | Pointer to **[]string** | Impact of the validation message on the Recovery Plan actions.  | [optional] 
**CauseAndResolutionMessageList** | Pointer to [**[]RecoveryPlanValidationMessageCauseAndResolutionMessageListInner**](RecoveryPlanValidationMessageCauseAndResolutionMessageListInner.md) | List of causes and resolutions for the validation warning or error.  | [optional] 

## Methods

### NewRecoveryPlanValidationMessage

`func NewRecoveryPlanValidationMessage(message string, validationType string, ) *RecoveryPlanValidationMessage`

NewRecoveryPlanValidationMessage instantiates a new RecoveryPlanValidationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanValidationMessageWithDefaults

`func NewRecoveryPlanValidationMessageWithDefaults() *RecoveryPlanValidationMessage`

NewRecoveryPlanValidationMessageWithDefaults instantiates a new RecoveryPlanValidationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RecoveryPlanValidationMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RecoveryPlanValidationMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RecoveryPlanValidationMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetValidationType

`func (o *RecoveryPlanValidationMessage) GetValidationType() string`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *RecoveryPlanValidationMessage) GetValidationTypeOk() (*string, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *RecoveryPlanValidationMessage) SetValidationType(v string)`

SetValidationType sets ValidationType field to given value.


### GetAffectedAnyReferenceList

`func (o *RecoveryPlanValidationMessage) GetAffectedAnyReferenceList() []Reference`

GetAffectedAnyReferenceList returns the AffectedAnyReferenceList field if non-nil, zero value otherwise.

### GetAffectedAnyReferenceListOk

`func (o *RecoveryPlanValidationMessage) GetAffectedAnyReferenceListOk() (*[]Reference, bool)`

GetAffectedAnyReferenceListOk returns a tuple with the AffectedAnyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedAnyReferenceList

`func (o *RecoveryPlanValidationMessage) SetAffectedAnyReferenceList(v []Reference)`

SetAffectedAnyReferenceList sets AffectedAnyReferenceList field to given value.

### HasAffectedAnyReferenceList

`func (o *RecoveryPlanValidationMessage) HasAffectedAnyReferenceList() bool`

HasAffectedAnyReferenceList returns a boolean if a field has been set.

### GetImpactMessageList

`func (o *RecoveryPlanValidationMessage) GetImpactMessageList() []string`

GetImpactMessageList returns the ImpactMessageList field if non-nil, zero value otherwise.

### GetImpactMessageListOk

`func (o *RecoveryPlanValidationMessage) GetImpactMessageListOk() (*[]string, bool)`

GetImpactMessageListOk returns a tuple with the ImpactMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactMessageList

`func (o *RecoveryPlanValidationMessage) SetImpactMessageList(v []string)`

SetImpactMessageList sets ImpactMessageList field to given value.

### HasImpactMessageList

`func (o *RecoveryPlanValidationMessage) HasImpactMessageList() bool`

HasImpactMessageList returns a boolean if a field has been set.

### GetCauseAndResolutionMessageList

`func (o *RecoveryPlanValidationMessage) GetCauseAndResolutionMessageList() []RecoveryPlanValidationMessageCauseAndResolutionMessageListInner`

GetCauseAndResolutionMessageList returns the CauseAndResolutionMessageList field if non-nil, zero value otherwise.

### GetCauseAndResolutionMessageListOk

`func (o *RecoveryPlanValidationMessage) GetCauseAndResolutionMessageListOk() (*[]RecoveryPlanValidationMessageCauseAndResolutionMessageListInner, bool)`

GetCauseAndResolutionMessageListOk returns a tuple with the CauseAndResolutionMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauseAndResolutionMessageList

`func (o *RecoveryPlanValidationMessage) SetCauseAndResolutionMessageList(v []RecoveryPlanValidationMessageCauseAndResolutionMessageListInner)`

SetCauseAndResolutionMessageList sets CauseAndResolutionMessageList field to given value.

### HasCauseAndResolutionMessageList

`func (o *RecoveryPlanValidationMessage) HasCauseAndResolutionMessageList() bool`

HasCauseAndResolutionMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


