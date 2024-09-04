# RoutingPolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]RoutingPolicyIntentResource**](RoutingPolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RoutingPolicyListMetadataOutput**](RoutingPolicyListMetadataOutput.md) |  | 

## Methods

### NewRoutingPolicyListIntentResponse

`func NewRoutingPolicyListIntentResponse(apiVersion string, metadata RoutingPolicyListMetadataOutput, ) *RoutingPolicyListIntentResponse`

NewRoutingPolicyListIntentResponse instantiates a new RoutingPolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyListIntentResponseWithDefaults

`func NewRoutingPolicyListIntentResponseWithDefaults() *RoutingPolicyListIntentResponse`

NewRoutingPolicyListIntentResponseWithDefaults instantiates a new RoutingPolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RoutingPolicyListIntentResponse) GetEntities() []RoutingPolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RoutingPolicyListIntentResponse) GetEntitiesOk() (*[]RoutingPolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RoutingPolicyListIntentResponse) SetEntities(v []RoutingPolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RoutingPolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoutingPolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingPolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingPolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RoutingPolicyListIntentResponse) GetMetadata() RoutingPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoutingPolicyListIntentResponse) GetMetadataOk() (*RoutingPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoutingPolicyListIntentResponse) SetMetadata(v RoutingPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


