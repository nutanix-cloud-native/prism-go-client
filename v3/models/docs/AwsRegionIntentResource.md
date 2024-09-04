# AwsRegionIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsRegionDefStatus**](AwsRegionDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsRegionMetadata**](AwsRegionMetadata.md) |  | 

## Methods

### NewAwsRegionIntentResource

`func NewAwsRegionIntentResource(apiVersion string, metadata AwsRegionMetadata, ) *AwsRegionIntentResource`

NewAwsRegionIntentResource instantiates a new AwsRegionIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRegionIntentResourceWithDefaults

`func NewAwsRegionIntentResourceWithDefaults() *AwsRegionIntentResource`

NewAwsRegionIntentResourceWithDefaults instantiates a new AwsRegionIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsRegionIntentResource) GetStatus() AwsRegionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsRegionIntentResource) GetStatusOk() (*AwsRegionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsRegionIntentResource) SetStatus(v AwsRegionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsRegionIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsRegionIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsRegionIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsRegionIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsRegionIntentResource) GetMetadata() AwsRegionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsRegionIntentResource) GetMetadataOk() (*AwsRegionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsRegionIntentResource) SetMetadata(v AwsRegionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


