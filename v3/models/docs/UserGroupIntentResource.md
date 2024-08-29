# UserGroupIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**UserGroupDefStatus**](UserGroupDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**UserGroupMetadata**](UserGroupMetadata.md) |  | 

## Methods

### NewUserGroupIntentResource

`func NewUserGroupIntentResource(metadata UserGroupMetadata, ) *UserGroupIntentResource`

NewUserGroupIntentResource instantiates a new UserGroupIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupIntentResourceWithDefaults

`func NewUserGroupIntentResourceWithDefaults() *UserGroupIntentResource`

NewUserGroupIntentResourceWithDefaults instantiates a new UserGroupIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserGroupIntentResource) GetStatus() UserGroupDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserGroupIntentResource) GetStatusOk() (*UserGroupDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserGroupIntentResource) SetStatus(v UserGroupDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserGroupIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *UserGroupIntentResource) GetSpec() UserGroup`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UserGroupIntentResource) GetSpecOk() (*UserGroup, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UserGroupIntentResource) SetSpec(v UserGroup)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UserGroupIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *UserGroupIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UserGroupIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UserGroupIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UserGroupIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *UserGroupIntentResource) GetMetadata() UserGroupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserGroupIntentResource) GetMetadataOk() (*UserGroupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserGroupIntentResource) SetMetadata(v UserGroupMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


