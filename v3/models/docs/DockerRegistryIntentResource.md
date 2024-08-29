# DockerRegistryIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DockerRegistryDefStatus**](DockerRegistryDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DockerRegistry**](DockerRegistry.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DockerRegistryMetadata**](DockerRegistryMetadata.md) |  | 

## Methods

### NewDockerRegistryIntentResource

`func NewDockerRegistryIntentResource(metadata DockerRegistryMetadata, ) *DockerRegistryIntentResource`

NewDockerRegistryIntentResource instantiates a new DockerRegistryIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerRegistryIntentResourceWithDefaults

`func NewDockerRegistryIntentResourceWithDefaults() *DockerRegistryIntentResource`

NewDockerRegistryIntentResourceWithDefaults instantiates a new DockerRegistryIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DockerRegistryIntentResource) GetStatus() DockerRegistryDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DockerRegistryIntentResource) GetStatusOk() (*DockerRegistryDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DockerRegistryIntentResource) SetStatus(v DockerRegistryDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DockerRegistryIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DockerRegistryIntentResource) GetSpec() DockerRegistry`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DockerRegistryIntentResource) GetSpecOk() (*DockerRegistry, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DockerRegistryIntentResource) SetSpec(v DockerRegistry)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DockerRegistryIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DockerRegistryIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DockerRegistryIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DockerRegistryIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DockerRegistryIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DockerRegistryIntentResource) GetMetadata() DockerRegistryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DockerRegistryIntentResource) GetMetadataOk() (*DockerRegistryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DockerRegistryIntentResource) SetMetadata(v DockerRegistryMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


