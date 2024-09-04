# MhVmIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**MhVmDefStatus**](MhVmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**MhVm**](MhVm.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**MhVmMetadata**](MhVmMetadata.md) |  | 

## Methods

### NewMhVmIntentResource

`func NewMhVmIntentResource(metadata MhVmMetadata, ) *MhVmIntentResource`

NewMhVmIntentResource instantiates a new MhVmIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmIntentResourceWithDefaults

`func NewMhVmIntentResourceWithDefaults() *MhVmIntentResource`

NewMhVmIntentResourceWithDefaults instantiates a new MhVmIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MhVmIntentResource) GetStatus() MhVmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MhVmIntentResource) GetStatusOk() (*MhVmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MhVmIntentResource) SetStatus(v MhVmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MhVmIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *MhVmIntentResource) GetSpec() MhVm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MhVmIntentResource) GetSpecOk() (*MhVm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MhVmIntentResource) SetSpec(v MhVm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MhVmIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *MhVmIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MhVmIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MhVmIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MhVmIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *MhVmIntentResource) GetMetadata() MhVmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MhVmIntentResource) GetMetadataOk() (*MhVmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MhVmIntentResource) SetMetadata(v MhVmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


