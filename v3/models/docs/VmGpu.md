# VmGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | Pointer to **string** | The vendor of the GPU. | [optional] 
**Mode** | Pointer to **string** | The mode of this GPU. | [optional] 
**DeviceId** | Pointer to **int32** | The device ID of the GPU. | [optional] 

## Methods

### NewVmGpu

`func NewVmGpu() *VmGpu`

NewVmGpu instantiates a new VmGpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmGpuWithDefaults

`func NewVmGpuWithDefaults() *VmGpu`

NewVmGpuWithDefaults instantiates a new VmGpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *VmGpu) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VmGpu) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VmGpu) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *VmGpu) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetMode

`func (o *VmGpu) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VmGpu) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VmGpu) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VmGpu) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDeviceId

`func (o *VmGpu) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *VmGpu) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *VmGpu) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *VmGpu) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


