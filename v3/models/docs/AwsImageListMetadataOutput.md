# AwsImageListMetadataOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "aws_image"]
**TotalMatches** | Pointer to **int32** | Total matches found | [optional] 
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Length** | Pointer to **int32** | The number of records retrieved relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewAwsImageListMetadataOutput

`func NewAwsImageListMetadataOutput() *AwsImageListMetadataOutput`

NewAwsImageListMetadataOutput instantiates a new AwsImageListMetadataOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsImageListMetadataOutputWithDefaults

`func NewAwsImageListMetadataOutputWithDefaults() *AwsImageListMetadataOutput`

NewAwsImageListMetadataOutputWithDefaults instantiates a new AwsImageListMetadataOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AwsImageListMetadataOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsImageListMetadataOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsImageListMetadataOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AwsImageListMetadataOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTotalMatches

`func (o *AwsImageListMetadataOutput) GetTotalMatches() int32`

GetTotalMatches returns the TotalMatches field if non-nil, zero value otherwise.

### GetTotalMatchesOk

`func (o *AwsImageListMetadataOutput) GetTotalMatchesOk() (*int32, bool)`

GetTotalMatchesOk returns a tuple with the TotalMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatches

`func (o *AwsImageListMetadataOutput) SetTotalMatches(v int32)`

SetTotalMatches sets TotalMatches field to given value.

### HasTotalMatches

`func (o *AwsImageListMetadataOutput) HasTotalMatches() bool`

HasTotalMatches returns a boolean if a field has been set.

### GetSortAttribute

`func (o *AwsImageListMetadataOutput) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *AwsImageListMetadataOutput) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *AwsImageListMetadataOutput) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *AwsImageListMetadataOutput) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetLength

`func (o *AwsImageListMetadataOutput) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AwsImageListMetadataOutput) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AwsImageListMetadataOutput) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AwsImageListMetadataOutput) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *AwsImageListMetadataOutput) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AwsImageListMetadataOutput) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AwsImageListMetadataOutput) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AwsImageListMetadataOutput) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *AwsImageListMetadataOutput) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsImageListMetadataOutput) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsImageListMetadataOutput) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AwsImageListMetadataOutput) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


