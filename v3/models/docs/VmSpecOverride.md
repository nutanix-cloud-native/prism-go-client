# VmSpecOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldReconfigureNutanixGuestTools** | Pointer to **bool** | Whether to reconfigure NGT inside the guest VM if it was installed at the time of snapshot.  | [optional] [default to true]
**Name** | Pointer to **string** | VM Name. | [optional] 
**Resources** | Pointer to [**VMResources1**](VMResources1.md) |  | [optional] 

## Methods

### NewVmSpecOverride

`func NewVmSpecOverride() *VmSpecOverride`

NewVmSpecOverride instantiates a new VmSpecOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSpecOverrideWithDefaults

`func NewVmSpecOverrideWithDefaults() *VmSpecOverride`

NewVmSpecOverrideWithDefaults instantiates a new VmSpecOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldReconfigureNutanixGuestTools

`func (o *VmSpecOverride) GetShouldReconfigureNutanixGuestTools() bool`

GetShouldReconfigureNutanixGuestTools returns the ShouldReconfigureNutanixGuestTools field if non-nil, zero value otherwise.

### GetShouldReconfigureNutanixGuestToolsOk

`func (o *VmSpecOverride) GetShouldReconfigureNutanixGuestToolsOk() (*bool, bool)`

GetShouldReconfigureNutanixGuestToolsOk returns a tuple with the ShouldReconfigureNutanixGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReconfigureNutanixGuestTools

`func (o *VmSpecOverride) SetShouldReconfigureNutanixGuestTools(v bool)`

SetShouldReconfigureNutanixGuestTools sets ShouldReconfigureNutanixGuestTools field to given value.

### HasShouldReconfigureNutanixGuestTools

`func (o *VmSpecOverride) HasShouldReconfigureNutanixGuestTools() bool`

HasShouldReconfigureNutanixGuestTools returns a boolean if a field has been set.

### GetName

`func (o *VmSpecOverride) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmSpecOverride) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmSpecOverride) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmSpecOverride) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *VmSpecOverride) GetResources() VMResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VmSpecOverride) GetResourcesOk() (*VMResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VmSpecOverride) SetResources(v VMResources1)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VmSpecOverride) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


