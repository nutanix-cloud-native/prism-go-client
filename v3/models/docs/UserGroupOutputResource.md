# UserGroupOutputResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectsReferenceList** | Pointer to [**[]ProjectReference**](ProjectReference.md) | A list of projects the user group is part of. | [optional] 
**DirectoryServiceUserGroup** | Pointer to [**DirectoryServiceUserGroupStatus**](DirectoryServiceUserGroupStatus.md) |  | [optional] 
**AccessControlPolicyReferenceList** | Pointer to [**[]AccessControlPolicyReference**](AccessControlPolicyReference.md) | List of ACP references. | [optional] 
**DisplayName** | **string** | The display name for the user group. | 
**UserGroupType** | Pointer to **string** |  | [optional] 

## Methods

### NewUserGroupOutputResource

`func NewUserGroupOutputResource(displayName string, ) *UserGroupOutputResource`

NewUserGroupOutputResource instantiates a new UserGroupOutputResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupOutputResourceWithDefaults

`func NewUserGroupOutputResourceWithDefaults() *UserGroupOutputResource`

NewUserGroupOutputResourceWithDefaults instantiates a new UserGroupOutputResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectsReferenceList

`func (o *UserGroupOutputResource) GetProjectsReferenceList() []ProjectReference`

GetProjectsReferenceList returns the ProjectsReferenceList field if non-nil, zero value otherwise.

### GetProjectsReferenceListOk

`func (o *UserGroupOutputResource) GetProjectsReferenceListOk() (*[]ProjectReference, bool)`

GetProjectsReferenceListOk returns a tuple with the ProjectsReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsReferenceList

`func (o *UserGroupOutputResource) SetProjectsReferenceList(v []ProjectReference)`

SetProjectsReferenceList sets ProjectsReferenceList field to given value.

### HasProjectsReferenceList

`func (o *UserGroupOutputResource) HasProjectsReferenceList() bool`

HasProjectsReferenceList returns a boolean if a field has been set.

### GetDirectoryServiceUserGroup

`func (o *UserGroupOutputResource) GetDirectoryServiceUserGroup() DirectoryServiceUserGroupStatus`

GetDirectoryServiceUserGroup returns the DirectoryServiceUserGroup field if non-nil, zero value otherwise.

### GetDirectoryServiceUserGroupOk

`func (o *UserGroupOutputResource) GetDirectoryServiceUserGroupOk() (*DirectoryServiceUserGroupStatus, bool)`

GetDirectoryServiceUserGroupOk returns a tuple with the DirectoryServiceUserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceUserGroup

`func (o *UserGroupOutputResource) SetDirectoryServiceUserGroup(v DirectoryServiceUserGroupStatus)`

SetDirectoryServiceUserGroup sets DirectoryServiceUserGroup field to given value.

### HasDirectoryServiceUserGroup

`func (o *UserGroupOutputResource) HasDirectoryServiceUserGroup() bool`

HasDirectoryServiceUserGroup returns a boolean if a field has been set.

### GetAccessControlPolicyReferenceList

`func (o *UserGroupOutputResource) GetAccessControlPolicyReferenceList() []AccessControlPolicyReference`

GetAccessControlPolicyReferenceList returns the AccessControlPolicyReferenceList field if non-nil, zero value otherwise.

### GetAccessControlPolicyReferenceListOk

`func (o *UserGroupOutputResource) GetAccessControlPolicyReferenceListOk() (*[]AccessControlPolicyReference, bool)`

GetAccessControlPolicyReferenceListOk returns a tuple with the AccessControlPolicyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicyReferenceList

`func (o *UserGroupOutputResource) SetAccessControlPolicyReferenceList(v []AccessControlPolicyReference)`

SetAccessControlPolicyReferenceList sets AccessControlPolicyReferenceList field to given value.

### HasAccessControlPolicyReferenceList

`func (o *UserGroupOutputResource) HasAccessControlPolicyReferenceList() bool`

HasAccessControlPolicyReferenceList returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserGroupOutputResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserGroupOutputResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserGroupOutputResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUserGroupType

`func (o *UserGroupOutputResource) GetUserGroupType() string`

GetUserGroupType returns the UserGroupType field if non-nil, zero value otherwise.

### GetUserGroupTypeOk

`func (o *UserGroupOutputResource) GetUserGroupTypeOk() (*string, bool)`

GetUserGroupTypeOk returns a tuple with the UserGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupType

`func (o *UserGroupOutputResource) SetUserGroupType(v string)`

SetUserGroupType sets UserGroupType field to given value.

### HasUserGroupType

`func (o *UserGroupOutputResource) HasUserGroupType() bool`

HasUserGroupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


