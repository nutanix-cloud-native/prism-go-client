# SyslogServerResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleList** | Pointer to [**[]SyslogModule**](SyslogModule.md) |  | [optional] 
**IpAddress** | Pointer to **string** | IP address of the Remote Syslog server | [optional] 
**Port** | Pointer to **int32** | Port of the Remote Syslog server | [optional] 
**ServerName** | Pointer to **string** | Name of the Remote Syslog server | [optional] 
**NetworkProtocol** | Pointer to **string** | Network Protocol to be used | [optional] 

## Methods

### NewSyslogServerResources

`func NewSyslogServerResources() *SyslogServerResources`

NewSyslogServerResources instantiates a new SyslogServerResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServerResourcesWithDefaults

`func NewSyslogServerResourcesWithDefaults() *SyslogServerResources`

NewSyslogServerResourcesWithDefaults instantiates a new SyslogServerResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleList

`func (o *SyslogServerResources) GetModuleList() []SyslogModule`

GetModuleList returns the ModuleList field if non-nil, zero value otherwise.

### GetModuleListOk

`func (o *SyslogServerResources) GetModuleListOk() (*[]SyslogModule, bool)`

GetModuleListOk returns a tuple with the ModuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleList

`func (o *SyslogServerResources) SetModuleList(v []SyslogModule)`

SetModuleList sets ModuleList field to given value.

### HasModuleList

`func (o *SyslogServerResources) HasModuleList() bool`

HasModuleList returns a boolean if a field has been set.

### GetIpAddress

`func (o *SyslogServerResources) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SyslogServerResources) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SyslogServerResources) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SyslogServerResources) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *SyslogServerResources) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogServerResources) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogServerResources) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogServerResources) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServerName

`func (o *SyslogServerResources) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *SyslogServerResources) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *SyslogServerResources) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *SyslogServerResources) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetNetworkProtocol

`func (o *SyslogServerResources) GetNetworkProtocol() string`

GetNetworkProtocol returns the NetworkProtocol field if non-nil, zero value otherwise.

### GetNetworkProtocolOk

`func (o *SyslogServerResources) GetNetworkProtocolOk() (*string, bool)`

GetNetworkProtocolOk returns a tuple with the NetworkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtocol

`func (o *SyslogServerResources) SetNetworkProtocol(v string)`

SetNetworkProtocol sets NetworkProtocol field to given value.

### HasNetworkProtocol

`func (o *SyslogServerResources) HasNetworkProtocol() bool`

HasNetworkProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


