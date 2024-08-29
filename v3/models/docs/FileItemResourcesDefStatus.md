# FileItemResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetrievalUriList** | Pointer to **[]string** | List of URIs where the raw file_item data can be accessed. | [optional] 
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 
**SourceUri** | Pointer to **string** | URI that points at the file to create the file_item from. | [optional] 
**SourceAuth** | Pointer to [**SourceAuth**](SourceAuth.md) |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 

## Methods

### NewFileItemResourcesDefStatus

`func NewFileItemResourcesDefStatus() *FileItemResourcesDefStatus`

NewFileItemResourcesDefStatus instantiates a new FileItemResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileItemResourcesDefStatusWithDefaults

`func NewFileItemResourcesDefStatusWithDefaults() *FileItemResourcesDefStatus`

NewFileItemResourcesDefStatusWithDefaults instantiates a new FileItemResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetrievalUriList

`func (o *FileItemResourcesDefStatus) GetRetrievalUriList() []string`

GetRetrievalUriList returns the RetrievalUriList field if non-nil, zero value otherwise.

### GetRetrievalUriListOk

`func (o *FileItemResourcesDefStatus) GetRetrievalUriListOk() (*[]string, bool)`

GetRetrievalUriListOk returns a tuple with the RetrievalUriList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalUriList

`func (o *FileItemResourcesDefStatus) SetRetrievalUriList(v []string)`

SetRetrievalUriList sets RetrievalUriList field to given value.

### HasRetrievalUriList

`func (o *FileItemResourcesDefStatus) HasRetrievalUriList() bool`

HasRetrievalUriList returns a boolean if a field has been set.

### GetChecksum

`func (o *FileItemResourcesDefStatus) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *FileItemResourcesDefStatus) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *FileItemResourcesDefStatus) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *FileItemResourcesDefStatus) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetSourceUri

`func (o *FileItemResourcesDefStatus) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *FileItemResourcesDefStatus) GetSourceUriOk() (*string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUri

`func (o *FileItemResourcesDefStatus) SetSourceUri(v string)`

SetSourceUri sets SourceUri field to given value.

### HasSourceUri

`func (o *FileItemResourcesDefStatus) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### GetSourceAuth

`func (o *FileItemResourcesDefStatus) GetSourceAuth() SourceAuth`

GetSourceAuth returns the SourceAuth field if non-nil, zero value otherwise.

### GetSourceAuthOk

`func (o *FileItemResourcesDefStatus) GetSourceAuthOk() (*SourceAuth, bool)`

GetSourceAuthOk returns a tuple with the SourceAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAuth

`func (o *FileItemResourcesDefStatus) SetSourceAuth(v SourceAuth)`

SetSourceAuth sets SourceAuth field to given value.

### HasSourceAuth

`func (o *FileItemResourcesDefStatus) HasSourceAuth() bool`

HasSourceAuth returns a boolean if a field has been set.

### GetSizeBytes

`func (o *FileItemResourcesDefStatus) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *FileItemResourcesDefStatus) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *FileItemResourcesDefStatus) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *FileItemResourcesDefStatus) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *FileItemResourcesDefStatus) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *FileItemResourcesDefStatus) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *FileItemResourcesDefStatus) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *FileItemResourcesDefStatus) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


