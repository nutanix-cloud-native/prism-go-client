# CategoryMappingResourcesAdMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIdentifier** | **string** | The objectGUID for the object in AD. | 
**DirectoryServiceReference** | [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | 

## Methods

### NewCategoryMappingResourcesAdMapping

`func NewCategoryMappingResourcesAdMapping(objectIdentifier string, directoryServiceReference DirectoryServiceReference, ) *CategoryMappingResourcesAdMapping`

NewCategoryMappingResourcesAdMapping instantiates a new CategoryMappingResourcesAdMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingResourcesAdMappingWithDefaults

`func NewCategoryMappingResourcesAdMappingWithDefaults() *CategoryMappingResourcesAdMapping`

NewCategoryMappingResourcesAdMappingWithDefaults instantiates a new CategoryMappingResourcesAdMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIdentifier

`func (o *CategoryMappingResourcesAdMapping) GetObjectIdentifier() string`

GetObjectIdentifier returns the ObjectIdentifier field if non-nil, zero value otherwise.

### GetObjectIdentifierOk

`func (o *CategoryMappingResourcesAdMapping) GetObjectIdentifierOk() (*string, bool)`

GetObjectIdentifierOk returns a tuple with the ObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifier

`func (o *CategoryMappingResourcesAdMapping) SetObjectIdentifier(v string)`

SetObjectIdentifier sets ObjectIdentifier field to given value.


### GetDirectoryServiceReference

`func (o *CategoryMappingResourcesAdMapping) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *CategoryMappingResourcesAdMapping) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *CategoryMappingResourcesAdMapping) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


