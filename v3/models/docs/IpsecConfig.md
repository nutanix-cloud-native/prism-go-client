# IpsecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteVtiIp** | Pointer to **string** | Virtual tunnel interface IP of the remote VPN gateway. | [optional] 
**RemoteAuthenticationId** | Pointer to **string** | IKE Authentication ID of the remote peer. | [optional] 
**LocalVtiIp** | Pointer to **string** | Virtual tunnel interface IP of the local VPN gateway. | [optional] 
**IkeLifetimeSecs** | Pointer to **int32** | IKE lifetime (in seconds) | [optional] 
**EspPfsDhGroupNumber** | Pointer to **int32** | Diffie-Hellman group to be used for Perfect Forward Secrecy (PFS). Supported DH groups are 14, 19 and 20.  | [optional] 
**LocalAuthenticationId** | Pointer to **string** | Local IKE authentication ID used for this connection. | [optional] 
**IpsecLifetimeSecs** | Pointer to **int32** | IPSec lifetime (in seconds) | [optional] 
**VtiIpPrefixLength** | Pointer to **int32** | IP prefix length of the virtual tunnel interface. | [optional] 
**PreSharedKey** | **string** | Shared secret for authentication between gateway peers. | 

## Methods

### NewIpsecConfig

`func NewIpsecConfig(preSharedKey string, ) *IpsecConfig`

NewIpsecConfig instantiates a new IpsecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecConfigWithDefaults

`func NewIpsecConfigWithDefaults() *IpsecConfig`

NewIpsecConfigWithDefaults instantiates a new IpsecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteVtiIp

`func (o *IpsecConfig) GetRemoteVtiIp() string`

GetRemoteVtiIp returns the RemoteVtiIp field if non-nil, zero value otherwise.

### GetRemoteVtiIpOk

`func (o *IpsecConfig) GetRemoteVtiIpOk() (*string, bool)`

GetRemoteVtiIpOk returns a tuple with the RemoteVtiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtiIp

`func (o *IpsecConfig) SetRemoteVtiIp(v string)`

SetRemoteVtiIp sets RemoteVtiIp field to given value.

### HasRemoteVtiIp

`func (o *IpsecConfig) HasRemoteVtiIp() bool`

HasRemoteVtiIp returns a boolean if a field has been set.

### GetRemoteAuthenticationId

`func (o *IpsecConfig) GetRemoteAuthenticationId() string`

GetRemoteAuthenticationId returns the RemoteAuthenticationId field if non-nil, zero value otherwise.

### GetRemoteAuthenticationIdOk

`func (o *IpsecConfig) GetRemoteAuthenticationIdOk() (*string, bool)`

GetRemoteAuthenticationIdOk returns a tuple with the RemoteAuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationId

`func (o *IpsecConfig) SetRemoteAuthenticationId(v string)`

SetRemoteAuthenticationId sets RemoteAuthenticationId field to given value.

### HasRemoteAuthenticationId

`func (o *IpsecConfig) HasRemoteAuthenticationId() bool`

HasRemoteAuthenticationId returns a boolean if a field has been set.

### GetLocalVtiIp

`func (o *IpsecConfig) GetLocalVtiIp() string`

GetLocalVtiIp returns the LocalVtiIp field if non-nil, zero value otherwise.

### GetLocalVtiIpOk

`func (o *IpsecConfig) GetLocalVtiIpOk() (*string, bool)`

GetLocalVtiIpOk returns a tuple with the LocalVtiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtiIp

`func (o *IpsecConfig) SetLocalVtiIp(v string)`

SetLocalVtiIp sets LocalVtiIp field to given value.

### HasLocalVtiIp

`func (o *IpsecConfig) HasLocalVtiIp() bool`

HasLocalVtiIp returns a boolean if a field has been set.

### GetIkeLifetimeSecs

`func (o *IpsecConfig) GetIkeLifetimeSecs() int32`

GetIkeLifetimeSecs returns the IkeLifetimeSecs field if non-nil, zero value otherwise.

