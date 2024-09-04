# AwsVmIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsVmDefStatus**](AwsVmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AwsVm**](AwsVm.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVmMetadata**](AwsVmMetadata.md) |  | 

## Methods

### NewAwsVmIntentResource

`func NewAwsVmIntentResource(metadata AwsVmMetadata, ) *AwsVmIntentResource`

NewAwsVmIntentResource instantiates a new AwsVmIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmIntentResourceWithDefaults

`func NewAwsVmIntentResourceWithDefaults() *AwsVmIntentResource`

NewAwsVmIntentResourceWithDefaults instantiates a new AwsVmIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsVmIntentResource) GetStatus() AwsVmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsVmIntentResource) GetStatusOk() (*AwsVmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsVmIntentResource) SetStatus(v AwsVmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsVmIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AwsVmIntentResource) GetSpec() AwsVm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AwsVmIntentResource) GetSpecOk() (*AwsVm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AwsVmIntentResource) SetSpec(v AwsVm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AwsVmIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVmIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVmIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVmIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AwsVmIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AwsVmIntentResource) GetMetadata() AwsVmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVmIntentResource) GetMetadataOk() (*AwsVmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVmIntentResource) SetMetadata(v AwsVmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


