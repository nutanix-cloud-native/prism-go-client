# BgpConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** | ASN used in BGP. | [optional] 
**Password** | Pointer to **string** | Password for authentication. Note that the clear-text password value specfied in the input spec is never revealed in the status. Use this field only as means to verify if the password is currently set or not.  | [optional] 
**DistributeConnected** | Pointer to **bool** | Boolean flag to indicate users choice on whether connected routes must be redistributed over eBGP. (Applicable to VLAN backed gateways only)  | [optional] 
**PeerIp** | Pointer to **string** | IP address of the peer. | [optional] 

## Methods

### NewBgpConfigStatus

`func NewBgpConfigStatus() *BgpConfigStatus`

NewBgpConfigStatus instantiates a new BgpConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigStatusWithDefaults

`func NewBgpConfigStatusWithDefaults() *BgpConfigStatus`

NewBgpConfigStatusWithDefaults instantiates a new BgpConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *BgpConfigStatus) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *BgpConfigStatus) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *BgpConfigStatus) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *BgpConfigStatus) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetPassword

`func (o *BgpConfigStatus) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BgpConfigStatus) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BgpConfigStatus) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BgpConfigStatus) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDistributeConnected

`func (o *BgpConfigStatus) GetDistributeConnected() bool`

GetDistributeConnected returns the DistributeConnected field if non-nil, zero value otherwise.

### GetDistributeConnectedOk

`func (o *BgpConfigStatus) GetDistributeConnectedOk() (*bool, bool)`

GetDistributeConnectedOk returns a tuple with the DistributeConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeConnected

`func (o *BgpConfigStatus) SetDistributeConnected(v bool)`

SetDistributeConnected sets DistributeConnected field to given value.

### HasDistributeConnected

`func (o *BgpConfigStatus) HasDistributeConnected() bool`

HasDistributeConnected returns a boolean if a field has been set.

### GetPeerIp

`func (o *BgpConfigStatus) GetPeerIp() string`

GetPeerIp returns the PeerIp field if non-nil, zero value otherwise.

### GetPeerIpOk

`func (o *BgpConfigStatus) GetPeerIpOk() (*string, bool)`

GetPeerIpOk returns a tuple with the PeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIp

`func (o *BgpConfigStatus) SetPeerIp(v string)`

SetPeerIp sets PeerIp field to given value.

### HasPeerIp

`func (o *BgpConfigStatus) HasPeerIp() bool`

HasPeerIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


