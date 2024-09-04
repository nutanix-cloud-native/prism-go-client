# AwsAvailabilityZoneListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "aws_availability_zone"]
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewAwsAvailabilityZoneListMetadata

`func NewAwsAvailabilityZoneListMetadata() *AwsAvailabilityZoneListMetadata`

NewAwsAvailabilityZoneListMetadata instantiates a new AwsAvailabilityZoneListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAvailabilityZoneListMetadataWithDefaults

`func NewAwsAvailabilityZoneListMetadataWithDefaults() *AwsAvailabilityZoneListMetadata`

NewAwsAvailabilityZoneListMetadataWithDefaults instantiates a new AwsAvailabilityZoneListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AwsAvailabilityZoneListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AwsAvailabilityZoneListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AwsAvailabilityZoneListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AwsAvailabilityZoneListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortOrder

`func (o *AwsAvailabilityZoneListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AwsAvailabilityZoneListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AwsAvailabilityZoneListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AwsAvailabilityZoneListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetLength

`func (o *AwsAvailabilityZoneListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AwsAvailabilityZoneListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AwsAvailabilityZoneListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AwsAvailabilityZoneListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortAttribute

`func (o *AwsAvailabilityZoneListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *AwsAvailabilityZoneListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *AwsAvailabilityZoneListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *AwsAvailabilityZoneListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetOffset

`func (o *AwsAvailabilityZoneListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AwsAvailabilityZoneListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AwsAvailabilityZoneListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AwsAvailabilityZoneListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


