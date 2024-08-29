# BgpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** | Autonomous system number | [optional] 
**Password** | Pointer to **string** | Password for authentication. | [optional] 
**DistributeConnected** | Pointer to **bool** | Boolean flag to indicate users choice on whether connected routes must be redistributed over eBGP. (Applicable to VLAN backed gateways only)  | [optional] 
**PeerIp** | Pointer to **string** | IP address of the peer. | [optional] 

## Methods

### NewBgpConfig

`func NewBgpConfig() *BgpConfig`

NewBgpConfig instantiates a new BgpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigWithDefaults

`func NewBgpConfigWithDefaults() *BgpConfig`

NewBgpConfigWithDefaults instantiates a new BgpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *BgpConfig) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpConfig) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpConfig) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *BgpConfig) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetPassword

`func (o *BgpConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BgpConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDistributeConnected

`func (o *BgpConfig) GetDistributeConnected() bool`

GetDistributeConnected returns the DistributeConnected field if non-nil, zero value otherwise.

### GetDistributeConnectedOk

`func (o *BgpConfig) GetDistributeConnectedOk() (*bool, bool)`

GetDistributeConnectedOk returns a tuple with the DistributeConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeConnected

`func (o *BgpConfig) SetDistributeConnected(v bool)`

SetDistributeConnected sets DistributeConnected field to given value.

### HasDistributeConnected

`func (o *BgpConfig) HasDistributeConnected() bool`

HasDistributeConnected returns a boolean if a field has been set.

### GetPeerIp

`func (o *BgpConfig) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *BgpConfig) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *BgpConfig) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *BgpConfig) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


