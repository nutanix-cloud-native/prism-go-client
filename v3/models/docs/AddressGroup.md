# AddressGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressGroupString** | Pointer to **string** | List of original addresses input. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IpAddressBlockList** | Pointer to [**[]IpSubnet**](IpSubnet.md) | list of subnets and CIDR blocks in the address group | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressGroup

`func NewAddressGroup() *AddressGroup`

NewAddressGroup instantiates a new AddressGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGroupWithDefaults

`func NewAddressGroupWithDefaults() *AddressGroup`

NewAddressGroupWithDefaults instantiates a new AddressGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressGroupString

`func (o *AddressGroup) GetAddressGroupString() string`

GetAddressGroupString returns the AddressGroupString field if non-nil, zero value otherwise.

### GetAddressGroupStringOk

`func (o *AddressGroup) GetAddressGroupStringOk() (*string, bool)`

GetAddressGroupStringOk returns a tuple with the AddressGroupString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressGroupString

`func (o *AddressGroup) SetAddressGroupString(v string)`

SetAddressGroupString sets AddressGroupString field to given value.

### HasAddressGroupString

`func (o *AddressGroup) HasAddressGroupString() bool`

HasAddressGroupString returns a boolean if a field has been set.

### GetName

`func (o *AddressGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpAddressBlockList

`func (o *AddressGroup) GetIpAddressBlockList() []IpSubnet`

GetIpAddressBlockList returns the IpAddressBlockList field if non-nil, zero value otherwise.

### GetIpAddressBlockListOk

`func (o *AddressGroup) GetIpAddressBlockListOk() (*[]IpSubnet, bool)`

GetIpAddressBlockListOk returns a tuple with the IpAddressBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressBlockList

`func (o *AddressGroup) SetIpAddressBlockList(v []IpSubnet)`

SetIpAddressBlockList sets IpAddressBlockList field to given value.

### HasIpAddressBlockList

`func (o *AddressGroup) HasIpAddressBlockList() bool`

HasIpAddressBlockList returns a boolean if a field has been set.

### GetDescription

`func (o *AddressGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddressGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddressGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddressGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


