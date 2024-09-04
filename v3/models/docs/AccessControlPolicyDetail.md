# AccessControlPolicyDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**RoleAcp**](RoleAcp.md) |  | [optional] 
**FilterList** | Pointer to [**AccessControlPolicyDetailFilterList**](AccessControlPolicyDetailFilterList.md) |  | [optional] 

## Methods

### NewAccessControlPolicyDetail

`func NewAccessControlPolicyDetail() *AccessControlPolicyDetail`

NewAccessControlPolicyDetail instantiates a new AccessControlPolicyDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyDetailWithDefaults

`func NewAccessControlPolicyDetailWithDefaults() *AccessControlPolicyDetail`

NewAccessControlPolicyDetailWithDefaults instantiates a new AccessControlPolicyDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AccessControlPolicyDetail) GetRole() RoleAcp`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AccessControlPolicyDetail) GetRoleOk() (*RoleAcp, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AccessControlPolicyDetail) SetRole(v RoleAcp)`

SetRole sets Role field to given value.

### HasRole

`func (o *AccessControlPolicyDetail) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFilterList

`func (o *AccessControlPolicyDetail) GetFilterList() AccessControlPolicyDetailFilterList`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *AccessControlPolicyDetail) GetFilterListOk() (*AccessControlPolicyDetailFilterList, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *AccessControlPolicyDetail) SetFilterList(v AccessControlPolicyDetailFilterList)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *AccessControlPolicyDetail) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


