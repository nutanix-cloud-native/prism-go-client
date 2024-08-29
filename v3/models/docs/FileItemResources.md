# FileItemResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to [**Checksum**](Checksum.md) |  | [optional] 
**SourceUri** | Pointer to **string** | URI that points at the file to create the file_item from. | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**SourceAuth** | Pointer to [**SourceAuth**](SourceAuth.md) |  | [optional] 

## Methods

### NewFileItemResources

`func NewFileItemResources() *FileItemResources`

NewFileItemResources instantiates a new FileItemResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileItemResourcesWithDefaults

`func NewFileItemResourcesWithDefaults() *FileItemResources`

NewFileItemResourcesWithDefaults instantiates a new FileItemResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *FileItemResources) GetChecksum() Checksum`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *FileItemResources) GetChecksumOk() (*Checksum, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *FileItemResources) SetChecksum(v Checksum)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *FileItemResources) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetSourceUri

`func (o *FileItemResources) GetSourceUri() string`

GetSourceUri returns the SourceUri field if non-nil, zero value otherwise.

### GetSourceUriOk

`func (o *FileItemResources) GetSourceUriOk() (*string, bool)`

GetSourceUriOk returns a tuple with the SourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUri

`func (o *FileItemResources) SetSourceUri(v string)`

SetSourceUri sets SourceUri field to given value.

### HasSourceUri

`func (o *FileItemResources) HasSourceUri() bool`

HasSourceUri returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *FileItemResources) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *FileItemResources) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *FileItemResources) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *FileItemResources) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetSourceAuth

`func (o *FileItemResources) GetSourceAuth() SourceAuth`

GetSourceAuth returns the SourceAuth field if non-nil, zero value otherwise.

### GetSourceAuthOk

`func (o *FileItemResources) GetSourceAuthOk() (*SourceAuth, bool)`

GetSourceAuthOk returns a tuple with the SourceAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAuth

`func (o *FileItemResources) SetSourceAuth(v SourceAuth)`

SetSourceAuth sets SourceAuth field to given value.

### HasSourceAuth

`func (o *FileItemResources) HasSourceAuth() bool`

HasSourceAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


