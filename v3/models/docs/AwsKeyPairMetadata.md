# AwsKeyPairMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when aws_key_pair was last updated  | [optional] [readonly] 
**Kind** | **string** | The kind name | [readonly] [default to "aws_key_pair"]
**UUID** | Pointer to **string** | aws_key_pair uuid | [optional] 
**ProjectReference** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when aws_key_pair was created  | [optional] [readonly] 
**SpecVersion** | Pointer to **int32** | Version number of the latest spec. | [optional] 
**OwnerReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**Categories** | Pointer to **map[string]string** | Categories for the aws_key_pair | [optional] 
**Name** | Pointer to **string** | aws_key_pair name | [optional] [readonly] 

## Methods

### NewAwsKeyPairMetadata

`func NewAwsKeyPairMetadata(kind string, ) *AwsKeyPairMetadata`

NewAwsKeyPairMetadata instantiates a new AwsKeyPairMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKeyPairMetadataWithDefaults

`func NewAwsKeyPairMetadataWithDefaults() *AwsKeyPairMetadata`

NewAwsKeyPairMetadataWithDefaults instantiates a new AwsKeyPairMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *AwsKeyPairMetadata) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *AwsKeyPairMetadata) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *AwsKeyPairMetadata) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *AwsKeyPairMetadata) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetKind

`func (o *AwsKeyPairMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsKeyPairMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsKeyPairMetadata) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUUID

`func (o *AwsKeyPairMetadata) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AwsKeyPairMetadata) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AwsKeyPairMetadata) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AwsKeyPairMetadata) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetProjectReference

`func (o *AwsKeyPairMetadata) GetProjectReference() ProjectReference`

GetProjectReference returns the ProjectReference field if non-nil, zero value otherwise.

### GetProjectReferenceOk

`func (o *AwsKeyPairMetadata) GetProjectReferenceOk() (*ProjectReference, bool)`

GetProjectReferenceOk returns a tuple with the ProjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReference

`func (o *AwsKeyPairMetadata) SetProjectReference(v ProjectReference)`

SetProjectReference sets ProjectReference field to given value.

### HasProjectReference

`func (o *AwsKeyPairMetadata) HasProjectReference() bool`

HasProjectReference returns a boolean if a field has been set.

### GetCreationTime

`func (o *AwsKeyPairMetadata) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AwsKeyPairMetadata) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AwsKeyPairMetadata) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AwsKeyPairMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetSpecVersion

`func (o *AwsKeyPairMetadata) GetSpecVersion() int32`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *AwsKeyPairMetadata) GetSpecVersionOk() (*int32, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *AwsKeyPairMetadata) SetSpecVersion(v int32)`

SetSpecVersion sets SpecVersion field to given value.

### HasSpecVersion

`func (o *AwsKeyPairMetadata) HasSpecVersion() bool`

HasSpecVersion returns a boolean if a field has been set.

### GetOwnerReference

`func (o *AwsKeyPairMetadata) GetOwnerReference() UserReference`

GetOwnerReference returns the OwnerReference field if non-nil, zero value otherwise.

### GetOwnerReferenceOk

`func (o *AwsKeyPairMetadata) GetOwnerReferenceOk() (*UserReference, bool)`

GetOwnerReferenceOk returns a tuple with the OwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReference

`func (o *AwsKeyPairMetadata) SetOwnerReference(v UserReference)`

SetOwnerReference sets OwnerReference field to given value.

### HasOwnerReference

`func (o *AwsKeyPairMetadata) HasOwnerReference() bool`

HasOwnerReference returns a boolean if a field has been set.

### GetCategories

`func (o *AwsKeyPairMetadata) GetCategories() map[string]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AwsKeyPairMetadata) GetCategoriesOk() (*map[string]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AwsKeyPairMetadata) SetCategories(v map[string]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AwsKeyPairMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetName

`func (o *AwsKeyPairMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsKeyPairMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsKeyPairMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsKeyPairMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


