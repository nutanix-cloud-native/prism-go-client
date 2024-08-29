# VpnConnectionListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VpnConnectionIntentResource**](VpnConnectionIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**VpnConnectionListMetadataOutput**](VpnConnectionListMetadataOutput.md) |  | 

## Methods

### NewVpnConnectionListIntentResponse

`func NewVpnConnectionListIntentResponse(apiVersion string, metadata VpnConnectionListMetadataOutput, ) *VpnConnectionListIntentResponse`

NewVpnConnectionListIntentResponse instantiates a new VpnConnectionListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnConnectionListIntentResponseWithDefaults

`func NewVpnConnectionListIntentResponseWithDefaults() *VpnConnectionListIntentResponse`

NewVpnConnectionListIntentResponseWithDefaults instantiates a new VpnConnectionListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *VpnConnectionListIntentResponse) GetEntities() []VpnConnectionIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *VpnConnectionListIntentResponse) GetEntitiesOk() (*[]VpnConnectionIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *VpnConnectionListIntentResponse) SetEntities(v []VpnConnectionIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *VpnConnectionListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *VpnConnectionListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpnConnectionListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpnConnectionListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *VpnConnectionListIntentResponse) GetMetadata() VpnConnectionListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpnConnectionListIntentResponse) GetMetadataOk() (*VpnConnectionListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpnConnectionListIntentResponse) SetMetadata(v VpnConnectionListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


