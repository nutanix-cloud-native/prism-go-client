# IdentityProviderIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**IdentityProviderDefStatus**](IdentityProviderDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**IdentityProvider**](IdentityProvider.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**IdentityProviderMetadata**](IdentityProviderMetadata.md) |  | 

## Methods

### NewIdentityProviderIntentResponse

`func NewIdentityProviderIntentResponse(apiVersion string, metadata IdentityProviderMetadata, ) *IdentityProviderIntentResponse`

NewIdentityProviderIntentResponse instantiates a new IdentityProviderIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderIntentResponseWithDefaults

`func NewIdentityProviderIntentResponseWithDefaults() *IdentityProviderIntentResponse`

NewIdentityProviderIntentResponseWithDefaults instantiates a new IdentityProviderIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IdentityProviderIntentResponse) GetStatus() IdentityProviderDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProviderIntentResponse) GetStatusOk() (*IdentityProviderDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProviderIntentResponse) SetStatus(v IdentityProviderDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProviderIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *IdentityProviderIntentResponse) GetSpec() IdentityProvider`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IdentityProviderIntentResponse) GetSpecOk() (*IdentityProvider, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IdentityProviderIntentResponse) SetSpec(v IdentityProvider)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IdentityProviderIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *IdentityProviderIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IdentityProviderIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IdentityProviderIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *IdentityProviderIntentResponse) GetMetadata() IdentityProviderMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityProviderIntentResponse) GetMetadataOk() (*IdentityProviderMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityProviderIntentResponse) SetMetadata(v IdentityProviderMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


