# CellDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the cell | 
**Resources** | [**CellResources**](CellResources.md) |  | 

## Methods

### NewCellDefStatus

`func NewCellDefStatus(name string, resources CellResources, ) *CellDefStatus`

NewCellDefStatus instantiates a new CellDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellDefStatusWithDefaults

`func NewCellDefStatusWithDefaults() *CellDefStatus`

NewCellDefStatusWithDefaults instantiates a new CellDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CellDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CellDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CellDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *CellDefStatus) GetResources() CellResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CellDefStatus) GetResourcesOk() (*CellResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CellDefStatus) SetResources(v CellResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


