# VirtualNetworkListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VirtualNetworkIntentResource**](VirtualNetworkIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VirtualNetworkListMetadataOutput**](VirtualNetworkListMetadataOutput.md) |  | 

## Methods

### NewVirtualNetworkListIntentResponse

`func NewVirtualNetworkListIntentResponse(apiVersion string, metadata VirtualNetworkListMetadataOutput, ) *VirtualNetworkListIntentResponse`

NewVirtualNetworkListIntentResponse instantiates a new VirtualNetworkListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkListIntentResponseWithDefaults

`func NewVirtualNetworkListIntentResponseWithDefaults() *VirtualNetworkListIntentResponse`

NewVirtualNetworkListIntentResponseWithDefaults instantiates a new VirtualNetworkListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VirtualNetworkListIntentResponse) GetEntities() []VirtualNetworkIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VirtualNetworkListIntentResponse) GetEntitiesOk() (*[]VirtualNetworkIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VirtualNetworkListIntentResponse) SetEntities(v []VirtualNetworkIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VirtualNetworkListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *VirtualNetworkListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VirtualNetworkListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VirtualNetworkListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VirtualNetworkListIntentResponse) GetMetadata() VirtualNetworkListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VirtualNetworkListIntentResponse) GetMetadataOk() (*VirtualNetworkListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VirtualNetworkListIntentResponse) SetMetadata(v VirtualNetworkListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


