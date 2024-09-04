# DirectoryServiceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenLdapConfiguration** | Pointer to [**OpenLdapConfigurationDefStatus**](OpenLdapConfigurationDefStatus.md) |  | [optional] 
**Url** | **string** | URL of the directory. | 
**DirectoryType** | **string** | Type of the directory service. | 
**AdminUserReferenceList** | Pointer to [**[]UserReference**](UserReference.md) | The list of admin users available in the directory service.  | [optional] 
**SecondaryUrls** | Pointer to **[]string** | Secondary URLs of the directory | [optional] 
**DomainName** | **string** | The domain name of the directory service. | 
**ServiceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 
**AdminGroupReferenceList** | Pointer to [**[]UserGroupReference**](UserGroupReference.md) | List of admin user groups available in the directory service.  | [optional] 

## Methods

### NewDirectoryServiceResources

`func NewDirectoryServiceResources(url string, directoryType string, domainName string, serviceAccount ServiceAccount, ) *DirectoryServiceResources`

NewDirectoryServiceResources instantiates a new DirectoryServiceResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceResourcesWithDefaults

`func NewDirectoryServiceResourcesWithDefaults() *DirectoryServiceResources`

NewDirectoryServiceResourcesWithDefaults instantiates a new DirectoryServiceResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenLdapConfiguration

`func (o *DirectoryServiceResources) GetOpenLdapConfiguration() OpenLdapConfigurationDefStatus`

GetOpenLdapConfiguration returns the OpenLdapConfiguration field if non-nil, zero value otherwise.

### GetOpenLdapConfigurationOk

`func (o *DirectoryServiceResources) GetOpenLdapConfigurationOk() (*OpenLdapConfigurationDefStatus, bool)`

GetOpenLdapConfigurationOk returns a tuple with the OpenLdapConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLdapConfiguration

`func (o *DirectoryServiceResources) SetOpenLdapConfiguration(v OpenLdapConfigurationDefStatus)`

SetOpenLdapConfiguration sets OpenLdapConfiguration field to given value.

### HasOpenLdapConfiguration

`func (o *DirectoryServiceResources) HasOpenLdapConfiguration() bool`

HasOpenLdapConfiguration returns a boolean if a field has been set.

### GetUrl

`func (o *DirectoryServiceResources) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DirectoryServiceResources) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DirectoryServiceResources) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDirectoryType

`func (o *DirectoryServiceResources) GetDirectoryType() string`

GetDirectoryType returns the DirectoryType field if non-nil, zero value otherwise.

### GetDirectoryTypeOk

`func (o *DirectoryServiceResources) GetDirectoryTypeOk() (*string, bool)`

GetDirectoryTypeOk returns a tuple with the DirectoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryType

`func (o *DirectoryServiceResources) SetDirectoryType(v string)`

SetDirectoryType sets DirectoryType field to given value.


### GetAdminUserReferenceList

`func (o *DirectoryServiceResources) GetAdminUserReferenceList() []UserReference`

GetAdminUserReferenceList returns the AdminUserReferenceList field if non-nil, zero value otherwise.

### GetAdminUserReferenceListOk

`func (o *DirectoryServiceResources) GetAdminUserReferenceListOk() (*[]UserReference, bool)`

GetAdminUserReferenceListOk returns a tuple with the AdminUserReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserReferenceList

`func (o *DirectoryServiceResources) SetAdminUserReferenceList(v []UserReference)`

SetAdminUserReferenceList sets AdminUserReferenceList field to given value.

### HasAdminUserReferenceList

`func (o *DirectoryServiceResources) HasAdminUserReferenceList() bool`

HasAdminUserReferenceList returns a boolean if a field has been set.

### GetSecondaryUrls

`func (o *DirectoryServiceResources) GetSecondaryUrls() []string`

GetSecondaryUrls returns the SecondaryUrls field if non-nil, zero value otherwise.

### GetSecondaryUrlsOk

`func (o *DirectoryServiceResources) GetSecondaryUrlsOk() (*[]string, bool)`

GetSecondaryUrlsOk returns a tuple with the SecondaryUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryUrls

`func (o *DirectoryServiceResources) SetSecondaryUrls(v []string)`

SetSecondaryUrls sets SecondaryUrls field to given value.

### HasSecondaryUrls

`func (o *DirectoryServiceResources) HasSecondaryUrls() bool`

HasSecondaryUrls returns a boolean if a field has been set.

### GetDomainName

`func (o *DirectoryServiceResources) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DirectoryServiceResources) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DirectoryServiceResources) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetServiceAccount

`func (o *DirectoryServiceResources) GetServiceAccount() ServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DirectoryServiceResources) GetServiceAccountOk() (*ServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DirectoryServiceResources) SetServiceAccount(v ServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetAdminGroupReferenceList

`func (o *DirectoryServiceResources) GetAdminGroupReferenceList() []UserGroupReference`

GetAdminGroupReferenceList returns the AdminGroupReferenceList field if non-nil, zero value otherwise.

### GetAdminGroupReferenceListOk

`func (o *DirectoryServiceResources) GetAdminGroupReferenceListOk() (*[]UserGroupReference, bool)`

GetAdminGroupReferenceListOk returns a tuple with the AdminGroupReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupReferenceList

`func (o *DirectoryServiceResources) SetAdminGroupReferenceList(v []UserGroupReference)`

SetAdminGroupReferenceList sets AdminGroupReferenceList field to given value.

### HasAdminGroupReferenceList

`func (o *DirectoryServiceResources) HasAdminGroupReferenceList() bool`

HasAdminGroupReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


