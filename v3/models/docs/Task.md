# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current state of the task. | [optional] 
**LastUpdateTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when task was last updated.  | [optional] 
**ErrorDetail** | Pointer to **string** | In case of task failure this field will provide the error description.  | [optional] 
**LogicalTimestamp** | Pointer to **int64** | Number of times the task has been updated. The value increases sequentially with each update of the task and can be used to verify if there have been changes to the task.  | [optional] 
**RequestedStatus** | Pointer to **string** | Final expected state of the task. It is set when the task is aborted.  | [optional] 
**EntityReferenceList** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when Task execution started.  | [optional] 
**CreationTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when task was created.  | [optional] 
**UUID** | Pointer to **string** | UUID of the task. | [optional] 
**StartTimeUsecs** | Pointer to **int64** | Time in microseconds from epoch when the task execution started.  | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**SubtaskReferenceList** | Pointer to [**[]TaskReference**](TaskReference.md) | Reference to the sub-tasks. | [optional] 
**CompletionTime** | Pointer to **time.Time** | UTC date and time in RFC-3339 format when Task execution completed.  | [optional] 
**CreationTimeUsecs** | Pointer to **int64** | Time in microseconds from epoch when task was created.  | [optional] 
**ProgressMessage** | Pointer to **string** | Description of what currently the task is doing. | [optional] 
**OperationType** | Pointer to **string** | Type of the operation tracked by the task. | [optional] 
**CompletionTimeUsecs** | Pointer to **int64** | Time in microseconds from epoch when task execution completed.  | [optional] 
**ErrorCode** | Pointer to **string** | In case of task failure this field will provide the error code.  | [optional] 
**PercentageComplete** | Pointer to **int32** | The completion percentage for the task. | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**ParentTaskReference** | Pointer to [**TaskReference**](TaskReference.md) |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *Task) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *Task) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *Task) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *Task) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetErrorDetail

`func (o *Task) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *Task) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *Task) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *Task) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetLogicalTimestamp

`func (o *Task) GetLogicalTimestamp() int64`

GetLogicalTimestamp returns the LogicalTimestamp field if non-nil, zero value otherwise.

### GetLogicalTimestampOk

`func (o *Task) GetLogicalTimestampOk() (*int64, bool)`

GetLogicalTimestampOk returns a tuple with the LogicalTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalTimestamp

`func (o *Task) SetLogicalTimestamp(v int64)`

SetLogicalTimestamp sets LogicalTimestamp field to given value.

### HasLogicalTimestamp

`func (o *Task) HasLogicalTimestamp() bool`

HasLogicalTimestamp returns a boolean if a field has been set.

### GetRequestedStatus

`func (o *Task) GetRequestedStatus() string`

GetRequestedStatus returns the RequestedStatus field if non-nil, zero value otherwise.

### GetRequestedStatusOk

`func (o *Task) GetRequestedStatusOk() (*string, bool)`

GetRequestedStatusOk returns a tuple with the RequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedStatus

`func (o *Task) SetRequestedStatus(v string)`

SetRequestedStatus sets RequestedStatus field to given value.

### HasRequestedStatus

`func (o *Task) HasRequestedStatus() bool`

HasRequestedStatus returns a boolean if a field has been set.

### GetEntityReferenceList

`func (o *Task) GetEntityReferenceList() []Reference`

GetEntityReferenceList returns the EntityReferenceList field if non-nil, zero value otherwise.

### GetEntityReferenceListOk

`func (o *Task) GetEntityReferenceListOk() (*[]Reference, bool)`

GetEntityReferenceListOk returns a tuple with the EntityReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityReferenceList

`func (o *Task) SetEntityReferenceList(v []Reference)`

SetEntityReferenceList sets EntityReferenceList field to given value.

### HasEntityReferenceList

`func (o *Task) HasEntityReferenceList() bool`

HasEntityReferenceList returns a boolean if a field has been set.

### GetStartTime

`func (o *Task) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Task) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Task) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Task) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *Task) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Task) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Task) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Task) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUUID

`func (o *Task) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Task) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Task) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *Task) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetStartTimeUsecs

`func (o *Task) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *Task) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *Task) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *Task) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### GetClusterReference

`func (o *Task) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *Task) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *Task) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *Task) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetSubtaskReferenceList

`func (o *Task) GetSubtaskReferenceList() []TaskReference`

GetSubtaskReferenceList returns the SubtaskReferenceList field if non-nil, zero value otherwise.

### GetSubtaskReferenceListOk

`func (o *Task) GetSubtaskReferenceListOk() (*[]TaskReference, bool)`

GetSubtaskReferenceListOk returns a tuple with the SubtaskReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtaskReferenceList

`func (o *Task) SetSubtaskReferenceList(v []TaskReference)`

SetSubtaskReferenceList sets SubtaskReferenceList field to given value.

### HasSubtaskReferenceList

`func (o *Task) HasSubtaskReferenceList() bool`

HasSubtaskReferenceList returns a boolean if a field has been set.

### GetCompletionTime

`func (o *Task) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *Task) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *Task) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *Task) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetCreationTimeUsecs

`func (o *Task) GetCreationTimeUsecs() int64`

GetCreationTimeUsecs returns the CreationTimeUsecs field if non-nil, zero value otherwise.

### GetCreationTimeUsecsOk

`func (o *Task) GetCreationTimeUsecsOk() (*int64, bool)`

GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeUsecs

`func (o *Task) SetCreationTimeUsecs(v int64)`

SetCreationTimeUsecs sets CreationTimeUsecs field to given value.

### HasCreationTimeUsecs

`func (o *Task) HasCreationTimeUsecs() bool`

HasCreationTimeUsecs returns a boolean if a field has been set.

### GetProgressMessage

`func (o *Task) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *Task) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *Task) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *Task) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetOperationType

`func (o *Task) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *Task) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *Task) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *Task) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetCompletionTimeUsecs

`func (o *Task) GetCompletionTimeUsecs() int64`

GetCompletionTimeUsecs returns the CompletionTimeUsecs field if non-nil, zero value otherwise.

### GetCompletionTimeUsecsOk

`func (o *Task) GetCompletionTimeUsecsOk() (*int64, bool)`

GetCompletionTimeUsecsOk returns a tuple with the CompletionTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimeUsecs

`func (o *Task) SetCompletionTimeUsecs(v int64)`

SetCompletionTimeUsecs sets CompletionTimeUsecs field to given value.

### HasCompletionTimeUsecs

`func (o *Task) HasCompletionTimeUsecs() bool`

HasCompletionTimeUsecs returns a boolean if a field has been set.

### GetErrorCode

`func (o *Task) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Task) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Task) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Task) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *Task) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *Task) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *Task) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.

### HasPercentageComplete

`func (o *Task) HasPercentageComplete() bool`

HasPercentageComplete returns a boolean if a field has been set.

### GetApiVersion

`func (o *Task) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Task) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Task) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Task) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetParentTaskReference

`func (o *Task) GetParentTaskReference() TaskReference`

GetParentTaskReference returns the ParentTaskReference field if non-nil, zero value otherwise.

### GetParentTaskReferenceOk

`func (o *Task) GetParentTaskReferenceOk() (*TaskReference, bool)`

GetParentTaskReferenceOk returns a tuple with the ParentTaskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskReference

`func (o *Task) SetParentTaskReference(v TaskReference)`

SetParentTaskReference sets ParentTaskReference field to given value.

### HasParentTaskReference

`func (o *Task) HasParentTaskReference() bool`

HasParentTaskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


