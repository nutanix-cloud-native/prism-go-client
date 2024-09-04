# IdentityProviderIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**IdentityProviderDefStatus**](IdentityProviderDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**IdentityProvider**](IdentityProvider.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**IdentityProviderMetadata**](IdentityProviderMetadata.md) |  | 

## Methods

### NewIdentityProviderIntentResource

`func NewIdentityProviderIntentResource(metadata IdentityProviderMetadata, ) *IdentityProviderIntentResource`

NewIdentityProviderIntentResource instantiates a new IdentityProviderIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIntentResourceWithDefaults

`func NewIdentityProviderIntentResourceWithDefaults() *IdentityProviderIntentResource`

NewIdentityProviderIntentResourceWithDefaults instantiates a new IdentityProviderIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IdentityProviderIntentResource) GetStatus() IdentityProviderDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProviderIntentResource) GetStatusOk() (*IdentityProviderDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProviderIntentResource) SetStatus(v IdentityProviderDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProviderIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *IdentityProviderIntentResource) GetSpec() IdentityProvider`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IdentityProviderIntentResource) GetSpecOk() (*IdentityProvider, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IdentityProviderIntentResource) SetSpec(v IdentityProvider)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IdentityProviderIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *IdentityProviderIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IdentityProviderIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IdentityProviderIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IdentityProviderIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityProviderIntentResource) GetMetadata() IdentityProviderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityProviderIntentResource) GetMetadataOk() (*IdentityProviderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityProviderIntentResource) SetMetadata(v IdentityProviderMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


