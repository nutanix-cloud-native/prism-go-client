# CloudTenantIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CloudTenantDefStatus**](CloudTenantDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CloudTenant**](CloudTenant.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTenantMetadata**](CloudTenantMetadata.md) |  | 

## Methods

### NewCloudTenantIntentResource

`func NewCloudTenantIntentResource(metadata CloudTenantMetadata, ) *CloudTenantIntentResource`

NewCloudTenantIntentResource instantiates a new CloudTenantIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTenantIntentResourceWithDefaults

`func NewCloudTenantIntentResourceWithDefaults() *CloudTenantIntentResource`

NewCloudTenantIntentResourceWithDefaults instantiates a new CloudTenantIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudTenantIntentResource) GetStatus() CloudTenantDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTenantIntentResource) GetStatusOk() (*CloudTenantDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTenantIntentResource) SetStatus(v CloudTenantDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTenantIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CloudTenantIntentResource) GetSpec() CloudTenant`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTenantIntentResource) GetSpecOk() (*CloudTenant, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTenantIntentResource) SetSpec(v CloudTenant)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudTenantIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudTenantIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTenantIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTenantIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudTenantIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudTenantIntentResource) GetMetadata() CloudTenantMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTenantIntentResource) GetMetadataOk() (*CloudTenantMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTenantIntentResource) SetMetadata(v CloudTenantMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


