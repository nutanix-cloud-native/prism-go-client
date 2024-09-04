# AccessControlPolicyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleReference** | Pointer to [**RoleReference**](RoleReference.md) |  | [optional] 
**UserReferenceList** | Pointer to [**[]UserReference**](UserReference.md) | The User(s) being assigned a given role. | [optional] 
**FilterList** | Pointer to [**AccessControlPolicyDetailFilterList**](AccessControlPolicyDetailFilterList.md) |  | [optional] 
**UserGroupReferenceList** | Pointer to [**[]UserGroupReference**](UserGroupReference.md) | The User group(s) being assigned a given role. | [optional] 

## Methods

### NewAccessControlPolicyResources

`func NewAccessControlPolicyResources() *AccessControlPolicyResources`

NewAccessControlPolicyResources instantiates a new AccessControlPolicyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyResourcesWithDefaults

`func NewAccessControlPolicyResourcesWithDefaults() *AccessControlPolicyResources`

NewAccessControlPolicyResourcesWithDefaults instantiates a new AccessControlPolicyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleReference

`func (o *AccessControlPolicyResources) GetRoleReference() RoleReference`

GetRoleReference returns the RoleReference field if non-nil, zero value otherwise.

### GetRoleReferenceOk

`func (o *AccessControlPolicyResources) GetRoleReferenceOk() (*RoleReference, bool)`

GetRoleReferenceOk returns a tuple with the RoleReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleReference

`func (o *AccessControlPolicyResources) SetRoleReference(v RoleReference)`

SetRoleReference sets RoleReference field to given value.

### HasRoleReference

`func (o *AccessControlPolicyResources) HasRoleReference() bool`

HasRoleReference returns a boolean if a field has been set.

### GetUserReferenceList

`func (o *AccessControlPolicyResources) GetUserReferenceList() []UserReference`

GetUserReferenceList returns the UserReferenceList field if non-nil, zero value otherwise.

### GetUserReferenceListOk

`func (o *AccessControlPolicyResources) GetUserReferenceListOk() (*[]UserReference, bool)`

GetUserReferenceListOk returns a tuple with the UserReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReferenceList

`func (o *AccessControlPolicyResources) SetUserReferenceList(v []UserReference)`

SetUserReferenceList sets UserReferenceList field to given value.

### HasUserReferenceList

`func (o *AccessControlPolicyResources) HasUserReferenceList() bool`

HasUserReferenceList returns a boolean if a field has been set.

### GetFilterList

`func (o *AccessControlPolicyResources) GetFilterList() AccessControlPolicyDetailFilterList`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *AccessControlPolicyResources) GetFilterListOk() (*AccessControlPolicyDetailFilterList, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *AccessControlPolicyResources) SetFilterList(v AccessControlPolicyDetailFilterList)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *AccessControlPolicyResources) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.

### GetUserGroupReferenceList

`func (o *AccessControlPolicyResources) GetUserGroupReferenceList() []UserGroupReference`

GetUserGroupReferenceList returns the UserGroupReferenceList field if non-nil, zero value otherwise.

### GetUserGroupReferenceListOk

`func (o *AccessControlPolicyResources) GetUserGroupReferenceListOk() (*[]UserGroupReference, bool)`

GetUserGroupReferenceListOk returns a tuple with the UserGroupReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupReferenceList

`func (o *AccessControlPolicyResources) SetUserGroupReferenceList(v []UserGroupReference)`

SetUserGroupReferenceList sets UserGroupReferenceList field to given value.

### HasUserGroupReferenceList

`func (o *AccessControlPolicyResources) HasUserGroupReferenceList() bool`

HasUserGroupReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


