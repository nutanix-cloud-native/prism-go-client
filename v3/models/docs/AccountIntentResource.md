# AccountIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AccountDefStatus**](AccountDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Account**](Account.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AccountMetadata**](AccountMetadata.md) |  | 

## Methods

### NewAccountIntentResource

`func NewAccountIntentResource(metadata AccountMetadata, ) *AccountIntentResource`

NewAccountIntentResource instantiates a new AccountIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIntentResourceWithDefaults

`func NewAccountIntentResourceWithDefaults() *AccountIntentResource`

NewAccountIntentResourceWithDefaults instantiates a new AccountIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountIntentResource) GetStatus() AccountDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountIntentResource) GetStatusOk() (*AccountDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountIntentResource) SetStatus(v AccountDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AccountIntentResource) GetSpec() Account`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AccountIntentResource) GetSpecOk() (*Account, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AccountIntentResource) SetSpec(v Account)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AccountIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AccountIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AccountIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AccountIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AccountIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountIntentResource) GetMetadata() AccountMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountIntentResource) GetMetadataOk() (*AccountMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountIntentResource) SetMetadata(v AccountMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


