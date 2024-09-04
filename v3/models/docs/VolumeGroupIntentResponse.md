# VolumeGroupIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VolumeGroupDefStatus**](VolumeGroupDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VolumeGroup**](VolumeGroup.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VolumeGroupMetadata**](VolumeGroupMetadata.md) |  | 

## Methods

### NewVolumeGroupIntentResponse

`func NewVolumeGroupIntentResponse(apiVersion string, metadata VolumeGroupMetadata, ) *VolumeGroupIntentResponse`

NewVolumeGroupIntentResponse instantiates a new VolumeGroupIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupIntentResponseWithDefaults

`func NewVolumeGroupIntentResponseWithDefaults() *VolumeGroupIntentResponse`

NewVolumeGroupIntentResponseWithDefaults instantiates a new VolumeGroupIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VolumeGroupIntentResponse) GetStatus() VolumeGroupDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeGroupIntentResponse) GetStatusOk() (*VolumeGroupDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeGroupIntentResponse) SetStatus(v VolumeGroupDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeGroupIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VolumeGroupIntentResponse) GetSpec() VolumeGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VolumeGroupIntentResponse) GetSpecOk() (*VolumeGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VolumeGroupIntentResponse) SetSpec(v VolumeGroup)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VolumeGroupIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VolumeGroupIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VolumeGroupIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VolumeGroupIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VolumeGroupIntentResponse) GetMetadata() VolumeGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VolumeGroupIntentResponse) GetMetadataOk() (*VolumeGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VolumeGroupIntentResponse) SetMetadata(v VolumeGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


