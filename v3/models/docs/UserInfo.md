# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlPolicyList** | Pointer to [**[]AccessControlPolicyDetail**](AccessControlPolicyDetail.md) | The list of Access Control Policies | [optional] 
**UserDetails** | Pointer to [**[]UserDetailsObject**](UserDetailsObject.md) | User details of the logged in user | [optional] 
**UserLegacyName** | Pointer to **string** | The legacy name of the logged in user | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlPolicyList

`func (o *UserInfo) GetAccessControlPolicyList() []AccessControlPolicyDetail`

GetAccessControlPolicyList returns the AccessControlPolicyList field if non-nil, zero value otherwise.

### GetAccessControlPolicyListOk

`func (o *UserInfo) GetAccessControlPolicyListOk() (*[]AccessControlPolicyDetail, bool)`

GetAccessControlPolicyListOk returns a tuple with the AccessControlPolicyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicyList

`func (o *UserInfo) SetAccessControlPolicyList(v []AccessControlPolicyDetail)`

SetAccessControlPolicyList sets AccessControlPolicyList field to given value.

### HasAccessControlPolicyList

`func (o *UserInfo) HasAccessControlPolicyList() bool`

HasAccessControlPolicyList returns a boolean if a field has been set.

### GetUserDetails

`func (o *UserInfo) GetUserDetails() []UserDetailsObject`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *UserInfo) GetUserDetailsOk() (*[]UserDetailsObject, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *UserInfo) SetUserDetails(v []UserDetailsObject)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *UserInfo) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.

### GetUserLegacyName

`func (o *UserInfo) GetUserLegacyName() string`

GetUserLegacyName returns the UserLegacyName field if non-nil, zero value otherwise.

### GetUserLegacyNameOk

`func (o *UserInfo) GetUserLegacyNameOk() (*string, bool)`

GetUserLegacyNameOk returns a tuple with the UserLegacyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLegacyName

`func (o *UserInfo) SetUserLegacyName(v string)`

SetUserLegacyName sets UserLegacyName field to given value.

### HasUserLegacyName

`func (o *UserInfo) HasUserLegacyName() bool`

HasUserLegacyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


