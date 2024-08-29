# OpenLdapConfigurationDefStatusUserGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupSearchBase** | **string** | The base DN for group search.  | 
**GroupMemberAttribute** | **string** | The attribute in a group that associates users to the group.  | 
**GroupObjectClass** | **string** | The object class in the OpenLDAP system that corresponds to groups.  | 
**GroupMemberAttributeValue** | **string** | The user attribute value that will be used in group entity to associate user to the group.  | 

## Methods

### NewOpenLdapConfigurationDefStatusUserGroupConfiguration

`func NewOpenLdapConfigurationDefStatusUserGroupConfiguration(groupSearchBase string, groupMemberAttribute string, groupObjectClass string, groupMemberAttributeValue string, ) *OpenLdapConfigurationDefStatusUserGroupConfiguration`

NewOpenLdapConfigurationDefStatusUserGroupConfiguration instantiates a new OpenLdapConfigurationDefStatusUserGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLdapConfigurationDefStatusUserGroupConfigurationWithDefaults

`func NewOpenLdapConfigurationDefStatusUserGroupConfigurationWithDefaults() *OpenLdapConfigurationDefStatusUserGroupConfiguration`

NewOpenLdapConfigurationDefStatusUserGroupConfigurationWithDefaults instantiates a new OpenLdapConfigurationDefStatusUserGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupSearchBase

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupSearchBase() string`

GetGroupSearchBase returns the GroupSearchBase field if non-nil, zero value otherwise.

### GetGroupSearchBaseOk

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupSearchBaseOk() (*string, bool)`

GetGroupSearchBaseOk returns a tuple with the GroupSearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBase

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) SetGroupSearchBase(v string)`

SetGroupSearchBase sets GroupSearchBase field to given value.


### GetGroupMemberAttribute

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.


### GetGroupObjectClass

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.


### GetGroupMemberAttributeValue

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupMemberAttributeValue() string`

GetGroupMemberAttributeValue returns the GroupMemberAttributeValue field if non-nil, zero value otherwise.

### GetGroupMemberAttributeValueOk

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) GetGroupMemberAttributeValueOk() (*string, bool)`

GetGroupMemberAttributeValueOk returns a tuple with the GroupMemberAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttributeValue

`func (o *OpenLdapConfigurationDefStatusUserGroupConfiguration) SetGroupMemberAttributeValue(v string)`

SetGroupMemberAttributeValue sets GroupMemberAttributeValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


