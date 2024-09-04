# SecurityPlanningPrechecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasPassed** | Pointer to **bool** | Whether this type of prechecks passed | [optional] 
**PrecheckResultList** | Pointer to [**[]ValidationResult**](ValidationResult.md) | Enumerates prechecks for this type and it&#39;s status (pass or fail)  | [optional] 

## Methods

### NewSecurityPlanningPrechecks

`func NewSecurityPlanningPrechecks() *SecurityPlanningPrechecks`

NewSecurityPlanningPrechecks instantiates a new SecurityPlanningPrechecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPlanningPrechecksWithDefaults

`func NewSecurityPlanningPrechecksWithDefaults() *SecurityPlanningPrechecks`

NewSecurityPlanningPrechecksWithDefaults instantiates a new SecurityPlanningPrechecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasPassed

`func (o *SecurityPlanningPrechecks) GetHasPassed() bool`

GetHasPassed returns the HasPassed field if non-nil, zero value otherwise.

### GetHasPassedOk

`func (o *SecurityPlanningPrechecks) GetHasPassedOk() (*bool, bool)`

GetHasPassedOk returns a tuple with the HasPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassed

`func (o *SecurityPlanningPrechecks) SetHasPassed(v bool)`

SetHasPassed sets HasPassed field to given value.

### HasHasPassed

`func (o *SecurityPlanningPrechecks) HasHasPassed() bool`

HasHasPassed returns a boolean if a field has been set.

### GetPrecheckResultList

`func (o *SecurityPlanningPrechecks) GetPrecheckResultList() []ValidationResult`

GetPrecheckResultList returns the PrecheckResultList field if non-nil, zero value otherwise.

### GetPrecheckResultListOk

`func (o *SecurityPlanningPrechecks) GetPrecheckResultListOk() (*[]ValidationResult, bool)`

GetPrecheckResultListOk returns a tuple with the PrecheckResultList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckResultList

`func (o *SecurityPlanningPrechecks) SetPrecheckResultList(v []ValidationResult)`

SetPrecheckResultList sets PrecheckResultList field to given value.

### HasPrecheckResultList

`func (o *SecurityPlanningPrechecks) HasPrecheckResultList() bool`

HasPrecheckResultList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


