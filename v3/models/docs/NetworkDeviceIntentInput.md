# NetworkDeviceIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**NetworkDevice**](NetworkDevice.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkDeviceMetadata**](NetworkDeviceMetadata.md) |  | 

## Methods

### NewNetworkDeviceIntentInput

`func NewNetworkDeviceIntentInput(spec NetworkDevice, metadata NetworkDeviceMetadata, ) *NetworkDeviceIntentInput`

NewNetworkDeviceIntentInput instantiates a new NetworkDeviceIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceIntentInputWithDefaults

`func NewNetworkDeviceIntentInputWithDefaults() *NetworkDeviceIntentInput`

NewNetworkDeviceIntentInputWithDefaults instantiates a new NetworkDeviceIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *NetworkDeviceIntentInput) GetSpec() NetworkDevice`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkDeviceIntentInput) GetSpecOk() (*NetworkDevice, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkDeviceIntentInput) SetSpec(v NetworkDevice)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *NetworkDeviceIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkDeviceIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkDeviceIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkDeviceIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkDeviceIntentInput) GetMetadata() NetworkDeviceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkDeviceIntentInput) GetMetadataOk() (*NetworkDeviceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkDeviceIntentInput) SetMetadata(v NetworkDeviceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


