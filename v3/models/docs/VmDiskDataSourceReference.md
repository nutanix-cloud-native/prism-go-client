# VmDiskDataSourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDirectAttach** | Pointer to **bool** | This is to indicate if attaching the referenced disk directly. Important: this should only be used by internal services. Direct attaching a disk that is used by another VM will result in data loss.  | [optional] 
**Url** | Pointer to **string** | AHV sync rep uses this field to identify a dormant VM disk.  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewVmDiskDataSourceReference

`func NewVmDiskDataSourceReference() *VmDiskDataSourceReference`

NewVmDiskDataSourceReference instantiates a new VmDiskDataSourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiskDataSourceReferenceWithDefaults

`func NewVmDiskDataSourceReferenceWithDefaults() *VmDiskDataSourceReference`

NewVmDiskDataSourceReferenceWithDefaults instantiates a new VmDiskDataSourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDirectAttach

`func (o *VmDiskDataSourceReference) GetIsDirectAttach() bool`

GetIsDirectAttach returns the IsDirectAttach field if non-nil, zero value otherwise.

### GetIsDirectAttachOk

`func (o *VmDiskDataSourceReference) GetIsDirectAttachOk() (*bool, bool)`

GetIsDirectAttachOk returns a tuple with the IsDirectAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectAttach

`func (o *VmDiskDataSourceReference) SetIsDirectAttach(v bool)`

SetIsDirectAttach sets IsDirectAttach field to given value.

### HasIsDirectAttach

`func (o *VmDiskDataSourceReference) HasIsDirectAttach() bool`

HasIsDirectAttach returns a boolean if a field has been set.

### GetUrl

`func (o *VmDiskDataSourceReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VmDiskDataSourceReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VmDiskDataSourceReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VmDiskDataSourceReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKind

`func (o *VmDiskDataSourceReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VmDiskDataSourceReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VmDiskDataSourceReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VmDiskDataSourceReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUUID

`func (o *VmDiskDataSourceReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmDiskDataSourceReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmDiskDataSourceReference) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmDiskDataSourceReference) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *VmDiskDataSourceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmDiskDataSourceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmDiskDataSourceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmDiskDataSourceReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


