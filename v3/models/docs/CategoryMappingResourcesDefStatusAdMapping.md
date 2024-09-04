# CategoryMappingResourcesDefStatusAdMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIdentifier** | Pointer to **string** | The objectGUID for the object in AD. | [optional] 
**ObjectCommonName** | Pointer to **string** | The CN for the mapped object in AD. | [optional] [readonly] 
**Status** | Pointer to **string** | Whether the mapping is usable or not; USABLE means it is usable. DELETED means the mapped object has been removed from AD, and DIRECTORY_NOT_CONFIGURED means either the directory service the mapping references has been removed or that directory service is not currently in use for identity categorization.  | [optional] [readonly] 
**DirectoryServiceReference** | Pointer to [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | [optional] 
**ObjectPath** | Pointer to **string** | The path for the mapped object in AD. | [optional] [readonly] 

## Methods

### NewCategoryMappingResourcesDefStatusAdMapping

`func NewCategoryMappingResourcesDefStatusAdMapping() *CategoryMappingResourcesDefStatusAdMapping`

NewCategoryMappingResourcesDefStatusAdMapping instantiates a new CategoryMappingResourcesDefStatusAdMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryMappingResourcesDefStatusAdMappingWithDefaults

`func NewCategoryMappingResourcesDefStatusAdMappingWithDefaults() *CategoryMappingResourcesDefStatusAdMapping`

NewCategoryMappingResourcesDefStatusAdMappingWithDefaults instantiates a new CategoryMappingResourcesDefStatusAdMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIdentifier

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectIdentifier() string`

GetObjectIdentifier returns the ObjectIdentifier field if non-nil, zero value otherwise.

### GetObjectIdentifierOk

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectIdentifierOk() (*string, bool)`

GetObjectIdentifierOk returns a tuple with the ObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifier

`func (o *CategoryMappingResourcesDefStatusAdMapping) SetObjectIdentifier(v string)`

SetObjectIdentifier sets ObjectIdentifier field to given value.

### HasObjectIdentifier

`func (o *CategoryMappingResourcesDefStatusAdMapping) HasObjectIdentifier() bool`

HasObjectIdentifier returns a boolean if a field has been set.

### GetObjectCommonName

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectCommonName() string`

GetObjectCommonName returns the ObjectCommonName field if non-nil, zero value otherwise.

### GetObjectCommonNameOk

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectCommonNameOk() (*string, bool)`

GetObjectCommonNameOk returns a tuple with the ObjectCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCommonName

`func (o *CategoryMappingResourcesDefStatusAdMapping) SetObjectCommonName(v string)`

SetObjectCommonName sets ObjectCommonName field to given value.

### HasObjectCommonName

`func (o *CategoryMappingResourcesDefStatusAdMapping) HasObjectCommonName() bool`

HasObjectCommonName returns a boolean if a field has been set.

### GetStatus

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CategoryMappingResourcesDefStatusAdMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CategoryMappingResourcesDefStatusAdMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirectoryServiceReference

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *CategoryMappingResourcesDefStatusAdMapping) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.

### HasDirectoryServiceReference

`func (o *CategoryMappingResourcesDefStatusAdMapping) HasDirectoryServiceReference() bool`

HasDirectoryServiceReference returns a boolean if a field has been set.

### GetObjectPath

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectPath() string`

GetObjectPath returns the ObjectPath field if non-nil, zero value otherwise.

### GetObjectPathOk

`func (o *CategoryMappingResourcesDefStatusAdMapping) GetObjectPathOk() (*string, bool)`

GetObjectPathOk returns a tuple with the ObjectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPath

`func (o *CategoryMappingResourcesDefStatusAdMapping) SetObjectPath(v string)`

SetObjectPath sets ObjectPath field to given value.

### HasObjectPath

`func (o *CategoryMappingResourcesDefStatusAdMapping) HasObjectPath() bool`

HasObjectPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


