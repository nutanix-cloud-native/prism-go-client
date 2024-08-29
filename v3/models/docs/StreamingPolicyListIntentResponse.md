# StreamingPolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]StreamingPolicyIntentResource**](StreamingPolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**StreamingPolicyListMetadataOutput**](StreamingPolicyListMetadataOutput.md) |  | 

## Methods

### NewStreamingPolicyListIntentResponse

`func NewStreamingPolicyListIntentResponse(apiVersion string, metadata StreamingPolicyListMetadataOutput, ) *StreamingPolicyListIntentResponse`

NewStreamingPolicyListIntentResponse instantiates a new StreamingPolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyListIntentResponseWithDefaults

`func NewStreamingPolicyListIntentResponseWithDefaults() *StreamingPolicyListIntentResponse`

NewStreamingPolicyListIntentResponseWithDefaults instantiates a new StreamingPolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *StreamingPolicyListIntentResponse) GetEntities() []StreamingPolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *StreamingPolicyListIntentResponse) GetEntitiesOk() (*[]StreamingPolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *StreamingPolicyListIntentResponse) SetEntities(v []StreamingPolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *StreamingPolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *StreamingPolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StreamingPolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StreamingPolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *StreamingPolicyListIntentResponse) GetMetadata() StreamingPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamingPolicyListIntentResponse) GetMetadataOk() (*StreamingPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamingPolicyListIntentResponse) SetMetadata(v StreamingPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


