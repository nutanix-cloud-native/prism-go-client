# InternalDirectoryServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenldapConfig** | Pointer to [**InternalOpenLdapConfig**](InternalOpenLdapConfig.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the directory. | [optional] 
**DirectoryType** | Pointer to **string** | Type of the directory service. | [optional] 
**DomainName** | Pointer to **string** | The domain name of the directory service. | [optional] 
**ServiceAccountConfig** | Pointer to [**InternalServiceAccountConfig**](InternalServiceAccountConfig.md) |  | [optional] 
**RecursiveGroupSearch** | Pointer to **bool** | Search group recursively | [optional] 
**DomainDisplayName** | Pointer to **string** | Domain display name | [optional] 
**Name** | Pointer to **string** | Name of the directory service configuration. | [optional] 

## Methods

### NewInternalDirectoryServiceConfig

`func NewInternalDirectoryServiceConfig() *InternalDirectoryServiceConfig`

NewInternalDirectoryServiceConfig instantiates a new InternalDirectoryServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDirectoryServiceConfigWithDefaults

`func NewInternalDirectoryServiceConfigWithDefaults() *InternalDirectoryServiceConfig`

NewInternalDirectoryServiceConfigWithDefaults instantiates a new InternalDirectoryServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenldapConfig

`func (o *InternalDirectoryServiceConfig) GetOpenldapConfig() InternalOpenLdapConfig`

GetOpenldapConfig returns the OpenldapConfig field if non-nil, zero value otherwise.

### GetOpenldapConfigOk

`func (o *InternalDirectoryServiceConfig) GetOpenldapConfigOk() (*InternalOpenLdapConfig, bool)`

GetOpenldapConfigOk returns a tuple with the OpenldapConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenldapConfig

`func (o *InternalDirectoryServiceConfig) SetOpenldapConfig(v InternalOpenLdapConfig)`

SetOpenldapConfig sets OpenldapConfig field to given value.

### HasOpenldapConfig

`func (o *InternalDirectoryServiceConfig) HasOpenldapConfig() bool`

HasOpenldapConfig returns a boolean if a field has been set.

### GetUUID

`func (o *InternalDirectoryServiceConfig) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *InternalDirectoryServiceConfig) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *InternalDirectoryServiceConfig) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *InternalDirectoryServiceConfig) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetUrl

`func (o *InternalDirectoryServiceConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InternalDirectoryServiceConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InternalDirectoryServiceConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InternalDirectoryServiceConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDirectoryType

`func (o *InternalDirectoryServiceConfig) GetDirectoryType() string`

GetDirectoryType returns the DirectoryType field if non-nil, zero value otherwise.

### GetDirectoryTypeOk

`func (o *InternalDirectoryServiceConfig) GetDirectoryTypeOk() (*string, bool)`

GetDirectoryTypeOk returns a tuple with the DirectoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryType

`func (o *InternalDirectoryServiceConfig) SetDirectoryType(v string)`

SetDirectoryType sets DirectoryType field to given value.

### HasDirectoryType

`func (o *InternalDirectoryServiceConfig) HasDirectoryType() bool`

HasDirectoryType returns a boolean if a field has been set.

### GetDomainName

`func (o *InternalDirectoryServiceConfig) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *InternalDirectoryServiceConfig) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *InternalDirectoryServiceConfig) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *InternalDirectoryServiceConfig) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetServiceAccountConfig

`func (o *InternalDirectoryServiceConfig) GetServiceAccountConfig() InternalServiceAccountConfig`

GetServiceAccountConfig returns the ServiceAccountConfig field if non-nil, zero value otherwise.

### GetServiceAccountConfigOk

`func (o *InternalDirectoryServiceConfig) GetServiceAccountConfigOk() (*InternalServiceAccountConfig, bool)`

GetServiceAccountConfigOk returns a tuple with the ServiceAccountConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountConfig

`func (o *InternalDirectoryServiceConfig) SetServiceAccountConfig(v InternalServiceAccountConfig)`

SetServiceAccountConfig sets ServiceAccountConfig field to given value.

### HasServiceAccountConfig

`func (o *InternalDirectoryServiceConfig) HasServiceAccountConfig() bool`

HasServiceAccountConfig returns a boolean if a field has been set.

### GetRecursiveGroupSearch

`func (o *InternalDirectoryServiceConfig) GetRecursiveGroupSearch() bool`

GetRecursiveGroupSearch returns the RecursiveGroupSearch field if non-nil, zero value otherwise.

### GetRecursiveGroupSearchOk

`func (o *InternalDirectoryServiceConfig) GetRecursiveGroupSearchOk() (*bool, bool)`

GetRecursiveGroupSearchOk returns a tuple with the RecursiveGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveGroupSearch

`func (o *InternalDirectoryServiceConfig) SetRecursiveGroupSearch(v bool)`

SetRecursiveGroupSearch sets RecursiveGroupSearch field to given value.

### HasRecursiveGroupSearch

`func (o *InternalDirectoryServiceConfig) HasRecursiveGroupSearch() bool`

HasRecursiveGroupSearch returns a boolean if a field has been set.

### GetDomainDisplayName

`func (o *InternalDirectoryServiceConfig) GetDomainDisplayName() string`

GetDomainDisplayName returns the DomainDisplayName field if non-nil, zero value otherwise.

### GetDomainDisplayNameOk

`func (o *InternalDirectoryServiceConfig) GetDomainDisplayNameOk() (*string, bool)`

GetDomainDisplayNameOk returns a tuple with the DomainDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDisplayName

`func (o *InternalDirectoryServiceConfig) SetDomainDisplayName(v string)`

SetDomainDisplayName sets DomainDisplayName field to given value.

### HasDomainDisplayName

`func (o *InternalDirectoryServiceConfig) HasDomainDisplayName() bool`

HasDomainDisplayName returns a boolean if a field has been set.

### GetName

`func (o *InternalDirectoryServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalDirectoryServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalDirectoryServiceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalDirectoryServiceConfig) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


