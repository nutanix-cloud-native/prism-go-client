# DirectoryServiceSearchMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The search string. | 
**SearchedAttributeList** | Pointer to **[]string** | The attributes for search operation. If not specified, search is performed with the common name.  | [optional] 
**ReturnedAttributeList** | Pointer to **[]string** | The attributes the search operation returns. If not available, a list of default attributes is returned.  | [optional] 
**IsWildcardSearch** | Pointer to **bool** | The attribute that tells if the query is a wildcard match or exact match query.  | [optional] [default to true]

## Methods

### NewDirectoryServiceSearchMetadata

`func NewDirectoryServiceSearchMetadata(query string, ) *DirectoryServiceSearchMetadata`

NewDirectoryServiceSearchMetadata instantiates a new DirectoryServiceSearchMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceSearchMetadataWithDefaults

`func NewDirectoryServiceSearchMetadataWithDefaults() *DirectoryServiceSearchMetadata`

NewDirectoryServiceSearchMetadataWithDefaults instantiates a new DirectoryServiceSearchMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *DirectoryServiceSearchMetadata) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DirectoryServiceSearchMetadata) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DirectoryServiceSearchMetadata) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetSearchedAttributeList

`func (o *DirectoryServiceSearchMetadata) GetSearchedAttributeList() []string`

GetSearchedAttributeList returns the SearchedAttributeList field if non-nil, zero value otherwise.

### GetSearchedAttributeListOk

`func (o *DirectoryServiceSearchMetadata) GetSearchedAttributeListOk() (*[]string, bool)`

GetSearchedAttributeListOk returns a tuple with the SearchedAttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchedAttributeList

`func (o *DirectoryServiceSearchMetadata) SetSearchedAttributeList(v []string)`

SetSearchedAttributeList sets SearchedAttributeList field to given value.

### HasSearchedAttributeList

`func (o *DirectoryServiceSearchMetadata) HasSearchedAttributeList() bool`

HasSearchedAttributeList returns a boolean if a field has been set.

### GetReturnedAttributeList

`func (o *DirectoryServiceSearchMetadata) GetReturnedAttributeList() []string`

GetReturnedAttributeList returns the ReturnedAttributeList field if non-nil, zero value otherwise.

### GetReturnedAttributeListOk

`func (o *DirectoryServiceSearchMetadata) GetReturnedAttributeListOk() (*[]string, bool)`

GetReturnedAttributeListOk returns a tuple with the ReturnedAttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedAttributeList

`func (o *DirectoryServiceSearchMetadata) SetReturnedAttributeList(v []string)`

SetReturnedAttributeList sets ReturnedAttributeList field to given value.

### HasReturnedAttributeList

`func (o *DirectoryServiceSearchMetadata) HasReturnedAttributeList() bool`

HasReturnedAttributeList returns a boolean if a field has been set.

### GetIsWildcardSearch

`func (o *DirectoryServiceSearchMetadata) GetIsWildcardSearch() bool`

GetIsWildcardSearch returns the IsWildcardSearch field if non-nil, zero value otherwise.

### GetIsWildcardSearchOk

`func (o *DirectoryServiceSearchMetadata) GetIsWildcardSearchOk() (*bool, bool)`

GetIsWildcardSearchOk returns a tuple with the IsWildcardSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWildcardSearch

`func (o *DirectoryServiceSearchMetadata) SetIsWildcardSearch(v bool)`

SetIsWildcardSearch sets IsWildcardSearch field to given value.

### HasIsWildcardSearch

`func (o *DirectoryServiceSearchMetadata) HasIsWildcardSearch() bool`

HasIsWildcardSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


