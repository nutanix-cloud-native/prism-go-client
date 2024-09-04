# CategoryMappingIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CategoryMappingDefStatus**](CategoryMappingDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CategoryMapping**](CategoryMapping.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CategoryMappingMetadata**](CategoryMappingMetadata.md) |  | 

## Methods

### NewCategoryMappingIntentResource

`func NewCategoryMappingIntentResource(metadata CategoryMappingMetadata, ) *CategoryMappingIntentResource`

NewCategoryMappingIntentResource instantiates a new CategoryMappingIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingIntentResourceWithDefaults

`func NewCategoryMappingIntentResourceWithDefaults() *CategoryMappingIntentResource`

NewCategoryMappingIntentResourceWithDefaults instantiates a new CategoryMappingIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CategoryMappingIntentResource) GetStatus() CategoryMappingDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CategoryMappingIntentResource) GetStatusOk() (*CategoryMappingDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CategoryMappingIntentResource) SetStatus(v CategoryMappingDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CategoryMappingIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CategoryMappingIntentResource) GetSpec() CategoryMapping`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CategoryMappingIntentResource) GetSpecOk() (*CategoryMapping, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CategoryMappingIntentResource) SetSpec(v CategoryMapping)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CategoryMappingIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryMappingIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryMappingIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryMappingIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryMappingIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CategoryMappingIntentResource) GetMetadata() CategoryMappingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryMappingIntentResource) GetMetadataOk() (*CategoryMappingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryMappingIntentResource) SetMetadata(v CategoryMappingMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


