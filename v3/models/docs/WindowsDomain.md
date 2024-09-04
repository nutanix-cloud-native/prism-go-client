# WindowsDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the node to be renamed to during domain-join. If not given,a new name will be automatically assigned.  | [optional] 
**DomainName** | Pointer to **string** | Full name of domain. | [optional] 
**DomainCredential** | [**Credentials**](Credentials.md) |  | 
**OrganizationUnitPath** | Pointer to **string** | Path to organization unit in the domain. | [optional] 
**NamePrefix** | Pointer to **string** | The name prefix in the domain in case of CPS deployment. | [optional] 
**NameServerIp** | Pointer to **string** | The ip of name server that can resolve the domain name. Required during joining domain.  | [optional] 

## Methods

### NewWindowsDomain

`func NewWindowsDomain(domainCredential Credentials, ) *WindowsDomain`

NewWindowsDomain instantiates a new WindowsDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsDomainWithDefaults

`func NewWindowsDomainWithDefaults() *WindowsDomain`

NewWindowsDomainWithDefaults instantiates a new WindowsDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WindowsDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WindowsDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WindowsDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WindowsDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomainName

`func (o *WindowsDomain) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *WindowsDomain) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *WindowsDomain) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *WindowsDomain) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainCredential

`func (o *WindowsDomain) GetDomainCredential() Credentials`

GetDomainCredential returns the DomainCredential field if non-nil, zero value otherwise.

### GetDomainCredentialOk

`func (o *WindowsDomain) GetDomainCredentialOk() (*Credentials, bool)`

GetDomainCredentialOk returns a tuple with the DomainCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCredential

`func (o *WindowsDomain) SetDomainCredential(v Credentials)`

SetDomainCredential sets DomainCredential field to given value.


### GetOrganizationUnitPath

`func (o *WindowsDomain) GetOrganizationUnitPath() string`

GetOrganizationUnitPath returns the OrganizationUnitPath field if non-nil, zero value otherwise.

### GetOrganizationUnitPathOk

`func (o *WindowsDomain) GetOrganizationUnitPathOk() (*string, bool)`

GetOrganizationUnitPathOk returns a tuple with the OrganizationUnitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnitPath

`func (o *WindowsDomain) SetOrganizationUnitPath(v string)`

SetOrganizationUnitPath sets OrganizationUnitPath field to given value.

### HasOrganizationUnitPath

`func (o *WindowsDomain) HasOrganizationUnitPath() bool`

HasOrganizationUnitPath returns a boolean if a field has been set.

### GetNamePrefix

`func (o *WindowsDomain) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *WindowsDomain) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *WindowsDomain) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *WindowsDomain) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetNameServerIp

`func (o *WindowsDomain) GetNameServerIp() string`

GetNameServerIp returns the NameServerIp field if non-nil, zero value otherwise.

### GetNameServerIpOk

`func (o *WindowsDomain) GetNameServerIpOk() (*string, bool)`

GetNameServerIpOk returns a tuple with the NameServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerIp

`func (o *WindowsDomain) SetNameServerIp(v string)`

SetNameServerIp sets NameServerIp field to given value.

### HasNameServerIp

`func (o *WindowsDomain) HasNameServerIp() bool`

HasNameServerIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


