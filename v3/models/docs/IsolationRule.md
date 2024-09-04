# IsolationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of action. | [optional] 
**FirstEntityFilter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 
**SecondEntityFilter** | Pointer to [**CategoryFilter**](CategoryFilter.md) |  | [optional] 

## Methods

### NewIsolationRule

`func NewIsolationRule() *IsolationRule`

NewIsolationRule instantiates a new IsolationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolationRuleWithDefaults

`func NewIsolationRuleWithDefaults() *IsolationRule`

NewIsolationRuleWithDefaults instantiates a new IsolationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IsolationRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IsolationRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IsolationRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IsolationRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFirstEntityFilter

`func (o *IsolationRule) GetFirstEntityFilter() CategoryFilter`

GetFirstEntityFilter returns the FirstEntityFilter field if non-nil, zero value otherwise.

### GetFirstEntityFilterOk

`func (o *IsolationRule) GetFirstEntityFilterOk() (*CategoryFilter, bool)`

GetFirstEntityFilterOk returns a tuple with the FirstEntityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEntityFilter

`func (o *IsolationRule) SetFirstEntityFilter(v CategoryFilter)`

SetFirstEntityFilter sets FirstEntityFilter field to given value.

### HasFirstEntityFilter

`func (o *IsolationRule) HasFirstEntityFilter() bool`

HasFirstEntityFilter returns a boolean if a field has been set.

### GetSecondEntityFilter

`func (o *IsolationRule) GetSecondEntityFilter() CategoryFilter`

GetSecondEntityFilter returns the SecondEntityFilter field if non-nil, zero value otherwise.

### GetSecondEntityFilterOk

`func (o *IsolationRule) GetSecondEntityFilterOk() (*CategoryFilter, bool)`

GetSecondEntityFilterOk returns a tuple with the SecondEntityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondEntityFilter

`func (o *IsolationRule) SetSecondEntityFilter(v CategoryFilter)`

SetSecondEntityFilter sets SecondEntityFilter field to given value.

### HasSecondEntityFilter

`func (o *IsolationRule) HasSecondEntityFilter() bool`

HasSecondEntityFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


