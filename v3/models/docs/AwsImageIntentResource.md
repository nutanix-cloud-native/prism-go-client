# AwsImageIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsImageDefStatus**](AwsImageDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsImageMetadata**](AwsImageMetadata.md) |  | 

## Methods

### NewAwsImageIntentResource

`func NewAwsImageIntentResource(apiVersion string, metadata AwsImageMetadata, ) *AwsImageIntentResource`

NewAwsImageIntentResource instantiates a new AwsImageIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsImageIntentResourceWithDefaults

`func NewAwsImageIntentResourceWithDefaults() *AwsImageIntentResource`

NewAwsImageIntentResourceWithDefaults instantiates a new AwsImageIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsImageIntentResource) GetStatus() AwsImageDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsImageIntentResource) GetStatusOk() (*AwsImageDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsImageIntentResource) SetStatus(v AwsImageDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsImageIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsImageIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsImageIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsImageIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsImageIntentResource) GetMetadata() AwsImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsImageIntentResource) GetMetadataOk() (*AwsImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsImageIntentResource) SetMetadata(v AwsImageMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


