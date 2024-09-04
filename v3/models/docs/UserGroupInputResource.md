# UserGroupInputResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SamlUserGroup** | Pointer to [**SamlUserGroupInput**](SamlUserGroupInput.md) |  | [optional] 
**DirectoryServiceUserGroup** | Pointer to [**DirectoryServiceUserGroupInput**](DirectoryServiceUserGroupInput.md) |  | [optional] 
**DirectoryServiceOu** | Pointer to [**DirectoryServiceOuInput**](DirectoryServiceOuInput.md) |  | [optional] 

## Methods

### NewUserGroupInputResource

`func NewUserGroupInputResource() *UserGroupInputResource`

NewUserGroupInputResource instantiates a new UserGroupInputResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupInputResourceWithDefaults

`func NewUserGroupInputResourceWithDefaults() *UserGroupInputResource`

NewUserGroupInputResourceWithDefaults instantiates a new UserGroupInputResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamlUserGroup

`func (o *UserGroupInputResource) GetSamlUserGroup() SamlUserGroupInput`

GetSamlUserGroup returns the SamlUserGroup field if non-nil, zero value otherwise.

### GetSamlUserGroupOk

`func (o *UserGroupInputResource) GetSamlUserGroupOk() (*SamlUserGroupInput, bool)`

GetSamlUserGroupOk returns a tuple with the SamlUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlUserGroup

`func (o *UserGroupInputResource) SetSamlUserGroup(v SamlUserGroupInput)`

SetSamlUserGroup sets SamlUserGroup field to given value.

### HasSamlUserGroup

`func (o *UserGroupInputResource) HasSamlUserGroup() bool`

HasSamlUserGroup returns a boolean if a field has been set.

### GetDirectoryServiceUserGroup

`func (o *UserGroupInputResource) GetDirectoryServiceUserGroup() DirectoryServiceUserGroupInput`

GetDirectoryServiceUserGroup returns the DirectoryServiceUserGroup field if non-nil, zero value otherwise.

### GetDirectoryServiceUserGroupOk

`func (o *UserGroupInputResource) GetDirectoryServiceUserGroupOk() (*DirectoryServiceUserGroupInput, bool)`

GetDirectoryServiceUserGroupOk returns a tuple with the DirectoryServiceUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceUserGroup

`func (o *UserGroupInputResource) SetDirectoryServiceUserGroup(v DirectoryServiceUserGroupInput)`

SetDirectoryServiceUserGroup sets DirectoryServiceUserGroup field to given value.

### HasDirectoryServiceUserGroup

`func (o *UserGroupInputResource) HasDirectoryServiceUserGroup() bool`

HasDirectoryServiceUserGroup returns a boolean if a field has been set.

### GetDirectoryServiceOu

`func (o *UserGroupInputResource) GetDirectoryServiceOu() DirectoryServiceOuInput`

GetDirectoryServiceOu returns the DirectoryServiceOu field if non-nil, zero value otherwise.

### GetDirectoryServiceOuOk

`func (o *UserGroupInputResource) GetDirectoryServiceOuOk() (*DirectoryServiceOuInput, bool)`

GetDirectoryServiceOuOk returns a tuple with the DirectoryServiceOu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceOu

`func (o *UserGroupInputResource) SetDirectoryServiceOu(v DirectoryServiceOuInput)`

SetDirectoryServiceOu sets DirectoryServiceOu field to given value.

### HasDirectoryServiceOu

`func (o *UserGroupInputResource) HasDirectoryServiceOu() bool`

HasDirectoryServiceOu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


