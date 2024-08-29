# ProjectInternalDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlPolicyListStatus** | Pointer to [**[]ProjectInternalDefStatusAccessControlPolicyListStatusInner**](ProjectInternalDefStatusAccessControlPolicyListStatusInner.md) | The list of access control policies associates with users in the project.  | [optional] 
**ProjectStatus** | [**ProjectResourceStatus**](ProjectResourceStatus.md) |  | 

## Methods

### NewProjectInternalDefStatus

`func NewProjectInternalDefStatus(projectStatus ProjectResourceStatus, ) *ProjectInternalDefStatus`

NewProjectInternalDefStatus instantiates a new ProjectInternalDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalDefStatusWithDefaults

`func NewProjectInternalDefStatusWithDefaults() *ProjectInternalDefStatus`

NewProjectInternalDefStatusWithDefaults instantiates a new ProjectInternalDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlPolicyListStatus

`func (o *ProjectInternalDefStatus) GetAccessControlPolicyListStatus() []ProjectInternalDefStatusAccessControlPolicyListStatusInner`

GetAccessControlPolicyListStatus returns the AccessControlPolicyListStatus field if non-nil, zero value otherwise.

### GetAccessControlPolicyListStatusOk

`func (o *ProjectInternalDefStatus) GetAccessControlPolicyListStatusOk() (*[]ProjectInternalDefStatusAccessControlPolicyListStatusInner, bool)`

GetAccessControlPolicyListStatusOk returns a tuple with the AccessControlPolicyListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicyListStatus

`func (o *ProjectInternalDefStatus) SetAccessControlPolicyListStatus(v []ProjectInternalDefStatusAccessControlPolicyListStatusInner)`

SetAccessControlPolicyListStatus sets AccessControlPolicyListStatus field to given value.

### HasAccessControlPolicyListStatus

`func (o *ProjectInternalDefStatus) HasAccessControlPolicyListStatus() bool`

HasAccessControlPolicyListStatus returns a boolean if a field has been set.

### GetProjectStatus

`func (o *ProjectInternalDefStatus) GetProjectStatus() ProjectResourceStatus`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *ProjectInternalDefStatus) GetProjectStatusOk() (*ProjectResourceStatus, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *ProjectInternalDefStatus) SetProjectStatus(v ProjectResourceStatus)`

SetProjectStatus sets ProjectStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


