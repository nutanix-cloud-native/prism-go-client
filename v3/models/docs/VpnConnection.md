# VpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for vpn_connection. | [optional] 
**Resources** | [**VpnConnectionResources**](VpnConnectionResources.md) |  | 
**Name** | **string** | vpn_connection Name. | 

## Methods

### NewVpnConnection

`func NewVpnConnection(resources VpnConnectionResources, name string, ) *VpnConnection`

NewVpnConnection instantiates a new VpnConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionWithDefaults

`func NewVpnConnectionWithDefaults() *VpnConnection`

NewVpnConnectionWithDefaults instantiates a new VpnConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *VpnConnection) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VpnConnection) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VpnConnection) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *VpnConnection) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *VpnConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *VpnConnection) GetResources() VpnConnectionResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpnConnection) GetResourcesOk() (*VpnConnectionResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpnConnection) SetResources(v VpnConnectionResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *VpnConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnConnection) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


