# VmRecoveryPointIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VmRecoveryPointDefStatus**](VmRecoveryPointDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VmRecoveryPoint**](VmRecoveryPoint.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VmRecoveryPointMetadata**](VmRecoveryPointMetadata.md) |  | 

## Methods

### NewVmRecoveryPointIntentResource

`func NewVmRecoveryPointIntentResource(metadata VmRecoveryPointMetadata, ) *VmRecoveryPointIntentResource`

NewVmRecoveryPointIntentResource instantiates a new VmRecoveryPointIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointIntentResourceWithDefaults

`func NewVmRecoveryPointIntentResourceWithDefaults() *VmRecoveryPointIntentResource`

NewVmRecoveryPointIntentResourceWithDefaults instantiates a new VmRecoveryPointIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VmRecoveryPointIntentResource) GetStatus() VmRecoveryPointDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VmRecoveryPointIntentResource) GetStatusOk() (*VmRecoveryPointDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VmRecoveryPointIntentResource) SetStatus(v VmRecoveryPointDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VmRecoveryPointIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VmRecoveryPointIntentResource) GetSpec() VmRecoveryPoint`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmRecoveryPointIntentResource) GetSpecOk() (*VmRecoveryPoint, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmRecoveryPointIntentResource) SetSpec(v VmRecoveryPoint)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VmRecoveryPointIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VmRecoveryPointIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmRecoveryPointIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmRecoveryPointIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VmRecoveryPointIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VmRecoveryPointIntentResource) GetMetadata() VmRecoveryPointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointIntentResource) GetMetadataOk() (*VmRecoveryPointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointIntentResource) SetMetadata(v VmRecoveryPointMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


