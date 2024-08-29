# DirectoryServiceUserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserPrincipalName** | Pointer to **string** | The UserPrincipalName of the user from the directory service. It will be same as default user principal name if no upn suffix is enabled for user logon name in directory service.  | [optional] 
**DirectoryServiceReference** | Pointer to [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | [optional] 
**DefaultUserPrincipalName** | Pointer to **string** | The Default UserPrincipalName of the user from the directory service. This is of format samAccountName@domain_name.  | [optional] 

## Methods

### NewDirectoryServiceUserStatus

`func NewDirectoryServiceUserStatus() *DirectoryServiceUserStatus`

NewDirectoryServiceUserStatus instantiates a new DirectoryServiceUserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceUserStatusWithDefaults

`func NewDirectoryServiceUserStatusWithDefaults() *DirectoryServiceUserStatus`

NewDirectoryServiceUserStatusWithDefaults instantiates a new DirectoryServiceUserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserPrincipalName

`func (o *DirectoryServiceUserStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DirectoryServiceUserStatus) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *DirectoryServiceUserStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *DirectoryServiceUserStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetDirectoryServiceReference

`func (o *DirectoryServiceUserStatus) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *DirectoryServiceUserStatus) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *DirectoryServiceUserStatus) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.

### HasDirectoryServiceReference

`func (o *DirectoryServiceUserStatus) HasDirectoryServiceReference() bool`

HasDirectoryServiceReference returns a boolean if a field has been set.

### GetDefaultUserPrincipalName

`func (o *DirectoryServiceUserStatus) GetDefaultUserPrincipalName() string`

GetDefaultUserPrincipalName returns the DefaultUserPrincipalName field if non-nil, zero value otherwise.

### GetDefaultUserPrincipalNameOk

`func (o *DirectoryServiceUserStatus) GetDefaultUserPrincipalNameOk() (*string, bool)`

GetDefaultUserPrincipalNameOk returns a tuple with the DefaultUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserPrincipalName

`func (o *DirectoryServiceUserStatus) SetDefaultUserPrincipalName(v string)`

SetDefaultUserPrincipalName sets DefaultUserPrincipalName field to given value.

### HasDefaultUserPrincipalName

`func (o *DirectoryServiceUserStatus) HasDefaultUserPrincipalName() bool`

HasDefaultUserPrincipalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


