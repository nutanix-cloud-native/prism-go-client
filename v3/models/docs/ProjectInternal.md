# ProjectInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlPolicyList** | Pointer to [**[]ProjectInternalAccessControlPolicyListInner**](ProjectInternalAccessControlPolicyListInner.md) | The list of ACPs to be attached to the users belonging to a project.  | [optional] 
**ProjectDetail** | [**ProjectDetails**](ProjectDetails.md) |  | 
**UserList** | Pointer to [**[]ProjectInternalUserListInner**](ProjectInternalUserListInner.md) | The list of user specification to be associated with the project.  | [optional] 
**UserGroupList** | Pointer to [**[]ProjectInternalUserGroupListInner**](ProjectInternalUserGroupListInner.md) | The list of user group specification to be associated with the project.  | [optional] 

## Methods

### NewProjectInternal

`func NewProjectInternal(projectDetail ProjectDetails, ) *ProjectInternal`

NewProjectInternal instantiates a new ProjectInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalWithDefaults

`func NewProjectInternalWithDefaults() *ProjectInternal`

NewProjectInternalWithDefaults instantiates a new ProjectInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlPolicyList

`func (o *ProjectInternal) GetAccessControlPolicyList() []ProjectInternalAccessControlPolicyListInner`

GetAccessControlPolicyList returns the AccessControlPolicyList field if non-nil, zero value otherwise.

### GetAccessControlPolicyListOk

`func (o *ProjectInternal) GetAccessControlPolicyListOk() (*[]ProjectInternalAccessControlPolicyListInner, bool)`

GetAccessControlPolicyListOk returns a tuple with the AccessControlPolicyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicyList

`func (o *ProjectInternal) SetAccessControlPolicyList(v []ProjectInternalAccessControlPolicyListInner)`

SetAccessControlPolicyList sets AccessControlPolicyList field to given value.

### HasAccessControlPolicyList

`func (o *ProjectInternal) HasAccessControlPolicyList() bool`

HasAccessControlPolicyList returns a boolean if a field has been set.

### GetProjectDetail

`func (o *ProjectInternal) GetProjectDetail() ProjectDetails`

GetProjectDetail returns the ProjectDetail field if non-nil, zero value otherwise.

### GetProjectDetailOk

`func (o *ProjectInternal) GetProjectDetailOk() (*ProjectDetails, bool)`

GetProjectDetailOk returns a tuple with the ProjectDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDetail

`func (o *ProjectInternal) SetProjectDetail(v ProjectDetails)`

SetProjectDetail sets ProjectDetail field to given value.


### GetUserList

`func (o *ProjectInternal) GetUserList() []ProjectInternalUserListInner`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *ProjectInternal) GetUserListOk() (*[]ProjectInternalUserListInner, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *ProjectInternal) SetUserList(v []ProjectInternalUserListInner)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *ProjectInternal) HasUserList() bool`

HasUserList returns a boolean if a field has been set.

### GetUserGroupList

`func (o *ProjectInternal) GetUserGroupList() []ProjectInternalUserGroupListInner`

GetUserGroupList returns the UserGroupList field if non-nil, zero value otherwise.

### GetUserGroupListOk

`func (o *ProjectInternal) GetUserGroupListOk() (*[]ProjectInternalUserGroupListInner, bool)`

GetUserGroupListOk returns a tuple with the UserGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupList

`func (o *ProjectInternal) SetUserGroupList(v []ProjectInternalUserGroupListInner)`

SetUserGroupList sets UserGroupList field to given value.

### HasUserGroupList

`func (o *ProjectInternal) HasUserGroupList() bool`

HasUserGroupList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


