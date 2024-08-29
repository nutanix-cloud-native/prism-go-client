# VirtualNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Resources** | Pointer to [**VirtualNetworkResources**](VirtualNetworkResources.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualNetwork

`func NewVirtualNetwork(name string, ) *VirtualNetwork`

NewVirtualNetwork instantiates a new VirtualNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkWithDefaults

`func NewVirtualNetworkWithDefaults() *VirtualNetwork`

NewVirtualNetworkWithDefaults instantiates a new VirtualNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *VirtualNetwork) GetResources() VirtualNetworkResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VirtualNetwork) GetResourcesOk() (*VirtualNetworkResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VirtualNetwork) SetResources(v VirtualNetworkResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VirtualNetwork) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


