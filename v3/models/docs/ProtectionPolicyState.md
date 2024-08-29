# ProtectionPolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyReference** | [**Reference**](Reference.md) |  | 
**PolicyInfo** | Pointer to [**ProtectionInfo**](ProtectionInfo.md) |  | [optional] 
**ComplianceStatus** | **string** | Compliance state enum. | 
**EnforcementMode** | **string** | Policy enforcement mode informs us about what the policy engine is currently doing to enforce the policy on the entity. Monitoring indicates that the policy engine is simply monitoring the entity&#39;s state. Enforcing means that the policy engine is currently trying to enforce the policy on the entity. Enforcement failed indicates that the policy engine encountered a non-transient error and requires user intervention to fix the problem, error message gives the reason for error in this case.  | 
**ErrorMessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 

## Methods

### NewProtectionPolicyState

`func NewProtectionPolicyState(policyReference Reference, complianceStatus string, enforcementMode string, ) *ProtectionPolicyState`

NewProtectionPolicyState instantiates a new ProtectionPolicyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyStateWithDefaults

`func NewProtectionPolicyStateWithDefaults() *ProtectionPolicyState`

NewProtectionPolicyStateWithDefaults instantiates a new ProtectionPolicyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyReference

`func (o *ProtectionPolicyState) GetPolicyReference() Reference`

GetPolicyReference returns the PolicyReference field if non-nil, zero value otherwise.

### GetPolicyReferenceOk

`func (o *ProtectionPolicyState) GetPolicyReferenceOk() (*Reference, bool)`

GetPolicyReferenceOk returns a tuple with the PolicyReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReference

`func (o *ProtectionPolicyState) SetPolicyReference(v Reference)`

SetPolicyReference sets PolicyReference field to given value.


### GetPolicyInfo

`func (o *ProtectionPolicyState) GetPolicyInfo() ProtectionInfo`

GetPolicyInfo returns the PolicyInfo field if non-nil, zero value otherwise.

### GetPolicyInfoOk

`func (o *ProtectionPolicyState) GetPolicyInfoOk() (*ProtectionInfo, bool)`

GetPolicyInfoOk returns a tuple with the PolicyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInfo

`func (o *ProtectionPolicyState) SetPolicyInfo(v ProtectionInfo)`

SetPolicyInfo sets PolicyInfo field to given value.

### HasPolicyInfo

`func (o *ProtectionPolicyState) HasPolicyInfo() bool`

HasPolicyInfo returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *ProtectionPolicyState) GetComplianceStatus() string`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *ProtectionPolicyState) GetComplianceStatusOk() (*string, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *ProtectionPolicyState) SetComplianceStatus(v string)`

SetComplianceStatus sets ComplianceStatus field to given value.


### GetEnforcementMode

`func (o *ProtectionPolicyState) GetEnforcementMode() string`

GetEnforcementMode returns the EnforcementMode field if non-nil, zero value otherwise.

### GetEnforcementModeOk

`func (o *ProtectionPolicyState) GetEnforcementModeOk() (*string, bool)`

GetEnforcementModeOk returns a tuple with the EnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMode

`func (o *ProtectionPolicyState) SetEnforcementMode(v string)`

SetEnforcementMode sets EnforcementMode field to given value.


### GetErrorMessageList

`func (o *ProtectionPolicyState) GetErrorMessageList() []MessageResource`

GetErrorMessageList returns the ErrorMessageList field if non-nil, zero value otherwise.

### GetErrorMessageListOk

`func (o *ProtectionPolicyState) GetErrorMessageListOk() (*[]MessageResource, bool)`

GetErrorMessageListOk returns a tuple with the ErrorMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessageList

`func (o *ProtectionPolicyState) SetErrorMessageList(v []MessageResource)`

SetErrorMessageList sets ErrorMessageList field to given value.

### HasErrorMessageList

`func (o *ProtectionPolicyState) HasErrorMessageList() bool`

HasErrorMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


