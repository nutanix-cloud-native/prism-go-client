# AwsVpcIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsVpcDefStatus**](AwsVpcDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVpcMetadata**](AwsVpcMetadata.md) |  | 

## Methods

### NewAwsVpcIntentResource

`func NewAwsVpcIntentResource(apiVersion string, metadata AwsVpcMetadata, ) *AwsVpcIntentResource`

NewAwsVpcIntentResource instantiates a new AwsVpcIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcIntentResourceWithDefaults

`func NewAwsVpcIntentResourceWithDefaults() *AwsVpcIntentResource`

NewAwsVpcIntentResourceWithDefaults instantiates a new AwsVpcIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsVpcIntentResource) GetStatus() AwsVpcDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsVpcIntentResource) GetStatusOk() (*AwsVpcDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsVpcIntentResource) SetStatus(v AwsVpcDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsVpcIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVpcIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVpcIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVpcIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsVpcIntentResource) GetMetadata() AwsVpcMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVpcIntentResource) GetMetadataOk() (*AwsVpcMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVpcIntentResource) SetMetadata(v AwsVpcMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


