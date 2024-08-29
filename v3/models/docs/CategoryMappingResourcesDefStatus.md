# CategoryMappingResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryValue** | Pointer to **string** | The value for the category that this mapping is for. | [optional] 
**AdMapping** | Pointer to [**CategoryMappingResourcesDefStatusAdMapping**](CategoryMappingResourcesDefStatusAdMapping.md) |  | [optional] 
**CategoryName** | Pointer to **string** | The name for the category that this mapping is for. | [optional] 

## Methods

### NewCategoryMappingResourcesDefStatus

`func NewCategoryMappingResourcesDefStatus() *CategoryMappingResourcesDefStatus`

NewCategoryMappingResourcesDefStatus instantiates a new CategoryMappingResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingResourcesDefStatusWithDefaults

`func NewCategoryMappingResourcesDefStatusWithDefaults() *CategoryMappingResourcesDefStatus`

NewCategoryMappingResourcesDefStatusWithDefaults instantiates a new CategoryMappingResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryValue

`func (o *CategoryMappingResourcesDefStatus) GetCategoryValue() string`

GetCategoryValue returns the CategoryValue field if non-nil, zero value otherwise.

### GetCategoryValueOk

`func (o *CategoryMappingResourcesDefStatus) GetCategoryValueOk() (*string, bool)`

GetCategoryValueOk returns a tuple with the CategoryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryValue

`func (o *CategoryMappingResourcesDefStatus) SetCategoryValue(v string)`

SetCategoryValue sets CategoryValue field to given value.

### HasCategoryValue

`func (o *CategoryMappingResourcesDefStatus) HasCategoryValue() bool`

HasCategoryValue returns a boolean if a field has been set.

### GetAdMapping

`func (o *CategoryMappingResourcesDefStatus) GetAdMapping() CategoryMappingResourcesDefStatusAdMapping`

GetAdMapping returns the AdMapping field if non-nil, zero value otherwise.

### GetAdMappingOk

`func (o *CategoryMappingResourcesDefStatus) GetAdMappingOk() (*CategoryMappingResourcesDefStatusAdMapping, bool)`

GetAdMappingOk returns a tuple with the AdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdMapping

`func (o *CategoryMappingResourcesDefStatus) SetAdMapping(v CategoryMappingResourcesDefStatusAdMapping)`

SetAdMapping sets AdMapping field to given value.

### HasAdMapping

`func (o *CategoryMappingResourcesDefStatus) HasAdMapping() bool`

HasAdMapping returns a boolean if a field has been set.

### GetCategoryName

`func (o *CategoryMappingResourcesDefStatus) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *CategoryMappingResourcesDefStatus) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *CategoryMappingResourcesDefStatus) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *CategoryMappingResourcesDefStatus) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


