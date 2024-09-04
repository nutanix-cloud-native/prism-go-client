# AwsRoleIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsRoleDefStatus**](AwsRoleDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsRoleMetadata**](AwsRoleMetadata.md) |  | 

## Methods

### NewAwsRoleIntentResource

`func NewAwsRoleIntentResource(apiVersion string, metadata AwsRoleMetadata, ) *AwsRoleIntentResource`

NewAwsRoleIntentResource instantiates a new AwsRoleIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleIntentResourceWithDefaults

`func NewAwsRoleIntentResourceWithDefaults() *AwsRoleIntentResource`

NewAwsRoleIntentResourceWithDefaults instantiates a new AwsRoleIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsRoleIntentResource) GetStatus() AwsRoleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsRoleIntentResource) GetStatusOk() (*AwsRoleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsRoleIntentResource) SetStatus(v AwsRoleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsRoleIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsRoleIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsRoleIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsRoleIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsRoleIntentResource) GetMetadata() AwsRoleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsRoleIntentResource) GetMetadataOk() (*AwsRoleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsRoleIntentResource) SetMetadata(v AwsRoleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


