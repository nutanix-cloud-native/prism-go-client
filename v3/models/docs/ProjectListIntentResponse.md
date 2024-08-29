# ProjectListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ProjectIntentResource**](ProjectIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ProjectListMetadataOutput**](ProjectListMetadataOutput.md) |  | 

## Methods

### NewProjectListIntentResponse

`func NewProjectListIntentResponse(apiVersion string, metadata ProjectListMetadataOutput, ) *ProjectListIntentResponse`

NewProjectListIntentResponse instantiates a new ProjectListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListIntentResponseWithDefaults

`func NewProjectListIntentResponseWithDefaults() *ProjectListIntentResponse`

NewProjectListIntentResponseWithDefaults instantiates a new ProjectListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ProjectListIntentResponse) GetEntities() []ProjectIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ProjectListIntentResponse) GetEntitiesOk() (*[]ProjectIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ProjectListIntentResponse) SetEntities(v []ProjectIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ProjectListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProjectListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProjectListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProjectListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ProjectListIntentResponse) GetMetadata() ProjectListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectListIntentResponse) GetMetadataOk() (*ProjectListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectListIntentResponse) SetMetadata(v ProjectListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


