# TaskPollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**HasPollTimedOut** | Pointer to **bool** | Did the poll timeout before all polled tasks completed | [optional] 

## Methods

### NewTaskPollResponse

`func NewTaskPollResponse() *TaskPollResponse`

NewTaskPollResponse instantiates a new TaskPollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPollResponseWithDefaults

`func NewTaskPollResponseWithDefaults() *TaskPollResponse`

NewTaskPollResponseWithDefaults instantiates a new TaskPollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *TaskPollResponse) GetEntities() []Task`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TaskPollResponse) GetEntitiesOk() (*[]Task, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *TaskPollResponse) SetEntities(v []Task)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *TaskPollResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetHasPollTimedOut

`func (o *TaskPollResponse) GetHasPollTimedOut() bool`

GetHasPollTimedOut returns the HasPollTimedOut field if non-nil, zero value otherwise.

### GetHasPollTimedOutOk

`func (o *TaskPollResponse) GetHasPollTimedOutOk() (*bool, bool)`

GetHasPollTimedOutOk returns a tuple with the HasPollTimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPollTimedOut

`func (o *TaskPollResponse) SetHasPollTimedOut(v bool)`

SetHasPollTimedOut sets HasPollTimedOut field to given value.

### HasHasPollTimedOut

`func (o *TaskPollResponse) HasHasPollTimedOut() bool`

HasHasPollTimedOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


