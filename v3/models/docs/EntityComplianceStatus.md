# EntityComplianceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | If the entity in not in compliance, this field shows the reason as to why not.  | [optional] 
**ComplianceState** | Pointer to **string** | Compliance state enum. | [optional] 
**EntityUuid** | Pointer to **string** | The entity UUID. | [optional] 
**PolicyUuid** | Pointer to **string** | The policy UUID. | [optional] 

## Methods

### NewEntityComplianceStatus

`func NewEntityComplianceStatus() *EntityComplianceStatus`

NewEntityComplianceStatus instantiates a new EntityComplianceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityComplianceStatusWithDefaults

`func NewEntityComplianceStatusWithDefaults() *EntityComplianceStatus`

NewEntityComplianceStatusWithDefaults instantiates a new EntityComplianceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *EntityComplianceStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EntityComplianceStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EntityComplianceStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EntityComplianceStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComplianceState

`func (o *EntityComplianceStatus) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *EntityComplianceStatus) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *EntityComplianceStatus) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *EntityComplianceStatus) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetEntityUuid

`func (o *EntityComplianceStatus) GetEntityUuid() string`

GetEntityUuid returns the EntityUuid field if non-nil, zero value otherwise.

### GetEntityUuidOk

`func (o *EntityComplianceStatus) GetEntityUuidOk() (*string, bool)`

GetEntityUuidOk returns a tuple with the EntityUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityUuid

`func (o *EntityComplianceStatus) SetEntityUuid(v string)`

SetEntityUuid sets EntityUuid field to given value.

### HasEntityUuid

`func (o *EntityComplianceStatus) HasEntityUuid() bool`

HasEntityUuid returns a boolean if a field has been set.

### GetPolicyUuid

`func (o *EntityComplianceStatus) GetPolicyUuid() string`

GetPolicyUuid returns the PolicyUuid field if non-nil, zero value otherwise.

### GetPolicyUuidOk

`func (o *EntityComplianceStatus) GetPolicyUuidOk() (*string, bool)`

GetPolicyUuidOk returns a tuple with the PolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUuid

`func (o *EntityComplianceStatus) SetPolicyUuid(v string)`

SetPolicyUuid sets PolicyUuid field to given value.

### HasPolicyUuid

`func (o *EntityComplianceStatus) HasPolicyUuid() bool`

HasPolicyUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


