# CloudTenantIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**CloudTenant**](CloudTenant.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTenantMetadata**](CloudTenantMetadata.md) |  | 

## Methods

### NewCloudTenantIntentInput

`func NewCloudTenantIntentInput(spec CloudTenant, metadata CloudTenantMetadata, ) *CloudTenantIntentInput`

NewCloudTenantIntentInput instantiates a new CloudTenantIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTenantIntentInputWithDefaults

`func NewCloudTenantIntentInputWithDefaults() *CloudTenantIntentInput`

NewCloudTenantIntentInputWithDefaults instantiates a new CloudTenantIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CloudTenantIntentInput) GetSpec() CloudTenant`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTenantIntentInput) GetSpecOk() (*CloudTenant, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTenantIntentInput) SetSpec(v CloudTenant)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *CloudTenantIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTenantIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTenantIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudTenantIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudTenantIntentInput) GetMetadata() CloudTenantMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTenantIntentInput) GetMetadataOk() (*CloudTenantMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTenantIntentInput) SetMetadata(v CloudTenantMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


