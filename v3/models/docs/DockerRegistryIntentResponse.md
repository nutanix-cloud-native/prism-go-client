# DockerRegistryIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DockerRegistryDefStatus**](DockerRegistryDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DockerRegistry**](DockerRegistry.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DockerRegistryMetadata**](DockerRegistryMetadata.md) |  | 

## Methods

### NewDockerRegistryIntentResponse

`func NewDockerRegistryIntentResponse(apiVersion string, metadata DockerRegistryMetadata, ) *DockerRegistryIntentResponse`

NewDockerRegistryIntentResponse instantiates a new DockerRegistryIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryIntentResponseWithDefaults

`func NewDockerRegistryIntentResponseWithDefaults() *DockerRegistryIntentResponse`

NewDockerRegistryIntentResponseWithDefaults instantiates a new DockerRegistryIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DockerRegistryIntentResponse) GetStatus() DockerRegistryDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DockerRegistryIntentResponse) GetStatusOk() (*DockerRegistryDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DockerRegistryIntentResponse) SetStatus(v DockerRegistryDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DockerRegistryIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DockerRegistryIntentResponse) GetSpec() DockerRegistry`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DockerRegistryIntentResponse) GetSpecOk() (*DockerRegistry, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DockerRegistryIntentResponse) SetSpec(v DockerRegistry)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DockerRegistryIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DockerRegistryIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DockerRegistryIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DockerRegistryIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DockerRegistryIntentResponse) GetMetadata() DockerRegistryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DockerRegistryIntentResponse) GetMetadataOk() (*DockerRegistryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DockerRegistryIntentResponse) SetMetadata(v DockerRegistryMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


