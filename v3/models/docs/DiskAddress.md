# DiskAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIndex** | **int32** |  | 
**AdapterType** | **string** |  | 

## Methods

### NewDiskAddress

`func NewDiskAddress(deviceIndex int32, adapterType string, ) *DiskAddress`

NewDiskAddress instantiates a new DiskAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskAddressWithDefaults

`func NewDiskAddressWithDefaults() *DiskAddress`

NewDiskAddressWithDefaults instantiates a new DiskAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIndex

`func (o *DiskAddress) GetDeviceIndex() int32`

GetDeviceIndex returns the DeviceIndex field if non-nil, zero value otherwise.

### GetDeviceIndexOk

`func (o *DiskAddress) GetDeviceIndexOk() (*int32, bool)`

GetDeviceIndexOk returns a tuple with the DeviceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIndex

`func (o *DiskAddress) SetDeviceIndex(v int32)`

SetDeviceIndex sets DeviceIndex field to given value.


### GetAdapterType

`func (o *DiskAddress) GetAdapterType() string`

GetAdapterType returns the AdapterType field if non-nil, zero value otherwise.

### GetAdapterTypeOk

`func (o *DiskAddress) GetAdapterTypeOk() (*string, bool)`

GetAdapterTypeOk returns a tuple with the AdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterType

`func (o *DiskAddress) SetAdapterType(v string)`

SetAdapterType sets AdapterType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


