# NetworkFunctionChainIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**NetworkFunctionChain**](NetworkFunctionChain.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkFunctionChainMetadata**](NetworkFunctionChainMetadata.md) |  | 

## Methods

### NewNetworkFunctionChainIntentInput

`func NewNetworkFunctionChainIntentInput(spec NetworkFunctionChain, metadata NetworkFunctionChainMetadata, ) *NetworkFunctionChainIntentInput`

NewNetworkFunctionChainIntentInput instantiates a new NetworkFunctionChainIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFunctionChainIntentInputWithDefaults

`func NewNetworkFunctionChainIntentInputWithDefaults() *NetworkFunctionChainIntentInput`

NewNetworkFunctionChainIntentInputWithDefaults instantiates a new NetworkFunctionChainIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *NetworkFunctionChainIntentInput) GetSpec() NetworkFunctionChain`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkFunctionChainIntentInput) GetSpecOk() (*NetworkFunctionChain, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkFunctionChainIntentInput) SetSpec(v NetworkFunctionChain)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *NetworkFunctionChainIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkFunctionChainIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkFunctionChainIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkFunctionChainIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkFunctionChainIntentInput) GetMetadata() NetworkFunctionChainMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkFunctionChainIntentInput) GetMetadataOk() (*NetworkFunctionChainMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkFunctionChainIntentInput) SetMetadata(v NetworkFunctionChainMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


