# IpfixExporterMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when ipfix_exporter was last updated  | [optional] [readonly] 
**UseCategoriesMapping** | Pointer to **bool** | Client need to specify this field as true if user want to use the newer way of assigning the categories. Without this things should work as it was earlier.  | [optional] [default to false]
**Kind** | **string** | The kind name | [readonly] [default to "ipfix_exporter"]
**UUID** | Pointer to **string** | ipfix_exporter uuid | [optional] 
**ProjectReference** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when ipfix_exporter was created  | [optional] [readonly] 
**SpecVersion** | Pointer to **int32** | Version number of the latest spec. | [optional] 
**SpecHash** | Pointer to **string** | Hash of the spec. This will be returned from server.  | [optional] 
**CategoriesMapping** | Pointer to **map[string][]string** | Categories for the ipfix_exporter. This allows setting up multiple values from a single key. Categories assigned using the older view will be present here. This is the new way of assigning categories.  | [optional] 
**ShouldForceTranslate** | Pointer to **bool** | Applied on Prism Central only. Indicate whether force to translate the spec of the fanout request to fit the target cluster API schema.  | [optional] 
**EntityVersion** | Pointer to **string** | Logical entity version that allows serializing updates to the entity across multiple API namespaces.  For kinds that support entity_version, it overrides spec_version described above.  | [optional] [readonly] 
**OwnerReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**Categories** | Pointer to **map[string]string** | Categories for the ipfix_exporter. This allows assigning one value of a key to any entity. Changes done in this will be reflected in the categories_mapping field.  | [optional] 
**Name** | Pointer to **string** | ipfix_exporter name | [optional] [readonly] 

## Methods

### NewIpfixExporterMetadata

`func NewIpfixExporterMetadata(kind string, ) *IpfixExporterMetadata`

NewIpfixExporterMetadata instantiates a new IpfixExporterMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixExporterMetadataWithDefaults

`func NewIpfixExporterMetadataWithDefaults() *IpfixExporterMetadata`

NewIpfixExporterMetadataWithDefaults instantiates a new IpfixExporterMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *IpfixExporterMetadata) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *IpfixExporterMetadata) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *IpfixExporterMetadata) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *IpfixExporterMetadata) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetUseCategoriesMapping

`func (o *IpfixExporterMetadata) GetUseCategoriesMapping() bool`

GetUseCategoriesMapping returns the UseCategoriesMapping field if non-nil, zero value otherwise.

### GetUseCategoriesMappingOk

`func (o *IpfixExporterMetadata) GetUseCategoriesMappingOk() (*bool, bool)`

GetUseCategoriesMappingOk returns a tuple with the UseCategoriesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCategoriesMapping

`func (o *IpfixExporterMetadata) SetUseCategoriesMapping(v bool)`

SetUseCategoriesMapping sets UseCategoriesMapping field to given value.

### HasUseCategoriesMapping

`func (o *IpfixExporterMetadata) HasUseCategoriesMapping() bool`

HasUseCategoriesMapping returns a boolean if a field has been set.

### GetKind

`func (o *IpfixExporterMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IpfixExporterMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IpfixExporterMetadata) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUUID

`func (o *IpfixExporterMetadata) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *IpfixExporterMetadata) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *IpfixExporterMetadata) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *IpfixExporterMetadata) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetProjectReference

`func (o *IpfixExporterMetadata) GetProjectReference() ProjectReference`

GetProjectReference returns the ProjectReference field if non-nil, zero value otherwise.

### GetProjectReferenceOk

`func (o *IpfixExporterMetadata) GetProjectReferenceOk() (*ProjectReference, bool)`

GetProjectReferenceOk returns a tuple with the ProjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReference

`func (o *IpfixExporterMetadata) SetProjectReference(v ProjectReference)`

SetProjectReference sets ProjectReference field to given value.

### HasProjectReference

`func (o *IpfixExporterMetadata) HasProjectReference() bool`

HasProjectReference returns a boolean if a field has been set.

### GetCreationTime

