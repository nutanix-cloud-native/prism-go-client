# VirtualNetworkIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**VirtualNetworkDefStatus**](VirtualNetworkDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**VirtualNetwork**](VirtualNetwork.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VirtualNetworkMetadata**](VirtualNetworkMetadata.md) |  | 

## Methods

### NewVirtualNetworkIntentResponse

`func NewVirtualNetworkIntentResponse(apiVersion string, metadata VirtualNetworkMetadata, ) *VirtualNetworkIntentResponse`

NewVirtualNetworkIntentResponse instantiates a new VirtualNetworkIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkIntentResponseWithDefaults

`func NewVirtualNetworkIntentResponseWithDefaults() *VirtualNetworkIntentResponse`

NewVirtualNetworkIntentResponseWithDefaults instantiates a new VirtualNetworkIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VirtualNetworkIntentResponse) GetStatus() VirtualNetworkDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualNetworkIntentResponse) GetStatusOk() (*VirtualNetworkDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualNetworkIntentResponse) SetStatus(v VirtualNetworkDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualNetworkIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *VirtualNetworkIntentResponse) GetSpec() VirtualNetwork`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VirtualNetworkIntentResponse) GetSpecOk() (*VirtualNetwork, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VirtualNetworkIntentResponse) SetSpec(v VirtualNetwork)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VirtualNetworkIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *VirtualNetworkIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VirtualNetworkIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VirtualNetworkIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VirtualNetworkIntentResponse) GetMetadata() VirtualNetworkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualNetworkIntentResponse) GetMetadataOk() (*VirtualNetworkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualNetworkIntentResponse) SetMetadata(v VirtualNetworkMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


