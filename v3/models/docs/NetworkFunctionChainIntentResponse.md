# NetworkFunctionChainIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NetworkFunctionChainDefStatus**](NetworkFunctionChainDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**NetworkFunctionChain**](NetworkFunctionChain.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkFunctionChainMetadata**](NetworkFunctionChainMetadata.md) |  | 

## Methods

### NewNetworkFunctionChainIntentResponse

`func NewNetworkFunctionChainIntentResponse(apiVersion string, metadata NetworkFunctionChainMetadata, ) *NetworkFunctionChainIntentResponse`

NewNetworkFunctionChainIntentResponse instantiates a new NetworkFunctionChainIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFunctionChainIntentResponseWithDefaults

`func NewNetworkFunctionChainIntentResponseWithDefaults() *NetworkFunctionChainIntentResponse`

NewNetworkFunctionChainIntentResponseWithDefaults instantiates a new NetworkFunctionChainIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NetworkFunctionChainIntentResponse) GetStatus() NetworkFunctionChainDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFunctionChainIntentResponse) GetStatusOk() (*NetworkFunctionChainDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFunctionChainIntentResponse) SetStatus(v NetworkFunctionChainDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkFunctionChainIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkFunctionChainIntentResponse) GetSpec() NetworkFunctionChain`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkFunctionChainIntentResponse) GetSpecOk() (*NetworkFunctionChain, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkFunctionChainIntentResponse) SetSpec(v NetworkFunctionChain)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkFunctionChainIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkFunctionChainIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkFunctionChainIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkFunctionChainIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NetworkFunctionChainIntentResponse) GetMetadata() NetworkFunctionChainMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkFunctionChainIntentResponse) GetMetadataOk() (*NetworkFunctionChainMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkFunctionChainIntentResponse) SetMetadata(v NetworkFunctionChainMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


