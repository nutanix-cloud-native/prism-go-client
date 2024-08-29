# VmBootDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskAddress** | Pointer to [**DiskAddress**](DiskAddress.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address of nic to boot from. | [optional] 

## Methods

### NewVmBootDevice

`func NewVmBootDevice() *VmBootDevice`

NewVmBootDevice instantiates a new VmBootDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmBootDeviceWithDefaults

`func NewVmBootDeviceWithDefaults() *VmBootDevice`

NewVmBootDeviceWithDefaults instantiates a new VmBootDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskAddress

`func (o *VmBootDevice) GetDiskAddress() DiskAddress`

GetDiskAddress returns the DiskAddress field if non-nil, zero value otherwise.

### GetDiskAddressOk

`func (o *VmBootDevice) GetDiskAddressOk() (*DiskAddress, bool)`

GetDiskAddressOk returns a tuple with the DiskAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskAddress

`func (o *VmBootDevice) SetDiskAddress(v DiskAddress)`

SetDiskAddress sets DiskAddress field to given value.

### HasDiskAddress

`func (o *VmBootDevice) HasDiskAddress() bool`

HasDiskAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *VmBootDevice) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VmBootDevice) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VmBootDevice) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VmBootDevice) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


