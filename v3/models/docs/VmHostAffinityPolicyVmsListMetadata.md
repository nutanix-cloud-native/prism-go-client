# VmHostAffinityPolicyVmsListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "vm"]
**Length** | Pointer to **int64** | The number of records to retrieve relative to the offset. | [optional] 
**Offset** | Pointer to **int64** | Offset from the start of the entity list. | [optional] 

## Methods

### NewVmHostAffinityPolicyVmsListMetadata

`func NewVmHostAffinityPolicyVmsListMetadata() *VmHostAffinityPolicyVmsListMetadata`

NewVmHostAffinityPolicyVmsListMetadata instantiates a new VmHostAffinityPolicyVmsListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyVmsListMetadataWithDefaults

`func NewVmHostAffinityPolicyVmsListMetadataWithDefaults() *VmHostAffinityPolicyVmsListMetadata`

NewVmHostAffinityPolicyVmsListMetadataWithDefaults instantiates a new VmHostAffinityPolicyVmsListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *VmHostAffinityPolicyVmsListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VmHostAffinityPolicyVmsListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VmHostAffinityPolicyVmsListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VmHostAffinityPolicyVmsListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLength

`func (o *VmHostAffinityPolicyVmsListMetadata) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *VmHostAffinityPolicyVmsListMetadata) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *VmHostAffinityPolicyVmsListMetadata) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *VmHostAffinityPolicyVmsListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetOffset

`func (o *VmHostAffinityPolicyVmsListMetadata) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VmHostAffinityPolicyVmsListMetadata) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VmHostAffinityPolicyVmsListMetadata) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VmHostAffinityPolicyVmsListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


