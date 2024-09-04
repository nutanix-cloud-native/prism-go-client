# PhysicalAvailabilityZoneResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellFqdnList** | Pointer to **[]string** | List of cell FQDN mapped to the cluster. | [optional] 
**OlbVirtualAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**XlbVirtualAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BillingDomain** | Pointer to **string** | Billing Domain Address. | [optional] 
**IsMultiTenant** | Pointer to **bool** | Flag indicates if the Prism Central can be shared by multiple tenants. Default is False, this must be set to True for XI multitenant PC.  | [optional] [default to false]
**IdpServiceProviderName** | Pointer to **string** | IDP Service Provider name. | [optional] 
**PulseConfiguration** | Pointer to [**PulseConfiguration**](PulseConfiguration.md) |  | [optional] 
**NtpServerList** | Pointer to **[]string** | List of NTP Server addresses. | [optional] 
**CloudUuid** | Pointer to **string** | Reference to the Infra Cloud this physical AZ belongs. | [optional] 
**MyNutanixDomain** | Pointer to **string** | My Nutanix Domain Address used for IDP Registration. | [optional] 
**ExternalUrl** | Pointer to **string** | External URL for the Physical Availability Zone, required when is_mutlti_tenant set to False.  | [optional] 

## Methods

### NewPhysicalAvailabilityZoneResources

`func NewPhysicalAvailabilityZoneResources() *PhysicalAvailabilityZoneResources`

NewPhysicalAvailabilityZoneResources instantiates a new PhysicalAvailabilityZoneResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalAvailabilityZoneResourcesWithDefaults

`func NewPhysicalAvailabilityZoneResourcesWithDefaults() *PhysicalAvailabilityZoneResources`

NewPhysicalAvailabilityZoneResourcesWithDefaults instantiates a new PhysicalAvailabilityZoneResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellFqdnList

`func (o *PhysicalAvailabilityZoneResources) GetCellFqdnList() []string`

GetCellFqdnList returns the CellFqdnList field if non-nil, zero value otherwise.

### GetCellFqdnListOk

`func (o *PhysicalAvailabilityZoneResources) GetCellFqdnListOk() (*[]string, bool)`

GetCellFqdnListOk returns a tuple with the CellFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellFqdnList

`func (o *PhysicalAvailabilityZoneResources) SetCellFqdnList(v []string)`

SetCellFqdnList sets CellFqdnList field to given value.

### HasCellFqdnList

`func (o *PhysicalAvailabilityZoneResources) HasCellFqdnList() bool`

HasCellFqdnList returns a boolean if a field has been set.

### GetOlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) GetOlbVirtualAddress() Address`

GetOlbVirtualAddress returns the OlbVirtualAddress field if non-nil, zero value otherwise.

### GetOlbVirtualAddressOk

`func (o *PhysicalAvailabilityZoneResources) GetOlbVirtualAddressOk() (*Address, bool)`

GetOlbVirtualAddressOk returns a tuple with the OlbVirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) SetOlbVirtualAddress(v Address)`

SetOlbVirtualAddress sets OlbVirtualAddress field to given value.

### HasOlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) HasOlbVirtualAddress() bool`

HasOlbVirtualAddress returns a boolean if a field has been set.

### GetXlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) GetXlbVirtualAddress() Address`

GetXlbVirtualAddress returns the XlbVirtualAddress field if non-nil, zero value otherwise.

### GetXlbVirtualAddressOk

`func (o *PhysicalAvailabilityZoneResources) GetXlbVirtualAddressOk() (*Address, bool)`

