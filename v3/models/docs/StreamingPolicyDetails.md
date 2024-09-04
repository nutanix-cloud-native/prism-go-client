# StreamingPolicyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If policy is enabled | [optional] [default to true]
**WindowDuration** | Pointer to **string** | Time window on which policy needs to be applied | [optional] [default to "FIVE_MINUTES"]
**ExecutionFrequency** | Pointer to **string** | How often policy needs to be evaluated | [optional] [default to "FIVE_MINUTES"]
**SqlQuery** | **string** | Policy as SQL string | 

## Methods

### NewStreamingPolicyDetails

`func NewStreamingPolicyDetails(sqlQuery string, ) *StreamingPolicyDetails`

NewStreamingPolicyDetails instantiates a new StreamingPolicyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyDetailsWithDefaults

`func NewStreamingPolicyDetailsWithDefaults() *StreamingPolicyDetails`

NewStreamingPolicyDetailsWithDefaults instantiates a new StreamingPolicyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *StreamingPolicyDetails) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *StreamingPolicyDetails) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *StreamingPolicyDetails) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *StreamingPolicyDetails) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetWindowDuration

`func (o *StreamingPolicyDetails) GetWindowDuration() string`

GetWindowDuration returns the WindowDuration field if non-nil, zero value otherwise.

### GetWindowDurationOk

`func (o *StreamingPolicyDetails) GetWindowDurationOk() (*string, bool)`

GetWindowDurationOk returns a tuple with the WindowDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDuration

`func (o *StreamingPolicyDetails) SetWindowDuration(v string)`

SetWindowDuration sets WindowDuration field to given value.

### HasWindowDuration

`func (o *StreamingPolicyDetails) HasWindowDuration() bool`

HasWindowDuration returns a boolean if a field has been set.

### GetExecutionFrequency

`func (o *StreamingPolicyDetails) GetExecutionFrequency() string`

GetExecutionFrequency returns the ExecutionFrequency field if non-nil, zero value otherwise.

### GetExecutionFrequencyOk

`func (o *StreamingPolicyDetails) GetExecutionFrequencyOk() (*string, bool)`

GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionFrequency

`func (o *StreamingPolicyDetails) SetExecutionFrequency(v string)`

SetExecutionFrequency sets ExecutionFrequency field to given value.

### HasExecutionFrequency

`func (o *StreamingPolicyDetails) HasExecutionFrequency() bool`

HasExecutionFrequency returns a boolean if a field has been set.

### GetSqlQuery

`func (o *StreamingPolicyDetails) GetSqlQuery() string`

GetSqlQuery returns the SqlQuery field if non-nil, zero value otherwise.

### GetSqlQueryOk

`func (o *StreamingPolicyDetails) GetSqlQueryOk() (*string, bool)`

GetSqlQueryOk returns a tuple with the SqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlQuery

`func (o *StreamingPolicyDetails) SetSqlQuery(v string)`

SetSqlQuery sets SqlQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


