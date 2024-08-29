# OpenLdapConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserConfiguration** | [**OpenLdapConfigurationUserConfiguration**](OpenLdapConfigurationUserConfiguration.md) |  | 
**UserGroupConfiguration** | [**OpenLdapConfigurationUserGroupConfiguration**](OpenLdapConfigurationUserGroupConfiguration.md) |  | 

## Methods

### NewOpenLdapConfiguration

`func NewOpenLdapConfiguration(userConfiguration OpenLdapConfigurationUserConfiguration, userGroupConfiguration OpenLdapConfigurationUserGroupConfiguration, ) *OpenLdapConfiguration`

NewOpenLdapConfiguration instantiates a new OpenLdapConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLdapConfigurationWithDefaults

`func NewOpenLdapConfigurationWithDefaults() *OpenLdapConfiguration`

NewOpenLdapConfigurationWithDefaults instantiates a new OpenLdapConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserConfiguration

`func (o *OpenLdapConfiguration) GetUserConfiguration() OpenLdapConfigurationUserConfiguration`

GetUserConfiguration returns the UserConfiguration field if non-nil, zero value otherwise.

### GetUserConfigurationOk

`func (o *OpenLdapConfiguration) GetUserConfigurationOk() (*OpenLdapConfigurationUserConfiguration, bool)`

GetUserConfigurationOk returns a tuple with the UserConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConfiguration

`func (o *OpenLdapConfiguration) SetUserConfiguration(v OpenLdapConfigurationUserConfiguration)`

SetUserConfiguration sets UserConfiguration field to given value.


### GetUserGroupConfiguration

`func (o *OpenLdapConfiguration) GetUserGroupConfiguration() OpenLdapConfigurationUserGroupConfiguration`

GetUserGroupConfiguration returns the UserGroupConfiguration field if non-nil, zero value otherwise.

### GetUserGroupConfigurationOk

`func (o *OpenLdapConfiguration) GetUserGroupConfigurationOk() (*OpenLdapConfigurationUserGroupConfiguration, bool)`

GetUserGroupConfigurationOk returns a tuple with the UserGroupConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupConfiguration

`func (o *OpenLdapConfiguration) SetUserGroupConfiguration(v OpenLdapConfigurationUserGroupConfiguration)`

SetUserGroupConfiguration sets UserGroupConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


