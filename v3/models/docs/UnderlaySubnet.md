# UnderlaySubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The underlay subnet name | [optional] 
**Resources** | [**UnderlaySubnetResources**](UnderlaySubnetResources.md) |  | 

## Methods

### NewUnderlaySubnet

`func NewUnderlaySubnet(resources UnderlaySubnetResources, ) *UnderlaySubnet`

NewUnderlaySubnet instantiates a new UnderlaySubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetWithDefaults

`func NewUnderlaySubnetWithDefaults() *UnderlaySubnet`

NewUnderlaySubnetWithDefaults instantiates a new UnderlaySubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UnderlaySubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnderlaySubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnderlaySubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnderlaySubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *UnderlaySubnet) GetResources() UnderlaySubnetResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UnderlaySubnet) GetResourcesOk() (*UnderlaySubnetResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UnderlaySubnet) SetResources(v UnderlaySubnetResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


