# DataStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressionSavingPct** | Pointer to **float32** |  | [optional] 
**CpuOvercommitRatio** | Pointer to **float32** |  | [optional] 
**CpuReservationPct** | Pointer to **float32** |  | [optional] 
**InlineDedupSavingPct** | Pointer to **float32** |  | [optional] 
**DedupSavingPct** | Pointer to **float32** |  | [optional] 
**OverallSavingPct** | Pointer to **float32** |  | [optional] 
**ErasureCodingSavingPct** | Pointer to **float32** |  | [optional] 
**RamOvercommitRatio** | Pointer to **float32** |  | [optional] 
**Rf** | Pointer to **int32** |  | [optional] 
**NPlus** | Pointer to **int32** |  | [optional] 
**RamReservationPct** | Pointer to **float32** |  | [optional] 
**StorageReservationPct** | Pointer to **float32** |  | [optional] 

## Methods

### NewDataStoreConfig

`func NewDataStoreConfig() *DataStoreConfig`

NewDataStoreConfig instantiates a new DataStoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreConfigWithDefaults

`func NewDataStoreConfigWithDefaults() *DataStoreConfig`

NewDataStoreConfigWithDefaults instantiates a new DataStoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressionSavingPct

`func (o *DataStoreConfig) GetCompressionSavingPct() float32`

GetCompressionSavingPct returns the CompressionSavingPct field if non-nil, zero value otherwise.

### GetCompressionSavingPctOk

`func (o *DataStoreConfig) GetCompressionSavingPctOk() (*float32, bool)`

GetCompressionSavingPctOk returns a tuple with the CompressionSavingPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionSavingPct

`func (o *DataStoreConfig) SetCompressionSavingPct(v float32)`

SetCompressionSavingPct sets CompressionSavingPct field to given value.

### HasCompressionSavingPct

`func (o *DataStoreConfig) HasCompressionSavingPct() bool`

HasCompressionSavingPct returns a boolean if a field has been set.

### GetCpuOvercommitRatio

`func (o *DataStoreConfig) GetCpuOvercommitRatio() float32`

GetCpuOvercommitRatio returns the CpuOvercommitRatio field if non-nil, zero value otherwise.

### GetCpuOvercommitRatioOk

`func (o *DataStoreConfig) GetCpuOvercommitRatioOk() (*float32, bool)`

GetCpuOvercommitRatioOk returns a tuple with the CpuOvercommitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOvercommitRatio

`func (o *DataStoreConfig) SetCpuOvercommitRatio(v float32)`

SetCpuOvercommitRatio sets CpuOvercommitRatio field to given value.

### HasCpuOvercommitRatio

`func (o *DataStoreConfig) HasCpuOvercommitRatio() bool`

HasCpuOvercommitRatio returns a boolean if a field has been set.

### GetCpuReservationPct

`func (o *DataStoreConfig) GetCpuReservationPct() float32`

GetCpuReservationPct returns the CpuReservationPct field if non-nil, zero value otherwise.

### GetCpuReservationPctOk

`func (o *DataStoreConfig) GetCpuReservationPctOk() (*float32, bool)`

GetCpuReservationPctOk returns a tuple with the CpuReservationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservationPct

`func (o *DataStoreConfig) SetCpuReservationPct(v float32)`

SetCpuReservationPct sets CpuReservationPct field to given value.

### HasCpuReservationPct

`func (o *DataStoreConfig) HasCpuReservationPct() bool`

HasCpuReservationPct returns a boolean if a field has been set.

### GetInlineDedupSavingPct

`func (o *DataStoreConfig) GetInlineDedupSavingPct() float32`

GetInlineDedupSavingPct returns the InlineDedupSavingPct field if non-nil, zero value otherwise.

### GetInlineDedupSavingPctOk

`func (o *DataStoreConfig) GetInlineDedupSavingPctOk() (*float32, bool)`

GetInlineDedupSavingPctOk returns a tuple with the InlineDedupSavingPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineDedupSavingPct

`func (o *DataStoreConfig) SetInlineDedupSavingPct(v float32)`

SetInlineDedupSavingPct sets InlineDedupSavingPct field to given value.

### HasInlineDedupSavingPct

`func (o *DataStoreConfig) HasInlineDedupSavingPct() bool`

HasInlineDedupSavingPct returns a boolean if a field has been set.

### GetDedupSavingPct

`func (o *DataStoreConfig) GetDedupSavingPct() float32`

GetDedupSavingPct returns the DedupSavingPct field if non-nil, zero value otherwise.

### GetDedupSavingPctOk

`func (o *DataStoreConfig) GetDedupSavingPctOk() (*float32, bool)`

GetDedupSavingPctOk returns a tuple with the DedupSavingPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupSavingPct

`func (o *DataStoreConfig) SetDedupSavingPct(v float32)`

SetDedupSavingPct sets DedupSavingPct field to given value.

### HasDedupSavingPct

