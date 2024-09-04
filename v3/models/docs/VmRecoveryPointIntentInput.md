# VmRecoveryPointIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VmRecoveryPoint**](VmRecoveryPoint.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VmRecoveryPointMetadata**](VmRecoveryPointMetadata.md) |  | 

## Methods

### NewVmRecoveryPointIntentInput

`func NewVmRecoveryPointIntentInput(spec VmRecoveryPoint, metadata VmRecoveryPointMetadata, ) *VmRecoveryPointIntentInput`

NewVmRecoveryPointIntentInput instantiates a new VmRecoveryPointIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointIntentInputWithDefaults

`func NewVmRecoveryPointIntentInputWithDefaults() *VmRecoveryPointIntentInput`

NewVmRecoveryPointIntentInputWithDefaults instantiates a new VmRecoveryPointIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VmRecoveryPointIntentInput) GetSpec() VmRecoveryPoint`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmRecoveryPointIntentInput) GetSpecOk() (*VmRecoveryPoint, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmRecoveryPointIntentInput) SetSpec(v VmRecoveryPoint)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VmRecoveryPointIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmRecoveryPointIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmRecoveryPointIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VmRecoveryPointIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VmRecoveryPointIntentInput) GetMetadata() VmRecoveryPointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointIntentInput) GetMetadataOk() (*VmRecoveryPointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointIntentInput) SetMetadata(v VmRecoveryPointMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


