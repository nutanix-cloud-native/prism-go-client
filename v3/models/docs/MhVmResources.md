# MhVmResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageConfig** | Pointer to [**MhVmStorageConfig**](MhVmStorageConfig.md) |  | [optional] 
**ParentReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**GuestTools** | Pointer to [**GuestToolsSpec**](GuestToolsSpec.md) |  | [optional] 

## Methods

### NewMhVmResources

`func NewMhVmResources() *MhVmResources`

NewMhVmResources instantiates a new MhVmResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmResourcesWithDefaults

`func NewMhVmResourcesWithDefaults() *MhVmResources`

NewMhVmResourcesWithDefaults instantiates a new MhVmResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageConfig

`func (o *MhVmResources) GetStorageConfig() MhVmStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *MhVmResources) GetStorageConfigOk() (*MhVmStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *MhVmResources) SetStorageConfig(v MhVmStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *MhVmResources) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetParentReference

`func (o *MhVmResources) GetParentReference() Reference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MhVmResources) GetParentReferenceOk() (*Reference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentReference

`func (o *MhVmResources) SetParentReference(v Reference)`

SetParentReference sets ParentReference field to given value.

### HasParentReference

`func (o *MhVmResources) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### GetGuestTools

`func (o *MhVmResources) GetGuestTools() GuestToolsSpec`

GetGuestTools returns the GuestTools field if non-nil, zero value otherwise.

### GetGuestToolsOk

`func (o *MhVmResources) GetGuestToolsOk() (*GuestToolsSpec, bool)`

GetGuestToolsOk returns a tuple with the GuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestTools

`func (o *MhVmResources) SetGuestTools(v GuestToolsSpec)`

SetGuestTools sets GuestTools field to given value.

### HasGuestTools

`func (o *MhVmResources) HasGuestTools() bool`

HasGuestTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


