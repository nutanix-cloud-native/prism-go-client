# MhVmNicOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UUID** | Pointer to **string** | UUID of the Virtual NIC. | [optional] 
**AdapterType** | Pointer to **string** | Adapter type. | [optional] 
**MacAddressType** | Pointer to **string** | The MAC address type for the Virtual NIC. | [optional] 
**MacAddress** | Pointer to **string** | The MAC address for the Virtual NIC. | [optional] 
**SubnetReference** | Pointer to [**SubnetReference**](SubnetReference.md) |  | [optional] 
**IsConnected** | Pointer to **bool** | Whether or not the Virtual NIC is connected. | [optional] [default to true]

## Methods

### NewMhVmNicOverride

`func NewMhVmNicOverride() *MhVmNicOverride`

NewMhVmNicOverride instantiates a new MhVmNicOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmNicOverrideWithDefaults

`func NewMhVmNicOverrideWithDefaults() *MhVmNicOverride`

NewMhVmNicOverrideWithDefaults instantiates a new MhVmNicOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUUID

`func (o *MhVmNicOverride) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *MhVmNicOverride) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *MhVmNicOverride) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *MhVmNicOverride) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetAdapterType

`func (o *MhVmNicOverride) GetAdapterType() string`

GetAdapterType returns the AdapterType field if non-nil, zero value otherwise.

### GetAdapterTypeOk

`func (o *MhVmNicOverride) GetAdapterTypeOk() (*string, bool)`

GetAdapterTypeOk returns a tuple with the AdapterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterType

`func (o *MhVmNicOverride) SetAdapterType(v string)`

SetAdapterType sets AdapterType field to given value.

### HasAdapterType

`func (o *MhVmNicOverride) HasAdapterType() bool`

HasAdapterType returns a boolean if a field has been set.

### GetMacAddressType

`func (o *MhVmNicOverride) GetMacAddressType() string`

GetMacAddressType returns the MacAddressType field if non-nil, zero value otherwise.

### GetMacAddressTypeOk

`func (o *MhVmNicOverride) GetMacAddressTypeOk() (*string, bool)`

GetMacAddressTypeOk returns a tuple with the MacAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressType

`func (o *MhVmNicOverride) SetMacAddressType(v string)`

SetMacAddressType sets MacAddressType field to given value.

### HasMacAddressType

`func (o *MhVmNicOverride) HasMacAddressType() bool`

HasMacAddressType returns a boolean if a field has been set.

### GetMacAddress

`func (o *MhVmNicOverride) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MhVmNicOverride) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MhVmNicOverride) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *MhVmNicOverride) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetSubnetReference

`func (o *MhVmNicOverride) GetSubnetReference() SubnetReference`

GetSubnetReference returns the SubnetReference field if non-nil, zero value otherwise.

### GetSubnetReferenceOk

`func (o *MhVmNicOverride) GetSubnetReferenceOk() (*SubnetReference, bool)`

GetSubnetReferenceOk returns a tuple with the SubnetReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetReference

`func (o *MhVmNicOverride) SetSubnetReference(v SubnetReference)`

SetSubnetReference sets SubnetReference field to given value.

### HasSubnetReference

`func (o *MhVmNicOverride) HasSubnetReference() bool`

HasSubnetReference returns a boolean if a field has been set.

### GetIsConnected

`func (o *MhVmNicOverride) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *MhVmNicOverride) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *MhVmNicOverride) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *MhVmNicOverride) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


