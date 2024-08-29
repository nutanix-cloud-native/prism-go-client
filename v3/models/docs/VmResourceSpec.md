# VmResourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RamGb** | Pointer to **float32** |  | [optional] 
**HddGb** | Pointer to **float32** |  | [optional] 
**NumVcpus** | Pointer to **int32** |  | [optional] 

## Methods

### NewVmResourceSpec

`func NewVmResourceSpec() *VmResourceSpec`

NewVmResourceSpec instantiates a new VmResourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmResourceSpecWithDefaults

`func NewVmResourceSpecWithDefaults() *VmResourceSpec`

NewVmResourceSpecWithDefaults instantiates a new VmResourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRamGb

`func (o *VmResourceSpec) GetRamGb() float32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *VmResourceSpec) GetRamGbOk() (*float32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *VmResourceSpec) SetRamGb(v float32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *VmResourceSpec) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetHddGb

`func (o *VmResourceSpec) GetHddGb() float32`

GetHddGb returns the HddGb field if non-nil, zero value otherwise.

### GetHddGbOk

`func (o *VmResourceSpec) GetHddGbOk() (*float32, bool)`

GetHddGbOk returns a tuple with the HddGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddGb

`func (o *VmResourceSpec) SetHddGb(v float32)`

SetHddGb sets HddGb field to given value.

### HasHddGb

`func (o *VmResourceSpec) HasHddGb() bool`

HasHddGb returns a boolean if a field has been set.

### GetNumVcpus

`func (o *VmResourceSpec) GetNumVcpus() int32`

GetNumVcpus returns the NumVcpus field if non-nil, zero value otherwise.

### GetNumVcpusOk

`func (o *VmResourceSpec) GetNumVcpusOk() (*int32, bool)`

GetNumVcpusOk returns a tuple with the NumVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVcpus

`func (o *VmResourceSpec) SetNumVcpus(v int32)`

SetNumVcpus sets NumVcpus field to given value.

### HasNumVcpus

`func (o *VmResourceSpec) HasNumVcpus() bool`

HasNumVcpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


