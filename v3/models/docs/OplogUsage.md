# OplogUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OplogDiskPct** | Pointer to **float32** | Oplog disk size used in percentage. | [optional] 
**OplogDiskSize** | Pointer to **int64** | Size of oplog disk in bytes. | [optional] 

## Methods

### NewOplogUsage

`func NewOplogUsage() *OplogUsage`

NewOplogUsage instantiates a new OplogUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOplogUsageWithDefaults

`func NewOplogUsageWithDefaults() *OplogUsage`

NewOplogUsageWithDefaults instantiates a new OplogUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOplogDiskPct

`func (o *OplogUsage) GetOplogDiskPct() float32`

GetOplogDiskPct returns the OplogDiskPct field if non-nil, zero value otherwise.

### GetOplogDiskPctOk

`func (o *OplogUsage) GetOplogDiskPctOk() (*float32, bool)`

GetOplogDiskPctOk returns a tuple with the OplogDiskPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogDiskPct

`func (o *OplogUsage) SetOplogDiskPct(v float32)`

SetOplogDiskPct sets OplogDiskPct field to given value.

### HasOplogDiskPct

`func (o *OplogUsage) HasOplogDiskPct() bool`

HasOplogDiskPct returns a boolean if a field has been set.

### GetOplogDiskSize

`func (o *OplogUsage) GetOplogDiskSize() int64`

GetOplogDiskSize returns the OplogDiskSize field if non-nil, zero value otherwise.

### GetOplogDiskSizeOk

`func (o *OplogUsage) GetOplogDiskSizeOk() (*int64, bool)`

GetOplogDiskSizeOk returns a tuple with the OplogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogDiskSize

`func (o *OplogUsage) SetOplogDiskSize(v int64)`

SetOplogDiskSize sets OplogDiskSize field to given value.

### HasOplogDiskSize

`func (o *OplogUsage) HasOplogDiskSize() bool`

HasOplogDiskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


