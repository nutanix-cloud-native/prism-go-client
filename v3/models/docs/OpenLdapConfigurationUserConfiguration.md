# OpenLdapConfigurationUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSearchBase** | **string** | The base DN for user search. | 
**UserObjectClass** | **string** | The object class in the OpenLDAP system that corresponds to users.  | 
**UsernameAttribute** | **string** | Unique identifier for each user which can be used in authentication.  | 

## Methods

### NewOpenLdapConfigurationUserConfiguration

`func NewOpenLdapConfigurationUserConfiguration(userSearchBase string, userObjectClass string, usernameAttribute string, ) *OpenLdapConfigurationUserConfiguration`

NewOpenLdapConfigurationUserConfiguration instantiates a new OpenLdapConfigurationUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLdapConfigurationUserConfigurationWithDefaults

`func NewOpenLdapConfigurationUserConfigurationWithDefaults() *OpenLdapConfigurationUserConfiguration`

NewOpenLdapConfigurationUserConfigurationWithDefaults instantiates a new OpenLdapConfigurationUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSearchBase

`func (o *OpenLdapConfigurationUserConfiguration) GetUserSearchBase() string`

GetUserSearchBase returns the UserSearchBase field if non-nil, zero value otherwise.

### GetUserSearchBaseOk

`func (o *OpenLdapConfigurationUserConfiguration) GetUserSearchBaseOk() (*string, bool)`

GetUserSearchBaseOk returns a tuple with the UserSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchBase

`func (o *OpenLdapConfigurationUserConfiguration) SetUserSearchBase(v string)`

SetUserSearchBase sets UserSearchBase field to given value.


### GetUserObjectClass

`func (o *OpenLdapConfigurationUserConfiguration) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *OpenLdapConfigurationUserConfiguration) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *OpenLdapConfigurationUserConfiguration) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.


### GetUsernameAttribute

`func (o *OpenLdapConfigurationUserConfiguration) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *OpenLdapConfigurationUserConfiguration) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *OpenLdapConfigurationUserConfiguration) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


