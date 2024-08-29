# ProjectResources2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountReferenceList** | Pointer to [**[]AccountReference**](AccountReference.md) | List of accounts associated with the project. | [optional] 
**ResourceDomain** | Pointer to [**ResourceDomainResourcesStatus**](ResourceDomainResourcesStatus.md) |  | [optional] 
**DirectoryReferenceList** | Pointer to [**[]DirectoryReference**](DirectoryReference.md) | List of directory references in the project | [optional] 
**IdentityProvidersReferenceList** | Pointer to [**[]IdentityProvidersReference**](IdentityProvidersReference.md) | List of identity provider references in the project | [optional] 
**UserReferenceList** | Pointer to [**[]UserReference**](UserReference.md) | List of users added directly to the project.  | [optional] 
**DefaultSubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**EnvironmentReferenceList** | Pointer to [**[]EnvironmentReference**](EnvironmentReference.md) | List of environments associated with the project. | [optional] 
**IsDefault** | Pointer to **bool** | Indicates if it is the default project.  | [optional] 
**TunnelReferenceList** | Pointer to [**[]TunnelReference**](TunnelReference.md) | List of tunnels associated with the project. | [optional] 
**ExternalUserGroupReferenceList** | Pointer to [**[]UserGroupReference**](UserGroupReference.md) | List of directory service groups reference. These groups are not managed by Nutanix.  | [optional] 
**DefaultEnvironmentReference** | Pointer to [**EnvironmentReference**](EnvironmentReference.md) |  | [optional] 
**SubnetReferenceList** | Pointer to [**[]SubnetReference**](SubnetReference.md) | List of subnets for the project. | [optional] 
**ClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | List of clusters associated with the project. | [optional] 
**VpcReferenceList** | Pointer to [**[]VpcReference**](VpcReference.md) | List of VPCs associated with the project. | [optional] 
**ExternalNetworkList** | Pointer to [**[]ExternalNetwork**](ExternalNetwork.md) | List of external network associated with the project. | [optional] 
**EnableDirectoryAndIdentityProviderWhitelist** | Pointer to **bool** | Indicates AD/IDP whitelisting | [optional] 

## Methods

### NewProjectResources2

`func NewProjectResources2() *ProjectResources2`

NewProjectResources2 instantiates a new ProjectResources2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResources2WithDefaults

`func NewProjectResources2WithDefaults() *ProjectResources2`

NewProjectResources2WithDefaults instantiates a new ProjectResources2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountReferenceList

`func (o *ProjectResources2) GetAccountReferenceList() []AccountReference`

GetAccountReferenceList returns the AccountReferenceList field if non-nil, zero value otherwise.

### GetAccountReferenceListOk

`func (o *ProjectResources2) GetAccountReferenceListOk() (*[]AccountReference, bool)`

GetAccountReferenceListOk returns a tuple with the AccountReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferenceList

`func (o *ProjectResources2) SetAccountReferenceList(v []AccountReference)`

SetAccountReferenceList sets AccountReferenceList field to given value.

### HasAccountReferenceList

`func (o *ProjectResources2) HasAccountReferenceList() bool`

HasAccountReferenceList returns a boolean if a field has been set.

### GetResourceDomain

`func (o *ProjectResources2) GetResourceDomain() ResourceDomainResourcesStatus`

GetResourceDomain returns the ResourceDomain field if non-nil, zero value otherwise.

### GetResourceDomainOk

`func (o *ProjectResources2) GetResourceDomainOk() (*ResourceDomainResourcesStatus, bool)`

GetResourceDomainOk returns a tuple with the ResourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDomain

`func (o *ProjectResources2) SetResourceDomain(v ResourceDomainResourcesStatus)`

SetResourceDomain sets ResourceDomain field to given value.

### HasResourceDomain

`func (o *ProjectResources2) HasResourceDomain() bool`

HasResourceDomain returns a boolean if a field has been set.

### GetDirectoryReferenceList

`func (o *ProjectResources2) GetDirectoryReferenceList() []DirectoryReference`

GetDirectoryReferenceList returns the DirectoryReferenceList field if non-nil, zero value otherwise.

### GetDirectoryReferenceListOk

`func (o *ProjectResources2) GetDirectoryReferenceListOk() (*[]DirectoryReference, bool)`

GetDirectoryReferenceListOk returns a tuple with the DirectoryReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryReferenceList

`func (o *ProjectResources2) SetDirectoryReferenceList(v []DirectoryReference)`

SetDirectoryReferenceList sets DirectoryReferenceList field to given value.

### HasDirectoryReferenceList

`func (o *ProjectResources2) HasDirectoryReferenceList() bool`

HasDirectoryReferenceList returns a boolean if a field has been set.

### GetIdentityProvidersReferenceList

`func (o *ProjectResources2) GetIdentityProvidersReferenceList() []IdentityProvidersReference`

GetIdentityProvidersReferenceList returns the IdentityProvidersReferenceList field if non-nil, zero value otherwise.

