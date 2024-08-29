# ProjectIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ProjectDefStatus**](ProjectDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Project**](Project.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ProjectMetadata**](ProjectMetadata.md) |  | 

## Methods

### NewProjectIntentResource

`func NewProjectIntentResource(metadata ProjectMetadata, ) *ProjectIntentResource`

NewProjectIntentResource instantiates a new ProjectIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIntentResourceWithDefaults

`func NewProjectIntentResourceWithDefaults() *ProjectIntentResource`

NewProjectIntentResourceWithDefaults instantiates a new ProjectIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectIntentResource) GetStatus() ProjectDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectIntentResource) GetStatusOk() (*ProjectDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectIntentResource) SetStatus(v ProjectDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ProjectIntentResource) GetSpec() Project`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProjectIntentResource) GetSpecOk() (*Project, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProjectIntentResource) SetSpec(v Project)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ProjectIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProjectIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProjectIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProjectIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ProjectIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ProjectIntentResource) GetMetadata() ProjectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectIntentResource) GetMetadataOk() (*ProjectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectIntentResource) SetMetadata(v ProjectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


