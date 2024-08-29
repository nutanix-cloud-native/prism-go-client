# AffectedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConflictingProtectionRuleList** | Pointer to [**[]AffectedEntityConflictingProtectionRuleListInner**](AffectedEntityConflictingProtectionRuleListInner.md) | List of reference to conflicting protection rules for an entity. If an entity is filtered by multiple protection rules, Kanon service can not protect that entity. We will flag those entities along with the protection rules. This is set only incase if the entity has conflicts.  | [optional] 
**VmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 

## Methods

### NewAffectedEntity

`func NewAffectedEntity() *AffectedEntity`

NewAffectedEntity instantiates a new AffectedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffectedEntityWithDefaults

`func NewAffectedEntityWithDefaults() *AffectedEntity`

NewAffectedEntityWithDefaults instantiates a new AffectedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflictingProtectionRuleList

`func (o *AffectedEntity) GetConflictingProtectionRuleList() []AffectedEntityConflictingProtectionRuleListInner`

GetConflictingProtectionRuleList returns the ConflictingProtectionRuleList field if non-nil, zero value otherwise.

### GetConflictingProtectionRuleListOk

`func (o *AffectedEntity) GetConflictingProtectionRuleListOk() (*[]AffectedEntityConflictingProtectionRuleListInner, bool)`

GetConflictingProtectionRuleListOk returns a tuple with the ConflictingProtectionRuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingProtectionRuleList

`func (o *AffectedEntity) SetConflictingProtectionRuleList(v []AffectedEntityConflictingProtectionRuleListInner)`

SetConflictingProtectionRuleList sets ConflictingProtectionRuleList field to given value.

### HasConflictingProtectionRuleList

`func (o *AffectedEntity) HasConflictingProtectionRuleList() bool`

HasConflictingProtectionRuleList returns a boolean if a field has been set.

### GetVmReference

`func (o *AffectedEntity) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *AffectedEntity) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *AffectedEntity) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.

### HasVmReference

`func (o *AffectedEntity) HasVmReference() bool`

HasVmReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


