# NetworkDeviceListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]NetworkDeviceIntentResource**](NetworkDeviceIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkDeviceListMetadataOutput**](NetworkDeviceListMetadataOutput.md) |  | 

## Methods

### NewNetworkDeviceListIntentResponse

`func NewNetworkDeviceListIntentResponse(apiVersion string, metadata NetworkDeviceListMetadataOutput, ) *NetworkDeviceListIntentResponse`

NewNetworkDeviceListIntentResponse instantiates a new NetworkDeviceListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceListIntentResponseWithDefaults

`func NewNetworkDeviceListIntentResponseWithDefaults() *NetworkDeviceListIntentResponse`

NewNetworkDeviceListIntentResponseWithDefaults instantiates a new NetworkDeviceListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *NetworkDeviceListIntentResponse) GetEntities() []NetworkDeviceIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NetworkDeviceListIntentResponse) GetEntitiesOk() (*[]NetworkDeviceIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NetworkDeviceListIntentResponse) SetEntities(v []NetworkDeviceIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *NetworkDeviceListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkDeviceListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkDeviceListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkDeviceListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NetworkDeviceListIntentResponse) GetMetadata() NetworkDeviceListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkDeviceListIntentResponse) GetMetadataOk() (*NetworkDeviceListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkDeviceListIntentResponse) SetMetadata(v NetworkDeviceListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


