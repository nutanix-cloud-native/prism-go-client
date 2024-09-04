# ClusterNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasqueradingIp** | Pointer to **string** | The cluster NAT&#39;d or proxy IP which maps to the cluster local IP.  | [optional] 
**MasqueradingPort** | Pointer to **int32** | Port used together with masquerading_ip to connect to the cluster.  | [optional] 
**ExternalIp** | Pointer to **string** | The local IP of cluster visible externally. | [optional] 
**HttpProxyList** | Pointer to [**[]ClusterNetworkEntity**](ClusterNetworkEntity.md) | List of proxies to connect to the service centers. | [optional] 
**SmtpServer** | Pointer to [**SmtpServer**](SmtpServer.md) |  | [optional] 
**NtpServerIpList** | Pointer to **[]string** | The list of IP addresses or FQDNs of the NTP servers. | [optional] 
**ExternalSubnet** | Pointer to **string** | External subnet for cross server communication. The format is IP/netmask.  | [optional] [default to "172.16.0.0/255.240.0.0"]
**NfsSubnetWhitelist** | Pointer to **[]string** | Comma separated list of subnets (of the form &#39;a.b.c.d/l.m.n.o&#39;) that are allowed to send NFS requests to this container. If not specified, the global NFS whitelist will be looked up for access permission. The internal subnet is always automatically considered part of the whitelist, even if the field below does not explicitly specify it. Similarly, all the hypervisor IPs are considered part of the whitelist. Finally, to permit debugging, all of the SVMs local IPs are considered to be implicitly part of the whitelist.  | [optional] 
**ExternalDataServicesIp** | Pointer to **string** | The cluster IP address that provides external entities access to various cluster data services.  | [optional] 
**DomainServer** | Pointer to [**ClusterDomainServer**](ClusterDomainServer.md) |  | [optional] 
**FullyQualifiedDomainName** | Pointer to **string** | fully qualified domain name of the cluster visible externally. | [optional] 
**NameServerIpList** | Pointer to **[]string** | The list of IP addresses of the name servers. | [optional] 
**HttpProxyWhitelist** | Pointer to [**[]HttpProxyWhitelist**](HttpProxyWhitelist.md) | HTTP proxy whitelist. | [optional] 
**InternalSubnet** | Pointer to **string** | The internal subnet is local to every server - its not visible outside.iSCSI requests generated internally within the appliance (by user VMs or VMFS) are sent to the internal subnet. The format is IP/netmask.  | [optional] [default to "192.168.5.0/255.255.255.0"]
**DefaultVswitchConfig** | Pointer to [**VswitchConfig**](VswitchConfig.md) |  | [optional] 

## Methods

### NewClusterNetwork

`func NewClusterNetwork() *ClusterNetwork`

NewClusterNetwork instantiates a new ClusterNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkWithDefaults

`func NewClusterNetworkWithDefaults() *ClusterNetwork`

NewClusterNetworkWithDefaults instantiates a new ClusterNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasqueradingIp

`func (o *ClusterNetwork) GetMasqueradingIp() string`

GetMasqueradingIp returns the MasqueradingIp field if non-nil, zero value otherwise.

### GetMasqueradingIpOk

`func (o *ClusterNetwork) GetMasqueradingIpOk() (*string, bool)`

GetMasqueradingIpOk returns a tuple with the MasqueradingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasqueradingIp

`func (o *ClusterNetwork) SetMasqueradingIp(v string)`

SetMasqueradingIp sets MasqueradingIp field to given value.

### HasMasqueradingIp

`func (o *ClusterNetwork) HasMasqueradingIp() bool`

HasMasqueradingIp returns a boolean if a field has been set.

### GetMasqueradingPort

`func (o *ClusterNetwork) GetMasqueradingPort() int32`

GetMasqueradingPort returns the MasqueradingPort field if non-nil, zero value otherwise.

### GetMasqueradingPortOk

`func (o *ClusterNetwork) GetMasqueradingPortOk() (*int32, bool)`

GetMasqueradingPortOk returns a tuple with the MasqueradingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasqueradingPort

`func (o *ClusterNetwork) SetMasqueradingPort(v int32)`

SetMasqueradingPort sets MasqueradingPort field to given value.

### HasMasqueradingPort

`func (o *ClusterNetwork) HasMasqueradingPort() bool`

HasMasqueradingPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ClusterNetwork) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ClusterNetwork) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ClusterNetwork) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ClusterNetwork) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetHttpProxyList

