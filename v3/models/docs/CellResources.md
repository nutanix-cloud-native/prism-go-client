# CellResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterReference** | [**DatacenterReference**](DatacenterReference.md) |  | 
**CellClass** | **string** | Human readable tag that denotes the type of hardware in a cell as well as the networking equipment and setup.  | 

## Methods

### NewCellResources

`func NewCellResources(datacenterReference DatacenterReference, cellClass string, ) *CellResources`

NewCellResources instantiates a new CellResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellResourcesWithDefaults

`func NewCellResourcesWithDefaults() *CellResources`

NewCellResourcesWithDefaults instantiates a new CellResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterReference

`func (o *CellResources) GetDatacenterReference() DatacenterReference`

GetDatacenterReference returns the DatacenterReference field if non-nil, zero value otherwise.

### GetDatacenterReferenceOk

`func (o *CellResources) GetDatacenterReferenceOk() (*DatacenterReference, bool)`

GetDatacenterReferenceOk returns a tuple with the DatacenterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterReference

`func (o *CellResources) SetDatacenterReference(v DatacenterReference)`

SetDatacenterReference sets DatacenterReference field to given value.


### GetCellClass

`func (o *CellResources) GetCellClass() string`

GetCellClass returns the CellClass field if non-nil, zero value otherwise.

### GetCellClassOk

`func (o *CellResources) GetCellClassOk() (*string, bool)`

GetCellClassOk returns a tuple with the CellClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellClass

`func (o *CellResources) SetCellClass(v string)`

SetCellClass sets CellClass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


