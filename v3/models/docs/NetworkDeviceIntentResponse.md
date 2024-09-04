# NetworkDeviceIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NetworkDeviceDefStatus**](NetworkDeviceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**NetworkDevice**](NetworkDevice.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkDeviceMetadata**](NetworkDeviceMetadata.md) |  | 

## Methods

### NewNetworkDeviceIntentResponse

`func NewNetworkDeviceIntentResponse(apiVersion string, metadata NetworkDeviceMetadata, ) *NetworkDeviceIntentResponse`

NewNetworkDeviceIntentResponse instantiates a new NetworkDeviceIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceIntentResponseWithDefaults

`func NewNetworkDeviceIntentResponseWithDefaults() *NetworkDeviceIntentResponse`

NewNetworkDeviceIntentResponseWithDefaults instantiates a new NetworkDeviceIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NetworkDeviceIntentResponse) GetStatus() NetworkDeviceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDeviceIntentResponse) GetStatusOk() (*NetworkDeviceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDeviceIntentResponse) SetStatus(v NetworkDeviceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkDeviceIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkDeviceIntentResponse) GetSpec() NetworkDevice`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkDeviceIntentResponse) GetSpecOk() (*NetworkDevice, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkDeviceIntentResponse) SetSpec(v NetworkDevice)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkDeviceIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkDeviceIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkDeviceIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkDeviceIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NetworkDeviceIntentResponse) GetMetadata() NetworkDeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkDeviceIntentResponse) GetMetadataOk() (*NetworkDeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkDeviceIntentResponse) SetMetadata(v NetworkDeviceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


