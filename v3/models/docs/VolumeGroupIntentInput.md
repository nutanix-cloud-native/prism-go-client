# VolumeGroupIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VolumeGroup**](VolumeGroup.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VolumeGroupMetadata**](VolumeGroupMetadata.md) |  | 

## Methods

### NewVolumeGroupIntentInput

`func NewVolumeGroupIntentInput(spec VolumeGroup, metadata VolumeGroupMetadata, ) *VolumeGroupIntentInput`

NewVolumeGroupIntentInput instantiates a new VolumeGroupIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupIntentInputWithDefaults

`func NewVolumeGroupIntentInputWithDefaults() *VolumeGroupIntentInput`

NewVolumeGroupIntentInputWithDefaults instantiates a new VolumeGroupIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VolumeGroupIntentInput) GetSpec() VolumeGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VolumeGroupIntentInput) GetSpecOk() (*VolumeGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VolumeGroupIntentInput) SetSpec(v VolumeGroup)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VolumeGroupIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VolumeGroupIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VolumeGroupIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VolumeGroupIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VolumeGroupIntentInput) GetMetadata() VolumeGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VolumeGroupIntentInput) GetMetadataOk() (*VolumeGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VolumeGroupIntentInput) SetMetadata(v VolumeGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


