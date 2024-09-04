# DirectoryServiceSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchResultList** | Pointer to [**[]Entity**](Entity.md) |  | [optional] 
**DomainName** | Pointer to **string** | The domain name of the directory service. | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | Pointer to [**DirectoryServiceSearchMetadata**](DirectoryServiceSearchMetadata.md) |  | [optional] 

## Methods

### NewDirectoryServiceSearchResponse

`func NewDirectoryServiceSearchResponse() *DirectoryServiceSearchResponse`

NewDirectoryServiceSearchResponse instantiates a new DirectoryServiceSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceSearchResponseWithDefaults

`func NewDirectoryServiceSearchResponseWithDefaults() *DirectoryServiceSearchResponse`

NewDirectoryServiceSearchResponseWithDefaults instantiates a new DirectoryServiceSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchResultList

`func (o *DirectoryServiceSearchResponse) GetSearchResultList() []Entity`

GetSearchResultList returns the SearchResultList field if non-nil, zero value otherwise.

### GetSearchResultListOk

`func (o *DirectoryServiceSearchResponse) GetSearchResultListOk() (*[]Entity, bool)`

GetSearchResultListOk returns a tuple with the SearchResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultList

`func (o *DirectoryServiceSearchResponse) SetSearchResultList(v []Entity)`

SetSearchResultList sets SearchResultList field to given value.

### HasSearchResultList

`func (o *DirectoryServiceSearchResponse) HasSearchResultList() bool`

HasSearchResultList returns a boolean if a field has been set.

### GetDomainName

`func (o *DirectoryServiceSearchResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DirectoryServiceSearchResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DirectoryServiceSearchResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DirectoryServiceSearchResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectoryServiceSearchResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectoryServiceSearchResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectoryServiceSearchResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectoryServiceSearchResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectoryServiceSearchResponse) GetMetadata() DirectoryServiceSearchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectoryServiceSearchResponse) GetMetadataOk() (*DirectoryServiceSearchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectoryServiceSearchResponse) SetMetadata(v DirectoryServiceSearchMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DirectoryServiceSearchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


