# NetworkDeviceDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | An optional name for the device | [optional] 
**Resources** | [**NetworkDeviceResources**](NetworkDeviceResources.md) |  | 
**Description** | Pointer to **string** | An optional description for the device | [optional] 

## Methods

### NewNetworkDeviceDefStatus

`func NewNetworkDeviceDefStatus(resources NetworkDeviceResources, ) *NetworkDeviceDefStatus`

NewNetworkDeviceDefStatus instantiates a new NetworkDeviceDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDefStatusWithDefaults

`func NewNetworkDeviceDefStatusWithDefaults() *NetworkDeviceDefStatus`

NewNetworkDeviceDefStatusWithDefaults instantiates a new NetworkDeviceDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkDeviceDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDeviceDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDeviceDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDeviceDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *NetworkDeviceDefStatus) GetResources() NetworkDeviceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *NetworkDeviceDefStatus) GetResourcesOk() (*NetworkDeviceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *NetworkDeviceDefStatus) SetResources(v NetworkDeviceResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *NetworkDeviceDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDeviceDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDeviceDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkDeviceDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


