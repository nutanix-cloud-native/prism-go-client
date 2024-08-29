# VmIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VmDefStatus**](VmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Vm**](Vm.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VmMetadata**](VmMetadata.md) |  | 

## Methods

### NewVmIntentResponse

`func NewVmIntentResponse(apiVersion string, metadata VmMetadata, ) *VmIntentResponse`

NewVmIntentResponse instantiates a new VmIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmIntentResponseWithDefaults

`func NewVmIntentResponseWithDefaults() *VmIntentResponse`

NewVmIntentResponseWithDefaults instantiates a new VmIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VmIntentResponse) GetStatus() VmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VmIntentResponse) GetStatusOk() (*VmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VmIntentResponse) SetStatus(v VmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VmIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VmIntentResponse) GetSpec() Vm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmIntentResponse) GetSpecOk() (*Vm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmIntentResponse) SetSpec(v Vm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VmIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VmIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VmIntentResponse) GetMetadata() VmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmIntentResponse) GetMetadataOk() (*VmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmIntentResponse) SetMetadata(v VmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


