# VpnGatewayListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VpnGatewayIntentResource**](VpnGatewayIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpnGatewayListMetadataOutput**](VpnGatewayListMetadataOutput.md) |  | 

## Methods

### NewVpnGatewayListIntentResponse

`func NewVpnGatewayListIntentResponse(apiVersion string, metadata VpnGatewayListMetadataOutput, ) *VpnGatewayListIntentResponse`

NewVpnGatewayListIntentResponse instantiates a new VpnGatewayListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayListIntentResponseWithDefaults

`func NewVpnGatewayListIntentResponseWithDefaults() *VpnGatewayListIntentResponse`

NewVpnGatewayListIntentResponseWithDefaults instantiates a new VpnGatewayListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VpnGatewayListIntentResponse) GetEntities() []VpnGatewayIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VpnGatewayListIntentResponse) GetEntitiesOk() (*[]VpnGatewayIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VpnGatewayListIntentResponse) SetEntities(v []VpnGatewayIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VpnGatewayListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnGatewayListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnGatewayListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnGatewayListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpnGatewayListIntentResponse) GetMetadata() VpnGatewayListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnGatewayListIntentResponse) GetMetadataOk() (*VpnGatewayListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnGatewayListIntentResponse) SetMetadata(v VpnGatewayListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


