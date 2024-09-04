# IdentityProviderDefStatusResourcesIdpProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginUrl** | **string** | Login URL of the Identity provider. | 
**Certificate** | **string** | Cert for verification. | 
**IdpUrl** | **string** | URL of the Identity provider. | 
**ErrorUrl** | Pointer to **string** | Error URL of the Identity provider. | [optional] 
**LogoutUrl** | Pointer to **string** | Logout URL of the Identity provider. | [optional] 

## Methods

### NewIdentityProviderDefStatusResourcesIdpProperties

`func NewIdentityProviderDefStatusResourcesIdpProperties(loginUrl string, certificate string, idpUrl string, ) *IdentityProviderDefStatusResourcesIdpProperties`

NewIdentityProviderDefStatusResourcesIdpProperties instantiates a new IdentityProviderDefStatusResourcesIdpProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderDefStatusResourcesIdpPropertiesWithDefaults

`func NewIdentityProviderDefStatusResourcesIdpPropertiesWithDefaults() *IdentityProviderDefStatusResourcesIdpProperties`

NewIdentityProviderDefStatusResourcesIdpPropertiesWithDefaults instantiates a new IdentityProviderDefStatusResourcesIdpProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.


### GetCertificate

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IdentityProviderDefStatusResourcesIdpProperties) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetIdpUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetIdpUrl() string`

GetIdpUrl returns the IdpUrl field if non-nil, zero value otherwise.

### GetIdpUrlOk

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetIdpUrlOk() (*string, bool)`

GetIdpUrlOk returns a tuple with the IdpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) SetIdpUrl(v string)`

SetIdpUrl sets IdpUrl field to given value.


### GetErrorUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetErrorUrl() string`

GetErrorUrl returns the ErrorUrl field if non-nil, zero value otherwise.

### GetErrorUrlOk

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetErrorUrlOk() (*string, bool)`

GetErrorUrlOk returns a tuple with the ErrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) SetErrorUrl(v string)`

SetErrorUrl sets ErrorUrl field to given value.

### HasErrorUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) HasErrorUrl() bool`

HasErrorUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *IdentityProviderDefStatusResourcesIdpProperties) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *IdentityProviderDefStatusResourcesIdpProperties) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


