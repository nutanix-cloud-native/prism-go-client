# VirtualServerWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumVms** | Pointer to **int32** |  | [optional] 
**ServerProfileType** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualServerWorkload

`func NewVirtualServerWorkload() *VirtualServerWorkload`

NewVirtualServerWorkload instantiates a new VirtualServerWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualServerWorkloadWithDefaults

`func NewVirtualServerWorkloadWithDefaults() *VirtualServerWorkload`

NewVirtualServerWorkloadWithDefaults instantiates a new VirtualServerWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumVms

`func (o *VirtualServerWorkload) GetNumVms() int32`

GetNumVms returns the NumVms field if non-nil, zero value otherwise.

### GetNumVmsOk

`func (o *VirtualServerWorkload) GetNumVmsOk() (*int32, bool)`

GetNumVmsOk returns a tuple with the NumVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVms

`func (o *VirtualServerWorkload) SetNumVms(v int32)`

SetNumVms sets NumVms field to given value.

### HasNumVms

`func (o *VirtualServerWorkload) HasNumVms() bool`

HasNumVms returns a boolean if a field has been set.

### GetServerProfileType

`func (o *VirtualServerWorkload) GetServerProfileType() string`

GetServerProfileType returns the ServerProfileType field if non-nil, zero value otherwise.

### GetServerProfileTypeOk

`func (o *VirtualServerWorkload) GetServerProfileTypeOk() (*string, bool)`

GetServerProfileTypeOk returns a tuple with the ServerProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProfileType

`func (o *VirtualServerWorkload) SetServerProfileType(v string)`

SetServerProfileType sets ServerProfileType field to given value.

### HasServerProfileType

`func (o *VirtualServerWorkload) HasServerProfileType() bool`

HasServerProfileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


