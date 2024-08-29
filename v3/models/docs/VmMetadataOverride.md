# VmMetadataOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoriesMapping** | Pointer to **map[string][]string** | Categories for the vm. This allows setting up multiple values from a single key.  | [optional] 
**UUID** | Pointer to **string** | vm uuid. | [optional] 

## Methods

### NewVmMetadataOverride

`func NewVmMetadataOverride() *VmMetadataOverride`

NewVmMetadataOverride instantiates a new VmMetadataOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmMetadataOverrideWithDefaults

`func NewVmMetadataOverrideWithDefaults() *VmMetadataOverride`

NewVmMetadataOverrideWithDefaults instantiates a new VmMetadataOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoriesMapping

`func (o *VmMetadataOverride) GetCategoriesMapping() map[string][]string`

GetCategoriesMapping returns the CategoriesMapping field if non-nil, zero value otherwise.

### GetCategoriesMappingOk

`func (o *VmMetadataOverride) GetCategoriesMappingOk() (*map[string][]string, bool)`

GetCategoriesMappingOk returns a tuple with the CategoriesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesMapping

`func (o *VmMetadataOverride) SetCategoriesMapping(v map[string][]string)`

SetCategoriesMapping sets CategoriesMapping field to given value.

### HasCategoriesMapping

`func (o *VmMetadataOverride) HasCategoriesMapping() bool`

HasCategoriesMapping returns a boolean if a field has been set.

### GetUUID

`func (o *VmMetadataOverride) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmMetadataOverride) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmMetadataOverride) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmMetadataOverride) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


