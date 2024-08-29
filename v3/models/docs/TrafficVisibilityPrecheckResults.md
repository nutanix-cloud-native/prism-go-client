# TrafficVisibilityPrecheckResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPassed** | Pointer to **bool** | IPFix export is capable or not on this cluster | [optional] 
**PrecheckResultList** | Pointer to [**[]TrafficVisibilityValidationResult**](TrafficVisibilityValidationResult.md) | Enumerates prechecks for this type and its status (pass or fail)  | [optional] 

## Methods

### NewTrafficVisibilityPrecheckResults

`func NewTrafficVisibilityPrecheckResults() *TrafficVisibilityPrecheckResults`

NewTrafficVisibilityPrecheckResults instantiates a new TrafficVisibilityPrecheckResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficVisibilityPrecheckResultsWithDefaults

`func NewTrafficVisibilityPrecheckResultsWithDefaults() *TrafficVisibilityPrecheckResults`

NewTrafficVisibilityPrecheckResultsWithDefaults instantiates a new TrafficVisibilityPrecheckResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPassed

`func (o *TrafficVisibilityPrecheckResults) GetHasPassed() bool`

GetHasPassed returns the HasPassed field if non-nil, zero value otherwise.

### GetHasPassedOk

`func (o *TrafficVisibilityPrecheckResults) GetHasPassedOk() (*bool, bool)`

GetHasPassedOk returns a tuple with the HasPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassed

`func (o *TrafficVisibilityPrecheckResults) SetHasPassed(v bool)`

SetHasPassed sets HasPassed field to given value.

### HasHasPassed

`func (o *TrafficVisibilityPrecheckResults) HasHasPassed() bool`

HasHasPassed returns a boolean if a field has been set.

### GetPrecheckResultList

`func (o *TrafficVisibilityPrecheckResults) GetPrecheckResultList() []TrafficVisibilityValidationResult`

GetPrecheckResultList returns the PrecheckResultList field if non-nil, zero value otherwise.

### GetPrecheckResultListOk

`func (o *TrafficVisibilityPrecheckResults) GetPrecheckResultListOk() (*[]TrafficVisibilityValidationResult, bool)`

GetPrecheckResultListOk returns a tuple with the PrecheckResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckResultList

`func (o *TrafficVisibilityPrecheckResults) SetPrecheckResultList(v []TrafficVisibilityValidationResult)`

SetPrecheckResultList sets PrecheckResultList field to given value.

### HasPrecheckResultList

`func (o *TrafficVisibilityPrecheckResults) HasPrecheckResultList() bool`

HasPrecheckResultList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


