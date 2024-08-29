# NetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional name for the device | [optional] 
**Resources** | [**NetworkDeviceResources**](NetworkDeviceResources.md) |  | 
**Description** | Pointer to **string** | An optional description for the device | [optional] 

## Methods

### NewNetworkDevice

`func NewNetworkDevice(resources NetworkDeviceResources, ) *NetworkDevice`

NewNetworkDevice instantiates a new NetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceWithDefaults

`func NewNetworkDeviceWithDefaults() *NetworkDevice`

NewNetworkDeviceWithDefaults instantiates a new NetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *NetworkDevice) GetResources() NetworkDeviceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkDevice) GetResourcesOk() (*NetworkDeviceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkDevice) SetResources(v NetworkDeviceResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *NetworkDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


