# EntityCountsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealizedCount** | Pointer to **bool** | Whether to compute counts of realized entities. | [optional] [default to false]
**ComplianceCount** | Pointer to **bool** | Whether to compute individual counts of entities which are COMPLIANT, NON_COMPLIANT and IN_PROGRESS.  | [optional] [default to false]
**UnrealizedCount** | Pointer to **bool** | Whether to compute counts of unrealized entities. | [optional] [default to false]

## Methods

### NewEntityCountsInput

`func NewEntityCountsInput() *EntityCountsInput`

NewEntityCountsInput instantiates a new EntityCountsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityCountsInputWithDefaults

`func NewEntityCountsInputWithDefaults() *EntityCountsInput`

NewEntityCountsInputWithDefaults instantiates a new EntityCountsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealizedCount

`func (o *EntityCountsInput) GetRealizedCount() bool`

GetRealizedCount returns the RealizedCount field if non-nil, zero value otherwise.

### GetRealizedCountOk

`func (o *EntityCountsInput) GetRealizedCountOk() (*bool, bool)`

GetRealizedCountOk returns a tuple with the RealizedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedCount

`func (o *EntityCountsInput) SetRealizedCount(v bool)`

SetRealizedCount sets RealizedCount field to given value.

### HasRealizedCount

`func (o *EntityCountsInput) HasRealizedCount() bool`

HasRealizedCount returns a boolean if a field has been set.

### GetComplianceCount

`func (o *EntityCountsInput) GetComplianceCount() bool`

GetComplianceCount returns the ComplianceCount field if non-nil, zero value otherwise.

### GetComplianceCountOk

`func (o *EntityCountsInput) GetComplianceCountOk() (*bool, bool)`

GetComplianceCountOk returns a tuple with the ComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCount

`func (o *EntityCountsInput) SetComplianceCount(v bool)`

SetComplianceCount sets ComplianceCount field to given value.

### HasComplianceCount

`func (o *EntityCountsInput) HasComplianceCount() bool`

HasComplianceCount returns a boolean if a field has been set.

### GetUnrealizedCount

`func (o *EntityCountsInput) GetUnrealizedCount() bool`

GetUnrealizedCount returns the UnrealizedCount field if non-nil, zero value otherwise.

### GetUnrealizedCountOk

`func (o *EntityCountsInput) GetUnrealizedCountOk() (*bool, bool)`

GetUnrealizedCountOk returns a tuple with the UnrealizedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedCount

`func (o *EntityCountsInput) SetUnrealizedCount(v bool)`

SetUnrealizedCount sets UnrealizedCount field to given value.

### HasUnrealizedCount

`func (o *EntityCountsInput) HasUnrealizedCount() bool`

HasUnrealizedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


