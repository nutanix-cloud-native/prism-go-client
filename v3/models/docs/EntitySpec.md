# EntitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeSpec** | Pointer to [**VolumeSpec**](VolumeSpec.md) |  | [optional] 
**VmSpec** | Pointer to [**VmSpec**](VmSpec.md) |  | [optional] 

## Methods

### NewEntitySpec

`func NewEntitySpec() *EntitySpec`

NewEntitySpec instantiates a new EntitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySpecWithDefaults

`func NewEntitySpecWithDefaults() *EntitySpec`

NewEntitySpecWithDefaults instantiates a new EntitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeSpec

`func (o *EntitySpec) GetVolumeSpec() VolumeSpec`

GetVolumeSpec returns the VolumeSpec field if non-nil, zero value otherwise.

### GetVolumeSpecOk

`func (o *EntitySpec) GetVolumeSpecOk() (*VolumeSpec, bool)`

GetVolumeSpecOk returns a tuple with the VolumeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSpec

`func (o *EntitySpec) SetVolumeSpec(v VolumeSpec)`

SetVolumeSpec sets VolumeSpec field to given value.

### HasVolumeSpec

`func (o *EntitySpec) HasVolumeSpec() bool`

HasVolumeSpec returns a boolean if a field has been set.

### GetVmSpec

`func (o *EntitySpec) GetVmSpec() VmSpec`

GetVmSpec returns the VmSpec field if non-nil, zero value otherwise.

### GetVmSpecOk

`func (o *EntitySpec) GetVmSpecOk() (*VmSpec, bool)`

GetVmSpecOk returns a tuple with the VmSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSpec

`func (o *EntitySpec) SetVmSpec(v VmSpec)`

SetVmSpec sets VmSpec field to given value.

### HasVmSpec

`func (o *EntitySpec) HasVmSpec() bool`

HasVmSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


