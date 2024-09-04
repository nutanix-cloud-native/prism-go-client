# PermissionIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**PermissionDefStatus**](PermissionDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**PermissionMetadata**](PermissionMetadata.md) |  | 

## Methods

### NewPermissionIntentResource

`func NewPermissionIntentResource(metadata PermissionMetadata, ) *PermissionIntentResource`

NewPermissionIntentResource instantiates a new PermissionIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionIntentResourceWithDefaults

`func NewPermissionIntentResourceWithDefaults() *PermissionIntentResource`

NewPermissionIntentResourceWithDefaults instantiates a new PermissionIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PermissionIntentResource) GetStatus() PermissionDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PermissionIntentResource) GetStatusOk() (*PermissionDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PermissionIntentResource) SetStatus(v PermissionDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PermissionIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *PermissionIntentResource) GetSpec() Permission`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PermissionIntentResource) GetSpecOk() (*Permission, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PermissionIntentResource) SetSpec(v Permission)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PermissionIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *PermissionIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PermissionIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PermissionIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PermissionIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *PermissionIntentResource) GetMetadata() PermissionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionIntentResource) GetMetadataOk() (*PermissionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionIntentResource) SetMetadata(v PermissionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


