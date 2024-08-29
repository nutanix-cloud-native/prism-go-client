# RemoteSyslogServerListMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | The kind name | [optional] [default to "remote_syslog_server"]
**SortAttribute** | Pointer to **string** | The attribute to perform sort on | [optional] 
**Filter** | Pointer to **string** | The filter in FIQL syntax used for the results. | [optional] 
**Length** | Pointer to **int32** | The number of records to retrieve relative to the offset | [optional] 
**SortOrder** | Pointer to **string** | The sort order in which results are returned | [optional] 
**Offset** | Pointer to **int32** | Offset from the start of the entity list | [optional] 

## Methods

### NewRemoteSyslogServerListMetadata

`func NewRemoteSyslogServerListMetadata() *RemoteSyslogServerListMetadata`

NewRemoteSyslogServerListMetadata instantiates a new RemoteSyslogServerListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSyslogServerListMetadataWithDefaults

`func NewRemoteSyslogServerListMetadataWithDefaults() *RemoteSyslogServerListMetadata`

NewRemoteSyslogServerListMetadataWithDefaults instantiates a new RemoteSyslogServerListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RemoteSyslogServerListMetadata) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RemoteSyslogServerListMetadata) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RemoteSyslogServerListMetadata) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RemoteSyslogServerListMetadata) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSortAttribute

`func (o *RemoteSyslogServerListMetadata) GetSortAttribute() string`

GetSortAttribute returns the SortAttribute field if non-nil, zero value otherwise.

### GetSortAttributeOk

`func (o *RemoteSyslogServerListMetadata) GetSortAttributeOk() (*string, bool)`

GetSortAttributeOk returns a tuple with the SortAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAttribute

`func (o *RemoteSyslogServerListMetadata) SetSortAttribute(v string)`

SetSortAttribute sets SortAttribute field to given value.

### HasSortAttribute

`func (o *RemoteSyslogServerListMetadata) HasSortAttribute() bool`

HasSortAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *RemoteSyslogServerListMetadata) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RemoteSyslogServerListMetadata) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RemoteSyslogServerListMetadata) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RemoteSyslogServerListMetadata) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLength

`func (o *RemoteSyslogServerListMetadata) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RemoteSyslogServerListMetadata) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RemoteSyslogServerListMetadata) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *RemoteSyslogServerListMetadata) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetSortOrder

`func (o *RemoteSyslogServerListMetadata) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RemoteSyslogServerListMetadata) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RemoteSyslogServerListMetadata) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *RemoteSyslogServerListMetadata) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetOffset

`func (o *RemoteSyslogServerListMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RemoteSyslogServerListMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RemoteSyslogServerListMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RemoteSyslogServerListMetadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


