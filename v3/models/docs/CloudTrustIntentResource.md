# CloudTrustIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CloudTrustDefStatus**](CloudTrustDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CloudTrust**](CloudTrust.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTrustMetadata**](CloudTrustMetadata.md) |  | 

## Methods

### NewCloudTrustIntentResource

`func NewCloudTrustIntentResource(metadata CloudTrustMetadata, ) *CloudTrustIntentResource`

NewCloudTrustIntentResource instantiates a new CloudTrustIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustIntentResourceWithDefaults

`func NewCloudTrustIntentResourceWithDefaults() *CloudTrustIntentResource`

NewCloudTrustIntentResourceWithDefaults instantiates a new CloudTrustIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudTrustIntentResource) GetStatus() CloudTrustDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTrustIntentResource) GetStatusOk() (*CloudTrustDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTrustIntentResource) SetStatus(v CloudTrustDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTrustIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CloudTrustIntentResource) GetSpec() CloudTrust`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTrustIntentResource) GetSpecOk() (*CloudTrust, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTrustIntentResource) SetSpec(v CloudTrust)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudTrustIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudTrustIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTrustIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTrustIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudTrustIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudTrustIntentResource) GetMetadata() CloudTrustMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTrustIntentResource) GetMetadataOk() (*CloudTrustMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTrustIntentResource) SetMetadata(v CloudTrustMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