### GetIdentityProvidersReferenceListOk

`func (o *ProjectResources2) GetIdentityProvidersReferenceListOk() (*[]IdentityProvidersReference, bool)`

GetIdentityProvidersReferenceListOk returns a tuple with the IdentityProvidersReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvidersReferenceList

`func (o *ProjectResources2) SetIdentityProvidersReferenceList(v []IdentityProvidersReference)`

SetIdentityProvidersReferenceList sets IdentityProvidersReferenceList field to given value.

### HasIdentityProvidersReferenceList

`func (o *ProjectResources2) HasIdentityProvidersReferenceList() bool`

HasIdentityProvidersReferenceList returns a boolean if a field has been set.

### GetUserReferenceList

`func (o *ProjectResources2) GetUserReferenceList() []UserReference`

GetUserReferenceList returns the UserReferenceList field if non-nil, zero value otherwise.

### GetUserReferenceListOk

`func (o *ProjectResources2) GetUserReferenceListOk() (*[]UserReference, bool)`

GetUserReferenceListOk returns a tuple with the UserReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReferenceList

`func (o *ProjectResources2) SetUserReferenceList(v []UserReference)`

SetUserReferenceList sets UserReferenceList field to given value.

### HasUserReferenceList

`func (o *ProjectResources2) HasUserReferenceList() bool`

HasUserReferenceList returns a boolean if a field has been set.

### GetDefaultSubnetReference

`func (o *ProjectResources2) GetDefaultSubnetReference() SubnetReference`

GetDefaultSubnetReference returns the DefaultSubnetReference field if non-nil, zero value otherwise.

### GetDefaultSubnetReferenceOk

`func (o *ProjectResources2) GetDefaultSubnetReferenceOk() (*SubnetReference, bool)`

GetDefaultSubnetReferenceOk returns a tuple with the DefaultSubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubnetReference

`func (o *ProjectResources2) SetDefaultSubnetReference(v SubnetReference)`

SetDefaultSubnetReference sets DefaultSubnetReference field to given value.

### HasDefaultSubnetReference

`func (o *ProjectResources2) HasDefaultSubnetReference() bool`

HasDefaultSubnetReference returns a boolean if a field has been set.

### GetEnvironmentReferenceList

`func (o *ProjectResources2) GetEnvironmentReferenceList() []EnvironmentReference`

GetEnvironmentReferenceList returns the EnvironmentReferenceList field if non-nil, zero value otherwise.

### GetEnvironmentReferenceListOk

`func (o *ProjectResources2) GetEnvironmentReferenceListOk() (*[]EnvironmentReference, bool)`

GetEnvironmentReferenceListOk returns a tuple with the EnvironmentReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentReferenceList

`func (o *ProjectResources2) SetEnvironmentReferenceList(v []EnvironmentReference)`

SetEnvironmentReferenceList sets EnvironmentReferenceList field to given value.

### HasEnvironmentReferenceList

`func (o *ProjectResources2) HasEnvironmentReferenceList() bool`

HasEnvironmentReferenceList returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProjectResources2) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProjectResources2) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProjectResources2) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProjectResources2) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetTunnelReferenceList

`func (o *ProjectResources2) GetTunnelReferenceList() []TunnelReference`

GetTunnelReferenceList returns the TunnelReferenceList field if non-nil, zero value otherwise.

### GetTunnelReferenceListOk

`func (o *ProjectResources2) GetTunnelReferenceListOk() (*[]TunnelReference, bool)`

GetTunnelReferenceListOk returns a tuple with the TunnelReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelReferenceList

`func (o *ProjectResources2) SetTunnelReferenceList(v []TunnelReference)`

SetTunnelReferenceList sets TunnelReferenceList field to given value.

### HasTunnelReferenceList

`func (o *ProjectResources2) HasTunnelReferenceList() bool`

HasTunnelReferenceList returns a boolean if a field has been set.

### GetExternalUserGroupReferenceList

`func (o *ProjectResources2) GetExternalUserGroupReferenceList() []UserGroupReference`

GetExternalUserGroupReferenceList returns the ExternalUserGroupReferenceList field if non-nil, zero value otherwise.

### GetExternalUserGroupReferenceListOk

`func (o *ProjectResources2) GetExternalUserGroupReferenceListOk() (*[]UserGroupReference, bool)`

GetExternalUserGroupReferenceListOk returns a tuple with the ExternalUserGroupReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserGroupReferenceList

`func (o *ProjectResources2) SetExternalUserGroupReferenceList(v []UserGroupReference)`

SetExternalUserGroupReferenceList sets ExternalUserGroupReferenceList field to given value.

### HasExternalUserGroupReferenceList

`func (o *ProjectResources2) HasExternalUserGroupReferenceList() bool`

HasExternalUserGroupReferenceList returns a boolean if a field has been set.

### GetDefaultEnvironmentReference

`func (o *ProjectResources2) GetDefaultEnvironmentReference() EnvironmentReference`