`func (o *ClusterNetwork) GetHttpProxyList() []ClusterNetworkEntity`

GetHttpProxyList returns the HttpProxyList field if non-nil, zero value otherwise.

### GetHttpProxyListOk

`func (o *ClusterNetwork) GetHttpProxyListOk() (*[]ClusterNetworkEntity, bool)`

GetHttpProxyListOk returns a tuple with the HttpProxyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyList

`func (o *ClusterNetwork) SetHttpProxyList(v []ClusterNetworkEntity)`

SetHttpProxyList sets HttpProxyList field to given value.

### HasHttpProxyList

`func (o *ClusterNetwork) HasHttpProxyList() bool`

HasHttpProxyList returns a boolean if a field has been set.

### GetSmtpServer

`func (o *ClusterNetwork) GetSmtpServer() SmtpServer`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *ClusterNetwork) GetSmtpServerOk() (*SmtpServer, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *ClusterNetwork) SetSmtpServer(v SmtpServer)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *ClusterNetwork) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetNtpServerIpList

`func (o *ClusterNetwork) GetNtpServerIpList() []string`

GetNtpServerIpList returns the NtpServerIpList field if non-nil, zero value otherwise.

### GetNtpServerIpListOk

`func (o *ClusterNetwork) GetNtpServerIpListOk() (*[]string, bool)`

GetNtpServerIpListOk returns a tuple with the NtpServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServerIpList

`func (o *ClusterNetwork) SetNtpServerIpList(v []string)`

SetNtpServerIpList sets NtpServerIpList field to given value.

### HasNtpServerIpList

`func (o *ClusterNetwork) HasNtpServerIpList() bool`

HasNtpServerIpList returns a boolean if a field has been set.

### GetExternalSubnet

`func (o *ClusterNetwork) GetExternalSubnet() string`

GetExternalSubnet returns the ExternalSubnet field if non-nil, zero value otherwise.

### GetExternalSubnetOk

`func (o *ClusterNetwork) GetExternalSubnetOk() (*string, bool)`

GetExternalSubnetOk returns a tuple with the ExternalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubnet

`func (o *ClusterNetwork) SetExternalSubnet(v string)`

SetExternalSubnet sets ExternalSubnet field to given value.

### HasExternalSubnet

`func (o *ClusterNetwork) HasExternalSubnet() bool`

HasExternalSubnet returns a boolean if a field has been set.

### GetNfsSubnetWhitelist

`func (o *ClusterNetwork) GetNfsSubnetWhitelist() []string`

GetNfsSubnetWhitelist returns the NfsSubnetWhitelist field if non-nil, zero value otherwise.

### GetNfsSubnetWhitelistOk

`func (o *ClusterNetwork) GetNfsSubnetWhitelistOk() (*[]string, bool)`

GetNfsSubnetWhitelistOk returns a tuple with the NfsSubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsSubnetWhitelist

`func (o *ClusterNetwork) SetNfsSubnetWhitelist(v []string)`

SetNfsSubnetWhitelist sets NfsSubnetWhitelist field to given value.

### HasNfsSubnetWhitelist

`func (o *ClusterNetwork) HasNfsSubnetWhitelist() bool`

HasNfsSubnetWhitelist returns a boolean if a field has been set.

### GetExternalDataServicesIp

`func (o *ClusterNetwork) GetExternalDataServicesIp() string`

GetExternalDataServicesIp returns the ExternalDataServicesIp field if non-nil, zero value otherwise.

### GetExternalDataServicesIpOk

`func (o *ClusterNetwork) GetExternalDataServicesIpOk() (*string, bool)`

GetExternalDataServicesIpOk returns a tuple with the ExternalDataServicesIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDataServicesIp

`func (o *ClusterNetwork) SetExternalDataServicesIp(v string)`

SetExternalDataServicesIp sets ExternalDataServicesIp field to given value.

### HasExternalDataServicesIp

`func (o *ClusterNetwork) HasExternalDataServicesIp() bool`

HasExternalDataServicesIp returns a boolean if a field has been set.

### GetDomainServer

`func (o *ClusterNetwork) GetDomainServer() ClusterDomainServer`

GetDomainServer returns the DomainServer field if non-nil, zero value otherwise.

### GetDomainServerOk

