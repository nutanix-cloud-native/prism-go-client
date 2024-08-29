# NetworkDeviceIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NetworkDeviceDefStatus**](NetworkDeviceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**NetworkDevice**](NetworkDevice.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkDeviceMetadata**](NetworkDeviceMetadata.md) |  | 

## Methods

### NewNetworkDeviceIntentResource

`func NewNetworkDeviceIntentResource(metadata NetworkDeviceMetadata, ) *NetworkDeviceIntentResource`

NewNetworkDeviceIntentResource instantiates a new NetworkDeviceIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceIntentResourceWithDefaults

`func NewNetworkDeviceIntentResourceWithDefaults() *NetworkDeviceIntentResource`

NewNetworkDeviceIntentResourceWithDefaults instantiates a new NetworkDeviceIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NetworkDeviceIntentResource) GetStatus() NetworkDeviceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDeviceIntentResource) GetStatusOk() (*NetworkDeviceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDeviceIntentResource) SetStatus(v NetworkDeviceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkDeviceIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkDeviceIntentResource) GetSpec() NetworkDevice`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkDeviceIntentResource) GetSpecOk() (*NetworkDevice, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkDeviceIntentResource) SetSpec(v NetworkDevice)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkDeviceIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkDeviceIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkDeviceIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkDeviceIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkDeviceIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkDeviceIntentResource) GetMetadata() NetworkDeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkDeviceIntentResource) GetMetadataOk() (*NetworkDeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkDeviceIntentResource) SetMetadata(v NetworkDeviceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


