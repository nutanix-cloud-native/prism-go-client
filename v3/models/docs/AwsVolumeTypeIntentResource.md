# AwsVolumeTypeIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsVolumeTypeDefStatus**](AwsVolumeTypeDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVolumeTypeMetadata**](AwsVolumeTypeMetadata.md) |  | 

## Methods

### NewAwsVolumeTypeIntentResource

`func NewAwsVolumeTypeIntentResource(apiVersion string, metadata AwsVolumeTypeMetadata, ) *AwsVolumeTypeIntentResource`

NewAwsVolumeTypeIntentResource instantiates a new AwsVolumeTypeIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVolumeTypeIntentResourceWithDefaults

`func NewAwsVolumeTypeIntentResourceWithDefaults() *AwsVolumeTypeIntentResource`

NewAwsVolumeTypeIntentResourceWithDefaults instantiates a new AwsVolumeTypeIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsVolumeTypeIntentResource) GetStatus() AwsVolumeTypeDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsVolumeTypeIntentResource) GetStatusOk() (*AwsVolumeTypeDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsVolumeTypeIntentResource) SetStatus(v AwsVolumeTypeDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsVolumeTypeIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVolumeTypeIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVolumeTypeIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVolumeTypeIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsVolumeTypeIntentResource) GetMetadata() AwsVolumeTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVolumeTypeIntentResource) GetMetadataOk() (*AwsVolumeTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVolumeTypeIntentResource) SetMetadata(v AwsVolumeTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


