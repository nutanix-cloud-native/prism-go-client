# AwsKeyPairIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsKeyPairDefStatus**](AwsKeyPairDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsKeyPairMetadata**](AwsKeyPairMetadata.md) |  | 

## Methods

### NewAwsKeyPairIntentResource

`func NewAwsKeyPairIntentResource(apiVersion string, metadata AwsKeyPairMetadata, ) *AwsKeyPairIntentResource`

NewAwsKeyPairIntentResource instantiates a new AwsKeyPairIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKeyPairIntentResourceWithDefaults

`func NewAwsKeyPairIntentResourceWithDefaults() *AwsKeyPairIntentResource`

NewAwsKeyPairIntentResourceWithDefaults instantiates a new AwsKeyPairIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsKeyPairIntentResource) GetStatus() AwsKeyPairDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsKeyPairIntentResource) GetStatusOk() (*AwsKeyPairDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsKeyPairIntentResource) SetStatus(v AwsKeyPairDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsKeyPairIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsKeyPairIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsKeyPairIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsKeyPairIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsKeyPairIntentResource) GetMetadata() AwsKeyPairMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsKeyPairIntentResource) GetMetadataOk() (*AwsKeyPairMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsKeyPairIntentResource) SetMetadata(v AwsKeyPairMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


