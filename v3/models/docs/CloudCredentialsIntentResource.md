# CloudCredentialsIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CloudCredentialsDefStatus**](CloudCredentialsDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CloudCredentials**](CloudCredentials.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudCredentialsMetadata**](CloudCredentialsMetadata.md) |  | 

## Methods

### NewCloudCredentialsIntentResource

`func NewCloudCredentialsIntentResource(metadata CloudCredentialsMetadata, ) *CloudCredentialsIntentResource`

NewCloudCredentialsIntentResource instantiates a new CloudCredentialsIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsIntentResourceWithDefaults

`func NewCloudCredentialsIntentResourceWithDefaults() *CloudCredentialsIntentResource`

NewCloudCredentialsIntentResourceWithDefaults instantiates a new CloudCredentialsIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudCredentialsIntentResource) GetStatus() CloudCredentialsDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCredentialsIntentResource) GetStatusOk() (*CloudCredentialsDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCredentialsIntentResource) SetStatus(v CloudCredentialsDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCredentialsIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CloudCredentialsIntentResource) GetSpec() CloudCredentials`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudCredentialsIntentResource) GetSpecOk() (*CloudCredentials, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudCredentialsIntentResource) SetSpec(v CloudCredentials)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudCredentialsIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudCredentialsIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudCredentialsIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudCredentialsIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudCredentialsIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudCredentialsIntentResource) GetMetadata() CloudCredentialsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudCredentialsIntentResource) GetMetadataOk() (*CloudCredentialsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudCredentialsIntentResource) SetMetadata(v CloudCredentialsMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


