# MhVmSpecOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldReconfigureNutanixGuestTools** | Pointer to **bool** | Whether to reconfigure NGT inside the guest VM if it was installed at the time of snapshot.  | [optional] [default to true]
**Name** | Pointer to **string** | VM name. | [optional] 
**Resources** | Pointer to [**VMResources**](VMResources.md) |  | [optional] 

## Methods

### NewMhVmSpecOverride

`func NewMhVmSpecOverride() *MhVmSpecOverride`

NewMhVmSpecOverride instantiates a new MhVmSpecOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmSpecOverrideWithDefaults

`func NewMhVmSpecOverrideWithDefaults() *MhVmSpecOverride`

NewMhVmSpecOverrideWithDefaults instantiates a new MhVmSpecOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldReconfigureNutanixGuestTools

`func (o *MhVmSpecOverride) GetShouldReconfigureNutanixGuestTools() bool`

GetShouldReconfigureNutanixGuestTools returns the ShouldReconfigureNutanixGuestTools field if non-nil, zero value otherwise.

### GetShouldReconfigureNutanixGuestToolsOk

`func (o *MhVmSpecOverride) GetShouldReconfigureNutanixGuestToolsOk() (*bool, bool)`

GetShouldReconfigureNutanixGuestToolsOk returns a tuple with the ShouldReconfigureNutanixGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReconfigureNutanixGuestTools

`func (o *MhVmSpecOverride) SetShouldReconfigureNutanixGuestTools(v bool)`

SetShouldReconfigureNutanixGuestTools sets ShouldReconfigureNutanixGuestTools field to given value.

### HasShouldReconfigureNutanixGuestTools

`func (o *MhVmSpecOverride) HasShouldReconfigureNutanixGuestTools() bool`

HasShouldReconfigureNutanixGuestTools returns a boolean if a field has been set.

### GetName

`func (o *MhVmSpecOverride) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MhVmSpecOverride) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MhVmSpecOverride) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MhVmSpecOverride) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *MhVmSpecOverride) GetResources() VMResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MhVmSpecOverride) GetResourcesOk() (*VMResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MhVmSpecOverride) SetResources(v VMResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MhVmSpecOverride) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


