# DiskSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSizeBytes** | **int64** | Total size of the disk in bytes. | 
**UUID** | **string** | The device ID which is used to uniquely identify this particular disk.  | 

## Methods

### NewDiskSpec

`func NewDiskSpec(totalSizeBytes int64, uUID string, ) *DiskSpec`

NewDiskSpec instantiates a new DiskSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSpecWithDefaults

`func NewDiskSpecWithDefaults() *DiskSpec`

NewDiskSpecWithDefaults instantiates a new DiskSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSizeBytes

`func (o *DiskSpec) GetTotalSizeBytes() int64`

GetTotalSizeBytes returns the TotalSizeBytes field if non-nil, zero value otherwise.

### GetTotalSizeBytesOk

`func (o *DiskSpec) GetTotalSizeBytesOk() (*int64, bool)`

GetTotalSizeBytesOk returns a tuple with the TotalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeBytes

`func (o *DiskSpec) SetTotalSizeBytes(v int64)`

SetTotalSizeBytes sets TotalSizeBytes field to given value.


### GetUUID

`func (o *DiskSpec) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *DiskSpec) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *DiskSpec) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


