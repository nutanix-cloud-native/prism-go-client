# RebootNgtPolicyParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyOnce** | Pointer to **bool** | Flag for policies to be applied only once. | [optional] [default to false]
**Schedule** | Pointer to [**RebootNgtPolicyParametersSchedule**](RebootNgtPolicyParametersSchedule.md) |  | [optional] 
**ScheduleType** | **string** | Type of Schedule for policy enforcement. | 

## Methods

### NewRebootNgtPolicyParameters

`func NewRebootNgtPolicyParameters(scheduleType string, ) *RebootNgtPolicyParameters`

NewRebootNgtPolicyParameters instantiates a new RebootNgtPolicyParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootNgtPolicyParametersWithDefaults

`func NewRebootNgtPolicyParametersWithDefaults() *RebootNgtPolicyParameters`

NewRebootNgtPolicyParametersWithDefaults instantiates a new RebootNgtPolicyParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyOnce

`func (o *RebootNgtPolicyParameters) GetApplyOnce() bool`

GetApplyOnce returns the ApplyOnce field if non-nil, zero value otherwise.

### GetApplyOnceOk

`func (o *RebootNgtPolicyParameters) GetApplyOnceOk() (*bool, bool)`

GetApplyOnceOk returns a tuple with the ApplyOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyOnce

`func (o *RebootNgtPolicyParameters) SetApplyOnce(v bool)`

SetApplyOnce sets ApplyOnce field to given value.

### HasApplyOnce

`func (o *RebootNgtPolicyParameters) HasApplyOnce() bool`

HasApplyOnce returns a boolean if a field has been set.

### GetSchedule

`func (o *RebootNgtPolicyParameters) GetSchedule() RebootNgtPolicyParametersSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RebootNgtPolicyParameters) GetScheduleOk() (*RebootNgtPolicyParametersSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RebootNgtPolicyParameters) SetSchedule(v RebootNgtPolicyParametersSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RebootNgtPolicyParameters) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleType

`func (o *RebootNgtPolicyParameters) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *RebootNgtPolicyParameters) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *RebootNgtPolicyParameters) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


