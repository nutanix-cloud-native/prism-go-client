# VolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | **string** | The UUID which is used to uniquely identify this Volume.  | 
**VolumeRecoveryPointReference** | Pointer to [**VolumeRecoveryPointReference**](VolumeRecoveryPointReference.md) |  | [optional] 
**DiskList** | Pointer to [**[]DiskSpec**](DiskSpec.md) | List of associated Volume virtual disks. | [optional] 

## Methods

### NewVolumeSpec

`func NewVolumeSpec(uUID string, ) *VolumeSpec`

NewVolumeSpec instantiates a new VolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSpecWithDefaults

`func NewVolumeSpecWithDefaults() *VolumeSpec`

NewVolumeSpecWithDefaults instantiates a new VolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *VolumeSpec) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VolumeSpec) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VolumeSpec) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetVolumeRecoveryPointReference

`func (o *VolumeSpec) GetVolumeRecoveryPointReference() VolumeRecoveryPointReference`

GetVolumeRecoveryPointReference returns the VolumeRecoveryPointReference field if non-nil, zero value otherwise.

### GetVolumeRecoveryPointReferenceOk

`func (o *VolumeSpec) GetVolumeRecoveryPointReferenceOk() (*VolumeRecoveryPointReference, bool)`

GetVolumeRecoveryPointReferenceOk returns a tuple with the VolumeRecoveryPointReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRecoveryPointReference

`func (o *VolumeSpec) SetVolumeRecoveryPointReference(v VolumeRecoveryPointReference)`

SetVolumeRecoveryPointReference sets VolumeRecoveryPointReference field to given value.

### HasVolumeRecoveryPointReference

`func (o *VolumeSpec) HasVolumeRecoveryPointReference() bool`

HasVolumeRecoveryPointReference returns a boolean if a field has been set.

### GetDiskList

`func (o *VolumeSpec) GetDiskList() []DiskSpec`

GetDiskList returns the DiskList field if non-nil, zero value otherwise.

### GetDiskListOk

`func (o *VolumeSpec) GetDiskListOk() (*[]DiskSpec, bool)`

GetDiskListOk returns a tuple with the DiskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskList

`func (o *VolumeSpec) SetDiskList(v []DiskSpec)`

SetDiskList sets DiskList field to given value.

### HasDiskList

`func (o *VolumeSpec) HasDiskList() bool`

HasDiskList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


