# VmRestoreOverrideSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | vm Name. | [optional] 
**Resources** | Pointer to [**VmResources**](VmResources.md) |  | [optional] 

## Methods

### NewVmRestoreOverrideSpec

`func NewVmRestoreOverrideSpec() *VmRestoreOverrideSpec`

NewVmRestoreOverrideSpec instantiates a new VmRestoreOverrideSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmRestoreOverrideSpecWithDefaults

`func NewVmRestoreOverrideSpecWithDefaults() *VmRestoreOverrideSpec`

NewVmRestoreOverrideSpecWithDefaults instantiates a new VmRestoreOverrideSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmRestoreOverrideSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmRestoreOverrideSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmRestoreOverrideSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmRestoreOverrideSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResources

`func (o *VmRestoreOverrideSpec) GetResources() VmResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VmRestoreOverrideSpec) GetResourcesOk() (*VmResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VmRestoreOverrideSpec) SetResources(v VmResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *VmRestoreOverrideSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


