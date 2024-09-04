# VirtualNetworkIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**VirtualNetwork**](VirtualNetwork.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VirtualNetworkMetadata**](VirtualNetworkMetadata.md) |  | 

## Methods

### NewVirtualNetworkIntentInput

`func NewVirtualNetworkIntentInput(spec VirtualNetwork, metadata VirtualNetworkMetadata, ) *VirtualNetworkIntentInput`

NewVirtualNetworkIntentInput instantiates a new VirtualNetworkIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkIntentInputWithDefaults

`func NewVirtualNetworkIntentInputWithDefaults() *VirtualNetworkIntentInput`

NewVirtualNetworkIntentInputWithDefaults instantiates a new VirtualNetworkIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VirtualNetworkIntentInput) GetSpec() VirtualNetwork`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VirtualNetworkIntentInput) GetSpecOk() (*VirtualNetwork, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VirtualNetworkIntentInput) SetSpec(v VirtualNetwork)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VirtualNetworkIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VirtualNetworkIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VirtualNetworkIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VirtualNetworkIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VirtualNetworkIntentInput) GetMetadata() VirtualNetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualNetworkIntentInput) GetMetadataOk() (*VirtualNetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualNetworkIntentInput) SetMetadata(v VirtualNetworkMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


