# VmBootConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootDeviceOrderList** | Pointer to **[]string** | Indicates the order of device types in which VM should try to boot from. If boot device order is not provided the system will decide appropriate boot device order.  | [optional] 
**BootDevice** | Pointer to [**VmBootDevice**](VmBootDevice.md) |  | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**BootType** | Pointer to **string** | Indicates whether the VM should use Secure boot, UEFI boot or Legacy boot.If UEFI or Secure boot is enabled then other legacy boot options (like boot_device and boot_device_order_list) are ignored. Secure boot depends on UEFI boot, i.e. enabling Secure boot means that UEFI boot is also enabled.  | [optional] 

## Methods

### NewVmBootConfig

`func NewVmBootConfig() *VmBootConfig`

NewVmBootConfig instantiates a new VmBootConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmBootConfigWithDefaults

`func NewVmBootConfigWithDefaults() *VmBootConfig`

NewVmBootConfigWithDefaults instantiates a new VmBootConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootDeviceOrderList

`func (o *VmBootConfig) GetBootDeviceOrderList() []string`

GetBootDeviceOrderList returns the BootDeviceOrderList field if non-nil, zero value otherwise.

### GetBootDeviceOrderListOk

`func (o *VmBootConfig) GetBootDeviceOrderListOk() (*[]string, bool)`

GetBootDeviceOrderListOk returns a tuple with the BootDeviceOrderList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDeviceOrderList

`func (o *VmBootConfig) SetBootDeviceOrderList(v []string)`

SetBootDeviceOrderList sets BootDeviceOrderList field to given value.

### HasBootDeviceOrderList

`func (o *VmBootConfig) HasBootDeviceOrderList() bool`

HasBootDeviceOrderList returns a boolean if a field has been set.

### GetBootDevice

`func (o *VmBootConfig) GetBootDevice() VmBootDevice`

GetBootDevice returns the BootDevice field if non-nil, zero value otherwise.

### GetBootDeviceOk

`func (o *VmBootConfig) GetBootDeviceOk() (*VmBootDevice, bool)`

GetBootDeviceOk returns a tuple with the BootDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDevice

`func (o *VmBootConfig) SetBootDevice(v VmBootDevice)`

SetBootDevice sets BootDevice field to given value.

### HasBootDevice

`func (o *VmBootConfig) HasBootDevice() bool`

HasBootDevice returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VmBootConfig) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VmBootConfig) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VmBootConfig) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VmBootConfig) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetBootType

`func (o *VmBootConfig) GetBootType() string`

GetBootType returns the BootType field if non-nil, zero value otherwise.

### GetBootTypeOk

`func (o *VmBootConfig) GetBootTypeOk() (*string, bool)`

GetBootTypeOk returns a tuple with the BootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootType

`func (o *VmBootConfig) SetBootType(v string)`

SetBootType sets BootType field to given value.

### HasBootType

`func (o *VmBootConfig) HasBootType() bool`

HasBootType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


