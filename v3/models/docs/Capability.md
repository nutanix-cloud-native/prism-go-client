# Capability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityDescription** | Pointer to **string** | Optional value describing the capability | [optional] 
**SelfCapability** | Pointer to [**CapabilityDetails**](CapabilityDetails.md) |  | [optional] 
**Name** | Pointer to **string** | Name of capability | [optional] 
**CapableMembers** | Pointer to [**[]CapabilityDetails**](CapabilityDetails.md) | List of members having the capability | [optional] 

## Methods

### NewCapability

`func NewCapability() *Capability`

NewCapability instantiates a new Capability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityWithDefaults

`func NewCapabilityWithDefaults() *Capability`

NewCapabilityWithDefaults instantiates a new Capability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilityDescription

`func (o *Capability) GetCapabilityDescription() string`

GetCapabilityDescription returns the CapabilityDescription field if non-nil, zero value otherwise.

### GetCapabilityDescriptionOk

`func (o *Capability) GetCapabilityDescriptionOk() (*string, bool)`

GetCapabilityDescriptionOk returns a tuple with the CapabilityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityDescription

`func (o *Capability) SetCapabilityDescription(v string)`

SetCapabilityDescription sets CapabilityDescription field to given value.

### HasCapabilityDescription

`func (o *Capability) HasCapabilityDescription() bool`

HasCapabilityDescription returns a boolean if a field has been set.

### GetSelfCapability

`func (o *Capability) GetSelfCapability() CapabilityDetails`

GetSelfCapability returns the SelfCapability field if non-nil, zero value otherwise.

### GetSelfCapabilityOk

`func (o *Capability) GetSelfCapabilityOk() (*CapabilityDetails, bool)`

GetSelfCapabilityOk returns a tuple with the SelfCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfCapability

`func (o *Capability) SetSelfCapability(v CapabilityDetails)`

SetSelfCapability sets SelfCapability field to given value.

### HasSelfCapability

`func (o *Capability) HasSelfCapability() bool`

HasSelfCapability returns a boolean if a field has been set.

### GetName

`func (o *Capability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Capability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Capability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Capability) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCapableMembers

`func (o *Capability) GetCapableMembers() []CapabilityDetails`

GetCapableMembers returns the CapableMembers field if non-nil, zero value otherwise.

### GetCapableMembersOk

`func (o *Capability) GetCapableMembersOk() (*[]CapabilityDetails, bool)`

GetCapableMembersOk returns a tuple with the CapableMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapableMembers

`func (o *Capability) SetCapableMembers(v []CapabilityDetails)`

SetCapableMembers sets CapableMembers field to given value.

### HasCapableMembers

`func (o *Capability) HasCapableMembers() bool`

HasCapableMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


