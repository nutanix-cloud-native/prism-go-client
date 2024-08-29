# ProjectInternalUserListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Indicates the action(add, delete, update) | 
**User** | [**User**](User.md) |  | 
**Metadata** | [**ProjectInternalAccessControlPolicyListInnerMetadata**](ProjectInternalAccessControlPolicyListInnerMetadata.md) |  | 

## Methods

### NewProjectInternalUserListInner

`func NewProjectInternalUserListInner(operation string, user User, metadata ProjectInternalAccessControlPolicyListInnerMetadata, ) *ProjectInternalUserListInner`

NewProjectInternalUserListInner instantiates a new ProjectInternalUserListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalUserListInnerWithDefaults

`func NewProjectInternalUserListInnerWithDefaults() *ProjectInternalUserListInner`

NewProjectInternalUserListInnerWithDefaults instantiates a new ProjectInternalUserListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ProjectInternalUserListInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectInternalUserListInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectInternalUserListInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetUser

`func (o *ProjectInternalUserListInner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProjectInternalUserListInner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProjectInternalUserListInner) SetUser(v User)`

SetUser sets User field to given value.


### GetMetadata

`func (o *ProjectInternalUserListInner) GetMetadata() ProjectInternalAccessControlPolicyListInnerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectInternalUserListInner) GetMetadataOk() (*ProjectInternalAccessControlPolicyListInnerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectInternalUserListInner) SetMetadata(v ProjectInternalAccessControlPolicyListInnerMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


