# VirtualNetworkIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VirtualNetworkDefStatus**](VirtualNetworkDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VirtualNetworkMetadata**](VirtualNetworkMetadata.md) |  | 

## Methods

### NewVirtualNetworkIntentResource

`func NewVirtualNetworkIntentResource(metadata VirtualNetworkMetadata, ) *VirtualNetworkIntentResource`

NewVirtualNetworkIntentResource instantiates a new VirtualNetworkIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkIntentResourceWithDefaults

`func NewVirtualNetworkIntentResourceWithDefaults() *VirtualNetworkIntentResource`

NewVirtualNetworkIntentResourceWithDefaults instantiates a new VirtualNetworkIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VirtualNetworkIntentResource) GetStatus() VirtualNetworkDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualNetworkIntentResource) GetStatusOk() (*VirtualNetworkDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualNetworkIntentResource) SetStatus(v VirtualNetworkDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualNetworkIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VirtualNetworkIntentResource) GetSpec() VirtualNetwork`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VirtualNetworkIntentResource) GetSpecOk() (*VirtualNetwork, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VirtualNetworkIntentResource) SetSpec(v VirtualNetwork)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VirtualNetworkIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VirtualNetworkIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VirtualNetworkIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VirtualNetworkIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VirtualNetworkIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VirtualNetworkIntentResource) GetMetadata() VirtualNetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualNetworkIntentResource) GetMetadataOk() (*VirtualNetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualNetworkIntentResource) SetMetadata(v VirtualNetworkMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


