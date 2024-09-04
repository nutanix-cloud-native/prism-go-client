# AppRunlogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AppRunlogResponse**](AppRunlogResponse.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppRunlogListMetadata**](AppRunlogListMetadata.md) |  | 

## Methods

### NewAppRunlogList

`func NewAppRunlogList(apiVersion string, metadata AppRunlogListMetadata, ) *AppRunlogList`

NewAppRunlogList instantiates a new AppRunlogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunlogListWithDefaults

`func NewAppRunlogListWithDefaults() *AppRunlogList`

NewAppRunlogListWithDefaults instantiates a new AppRunlogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AppRunlogList) GetEntities() []AppRunlogResponse`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AppRunlogList) GetEntitiesOk() (*[]AppRunlogResponse, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AppRunlogList) SetEntities(v []AppRunlogResponse)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AppRunlogList) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppRunlogList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppRunlogList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppRunlogList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppRunlogList) GetMetadata() AppRunlogListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppRunlogList) GetMetadataOk() (*AppRunlogListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppRunlogList) SetMetadata(v AppRunlogListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


