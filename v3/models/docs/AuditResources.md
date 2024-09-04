# AuditResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditMessage** | **string** | Audit message. | 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**OperationParameterList** | Pointer to [**[]AuditParameters**](AuditParameters.md) | Parameters associated with the operation captured in the audit. | [optional] 
**SourceEntityReference** | Pointer to [**EntityInfo**](EntityInfo.md) |  | [optional] 
**OperationStartTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when operation was started.  | [optional] 
**InitiatedUser** | Pointer to [**AuditUser**](AuditUser.md) |  | [optional] 
**OperationCompleteTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when operation was completed.  | [optional] 
**OperationType** | Pointer to **string** | The operation type the audit captures. | [optional] 
**AffectedEntityReferenceList** | Pointer to [**[]EntityInfo**](EntityInfo.md) | A list of entities causing and/or related to the audit.  | [optional] 

## Methods

### NewAuditResources

`func NewAuditResources(auditMessage string, ) *AuditResources`

NewAuditResources instantiates a new AuditResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditResourcesWithDefaults

`func NewAuditResourcesWithDefaults() *AuditResources`

NewAuditResourcesWithDefaults instantiates a new AuditResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditMessage

`func (o *AuditResources) GetAuditMessage() string`

GetAuditMessage returns the AuditMessage field if non-nil, zero value otherwise.

### GetAuditMessageOk

`func (o *AuditResources) GetAuditMessageOk() (*string, bool)`

GetAuditMessageOk returns a tuple with the AuditMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditMessage

`func (o *AuditResources) SetAuditMessage(v string)`

SetAuditMessage sets AuditMessage field to given value.


### GetClusterReference

`func (o *AuditResources) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *AuditResources) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *AuditResources) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *AuditResources) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetOperationParameterList

`func (o *AuditResources) GetOperationParameterList() []AuditParameters`

GetOperationParameterList returns the OperationParameterList field if non-nil, zero value otherwise.

### GetOperationParameterListOk

`func (o *AuditResources) GetOperationParameterListOk() (*[]AuditParameters, bool)`

GetOperationParameterListOk returns a tuple with the OperationParameterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationParameterList

`func (o *AuditResources) SetOperationParameterList(v []AuditParameters)`

SetOperationParameterList sets OperationParameterList field to given value.

### HasOperationParameterList

`func (o *AuditResources) HasOperationParameterList() bool`

HasOperationParameterList returns a boolean if a field has been set.

### GetSourceEntityReference

`func (o *AuditResources) GetSourceEntityReference() EntityInfo`

GetSourceEntityReference returns the SourceEntityReference field if non-nil, zero value otherwise.

### GetSourceEntityReferenceOk

`func (o *AuditResources) GetSourceEntityReferenceOk() (*EntityInfo, bool)`

GetSourceEntityReferenceOk returns a tuple with the SourceEntityReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityReference

`func (o *AuditResources) SetSourceEntityReference(v EntityInfo)`

SetSourceEntityReference sets SourceEntityReference field to given value.

### HasSourceEntityReference

`func (o *AuditResources) HasSourceEntityReference() bool`

HasSourceEntityReference returns a boolean if a field has been set.

### GetOperationStartTime

`func (o *AuditResources) GetOperationStartTime() time.Time`

GetOperationStartTime returns the OperationStartTime field if non-nil, zero value otherwise.

### GetOperationStartTimeOk

`func (o *AuditResources) GetOperationStartTimeOk() (*time.Time, bool)`

GetOperationStartTimeOk returns a tuple with the OperationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStartTime

`func (o *AuditResources) SetOperationStartTime(v time.Time)`

SetOperationStartTime sets OperationStartTime field to given value.

### HasOperationStartTime

`func (o *AuditResources) HasOperationStartTime() bool`

HasOperationStartTime returns a boolean if a field has been set.

### GetInitiatedUser

`func (o *AuditResources) GetInitiatedUser() AuditUser`

GetInitiatedUser returns the InitiatedUser field if non-nil, zero value otherwise.

### GetInitiatedUserOk

`func (o *AuditResources) GetInitiatedUserOk() (*AuditUser, bool)`

GetInitiatedUserOk returns a tuple with the InitiatedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedUser

`func (o *AuditResources) SetInitiatedUser(v AuditUser)`

SetInitiatedUser sets InitiatedUser field to given value.

### HasInitiatedUser

`func (o *AuditResources) HasInitiatedUser() bool`

HasInitiatedUser returns a boolean if a field has been set.

### GetOperationCompleteTime

`func (o *AuditResources) GetOperationCompleteTime() time.Time`

GetOperationCompleteTime returns the OperationCompleteTime field if non-nil, zero value otherwise.

### GetOperationCompleteTimeOk

`func (o *AuditResources) GetOperationCompleteTimeOk() (*time.Time, bool)`

GetOperationCompleteTimeOk returns a tuple with the OperationCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCompleteTime

`func (o *AuditResources) SetOperationCompleteTime(v time.Time)`

SetOperationCompleteTime sets OperationCompleteTime field to given value.

### HasOperationCompleteTime

`func (o *AuditResources) HasOperationCompleteTime() bool`

HasOperationCompleteTime returns a boolean if a field has been set.

### GetOperationType

`func (o *AuditResources) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AuditResources) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AuditResources) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *AuditResources) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetAffectedEntityReferenceList

`func (o *AuditResources) GetAffectedEntityReferenceList() []EntityInfo`

GetAffectedEntityReferenceList returns the AffectedEntityReferenceList field if non-nil, zero value otherwise.

### GetAffectedEntityReferenceListOk

`func (o *AuditResources) GetAffectedEntityReferenceListOk() (*[]EntityInfo, bool)`

GetAffectedEntityReferenceListOk returns a tuple with the AffectedEntityReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntityReferenceList

`func (o *AuditResources) SetAffectedEntityReferenceList(v []EntityInfo)`

SetAffectedEntityReferenceList sets AffectedEntityReferenceList field to given value.

### HasAffectedEntityReferenceList

`func (o *AuditResources) HasAffectedEntityReferenceList() bool`

HasAffectedEntityReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


