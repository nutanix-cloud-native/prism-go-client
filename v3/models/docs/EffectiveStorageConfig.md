# EffectiveStorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressionDelaySecs** | Pointer to **int32** | Deprecated. Delay in seconds after which the VM disk data will be compressed.  | [optional] 
**CompressionEnabled** | Pointer to **bool** | Deprecated. Indicates whether compression is enabled or not on the VM.  | [optional] 
**ThrottledIops** | Pointer to **int32** | Deprecated. Max IOs the VM is allowed to do in a second. | [optional] 
**ActiveStoragePolicyReferenceList** | Pointer to [**[]Reference**](Reference.md) | List of storage policies active on the VM. | [optional] 
**EncryptionEnabled** | Pointer to **bool** | Deprecated. Indicates whether encryption is enabled or not on the VM.  | [optional] 

## Methods

### NewEffectiveStorageConfig

`func NewEffectiveStorageConfig() *EffectiveStorageConfig`

NewEffectiveStorageConfig instantiates a new EffectiveStorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectiveStorageConfigWithDefaults

`func NewEffectiveStorageConfigWithDefaults() *EffectiveStorageConfig`

NewEffectiveStorageConfigWithDefaults instantiates a new EffectiveStorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressionDelaySecs

`func (o *EffectiveStorageConfig) GetCompressionDelaySecs() int32`

GetCompressionDelaySecs returns the CompressionDelaySecs field if non-nil, zero value otherwise.

### GetCompressionDelaySecsOk

`func (o *EffectiveStorageConfig) GetCompressionDelaySecsOk() (*int32, bool)`

GetCompressionDelaySecsOk returns a tuple with the CompressionDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionDelaySecs

`func (o *EffectiveStorageConfig) SetCompressionDelaySecs(v int32)`

SetCompressionDelaySecs sets CompressionDelaySecs field to given value.

### HasCompressionDelaySecs

`func (o *EffectiveStorageConfig) HasCompressionDelaySecs() bool`

HasCompressionDelaySecs returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *EffectiveStorageConfig) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *EffectiveStorageConfig) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *EffectiveStorageConfig) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *EffectiveStorageConfig) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetThrottledIops

`func (o *EffectiveStorageConfig) GetThrottledIops() int32`

GetThrottledIops returns the ThrottledIops field if non-nil, zero value otherwise.

### GetThrottledIopsOk

`func (o *EffectiveStorageConfig) GetThrottledIopsOk() (*int32, bool)`

GetThrottledIopsOk returns a tuple with the ThrottledIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledIops

`func (o *EffectiveStorageConfig) SetThrottledIops(v int32)`

SetThrottledIops sets ThrottledIops field to given value.

### HasThrottledIops

`func (o *EffectiveStorageConfig) HasThrottledIops() bool`

HasThrottledIops returns a boolean if a field has been set.

### GetActiveStoragePolicyReferenceList

`func (o *EffectiveStorageConfig) GetActiveStoragePolicyReferenceList() []Reference`

GetActiveStoragePolicyReferenceList returns the ActiveStoragePolicyReferenceList field if non-nil, zero value otherwise.

### GetActiveStoragePolicyReferenceListOk

`func (o *EffectiveStorageConfig) GetActiveStoragePolicyReferenceListOk() (*[]Reference, bool)`

GetActiveStoragePolicyReferenceListOk returns a tuple with the ActiveStoragePolicyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStoragePolicyReferenceList

`func (o *EffectiveStorageConfig) SetActiveStoragePolicyReferenceList(v []Reference)`

SetActiveStoragePolicyReferenceList sets ActiveStoragePolicyReferenceList field to given value.

### HasActiveStoragePolicyReferenceList

`func (o *EffectiveStorageConfig) HasActiveStoragePolicyReferenceList() bool`

HasActiveStoragePolicyReferenceList returns a boolean if a field has been set.

### GetEncryptionEnabled

`func (o *EffectiveStorageConfig) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *EffectiveStorageConfig) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *EffectiveStorageConfig) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *EffectiveStorageConfig) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


