# RackResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellReference** | Pointer to [**CellReference**](CellReference.md) |  | [optional] 
**Location** | Pointer to **string** | The rack location | [optional] 

## Methods

### NewRackResources

`func NewRackResources() *RackResources`

NewRackResources instantiates a new RackResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackResourcesWithDefaults

`func NewRackResourcesWithDefaults() *RackResources`

NewRackResourcesWithDefaults instantiates a new RackResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellReference

`func (o *RackResources) GetCellReference() CellReference`

GetCellReference returns the CellReference field if non-nil, zero value otherwise.

### GetCellReferenceOk

`func (o *RackResources) GetCellReferenceOk() (*CellReference, bool)`

GetCellReferenceOk returns a tuple with the CellReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellReference

`func (o *RackResources) SetCellReference(v CellReference)`

SetCellReference sets CellReference field to given value.

### HasCellReference

`func (o *RackResources) HasCellReference() bool`

HasCellReference returns a boolean if a field has been set.

### GetLocation

`func (o *RackResources) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RackResources) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RackResources) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RackResources) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


