# AccessControlPolicyResources1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleReference** | [**RoleReference**](RoleReference.md) |  | 
**UserReferenceList** | Pointer to [**[]UserReference**](UserReference.md) | The User(s) being assigned a given role. | [optional] 
**FilterList** | Pointer to [**AccessControlPolicyDetailFilterList**](AccessControlPolicyDetailFilterList.md) |  | [optional] 
**UserGroupReferenceList** | Pointer to [**[]UserGroupReference**](UserGroupReference.md) | The User group(s) being assigned a given role. | [optional] 

## Methods

### NewAccessControlPolicyResources1

`func NewAccessControlPolicyResources1(roleReference RoleReference, ) *AccessControlPolicyResources1`

NewAccessControlPolicyResources1 instantiates a new AccessControlPolicyResources1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyResources1WithDefaults

`func NewAccessControlPolicyResources1WithDefaults() *AccessControlPolicyResources1`

NewAccessControlPolicyResources1WithDefaults instantiates a new AccessControlPolicyResources1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleReference

`func (o *AccessControlPolicyResources1) GetRoleReference() RoleReference`

GetRoleReference returns the RoleReference field if non-nil, zero value otherwise.

### GetRoleReferenceOk

`func (o *AccessControlPolicyResources1) GetRoleReferenceOk() (*RoleReference, bool)`

GetRoleReferenceOk returns a tuple with the RoleReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleReference

`func (o *AccessControlPolicyResources1) SetRoleReference(v RoleReference)`

SetRoleReference sets RoleReference field to given value.


### GetUserReferenceList

`func (o *AccessControlPolicyResources1) GetUserReferenceList() []UserReference`

GetUserReferenceList returns the UserReferenceList field if non-nil, zero value otherwise.

### GetUserReferenceListOk

`func (o *AccessControlPolicyResources1) GetUserReferenceListOk() (*[]UserReference, bool)`

GetUserReferenceListOk returns a tuple with the UserReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReferenceList

`func (o *AccessControlPolicyResources1) SetUserReferenceList(v []UserReference)`

SetUserReferenceList sets UserReferenceList field to given value.

### HasUserReferenceList

`func (o *AccessControlPolicyResources1) HasUserReferenceList() bool`

HasUserReferenceList returns a boolean if a field has been set.

### GetFilterList

`func (o *AccessControlPolicyResources1) GetFilterList() AccessControlPolicyDetailFilterList`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *AccessControlPolicyResources1) GetFilterListOk() (*AccessControlPolicyDetailFilterList, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *AccessControlPolicyResources1) SetFilterList(v AccessControlPolicyDetailFilterList)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *AccessControlPolicyResources1) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.

### GetUserGroupReferenceList

`func (o *AccessControlPolicyResources1) GetUserGroupReferenceList() []UserGroupReference`

GetUserGroupReferenceList returns the UserGroupReferenceList field if non-nil, zero value otherwise.

### GetUserGroupReferenceListOk

`func (o *AccessControlPolicyResources1) GetUserGroupReferenceListOk() (*[]UserGroupReference, bool)`

GetUserGroupReferenceListOk returns a tuple with the UserGroupReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupReferenceList

`func (o *AccessControlPolicyResources1) SetUserGroupReferenceList(v []UserGroupReference)`

SetUserGroupReferenceList sets UserGroupReferenceList field to given value.

### HasUserGroupReferenceList

`func (o *AccessControlPolicyResources1) HasUserGroupReferenceList() bool`

HasUserGroupReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


