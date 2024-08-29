# VmSpecOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmSpec** | Pointer to [**VmIntentInput**](VmIntentInput.md) |  | [optional] 
**Warnings** | Pointer to [**[]VmSpecWarning**](VmSpecWarning.md) |  | [optional] 

## Methods

### NewVmSpecOutput

`func NewVmSpecOutput() *VmSpecOutput`

NewVmSpecOutput instantiates a new VmSpecOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSpecOutputWithDefaults

`func NewVmSpecOutputWithDefaults() *VmSpecOutput`

NewVmSpecOutputWithDefaults instantiates a new VmSpecOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmSpec

`func (o *VmSpecOutput) GetVmSpec() VmIntentInput`

GetVmSpec returns the VmSpec field if non-nil, zero value otherwise.

### GetVmSpecOk

`func (o *VmSpecOutput) GetVmSpecOk() (*VmIntentInput, bool)`

GetVmSpecOk returns a tuple with the VmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpec

`func (o *VmSpecOutput) SetVmSpec(v VmIntentInput)`

SetVmSpec sets VmSpec field to given value.

### HasVmSpec

`func (o *VmSpecOutput) HasVmSpec() bool`

HasVmSpec returns a boolean if a field has been set.

### GetWarnings

`func (o *VmSpecOutput) GetWarnings() []VmSpecWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VmSpecOutput) GetWarningsOk() (*[]VmSpecWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VmSpecOutput) SetWarnings(v []VmSpecWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VmSpecOutput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


