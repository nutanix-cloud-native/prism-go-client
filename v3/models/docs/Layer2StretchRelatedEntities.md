# Layer2StretchRelatedEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | Pointer to [**[]Layer2StretchSubnetInfo**](Layer2StretchSubnetInfo.md) |  | [optional] 
**VpnConnections** | Pointer to [**[]Layer2StretchVpnConnectionInfo**](Layer2StretchVpnConnectionInfo.md) |  | [optional] 

## Methods

### NewLayer2StretchRelatedEntities

`func NewLayer2StretchRelatedEntities() *Layer2StretchRelatedEntities`

NewLayer2StretchRelatedEntities instantiates a new Layer2StretchRelatedEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchRelatedEntitiesWithDefaults

`func NewLayer2StretchRelatedEntitiesWithDefaults() *Layer2StretchRelatedEntities`

NewLayer2StretchRelatedEntitiesWithDefaults instantiates a new Layer2StretchRelatedEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *Layer2StretchRelatedEntities) GetSubnets() []Layer2StretchSubnetInfo`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *Layer2StretchRelatedEntities) GetSubnetsOk() (*[]Layer2StretchSubnetInfo, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *Layer2StretchRelatedEntities) SetSubnets(v []Layer2StretchSubnetInfo)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *Layer2StretchRelatedEntities) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetVpnConnections

`func (o *Layer2StretchRelatedEntities) GetVpnConnections() []Layer2StretchVpnConnectionInfo`

GetVpnConnections returns the VpnConnections field if non-nil, zero value otherwise.

### GetVpnConnectionsOk

`func (o *Layer2StretchRelatedEntities) GetVpnConnectionsOk() (*[]Layer2StretchVpnConnectionInfo, bool)`

GetVpnConnectionsOk returns a tuple with the VpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnConnections

`func (o *Layer2StretchRelatedEntities) SetVpnConnections(v []Layer2StretchVpnConnectionInfo)`

SetVpnConnections sets VpnConnections field to given value.

### HasVpnConnections

`func (o *Layer2StretchRelatedEntities) HasVpnConnections() bool`

HasVpnConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


