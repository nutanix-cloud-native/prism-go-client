# VmRecoveryPointIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VmRecoveryPointDefStatus**](VmRecoveryPointDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VmRecoveryPoint**](VmRecoveryPoint.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VmRecoveryPointMetadata**](VmRecoveryPointMetadata.md) |  | 

## Methods

### NewVmRecoveryPointIntentResponse

`func NewVmRecoveryPointIntentResponse(apiVersion string, metadata VmRecoveryPointMetadata, ) *VmRecoveryPointIntentResponse`

NewVmRecoveryPointIntentResponse instantiates a new VmRecoveryPointIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRecoveryPointIntentResponseWithDefaults

`func NewVmRecoveryPointIntentResponseWithDefaults() *VmRecoveryPointIntentResponse`

NewVmRecoveryPointIntentResponseWithDefaults instantiates a new VmRecoveryPointIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VmRecoveryPointIntentResponse) GetStatus() VmRecoveryPointDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VmRecoveryPointIntentResponse) GetStatusOk() (*VmRecoveryPointDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VmRecoveryPointIntentResponse) SetStatus(v VmRecoveryPointDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VmRecoveryPointIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VmRecoveryPointIntentResponse) GetSpec() VmRecoveryPoint`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmRecoveryPointIntentResponse) GetSpecOk() (*VmRecoveryPoint, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmRecoveryPointIntentResponse) SetSpec(v VmRecoveryPoint)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VmRecoveryPointIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VmRecoveryPointIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VmRecoveryPointIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VmRecoveryPointIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VmRecoveryPointIntentResponse) GetMetadata() VmRecoveryPointMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmRecoveryPointIntentResponse) GetMetadataOk() (*VmRecoveryPointMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmRecoveryPointIntentResponse) SetMetadata(v VmRecoveryPointMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


