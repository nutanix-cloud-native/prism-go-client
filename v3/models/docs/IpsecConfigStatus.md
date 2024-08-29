# IpsecConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteVtiIp** | Pointer to **string** | Virtual tunnel interface IP of the remote VPN gateway. | [optional] 
**NegotiatedProposal** | Pointer to [**IpsecConfigStatusNegotiatedProposal**](IpsecConfigStatusNegotiatedProposal.md) |  | [optional] 
**RemoteAuthenticationId** | Pointer to **string** | IKE Authentication ID of the remote peer. | [optional] 
**LocalVtiIp** | Pointer to **string** | Virtual tunnel interface IP of the local VPN gateway. | [optional] 
**IkeLifetimeSecs** | Pointer to **int32** | IKE lifetime (in seconds) | [optional] 
**EspPfsDhGroupNumber** | Pointer to **int32** | Diffie-Hellman group configured for Perfect Forward Secrecy (PFS). Supported DH groups are 14, 19 and 20.  | [optional] 
**LocalAuthenticationId** | Pointer to **string** | Local IKE authentication ID used for this connection. | [optional] 
**IpsecLifetimeSecs** | Pointer to **int32** | IPSec lifetime (in seconds) | [optional] 
**VtiIpPrefixLength** | Pointer to **int32** | IP prefix length of the virtual tunnel interface. | [optional] 
**PreSharedKey** | Pointer to **string** | Shared secret for authentication between gateway peers. Note that the clear-text secret value specfied in the input spec is never revealed in the status. Use this field only as means to verify if the secret is currently set or not.  | [optional] 

## Methods

### NewIpsecConfigStatus

`func NewIpsecConfigStatus() *IpsecConfigStatus`

NewIpsecConfigStatus instantiates a new IpsecConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecConfigStatusWithDefaults

`func NewIpsecConfigStatusWithDefaults() *IpsecConfigStatus`

NewIpsecConfigStatusWithDefaults instantiates a new IpsecConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteVtiIp

`func (o *IpsecConfigStatus) GetRemoteVtiIp() string`

GetRemoteVtiIp returns the RemoteVtiIp field if non-nil, zero value otherwise.

### GetRemoteVtiIpOk

`func (o *IpsecConfigStatus) GetRemoteVtiIpOk() (*string, bool)`

GetRemoteVtiIpOk returns a tuple with the RemoteVtiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtiIp

`func (o *IpsecConfigStatus) SetRemoteVtiIp(v string)`

SetRemoteVtiIp sets RemoteVtiIp field to given value.

### HasRemoteVtiIp

`func (o *IpsecConfigStatus) HasRemoteVtiIp() bool`

HasRemoteVtiIp returns a boolean if a field has been set.

### GetNegotiatedProposal

`func (o *IpsecConfigStatus) GetNegotiatedProposal() IpsecConfigStatusNegotiatedProposal`

GetNegotiatedProposal returns the NegotiatedProposal field if non-nil, zero value otherwise.

### GetNegotiatedProposalOk

`func (o *IpsecConfigStatus) GetNegotiatedProposalOk() (*IpsecConfigStatusNegotiatedProposal, bool)`

GetNegotiatedProposalOk returns a tuple with the NegotiatedProposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiatedProposal

`func (o *IpsecConfigStatus) SetNegotiatedProposal(v IpsecConfigStatusNegotiatedProposal)`

SetNegotiatedProposal sets NegotiatedProposal field to given value.

### HasNegotiatedProposal

`func (o *IpsecConfigStatus) HasNegotiatedProposal() bool`

HasNegotiatedProposal returns a boolean if a field has been set.

### GetRemoteAuthenticationId

`func (o *IpsecConfigStatus) GetRemoteAuthenticationId() string`

GetRemoteAuthenticationId returns the RemoteAuthenticationId field if non-nil, zero value otherwise.

### GetRemoteAuthenticationIdOk

`func (o *IpsecConfigStatus) GetRemoteAuthenticationIdOk() (*string, bool)`

GetRemoteAuthenticationIdOk returns a tuple with the RemoteAuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationId

`func (o *IpsecConfigStatus) SetRemoteAuthenticationId(v string)`

SetRemoteAuthenticationId sets RemoteAuthenticationId field to given value.

### HasRemoteAuthenticationId

`func (o *IpsecConfigStatus) HasRemoteAuthenticationId() bool`

HasRemoteAuthenticationId returns a boolean if a field has been set.

### GetLocalVtiIp

`func (o *IpsecConfigStatus) GetLocalVtiIp() string`

GetLocalVtiIp returns the LocalVtiIp field if non-nil, zero value otherwise.

### GetLocalVtiIpOk

`func (o *IpsecConfigStatus) GetLocalVtiIpOk() (*string, bool)`

