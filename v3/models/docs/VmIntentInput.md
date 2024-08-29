# VmIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Vm**](Vm.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VmMetadata**](VmMetadata.md) |  | 

## Methods

### NewVmIntentInput

`func NewVmIntentInput(spec Vm, metadata VmMetadata, ) *VmIntentInput`

NewVmIntentInput instantiates a new VmIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmIntentInputWithDefaults

`func NewVmIntentInputWithDefaults() *VmIntentInput`

NewVmIntentInputWithDefaults instantiates a new VmIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VmIntentInput) GetSpec() Vm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmIntentInput) GetSpecOk() (*Vm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmIntentInput) SetSpec(v Vm)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VmIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VmIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VmIntentInput) GetMetadata() VmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmIntentInput) GetMetadataOk() (*VmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmIntentInput) SetMetadata(v VmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


