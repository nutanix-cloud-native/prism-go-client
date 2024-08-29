# AwsInstanceDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGb** | Pointer to **int64** | Size of volumes in GB. | [optional] 
**VolumeType** | Pointer to **string** | Type of the Volume. | [optional] 
**DeviceName** | Pointer to **string** | Device name e.g. /dev/sdb | [optional] 
**Iops** | Pointer to **int64** | The requested number of I/O operations per second that the volume can support for provisioned IOPS (SSD) volumes (io1). Minimum &#x3D; 100 IOPS and Maximum &#x3D; 20000 IOPS.  | [optional] 
**SnapshotId** | Pointer to **string** | AWS snapshot ID. | [optional] 
**DeleteOnTermination** | Pointer to **bool** | Status of delete on termination. | [optional] 

## Methods

### NewAwsInstanceDisk

`func NewAwsInstanceDisk() *AwsInstanceDisk`

NewAwsInstanceDisk instantiates a new AwsInstanceDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsInstanceDiskWithDefaults

`func NewAwsInstanceDiskWithDefaults() *AwsInstanceDisk`

NewAwsInstanceDiskWithDefaults instantiates a new AwsInstanceDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGb

`func (o *AwsInstanceDisk) GetSizeGb() int64`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *AwsInstanceDisk) GetSizeGbOk() (*int64, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *AwsInstanceDisk) SetSizeGb(v int64)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *AwsInstanceDisk) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetVolumeType

`func (o *AwsInstanceDisk) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AwsInstanceDisk) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AwsInstanceDisk) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AwsInstanceDisk) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *AwsInstanceDisk) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AwsInstanceDisk) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AwsInstanceDisk) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AwsInstanceDisk) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIops

`func (o *AwsInstanceDisk) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *AwsInstanceDisk) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *AwsInstanceDisk) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *AwsInstanceDisk) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetSnapshotId

`func (o *AwsInstanceDisk) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *AwsInstanceDisk) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *AwsInstanceDisk) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *AwsInstanceDisk) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetDeleteOnTermination

`func (o *AwsInstanceDisk) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *AwsInstanceDisk) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *AwsInstanceDisk) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *AwsInstanceDisk) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


