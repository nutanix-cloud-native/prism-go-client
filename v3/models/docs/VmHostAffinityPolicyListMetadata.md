# VmHostAffinityPolicyListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name. | [optional] [default to "vm_host_affinity_policy"]
**Length** | Pointer to **int64** | Number of records to retrieve relative to the offset. | [optional] 
**Offset** | Pointer to **int64** | Offset from the start of the entity list. | [optional] 

## Methods

### NewVmHostAffinityPolicyListMetadata

`func NewVmHostAffinityPolicyListMetadata() *VmHostAffinityPolicyListMetadata`

NewVmHostAffinityPolicyListMetadata instantiates a new VmHostAffinityPolicyListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyListMetadataWithDefaults

`func NewVmHostAffinityPolicyListMetadataWithDefaults() *VmHostAffinityPolicyListMetadata`

NewVmHostAffinityPolicyListMetadataWithDefaults instantiates a new VmHostAffinityPolicyListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *VmHostAffinityPolicyListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VmHostAffinityPolicyListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VmHostAffinityPolicyListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VmHostAffinityPolicyListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLength

`func (o *VmHostAffinityPolicyListMetadata) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VmHostAffinityPolicyListMetadata) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VmHostAffinityPolicyListMetadata) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *VmHostAffinityPolicyListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetOffset

`func (o *VmHostAffinityPolicyListMetadata) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VmHostAffinityPolicyListMetadata) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VmHostAffinityPolicyListMetadata) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VmHostAffinityPolicyListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


