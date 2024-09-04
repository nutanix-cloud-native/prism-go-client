# UnderlaySubnetResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterReference** | Pointer to [**DatacenterReference**](DatacenterReference.md) |  | [optional] 
**Netmask** | **string** | Netmask (could be in CIDR or IP format) | 
**Gateway** | **string** | Gateway IP address | 

## Methods

### NewUnderlaySubnetResources

`func NewUnderlaySubnetResources(netmask string, gateway string, ) *UnderlaySubnetResources`

NewUnderlaySubnetResources instantiates a new UnderlaySubnetResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetResourcesWithDefaults

`func NewUnderlaySubnetResourcesWithDefaults() *UnderlaySubnetResources`

NewUnderlaySubnetResourcesWithDefaults instantiates a new UnderlaySubnetResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterReference

`func (o *UnderlaySubnetResources) GetDatacenterReference() DatacenterReference`

GetDatacenterReference returns the DatacenterReference field if non-nil, zero value otherwise.

### GetDatacenterReferenceOk

`func (o *UnderlaySubnetResources) GetDatacenterReferenceOk() (*DatacenterReference, bool)`

GetDatacenterReferenceOk returns a tuple with the DatacenterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterReference

`func (o *UnderlaySubnetResources) SetDatacenterReference(v DatacenterReference)`

SetDatacenterReference sets DatacenterReference field to given value.

### HasDatacenterReference

`func (o *UnderlaySubnetResources) HasDatacenterReference() bool`

HasDatacenterReference returns a boolean if a field has been set.

### GetNetmask

`func (o *UnderlaySubnetResources) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *UnderlaySubnetResources) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *UnderlaySubnetResources) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.


### GetGateway

`func (o *UnderlaySubnetResources) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnderlaySubnetResources) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnderlaySubnetResources) SetGateway(v string)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


