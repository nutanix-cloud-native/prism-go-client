# DirectConnectVirtualInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for direct_connect_virtual_interface. | [optional] 
**Resources** | [**DirectConnectVirtualInterfaceResources**](DirectConnectVirtualInterfaceResources.md) |  | 
**Name** | **string** | direct_connect_virtual_interface Name. | 

## Methods

### NewDirectConnectVirtualInterface

`func NewDirectConnectVirtualInterface(resources DirectConnectVirtualInterfaceResources, name string, ) *DirectConnectVirtualInterface`

NewDirectConnectVirtualInterface instantiates a new DirectConnectVirtualInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVirtualInterfaceWithDefaults

`func NewDirectConnectVirtualInterfaceWithDefaults() *DirectConnectVirtualInterface`

NewDirectConnectVirtualInterfaceWithDefaults instantiates a new DirectConnectVirtualInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *DirectConnectVirtualInterface) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *DirectConnectVirtualInterface) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *DirectConnectVirtualInterface) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *DirectConnectVirtualInterface) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *DirectConnectVirtualInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectConnectVirtualInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectConnectVirtualInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectConnectVirtualInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *DirectConnectVirtualInterface) GetResources() DirectConnectVirtualInterfaceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DirectConnectVirtualInterface) GetResourcesOk() (*DirectConnectVirtualInterfaceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DirectConnectVirtualInterface) SetResources(v DirectConnectVirtualInterfaceResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *DirectConnectVirtualInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectConnectVirtualInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectConnectVirtualInterface) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


