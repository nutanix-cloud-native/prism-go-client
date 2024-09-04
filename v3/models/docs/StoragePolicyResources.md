# StoragePolicyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encryption** | Pointer to [**EncryptionParametersForTheEntities**](EncryptionParametersForTheEntities.md) |  | [optional] 
**FaultTolerance** | Pointer to [**FaultToleranceParametersForTheEntities**](FaultToleranceParametersForTheEntities.md) |  | [optional] 
**Qos** | Pointer to [**StorageQoSParametersForTheEntities**](StorageQoSParametersForTheEntities.md) |  | [optional] 
**Compression** | Pointer to [**CompressionParametersForTheEntities**](CompressionParametersForTheEntities.md) |  | [optional] 
**FilterList** | Pointer to [**[]Filter**](Filter.md) | Regex for entities on which the policy is to be applied | [optional] 

## Methods

### NewStoragePolicyResources

`func NewStoragePolicyResources() *StoragePolicyResources`

NewStoragePolicyResources instantiates a new StoragePolicyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyResourcesWithDefaults

`func NewStoragePolicyResourcesWithDefaults() *StoragePolicyResources`

NewStoragePolicyResourcesWithDefaults instantiates a new StoragePolicyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryption

`func (o *StoragePolicyResources) GetEncryption() EncryptionParametersForTheEntities`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *StoragePolicyResources) GetEncryptionOk() (*EncryptionParametersForTheEntities, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *StoragePolicyResources) SetEncryption(v EncryptionParametersForTheEntities)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *StoragePolicyResources) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFaultTolerance

`func (o *StoragePolicyResources) GetFaultTolerance() FaultToleranceParametersForTheEntities`

GetFaultTolerance returns the FaultTolerance field if non-nil, zero value otherwise.

### GetFaultToleranceOk

`func (o *StoragePolicyResources) GetFaultToleranceOk() (*FaultToleranceParametersForTheEntities, bool)`

GetFaultToleranceOk returns a tuple with the FaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultTolerance

`func (o *StoragePolicyResources) SetFaultTolerance(v FaultToleranceParametersForTheEntities)`

SetFaultTolerance sets FaultTolerance field to given value.

### HasFaultTolerance

`func (o *StoragePolicyResources) HasFaultTolerance() bool`

HasFaultTolerance returns a boolean if a field has been set.

### GetQos

`func (o *StoragePolicyResources) GetQos() StorageQoSParametersForTheEntities`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *StoragePolicyResources) GetQosOk() (*StorageQoSParametersForTheEntities, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *StoragePolicyResources) SetQos(v StorageQoSParametersForTheEntities)`

SetQos sets Qos field to given value.

### HasQos

`func (o *StoragePolicyResources) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetCompression

`func (o *StoragePolicyResources) GetCompression() CompressionParametersForTheEntities`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *StoragePolicyResources) GetCompressionOk() (*CompressionParametersForTheEntities, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *StoragePolicyResources) SetCompression(v CompressionParametersForTheEntities)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *StoragePolicyResources) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetFilterList

`func (o *StoragePolicyResources) GetFilterList() []Filter`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *StoragePolicyResources) GetFilterListOk() (*[]Filter, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *StoragePolicyResources) SetFilterList(v []Filter)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *StoragePolicyResources) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


