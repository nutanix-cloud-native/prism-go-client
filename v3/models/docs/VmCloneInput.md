# VmCloneInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**MetadataOfTheClonedVm**](MetadataOfTheClonedVm.md) |  | [optional] 
**OverrideSpec** | Pointer to [**VmCloneOverrideSpec**](VmCloneOverrideSpec.md) |  | [optional] 

## Methods

### NewVmCloneInput

`func NewVmCloneInput() *VmCloneInput`

NewVmCloneInput instantiates a new VmCloneInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCloneInputWithDefaults

`func NewVmCloneInputWithDefaults() *VmCloneInput`

NewVmCloneInputWithDefaults instantiates a new VmCloneInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *VmCloneInput) GetMetadata() MetadataOfTheClonedVm`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmCloneInput) GetMetadataOk() (*MetadataOfTheClonedVm, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmCloneInput) SetMetadata(v MetadataOfTheClonedVm)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VmCloneInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverrideSpec

`func (o *VmCloneInput) GetOverrideSpec() VmCloneOverrideSpec`

GetOverrideSpec returns the OverrideSpec field if non-nil, zero value otherwise.

### GetOverrideSpecOk

`func (o *VmCloneInput) GetOverrideSpecOk() (*VmCloneOverrideSpec, bool)`

GetOverrideSpecOk returns a tuple with the OverrideSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSpec

`func (o *VmCloneInput) SetOverrideSpec(v VmCloneOverrideSpec)`

SetOverrideSpec sets OverrideSpec field to given value.

### HasOverrideSpec

`func (o *VmCloneInput) HasOverrideSpec() bool`

HasOverrideSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


