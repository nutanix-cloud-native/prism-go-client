# MetadataOfTheRestoredVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoriesMapping** | Pointer to **map[string][]string** | Categories for the vm. This allows setting up multiple values from a single key.  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataOfTheRestoredVm

`func NewMetadataOfTheRestoredVm() *MetadataOfTheRestoredVm`

NewMetadataOfTheRestoredVm instantiates a new MetadataOfTheRestoredVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataOfTheRestoredVmWithDefaults

`func NewMetadataOfTheRestoredVmWithDefaults() *MetadataOfTheRestoredVm`

NewMetadataOfTheRestoredVmWithDefaults instantiates a new MetadataOfTheRestoredVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoriesMapping

`func (o *MetadataOfTheRestoredVm) GetCategoriesMapping() map[string][]string`

GetCategoriesMapping returns the CategoriesMapping field if non-nil, zero value otherwise.

### GetCategoriesMappingOk

`func (o *MetadataOfTheRestoredVm) GetCategoriesMappingOk() (*map[string][]string, bool)`

GetCategoriesMappingOk returns a tuple with the CategoriesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesMapping

`func (o *MetadataOfTheRestoredVm) SetCategoriesMapping(v map[string][]string)`

SetCategoriesMapping sets CategoriesMapping field to given value.

### HasCategoriesMapping

`func (o *MetadataOfTheRestoredVm) HasCategoriesMapping() bool`

HasCategoriesMapping returns a boolean if a field has been set.

### GetUUID

`func (o *MetadataOfTheRestoredVm) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *MetadataOfTheRestoredVm) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *MetadataOfTheRestoredVm) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *MetadataOfTheRestoredVm) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


