# Layer2StretchVpnConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerConnectionReference** | Pointer to [**VpnConnectionReference**](VpnConnectionReference.md) |  | [optional] 
**LocalVtiIp** | Pointer to **string** |  | [optional] 
**ConnectionReference** | Pointer to [**VpnConnectionReference**](VpnConnectionReference.md) |  | [optional] 

## Methods

### NewLayer2StretchVpnConnectionInfo

`func NewLayer2StretchVpnConnectionInfo() *Layer2StretchVpnConnectionInfo`

NewLayer2StretchVpnConnectionInfo instantiates a new Layer2StretchVpnConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchVpnConnectionInfoWithDefaults

`func NewLayer2StretchVpnConnectionInfoWithDefaults() *Layer2StretchVpnConnectionInfo`

NewLayer2StretchVpnConnectionInfoWithDefaults instantiates a new Layer2StretchVpnConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) GetPeerConnectionReference() VpnConnectionReference`

GetPeerConnectionReference returns the PeerConnectionReference field if non-nil, zero value otherwise.

### GetPeerConnectionReferenceOk

`func (o *Layer2StretchVpnConnectionInfo) GetPeerConnectionReferenceOk() (*VpnConnectionReference, bool)`

GetPeerConnectionReferenceOk returns a tuple with the PeerConnectionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) SetPeerConnectionReference(v VpnConnectionReference)`

SetPeerConnectionReference sets PeerConnectionReference field to given value.

### HasPeerConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) HasPeerConnectionReference() bool`

HasPeerConnectionReference returns a boolean if a field has been set.

### GetLocalVtiIp

`func (o *Layer2StretchVpnConnectionInfo) GetLocalVtiIp() string`

GetLocalVtiIp returns the LocalVtiIp field if non-nil, zero value otherwise.

### GetLocalVtiIpOk

`func (o *Layer2StretchVpnConnectionInfo) GetLocalVtiIpOk() (*string, bool)`

GetLocalVtiIpOk returns a tuple with the LocalVtiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtiIp

`func (o *Layer2StretchVpnConnectionInfo) SetLocalVtiIp(v string)`

SetLocalVtiIp sets LocalVtiIp field to given value.

### HasLocalVtiIp

`func (o *Layer2StretchVpnConnectionInfo) HasLocalVtiIp() bool`

HasLocalVtiIp returns a boolean if a field has been set.

### GetConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) GetConnectionReference() VpnConnectionReference`

GetConnectionReference returns the ConnectionReference field if non-nil, zero value otherwise.

### GetConnectionReferenceOk

`func (o *Layer2StretchVpnConnectionInfo) GetConnectionReferenceOk() (*VpnConnectionReference, bool)`

GetConnectionReferenceOk returns a tuple with the ConnectionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) SetConnectionReference(v VpnConnectionReference)`

SetConnectionReference sets ConnectionReference field to given value.

### HasConnectionReference

`func (o *Layer2StretchVpnConnectionInfo) HasConnectionReference() bool`

HasConnectionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


