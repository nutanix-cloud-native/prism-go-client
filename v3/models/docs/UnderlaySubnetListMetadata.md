# UnderlaySubnetListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "underlay_subnet"]
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Filter** | Pointer to **string** | The filter in FIQL syntax used for the results. | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewUnderlaySubnetListMetadata

`func NewUnderlaySubnetListMetadata() *UnderlaySubnetListMetadata`

NewUnderlaySubnetListMetadata instantiates a new UnderlaySubnetListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetListMetadataWithDefaults

`func NewUnderlaySubnetListMetadataWithDefaults() *UnderlaySubnetListMetadata`

NewUnderlaySubnetListMetadataWithDefaults instantiates a new UnderlaySubnetListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *UnderlaySubnetListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UnderlaySubnetListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UnderlaySubnetListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UnderlaySubnetListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortAttribute

`func (o *UnderlaySubnetListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *UnderlaySubnetListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *UnderlaySubnetListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *UnderlaySubnetListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *UnderlaySubnetListMetadata) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UnderlaySubnetListMetadata) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UnderlaySubnetListMetadata) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UnderlaySubnetListMetadata) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLength

`func (o *UnderlaySubnetListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *UnderlaySubnetListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *UnderlaySubnetListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *UnderlaySubnetListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *UnderlaySubnetListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UnderlaySubnetListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UnderlaySubnetListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UnderlaySubnetListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *UnderlaySubnetListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UnderlaySubnetListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UnderlaySubnetListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UnderlaySubnetListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


