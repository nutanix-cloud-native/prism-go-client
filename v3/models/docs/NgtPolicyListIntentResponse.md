# NgtPolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]NgtPolicyIntentResource**](NgtPolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NgtPolicyListMetadataOutput**](NgtPolicyListMetadataOutput.md) |  | 

## Methods

### NewNgtPolicyListIntentResponse

`func NewNgtPolicyListIntentResponse(apiVersion string, metadata NgtPolicyListMetadataOutput, ) *NgtPolicyListIntentResponse`

NewNgtPolicyListIntentResponse instantiates a new NgtPolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgtPolicyListIntentResponseWithDefaults

`func NewNgtPolicyListIntentResponseWithDefaults() *NgtPolicyListIntentResponse`

NewNgtPolicyListIntentResponseWithDefaults instantiates a new NgtPolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *NgtPolicyListIntentResponse) GetEntities() []NgtPolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NgtPolicyListIntentResponse) GetEntitiesOk() (*[]NgtPolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NgtPolicyListIntentResponse) SetEntities(v []NgtPolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *NgtPolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *NgtPolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NgtPolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NgtPolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NgtPolicyListIntentResponse) GetMetadata() NgtPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NgtPolicyListIntentResponse) GetMetadataOk() (*NgtPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NgtPolicyListIntentResponse) SetMetadata(v NgtPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


