# RoutingPolicyIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RoutingPolicyDefStatus**](RoutingPolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RoutingPolicy**](RoutingPolicy.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RoutingPolicyMetadata**](RoutingPolicyMetadata.md) |  | 

## Methods

### NewRoutingPolicyIntentResource

`func NewRoutingPolicyIntentResource(metadata RoutingPolicyMetadata, ) *RoutingPolicyIntentResource`

NewRoutingPolicyIntentResource instantiates a new RoutingPolicyIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyIntentResourceWithDefaults

`func NewRoutingPolicyIntentResourceWithDefaults() *RoutingPolicyIntentResource`

NewRoutingPolicyIntentResourceWithDefaults instantiates a new RoutingPolicyIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RoutingPolicyIntentResource) GetStatus() RoutingPolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoutingPolicyIntentResource) GetStatusOk() (*RoutingPolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoutingPolicyIntentResource) SetStatus(v RoutingPolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoutingPolicyIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RoutingPolicyIntentResource) GetSpec() RoutingPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoutingPolicyIntentResource) GetSpecOk() (*RoutingPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoutingPolicyIntentResource) SetSpec(v RoutingPolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RoutingPolicyIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoutingPolicyIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingPolicyIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingPolicyIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoutingPolicyIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RoutingPolicyIntentResource) GetMetadata() RoutingPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoutingPolicyIntentResource) GetMetadataOk() (*RoutingPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoutingPolicyIntentResource) SetMetadata(v RoutingPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


