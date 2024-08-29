# ProjectInternalIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ProjectInternal**](ProjectInternal.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ProjectMetadata**](ProjectMetadata.md) |  | 

## Methods

### NewProjectInternalIntentInput

`func NewProjectInternalIntentInput(spec ProjectInternal, apiVersion string, metadata ProjectMetadata, ) *ProjectInternalIntentInput`

NewProjectInternalIntentInput instantiates a new ProjectInternalIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalIntentInputWithDefaults

`func NewProjectInternalIntentInputWithDefaults() *ProjectInternalIntentInput`

NewProjectInternalIntentInputWithDefaults instantiates a new ProjectInternalIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ProjectInternalIntentInput) GetSpec() ProjectInternal`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProjectInternalIntentInput) GetSpecOk() (*ProjectInternal, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProjectInternalIntentInput) SetSpec(v ProjectInternal)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ProjectInternalIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProjectInternalIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProjectInternalIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ProjectInternalIntentInput) GetMetadata() ProjectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectInternalIntentInput) GetMetadataOk() (*ProjectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectInternalIntentInput) SetMetadata(v ProjectMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


