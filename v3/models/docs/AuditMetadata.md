# AuditMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when audit was last updated  | [optional] [readonly] 
**Kind** | Pointer to **string** | The kind name | [optional] [readonly] [default to "audit"]
**UUID** | Pointer to **string** | audit uuid | [optional] 
**CreationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when audit was created  | [optional] [readonly] 
**OwnerReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**Name** | Pointer to **string** | audit name | [optional] [readonly] 

## Methods

### NewAuditMetadata

`func NewAuditMetadata() *AuditMetadata`

NewAuditMetadata instantiates a new AuditMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditMetadataWithDefaults

`func NewAuditMetadataWithDefaults() *AuditMetadata`

NewAuditMetadataWithDefaults instantiates a new AuditMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *AuditMetadata) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *AuditMetadata) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *AuditMetadata) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *AuditMetadata) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetKind

`func (o *AuditMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuditMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuditMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuditMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUUID

`func (o *AuditMetadata) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AuditMetadata) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AuditMetadata) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AuditMetadata) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetCreationTime

`func (o *AuditMetadata) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AuditMetadata) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AuditMetadata) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AuditMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetOwnerReference

`func (o *AuditMetadata) GetOwnerReference() UserReference`

GetOwnerReference returns the OwnerReference field if non-nil, zero value otherwise.

### GetOwnerReferenceOk

`func (o *AuditMetadata) GetOwnerReferenceOk() (*UserReference, bool)`

GetOwnerReferenceOk returns a tuple with the OwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReference

`func (o *AuditMetadata) SetOwnerReference(v UserReference)`

SetOwnerReference sets OwnerReference field to given value.

### HasOwnerReference

`func (o *AuditMetadata) HasOwnerReference() bool`

HasOwnerReference returns a boolean if a field has been set.

### GetName

`func (o *AuditMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