`func (o *IpfixExporterMetadata) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IpfixExporterMetadata) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IpfixExporterMetadata) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *IpfixExporterMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetSpecVersion

`func (o *IpfixExporterMetadata) GetSpecVersion() int32`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *IpfixExporterMetadata) GetSpecVersionOk() (*int32, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *IpfixExporterMetadata) SetSpecVersion(v int32)`

SetSpecVersion sets SpecVersion field to given value.

### HasSpecVersion

`func (o *IpfixExporterMetadata) HasSpecVersion() bool`

HasSpecVersion returns a boolean if a field has been set.

### GetSpecHash

`func (o *IpfixExporterMetadata) GetSpecHash() string`

GetSpecHash returns the SpecHash field if non-nil, zero value otherwise.

### GetSpecHashOk

`func (o *IpfixExporterMetadata) GetSpecHashOk() (*string, bool)`

GetSpecHashOk returns a tuple with the SpecHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecHash

`func (o *IpfixExporterMetadata) SetSpecHash(v string)`

SetSpecHash sets SpecHash field to given value.

### HasSpecHash

`func (o *IpfixExporterMetadata) HasSpecHash() bool`

HasSpecHash returns a boolean if a field has been set.

### GetCategoriesMapping

`func (o *IpfixExporterMetadata) GetCategoriesMapping() map[string][]string`

GetCategoriesMapping returns the CategoriesMapping field if non-nil, zero value otherwise.

### GetCategoriesMappingOk

`func (o *IpfixExporterMetadata) GetCategoriesMappingOk() (*map[string][]string, bool)`

GetCategoriesMappingOk returns a tuple with the CategoriesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesMapping

`func (o *IpfixExporterMetadata) SetCategoriesMapping(v map[string][]string)`

SetCategoriesMapping sets CategoriesMapping field to given value.

### HasCategoriesMapping

`func (o *IpfixExporterMetadata) HasCategoriesMapping() bool`

HasCategoriesMapping returns a boolean if a field has been set.

### GetShouldForceTranslate

`func (o *IpfixExporterMetadata) GetShouldForceTranslate() bool`

GetShouldForceTranslate returns the ShouldForceTranslate field if non-nil, zero value otherwise.

### GetShouldForceTranslateOk

`func (o *IpfixExporterMetadata) GetShouldForceTranslateOk() (*bool, bool)`

GetShouldForceTranslateOk returns a tuple with the ShouldForceTranslate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldForceTranslate

`func (o *IpfixExporterMetadata) SetShouldForceTranslate(v bool)`

SetShouldForceTranslate sets ShouldForceTranslate field to given value.

### HasShouldForceTranslate

`func (o *IpfixExporterMetadata) HasShouldForceTranslate() bool`

HasShouldForceTranslate returns a boolean if a field has been set.

### GetEntityVersion

`func (o *IpfixExporterMetadata) GetEntityVersion() string`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *IpfixExporterMetadata) GetEntityVersionOk() (*string, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *IpfixExporterMetadata) SetEntityVersion(v string)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *IpfixExporterMetadata) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.

### GetOwnerReference

`func (o *IpfixExporterMetadata) GetOwnerReference() UserReference`

GetOwnerReference returns the OwnerReference field if non-nil, zero value otherwise.

### GetOwnerReferenceOk

`func (o *IpfixExporterMetadata) GetOwnerReferenceOk() (*UserReference, bool)`

GetOwnerReferenceOk returns a tuple with the OwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReference

`func (o *IpfixExporterMetadata) SetOwnerReference(v UserReference)`

SetOwnerReference sets OwnerReference field to given value.

### HasOwnerReference

`func (o *IpfixExporterMetadata) HasOwnerReference() bool`

HasOwnerReference returns a boolean if a field has been set.

### GetCategories

`func (o *IpfixExporterMetadata) GetCategories() map[string]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *IpfixExporterMetadata) GetCategoriesOk() (*map[string]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *IpfixExporterMetadata) SetCategories(v map[string]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *IpfixExporterMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetName

`func (o *IpfixExporterMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpfixExporterMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpfixExporterMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpfixExporterMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


