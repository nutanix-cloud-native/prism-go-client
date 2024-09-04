# ImagePlacementPolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ImagePlacementPolicyIntentResource**](ImagePlacementPolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ImagePlacementPolicyListMetadataOutput**](ImagePlacementPolicyListMetadataOutput.md) |  | 

## Methods

### NewImagePlacementPolicyListIntentResponse

`func NewImagePlacementPolicyListIntentResponse(apiVersion string, metadata ImagePlacementPolicyListMetadataOutput, ) *ImagePlacementPolicyListIntentResponse`

NewImagePlacementPolicyListIntentResponse instantiates a new ImagePlacementPolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyListIntentResponseWithDefaults

`func NewImagePlacementPolicyListIntentResponseWithDefaults() *ImagePlacementPolicyListIntentResponse`

NewImagePlacementPolicyListIntentResponseWithDefaults instantiates a new ImagePlacementPolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ImagePlacementPolicyListIntentResponse) GetEntities() []ImagePlacementPolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ImagePlacementPolicyListIntentResponse) GetEntitiesOk() (*[]ImagePlacementPolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ImagePlacementPolicyListIntentResponse) SetEntities(v []ImagePlacementPolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ImagePlacementPolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ImagePlacementPolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImagePlacementPolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImagePlacementPolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ImagePlacementPolicyListIntentResponse) GetMetadata() ImagePlacementPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImagePlacementPolicyListIntentResponse) GetMetadataOk() (*ImagePlacementPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImagePlacementPolicyListIntentResponse) SetMetadata(v ImagePlacementPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


