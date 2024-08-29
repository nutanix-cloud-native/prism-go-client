# MetadataOfTheClonedVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | Pointer to **string** |  | [optional] 
**EntityVersion** | Pointer to **string** | Logical entity version of the VM from which to clone the new VM.  | [optional] 

## Methods

### NewMetadataOfTheClonedVm

`func NewMetadataOfTheClonedVm() *MetadataOfTheClonedVm`

NewMetadataOfTheClonedVm instantiates a new MetadataOfTheClonedVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataOfTheClonedVmWithDefaults

`func NewMetadataOfTheClonedVmWithDefaults() *MetadataOfTheClonedVm`

NewMetadataOfTheClonedVmWithDefaults instantiates a new MetadataOfTheClonedVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *MetadataOfTheClonedVm) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *MetadataOfTheClonedVm) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *MetadataOfTheClonedVm) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *MetadataOfTheClonedVm) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetEntityVersion

`func (o *MetadataOfTheClonedVm) GetEntityVersion() string`

GetEntityVersion returns the EntityVersion field if non-nil, zero value otherwise.

### GetEntityVersionOk

`func (o *MetadataOfTheClonedVm) GetEntityVersionOk() (*string, bool)`

GetEntityVersionOk returns a tuple with the EntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityVersion

`func (o *MetadataOfTheClonedVm) SetEntityVersion(v string)`

SetEntityVersion sets EntityVersion field to given value.

### HasEntityVersion

`func (o *MetadataOfTheClonedVm) HasEntityVersion() bool`

HasEntityVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


