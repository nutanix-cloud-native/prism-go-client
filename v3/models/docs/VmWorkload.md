# VmWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumVms** | Pointer to **int32** |  | [optional] 
**ResourceSpec** | Pointer to [**VmResourceSpec**](VmResourceSpec.md) |  | [optional] 

## Methods

### NewVmWorkload

`func NewVmWorkload() *VmWorkload`

NewVmWorkload instantiates a new VmWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmWorkloadWithDefaults

`func NewVmWorkloadWithDefaults() *VmWorkload`

NewVmWorkloadWithDefaults instantiates a new VmWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumVms

`func (o *VmWorkload) GetNumVms() int32`

GetNumVms returns the NumVms field if non-nil, zero value otherwise.

### GetNumVmsOk

`func (o *VmWorkload) GetNumVmsOk() (*int32, bool)`

GetNumVmsOk returns a tuple with the NumVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVms

`func (o *VmWorkload) SetNumVms(v int32)`

SetNumVms sets NumVms field to given value.

### HasNumVms

`func (o *VmWorkload) HasNumVms() bool`

HasNumVms returns a boolean if a field has been set.

### GetResourceSpec

`func (o *VmWorkload) GetResourceSpec() VmResourceSpec`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *VmWorkload) GetResourceSpecOk() (*VmResourceSpec, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *VmWorkload) SetResourceSpec(v VmResourceSpec)`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *VmWorkload) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


