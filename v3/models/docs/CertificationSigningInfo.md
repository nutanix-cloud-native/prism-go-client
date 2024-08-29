# CertificationSigningInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The Town or City where customer&#39;s business is located. | [optional] 
**CommonNameSuffix** | Pointer to **string** | Common name is by default &lt;node_uuid&gt;.nutanix.com, but if a customer wants something instead of nutanix.com they can specify it here.  | [optional] 
**State** | Pointer to **string** | The Province, Region, County or State where customer business is is located.  | [optional] 
**CountryCode** | Pointer to **string** | Two-letter ISO code for Country where customer&#39;s organization is located.  | [optional] 
**CommonName** | Pointer to **string** | Common name of the organization or host server | [optional] 
**Organization** | Pointer to **string** | Name of the customer business. | [optional] 
**EmailAddress** | Pointer to **string** | Email address of the certificate administrator. | [optional] 

## Methods

### NewCertificationSigningInfo

`func NewCertificationSigningInfo() *CertificationSigningInfo`

NewCertificationSigningInfo instantiates a new CertificationSigningInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationSigningInfoWithDefaults

`func NewCertificationSigningInfoWithDefaults() *CertificationSigningInfo`

NewCertificationSigningInfoWithDefaults instantiates a new CertificationSigningInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CertificationSigningInfo) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CertificationSigningInfo) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CertificationSigningInfo) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CertificationSigningInfo) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCommonNameSuffix

`func (o *CertificationSigningInfo) GetCommonNameSuffix() string`

GetCommonNameSuffix returns the CommonNameSuffix field if non-nil, zero value otherwise.

### GetCommonNameSuffixOk

`func (o *CertificationSigningInfo) GetCommonNameSuffixOk() (*string, bool)`

GetCommonNameSuffixOk returns a tuple with the CommonNameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonNameSuffix

`func (o *CertificationSigningInfo) SetCommonNameSuffix(v string)`

SetCommonNameSuffix sets CommonNameSuffix field to given value.

### HasCommonNameSuffix

`func (o *CertificationSigningInfo) HasCommonNameSuffix() bool`

HasCommonNameSuffix returns a boolean if a field has been set.

### GetState

`func (o *CertificationSigningInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificationSigningInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificationSigningInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CertificationSigningInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountryCode

`func (o *CertificationSigningInfo) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CertificationSigningInfo) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CertificationSigningInfo) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CertificationSigningInfo) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCommonName

`func (o *CertificationSigningInfo) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificationSigningInfo) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificationSigningInfo) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificationSigningInfo) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificationSigningInfo) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificationSigningInfo) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificationSigningInfo) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificationSigningInfo) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetEmailAddress

`func (o *CertificationSigningInfo) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CertificationSigningInfo) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CertificationSigningInfo) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CertificationSigningInfo) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


