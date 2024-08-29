# StreamingPolicyRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If policy is enabled | [optional] [default to true]
**WindowDuration** | Pointer to **string** | Time window on which policy needs to be applied | [optional] [default to "FIVE_MINUTES"]
**ExecutionFrequency** | Pointer to **string** | How often policy needs to be evaluated | [optional] [default to "FIVE_MINUTES"]
**SqlQuery** | **string** | Policy as SQL string | 

## Methods

### NewStreamingPolicyRequestDetails

`func NewStreamingPolicyRequestDetails(sqlQuery string, ) *StreamingPolicyRequestDetails`

NewStreamingPolicyRequestDetails instantiates a new StreamingPolicyRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyRequestDetailsWithDefaults

`func NewStreamingPolicyRequestDetailsWithDefaults() *StreamingPolicyRequestDetails`

NewStreamingPolicyRequestDetailsWithDefaults instantiates a new StreamingPolicyRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *StreamingPolicyRequestDetails) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *StreamingPolicyRequestDetails) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *StreamingPolicyRequestDetails) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *StreamingPolicyRequestDetails) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetWindowDuration

`func (o *StreamingPolicyRequestDetails) GetWindowDuration() string`

GetWindowDuration returns the WindowDuration field if non-nil, zero value otherwise.

### GetWindowDurationOk

`func (o *StreamingPolicyRequestDetails) GetWindowDurationOk() (*string, bool)`

GetWindowDurationOk returns a tuple with the WindowDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDuration

`func (o *StreamingPolicyRequestDetails) SetWindowDuration(v string)`

SetWindowDuration sets WindowDuration field to given value.

### HasWindowDuration

`func (o *StreamingPolicyRequestDetails) HasWindowDuration() bool`

HasWindowDuration returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *StreamingPolicyRequestDetails) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *StreamingPolicyRequestDetails) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *StreamingPolicyRequestDetails) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *StreamingPolicyRequestDetails) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetSqlQuery

`func (o *StreamingPolicyRequestDetails) GetSqlQuery() string`

GetSqlQuery returns the SqlQuery field if non-nil, zero value otherwise.

### GetSqlQueryOk

`func (o *StreamingPolicyRequestDetails) GetSqlQueryOk() (*string, bool)`

GetSqlQueryOk returns a tuple with the SqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlQuery

`func (o *StreamingPolicyRequestDetails) SetSqlQuery(v string)`

SetSqlQuery sets SqlQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


