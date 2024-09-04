# GetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineCompression** | Pointer to **bool** | Indicates whether compression is inline or post-process on the entity.  | [optional] 
**ComplianceState** | Pointer to **string** | Compliance state of the entity with the storage policy.  | [optional] [default to "IN_PROGRESS"]
**ThrottledIops** | Pointer to **int32** | Max IOs the entity is allowed to do in a second. | [optional] 
**ActiveStoragePolicyReferenceList** | Pointer to [**[]Reference**](Reference.md) | List of storage policies active on the entity. | [optional] 
**CompressionEnabled** | Pointer to **bool** | Indicates whether compression is enabled or not on the entity.  | [optional] 
**ReplicationFactor** | Pointer to **int64** | The replication factor effective on the entity. | [optional] 
**EncryptionEnabled** | Pointer to **bool** | Indicates whether encryption is enabled or not on the entity.  | [optional] 
**NonCompliantErrorCodes** | Pointer to **[]string** | Indicates the reasons of non-compliance with the applied policy.  | [optional] 

## Methods

### NewGetInfo

`func NewGetInfo() *GetInfo`

NewGetInfo instantiates a new GetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInfoWithDefaults

`func NewGetInfoWithDefaults() *GetInfo`

NewGetInfoWithDefaults instantiates a new GetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineCompression

`func (o *GetInfo) GetInlineCompression() bool`

GetInlineCompression returns the InlineCompression field if non-nil, zero value otherwise.

### GetInlineCompressionOk

`func (o *GetInfo) GetInlineCompressionOk() (*bool, bool)`

GetInlineCompressionOk returns a tuple with the InlineCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineCompression

`func (o *GetInfo) SetInlineCompression(v bool)`

SetInlineCompression sets InlineCompression field to given value.

### HasInlineCompression

`func (o *GetInfo) HasInlineCompression() bool`

HasInlineCompression returns a boolean if a field has been set.

### GetComplianceState

`func (o *GetInfo) GetComplianceState() string`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *GetInfo) GetComplianceStateOk() (*string, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceState

`func (o *GetInfo) SetComplianceState(v string)`

SetComplianceState sets ComplianceState field to given value.

### HasComplianceState

`func (o *GetInfo) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### GetThrottledIops

`func (o *GetInfo) GetThrottledIops() int32`

GetThrottledIops returns the ThrottledIops field if non-nil, zero value otherwise.

### GetThrottledIopsOk

`func (o *GetInfo) GetThrottledIopsOk() (*int32, bool)`

GetThrottledIopsOk returns a tuple with the ThrottledIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledIops

`func (o *GetInfo) SetThrottledIops(v int32)`

SetThrottledIops sets ThrottledIops field to given value.

### HasThrottledIops

`func (o *GetInfo) HasThrottledIops() bool`

HasThrottledIops returns a boolean if a field has been set.

### GetActiveStoragePolicyReferenceList

`func (o *GetInfo) GetActiveStoragePolicyReferenceList() []Reference`

GetActiveStoragePolicyReferenceList returns the ActiveStoragePolicyReferenceList field if non-nil, zero value otherwise.

### GetActiveStoragePolicyReferenceListOk

`func (o *GetInfo) GetActiveStoragePolicyReferenceListOk() (*[]Reference, bool)`

GetActiveStoragePolicyReferenceListOk returns a tuple with the ActiveStoragePolicyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStoragePolicyReferenceList

`func (o *GetInfo) SetActiveStoragePolicyReferenceList(v []Reference)`

SetActiveStoragePolicyReferenceList sets ActiveStoragePolicyReferenceList field to given value.

### HasActiveStoragePolicyReferenceList

`func (o *GetInfo) HasActiveStoragePolicyReferenceList() bool`

HasActiveStoragePolicyReferenceList returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *GetInfo) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *GetInfo) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *GetInfo) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *GetInfo) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *GetInfo) GetReplicationFactor() int64`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *GetInfo) GetReplicationFactorOk() (*int64, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *GetInfo) SetReplicationFactor(v int64)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *GetInfo) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *GetInfo) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *GetInfo) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *GetInfo) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *GetInfo) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetNonCompliantErrorCodes

`func (o *GetInfo) GetNonCompliantErrorCodes() []string`

GetNonCompliantErrorCodes returns the NonCompliantErrorCodes field if non-nil, zero value otherwise.

### GetNonCompliantErrorCodesOk

`func (o *GetInfo) GetNonCompliantErrorCodesOk() (*[]string, bool)`

GetNonCompliantErrorCodesOk returns a tuple with the NonCompliantErrorCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCompliantErrorCodes

`func (o *GetInfo) SetNonCompliantErrorCodes(v []string)`

SetNonCompliantErrorCodes sets NonCompliantErrorCodes field to given value.

### HasNonCompliantErrorCodes

`func (o *GetInfo) HasNonCompliantErrorCodes() bool`

HasNonCompliantErrorCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


