# UserDetailsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | username of the logged in user. | [optional] 
**UserUuid** | Pointer to **string** | UUID of the logged in user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the logged in user. | [optional] 
**UserType** | Pointer to **string** | Type of the logged in user. | [optional] 

## Methods

### NewUserDetailsObject

`func NewUserDetailsObject() *UserDetailsObject`

NewUserDetailsObject instantiates a new UserDetailsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailsObjectWithDefaults

`func NewUserDetailsObjectWithDefaults() *UserDetailsObject`

NewUserDetailsObjectWithDefaults instantiates a new UserDetailsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserDetailsObject) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDetailsObject) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDetailsObject) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserDetailsObject) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUserUuid

`func (o *UserDetailsObject) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *UserDetailsObject) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *UserDetailsObject) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *UserDetailsObject) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDetailsObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDetailsObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDetailsObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserDetailsObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserType

`func (o *UserDetailsObject) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserDetailsObject) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserDetailsObject) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserDetailsObject) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


