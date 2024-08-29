# AwsElasticIpIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsElasticIpDefStatus**](AwsElasticIpDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsElasticIpMetadata**](AwsElasticIpMetadata.md) |  | 

## Methods

### NewAwsElasticIpIntentResource

`func NewAwsElasticIpIntentResource(apiVersion string, metadata AwsElasticIpMetadata, ) *AwsElasticIpIntentResource`

NewAwsElasticIpIntentResource instantiates a new AwsElasticIpIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsElasticIpIntentResourceWithDefaults

`func NewAwsElasticIpIntentResourceWithDefaults() *AwsElasticIpIntentResource`

NewAwsElasticIpIntentResourceWithDefaults instantiates a new AwsElasticIpIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsElasticIpIntentResource) GetStatus() AwsElasticIpDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsElasticIpIntentResource) GetStatusOk() (*AwsElasticIpDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsElasticIpIntentResource) SetStatus(v AwsElasticIpDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsElasticIpIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsElasticIpIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsElasticIpIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsElasticIpIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsElasticIpIntentResource) GetMetadata() AwsElasticIpMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsElasticIpIntentResource) GetMetadataOk() (*AwsElasticIpMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsElasticIpIntentResource) SetMetadata(v AwsElasticIpMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


