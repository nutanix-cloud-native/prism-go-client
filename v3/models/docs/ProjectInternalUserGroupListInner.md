# ProjectInternalUserGroupListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Indicates the action(add, delete, update) | 
**Metadata** | [**ProjectInternalAccessControlPolicyListInnerMetadata**](ProjectInternalAccessControlPolicyListInnerMetadata.md) |  | 
**UserGroup** | [**UserGroup**](UserGroup.md) |  | 

## Methods

### NewProjectInternalUserGroupListInner

`func NewProjectInternalUserGroupListInner(operation string, metadata ProjectInternalAccessControlPolicyListInnerMetadata, userGroup UserGroup, ) *ProjectInternalUserGroupListInner`

NewProjectInternalUserGroupListInner instantiates a new ProjectInternalUserGroupListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalUserGroupListInnerWithDefaults

`func NewProjectInternalUserGroupListInnerWithDefaults() *ProjectInternalUserGroupListInner`

NewProjectInternalUserGroupListInnerWithDefaults instantiates a new ProjectInternalUserGroupListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ProjectInternalUserGroupListInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectInternalUserGroupListInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectInternalUserGroupListInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetMetadata

`func (o *ProjectInternalUserGroupListInner) GetMetadata() ProjectInternalAccessControlPolicyListInnerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectInternalUserGroupListInner) GetMetadataOk() (*ProjectInternalAccessControlPolicyListInnerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectInternalUserGroupListInner) SetMetadata(v ProjectInternalAccessControlPolicyListInnerMetadata)`

SetMetadata sets Metadata field to given value.


### GetUserGroup

`func (o *ProjectInternalUserGroupListInner) GetUserGroup() UserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ProjectInternalUserGroupListInner) GetUserGroupOk() (*UserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ProjectInternalUserGroupListInner) SetUserGroup(v UserGroup)`

SetUserGroup sets UserGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


