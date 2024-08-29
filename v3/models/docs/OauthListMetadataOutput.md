# OauthListMetadataOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "oauth"]
**TotalMatches** | Pointer to **int32** | Total matches found | [optional] 
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Filter** | Pointer to **string** | The filter used for the results | [optional] 
**Length** | Pointer to **int32** | The number of records retrieved relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewOauthListMetadataOutput

`func NewOauthListMetadataOutput() *OauthListMetadataOutput`

NewOauthListMetadataOutput instantiates a new OauthListMetadataOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthListMetadataOutputWithDefaults

`func NewOauthListMetadataOutputWithDefaults() *OauthListMetadataOutput`

NewOauthListMetadataOutputWithDefaults instantiates a new OauthListMetadataOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *OauthListMetadataOutput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OauthListMetadataOutput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OauthListMetadataOutput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OauthListMetadataOutput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTotalMatches

`func (o *OauthListMetadataOutput) GetTotalMatches() int32`

GetTotalMatches returns the TotalMatches field if non-nil, zero value otherwise.

### GetTotalMatchesOk

`func (o *OauthListMetadataOutput) GetTotalMatchesOk() (*int32, bool)`

GetTotalMatchesOk returns a tuple with the TotalMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatches

`func (o *OauthListMetadataOutput) SetTotalMatches(v int32)`

SetTotalMatches sets TotalMatches field to given value.

### HasTotalMatches

`func (o *OauthListMetadataOutput) HasTotalMatches() bool`

HasTotalMatches returns a boolean if a field has been set.

### GetSortAttribute

`func (o *OauthListMetadataOutput) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *OauthListMetadataOutput) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *OauthListMetadataOutput) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *OauthListMetadataOutput) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *OauthListMetadataOutput) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OauthListMetadataOutput) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OauthListMetadataOutput) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *OauthListMetadataOutput) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLength

`func (o *OauthListMetadataOutput) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OauthListMetadataOutput) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OauthListMetadataOutput) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *OauthListMetadataOutput) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *OauthListMetadataOutput) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *OauthListMetadataOutput) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *OauthListMetadataOutput) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *OauthListMetadataOutput) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *OauthListMetadataOutput) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *OauthListMetadataOutput) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *OauthListMetadataOutput) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *OauthListMetadataOutput) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


