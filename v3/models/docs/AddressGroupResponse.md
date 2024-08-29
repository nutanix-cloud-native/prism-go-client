# AddressGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressGroup** | Pointer to [**AddressGroup**](AddressGroup.md) |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressGroupResponse

`func NewAddressGroupResponse() *AddressGroupResponse`

NewAddressGroupResponse instantiates a new AddressGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGroupResponseWithDefaults

`func NewAddressGroupResponseWithDefaults() *AddressGroupResponse`

NewAddressGroupResponseWithDefaults instantiates a new AddressGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressGroup

`func (o *AddressGroupResponse) GetAddressGroup() AddressGroup`

GetAddressGroup returns the AddressGroup field if non-nil, zero value otherwise.

### GetAddressGroupOk

`func (o *AddressGroupResponse) GetAddressGroupOk() (*AddressGroup, bool)`

GetAddressGroupOk returns a tuple with the AddressGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressGroup

`func (o *AddressGroupResponse) SetAddressGroup(v AddressGroup)`

SetAddressGroup sets AddressGroup field to given value.

### HasAddressGroup

`func (o *AddressGroupResponse) HasAddressGroup() bool`

HasAddressGroup returns a boolean if a field has been set.

### GetUUID

`func (o *AddressGroupResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AddressGroupResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AddressGroupResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AddressGroupResponse) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