GetLocalVtiIpOk returns a tuple with the LocalVtiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtiIp

`func (o *IpsecConfigStatus) SetLocalVtiIp(v string)`

SetLocalVtiIp sets LocalVtiIp field to given value.

### HasLocalVtiIp

`func (o *IpsecConfigStatus) HasLocalVtiIp() bool`

HasLocalVtiIp returns a boolean if a field has been set.

### GetIkeLifetimeSecs

`func (o *IpsecConfigStatus) GetIkeLifetimeSecs() int32`

GetIkeLifetimeSecs returns the IkeLifetimeSecs field if non-nil, zero value otherwise.

### GetIkeLifetimeSecsOk

`func (o *IpsecConfigStatus) GetIkeLifetimeSecsOk() (*int32, bool)`

GetIkeLifetimeSecsOk returns a tuple with the IkeLifetimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetimeSecs

`func (o *IpsecConfigStatus) SetIkeLifetimeSecs(v int32)`

SetIkeLifetimeSecs sets IkeLifetimeSecs field to given value.

### HasIkeLifetimeSecs

`func (o *IpsecConfigStatus) HasIkeLifetimeSecs() bool`

HasIkeLifetimeSecs returns a boolean if a field has been set.

### GetEspPfsDhGroupNumber

`func (o *IpsecConfigStatus) GetEspPfsDhGroupNumber() int32`

GetEspPfsDhGroupNumber returns the EspPfsDhGroupNumber field if non-nil, zero value otherwise.

### GetEspPfsDhGroupNumberOk

`func (o *IpsecConfigStatus) GetEspPfsDhGroupNumberOk() (*int32, bool)`

GetEspPfsDhGroupNumberOk returns a tuple with the EspPfsDhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspPfsDhGroupNumber

`func (o *IpsecConfigStatus) SetEspPfsDhGroupNumber(v int32)`

SetEspPfsDhGroupNumber sets EspPfsDhGroupNumber field to given value.

### HasEspPfsDhGroupNumber

`func (o *IpsecConfigStatus) HasEspPfsDhGroupNumber() bool`

HasEspPfsDhGroupNumber returns a boolean if a field has been set.

### GetLocalAuthenticationId

`func (o *IpsecConfigStatus) GetLocalAuthenticationId() string`

GetLocalAuthenticationId returns the LocalAuthenticationId field if non-nil, zero value otherwise.

### GetLocalAuthenticationIdOk

`func (o *IpsecConfigStatus) GetLocalAuthenticationIdOk() (*string, bool)`

GetLocalAuthenticationIdOk returns a tuple with the LocalAuthenticationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAuthenticationId

`func (o *IpsecConfigStatus) SetLocalAuthenticationId(v string)`

SetLocalAuthenticationId sets LocalAuthenticationId field to given value.

### HasLocalAuthenticationId

`func (o *IpsecConfigStatus) HasLocalAuthenticationId() bool`

HasLocalAuthenticationId returns a boolean if a field has been set.

### GetIpsecLifetimeSecs

`func (o *IpsecConfigStatus) GetIpsecLifetimeSecs() int32`

GetIpsecLifetimeSecs returns the IpsecLifetimeSecs field if non-nil, zero value otherwise.

### GetIpsecLifetimeSecsOk

`func (o *IpsecConfigStatus) GetIpsecLifetimeSecsOk() (*int32, bool)`

GetIpsecLifetimeSecsOk returns a tuple with the IpsecLifetimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLifetimeSecs

`func (o *IpsecConfigStatus) SetIpsecLifetimeSecs(v int32)`

SetIpsecLifetimeSecs sets IpsecLifetimeSecs field to given value.

### HasIpsecLifetimeSecs

`func (o *IpsecConfigStatus) HasIpsecLifetimeSecs() bool`

HasIpsecLifetimeSecs returns a boolean if a field has been set.

### GetVtiIpPrefixLength

`func (o *IpsecConfigStatus) GetVtiIpPrefixLength() int32`

GetVtiIpPrefixLength returns the VtiIpPrefixLength field if non-nil, zero value otherwise.

### GetVtiIpPrefixLengthOk

`func (o *IpsecConfigStatus) GetVtiIpPrefixLengthOk() (*int32, bool)`

GetVtiIpPrefixLengthOk returns a tuple with the VtiIpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtiIpPrefixLength

`func (o *IpsecConfigStatus) SetVtiIpPrefixLength(v int32)`

SetVtiIpPrefixLength sets VtiIpPrefixLength field to given value.

### HasVtiIpPrefixLength

`func (o *IpsecConfigStatus) HasVtiIpPrefixLength() bool`

HasVtiIpPrefixLength returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *IpsecConfigStatus) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *IpsecConfigStatus) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *IpsecConfigStatus) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *IpsecConfigStatus) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


