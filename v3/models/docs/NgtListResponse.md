# NgtListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitiesList** | Pointer to [**[]NgtResponse**](NgtResponse.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | Pointer to [**NgtListMetadataOutput**](NgtListMetadataOutput.md) |  | [optional] 

## Methods

### NewNgtListResponse

`func NewNgtListResponse() *NgtListResponse`

NewNgtListResponse instantiates a new NgtListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtListResponseWithDefaults

`func NewNgtListResponseWithDefaults() *NgtListResponse`

NewNgtListResponseWithDefaults instantiates a new NgtListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitiesList

`func (o *NgtListResponse) GetEntitiesList() []NgtResponse`

GetEntitiesList returns the EntitiesList field if non-nil, zero value otherwise.

### GetEntitiesListOk

`func (o *NgtListResponse) GetEntitiesListOk() (*[]NgtResponse, bool)`

GetEntitiesListOk returns a tuple with the EntitiesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitiesList

`func (o *NgtListResponse) SetEntitiesList(v []NgtResponse)`

SetEntitiesList sets EntitiesList field to given value.

### HasEntitiesList

`func (o *NgtListResponse) HasEntitiesList() bool`

HasEntitiesList returns a boolean if a field has been set.

### GetApiVersion

`func (o *NgtListResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NgtListResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NgtListResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NgtListResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NgtListResponse) GetMetadata() NgtListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NgtListResponse) GetMetadataOk() (*NgtListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NgtListResponse) SetMetadata(v NgtListMetadataOutput)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NgtListResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