### GetIkeLifetimeSecsOk

`func (o *IpsecConfig) GetIkeLifetimeSecsOk() (*int32, bool)`

GetIkeLifetimeSecsOk returns a tuple with the IkeLifetimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetimeSecs

`func (o *IpsecConfig) SetIkeLifetimeSecs(v int32)`

SetIkeLifetimeSecs sets IkeLifetimeSecs field to given value.

### HasIkeLifetimeSecs

`func (o *IpsecConfig) HasIkeLifetimeSecs() bool`

HasIkeLifetimeSecs returns a boolean if a field has been set.

### GetEspPfsDhGroupNumber

`func (o *IpsecConfig) GetEspPfsDhGroupNumber() int32`

GetEspPfsDhGroupNumber returns the EspPfsDhGroupNumber field if non-nil, zero value otherwise.

### GetEspPfsDhGroupNumberOk

`func (o *IpsecConfig) GetEspPfsDhGroupNumberOk() (*int32, bool)`

GetEspPfsDhGroupNumberOk returns a tuple with the EspPfsDhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspPfsDhGroupNumber

`func (o *IpsecConfig) SetEspPfsDhGroupNumber(v int32)`

SetEspPfsDhGroupNumber sets EspPfsDhGroupNumber field to given value.

### HasEspPfsDhGroupNumber

`func (o *IpsecConfig) HasEspPfsDhGroupNumber() bool`

HasEspPfsDhGroupNumber returns a boolean if a field has been set.

### GetLocalAuthenticationId

`func (o *IpsecConfig) GetLocalAuthenticationId() string`

GetLocalAuthenticationId returns the LocalAuthenticationId field if non-nil, zero value otherwise.

### GetLocalAuthenticationIdOk

`func (o *IpsecConfig) GetLocalAuthenticationIdOk() (*string, bool)`

GetLocalAuthenticationIdOk returns a tuple with the LocalAuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuthenticationId

`func (o *IpsecConfig) SetLocalAuthenticationId(v string)`

SetLocalAuthenticationId sets LocalAuthenticationId field to given value.

### HasLocalAuthenticationId

`func (o *IpsecConfig) HasLocalAuthenticationId() bool`

HasLocalAuthenticationId returns a boolean if a field has been set.

### GetIpsecLifetimeSecs

`func (o *IpsecConfig) GetIpsecLifetimeSecs() int32`

GetIpsecLifetimeSecs returns the IpsecLifetimeSecs field if non-nil, zero value otherwise.

### GetIpsecLifetimeSecsOk

`func (o *IpsecConfig) GetIpsecLifetimeSecsOk() (*int32, bool)`

GetIpsecLifetimeSecsOk returns a tuple with the IpsecLifetimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLifetimeSecs

`func (o *IpsecConfig) SetIpsecLifetimeSecs(v int32)`

SetIpsecLifetimeSecs sets IpsecLifetimeSecs field to given value.

### HasIpsecLifetimeSecs

`func (o *IpsecConfig) HasIpsecLifetimeSecs() bool`

HasIpsecLifetimeSecs returns a boolean if a field has been set.

### GetVtiIpPrefixLength

`func (o *IpsecConfig) GetVtiIpPrefixLength() int32`

GetVtiIpPrefixLength returns the VtiIpPrefixLength field if non-nil, zero value otherwise.

### GetVtiIpPrefixLengthOk

`func (o *IpsecConfig) GetVtiIpPrefixLengthOk() (*int32, bool)`

GetVtiIpPrefixLengthOk returns a tuple with the VtiIpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtiIpPrefixLength

`func (o *IpsecConfig) SetVtiIpPrefixLength(v int32)`

SetVtiIpPrefixLength sets VtiIpPrefixLength field to given value.

### HasVtiIpPrefixLength

`func (o *IpsecConfig) HasVtiIpPrefixLength() bool`

HasVtiIpPrefixLength returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *IpsecConfig) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *IpsecConfig) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *IpsecConfig) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


