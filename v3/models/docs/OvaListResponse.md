# OvaListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]OvaGetResponse**](OvaGetResponse.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**OvaListMetadataOutput**](OvaListMetadataOutput.md) |  | 

## Methods

### NewOvaListResponse

`func NewOvaListResponse(apiVersion string, metadata OvaListMetadataOutput, ) *OvaListResponse`

NewOvaListResponse instantiates a new OvaListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvaListResponseWithDefaults

`func NewOvaListResponseWithDefaults() *OvaListResponse`

NewOvaListResponseWithDefaults instantiates a new OvaListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *OvaListResponse) GetEntities() []OvaGetResponse`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *OvaListResponse) GetEntitiesOk() (*[]OvaGetResponse, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *OvaListResponse) SetEntities(v []OvaGetResponse)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *OvaListResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *OvaListResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OvaListResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OvaListResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *OvaListResponse) GetMetadata() OvaListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OvaListResponse) GetMetadataOk() (*OvaListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OvaListResponse) SetMetadata(v OvaListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


