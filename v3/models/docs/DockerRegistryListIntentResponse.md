# DockerRegistryListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]DockerRegistryIntentResource**](DockerRegistryIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DockerRegistryListMetadataOutput**](DockerRegistryListMetadataOutput.md) |  | 

## Methods

### NewDockerRegistryListIntentResponse

`func NewDockerRegistryListIntentResponse(apiVersion string, metadata DockerRegistryListMetadataOutput, ) *DockerRegistryListIntentResponse`

NewDockerRegistryListIntentResponse instantiates a new DockerRegistryListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryListIntentResponseWithDefaults

`func NewDockerRegistryListIntentResponseWithDefaults() *DockerRegistryListIntentResponse`

NewDockerRegistryListIntentResponseWithDefaults instantiates a new DockerRegistryListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *DockerRegistryListIntentResponse) GetEntities() []DockerRegistryIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DockerRegistryListIntentResponse) GetEntitiesOk() (*[]DockerRegistryIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DockerRegistryListIntentResponse) SetEntities(v []DockerRegistryIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DockerRegistryListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *DockerRegistryListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DockerRegistryListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DockerRegistryListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DockerRegistryListIntentResponse) GetMetadata() DockerRegistryListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DockerRegistryListIntentResponse) GetMetadataOk() (*DockerRegistryListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DockerRegistryListIntentResponse) SetMetadata(v DockerRegistryListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


