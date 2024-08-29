# RoleAcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionUuidList** | Pointer to **[]string** | The list of Permissions. | [optional] 
**Name** | **string** | Role name. | 

## Methods

### NewRoleAcp

`func NewRoleAcp(name string, ) *RoleAcp`

NewRoleAcp instantiates a new RoleAcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAcpWithDefaults

`func NewRoleAcpWithDefaults() *RoleAcp`

NewRoleAcpWithDefaults instantiates a new RoleAcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionUuidList

`func (o *RoleAcp) GetPermissionUuidList() []string`

GetPermissionUuidList returns the PermissionUuidList field if non-nil, zero value otherwise.

### GetPermissionUuidListOk

`func (o *RoleAcp) GetPermissionUuidListOk() (*[]string, bool)`

GetPermissionUuidListOk returns a tuple with the PermissionUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionUuidList

`func (o *RoleAcp) SetPermissionUuidList(v []string)`

SetPermissionUuidList sets PermissionUuidList field to given value.

### HasPermissionUuidList

`func (o *RoleAcp) HasPermissionUuidList() bool`

HasPermissionUuidList returns a boolean if a field has been set.

### GetName

`func (o *RoleAcp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleAcp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleAcp) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


