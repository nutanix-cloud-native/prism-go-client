# AwsSecurityGroupListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "aws_security_group"]
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewAwsSecurityGroupListMetadata

`func NewAwsSecurityGroupListMetadata() *AwsSecurityGroupListMetadata`

NewAwsSecurityGroupListMetadata instantiates a new AwsSecurityGroupListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecurityGroupListMetadataWithDefaults

`func NewAwsSecurityGroupListMetadataWithDefaults() *AwsSecurityGroupListMetadata`

NewAwsSecurityGroupListMetadataWithDefaults instantiates a new AwsSecurityGroupListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AwsSecurityGroupListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsSecurityGroupListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsSecurityGroupListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AwsSecurityGroupListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortOrder

`func (o *AwsSecurityGroupListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AwsSecurityGroupListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AwsSecurityGroupListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AwsSecurityGroupListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLength

`func (o *AwsSecurityGroupListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AwsSecurityGroupListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AwsSecurityGroupListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AwsSecurityGroupListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortAttribute

`func (o *AwsSecurityGroupListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *AwsSecurityGroupListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *AwsSecurityGroupListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *AwsSecurityGroupListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetOffset

`func (o *AwsSecurityGroupListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsSecurityGroupListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsSecurityGroupListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AwsSecurityGroupListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


