# RecoveryPlanValidationMessageCauseAndResolutionMessageListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolutionList** | Pointer to **[]string** | Steps to resolve the warning or error. | [optional] 
**Cause** | **string** | Cause of the validation warning or error. | 

## Methods

### NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInner

`func NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInner(cause string, ) *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner`

NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInner instantiates a new RecoveryPlanValidationMessageCauseAndResolutionMessageListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInnerWithDefaults

`func NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInnerWithDefaults() *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner`

NewRecoveryPlanValidationMessageCauseAndResolutionMessageListInnerWithDefaults instantiates a new RecoveryPlanValidationMessageCauseAndResolutionMessageListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutionList

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) GetResolutionList() []string`

GetResolutionList returns the ResolutionList field if non-nil, zero value otherwise.

### GetResolutionListOk

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) GetResolutionListOk() (*[]string, bool)`

GetResolutionListOk returns a tuple with the ResolutionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionList

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) SetResolutionList(v []string)`

SetResolutionList sets ResolutionList field to given value.

### HasResolutionList

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) HasResolutionList() bool`

HasResolutionList returns a boolean if a field has been set.

### GetCause

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *RecoveryPlanValidationMessageCauseAndResolutionMessageListInner) SetCause(v string)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


