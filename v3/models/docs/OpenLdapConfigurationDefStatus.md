# OpenLdapConfigurationDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserConfiguration** | [**OpenLdapConfigurationDefStatusUserConfiguration**](OpenLdapConfigurationDefStatusUserConfiguration.md) |  | 
**UserGroupConfiguration** | [**OpenLdapConfigurationDefStatusUserGroupConfiguration**](OpenLdapConfigurationDefStatusUserGroupConfiguration.md) |  | 

## Methods

### NewOpenLdapConfigurationDefStatus

`func NewOpenLdapConfigurationDefStatus(userConfiguration OpenLdapConfigurationDefStatusUserConfiguration, userGroupConfiguration OpenLdapConfigurationDefStatusUserGroupConfiguration, ) *OpenLdapConfigurationDefStatus`

NewOpenLdapConfigurationDefStatus instantiates a new OpenLdapConfigurationDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLdapConfigurationDefStatusWithDefaults

`func NewOpenLdapConfigurationDefStatusWithDefaults() *OpenLdapConfigurationDefStatus`

NewOpenLdapConfigurationDefStatusWithDefaults instantiates a new OpenLdapConfigurationDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserConfiguration

`func (o *OpenLdapConfigurationDefStatus) GetUserConfiguration() OpenLdapConfigurationDefStatusUserConfiguration`

GetUserConfiguration returns the UserConfiguration field if non-nil, zero value otherwise.

### GetUserConfigurationOk

`func (o *OpenLdapConfigurationDefStatus) GetUserConfigurationOk() (*OpenLdapConfigurationDefStatusUserConfiguration, bool)`

GetUserConfigurationOk returns a tuple with the UserConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConfiguration

`func (o *OpenLdapConfigurationDefStatus) SetUserConfiguration(v OpenLdapConfigurationDefStatusUserConfiguration)`

SetUserConfiguration sets UserConfiguration field to given value.


### GetUserGroupConfiguration

`func (o *OpenLdapConfigurationDefStatus) GetUserGroupConfiguration() OpenLdapConfigurationDefStatusUserGroupConfiguration`

GetUserGroupConfiguration returns the UserGroupConfiguration field if non-nil, zero value otherwise.

### GetUserGroupConfigurationOk

`func (o *OpenLdapConfigurationDefStatus) GetUserGroupConfigurationOk() (*OpenLdapConfigurationDefStatusUserGroupConfiguration, bool)`

GetUserGroupConfigurationOk returns a tuple with the UserGroupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupConfiguration

`func (o *OpenLdapConfigurationDefStatus) SetUserGroupConfiguration(v OpenLdapConfigurationDefStatusUserGroupConfiguration)`

SetUserGroupConfiguration sets UserGroupConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


