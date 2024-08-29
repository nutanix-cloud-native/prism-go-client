# QueryEntitiesResponseEntityListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorReasonList** | Pointer to [**[]QueryEntitiesResponseEntityListInnerErrorReasonListInner**](QueryEntitiesResponseEntityListInnerErrorReasonListInner.md) | List of errors because of which protection of entity is not possible.  | [optional] 
**EntityReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**ProtectionStatus** | Pointer to **string** | Protection status of the entity. CAN_BE_PROTECTED would mean that there are no errors for this entity and it can be protected. PROTECTED status would mean that the entity is protected by the same protection policy passed in the arg. PARTIALLY_PROTECTED would mean that the entity is protected by the same protection policy passed in the arg but also has some errors or warnings associated with it. CANNOT_BE_PROTECTED would mean that the entities aren&#39;t protected by the protection policy passed in the arg and have errors associated with it. CAN_BE_PARTIALLY_PROTECTED would mean that entity can be protected but has some warnings associated with it.  | [optional] 
**ProtectionRuleReference** | Pointer to [**ProtectionRuleReference**](ProtectionRuleReference.md) |  | [optional] 
**WarningReasonList** | Pointer to [**[]QueryEntitiesResponseEntityListInnerWarningReasonListInner**](QueryEntitiesResponseEntityListInnerWarningReasonListInner.md) | List of warnings associated with the entity that needs to be protected.  | [optional] 

## Methods

### NewQueryEntitiesResponseEntityListInner

`func NewQueryEntitiesResponseEntityListInner() *QueryEntitiesResponseEntityListInner`

NewQueryEntitiesResponseEntityListInner instantiates a new QueryEntitiesResponseEntityListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryEntitiesResponseEntityListInnerWithDefaults

`func NewQueryEntitiesResponseEntityListInnerWithDefaults() *QueryEntitiesResponseEntityListInner`

NewQueryEntitiesResponseEntityListInnerWithDefaults instantiates a new QueryEntitiesResponseEntityListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorReasonList

`func (o *QueryEntitiesResponseEntityListInner) GetErrorReasonList() []QueryEntitiesResponseEntityListInnerErrorReasonListInner`

GetErrorReasonList returns the ErrorReasonList field if non-nil, zero value otherwise.

### GetErrorReasonListOk

`func (o *QueryEntitiesResponseEntityListInner) GetErrorReasonListOk() (*[]QueryEntitiesResponseEntityListInnerErrorReasonListInner, bool)`

GetErrorReasonListOk returns a tuple with the ErrorReasonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReasonList

`func (o *QueryEntitiesResponseEntityListInner) SetErrorReasonList(v []QueryEntitiesResponseEntityListInnerErrorReasonListInner)`

SetErrorReasonList sets ErrorReasonList field to given value.

### HasErrorReasonList

`func (o *QueryEntitiesResponseEntityListInner) HasErrorReasonList() bool`

HasErrorReasonList returns a boolean if a field has been set.

### GetEntityReference

`func (o *QueryEntitiesResponseEntityListInner) GetEntityReference() Reference`

GetEntityReference returns the EntityReference field if non-nil, zero value otherwise.

### GetEntityReferenceOk

`func (o *QueryEntitiesResponseEntityListInner) GetEntityReferenceOk() (*Reference, bool)`

GetEntityReferenceOk returns a tuple with the EntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityReference

`func (o *QueryEntitiesResponseEntityListInner) SetEntityReference(v Reference)`

SetEntityReference sets EntityReference field to given value.

### HasEntityReference

`func (o *QueryEntitiesResponseEntityListInner) HasEntityReference() bool`

HasEntityReference returns a boolean if a field has been set.

### GetProtectionStatus

`func (o *QueryEntitiesResponseEntityListInner) GetProtectionStatus() string`

GetProtectionStatus returns the ProtectionStatus field if non-nil, zero value otherwise.

### GetProtectionStatusOk

`func (o *QueryEntitiesResponseEntityListInner) GetProtectionStatusOk() (*string, bool)`

GetProtectionStatusOk returns a tuple with the ProtectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionStatus

`func (o *QueryEntitiesResponseEntityListInner) SetProtectionStatus(v string)`

SetProtectionStatus sets ProtectionStatus field to given value.

### HasProtectionStatus

`func (o *QueryEntitiesResponseEntityListInner) HasProtectionStatus() bool`

HasProtectionStatus returns a boolean if a field has been set.

### GetProtectionRuleReference

`func (o *QueryEntitiesResponseEntityListInner) GetProtectionRuleReference() ProtectionRuleReference`

GetProtectionRuleReference returns the ProtectionRuleReference field if non-nil, zero value otherwise.

### GetProtectionRuleReferenceOk

`func (o *QueryEntitiesResponseEntityListInner) GetProtectionRuleReferenceOk() (*ProtectionRuleReference, bool)`

GetProtectionRuleReferenceOk returns a tuple with the ProtectionRuleReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionRuleReference

`func (o *QueryEntitiesResponseEntityListInner) SetProtectionRuleReference(v ProtectionRuleReference)`

SetProtectionRuleReference sets ProtectionRuleReference field to given value.

### HasProtectionRuleReference

`func (o *QueryEntitiesResponseEntityListInner) HasProtectionRuleReference() bool`

HasProtectionRuleReference returns a boolean if a field has been set.

### GetWarningReasonList

`func (o *QueryEntitiesResponseEntityListInner) GetWarningReasonList() []QueryEntitiesResponseEntityListInnerWarningReasonListInner`

GetWarningReasonList returns the WarningReasonList field if non-nil, zero value otherwise.

### GetWarningReasonListOk

`func (o *QueryEntitiesResponseEntityListInner) GetWarningReasonListOk() (*[]QueryEntitiesResponseEntityListInnerWarningReasonListInner, bool)`

GetWarningReasonListOk returns a tuple with the WarningReasonList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningReasonList

`func (o *QueryEntitiesResponseEntityListInner) SetWarningReasonList(v []QueryEntitiesResponseEntityListInnerWarningReasonListInner)`

SetWarningReasonList sets WarningReasonList field to given value.

### HasWarningReasonList

`func (o *QueryEntitiesResponseEntityListInner) HasWarningReasonList() bool`

HasWarningReasonList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


