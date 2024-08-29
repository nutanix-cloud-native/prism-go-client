# DirectoryServiceResources1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenLdapConfiguration** | Pointer to [**OpenLdapConfiguration**](OpenLdapConfiguration.md) |  | [optional] 
**Url** | **string** | URL of the directory. | 
**DirectoryType** | **string** | Type of the directory service. | 
**AdminUserReferenceList** | Pointer to [**[]UserReference**](UserReference.md) | The list of admin users available in the directory service.  | [optional] 
**SecondaryUrls** | Pointer to **[]string** | Secondary URLs of the directory | [optional] 
**DomainName** | **string** | The domain name of the directory service. | 
**ServiceAccount** | [**ServiceAccount**](ServiceAccount.md) |  | 
**AdminGroupReferenceList** | Pointer to [**[]UserGroupReference**](UserGroupReference.md) | List of admin user groups available in the directory service.  | [optional] 

## Methods

### NewDirectoryServiceResources1

`func NewDirectoryServiceResources1(url string, directoryType string, domainName string, serviceAccount ServiceAccount, ) *DirectoryServiceResources1`

NewDirectoryServiceResources1 instantiates a new DirectoryServiceResources1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceResources1WithDefaults

`func NewDirectoryServiceResources1WithDefaults() *DirectoryServiceResources1`

NewDirectoryServiceResources1WithDefaults instantiates a new DirectoryServiceResources1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenLdapConfiguration

`func (o *DirectoryServiceResources1) GetOpenLdapConfiguration() OpenLdapConfiguration`

GetOpenLdapConfiguration returns the OpenLdapConfiguration field if non-nil, zero value otherwise.

### GetOpenLdapConfigurationOk

`func (o *DirectoryServiceResources1) GetOpenLdapConfigurationOk() (*OpenLdapConfiguration, bool)`

GetOpenLdapConfigurationOk returns a tuple with the OpenLdapConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLdapConfiguration

`func (o *DirectoryServiceResources1) SetOpenLdapConfiguration(v OpenLdapConfiguration)`

SetOpenLdapConfiguration sets OpenLdapConfiguration field to given value.

### HasOpenLdapConfiguration

`func (o *DirectoryServiceResources1) HasOpenLdapConfiguration() bool`

HasOpenLdapConfiguration returns a boolean if a field has been set.

### GetUrl

`func (o *DirectoryServiceResources1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DirectoryServiceResources1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DirectoryServiceResources1) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDirectoryType

`func (o *DirectoryServiceResources1) GetDirectoryType() string`

GetDirectoryType returns the DirectoryType field if non-nil, zero value otherwise.

### GetDirectoryTypeOk

`func (o *DirectoryServiceResources1) GetDirectoryTypeOk() (*string, bool)`

GetDirectoryTypeOk returns a tuple with the DirectoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryType

`func (o *DirectoryServiceResources1) SetDirectoryType(v string)`

SetDirectoryType sets DirectoryType field to given value.


### GetAdminUserReferenceList

`func (o *DirectoryServiceResources1) GetAdminUserReferenceList() []UserReference`

GetAdminUserReferenceList returns the AdminUserReferenceList field if non-nil, zero value otherwise.

### GetAdminUserReferenceListOk

`func (o *DirectoryServiceResources1) GetAdminUserReferenceListOk() (*[]UserReference, bool)`

GetAdminUserReferenceListOk returns a tuple with the AdminUserReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserReferenceList

`func (o *DirectoryServiceResources1) SetAdminUserReferenceList(v []UserReference)`

SetAdminUserReferenceList sets AdminUserReferenceList field to given value.

### HasAdminUserReferenceList

`func (o *DirectoryServiceResources1) HasAdminUserReferenceList() bool`

HasAdminUserReferenceList returns a boolean if a field has been set.

### GetSecondaryUrls

`func (o *DirectoryServiceResources1) GetSecondaryUrls() []string`

GetSecondaryUrls returns the SecondaryUrls field if non-nil, zero value otherwise.

### GetSecondaryUrlsOk

`func (o *DirectoryServiceResources1) GetSecondaryUrlsOk() (*[]string, bool)`

GetSecondaryUrlsOk returns a tuple with the SecondaryUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryUrls

`func (o *DirectoryServiceResources1) SetSecondaryUrls(v []string)`

SetSecondaryUrls sets SecondaryUrls field to given value.

### HasSecondaryUrls

`func (o *DirectoryServiceResources1) HasSecondaryUrls() bool`

HasSecondaryUrls returns a boolean if a field has been set.

### GetDomainName

`func (o *DirectoryServiceResources1) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DirectoryServiceResources1) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DirectoryServiceResources1) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetServiceAccount

`func (o *DirectoryServiceResources1) GetServiceAccount() ServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DirectoryServiceResources1) GetServiceAccountOk() (*ServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DirectoryServiceResources1) SetServiceAccount(v ServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetAdminGroupReferenceList

`func (o *DirectoryServiceResources1) GetAdminGroupReferenceList() []UserGroupReference`

GetAdminGroupReferenceList returns the AdminGroupReferenceList field if non-nil, zero value otherwise.

### GetAdminGroupReferenceListOk

`func (o *DirectoryServiceResources1) GetAdminGroupReferenceListOk() (*[]UserGroupReference, bool)`

GetAdminGroupReferenceListOk returns a tuple with the AdminGroupReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroupReferenceList

`func (o *DirectoryServiceResources1) SetAdminGroupReferenceList(v []UserGroupReference)`

SetAdminGroupReferenceList sets AdminGroupReferenceList field to given value.

### HasAdminGroupReferenceList

`func (o *DirectoryServiceResources1) HasAdminGroupReferenceList() bool`

HasAdminGroupReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


