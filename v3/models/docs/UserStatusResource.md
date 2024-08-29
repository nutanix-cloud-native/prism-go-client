# UserStatusResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlPolicyReferenceList** | Pointer to [**[]AccessControlPolicyReference**](AccessControlPolicyReference.md) | List of ACP references. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user (common name) provided by the directory service.  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**ResourceUsageSummary** | Pointer to [**UserStatusResourceResourceUsageSummary**](UserStatusResourceResourceUsageSummary.md) |  | [optional] 
**ProjectsReferenceList** | Pointer to [**[]ProjectReference**](ProjectReference.md) | A list of projects the user is part of. | [optional] 
**IdentityProviderUser** | Pointer to [**IdentityProviderUser**](IdentityProviderUser.md) |  | [optional] 
**DirectoryServiceUser** | Pointer to [**DirectoryServiceUserStatus**](DirectoryServiceUserStatus.md) |  | [optional] 

## Methods

### NewUserStatusResource

`func NewUserStatusResource() *UserStatusResource`

NewUserStatusResource instantiates a new UserStatusResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusResourceWithDefaults

`func NewUserStatusResourceWithDefaults() *UserStatusResource`

NewUserStatusResourceWithDefaults instantiates a new UserStatusResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlPolicyReferenceList

`func (o *UserStatusResource) GetAccessControlPolicyReferenceList() []AccessControlPolicyReference`

GetAccessControlPolicyReferenceList returns the AccessControlPolicyReferenceList field if non-nil, zero value otherwise.

### GetAccessControlPolicyReferenceListOk

`func (o *UserStatusResource) GetAccessControlPolicyReferenceListOk() (*[]AccessControlPolicyReference, bool)`

GetAccessControlPolicyReferenceListOk returns a tuple with the AccessControlPolicyReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicyReferenceList

`func (o *UserStatusResource) SetAccessControlPolicyReferenceList(v []AccessControlPolicyReference)`

SetAccessControlPolicyReferenceList sets AccessControlPolicyReferenceList field to given value.

### HasAccessControlPolicyReferenceList

`func (o *UserStatusResource) HasAccessControlPolicyReferenceList() bool`

HasAccessControlPolicyReferenceList returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserStatusResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserStatusResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserStatusResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserStatusResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserType

`func (o *UserStatusResource) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserStatusResource) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserStatusResource) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserStatusResource) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetResourceUsageSummary

`func (o *UserStatusResource) GetResourceUsageSummary() UserStatusResourceResourceUsageSummary`

GetResourceUsageSummary returns the ResourceUsageSummary field if non-nil, zero value otherwise.

### GetResourceUsageSummaryOk

`func (o *UserStatusResource) GetResourceUsageSummaryOk() (*UserStatusResourceResourceUsageSummary, bool)`

GetResourceUsageSummaryOk returns a tuple with the ResourceUsageSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUsageSummary

`func (o *UserStatusResource) SetResourceUsageSummary(v UserStatusResourceResourceUsageSummary)`

SetResourceUsageSummary sets ResourceUsageSummary field to given value.

### HasResourceUsageSummary

`func (o *UserStatusResource) HasResourceUsageSummary() bool`

HasResourceUsageSummary returns a boolean if a field has been set.

### GetProjectsReferenceList

`func (o *UserStatusResource) GetProjectsReferenceList() []ProjectReference`

GetProjectsReferenceList returns the ProjectsReferenceList field if non-nil, zero value otherwise.

### GetProjectsReferenceListOk

`func (o *UserStatusResource) GetProjectsReferenceListOk() (*[]ProjectReference, bool)`

GetProjectsReferenceListOk returns a tuple with the ProjectsReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsReferenceList

`func (o *UserStatusResource) SetProjectsReferenceList(v []ProjectReference)`

SetProjectsReferenceList sets ProjectsReferenceList field to given value.

### HasProjectsReferenceList

`func (o *UserStatusResource) HasProjectsReferenceList() bool`

HasProjectsReferenceList returns a boolean if a field has been set.

### GetIdentityProviderUser

`func (o *UserStatusResource) GetIdentityProviderUser() IdentityProviderUser`

GetIdentityProviderUser returns the IdentityProviderUser field if non-nil, zero value otherwise.

### GetIdentityProviderUserOk

`func (o *UserStatusResource) GetIdentityProviderUserOk() (*IdentityProviderUser, bool)`

GetIdentityProviderUserOk returns a tuple with the IdentityProviderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderUser

`func (o *UserStatusResource) SetIdentityProviderUser(v IdentityProviderUser)`

SetIdentityProviderUser sets IdentityProviderUser field to given value.

### HasIdentityProviderUser

`func (o *UserStatusResource) HasIdentityProviderUser() bool`

HasIdentityProviderUser returns a boolean if a field has been set.

### GetDirectoryServiceUser

`func (o *UserStatusResource) GetDirectoryServiceUser() DirectoryServiceUserStatus`

GetDirectoryServiceUser returns the DirectoryServiceUser field if non-nil, zero value otherwise.

### GetDirectoryServiceUserOk

`func (o *UserStatusResource) GetDirectoryServiceUserOk() (*DirectoryServiceUserStatus, bool)`

GetDirectoryServiceUserOk returns a tuple with the DirectoryServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceUser

`func (o *UserStatusResource) SetDirectoryServiceUser(v DirectoryServiceUserStatus)`

SetDirectoryServiceUser sets DirectoryServiceUser field to given value.

### HasDirectoryServiceUser

`func (o *UserStatusResource) HasDirectoryServiceUser() bool`

HasDirectoryServiceUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


