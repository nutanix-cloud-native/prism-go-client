# DockerRegistryImageSearchListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityList** | Pointer to [**[]DockerRegistryImageStatus**](DockerRegistryImageStatus.md) |  | [optional] [readonly] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DockerRegistryListMetadata**](DockerRegistryListMetadata.md) |  | 

## Methods

### NewDockerRegistryImageSearchListIntentResponse

`func NewDockerRegistryImageSearchListIntentResponse(apiVersion string, metadata DockerRegistryListMetadata, ) *DockerRegistryImageSearchListIntentResponse`

NewDockerRegistryImageSearchListIntentResponse instantiates a new DockerRegistryImageSearchListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryImageSearchListIntentResponseWithDefaults

`func NewDockerRegistryImageSearchListIntentResponseWithDefaults() *DockerRegistryImageSearchListIntentResponse`

NewDockerRegistryImageSearchListIntentResponseWithDefaults instantiates a new DockerRegistryImageSearchListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityList

`func (o *DockerRegistryImageSearchListIntentResponse) GetEntityList() []DockerRegistryImageStatus`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *DockerRegistryImageSearchListIntentResponse) GetEntityListOk() (*[]DockerRegistryImageStatus, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *DockerRegistryImageSearchListIntentResponse) SetEntityList(v []DockerRegistryImageStatus)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *DockerRegistryImageSearchListIntentResponse) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.

### GetApiVersion

`func (o *DockerRegistryImageSearchListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DockerRegistryImageSearchListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DockerRegistryImageSearchListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DockerRegistryImageSearchListIntentResponse) GetMetadata() DockerRegistryListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DockerRegistryImageSearchListIntentResponse) GetMetadataOk() (*DockerRegistryListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DockerRegistryImageSearchListIntentResponse) SetMetadata(v DockerRegistryListMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


