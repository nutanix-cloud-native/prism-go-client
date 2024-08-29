# ImagePlacementPolicyListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "image_placement_policy"]
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Filter** | Pointer to **string** | The filter in FIQL syntax used for the results. | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewImagePlacementPolicyListMetadata

`func NewImagePlacementPolicyListMetadata() *ImagePlacementPolicyListMetadata`

NewImagePlacementPolicyListMetadata instantiates a new ImagePlacementPolicyListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyListMetadataWithDefaults

`func NewImagePlacementPolicyListMetadataWithDefaults() *ImagePlacementPolicyListMetadata`

NewImagePlacementPolicyListMetadataWithDefaults instantiates a new ImagePlacementPolicyListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ImagePlacementPolicyListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ImagePlacementPolicyListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ImagePlacementPolicyListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ImagePlacementPolicyListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortAttribute

`func (o *ImagePlacementPolicyListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *ImagePlacementPolicyListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *ImagePlacementPolicyListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *ImagePlacementPolicyListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *ImagePlacementPolicyListMetadata) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ImagePlacementPolicyListMetadata) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ImagePlacementPolicyListMetadata) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ImagePlacementPolicyListMetadata) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLength

`func (o *ImagePlacementPolicyListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ImagePlacementPolicyListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ImagePlacementPolicyListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ImagePlacementPolicyListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *ImagePlacementPolicyListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ImagePlacementPolicyListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ImagePlacementPolicyListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ImagePlacementPolicyListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *ImagePlacementPolicyListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ImagePlacementPolicyListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ImagePlacementPolicyListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ImagePlacementPolicyListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


