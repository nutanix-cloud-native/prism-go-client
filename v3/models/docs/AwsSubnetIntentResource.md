# AwsSubnetIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsSubnetDefStatus**](AwsSubnetDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsSubnetMetadata**](AwsSubnetMetadata.md) |  | 

## Methods

### NewAwsSubnetIntentResource

`func NewAwsSubnetIntentResource(apiVersion string, metadata AwsSubnetMetadata, ) *AwsSubnetIntentResource`

NewAwsSubnetIntentResource instantiates a new AwsSubnetIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSubnetIntentResourceWithDefaults

`func NewAwsSubnetIntentResourceWithDefaults() *AwsSubnetIntentResource`

NewAwsSubnetIntentResourceWithDefaults instantiates a new AwsSubnetIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsSubnetIntentResource) GetStatus() AwsSubnetDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsSubnetIntentResource) GetStatusOk() (*AwsSubnetDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsSubnetIntentResource) SetStatus(v AwsSubnetDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsSubnetIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsSubnetIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsSubnetIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsSubnetIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsSubnetIntentResource) GetMetadata() AwsSubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsSubnetIntentResource) GetMetadataOk() (*AwsSubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsSubnetIntentResource) SetMetadata(v AwsSubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


