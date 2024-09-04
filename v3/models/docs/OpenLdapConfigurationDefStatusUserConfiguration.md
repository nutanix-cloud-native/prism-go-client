# OpenLdapConfigurationDefStatusUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSearchBase** | **string** | The base DN for user search. | 
**UserObjectClass** | **string** | The object class in the OpenLDAP system that corresponds to users.  | 
**UsernameAttribute** | **string** | Unique identifier for each user which can be used in authentication.  | 

## Methods

### NewOpenLdapConfigurationDefStatusUserConfiguration

`func NewOpenLdapConfigurationDefStatusUserConfiguration(userSearchBase string, userObjectClass string, usernameAttribute string, ) *OpenLdapConfigurationDefStatusUserConfiguration`

NewOpenLdapConfigurationDefStatusUserConfiguration instantiates a new OpenLdapConfigurationDefStatusUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLdapConfigurationDefStatusUserConfigurationWithDefaults

`func NewOpenLdapConfigurationDefStatusUserConfigurationWithDefaults() *OpenLdapConfigurationDefStatusUserConfiguration`

NewOpenLdapConfigurationDefStatusUserConfigurationWithDefaults instantiates a new OpenLdapConfigurationDefStatusUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSearchBase

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUserSearchBase() string`

GetUserSearchBase returns the UserSearchBase field if non-nil, zero value otherwise.

### GetUserSearchBaseOk

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUserSearchBaseOk() (*string, bool)`

GetUserSearchBaseOk returns a tuple with the UserSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSearchBase

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) SetUserSearchBase(v string)`

SetUserSearchBase sets UserSearchBase field to given value.


### GetUserObjectClass

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.


### GetUsernameAttribute

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *OpenLdapConfigurationDefStatusUserConfiguration) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


