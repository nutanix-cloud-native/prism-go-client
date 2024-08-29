# AwsVpcListMetadataOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "aws_vpc"]
**TotalMatches** | Pointer to **int32** | Total matches found | [optional] 
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Length** | Pointer to **int32** | The number of records retrieved relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewAwsVpcListMetadataOutput

`func NewAwsVpcListMetadataOutput() *AwsVpcListMetadataOutput`

NewAwsVpcListMetadataOutput instantiates a new AwsVpcListMetadataOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcListMetadataOutputWithDefaults

`func NewAwsVpcListMetadataOutputWithDefaults() *AwsVpcListMetadataOutput`

NewAwsVpcListMetadataOutputWithDefaults instantiates a new AwsVpcListMetadataOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AwsVpcListMetadataOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsVpcListMetadataOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsVpcListMetadataOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AwsVpcListMetadataOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTotalMatches

`func (o *AwsVpcListMetadataOutput) GetTotalMatches() int32`

GetTotalMatches returns the TotalMatches field if non-nil, zero value otherwise.

### GetTotalMatchesOk

`func (o *AwsVpcListMetadataOutput) GetTotalMatchesOk() (*int32, bool)`

GetTotalMatchesOk returns a tuple with the TotalMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatches

`func (o *AwsVpcListMetadataOutput) SetTotalMatches(v int32)`

SetTotalMatches sets TotalMatches field to given value.

### HasTotalMatches

`func (o *AwsVpcListMetadataOutput) HasTotalMatches() bool`

HasTotalMatches returns a boolean if a field has been set.

### GetSortAttribute

`func (o *AwsVpcListMetadataOutput) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *AwsVpcListMetadataOutput) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *AwsVpcListMetadataOutput) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *AwsVpcListMetadataOutput) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetLength

`func (o *AwsVpcListMetadataOutput) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AwsVpcListMetadataOutput) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AwsVpcListMetadataOutput) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AwsVpcListMetadataOutput) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *AwsVpcListMetadataOutput) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AwsVpcListMetadataOutput) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AwsVpcListMetadataOutput) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AwsVpcListMetadataOutput) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *AwsVpcListMetadataOutput) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsVpcListMetadataOutput) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsVpcListMetadataOutput) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AwsVpcListMetadataOutput) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


