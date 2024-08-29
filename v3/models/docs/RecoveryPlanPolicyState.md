# RecoveryPlanPolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyReference** | [**Reference**](Reference.md) |  | 
**PolicyInfo** | Pointer to [**RecoveryPlanInfo**](RecoveryPlanInfo.md) |  | [optional] 
**ComplianceStatus** | **string** | Compliance state enum. | 
**EnforcementMode** | **string** | Policy enforcement mode informs us about what the policy engine is currently doing to enforce the policy on the entity. Monitoring indicates that the policy engine is simply monitoring the entity&#39;s state. Enforcing means that the policy engine is currently trying to enforce the policy on the entity. Enforcement failed indicates that the policy engine encountered a non-transient error and requires user intervention to fix the problem, error message gives the reason for error in this case.  | 
**ErrorMessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 

## Methods

### NewRecoveryPlanPolicyState

`func NewRecoveryPlanPolicyState(policyReference Reference, complianceStatus string, enforcementMode string, ) *RecoveryPlanPolicyState`

NewRecoveryPlanPolicyState instantiates a new RecoveryPlanPolicyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanPolicyStateWithDefaults

`func NewRecoveryPlanPolicyStateWithDefaults() *RecoveryPlanPolicyState`

NewRecoveryPlanPolicyStateWithDefaults instantiates a new RecoveryPlanPolicyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyReference

`func (o *RecoveryPlanPolicyState) GetPolicyReference() Reference`

GetPolicyReference returns the PolicyReference field if non-nil, zero value otherwise.

### GetPolicyReferenceOk

`func (o *RecoveryPlanPolicyState) GetPolicyReferenceOk() (*Reference, bool)`

GetPolicyReferenceOk returns a tuple with the PolicyReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReference

`func (o *RecoveryPlanPolicyState) SetPolicyReference(v Reference)`

SetPolicyReference sets PolicyReference field to given value.


### GetPolicyInfo

`func (o *RecoveryPlanPolicyState) GetPolicyInfo() RecoveryPlanInfo`

GetPolicyInfo returns the PolicyInfo field if non-nil, zero value otherwise.

### GetPolicyInfoOk

`func (o *RecoveryPlanPolicyState) GetPolicyInfoOk() (*RecoveryPlanInfo, bool)`

GetPolicyInfoOk returns a tuple with the PolicyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInfo

`func (o *RecoveryPlanPolicyState) SetPolicyInfo(v RecoveryPlanInfo)`

SetPolicyInfo sets PolicyInfo field to given value.

### HasPolicyInfo

`func (o *RecoveryPlanPolicyState) HasPolicyInfo() bool`

HasPolicyInfo returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *RecoveryPlanPolicyState) GetComplianceStatus() string`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *RecoveryPlanPolicyState) GetComplianceStatusOk() (*string, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *RecoveryPlanPolicyState) SetComplianceStatus(v string)`

SetComplianceStatus sets ComplianceStatus field to given value.


### GetEnforcementMode

`func (o *RecoveryPlanPolicyState) GetEnforcementMode() string`

GetEnforcementMode returns the EnforcementMode field if non-nil, zero value otherwise.

### GetEnforcementModeOk

`func (o *RecoveryPlanPolicyState) GetEnforcementModeOk() (*string, bool)`

GetEnforcementModeOk returns a tuple with the EnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMode

`func (o *RecoveryPlanPolicyState) SetEnforcementMode(v string)`

SetEnforcementMode sets EnforcementMode field to given value.


### GetErrorMessageList

`func (o *RecoveryPlanPolicyState) GetErrorMessageList() []MessageResource`

GetErrorMessageList returns the ErrorMessageList field if non-nil, zero value otherwise.

### GetErrorMessageListOk

`func (o *RecoveryPlanPolicyState) GetErrorMessageListOk() (*[]MessageResource, bool)`

GetErrorMessageListOk returns a tuple with the ErrorMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessageList

`func (o *RecoveryPlanPolicyState) SetErrorMessageList(v []MessageResource)`

SetErrorMessageList sets ErrorMessageList field to given value.

### HasErrorMessageList

`func (o *RecoveryPlanPolicyState) HasErrorMessageList() bool`

HasErrorMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