`func (o *ClusterNetwork) GetDomainServerOk() (*ClusterDomainServer, bool)`

GetDomainServerOk returns a tuple with the DomainServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainServer

`func (o *ClusterNetwork) SetDomainServer(v ClusterDomainServer)`

SetDomainServer sets DomainServer field to given value.

### HasDomainServer

`func (o *ClusterNetwork) HasDomainServer() bool`

HasDomainServer returns a boolean if a field has been set.

### GetFullyQualifiedDomainName

`func (o *ClusterNetwork) GetFullyQualifiedDomainName() string`

GetFullyQualifiedDomainName returns the FullyQualifiedDomainName field if non-nil, zero value otherwise.

### GetFullyQualifiedDomainNameOk

`func (o *ClusterNetwork) GetFullyQualifiedDomainNameOk() (*string, bool)`

GetFullyQualifiedDomainNameOk returns a tuple with the FullyQualifiedDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedDomainName

`func (o *ClusterNetwork) SetFullyQualifiedDomainName(v string)`

SetFullyQualifiedDomainName sets FullyQualifiedDomainName field to given value.

### HasFullyQualifiedDomainName

`func (o *ClusterNetwork) HasFullyQualifiedDomainName() bool`

HasFullyQualifiedDomainName returns a boolean if a field has been set.

### GetNameServerIpList

`func (o *ClusterNetwork) GetNameServerIpList() []string`

GetNameServerIpList returns the NameServerIpList field if non-nil, zero value otherwise.

### GetNameServerIpListOk

`func (o *ClusterNetwork) GetNameServerIpListOk() (*[]string, bool)`

GetNameServerIpListOk returns a tuple with the NameServerIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServerIpList

`func (o *ClusterNetwork) SetNameServerIpList(v []string)`

SetNameServerIpList sets NameServerIpList field to given value.

### HasNameServerIpList

`func (o *ClusterNetwork) HasNameServerIpList() bool`

HasNameServerIpList returns a boolean if a field has been set.

### GetHttpProxyWhitelist

`func (o *ClusterNetwork) GetHttpProxyWhitelist() []HttpProxyWhitelist`

GetHttpProxyWhitelist returns the HttpProxyWhitelist field if non-nil, zero value otherwise.

### GetHttpProxyWhitelistOk

`func (o *ClusterNetwork) GetHttpProxyWhitelistOk() (*[]HttpProxyWhitelist, bool)`

GetHttpProxyWhitelistOk returns a tuple with the HttpProxyWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyWhitelist

`func (o *ClusterNetwork) SetHttpProxyWhitelist(v []HttpProxyWhitelist)`

SetHttpProxyWhitelist sets HttpProxyWhitelist field to given value.

### HasHttpProxyWhitelist

`func (o *ClusterNetwork) HasHttpProxyWhitelist() bool`

HasHttpProxyWhitelist returns a boolean if a field has been set.

### GetInternalSubnet

`func (o *ClusterNetwork) GetInternalSubnet() string`

GetInternalSubnet returns the InternalSubnet field if non-nil, zero value otherwise.

### GetInternalSubnetOk

`func (o *ClusterNetwork) GetInternalSubnetOk() (*string, bool)`

GetInternalSubnetOk returns a tuple with the InternalSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSubnet

`func (o *ClusterNetwork) SetInternalSubnet(v string)`

SetInternalSubnet sets InternalSubnet field to given value.

### HasInternalSubnet

`func (o *ClusterNetwork) HasInternalSubnet() bool`

HasInternalSubnet returns a boolean if a field has been set.

### GetDefaultVswitchConfig

`func (o *ClusterNetwork) GetDefaultVswitchConfig() VswitchConfig`

GetDefaultVswitchConfig returns the DefaultVswitchConfig field if non-nil, zero value otherwise.

### GetDefaultVswitchConfigOk

`func (o *ClusterNetwork) GetDefaultVswitchConfigOk() (*VswitchConfig, bool)`

GetDefaultVswitchConfigOk returns a tuple with the DefaultVswitchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVswitchConfig

`func (o *ClusterNetwork) SetDefaultVswitchConfig(v VswitchConfig)`

SetDefaultVswitchConfig sets DefaultVswitchConfig field to given value.

### HasDefaultVswitchConfig

`func (o *ClusterNetwork) HasDefaultVswitchConfig() bool`

HasDefaultVswitchConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


