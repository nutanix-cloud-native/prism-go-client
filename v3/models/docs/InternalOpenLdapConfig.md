# InternalOpenLdapConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupObjectClass** | Pointer to **string** | The object class in the OpenLDAP system that corresponds to groups. | [optional] 
**GroupSearchBase** | Pointer to **string** | The base DN for group search. | [optional] 
**GroupMemberAttribute** | Pointer to **string** | The attribute in a group that associates users to the group. | [optional] 
**UserObjectClass** | Pointer to **string** | The object class in the OpenLDAP system that corresponds to users. | [optional] 
**UsernameAttribute** | Pointer to **string** | Unique identifier for each user which can be used in authentication. | [optional] 
**UserSearchBase** | Pointer to **string** | The base DN for user search. | [optional] 
**GroupMemberAttributeValue** | Pointer to **string** | The user attribute value that will be used in group entity to associate user to the group.  | [optional] 

## Methods

### NewInternalOpenLdapConfig

`func NewInternalOpenLdapConfig() *InternalOpenLdapConfig`

NewInternalOpenLdapConfig instantiates a new InternalOpenLdapConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalOpenLdapConfigWithDefaults

`func NewInternalOpenLdapConfigWithDefaults() *InternalOpenLdapConfig`

NewInternalOpenLdapConfigWithDefaults instantiates a new InternalOpenLdapConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupObjectClass

`func (o *InternalOpenLdapConfig) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *InternalOpenLdapConfig) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *InternalOpenLdapConfig) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *InternalOpenLdapConfig) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetGroupSearchBase

`func (o *InternalOpenLdapConfig) GetGroupSearchBase() string`

GetGroupSearchBase returns the GroupSearchBase field if non-nil, zero value otherwise.

### GetGroupSearchBaseOk

`func (o *InternalOpenLdapConfig) GetGroupSearchBaseOk() (*string, bool)`

GetGroupSearchBaseOk returns a tuple with the GroupSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBase

`func (o *InternalOpenLdapConfig) SetGroupSearchBase(v string)`

SetGroupSearchBase sets GroupSearchBase field to given value.

### HasGroupSearchBase

`func (o *InternalOpenLdapConfig) HasGroupSearchBase() bool`

HasGroupSearchBase returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *InternalOpenLdapConfig) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *InternalOpenLdapConfig) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *InternalOpenLdapConfig) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *InternalOpenLdapConfig) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *InternalOpenLdapConfig) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *InternalOpenLdapConfig) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *InternalOpenLdapConfig) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *InternalOpenLdapConfig) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.

### GetUsernameAttribute

`func (o *InternalOpenLdapConfig) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *InternalOpenLdapConfig) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *InternalOpenLdapConfig) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.

### HasUsernameAttribute

`func (o *InternalOpenLdapConfig) HasUsernameAttribute() bool`

HasUsernameAttribute returns a boolean if a field has been set.

### GetUserSearchBase

`func (o *InternalOpenLdapConfig) GetUserSearchBase() string`

GetUserSearchBase returns the UserSearchBase field if non-nil, zero value otherwise.

### GetUserSearchBaseOk

`func (o *InternalOpenLdapConfig) GetUserSearchBaseOk() (*string, bool)`

GetUserSearchBaseOk returns a tuple with the UserSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchBase

`func (o *InternalOpenLdapConfig) SetUserSearchBase(v string)`

SetUserSearchBase sets UserSearchBase field to given value.

### HasUserSearchBase

`func (o *InternalOpenLdapConfig) HasUserSearchBase() bool`

HasUserSearchBase returns a boolean if a field has been set.

### GetGroupMemberAttributeValue

`func (o *InternalOpenLdapConfig) GetGroupMemberAttributeValue() string`

GetGroupMemberAttributeValue returns the GroupMemberAttributeValue field if non-nil, zero value otherwise.

### GetGroupMemberAttributeValueOk

`func (o *InternalOpenLdapConfig) GetGroupMemberAttributeValueOk() (*string, bool)`

GetGroupMemberAttributeValueOk returns a tuple with the GroupMemberAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttributeValue

`func (o *InternalOpenLdapConfig) SetGroupMemberAttributeValue(v string)`

SetGroupMemberAttributeValue sets GroupMemberAttributeValue field to given value.

### HasGroupMemberAttributeValue

`func (o *InternalOpenLdapConfig) HasGroupMemberAttributeValue() bool`

HasGroupMemberAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


