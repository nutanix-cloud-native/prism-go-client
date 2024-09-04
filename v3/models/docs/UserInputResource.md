# UserInputResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderUser** | Pointer to [**IdentityProviderUser**](IdentityProviderUser.md) |  | [optional] 
**DirectoryServiceUser** | Pointer to [**DirectoryServiceUser**](DirectoryServiceUser.md) |  | [optional] 

## Methods

### NewUserInputResource

`func NewUserInputResource() *UserInputResource`

NewUserInputResource instantiates a new UserInputResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInputResourceWithDefaults

`func NewUserInputResourceWithDefaults() *UserInputResource`

NewUserInputResourceWithDefaults instantiates a new UserInputResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviderUser

`func (o *UserInputResource) GetIdentityProviderUser() IdentityProviderUser`

GetIdentityProviderUser returns the IdentityProviderUser field if non-nil, zero value otherwise.

### GetIdentityProviderUserOk

`func (o *UserInputResource) GetIdentityProviderUserOk() (*IdentityProviderUser, bool)`

GetIdentityProviderUserOk returns a tuple with the IdentityProviderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderUser

`func (o *UserInputResource) SetIdentityProviderUser(v IdentityProviderUser)`

SetIdentityProviderUser sets IdentityProviderUser field to given value.

### HasIdentityProviderUser

`func (o *UserInputResource) HasIdentityProviderUser() bool`

HasIdentityProviderUser returns a boolean if a field has been set.

### GetDirectoryServiceUser

`func (o *UserInputResource) GetDirectoryServiceUser() DirectoryServiceUser`

GetDirectoryServiceUser returns the DirectoryServiceUser field if non-nil, zero value otherwise.

### GetDirectoryServiceUserOk

`func (o *UserInputResource) GetDirectoryServiceUserOk() (*DirectoryServiceUser, bool)`

GetDirectoryServiceUserOk returns a tuple with the DirectoryServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceUser

`func (o *UserInputResource) SetDirectoryServiceUser(v DirectoryServiceUser)`

SetDirectoryServiceUser sets DirectoryServiceUser field to given value.

### HasDirectoryServiceUser

`func (o *UserInputResource) HasDirectoryServiceUser() bool`

HasDirectoryServiceUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


