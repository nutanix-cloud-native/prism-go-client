# ActionTemplateListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ActionTemplateIntentResource**](ActionTemplateIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTemplateListMetadataOutput**](ActionTemplateListMetadataOutput.md) |  | 

## Methods

### NewActionTemplateListIntentResponse

`func NewActionTemplateListIntentResponse(apiVersion string, metadata ActionTemplateListMetadataOutput, ) *ActionTemplateListIntentResponse`

NewActionTemplateListIntentResponse instantiates a new ActionTemplateListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTemplateListIntentResponseWithDefaults

`func NewActionTemplateListIntentResponseWithDefaults() *ActionTemplateListIntentResponse`

NewActionTemplateListIntentResponseWithDefaults instantiates a new ActionTemplateListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ActionTemplateListIntentResponse) GetEntities() []ActionTemplateIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ActionTemplateListIntentResponse) GetEntitiesOk() (*[]ActionTemplateIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ActionTemplateListIntentResponse) SetEntities(v []ActionTemplateIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ActionTemplateListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTemplateListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTemplateListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTemplateListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionTemplateListIntentResponse) GetMetadata() ActionTemplateListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTemplateListIntentResponse) GetMetadataOk() (*ActionTemplateListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTemplateListIntentResponse) SetMetadata(v ActionTemplateListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


