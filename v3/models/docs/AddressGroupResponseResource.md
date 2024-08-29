# AddressGroupResponseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressGroup** | Pointer to [**AddressGroup**](AddressGroup.md) |  | [optional] 
**AssociatedPoliciesList** | Pointer to [**[]Reference**](Reference.md) | The policies where the address_group is being used | [optional] 

## Methods

### NewAddressGroupResponseResource

`func NewAddressGroupResponseResource() *AddressGroupResponseResource`

NewAddressGroupResponseResource instantiates a new AddressGroupResponseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressGroupResponseResourceWithDefaults

`func NewAddressGroupResponseResourceWithDefaults() *AddressGroupResponseResource`

NewAddressGroupResponseResourceWithDefaults instantiates a new AddressGroupResponseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressGroup

`func (o *AddressGroupResponseResource) GetAddressGroup() AddressGroup`

GetAddressGroup returns the AddressGroup field if non-nil, zero value otherwise.

### GetAddressGroupOk

`func (o *AddressGroupResponseResource) GetAddressGroupOk() (*AddressGroup, bool)`

GetAddressGroupOk returns a tuple with the AddressGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressGroup

`func (o *AddressGroupResponseResource) SetAddressGroup(v AddressGroup)`

SetAddressGroup sets AddressGroup field to given value.

### HasAddressGroup

`func (o *AddressGroupResponseResource) HasAddressGroup() bool`

HasAddressGroup returns a boolean if a field has been set.

### GetAssociatedPoliciesList

`func (o *AddressGroupResponseResource) GetAssociatedPoliciesList() []Reference`

GetAssociatedPoliciesList returns the AssociatedPoliciesList field if non-nil, zero value otherwise.

### GetAssociatedPoliciesListOk

`func (o *AddressGroupResponseResource) GetAssociatedPoliciesListOk() (*[]Reference, bool)`

GetAssociatedPoliciesListOk returns a tuple with the AssociatedPoliciesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPoliciesList

`func (o *AddressGroupResponseResource) SetAssociatedPoliciesList(v []Reference)`

SetAssociatedPoliciesList sets AssociatedPoliciesList field to given value.

### HasAssociatedPoliciesList

`func (o *AddressGroupResponseResource) HasAssociatedPoliciesList() bool`

HasAssociatedPoliciesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


