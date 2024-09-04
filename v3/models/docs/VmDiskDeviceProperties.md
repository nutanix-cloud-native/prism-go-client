# VmDiskDeviceProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** |  | [optional] [default to "DISK"]
**DiskAddress** | Pointer to [**DiskAddress**](DiskAddress.md) |  | [optional] 

## Methods

### NewVmDiskDeviceProperties

`func NewVmDiskDeviceProperties() *VmDiskDeviceProperties`

NewVmDiskDeviceProperties instantiates a new VmDiskDeviceProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiskDevicePropertiesWithDefaults

`func NewVmDiskDevicePropertiesWithDefaults() *VmDiskDeviceProperties`

NewVmDiskDevicePropertiesWithDefaults instantiates a new VmDiskDeviceProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *VmDiskDeviceProperties) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *VmDiskDeviceProperties) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *VmDiskDeviceProperties) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *VmDiskDeviceProperties) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDiskAddress

`func (o *VmDiskDeviceProperties) GetDiskAddress() DiskAddress`

GetDiskAddress returns the DiskAddress field if non-nil, zero value otherwise.

### GetDiskAddressOk

`func (o *VmDiskDeviceProperties) GetDiskAddressOk() (*DiskAddress, bool)`

GetDiskAddressOk returns a tuple with the DiskAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAddress

`func (o *VmDiskDeviceProperties) SetDiskAddress(v DiskAddress)`

SetDiskAddress sets DiskAddress field to given value.

### HasDiskAddress

`func (o *VmDiskDeviceProperties) HasDiskAddress() bool`

HasDiskAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


