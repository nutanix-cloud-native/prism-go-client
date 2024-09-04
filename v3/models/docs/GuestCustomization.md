# GuestCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInit** | Pointer to [**GuestCustomizationCloudInit**](GuestCustomizationCloudInit.md) |  | [optional] 
**IsOverridable** | Pointer to **bool** | Flag to allow override of customization by deployer. | [optional] [default to false]
**Sysprep** | Pointer to [**GuestCustomizationSysprep**](GuestCustomizationSysprep.md) |  | [optional] 

## Methods

### NewGuestCustomization

`func NewGuestCustomization() *GuestCustomization`

NewGuestCustomization instantiates a new GuestCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestCustomizationWithDefaults

`func NewGuestCustomizationWithDefaults() *GuestCustomization`

NewGuestCustomizationWithDefaults instantiates a new GuestCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInit

`func (o *GuestCustomization) GetCloudInit() GuestCustomizationCloudInit`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *GuestCustomization) GetCloudInitOk() (*GuestCustomizationCloudInit, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *GuestCustomization) SetCloudInit(v GuestCustomizationCloudInit)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *GuestCustomization) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.

### GetIsOverridable

`func (o *GuestCustomization) GetIsOverridable() bool`

GetIsOverridable returns the IsOverridable field if non-nil, zero value otherwise.

### GetIsOverridableOk

`func (o *GuestCustomization) GetIsOverridableOk() (*bool, bool)`

GetIsOverridableOk returns a tuple with the IsOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridable

`func (o *GuestCustomization) SetIsOverridable(v bool)`

SetIsOverridable sets IsOverridable field to given value.

### HasIsOverridable

`func (o *GuestCustomization) HasIsOverridable() bool`

HasIsOverridable returns a boolean if a field has been set.

### GetSysprep

`func (o *GuestCustomization) GetSysprep() GuestCustomizationSysprep`

GetSysprep returns the Sysprep field if non-nil, zero value otherwise.

### GetSysprepOk

`func (o *GuestCustomization) GetSysprepOk() (*GuestCustomizationSysprep, bool)`

GetSysprepOk returns a tuple with the Sysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprep

`func (o *GuestCustomization) SetSysprep(v GuestCustomizationSysprep)`

SetSysprep sets Sysprep field to given value.

### HasSysprep

`func (o *GuestCustomization) HasSysprep() bool`

HasSysprep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


