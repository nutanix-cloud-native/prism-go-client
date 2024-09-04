# VolumeGroupIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VolumeGroupDefStatus**](VolumeGroupDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VolumeGroup**](VolumeGroup.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VolumeGroupMetadata**](VolumeGroupMetadata.md) |  | 

## Methods

### NewVolumeGroupIntentResource

`func NewVolumeGroupIntentResource(metadata VolumeGroupMetadata, ) *VolumeGroupIntentResource`

NewVolumeGroupIntentResource instantiates a new VolumeGroupIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupIntentResourceWithDefaults

`func NewVolumeGroupIntentResourceWithDefaults() *VolumeGroupIntentResource`

NewVolumeGroupIntentResourceWithDefaults instantiates a new VolumeGroupIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VolumeGroupIntentResource) GetStatus() VolumeGroupDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeGroupIntentResource) GetStatusOk() (*VolumeGroupDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeGroupIntentResource) SetStatus(v VolumeGroupDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeGroupIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VolumeGroupIntentResource) GetSpec() VolumeGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VolumeGroupIntentResource) GetSpecOk() (*VolumeGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VolumeGroupIntentResource) SetSpec(v VolumeGroup)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VolumeGroupIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VolumeGroupIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VolumeGroupIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VolumeGroupIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VolumeGroupIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VolumeGroupIntentResource) GetMetadata() VolumeGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VolumeGroupIntentResource) GetMetadataOk() (*VolumeGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VolumeGroupIntentResource) SetMetadata(v VolumeGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