`func (o *DataStoreConfig) HasDedupSavingPct() bool`

HasDedupSavingPct returns a boolean if a field has been set.

### GetOverallSavingPct

`func (o *DataStoreConfig) GetOverallSavingPct() float32`

GetOverallSavingPct returns the OverallSavingPct field if non-nil, zero value otherwise.

### GetOverallSavingPctOk

`func (o *DataStoreConfig) GetOverallSavingPctOk() (*float32, bool)`

GetOverallSavingPctOk returns a tuple with the OverallSavingPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallSavingPct

`func (o *DataStoreConfig) SetOverallSavingPct(v float32)`

SetOverallSavingPct sets OverallSavingPct field to given value.

### HasOverallSavingPct

`func (o *DataStoreConfig) HasOverallSavingPct() bool`

HasOverallSavingPct returns a boolean if a field has been set.

### GetErasureCodingSavingPct

`func (o *DataStoreConfig) GetErasureCodingSavingPct() float32`

GetErasureCodingSavingPct returns the ErasureCodingSavingPct field if non-nil, zero value otherwise.

### GetErasureCodingSavingPctOk

`func (o *DataStoreConfig) GetErasureCodingSavingPctOk() (*float32, bool)`

GetErasureCodingSavingPctOk returns a tuple with the ErasureCodingSavingPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodingSavingPct

`func (o *DataStoreConfig) SetErasureCodingSavingPct(v float32)`

SetErasureCodingSavingPct sets ErasureCodingSavingPct field to given value.

### HasErasureCodingSavingPct

`func (o *DataStoreConfig) HasErasureCodingSavingPct() bool`

HasErasureCodingSavingPct returns a boolean if a field has been set.

### GetRamOvercommitRatio

`func (o *DataStoreConfig) GetRamOvercommitRatio() float32`

GetRamOvercommitRatio returns the RamOvercommitRatio field if non-nil, zero value otherwise.

### GetRamOvercommitRatioOk

`func (o *DataStoreConfig) GetRamOvercommitRatioOk() (*float32, bool)`

GetRamOvercommitRatioOk returns a tuple with the RamOvercommitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamOvercommitRatio

`func (o *DataStoreConfig) SetRamOvercommitRatio(v float32)`

SetRamOvercommitRatio sets RamOvercommitRatio field to given value.

### HasRamOvercommitRatio

`func (o *DataStoreConfig) HasRamOvercommitRatio() bool`

HasRamOvercommitRatio returns a boolean if a field has been set.

### GetRf

`func (o *DataStoreConfig) GetRf() int32`

GetRf returns the Rf field if non-nil, zero value otherwise.

### GetRfOk

`func (o *DataStoreConfig) GetRfOk() (*int32, bool)`

GetRfOk returns a tuple with the Rf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRf

`func (o *DataStoreConfig) SetRf(v int32)`

SetRf sets Rf field to given value.

### HasRf

`func (o *DataStoreConfig) HasRf() bool`

HasRf returns a boolean if a field has been set.

### GetNPlus

`func (o *DataStoreConfig) GetNPlus() int32`

GetNPlus returns the NPlus field if non-nil, zero value otherwise.

### GetNPlusOk

`func (o *DataStoreConfig) GetNPlusOk() (*int32, bool)`

GetNPlusOk returns a tuple with the NPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPlus

`func (o *DataStoreConfig) SetNPlus(v int32)`

SetNPlus sets NPlus field to given value.

### HasNPlus

`func (o *DataStoreConfig) HasNPlus() bool`

HasNPlus returns a boolean if a field has been set.

### GetRamReservationPct

`func (o *DataStoreConfig) GetRamReservationPct() float32`

GetRamReservationPct returns the RamReservationPct field if non-nil, zero value otherwise.

### GetRamReservationPctOk

`func (o *DataStoreConfig) GetRamReservationPctOk() (*float32, bool)`

GetRamReservationPctOk returns a tuple with the RamReservationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamReservationPct

`func (o *DataStoreConfig) SetRamReservationPct(v float32)`

SetRamReservationPct sets RamReservationPct field to given value.

### HasRamReservationPct

`func (o *DataStoreConfig) HasRamReservationPct() bool`

HasRamReservationPct returns a boolean if a field has been set.

### GetStorageReservationPct

`func (o *DataStoreConfig) GetStorageReservationPct() float32`

GetStorageReservationPct returns the StorageReservationPct field if non-nil, zero value otherwise.

### GetStorageReservationPctOk

`func (o *DataStoreConfig) GetStorageReservationPctOk() (*float32, bool)`

GetStorageReservationPctOk returns a tuple with the StorageReservationPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageReservationPct

`func (o *DataStoreConfig) SetStorageReservationPct(v float32)`

SetStorageReservationPct sets StorageReservationPct field to given value.

### HasStorageReservationPct

`func (o *DataStoreConfig) HasStorageReservationPct() bool`

HasStorageReservationPct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


