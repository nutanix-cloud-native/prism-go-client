# ClusterListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ClusterIntentResource**](ClusterIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ClusterListMetadataOutput**](ClusterListMetadataOutput.md) |  | 

## Methods

### NewClusterListIntentResponse

`func NewClusterListIntentResponse(apiVersion string, metadata ClusterListMetadataOutput, ) *ClusterListIntentResponse`

NewClusterListIntentResponse instantiates a new ClusterListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListIntentResponseWithDefaults

`func NewClusterListIntentResponseWithDefaults() *ClusterListIntentResponse`

NewClusterListIntentResponseWithDefaults instantiates a new ClusterListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ClusterListIntentResponse) GetEntities() []ClusterIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ClusterListIntentResponse) GetEntitiesOk() (*[]ClusterIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ClusterListIntentResponse) SetEntities(v []ClusterIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ClusterListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ClusterListIntentResponse) GetMetadata() ClusterListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterListIntentResponse) GetMetadataOk() (*ClusterListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterListIntentResponse) SetMetadata(v ClusterListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


