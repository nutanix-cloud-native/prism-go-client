# VmIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VmDefStatus**](VmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Vm**](Vm.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VmMetadata**](VmMetadata.md) |  | 

## Methods

### NewVmIntentResource

`func NewVmIntentResource(metadata VmMetadata, ) *VmIntentResource`

NewVmIntentResource instantiates a new VmIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmIntentResourceWithDefaults

`func NewVmIntentResourceWithDefaults() *VmIntentResource`

NewVmIntentResourceWithDefaults instantiates a new VmIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VmIntentResource) GetStatus() VmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VmIntentResource) GetStatusOk() (*VmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VmIntentResource) SetStatus(v VmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VmIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VmIntentResource) GetSpec() Vm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmIntentResource) GetSpecOk() (*Vm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmIntentResource) SetSpec(v Vm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VmIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VmIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VmIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VmIntentResource) GetMetadata() VmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmIntentResource) GetMetadataOk() (*VmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmIntentResource) SetMetadata(v VmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


