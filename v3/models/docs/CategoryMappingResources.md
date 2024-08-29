# CategoryMappingResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryValue** | **string** | The value for the category that this mapping is for. | 
**AdMapping** | [**CategoryMappingResourcesAdMapping**](CategoryMappingResourcesAdMapping.md) |  | 
**CategoryName** | **string** | The name for the category that this mapping is for. | 

## Methods

### NewCategoryMappingResources

`func NewCategoryMappingResources(categoryValue string, adMapping CategoryMappingResourcesAdMapping, categoryName string, ) *CategoryMappingResources`

NewCategoryMappingResources instantiates a new CategoryMappingResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingResourcesWithDefaults

`func NewCategoryMappingResourcesWithDefaults() *CategoryMappingResources`

NewCategoryMappingResourcesWithDefaults instantiates a new CategoryMappingResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryValue

`func (o *CategoryMappingResources) GetCategoryValue() string`

GetCategoryValue returns the CategoryValue field if non-nil, zero value otherwise.

### GetCategoryValueOk

`func (o *CategoryMappingResources) GetCategoryValueOk() (*string, bool)`

GetCategoryValueOk returns a tuple with the CategoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryValue

`func (o *CategoryMappingResources) SetCategoryValue(v string)`

SetCategoryValue sets CategoryValue field to given value.


### GetAdMapping

`func (o *CategoryMappingResources) GetAdMapping() CategoryMappingResourcesAdMapping`

GetAdMapping returns the AdMapping field if non-nil, zero value otherwise.

### GetAdMappingOk

`func (o *CategoryMappingResources) GetAdMappingOk() (*CategoryMappingResourcesAdMapping, bool)`

GetAdMappingOk returns a tuple with the AdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdMapping

`func (o *CategoryMappingResources) SetAdMapping(v CategoryMappingResourcesAdMapping)`

SetAdMapping sets AdMapping field to given value.


### GetCategoryName

`func (o *CategoryMappingResources) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *CategoryMappingResources) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *CategoryMappingResources) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


