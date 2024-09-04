# AwsInstanceDiskOutputStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the Volume. | [optional] 
**SizeGb** | Pointer to **int64** | Size of volumes in GB. | [optional] 
**AvailabilityZone** | Pointer to **string** | The zone on which the EBS volume is created | [optional] 
**VolumeType** | Pointer to **string** | Type of the Volume. | [optional] 
**DeviceName** | Pointer to **string** | Device name e.g. /dev/sdb | [optional] 
**Iops** | Pointer to **int64** | The requested number of I/O operations per second that the volume can support for provisioned IOPS (SSD) volumes (io1). Minimum &#x3D; 100 IOPS and Maximum &#x3D; 20000 IOPS.  | [optional] 
**SnapshotId** | Pointer to **string** | AWS snapshot ID. | [optional] 
**Id** | Pointer to **string** | AWS ID of the volume. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**DeleteOnTermination** | Pointer to **bool** | Status of delete on termination. | [optional] 

## Methods

### NewAwsInstanceDiskOutputStatus

`func NewAwsInstanceDiskOutputStatus() *AwsInstanceDiskOutputStatus`

NewAwsInstanceDiskOutputStatus instantiates a new AwsInstanceDiskOutputStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsInstanceDiskOutputStatusWithDefaults

`func NewAwsInstanceDiskOutputStatusWithDefaults() *AwsInstanceDiskOutputStatus`

NewAwsInstanceDiskOutputStatusWithDefaults instantiates a new AwsInstanceDiskOutputStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsInstanceDiskOutputStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsInstanceDiskOutputStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsInstanceDiskOutputStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsInstanceDiskOutputStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSizeGb

`func (o *AwsInstanceDiskOutputStatus) GetSizeGb() int64`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *AwsInstanceDiskOutputStatus) GetSizeGbOk() (*int64, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *AwsInstanceDiskOutputStatus) SetSizeGb(v int64)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *AwsInstanceDiskOutputStatus) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AwsInstanceDiskOutputStatus) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AwsInstanceDiskOutputStatus) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AwsInstanceDiskOutputStatus) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AwsInstanceDiskOutputStatus) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetVolumeType

`func (o *AwsInstanceDiskOutputStatus) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AwsInstanceDiskOutputStatus) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AwsInstanceDiskOutputStatus) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AwsInstanceDiskOutputStatus) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDeviceName

`func (o *AwsInstanceDiskOutputStatus) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AwsInstanceDiskOutputStatus) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AwsInstanceDiskOutputStatus) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AwsInstanceDiskOutputStatus) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIops

`func (o *AwsInstanceDiskOutputStatus) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *AwsInstanceDiskOutputStatus) GetIopsOk() (*int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *AwsInstanceDiskOutputStatus) SetIops(v int64)`

SetIops sets Iops field to given value.

### HasIops

`func (o *AwsInstanceDiskOutputStatus) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetSnapshotId

`func (o *AwsInstanceDiskOutputStatus) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *AwsInstanceDiskOutputStatus) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *AwsInstanceDiskOutputStatus) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *AwsInstanceDiskOutputStatus) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetId

`func (o *AwsInstanceDiskOutputStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsInstanceDiskOutputStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsInstanceDiskOutputStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsInstanceDiskOutputStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTagList

`func (o *AwsInstanceDiskOutputStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsInstanceDiskOutputStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsInstanceDiskOutputStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsInstanceDiskOutputStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetDeleteOnTermination

`func (o *AwsInstanceDiskOutputStatus) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *AwsInstanceDiskOutputStatus) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *AwsInstanceDiskOutputStatus) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *AwsInstanceDiskOutputStatus) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


