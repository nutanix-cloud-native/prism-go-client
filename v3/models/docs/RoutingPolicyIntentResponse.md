# RoutingPolicyIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RoutingPolicyDefStatus**](RoutingPolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RoutingPolicyMetadata**](RoutingPolicyMetadata.md) |  | 

## Methods

### NewRoutingPolicyIntentResponse

`func NewRoutingPolicyIntentResponse(apiVersion string, metadata RoutingPolicyMetadata, ) *RoutingPolicyIntentResponse`

NewRoutingPolicyIntentResponse instantiates a new RoutingPolicyIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyIntentResponseWithDefaults

`func NewRoutingPolicyIntentResponseWithDefaults() *RoutingPolicyIntentResponse`

NewRoutingPolicyIntentResponseWithDefaults instantiates a new RoutingPolicyIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RoutingPolicyIntentResponse) GetStatus() RoutingPolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoutingPolicyIntentResponse) GetStatusOk() (*RoutingPolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoutingPolicyIntentResponse) SetStatus(v RoutingPolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoutingPolicyIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RoutingPolicyIntentResponse) GetSpec() RoutingPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoutingPolicyIntentResponse) GetSpecOk() (*RoutingPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoutingPolicyIntentResponse) SetSpec(v RoutingPolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RoutingPolicyIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoutingPolicyIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingPolicyIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingPolicyIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RoutingPolicyIntentResponse) GetMetadata() RoutingPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoutingPolicyIntentResponse) GetMetadataOk() (*RoutingPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoutingPolicyIntentResponse) SetMetadata(v RoutingPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


