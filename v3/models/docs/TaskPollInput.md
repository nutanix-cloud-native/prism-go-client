# TaskPollInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PollTimeoutSeconds** | **int32** | Poll timeout in seconds | 
**TaskUuidList** | **[]string** | List of task UUIDs to poll on | 

## Methods

### NewTaskPollInput

`func NewTaskPollInput(pollTimeoutSeconds int32, taskUuidList []string, ) *TaskPollInput`

NewTaskPollInput instantiates a new TaskPollInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPollInputWithDefaults

`func NewTaskPollInputWithDefaults() *TaskPollInput`

NewTaskPollInputWithDefaults instantiates a new TaskPollInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPollTimeoutSeconds

`func (o *TaskPollInput) GetPollTimeoutSeconds() int32`

GetPollTimeoutSeconds returns the PollTimeoutSeconds field if non-nil, zero value otherwise.

### GetPollTimeoutSecondsOk

`func (o *TaskPollInput) GetPollTimeoutSecondsOk() (*int32, bool)`

GetPollTimeoutSecondsOk returns a tuple with the PollTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollTimeoutSeconds

`func (o *TaskPollInput) SetPollTimeoutSeconds(v int32)`

SetPollTimeoutSeconds sets PollTimeoutSeconds field to given value.


### GetTaskUuidList

`func (o *TaskPollInput) GetTaskUuidList() []string`

GetTaskUuidList returns the TaskUuidList field if non-nil, zero value otherwise.

### GetTaskUuidListOk

`func (o *TaskPollInput) GetTaskUuidListOk() (*[]string, bool)`

GetTaskUuidListOk returns a tuple with the TaskUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuidList

`func (o *TaskPollInput) SetTaskUuidList(v []string)`

SetTaskUuidList sets TaskUuidList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


