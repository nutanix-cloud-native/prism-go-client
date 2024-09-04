# VmPowerStateMechanism

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestTransitionConfig** | Pointer to [**VmGuestPowerStateTransitionConfig**](VmGuestPowerStateTransitionConfig.md) |  | [optional] 
**Mechanism** | Pointer to **string** | Power state mechanism (ACPI/GUEST/HARD). | [optional] 

## Methods

### NewVmPowerStateMechanism

`func NewVmPowerStateMechanism() *VmPowerStateMechanism`

NewVmPowerStateMechanism instantiates a new VmPowerStateMechanism object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmPowerStateMechanismWithDefaults

`func NewVmPowerStateMechanismWithDefaults() *VmPowerStateMechanism`

NewVmPowerStateMechanismWithDefaults instantiates a new VmPowerStateMechanism object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestTransitionConfig

`func (o *VmPowerStateMechanism) GetGuestTransitionConfig() VmGuestPowerStateTransitionConfig`

GetGuestTransitionConfig returns the GuestTransitionConfig field if non-nil, zero value otherwise.

### GetGuestTransitionConfigOk

`func (o *VmPowerStateMechanism) GetGuestTransitionConfigOk() (*VmGuestPowerStateTransitionConfig, bool)`

GetGuestTransitionConfigOk returns a tuple with the GuestTransitionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTransitionConfig

`func (o *VmPowerStateMechanism) SetGuestTransitionConfig(v VmGuestPowerStateTransitionConfig)`

SetGuestTransitionConfig sets GuestTransitionConfig field to given value.

### HasGuestTransitionConfig

`func (o *VmPowerStateMechanism) HasGuestTransitionConfig() bool`

HasGuestTransitionConfig returns a boolean if a field has been set.

### GetMechanism

`func (o *VmPowerStateMechanism) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *VmPowerStateMechanism) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *VmPowerStateMechanism) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *VmPowerStateMechanism) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


