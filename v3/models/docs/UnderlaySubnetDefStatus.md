# UnderlaySubnetDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The underlay subnet name | [optional] 
**Resources** | [**UnderlaySubnetResources**](UnderlaySubnetResources.md) |  | 

## Methods

### NewUnderlaySubnetDefStatus

`func NewUnderlaySubnetDefStatus(resources UnderlaySubnetResources, ) *UnderlaySubnetDefStatus`

NewUnderlaySubnetDefStatus instantiates a new UnderlaySubnetDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetDefStatusWithDefaults

`func NewUnderlaySubnetDefStatusWithDefaults() *UnderlaySubnetDefStatus`

NewUnderlaySubnetDefStatusWithDefaults instantiates a new UnderlaySubnetDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UnderlaySubnetDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnderlaySubnetDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnderlaySubnetDefStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnderlaySubnetDefStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *UnderlaySubnetDefStatus) GetResources() UnderlaySubnetResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UnderlaySubnetDefStatus) GetResourcesOk() (*UnderlaySubnetResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UnderlaySubnetDefStatus) SetResources(v UnderlaySubnetResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


