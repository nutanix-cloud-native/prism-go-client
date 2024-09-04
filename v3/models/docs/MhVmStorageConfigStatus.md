# MhVmStorageConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosPolicy** | Pointer to [**StorageQosPolicyConfigOutput**](StorageQosPolicyConfigOutput.md) |  | [optional] 
**FlashMode** | Pointer to **string** | State of the storage policy to pin virtual disks to the hot tier. When specified as a VM attribute, the storage policy applies to all virtual disks of the VM unless overridden by the same attribute specified for a virtual disk.  | [optional] 

## Methods

### NewMhVmStorageConfigStatus

`func NewMhVmStorageConfigStatus() *MhVmStorageConfigStatus`

NewMhVmStorageConfigStatus instantiates a new MhVmStorageConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmStorageConfigStatusWithDefaults

`func NewMhVmStorageConfigStatusWithDefaults() *MhVmStorageConfigStatus`

NewMhVmStorageConfigStatusWithDefaults instantiates a new MhVmStorageConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosPolicy

`func (o *MhVmStorageConfigStatus) GetQosPolicy() StorageQosPolicyConfigOutput`

GetQosPolicy returns the QosPolicy field if non-nil, zero value otherwise.

### GetQosPolicyOk

`func (o *MhVmStorageConfigStatus) GetQosPolicyOk() (*StorageQosPolicyConfigOutput, bool)`

GetQosPolicyOk returns a tuple with the QosPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicy

`func (o *MhVmStorageConfigStatus) SetQosPolicy(v StorageQosPolicyConfigOutput)`

SetQosPolicy sets QosPolicy field to given value.

### HasQosPolicy

`func (o *MhVmStorageConfigStatus) HasQosPolicy() bool`

HasQosPolicy returns a boolean if a field has been set.

### GetFlashMode

`func (o *MhVmStorageConfigStatus) GetFlashMode() string`

GetFlashMode returns the FlashMode field if non-nil, zero value otherwise.

### GetFlashModeOk

`func (o *MhVmStorageConfigStatus) GetFlashModeOk() (*string, bool)`

GetFlashModeOk returns a tuple with the FlashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashMode

`func (o *MhVmStorageConfigStatus) SetFlashMode(v string)`

SetFlashMode sets FlashMode field to given value.

### HasFlashMode

`func (o *MhVmStorageConfigStatus) HasFlashMode() bool`

HasFlashMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


