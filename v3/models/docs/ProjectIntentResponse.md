# ProjectIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ProjectDefStatus**](ProjectDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Project**](Project.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ProjectMetadata**](ProjectMetadata.md) |  | 

## Methods

### NewProjectIntentResponse

`func NewProjectIntentResponse(apiVersion string, metadata ProjectMetadata, ) *ProjectIntentResponse`

NewProjectIntentResponse instantiates a new ProjectIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIntentResponseWithDefaults

`func NewProjectIntentResponseWithDefaults() *ProjectIntentResponse`

NewProjectIntentResponseWithDefaults instantiates a new ProjectIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectIntentResponse) GetStatus() ProjectDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectIntentResponse) GetStatusOk() (*ProjectDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectIntentResponse) SetStatus(v ProjectDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ProjectIntentResponse) GetSpec() Project`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProjectIntentResponse) GetSpecOk() (*Project, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProjectIntentResponse) SetSpec(v Project)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ProjectIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProjectIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProjectIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProjectIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ProjectIntentResponse) GetMetadata() ProjectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectIntentResponse) GetMetadataOk() (*ProjectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectIntentResponse) SetMetadata(v ProjectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


