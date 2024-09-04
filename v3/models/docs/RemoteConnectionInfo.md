# RemoteConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | public key | [optional] 
**ClusterFqdn** | Pointer to **string** | Cluster Fully Qualified Domain Name | [optional] 
**PeerRemoteConnectionUuid** | Pointer to **string** | Connection uuid for remote Peer | [optional] 
**AdditionalCapabilities** | Pointer to **[]string** |  | [optional] 
**NodeAddressList** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**Token** | Pointer to [**RemoteConnectionToken**](RemoteConnectionToken.md) |  | [optional] 
**ClusterUuid** | Pointer to **string** | UUID of cluster | [optional] 
**ClusterFunction** | Pointer to **string** | cluster function string | [optional] 

## Methods

### NewRemoteConnectionInfo

`func NewRemoteConnectionInfo() *RemoteConnectionInfo`

NewRemoteConnectionInfo instantiates a new RemoteConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionInfoWithDefaults

`func NewRemoteConnectionInfoWithDefaults() *RemoteConnectionInfo`

NewRemoteConnectionInfoWithDefaults instantiates a new RemoteConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *RemoteConnectionInfo) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RemoteConnectionInfo) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RemoteConnectionInfo) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RemoteConnectionInfo) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetClusterFqdn

`func (o *RemoteConnectionInfo) GetClusterFqdn() string`

GetClusterFqdn returns the ClusterFqdn field if non-nil, zero value otherwise.

### GetClusterFqdnOk

`func (o *RemoteConnectionInfo) GetClusterFqdnOk() (*string, bool)`

GetClusterFqdnOk returns a tuple with the ClusterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFqdn

`func (o *RemoteConnectionInfo) SetClusterFqdn(v string)`

SetClusterFqdn sets ClusterFqdn field to given value.

### HasClusterFqdn

`func (o *RemoteConnectionInfo) HasClusterFqdn() bool`

HasClusterFqdn returns a boolean if a field has been set.

### GetPeerRemoteConnectionUuid

`func (o *RemoteConnectionInfo) GetPeerRemoteConnectionUuid() string`

GetPeerRemoteConnectionUuid returns the PeerRemoteConnectionUuid field if non-nil, zero value otherwise.

### GetPeerRemoteConnectionUuidOk

`func (o *RemoteConnectionInfo) GetPeerRemoteConnectionUuidOk() (*string, bool)`

GetPeerRemoteConnectionUuidOk returns a tuple with the PeerRemoteConnectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRemoteConnectionUuid

`func (o *RemoteConnectionInfo) SetPeerRemoteConnectionUuid(v string)`

SetPeerRemoteConnectionUuid sets PeerRemoteConnectionUuid field to given value.

### HasPeerRemoteConnectionUuid

`func (o *RemoteConnectionInfo) HasPeerRemoteConnectionUuid() bool`

HasPeerRemoteConnectionUuid returns a boolean if a field has been set.

### GetAdditionalCapabilities

`func (o *RemoteConnectionInfo) GetAdditionalCapabilities() []string`

GetAdditionalCapabilities returns the AdditionalCapabilities field if non-nil, zero value otherwise.

### GetAdditionalCapabilitiesOk

`func (o *RemoteConnectionInfo) GetAdditionalCapabilitiesOk() (*[]string, bool)`

GetAdditionalCapabilitiesOk returns a tuple with the AdditionalCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCapabilities

`func (o *RemoteConnectionInfo) SetAdditionalCapabilities(v []string)`

SetAdditionalCapabilities sets AdditionalCapabilities field to given value.

### HasAdditionalCapabilities

`func (o *RemoteConnectionInfo) HasAdditionalCapabilities() bool`

HasAdditionalCapabilities returns a boolean if a field has been set.

### GetNodeAddressList

`func (o *RemoteConnectionInfo) GetNodeAddressList() []Address`

GetNodeAddressList returns the NodeAddressList field if non-nil, zero value otherwise.

### GetNodeAddressListOk

`func (o *RemoteConnectionInfo) GetNodeAddressListOk() (*[]Address, bool)`

GetNodeAddressListOk returns a tuple with the NodeAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddressList

`func (o *RemoteConnectionInfo) SetNodeAddressList(v []Address)`

SetNodeAddressList sets NodeAddressList field to given value.

### HasNodeAddressList

`func (o *RemoteConnectionInfo) HasNodeAddressList() bool`

HasNodeAddressList returns a boolean if a field has been set.

### GetToken

`func (o *RemoteConnectionInfo) GetToken() RemoteConnectionToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoteConnectionInfo) GetTokenOk() (*RemoteConnectionToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoteConnectionInfo) SetToken(v RemoteConnectionToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *RemoteConnectionInfo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetClusterUuid

`func (o *RemoteConnectionInfo) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *RemoteConnectionInfo) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *RemoteConnectionInfo) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *RemoteConnectionInfo) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetClusterFunction

`func (o *RemoteConnectionInfo) GetClusterFunction() string`

GetClusterFunction returns the ClusterFunction field if non-nil, zero value otherwise.

### GetClusterFunctionOk

`func (o *RemoteConnectionInfo) GetClusterFunctionOk() (*string, bool)`

GetClusterFunctionOk returns a tuple with the ClusterFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterFunction

`func (o *RemoteConnectionInfo) SetClusterFunction(v string)`

SetClusterFunction sets ClusterFunction field to given value.

### HasClusterFunction

`func (o *RemoteConnectionInfo) HasClusterFunction() bool`

HasClusterFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


