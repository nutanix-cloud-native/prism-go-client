# AwsBlockDeviceMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataDiskList** | Pointer to [**[]AwsInstanceDisk**](AwsInstanceDisk.md) | Data disks of the instance. | [optional] 
**RootDisk** | Pointer to [**AwsInstanceDisk**](AwsInstanceDisk.md) |  | [optional] 

## Methods

### NewAwsBlockDeviceMap

`func NewAwsBlockDeviceMap() *AwsBlockDeviceMap`

NewAwsBlockDeviceMap instantiates a new AwsBlockDeviceMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsBlockDeviceMapWithDefaults

`func NewAwsBlockDeviceMapWithDefaults() *AwsBlockDeviceMap`

NewAwsBlockDeviceMapWithDefaults instantiates a new AwsBlockDeviceMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataDiskList

`func (o *AwsBlockDeviceMap) GetDataDiskList() []AwsInstanceDisk`

GetDataDiskList returns the DataDiskList field if non-nil, zero value otherwise.

### GetDataDiskListOk

`func (o *AwsBlockDeviceMap) GetDataDiskListOk() (*[]AwsInstanceDisk, bool)`

GetDataDiskListOk returns a tuple with the DataDiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskList

`func (o *AwsBlockDeviceMap) SetDataDiskList(v []AwsInstanceDisk)`

SetDataDiskList sets DataDiskList field to given value.

### HasDataDiskList

`func (o *AwsBlockDeviceMap) HasDataDiskList() bool`

HasDataDiskList returns a boolean if a field has been set.

### GetRootDisk

`func (o *AwsBlockDeviceMap) GetRootDisk() AwsInstanceDisk`

GetRootDisk returns the RootDisk field if non-nil, zero value otherwise.

### GetRootDiskOk

`func (o *AwsBlockDeviceMap) GetRootDiskOk() (*AwsInstanceDisk, bool)`

GetRootDiskOk returns a tuple with the RootDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDisk

`func (o *AwsBlockDeviceMap) SetRootDisk(v AwsInstanceDisk)`

SetRootDisk sets RootDisk field to given value.

### HasRootDisk

`func (o *AwsBlockDeviceMap) HasRootDisk() bool`

HasRootDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


