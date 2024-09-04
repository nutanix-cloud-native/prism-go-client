# VmDiskOutputStatusDataSourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDirectAttach** | Pointer to **bool** | This is to indicate if attaching the referenced disk directly. Important: this should only be used by internal services. Direct attaching a disk that is used by another VM will result in data loss.  | [optional] 
**Url** | Pointer to **string** | AHV sync rep uses this field to identify a dormant VM disk.  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewVmDiskOutputStatusDataSourceReference

`func NewVmDiskOutputStatusDataSourceReference() *VmDiskOutputStatusDataSourceReference`

NewVmDiskOutputStatusDataSourceReference instantiates a new VmDiskOutputStatusDataSourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDiskOutputStatusDataSourceReferenceWithDefaults

`func NewVmDiskOutputStatusDataSourceReferenceWithDefaults() *VmDiskOutputStatusDataSourceReference`

NewVmDiskOutputStatusDataSourceReferenceWithDefaults instantiates a new VmDiskOutputStatusDataSourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDirectAttach

`func (o *VmDiskOutputStatusDataSourceReference) GetIsDirectAttach() bool`

GetIsDirectAttach returns the IsDirectAttach field if non-nil, zero value otherwise.

### GetIsDirectAttachOk

`func (o *VmDiskOutputStatusDataSourceReference) GetIsDirectAttachOk() (*bool, bool)`

GetIsDirectAttachOk returns a tuple with the IsDirectAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectAttach

`func (o *VmDiskOutputStatusDataSourceReference) SetIsDirectAttach(v bool)`

SetIsDirectAttach sets IsDirectAttach field to given value.

### HasIsDirectAttach

`func (o *VmDiskOutputStatusDataSourceReference) HasIsDirectAttach() bool`

HasIsDirectAttach returns a boolean if a field has been set.

### GetUrl

`func (o *VmDiskOutputStatusDataSourceReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VmDiskOutputStatusDataSourceReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VmDiskOutputStatusDataSourceReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VmDiskOutputStatusDataSourceReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKind

`func (o *VmDiskOutputStatusDataSourceReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VmDiskOutputStatusDataSourceReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VmDiskOutputStatusDataSourceReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VmDiskOutputStatusDataSourceReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUUID

`func (o *VmDiskOutputStatusDataSourceReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *VmDiskOutputStatusDataSourceReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *VmDiskOutputStatusDataSourceReference) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *VmDiskOutputStatusDataSourceReference) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *VmDiskOutputStatusDataSourceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmDiskOutputStatusDataSourceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmDiskOutputStatusDataSourceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmDiskOutputStatusDataSourceReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


