# VmSetPowerStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUuid** | Pointer to **string** | The UUID of the power state task. | [optional] 

## Methods

### NewVmSetPowerStateResponse

`func NewVmSetPowerStateResponse() *VmSetPowerStateResponse`

NewVmSetPowerStateResponse instantiates a new VmSetPowerStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSetPowerStateResponseWithDefaults

`func NewVmSetPowerStateResponseWithDefaults() *VmSetPowerStateResponse`

NewVmSetPowerStateResponseWithDefaults instantiates a new VmSetPowerStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUuid

`func (o *VmSetPowerStateResponse) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *VmSetPowerStateResponse) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *VmSetPowerStateResponse) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *VmSetPowerStateResponse) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


