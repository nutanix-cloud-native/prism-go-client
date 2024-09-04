# StoragePolicyResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qos** | Pointer to [**StorageQoSParametersForTheEntities**](StorageQoSParametersForTheEntities.md) |  | [optional] 
**Compression** | Pointer to [**CompressionParametersForTheEntities**](CompressionParametersForTheEntities.md) |  | [optional] 
**Encryption** | Pointer to [**EncryptionParametersForTheEntities**](EncryptionParametersForTheEntities.md) |  | [optional] 
**FaultTolerance** | Pointer to [**FaultToleranceParametersForTheEntities**](FaultToleranceParametersForTheEntities.md) |  | [optional] 
**FilterList** | Pointer to [**[]Filter**](Filter.md) | Regex for entities on which the policy has been applied | [optional] 
**Type** | Pointer to **string** | Creator of the storage policy | [optional] [default to "USER"]

## Methods

### NewStoragePolicyResourcesDefStatus

`func NewStoragePolicyResourcesDefStatus() *StoragePolicyResourcesDefStatus`

NewStoragePolicyResourcesDefStatus instantiates a new StoragePolicyResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyResourcesDefStatusWithDefaults

`func NewStoragePolicyResourcesDefStatusWithDefaults() *StoragePolicyResourcesDefStatus`

NewStoragePolicyResourcesDefStatusWithDefaults instantiates a new StoragePolicyResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQos

`func (o *StoragePolicyResourcesDefStatus) GetQos() StorageQoSParametersForTheEntities`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *StoragePolicyResourcesDefStatus) GetQosOk() (*StorageQoSParametersForTheEntities, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *StoragePolicyResourcesDefStatus) SetQos(v StorageQoSParametersForTheEntities)`

SetQos sets Qos field to given value.

### HasQos

`func (o *StoragePolicyResourcesDefStatus) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetCompression

`func (o *StoragePolicyResourcesDefStatus) GetCompression() CompressionParametersForTheEntities`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *StoragePolicyResourcesDefStatus) GetCompressionOk() (*CompressionParametersForTheEntities, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *StoragePolicyResourcesDefStatus) SetCompression(v CompressionParametersForTheEntities)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *StoragePolicyResourcesDefStatus) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetEncryption

`func (o *StoragePolicyResourcesDefStatus) GetEncryption() EncryptionParametersForTheEntities`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *StoragePolicyResourcesDefStatus) GetEncryptionOk() (*EncryptionParametersForTheEntities, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *StoragePolicyResourcesDefStatus) SetEncryption(v EncryptionParametersForTheEntities)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *StoragePolicyResourcesDefStatus) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFaultTolerance

`func (o *StoragePolicyResourcesDefStatus) GetFaultTolerance() FaultToleranceParametersForTheEntities`

GetFaultTolerance returns the FaultTolerance field if non-nil, zero value otherwise.

### GetFaultToleranceOk

`func (o *StoragePolicyResourcesDefStatus) GetFaultToleranceOk() (*FaultToleranceParametersForTheEntities, bool)`

GetFaultToleranceOk returns a tuple with the FaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultTolerance

`func (o *StoragePolicyResourcesDefStatus) SetFaultTolerance(v FaultToleranceParametersForTheEntities)`

SetFaultTolerance sets FaultTolerance field to given value.

### HasFaultTolerance

`func (o *StoragePolicyResourcesDefStatus) HasFaultTolerance() bool`

HasFaultTolerance returns a boolean if a field has been set.

### GetFilterList

`func (o *StoragePolicyResourcesDefStatus) GetFilterList() []Filter`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *StoragePolicyResourcesDefStatus) GetFilterListOk() (*[]Filter, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *StoragePolicyResourcesDefStatus) SetFilterList(v []Filter)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *StoragePolicyResourcesDefStatus) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.

### GetType

`func (o *StoragePolicyResourcesDefStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoragePolicyResourcesDefStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoragePolicyResourcesDefStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoragePolicyResourcesDefStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


