# RoutingPolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RoutingPolicy**](RoutingPolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RoutingPolicyMetadata**](RoutingPolicyMetadata.md) |  | 

## Methods

### NewRoutingPolicyIntentInput

`func NewRoutingPolicyIntentInput(spec RoutingPolicy, metadata RoutingPolicyMetadata, ) *RoutingPolicyIntentInput`

NewRoutingPolicyIntentInput instantiates a new RoutingPolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyIntentInputWithDefaults

`func NewRoutingPolicyIntentInputWithDefaults() *RoutingPolicyIntentInput`

NewRoutingPolicyIntentInputWithDefaults instantiates a new RoutingPolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RoutingPolicyIntentInput) GetSpec() RoutingPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoutingPolicyIntentInput) GetSpecOk() (*RoutingPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoutingPolicyIntentInput) SetSpec(v RoutingPolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RoutingPolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingPolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingPolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoutingPolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RoutingPolicyIntentInput) GetMetadata() RoutingPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoutingPolicyIntentInput) GetMetadataOk() (*RoutingPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoutingPolicyIntentInput) SetMetadata(v RoutingPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


