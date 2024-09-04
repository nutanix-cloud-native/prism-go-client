# ResourceUtilizationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | **string** | The units of the resource type | 
**Limit** | Pointer to **int32** | The resource consumption limit (unspecified is unlimited) | [optional] 
**Value** | **int32** | The amount of resource consumed | 
**ResourceType** | **string** | The type of resource (for example storage, CPUs) | 

## Methods

### NewResourceUtilizationStatus

`func NewResourceUtilizationStatus(units string, value int32, resourceType string, ) *ResourceUtilizationStatus`

NewResourceUtilizationStatus instantiates a new ResourceUtilizationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUtilizationStatusWithDefaults

`func NewResourceUtilizationStatusWithDefaults() *ResourceUtilizationStatus`

NewResourceUtilizationStatusWithDefaults instantiates a new ResourceUtilizationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *ResourceUtilizationStatus) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ResourceUtilizationStatus) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ResourceUtilizationStatus) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetLimit

`func (o *ResourceUtilizationStatus) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResourceUtilizationStatus) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResourceUtilizationStatus) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResourceUtilizationStatus) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetValue

`func (o *ResourceUtilizationStatus) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceUtilizationStatus) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceUtilizationStatus) SetValue(v int32)`

SetValue sets Value field to given value.


### GetResourceType

`func (o *ResourceUtilizationStatus) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceUtilizationStatus) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceUtilizationStatus) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