GetDefaultEnvironmentReference returns the DefaultEnvironmentReference field if non-nil, zero value otherwise.

### GetDefaultEnvironmentReferenceOk

`func (o *ProjectResources2) GetDefaultEnvironmentReferenceOk() (*EnvironmentReference, bool)`

GetDefaultEnvironmentReferenceOk returns a tuple with the DefaultEnvironmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEnvironmentReference

`func (o *ProjectResources2) SetDefaultEnvironmentReference(v EnvironmentReference)`

SetDefaultEnvironmentReference sets DefaultEnvironmentReference field to given value.

### HasDefaultEnvironmentReference

`func (o *ProjectResources2) HasDefaultEnvironmentReference() bool`

HasDefaultEnvironmentReference returns a boolean if a field has been set.

### GetSubnetReferenceList

`func (o *ProjectResources2) GetSubnetReferenceList() []SubnetReference`

GetSubnetReferenceList returns the SubnetReferenceList field if non-nil, zero value otherwise.

### GetSubnetReferenceListOk

`func (o *ProjectResources2) GetSubnetReferenceListOk() (*[]SubnetReference, bool)`

GetSubnetReferenceListOk returns a tuple with the SubnetReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReferenceList

`func (o *ProjectResources2) SetSubnetReferenceList(v []SubnetReference)`

SetSubnetReferenceList sets SubnetReferenceList field to given value.

### HasSubnetReferenceList

`func (o *ProjectResources2) HasSubnetReferenceList() bool`

HasSubnetReferenceList returns a boolean if a field has been set.

### GetClusterReferenceList

`func (o *ProjectResources2) GetClusterReferenceList() []ClusterReference`

GetClusterReferenceList returns the ClusterReferenceList field if non-nil, zero value otherwise.

### GetClusterReferenceListOk

`func (o *ProjectResources2) GetClusterReferenceListOk() (*[]ClusterReference, bool)`

GetClusterReferenceListOk returns a tuple with the ClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReferenceList

`func (o *ProjectResources2) SetClusterReferenceList(v []ClusterReference)`

SetClusterReferenceList sets ClusterReferenceList field to given value.

### HasClusterReferenceList

`func (o *ProjectResources2) HasClusterReferenceList() bool`

HasClusterReferenceList returns a boolean if a field has been set.

### GetVpcReferenceList

`func (o *ProjectResources2) GetVpcReferenceList() []VpcReference`

GetVpcReferenceList returns the VpcReferenceList field if non-nil, zero value otherwise.

### GetVpcReferenceListOk

`func (o *ProjectResources2) GetVpcReferenceListOk() (*[]VpcReference, bool)`

GetVpcReferenceListOk returns a tuple with the VpcReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReferenceList

`func (o *ProjectResources2) SetVpcReferenceList(v []VpcReference)`

SetVpcReferenceList sets VpcReferenceList field to given value.

### HasVpcReferenceList

`func (o *ProjectResources2) HasVpcReferenceList() bool`

HasVpcReferenceList returns a boolean if a field has been set.

### GetExternalNetworkList

`func (o *ProjectResources2) GetExternalNetworkList() []ExternalNetwork`

GetExternalNetworkList returns the ExternalNetworkList field if non-nil, zero value otherwise.

### GetExternalNetworkListOk

`func (o *ProjectResources2) GetExternalNetworkListOk() (*[]ExternalNetwork, bool)`

GetExternalNetworkListOk returns a tuple with the ExternalNetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetworkList

`func (o *ProjectResources2) SetExternalNetworkList(v []ExternalNetwork)`

SetExternalNetworkList sets ExternalNetworkList field to given value.

### HasExternalNetworkList

`func (o *ProjectResources2) HasExternalNetworkList() bool`

HasExternalNetworkList returns a boolean if a field has been set.

### GetEnableDirectoryAndIdentityProviderWhitelist

`func (o *ProjectResources2) GetEnableDirectoryAndIdentityProviderWhitelist() bool`

GetEnableDirectoryAndIdentityProviderWhitelist returns the EnableDirectoryAndIdentityProviderWhitelist field if non-nil, zero value otherwise.

### GetEnableDirectoryAndIdentityProviderWhitelistOk

`func (o *ProjectResources2) GetEnableDirectoryAndIdentityProviderWhitelistOk() (*bool, bool)`

GetEnableDirectoryAndIdentityProviderWhitelistOk returns a tuple with the EnableDirectoryAndIdentityProviderWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectoryAndIdentityProviderWhitelist

`func (o *ProjectResources2) SetEnableDirectoryAndIdentityProviderWhitelist(v bool)`

SetEnableDirectoryAndIdentityProviderWhitelist sets EnableDirectoryAndIdentityProviderWhitelist field to given value.

### HasEnableDirectoryAndIdentityProviderWhitelist

`func (o *ProjectResources2) HasEnableDirectoryAndIdentityProviderWhitelist() bool`

HasEnableDirectoryAndIdentityProviderWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


