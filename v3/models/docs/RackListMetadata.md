# RackListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "rack"]
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Filter** | Pointer to **string** | The filter in FIQL syntax used for the results. | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewRackListMetadata

`func NewRackListMetadata() *RackListMetadata`

NewRackListMetadata instantiates a new RackListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackListMetadataWithDefaults

`func NewRackListMetadataWithDefaults() *RackListMetadata`

NewRackListMetadataWithDefaults instantiates a new RackListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RackListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RackListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RackListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RackListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortAttribute

`func (o *RackListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *RackListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *RackListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *RackListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *RackListMetadata) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RackListMetadata) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RackListMetadata) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RackListMetadata) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLength

`func (o *RackListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RackListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RackListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *RackListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *RackListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RackListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RackListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *RackListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *RackListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RackListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RackListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RackListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