GetXlbVirtualAddressOk returns a tuple with the XlbVirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) SetXlbVirtualAddress(v Address)`

SetXlbVirtualAddress sets XlbVirtualAddress field to given value.

### HasXlbVirtualAddress

`func (o *PhysicalAvailabilityZoneResources) HasXlbVirtualAddress() bool`

HasXlbVirtualAddress returns a boolean if a field has been set.

### GetBillingDomain

`func (o *PhysicalAvailabilityZoneResources) GetBillingDomain() string`

GetBillingDomain returns the BillingDomain field if non-nil, zero value otherwise.

### GetBillingDomainOk

`func (o *PhysicalAvailabilityZoneResources) GetBillingDomainOk() (*string, bool)`

GetBillingDomainOk returns a tuple with the BillingDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDomain

`func (o *PhysicalAvailabilityZoneResources) SetBillingDomain(v string)`

SetBillingDomain sets BillingDomain field to given value.

### HasBillingDomain

`func (o *PhysicalAvailabilityZoneResources) HasBillingDomain() bool`

HasBillingDomain returns a boolean if a field has been set.

### GetIsMultiTenant

`func (o *PhysicalAvailabilityZoneResources) GetIsMultiTenant() bool`

GetIsMultiTenant returns the IsMultiTenant field if non-nil, zero value otherwise.

### GetIsMultiTenantOk

`func (o *PhysicalAvailabilityZoneResources) GetIsMultiTenantOk() (*bool, bool)`

GetIsMultiTenantOk returns a tuple with the IsMultiTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiTenant

`func (o *PhysicalAvailabilityZoneResources) SetIsMultiTenant(v bool)`

SetIsMultiTenant sets IsMultiTenant field to given value.

### HasIsMultiTenant

`func (o *PhysicalAvailabilityZoneResources) HasIsMultiTenant() bool`

HasIsMultiTenant returns a boolean if a field has been set.

### GetIdpServiceProviderName

`func (o *PhysicalAvailabilityZoneResources) GetIdpServiceProviderName() string`

GetIdpServiceProviderName returns the IdpServiceProviderName field if non-nil, zero value otherwise.

### GetIdpServiceProviderNameOk

`func (o *PhysicalAvailabilityZoneResources) GetIdpServiceProviderNameOk() (*string, bool)`

GetIdpServiceProviderNameOk returns a tuple with the IdpServiceProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpServiceProviderName

`func (o *PhysicalAvailabilityZoneResources) SetIdpServiceProviderName(v string)`

SetIdpServiceProviderName sets IdpServiceProviderName field to given value.

### HasIdpServiceProviderName

`func (o *PhysicalAvailabilityZoneResources) HasIdpServiceProviderName() bool`

HasIdpServiceProviderName returns a boolean if a field has been set.

### GetPulseConfiguration

`func (o *PhysicalAvailabilityZoneResources) GetPulseConfiguration() PulseConfiguration`

GetPulseConfiguration returns the PulseConfiguration field if non-nil, zero value otherwise.

### GetPulseConfigurationOk

`func (o *PhysicalAvailabilityZoneResources) GetPulseConfigurationOk() (*PulseConfiguration, bool)`

GetPulseConfigurationOk returns a tuple with the PulseConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulseConfiguration

`func (o *PhysicalAvailabilityZoneResources) SetPulseConfiguration(v PulseConfiguration)`

SetPulseConfiguration sets PulseConfiguration field to given value.

### HasPulseConfiguration

`func (o *PhysicalAvailabilityZoneResources) HasPulseConfiguration() bool`

HasPulseConfiguration returns a boolean if a field has been set.

### GetNtpServerList

`func (o *PhysicalAvailabilityZoneResources) GetNtpServerList() []string`

GetNtpServerList returns the NtpServerList field if non-nil, zero value otherwise.

### GetNtpServerListOk

`func (o *PhysicalAvailabilityZoneResources) GetNtpServerListOk() (*[]string, bool)`

GetNtpServerListOk returns a tuple with the NtpServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerList

`func (o *PhysicalAvailabilityZoneResources) SetNtpServerList(v []string)`

SetNtpServerList sets NtpServerList field to given value.

### HasNtpServerList

`func (o *PhysicalAvailabilityZoneResources) HasNtpServerList() bool`

HasNtpServerList returns a boolean if a field has been set.

### GetCloudUuid

`func (o *PhysicalAvailabilityZoneResources) GetCloudUuid() string`

GetCloudUuid returns the CloudUuid field if non-nil, zero value otherwise.

### GetCloudUuidOk

`func (o *PhysicalAvailabilityZoneResources) GetCloudUuidOk() (*string, bool)`

GetCloudUuidOk returns a tuple with the CloudUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUuid

`func (o *PhysicalAvailabilityZoneResources) SetCloudUuid(v string)`

SetCloudUuid sets CloudUuid field to given value.

### HasCloudUuid

`func (o *PhysicalAvailabilityZoneResources) HasCloudUuid() bool`

HasCloudUuid returns a boolean if a field has been set.

### GetMyNutanixDomain

`func (o *PhysicalAvailabilityZoneResources) GetMyNutanixDomain() string`

GetMyNutanixDomain returns the MyNutanixDomain field if non-nil, zero value otherwise.

### GetMyNutanixDomainOk

`func (o *PhysicalAvailabilityZoneResources) GetMyNutanixDomainOk() (*string, bool)`

GetMyNutanixDomainOk returns a tuple with the MyNutanixDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyNutanixDomain

`func (o *PhysicalAvailabilityZoneResources) SetMyNutanixDomain(v string)`

SetMyNutanixDomain sets MyNutanixDomain field to given value.

### HasMyNutanixDomain

`func (o *PhysicalAvailabilityZoneResources) HasMyNutanixDomain() bool`

HasMyNutanixDomain returns a boolean if a field has been set.

### GetExternalUrl

`func (o *PhysicalAvailabilityZoneResources) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *PhysicalAvailabilityZoneResources) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *PhysicalAvailabilityZoneResources) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *PhysicalAvailabilityZoneResources) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


