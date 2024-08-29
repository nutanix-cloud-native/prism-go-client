# OvaMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerReference** | Pointer to [**UserReference**](UserReference.md) |  | [optional] 
**Kind** | Pointer to **string** | The kind name | [optional] [default to "ova"]
**CreationTime** | Pointer to **time.Time** | Creation time of OVA | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewOvaMetadata

`func NewOvaMetadata() *OvaMetadata`

NewOvaMetadata instantiates a new OvaMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvaMetadataWithDefaults

`func NewOvaMetadataWithDefaults() *OvaMetadata`

NewOvaMetadataWithDefaults instantiates a new OvaMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerReference

`func (o *OvaMetadata) GetOwnerReference() UserReference`

GetOwnerReference returns the OwnerReference field if non-nil, zero value otherwise.

### GetOwnerReferenceOk

`func (o *OvaMetadata) GetOwnerReferenceOk() (*UserReference, bool)`

GetOwnerReferenceOk returns a tuple with the OwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReference

`func (o *OvaMetadata) SetOwnerReference(v UserReference)`

SetOwnerReference sets OwnerReference field to given value.

### HasOwnerReference

`func (o *OvaMetadata) HasOwnerReference() bool`

HasOwnerReference returns a boolean if a field has been set.

### GetKind

`func (o *OvaMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OvaMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OvaMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OvaMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCreationTime

`func (o *OvaMetadata) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *OvaMetadata) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *OvaMetadata) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *OvaMetadata) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUUID

`func (o *OvaMetadata) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *OvaMetadata) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *OvaMetadata) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *OvaMetadata) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


