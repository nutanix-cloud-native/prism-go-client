# VMResources1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VnumaConfig** | Pointer to [**VmVnumaConfig**](VmVnumaConfig.md) |  | [optional] 
**NicList** | Pointer to [**[]VmNic**](VmNic.md) | List of Virtual NICs to be attached to the VM. | [optional] 
**GpuList** | Pointer to [**[]VmGpu**](VmGpu.md) | GPUs attached to the VM. | [optional] 

## Methods

### NewVMResources1

`func NewVMResources1() *VMResources1`

NewVMResources1 instantiates a new VMResources1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMResources1WithDefaults

`func NewVMResources1WithDefaults() *VMResources1`

NewVMResources1WithDefaults instantiates a new VMResources1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnumaConfig

`func (o *VMResources1) GetVnumaConfig() VmVnumaConfig`

GetVnumaConfig returns the VnumaConfig field if non-nil, zero value otherwise.

### GetVnumaConfigOk

`func (o *VMResources1) GetVnumaConfigOk() (*VmVnumaConfig, bool)`

GetVnumaConfigOk returns a tuple with the VnumaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnumaConfig

`func (o *VMResources1) SetVnumaConfig(v VmVnumaConfig)`

SetVnumaConfig sets VnumaConfig field to given value.

### HasVnumaConfig

`func (o *VMResources1) HasVnumaConfig() bool`

HasVnumaConfig returns a boolean if a field has been set.

### GetNicList

`func (o *VMResources1) GetNicList() []VmNic`

GetNicList returns the NicList field if non-nil, zero value otherwise.

### GetNicListOk

`func (o *VMResources1) GetNicListOk() (*[]VmNic, bool)`

GetNicListOk returns a tuple with the NicList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicList

`func (o *VMResources1) SetNicList(v []VmNic)`

SetNicList sets NicList field to given value.

### HasNicList

`func (o *VMResources1) HasNicList() bool`

HasNicList returns a boolean if a field has been set.

### GetGpuList

`func (o *VMResources1) GetGpuList() []VmGpu`

GetGpuList returns the GpuList field if non-nil, zero value otherwise.

### GetGpuListOk

`func (o *VMResources1) GetGpuListOk() (*[]VmGpu, bool)`

GetGpuListOk returns a tuple with the GpuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuList

`func (o *VMResources1) SetGpuList(v []VmGpu)`

SetGpuList sets GpuList field to given value.

### HasGpuList

`func (o *VMResources1) HasGpuList() bool`

HasGpuList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


