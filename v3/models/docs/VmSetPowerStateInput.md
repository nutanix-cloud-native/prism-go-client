# VmSetPowerStateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUuid** | Pointer to **string** | The UUID of the task (used for idempotency). | [optional] 

## Methods

### NewVmSetPowerStateInput

`func NewVmSetPowerStateInput() *VmSetPowerStateInput`

NewVmSetPowerStateInput instantiates a new VmSetPowerStateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmSetPowerStateInputWithDefaults

`func NewVmSetPowerStateInputWithDefaults() *VmSetPowerStateInput`

NewVmSetPowerStateInputWithDefaults instantiates a new VmSetPowerStateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUuid

`func (o *VmSetPowerStateInput) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *VmSetPowerStateInput) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *VmSetPowerStateInput) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *VmSetPowerStateInput) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


