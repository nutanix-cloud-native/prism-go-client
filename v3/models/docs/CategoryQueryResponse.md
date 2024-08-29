# CategoryQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]CategoryQueryResponseResultsInner**](CategoryQueryResponseResultsInner.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | Pointer to [**CategoryQueryResponseMetadata**](CategoryQueryResponseMetadata.md) |  | [optional] 

## Methods

### NewCategoryQueryResponse

`func NewCategoryQueryResponse() *CategoryQueryResponse`

NewCategoryQueryResponse instantiates a new CategoryQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryQueryResponseWithDefaults

`func NewCategoryQueryResponseWithDefaults() *CategoryQueryResponse`

NewCategoryQueryResponseWithDefaults instantiates a new CategoryQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CategoryQueryResponse) GetResults() []CategoryQueryResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CategoryQueryResponse) GetResultsOk() (*[]CategoryQueryResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CategoryQueryResponse) SetResults(v []CategoryQueryResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *CategoryQueryResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryQueryResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryQueryResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryQueryResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryQueryResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CategoryQueryResponse) GetMetadata() CategoryQueryResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CategoryQueryResponse) GetMetadataOk() (*CategoryQueryResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CategoryQueryResponse) SetMetadata(v CategoryQueryResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CategoryQueryResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


